package gapi

import (
	"context"
	"errors"

	db "blog_api/db/sqlc"
	"blog_api/pb"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateListUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListUsersParams{
		Limit:  int32(req.GetPageSize()),
		Offset: int32((req.GetPageId() - 1) * req.GetPageSize()),
	}

	var users_response []*pb.User

	user, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users")
	}

	countUser, err := server.store.CountUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to count users")
	}

	for _, u := range user {
		roles, err := server.store.GetRoleByUserId(ctx, u.ID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get role by user id")
		}

		newUserResponse := &pb.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
			Roles:     convertRoles(roles),
		}

		users_response = append(users_response, newUserResponse)
	}

	rsp := &pb.ListUserResponse{
		User:  users_response,
		Total: countUser,
	}
	return rsp, nil
}

func validateListUserRequest(req *pb.ListUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetPageId() < 0 {
		violations = append(violations, fieldViolation("page_id", errors.New("must be at least 0")))
	}

	if req.GetPageSize() < 0 {
		violations = append(violations, fieldViolation("page_size", errors.New("must be at least 0")))
	}

	return violations
}

func convertRoles(dbRoles []db.GetRoleByUserIdRow) []*pb.Role {
	var pbRoles []*pb.Role
	for _, dbRole := range dbRoles {
		pbRole := &pb.Role{
			Id:   dbRole.ID,
			Name: dbRole.Name,
		}
		pbRoles = append(pbRoles, pbRole)
	}
	return pbRoles
}
