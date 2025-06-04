package model

import (
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

	return user
}

func FromProtoToGetUsersByPreferencesRequest(req *gen.SelectUsersByPreferencesRequest) *GetUsersByPreferencesRequest {
	return &GetUsersByPreferencesRequest{
		UserId:          req.UserId,
		PreferredGender: req.PreferredGender,
		AgeMin:          int(req.AgeMin),
		AgeMax:          int(req.AgeMax),
		Position: &Position{
			Lat:    float64(req.Position.Lat),
			Lon:    float64(req.Position.Lon),
			Radius: int(req.Position.Radius),
		},
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
	}
}

func FromUserPreferencesRequestToProto(req *GetUsersByPreferencesRequest) *gen.SelectUsersByPreferencesRequest {
	return &gen.SelectUsersByPreferencesRequest{
		UserId:          req.UserId,
		PreferredGender: req.PreferredGender,
		AgeMin:          int32(req.AgeMin),
		AgeMax:          int32(req.AgeMax),
		Position: &gen.UserPosition{
			Lon:    float32(req.Position.Lon),
			Lat:    float32(req.Position.Lat),
			Radius: int32(req.Position.Radius),
		},
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
	}
}
