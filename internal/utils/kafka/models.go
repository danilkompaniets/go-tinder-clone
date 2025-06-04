package kafka

import "github.com/danilkompanites/tinder-clone/services/users/pkg/model"

type EventType string

var Events = struct {
	UserCreated            EventType
	UserUpdated            EventType
	UserPreferencesUpdated EventType
}{
	UserCreated:            "user.created",
	UserUpdated:            "user.updated",
	UserPreferencesUpdated: "user.preferences.updated",
}

type Topic string

var Topics = struct {
	User      Topic
	Auth      Topic
	Order     Topic
	Inventory Topic
}{
	User: "user-topic",
}

type PreferencesUpdate struct {
	ID              string          `json:"id"`
	PreferredGender *string         `json:"preferred_gender" binding:"required"`
	AgeMin          *int            `json:"age_min" binding:"required"`
	AgeMax          *int            `json:"age_max" binding:"required"`
	Position        *model.Position `json:"position" binding:"required"`
}

type ConsumerGroup string

var ConsumerGroups = struct {
	UserService ConsumerGroup
}{
	UserService: "user-service",
}
