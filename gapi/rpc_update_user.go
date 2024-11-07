package gapi

import (
	"context"
	"errors"

	db "blog_api/db/sqlc"
	"blog_api/pb"
	"blog_api/val"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if authPayload.UserId != req.GetUserId() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's info")
	}

	arg := db.UpdateUserParams{
		UserID: req.GetUserId(),
		FirstName: pgtype.Text{
			String: req.GetFirstName(),
			Valid:  req.FirstName != nil,
		},
		LastName: pgtype.Text{
			String: req.GetLastName(),
			Valid:  req.FirstName != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Phone: pgtype.Text{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.UserId == 0 {
		violations = append(violations, fieldViolation("user_id", errors.New("user_id is required")))
	}

	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.Phone != nil {
		if err := val.ValidatePhone(req.GetPhone()); err != nil {
			violations = append(violations, fieldViolation("phone", err))
		}
	}

	return violations
}
