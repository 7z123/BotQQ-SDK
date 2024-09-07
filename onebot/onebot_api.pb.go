// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: onebot_api.proto

package onebot

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionType int32

const (
	ActionType_TUnknown                ActionType = 0
	ActionType_send_private_msg        ActionType = 1
	ActionType_send_group_msg          ActionType = 2
	ActionType_send_msg                ActionType = 3
	ActionType_send_forward_msg        ActionType = 4
	ActionType_delete_msg              ActionType = 5
	ActionType_get_msg                 ActionType = 6
	ActionType_get_forward_msg         ActionType = 7
	ActionType_send_like               ActionType = 8
	ActionType_set_group_kick          ActionType = 9
	ActionType_set_group_ban           ActionType = 10
	ActionType_set_group_anonymous_ban ActionType = 11
	ActionType_set_group_whole_ban     ActionType = 12
	ActionType_set_group_admin         ActionType = 13
	ActionType_set_group_anonymous     ActionType = 14
	ActionType_set_group_card          ActionType = 15
	ActionType_set_group_name          ActionType = 16
	ActionType_set_group_leave         ActionType = 17
	ActionType_set_group_special_title ActionType = 18
	ActionType_set_friend_add_request  ActionType = 19
	ActionType_set_group_add_request   ActionType = 20
	ActionType_get_login_info          ActionType = 21
	ActionType_get_stranger_info       ActionType = 22
	ActionType_get_friend_list         ActionType = 23
	ActionType_get_group_info          ActionType = 24
	ActionType_get_group_list          ActionType = 25
	ActionType_get_group_member_info   ActionType = 26
	ActionType_get_group_member_list   ActionType = 27
	ActionType_get_group_honor_info    ActionType = 28
	ActionType_get_cookies             ActionType = 29
	ActionType_get_csrf_token          ActionType = 30
	ActionType_get_credentials         ActionType = 31
	ActionType_get_record              ActionType = 32
	ActionType_get_image               ActionType = 33
	ActionType_can_send_image          ActionType = 34
	ActionType_can_send_record         ActionType = 35
	ActionType_get_status              ActionType = 36
	ActionType_get_version_info        ActionType = 37
	ActionType_set_restart             ActionType = 38
	ActionType_clean_cache             ActionType = 39
	ActionType_group_poke              ActionType = 40
	ActionType_friend_poke             ActionType = 41
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0:  "TUnknown",
		1:  "send_private_msg",
		2:  "send_group_msg",
		3:  "send_msg",
		4:  "send_forward_msg",
		5:  "delete_msg",
		6:  "get_msg",
		7:  "get_forward_msg",
		8:  "send_like",
		9:  "set_group_kick",
		10: "set_group_ban",
		11: "set_group_anonymous_ban",
		12: "set_group_whole_ban",
		13: "set_group_admin",
		14: "set_group_anonymous",
		15: "set_group_card",
		16: "set_group_name",
		17: "set_group_leave",
		18: "set_group_special_title",
		19: "set_friend_add_request",
		20: "set_group_add_request",
		21: "get_login_info",
		22: "get_stranger_info",
		23: "get_friend_list",
		24: "get_group_info",
		25: "get_group_list",
		26: "get_group_member_info",
		27: "get_group_member_list",
		28: "get_group_honor_info",
		29: "get_cookies",
		30: "get_csrf_token",
		31: "get_credentials",
		32: "get_record",
		33: "get_image",
		34: "can_send_image",
		35: "can_send_record",
		36: "get_status",
		37: "get_version_info",
		38: "set_restart",
		39: "clean_cache",
		40: "group_poke",
		41: "friend_poke",
	}
	ActionType_value = map[string]int32{
		"TUnknown":                0,
		"send_private_msg":        1,
		"send_group_msg":          2,
		"send_msg":                3,
		"send_forward_msg":        4,
		"delete_msg":              5,
		"get_msg":                 6,
		"get_forward_msg":         7,
		"send_like":               8,
		"set_group_kick":          9,
		"set_group_ban":           10,
		"set_group_anonymous_ban": 11,
		"set_group_whole_ban":     12,
		"set_group_admin":         13,
		"set_group_anonymous":     14,
		"set_group_card":          15,
		"set_group_name":          16,
		"set_group_leave":         17,
		"set_group_special_title": 18,
		"set_friend_add_request":  19,
		"set_group_add_request":   20,
		"get_login_info":          21,
		"get_stranger_info":       22,
		"get_friend_list":         23,
		"get_group_info":          24,
		"get_group_list":          25,
		"get_group_member_info":   26,
		"get_group_member_list":   27,
		"get_group_honor_info":    28,
		"get_cookies":             29,
		"get_csrf_token":          30,
		"get_credentials":         31,
		"get_record":              32,
		"get_image":               33,
		"can_send_image":          34,
		"can_send_record":         35,
		"get_status":              36,
		"get_version_info":        37,
		"set_restart":             38,
		"clean_cache":             39,
		"group_poke":              40,
		"friend_poke":             41,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_onebot_api_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_onebot_api_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_onebot_api_proto_rawDescGZIP(), []int{0}
}

type Api struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string      `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Params *Api_Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Echo   string      `protobuf:"bytes,3,opt,name=echo,proto3" json:"echo,omitempty"`
}

func (x *Api) Reset() {
	*x = Api{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onebot_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Api) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Api) ProtoMessage() {}

func (x *Api) ProtoReflect() protoreflect.Message {
	mi := &file_onebot_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Api.ProtoReflect.Descriptor instead.
func (*Api) Descriptor() ([]byte, []int) {
	return file_onebot_api_proto_rawDescGZIP(), []int{0}
}

func (x *Api) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Api) GetParams() *Api_Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Api) GetEcho() string {
	if x != nil {
		return x.Echo
	}
	return ""
}

type Api_Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId          string    `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GuildId          string    `protobuf:"bytes,3,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	Message          *Message  `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Messages         *Messages `protobuf:"bytes,5,opt,name=messages,proto3" json:"messages,omitempty"`
	MessageType      string    `protobuf:"bytes,6,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	AutoEscape       bool      `protobuf:"varint,7,opt,name=auto_escape,json=autoEscape,proto3" json:"auto_escape,omitempty"`
	MessageId        string    `protobuf:"bytes,8,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Id               string    `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
	RejectAddRequest bool      `protobuf:"varint,10,opt,name=reject_add_request,json=rejectAddRequest,proto3" json:"reject_add_request,omitempty"`
	Duration         int32     `protobuf:"varint,11,opt,name=duration,proto3" json:"duration,omitempty"`
	Enable           bool      `protobuf:"varint,12,opt,name=enable,proto3" json:"enable,omitempty"`
	Card             string    `protobuf:"bytes,13,opt,name=card,proto3" json:"card,omitempty"`
	GroupName        string    `protobuf:"bytes,14,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Approve          bool      `protobuf:"varint,15,opt,name=approve,proto3" json:"approve,omitempty"`
	Remark           string    `protobuf:"bytes,16,opt,name=remark,proto3" json:"remark,omitempty"`
	Content          string    `protobuf:"bytes,17,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Api_Params) Reset() {
	*x = Api_Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onebot_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Api_Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Api_Params) ProtoMessage() {}

func (x *Api_Params) ProtoReflect() protoreflect.Message {
	mi := &file_onebot_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Api_Params.ProtoReflect.Descriptor instead.
func (*Api_Params) Descriptor() ([]byte, []int) {
	return file_onebot_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Api_Params) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Api_Params) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Api_Params) GetGuildId() string {
	if x != nil {
		return x.GuildId
	}
	return ""
}

func (x *Api_Params) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Api_Params) GetMessages() *Messages {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Api_Params) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *Api_Params) GetAutoEscape() bool {
	if x != nil {
		return x.AutoEscape
	}
	return false
}

func (x *Api_Params) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Api_Params) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Api_Params) GetRejectAddRequest() bool {
	if x != nil {
		return x.RejectAddRequest
	}
	return false
}

func (x *Api_Params) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Api_Params) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Api_Params) GetCard() string {
	if x != nil {
		return x.Card
	}
	return ""
}

func (x *Api_Params) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *Api_Params) GetApprove() bool {
	if x != nil {
		return x.Approve
	}
	return false
}

func (x *Api_Params) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Api_Params) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_onebot_api_proto protoreflect.FileDescriptor

var file_onebot_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x1a, 0x11, 0x6f, 0x6e, 0x65, 0x62,
	0x6f, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x04,
	0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x63, 0x68,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x1a, 0x84, 0x04,
	0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x6e, 0x65, 0x62, 0x6f,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x65, 0x73, 0x63, 0x61,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x45, 0x73,
	0x63, 0x61, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2a, 0xe8, 0x06, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x73, 0x67, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x73, 0x67, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x10, 0x04, 0x12,
	0x0e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x10, 0x05, 0x12,
	0x0b, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x10,
	0x07, 0x12, 0x0d, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x10, 0x08,
	0x12, 0x12, 0x0a, 0x0e, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x69,
	0x63, 0x6b, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x62, 0x61, 0x6e, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x73, 0x65, 0x74, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x5f, 0x62,
	0x61, 0x6e, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x77, 0x68, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x10, 0x0c, 0x12, 0x13, 0x0a,
	0x0f, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x73,
	0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x10, 0x0f, 0x12,
	0x12, 0x0a, 0x0e, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x10, 0x11, 0x12, 0x1b, 0x0a, 0x17, 0x73, 0x65, 0x74, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x10, 0x12, 0x12, 0x1a, 0x0a, 0x16, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10,
	0x13, 0x12, 0x19, 0x0a, 0x15, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x14, 0x12, 0x12, 0x0a, 0x0e,
	0x67, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x15,
	0x12, 0x15, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x16, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x10, 0x17, 0x12, 0x12, 0x0a, 0x0e,
	0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x18,
	0x12, 0x12, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x10, 0x19, 0x12, 0x19, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x1a, 0x12,
	0x19, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x10, 0x1b, 0x12, 0x18, 0x0a, 0x14, 0x67, 0x65,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x68, 0x6f, 0x6e, 0x6f, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x10, 0x1c, 0x12, 0x0f, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x73, 0x10, 0x1d, 0x12, 0x12, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x73, 0x72,
	0x66, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x1e, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x10, 0x1f, 0x12, 0x0e,
	0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x10, 0x20, 0x12, 0x0d,
	0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x21, 0x12, 0x12, 0x0a,
	0x0e, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x10,
	0x22, 0x12, 0x13, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x10, 0x23, 0x12, 0x0e, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x10, 0x24, 0x12, 0x14, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x25, 0x12, 0x0f, 0x0a, 0x0b,
	0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x10, 0x26, 0x12, 0x0f, 0x0a,
	0x0b, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x10, 0x27, 0x12, 0x0e,
	0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x6f, 0x6b, 0x65, 0x10, 0x28, 0x12, 0x0f,
	0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x6b, 0x65, 0x10, 0x29, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_onebot_api_proto_rawDescOnce sync.Once
	file_onebot_api_proto_rawDescData = file_onebot_api_proto_rawDesc
)

func file_onebot_api_proto_rawDescGZIP() []byte {
	file_onebot_api_proto_rawDescOnce.Do(func() {
		file_onebot_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_onebot_api_proto_rawDescData)
	})
	return file_onebot_api_proto_rawDescData
}

var file_onebot_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_onebot_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_onebot_api_proto_goTypes = []interface{}{
	(ActionType)(0),    // 0: onebot.ActionType
	(*Api)(nil),        // 1: onebot.Api
	(*Api_Params)(nil), // 2: onebot.Api.Params
	(*Message)(nil),    // 3: onebot.Message
	(*Messages)(nil),   // 4: onebot.Messages
}
var file_onebot_api_proto_depIdxs = []int32{
	2, // 0: onebot.Api.params:type_name -> onebot.Api.Params
	3, // 1: onebot.Api.Params.message:type_name -> onebot.Message
	4, // 2: onebot.Api.Params.messages:type_name -> onebot.Messages
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_onebot_api_proto_init() }
func file_onebot_api_proto_init() {
	if File_onebot_api_proto != nil {
		return
	}
	file_onebot_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_onebot_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Api); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_onebot_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Api_Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_onebot_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_onebot_api_proto_goTypes,
		DependencyIndexes: file_onebot_api_proto_depIdxs,
		EnumInfos:         file_onebot_api_proto_enumTypes,
		MessageInfos:      file_onebot_api_proto_msgTypes,
	}.Build()
	File_onebot_api_proto = out.File
	file_onebot_api_proto_rawDesc = nil
	file_onebot_api_proto_goTypes = nil
	file_onebot_api_proto_depIdxs = nil
}
