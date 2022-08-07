// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: id.proto

package pb

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

// 定义全局消息 ID
type MessageID int32

const (
	MessageID_LOGIN_REQUIRED MessageID = 0
	// 账号注册
	MessageID_CS_ACCOUNT_REGISTER MessageID = 1 //CS_AccountRegister_Req
	MessageID_SC_ACCOUNT_REGISTER MessageID = 2 //SC_AccountRegister_Res
	// 账号登录
	MessageID_CS_ACCOUNT_LOGIN MessageID = 3 //CS_AccountLogin_Req
	MessageID_SC_ACCOUNT_LOGIN MessageID = 4 //SC_AccountLogin_Res
	// 服务端定时向客户端更新体力值
	MessageID_SYNC_ACCOUNT_TP MessageID = 1201 // SC_SyncAccountTP
	// 客户端获取玩家体力
	MessageID_CS_GET_ACCOUNT_TP MessageID = 1202 // CS_GetAccountTP_Req
	MessageID_SC_GET_ACCOUNT_TP MessageID = 1203 // SC_GetAccountTP_Res
	// 玩家点击「选拔开始」消耗固定体力
	MessageID_CS_SELECTION_CONSUME_TP MessageID = 1204 //CS_SelectionConsumeTP_Req
	MessageID_SC_SELECTION_CONSUME_TP MessageID = 1205 //SC_SelectionConsumeTP_Res
	// 获取玩家可以使用的 base_role
	MessageID_CS_GET_BASE_ROLE MessageID = 1301 //CS_GetAccountBaseRole_Req
	MessageID_SC_GET_BASE_ROLE MessageID = 1302 //SC_GetAccountBaseRole_Res
	// 传入 base_role_id 和对应的进阶系数，生成一个训练角色
	// 玩家没有训练中的角色，则会根据选择的 base_role 生成一个对应的 training_role
	MessageID_CS_GENERATE_TRAINING_ROLE MessageID = 1401 //CS_GenerateTrainingRole_Req
	MessageID_SC_GENERATE_TRAINING_ROLE MessageID = 1402 //SC_GenerateTrainingRole_Res
	// 客户端请求获取玩家所有可用的支援卡
	MessageID_CS_LIST_SUPPORT_CARD MessageID = 1501 //CS_ListSupportCard_Req
	MessageID_SC_LIST_SUPPORT_CARD MessageID = 1502 //SC_ListSupportCard_Res
	// 客户端提交选定的所有支援卡
	MessageID_CS_COMMIT_SUPPORT_CARD MessageID = 1503 //CS_CommitSupportCard_Req
	MessageID_SC_COMMIT_SUPPORT_CARD MessageID = 1504 //SC_CommitSupportCard_Res
	// 客户端请求获取一批日程
	MessageID_CS_LIST_SCHEDULE MessageID = 1601 //CS_ListSchedule_Req
	MessageID_SC_LIST_SCHEDULE MessageID = 1602 //SC_ListSchedule_Res
	// 客户端提交选定的日程
	MessageID_CS_COMMIT_SCHEDULE MessageID = 1603 //CS_CommitSchedule_Req
	MessageID_SC_COMMIT_SCHEDULE MessageID = 1604 //SC_CommitSchedule_Res
	// 更新事件的阶段，默认 +1
	MessageID_CS_UPDATE_ENENT_PROCESS MessageID = 1701 //CS_UpdateEventProgress_Req
	MessageID_SC_UPDATE_EVENT_PROCESS MessageID = 1702 //SC_UpdateEventProgress_Res
	// 客户端提交选定的日程分支
	MessageID_CS_COMMIT_EVENT_BRANCH MessageID = 1703 //CS_CommitEventBranch_Req
	MessageID_SC_COMMIT_EVENT_BRANCH MessageID = 1704 //SC_CommitEventBranch_Res
	// 训练中购买芯片
	MessageID_CS_BUY_TRAINING_CHIP MessageID = 5001 //CS_BuyTrainingChip_Req
	MessageID_SC_BUY_TRAINING_CHIP MessageID = 5002 //SC_BuyTrainingChip_Res
	// 训练中购买遗物
	MessageID_CS_BUY_TRAINING_RELICS MessageID = 5003 //CS_BuyTrainingRelics_Req
	MessageID_SC_BUY_TRAINING_RELICS MessageID = 5004 //SC_BuyTrainingRelics_Res
	// 训练中购买训练道具
	MessageID_CS_BUY_TRAINING_CULTIVATE_GOODS MessageID = 5005 //CS_BuyTrainingCultivateGoods_Req
	MessageID_SC_BUY_TRAINING_CULTIVATE_GOODS MessageID = 5006 //SC_BuyTrainingCultivateGoods_Res
)

// Enum value maps for MessageID.
var (
	MessageID_name = map[int32]string{
		0:    "LOGIN_REQUIRED",
		1:    "CS_ACCOUNT_REGISTER",
		2:    "SC_ACCOUNT_REGISTER",
		3:    "CS_ACCOUNT_LOGIN",
		4:    "SC_ACCOUNT_LOGIN",
		1201: "SYNC_ACCOUNT_TP",
		1202: "CS_GET_ACCOUNT_TP",
		1203: "SC_GET_ACCOUNT_TP",
		1204: "CS_SELECTION_CONSUME_TP",
		1205: "SC_SELECTION_CONSUME_TP",
		1301: "CS_GET_BASE_ROLE",
		1302: "SC_GET_BASE_ROLE",
		1401: "CS_GENERATE_TRAINING_ROLE",
		1402: "SC_GENERATE_TRAINING_ROLE",
		1501: "CS_LIST_SUPPORT_CARD",
		1502: "SC_LIST_SUPPORT_CARD",
		1503: "CS_COMMIT_SUPPORT_CARD",
		1504: "SC_COMMIT_SUPPORT_CARD",
		1601: "CS_LIST_SCHEDULE",
		1602: "SC_LIST_SCHEDULE",
		1603: "CS_COMMIT_SCHEDULE",
		1604: "SC_COMMIT_SCHEDULE",
		1701: "CS_UPDATE_ENENT_PROCESS",
		1702: "SC_UPDATE_EVENT_PROCESS",
		1703: "CS_COMMIT_EVENT_BRANCH",
		1704: "SC_COMMIT_EVENT_BRANCH",
		5001: "CS_BUY_TRAINING_CHIP",
		5002: "SC_BUY_TRAINING_CHIP",
		5003: "CS_BUY_TRAINING_RELICS",
		5004: "SC_BUY_TRAINING_RELICS",
		5005: "CS_BUY_TRAINING_CULTIVATE_GOODS",
		5006: "SC_BUY_TRAINING_CULTIVATE_GOODS",
	}
	MessageID_value = map[string]int32{
		"LOGIN_REQUIRED":                  0,
		"CS_ACCOUNT_REGISTER":             1,
		"SC_ACCOUNT_REGISTER":             2,
		"CS_ACCOUNT_LOGIN":                3,
		"SC_ACCOUNT_LOGIN":                4,
		"SYNC_ACCOUNT_TP":                 1201,
		"CS_GET_ACCOUNT_TP":               1202,
		"SC_GET_ACCOUNT_TP":               1203,
		"CS_SELECTION_CONSUME_TP":         1204,
		"SC_SELECTION_CONSUME_TP":         1205,
		"CS_GET_BASE_ROLE":                1301,
		"SC_GET_BASE_ROLE":                1302,
		"CS_GENERATE_TRAINING_ROLE":       1401,
		"SC_GENERATE_TRAINING_ROLE":       1402,
		"CS_LIST_SUPPORT_CARD":            1501,
		"SC_LIST_SUPPORT_CARD":            1502,
		"CS_COMMIT_SUPPORT_CARD":          1503,
		"SC_COMMIT_SUPPORT_CARD":          1504,
		"CS_LIST_SCHEDULE":                1601,
		"SC_LIST_SCHEDULE":                1602,
		"CS_COMMIT_SCHEDULE":              1603,
		"SC_COMMIT_SCHEDULE":              1604,
		"CS_UPDATE_ENENT_PROCESS":         1701,
		"SC_UPDATE_EVENT_PROCESS":         1702,
		"CS_COMMIT_EVENT_BRANCH":          1703,
		"SC_COMMIT_EVENT_BRANCH":          1704,
		"CS_BUY_TRAINING_CHIP":            5001,
		"SC_BUY_TRAINING_CHIP":            5002,
		"CS_BUY_TRAINING_RELICS":          5003,
		"SC_BUY_TRAINING_RELICS":          5004,
		"CS_BUY_TRAINING_CULTIVATE_GOODS": 5005,
		"SC_BUY_TRAINING_CULTIVATE_GOODS": 5006,
	}
)

func (x MessageID) Enum() *MessageID {
	p := new(MessageID)
	*p = x
	return p
}

func (x MessageID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageID) Descriptor() protoreflect.EnumDescriptor {
	return file_id_proto_enumTypes[0].Descriptor()
}

func (MessageID) Type() protoreflect.EnumType {
	return &file_id_proto_enumTypes[0]
}

func (x MessageID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageID.Descriptor instead.
func (MessageID) EnumDescriptor() ([]byte, []int) {
	return file_id_proto_rawDescGZIP(), []int{0}
}

var File_id_proto protoreflect.FileDescriptor

var file_id_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xef, 0x06, 0x0a, 0x09, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x43, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x0f, 0x53, 0x59,
	0x4e, 0x43, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x50, 0x10, 0xb1, 0x09,
	0x12, 0x16, 0x0a, 0x11, 0x43, 0x53, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x54, 0x50, 0x10, 0xb2, 0x09, 0x12, 0x16, 0x0a, 0x11, 0x53, 0x43, 0x5f, 0x47,
	0x45, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x50, 0x10, 0xb3, 0x09,
	0x12, 0x1c, 0x0a, 0x17, 0x43, 0x53, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x5f, 0x54, 0x50, 0x10, 0xb4, 0x09, 0x12, 0x1c,
	0x0a, 0x17, 0x53, 0x43, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x5f, 0x54, 0x50, 0x10, 0xb5, 0x09, 0x12, 0x15, 0x0a, 0x10,
	0x43, 0x53, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x10, 0x95, 0x0a, 0x12, 0x15, 0x0a, 0x10, 0x53, 0x43, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0x96, 0x0a, 0x12, 0x1e, 0x0a, 0x19, 0x43, 0x53,
	0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0xf9, 0x0a, 0x12, 0x1e, 0x0a, 0x19, 0x53, 0x43,
	0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0xfa, 0x0a, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x53,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41,
	0x52, 0x44, 0x10, 0xdd, 0x0b, 0x12, 0x19, 0x0a, 0x14, 0x53, 0x43, 0x5f, 0x4c, 0x49, 0x53, 0x54,
	0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0xde, 0x0b,
	0x12, 0x1b, 0x0a, 0x16, 0x43, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x5f, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0xdf, 0x0b, 0x12, 0x1b, 0x0a,
	0x16, 0x53, 0x43, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f,
	0x52, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0xe0, 0x0b, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x53,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0xc1,
	0x0c, 0x12, 0x15, 0x0a, 0x10, 0x53, 0x43, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x53, 0x43, 0x48,
	0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0xc2, 0x0c, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x53, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0xc3,
	0x0c, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x43, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x5f, 0x53,
	0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0xc4, 0x0c, 0x12, 0x1c, 0x0a, 0x17, 0x43, 0x53,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0xa5, 0x0d, 0x12, 0x1c, 0x0a, 0x17, 0x53, 0x43, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x10, 0xa6, 0x0d, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x53, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x49, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48,
	0x10, 0xa7, 0x0d, 0x12, 0x1b, 0x0a, 0x16, 0x53, 0x43, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x10, 0xa8, 0x0d,
	0x12, 0x19, 0x0a, 0x14, 0x43, 0x53, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e,
	0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x49, 0x50, 0x10, 0x89, 0x27, 0x12, 0x19, 0x0a, 0x14, 0x53,
	0x43, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x43,
	0x48, 0x49, 0x50, 0x10, 0x8a, 0x27, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x53, 0x5f, 0x42, 0x55, 0x59,
	0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x43, 0x53,
	0x10, 0x8b, 0x27, 0x12, 0x1b, 0x0a, 0x16, 0x53, 0x43, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x54, 0x52,
	0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x43, 0x53, 0x10, 0x8c, 0x27,
	0x12, 0x24, 0x0a, 0x1f, 0x43, 0x53, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e,
	0x49, 0x4e, 0x47, 0x5f, 0x43, 0x55, 0x4c, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x4f,
	0x4f, 0x44, 0x53, 0x10, 0x8d, 0x27, 0x12, 0x24, 0x0a, 0x1f, 0x53, 0x43, 0x5f, 0x42, 0x55, 0x59,
	0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x55, 0x4c, 0x54, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x5f, 0x47, 0x4f, 0x4f, 0x44, 0x53, 0x10, 0x8e, 0x27, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_id_proto_rawDescOnce sync.Once
	file_id_proto_rawDescData = file_id_proto_rawDesc
)

func file_id_proto_rawDescGZIP() []byte {
	file_id_proto_rawDescOnce.Do(func() {
		file_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_id_proto_rawDescData)
	})
	return file_id_proto_rawDescData
}

var file_id_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_id_proto_goTypes = []interface{}{
	(MessageID)(0), // 0: MessageID
}
var file_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_id_proto_init() }
func file_id_proto_init() {
	if File_id_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_id_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_id_proto_goTypes,
		DependencyIndexes: file_id_proto_depIdxs,
		EnumInfos:         file_id_proto_enumTypes,
	}.Build()
	File_id_proto = out.File
	file_id_proto_rawDesc = nil
	file_id_proto_goTypes = nil
	file_id_proto_depIdxs = nil
}