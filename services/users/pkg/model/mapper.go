package model

import (
	"fmt"
	"github.com/danilkompanites/tinder-clone/gen"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

func FromProtoToUser(request *gen.User) *User {
	user := &User{
		ID:        request.Id,
		Username:  request.Username,
		Email:     request.Email,
		FirstName: &request.FirstName,
		Bio:       &request.Bio,
		Gender:    &request.Gender,
		BirthDate: request.BirthDate.AsTime(),
		City:      request.City,
		AvatarURL: request.AvatarUrl,
		CreatedAt: request.CreatedAt.AsTime(),
		UpdatedAt: request.UpdatedAt.AsTime(),
	}

	fmt.Println(user)

	return user
}
