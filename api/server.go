package api

import (
	"fmt"
	"net/http"

	db "blog_api/db/sqlc"
	"blog_api/token"
	"blog_api/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonResponseWithPaginate struct {
	jsonResponse
	Total int64 `json:"total"`
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

// Middleware to check the size of the request body
func LimitRequestBodySize(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		if err := c.Request.ParseForm(); err != nil {
			c.String(http.StatusRequestEntityTooLarge, "Request body too large")
			c.Abort()
			return
		}
		c.Next()
	}
}

func (server *Server) setupRouter() {

	// Load Config Environment
	config_env, _ := util.LoadConfig(".")

	// Add CORS middleware
	// Configure CORS settings
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{config_env.URL_FRONTEND}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	// Set Gin mode
	gin.SetMode(config_env.GIN_MODE)

	// Initialize Gin router
	router := gin.Default()

	// Apply middleware to limit request body size
	router.Use(LimitRequestBodySize(50 * 1024 * 1024))

	// Apply CORS middleware
	router.Use(cors.New(config))

	// Create a router group
	routerGroup := router.Group("/")

	// version
	routerGroup.GET("/api/version", server.getVersion)

	// Auth
	routerGroup.POST("/api/login", server.loginUser)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": true, "message": err.Error(), "data": nil}
}

func successResponse(data interface{}) gin.H {
	return gin.H{"error": false, "message": "successfully", "data": data}
}
