// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Activity_ActivityType int32

const (
	Activity_Play  Activity_ActivityType = 0
	Activity_Sleep Activity_ActivityType = 1
	Activity_Eat   Activity_ActivityType = 2
	Activity_Read  Activity_ActivityType = 3
)

var Activity_ActivityType_name = map[int32]string{
	0: "Play",
	1: "Sleep",
	2: "Eat",
	3: "Read",
}

var Activity_ActivityType_value = map[string]int32{
	"Play":  0,
	"Sleep": 1,
	"Eat":   2,
	"Read":  3,
}

func (x Activity_ActivityType) String() string {
	return proto.EnumName(Activity_ActivityType_name, int32(x))
}

func (Activity_ActivityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{5, 0}
}

type User struct {
	Userid               string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type UserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type QueryUserRequest struct {
	Userid               string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserRequest) Reset()         { *m = QueryUserRequest{} }
func (m *QueryUserRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserRequest) ProtoMessage()    {}
func (*QueryUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{3}
}

func (m *QueryUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserRequest.Unmarshal(m, b)
}
func (m *QueryUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserRequest.Marshal(b, m, deterministic)
}
func (m *QueryUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserRequest.Merge(m, src)
}
func (m *QueryUserRequest) XXX_Size() int {
	return xxx_messageInfo_QueryUserRequest.Size(m)
}
func (m *QueryUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserRequest proto.InternalMessageInfo

func (m *QueryUserRequest) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

type QueryUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserResponse) Reset()         { *m = QueryUserResponse{} }
func (m *QueryUserResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserResponse) ProtoMessage()    {}
func (*QueryUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{4}
}

func (m *QueryUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserResponse.Unmarshal(m, b)
}
func (m *QueryUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserResponse.Marshal(b, m, deterministic)
}
func (m *QueryUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserResponse.Merge(m, src)
}
func (m *QueryUserResponse) XXX_Size() int {
	return xxx_messageInfo_QueryUserResponse.Size(m)
}
func (m *QueryUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserResponse proto.InternalMessageInfo

func (m *QueryUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Activity struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ActivityType         Activity_ActivityType `protobuf:"varint,2,opt,name=activityType,proto3,enum=user.Activity_ActivityType" json:"activityType,omitempty"`
	Timestamp            int64                 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration             int64                 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Label                string                `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Email                string                `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{5}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Activity) GetActivityType() Activity_ActivityType {
	if m != nil {
		return m.ActivityType
	}
	return Activity_Play
}

func (m *Activity) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Activity) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Activity) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Activity) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ActivityRequest struct {
	Activity             *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ActivityRequest) Reset()         { *m = ActivityRequest{} }
func (m *ActivityRequest) String() string { return proto.CompactTextString(m) }
func (*ActivityRequest) ProtoMessage()    {}
func (*ActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{6}
}

func (m *ActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityRequest.Unmarshal(m, b)
}
func (m *ActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityRequest.Marshal(b, m, deterministic)
}
func (m *ActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityRequest.Merge(m, src)
}
func (m *ActivityRequest) XXX_Size() int {
	return xxx_messageInfo_ActivityRequest.Size(m)
}
func (m *ActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityRequest proto.InternalMessageInfo

func (m *ActivityRequest) GetActivity() *Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

type ActivityResponse struct {
	Activity             *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ActivityResponse) Reset()         { *m = ActivityResponse{} }
func (m *ActivityResponse) String() string { return proto.CompactTextString(m) }
func (*ActivityResponse) ProtoMessage()    {}
func (*ActivityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{7}
}

func (m *ActivityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityResponse.Unmarshal(m, b)
}
func (m *ActivityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityResponse.Marshal(b, m, deterministic)
}
func (m *ActivityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityResponse.Merge(m, src)
}
func (m *ActivityResponse) XXX_Size() int {
	return xxx_messageInfo_ActivityResponse.Size(m)
}
func (m *ActivityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityResponse proto.InternalMessageInfo

func (m *ActivityResponse) GetActivity() *Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

type QueryActivityRequest struct {
	Activityid           string   `protobuf:"bytes,1,opt,name=activityid,proto3" json:"activityid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryActivityRequest) Reset()         { *m = QueryActivityRequest{} }
func (m *QueryActivityRequest) String() string { return proto.CompactTextString(m) }
func (*QueryActivityRequest) ProtoMessage()    {}
func (*QueryActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{8}
}

func (m *QueryActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryActivityRequest.Unmarshal(m, b)
}
func (m *QueryActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryActivityRequest.Marshal(b, m, deterministic)
}
func (m *QueryActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryActivityRequest.Merge(m, src)
}
func (m *QueryActivityRequest) XXX_Size() int {
	return xxx_messageInfo_QueryActivityRequest.Size(m)
}
func (m *QueryActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryActivityRequest proto.InternalMessageInfo

func (m *QueryActivityRequest) GetActivityid() string {
	if m != nil {
		return m.Activityid
	}
	return ""
}

type QueryActivityResponse struct {
	Activity             *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryActivityResponse) Reset()         { *m = QueryActivityResponse{} }
func (m *QueryActivityResponse) String() string { return proto.CompactTextString(m) }
func (*QueryActivityResponse) ProtoMessage()    {}
func (*QueryActivityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{9}
}

func (m *QueryActivityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryActivityResponse.Unmarshal(m, b)
}
func (m *QueryActivityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryActivityResponse.Marshal(b, m, deterministic)
}
func (m *QueryActivityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryActivityResponse.Merge(m, src)
}
func (m *QueryActivityResponse) XXX_Size() int {
	return xxx_messageInfo_QueryActivityResponse.Size(m)
}
func (m *QueryActivityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryActivityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryActivityResponse proto.InternalMessageInfo

func (m *QueryActivityResponse) GetActivity() *Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

func init() {
	proto.RegisterEnum("user.Activity_ActivityType", Activity_ActivityType_name, Activity_ActivityType_value)
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserRequest)(nil), "user.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*QueryUserRequest)(nil), "user.QueryUserRequest")
	proto.RegisterType((*QueryUserResponse)(nil), "user.QueryUserResponse")
	proto.RegisterType((*Activity)(nil), "user.Activity")
	proto.RegisterType((*ActivityRequest)(nil), "user.ActivityRequest")
	proto.RegisterType((*ActivityResponse)(nil), "user.ActivityResponse")
	proto.RegisterType((*QueryActivityRequest)(nil), "user.QueryActivityRequest")
	proto.RegisterType((*QueryActivityResponse)(nil), "user.QueryActivityResponse")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0x6c, 0x27, 0x4d, 0x26, 0x21, 0x75, 0x47, 0x6d, 0xb0, 0x52, 0x54, 0x21, 0x9f, 0x50,
	0x25, 0x5c, 0x94, 0xa2, 0xc2, 0x85, 0xa2, 0x80, 0xb8, 0xf4, 0x04, 0x06, 0x2e, 0x1c, 0x90, 0x36,
	0xdd, 0x55, 0x63, 0x61, 0xc7, 0xc6, 0x5e, 0x23, 0xf9, 0x27, 0xf8, 0x5f, 0x6e, 0xc8, 0xb3, 0xf6,
	0x66, 0x63, 0x40, 0xb4, 0x17, 0x6b, 0xe7, 0xcd, 0xbc, 0x99, 0xd9, 0x37, 0xe3, 0x85, 0x83, 0xb2,
	0x10, 0xf9, 0x79, 0xfd, 0x09, 0xb2, 0x3c, 0x95, 0x29, 0x3a, 0xf5, 0xd9, 0xff, 0x0a, 0xce, 0xe7,
	0x42, 0xe4, 0x38, 0x83, 0x41, 0x6d, 0x47, 0xdc, 0xeb, 0x3d, 0xee, 0x3d, 0x19, 0x85, 0x8d, 0x85,
	0x08, 0xce, 0x86, 0x25, 0xc2, 0xb3, 0x08, 0xa5, 0x33, 0x1e, 0x41, 0x5f, 0x24, 0x2c, 0x8a, 0x3d,
	0x9b, 0x40, 0x65, 0xd4, 0x68, 0xb6, 0x4e, 0x37, 0xc2, 0x73, 0x14, 0x4a, 0x86, 0xff, 0x14, 0xc6,
	0x75, 0xfe, 0x50, 0x7c, 0x2f, 0x45, 0x21, 0xf1, 0x14, 0xa8, 0x2c, 0x15, 0x19, 0x2f, 0x20, 0xa0,
	0x7e, 0x28, 0x40, 0xb5, 0x13, 0xc0, 0x44, 0x85, 0x17, 0x59, 0xba, 0x29, 0xc4, 0x7f, 0xe3, 0xcf,
	0xc0, 0xfd, 0x50, 0x8a, 0xbc, 0x32, 0x6b, 0xfc, 0xe3, 0x2a, 0xfe, 0x05, 0x1c, 0x1a, 0xb1, 0x77,
	0x2c, 0xf0, 0xab, 0x07, 0xc3, 0xe5, 0x8d, 0x8c, 0x7e, 0x44, 0xb2, 0xc2, 0x29, 0x58, 0x3a, 0xab,
	0x15, 0x71, 0x7c, 0x0d, 0x13, 0xd6, 0xf8, 0x3e, 0x55, 0x99, 0x12, 0x69, 0xba, 0x38, 0x51, 0x49,
	0x5a, 0x96, 0x3e, 0xd4, 0x21, 0xe1, 0x0e, 0x01, 0x1f, 0xc1, 0x48, 0x46, 0x89, 0x28, 0x24, 0x4b,
	0x32, 0x52, 0xd3, 0x0e, 0xb7, 0x00, 0xce, 0x61, 0xc8, 0xcb, 0x9c, 0xc9, 0x28, 0xdd, 0x90, 0xa8,
	0x76, 0xa8, 0xed, 0x5a, 0xed, 0x98, 0xad, 0x44, 0xec, 0xf5, 0x95, 0xda, 0x64, 0x6c, 0x27, 0x33,
	0x30, 0x26, 0xe3, 0x5f, 0xc2, 0xc4, 0xec, 0x01, 0x87, 0xe0, 0xbc, 0x8f, 0x59, 0xe5, 0xee, 0xe1,
	0x08, 0xfa, 0x1f, 0x63, 0x21, 0x32, 0xb7, 0x87, 0xfb, 0x60, 0xbf, 0x63, 0xd2, 0xb5, 0x6a, 0x6f,
	0x28, 0x18, 0x77, 0x6d, 0xff, 0x15, 0x1c, 0xb4, 0xbc, 0x56, 0xdb, 0x33, 0x18, 0xb6, 0x17, 0x68,
	0x24, 0x9b, 0xee, 0xde, 0x36, 0xd4, 0x7e, 0xff, 0x0a, 0xdc, 0x2d, 0xbd, 0x91, 0xfb, 0x3e, 0xfc,
	0x4b, 0x38, 0xa2, 0x79, 0x75, 0x7b, 0x38, 0x05, 0x68, 0x63, 0xf4, 0x34, 0x0c, 0xc4, 0x7f, 0x0b,
	0xc7, 0x1d, 0xde, 0xfd, 0x8b, 0x2f, 0x7e, 0x5a, 0x6a, 0x13, 0xf5, 0xec, 0x17, 0xb0, 0xbf, 0xe4,
	0x9c, 0xfe, 0x95, 0x43, 0x63, 0x4b, 0x54, 0x4f, 0x73, 0x34, 0x21, 0x55, 0xce, 0xdf, 0xc3, 0x2b,
	0x18, 0xe9, 0x8d, 0xc3, 0x99, 0x0a, 0xe9, 0xae, 0xeb, 0xfc, 0xe1, 0x1f, 0xb8, 0xc1, 0x1f, 0x2f,
	0x39, 0xd7, 0x2d, 0x1c, 0x77, 0xba, 0x6d, 0x12, 0xcc, 0xba, 0xb0, 0xe6, 0x5f, 0xc3, 0x83, 0x1d,
	0x25, 0x70, 0x6e, 0xd4, 0xea, 0xa6, 0x39, 0xf9, 0xab, 0xaf, 0xcd, 0xf5, 0xe6, 0xd9, 0x97, 0xe0,
	0x36, 0x92, 0xeb, 0x72, 0x15, 0xdc, 0xa4, 0xc9, 0xf9, 0x35, 0xbb, 0x65, 0x5c, 0x88, 0x62, 0xfd,
	0xad, 0x4c, 0x58, 0xfe, 0xe2, 0xf9, 0x4b, 0x7a, 0x59, 0x5a, 0xf5, 0xc8, 0x58, 0x0d, 0xe8, 0x9d,
	0xb9, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x53, 0x71, 0xf2, 0x79, 0x7a, 0x04, 0x00, 0x00,
}