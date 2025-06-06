// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/api.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FirstName     string                 `protobuf:"bytes,4,opt,name=firstName,proto3" json:"firstName,omitempty"`
	Bio           string                 `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	Gender        string                 `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthDate     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=birthDate,proto3" json:"birthDate,omitempty"`
	City          string                 `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	AvatarUrl     string                 `protobuf:"bytes,9,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_api_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_api_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_api_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

type UserPosition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lon           float32                `protobuf:"fixed32,1,opt,name=lon,proto3" json:"lon,omitempty"`
	Lat           float32                `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Radius        int32                  `protobuf:"varint,3,opt,name=radius,proto3" json:"radius,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserPosition) Reset() {
	*x = UserPosition{}
	mi := &file_api_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPosition) ProtoMessage() {}

func (x *UserPosition) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPosition.ProtoReflect.Descriptor instead.
func (*UserPosition) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *UserPosition) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *UserPosition) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *UserPosition) GetRadius() int32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type Preferences struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PreferredGender string                 `protobuf:"bytes,3,opt,name=preferredGender,proto3" json:"preferredGender,omitempty"`
	AgeMin          int32                  `protobuf:"varint,4,opt,name=ageMin,proto3" json:"ageMin,omitempty"`
	AgeMax          int32                  `protobuf:"varint,5,opt,name=ageMax,proto3" json:"ageMax,omitempty"`
	Position        *UserPosition          `protobuf:"bytes,6,opt,name=position,proto3" json:"position,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Preferences) Reset() {
	*x = Preferences{}
	mi := &file_api_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Preferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preferences) ProtoMessage() {}

func (x *Preferences) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preferences.ProtoReflect.Descriptor instead.
func (*Preferences) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *Preferences) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Preferences) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Preferences) GetPreferredGender() string {
	if x != nil {
		return x.PreferredGender
	}
	return ""
}

func (x *Preferences) GetAgeMin() int32 {
	if x != nil {
		return x.AgeMin
	}
	return 0
}

func (x *Preferences) GetAgeMax() int32 {
	if x != nil {
		return x.AgeMax
	}
	return 0
}

func (x *Preferences) GetPosition() *UserPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

type SelectUsersByPreferencesRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PreferredGender string                 `protobuf:"bytes,2,opt,name=preferredGender,proto3" json:"preferredGender,omitempty"`
	AgeMin          int32                  `protobuf:"varint,3,opt,name=ageMin,proto3" json:"ageMin,omitempty"`
	AgeMax          int32                  `protobuf:"varint,4,opt,name=ageMax,proto3" json:"ageMax,omitempty"`
	Position        *UserPosition          `protobuf:"bytes,5,opt,name=position,proto3" json:"position,omitempty"`
	Limit           int32                  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset          int32                  `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SelectUsersByPreferencesRequest) Reset() {
	*x = SelectUsersByPreferencesRequest{}
	mi := &file_api_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SelectUsersByPreferencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectUsersByPreferencesRequest) ProtoMessage() {}

func (x *SelectUsersByPreferencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectUsersByPreferencesRequest.ProtoReflect.Descriptor instead.
func (*SelectUsersByPreferencesRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *SelectUsersByPreferencesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SelectUsersByPreferencesRequest) GetPreferredGender() string {
	if x != nil {
		return x.PreferredGender
	}
	return ""
}

func (x *SelectUsersByPreferencesRequest) GetAgeMin() int32 {
	if x != nil {
		return x.AgeMin
	}
	return 0
}

func (x *SelectUsersByPreferencesRequest) GetAgeMax() int32 {
	if x != nil {
		return x.AgeMax
	}
	return 0
}

func (x *SelectUsersByPreferencesRequest) GetPosition() *UserPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *SelectUsersByPreferencesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SelectUsersByPreferencesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type SelectUsersByPreferencesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SelectUsersByPreferencesResponse) Reset() {
	*x = SelectUsersByPreferencesResponse{}
	mi := &file_api_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SelectUsersByPreferencesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectUsersByPreferencesResponse) ProtoMessage() {}

func (x *SelectUsersByPreferencesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectUsersByPreferencesResponse.ProtoReflect.Descriptor instead.
func (*SelectUsersByPreferencesResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *SelectUsersByPreferencesResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_api_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{7}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	mi := &file_api_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{8}
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type MatchModel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromId        string                 `protobuf:"bytes,1,opt,name=fromId,proto3" json:"fromId,omitempty"`
	ToId          string                 `protobuf:"bytes,2,opt,name=toId,proto3" json:"toId,omitempty"`
	FromDecision  bool                   `protobuf:"varint,3,opt,name=fromDecision,proto3" json:"fromDecision,omitempty"`
	ToDecision    bool                   `protobuf:"varint,4,opt,name=toDecision,proto3" json:"toDecision,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchModel) Reset() {
	*x = MatchModel{}
	mi := &file_api_api_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchModel) ProtoMessage() {}

func (x *MatchModel) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchModel.ProtoReflect.Descriptor instead.
func (*MatchModel) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{9}
}

func (x *MatchModel) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *MatchModel) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *MatchModel) GetFromDecision() bool {
	if x != nil {
		return x.FromDecision
	}
	return false
}

func (x *MatchModel) GetToDecision() bool {
	if x != nil {
		return x.ToDecision
	}
	return false
}

type GetDecisionsUserIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDecisionsUserIdRequest) Reset() {
	*x = GetDecisionsUserIdRequest{}
	mi := &file_api_api_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDecisionsUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecisionsUserIdRequest) ProtoMessage() {}

func (x *GetDecisionsUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecisionsUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetDecisionsUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{10}
}

func (x *GetDecisionsUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetDecisionsUserIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserIds       []string               `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDecisionsUserIdResponse) Reset() {
	*x = GetDecisionsUserIdResponse{}
	mi := &file_api_api_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDecisionsUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecisionsUserIdResponse) ProtoMessage() {}

func (x *GetDecisionsUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecisionsUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetDecisionsUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{11}
}

func (x *GetDecisionsUserIdResponse) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_api_api_proto protoreflect.FileDescriptor

const file_api_api_proto_rawDesc = "" +
	"\n" +
	"\rapi/api.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf0\x02\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1c\n" +
	"\tfirstName\x18\x04 \x01(\tR\tfirstName\x12\x10\n" +
	"\x03bio\x18\x05 \x01(\tR\x03bio\x12\x16\n" +
	"\x06gender\x18\x06 \x01(\tR\x06gender\x128\n" +
	"\tbirthDate\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tbirthDate\x12\x12\n" +
	"\x04city\x18\b \x01(\tR\x04city\x12\x1c\n" +
	"\tavatarUrl\x18\t \x01(\tR\tavatarUrl\x128\n" +
	"\tcreatedAt\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x128\n" +
	"\tupdatedAt\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\".\n" +
	"\x11CreateUserRequest\x12\x19\n" +
	"\x04user\x18\x01 \x01(\v2\x05.UserR\x04user\"\x14\n" +
	"\x12CreateUserResponse\"J\n" +
	"\fUserPosition\x12\x10\n" +
	"\x03lon\x18\x01 \x01(\x02R\x03lon\x12\x10\n" +
	"\x03lat\x18\x02 \x01(\x02R\x03lat\x12\x16\n" +
	"\x06radius\x18\x03 \x01(\x05R\x06radius\"\xbb\x01\n" +
	"\vPreferences\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12(\n" +
	"\x0fpreferredGender\x18\x03 \x01(\tR\x0fpreferredGender\x12\x16\n" +
	"\x06ageMin\x18\x04 \x01(\x05R\x06ageMin\x12\x16\n" +
	"\x06ageMax\x18\x05 \x01(\x05R\x06ageMax\x12)\n" +
	"\bposition\x18\x06 \x01(\v2\r.UserPositionR\bposition\"\xec\x01\n" +
	"\x1fSelectUsersByPreferencesRequest\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userId\x12(\n" +
	"\x0fpreferredGender\x18\x02 \x01(\tR\x0fpreferredGender\x12\x16\n" +
	"\x06ageMin\x18\x03 \x01(\x05R\x06ageMin\x12\x16\n" +
	"\x06ageMax\x18\x04 \x01(\x05R\x06ageMax\x12)\n" +
	"\bposition\x18\x05 \x01(\v2\r.UserPositionR\bposition\x12\x14\n" +
	"\x05limit\x18\x06 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\a \x01(\x05R\x06offset\"?\n" +
	" SelectUsersByPreferencesResponse\x12\x1b\n" +
	"\x05users\x18\x01 \x03(\v2\x05.UserR\x05users\"9\n" +
	"\x13RefreshTokenRequest\x12\"\n" +
	"\frefreshToken\x18\x01 \x01(\tR\frefreshToken\"\\\n" +
	"\x14RefreshTokenResponse\x12 \n" +
	"\vaccessToken\x18\x01 \x01(\tR\vaccessToken\x12\"\n" +
	"\frefreshToken\x18\x02 \x01(\tR\frefreshToken\"|\n" +
	"\n" +
	"MatchModel\x12\x16\n" +
	"\x06fromId\x18\x01 \x01(\tR\x06fromId\x12\x12\n" +
	"\x04toId\x18\x02 \x01(\tR\x04toId\x12\"\n" +
	"\ffromDecision\x18\x03 \x01(\bR\ffromDecision\x12\x1e\n" +
	"\n" +
	"toDecision\x18\x04 \x01(\bR\n" +
	"toDecision\"3\n" +
	"\x19GetDecisionsUserIdRequest\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userId\"6\n" +
	"\x1aGetDecisionsUserIdResponse\x12\x18\n" +
	"\auserIds\x18\x01 \x03(\tR\auserIds2\xa6\x01\n" +
	"\x04user\x12=\n" +
	"\x12CreateUserFromAuth\x12\x12.CreateUserRequest\x1a\x13.CreateUserResponse\x12_\n" +
	"\x18SelectUsersByPreferences\x12 .SelectUsersByPreferencesRequest\x1a!.SelectUsersByPreferencesResponse2C\n" +
	"\x04auth\x12;\n" +
	"\fRefreshToken\x12\x14.RefreshTokenRequest\x1a\x15.RefreshTokenResponse2V\n" +
	"\x05match\x12M\n" +
	"\x12GetDecisionsUserId\x12\x1a.GetDecisionsUserIdRequest\x1a\x1b.GetDecisionsUserIdResponseB\x06Z\x04/genb\x06proto3"

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData []byte
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_api_proto_rawDesc), len(file_api_api_proto_rawDesc)))
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_api_proto_goTypes = []any{
	(*User)(nil),                             // 0: User
	(*CreateUserRequest)(nil),                // 1: CreateUserRequest
	(*CreateUserResponse)(nil),               // 2: CreateUserResponse
	(*UserPosition)(nil),                     // 3: UserPosition
	(*Preferences)(nil),                      // 4: Preferences
	(*SelectUsersByPreferencesRequest)(nil),  // 5: SelectUsersByPreferencesRequest
	(*SelectUsersByPreferencesResponse)(nil), // 6: SelectUsersByPreferencesResponse
	(*RefreshTokenRequest)(nil),              // 7: RefreshTokenRequest
	(*RefreshTokenResponse)(nil),             // 8: RefreshTokenResponse
	(*MatchModel)(nil),                       // 9: MatchModel
	(*GetDecisionsUserIdRequest)(nil),        // 10: GetDecisionsUserIdRequest
	(*GetDecisionsUserIdResponse)(nil),       // 11: GetDecisionsUserIdResponse
	(*timestamppb.Timestamp)(nil),            // 12: google.protobuf.Timestamp
}
var file_api_api_proto_depIdxs = []int32{
	12, // 0: User.birthDate:type_name -> google.protobuf.Timestamp
	12, // 1: User.createdAt:type_name -> google.protobuf.Timestamp
	12, // 2: User.updatedAt:type_name -> google.protobuf.Timestamp
	0,  // 3: CreateUserRequest.user:type_name -> User
	3,  // 4: Preferences.position:type_name -> UserPosition
	3,  // 5: SelectUsersByPreferencesRequest.position:type_name -> UserPosition
	0,  // 6: SelectUsersByPreferencesResponse.users:type_name -> User
	1,  // 7: user.CreateUserFromAuth:input_type -> CreateUserRequest
	5,  // 8: user.SelectUsersByPreferences:input_type -> SelectUsersByPreferencesRequest
	7,  // 9: auth.RefreshToken:input_type -> RefreshTokenRequest
	10, // 10: match.GetDecisionsUserId:input_type -> GetDecisionsUserIdRequest
	2,  // 11: user.CreateUserFromAuth:output_type -> CreateUserResponse
	6,  // 12: user.SelectUsersByPreferences:output_type -> SelectUsersByPreferencesResponse
	8,  // 13: auth.RefreshToken:output_type -> RefreshTokenResponse
	11, // 14: match.GetDecisionsUserId:output_type -> GetDecisionsUserIdResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_api_proto_rawDesc), len(file_api_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
