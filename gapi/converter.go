package gapi

import (
	db "blog_api/db/sqlc"
	"blog_api/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		Phone:             user.Phone,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
