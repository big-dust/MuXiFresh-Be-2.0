// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: intro.proto

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

// 返回某组文案
type GroupIntroReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName string `protobuf:"bytes,1,opt,name=GroupName,proto3" json:"GroupName,omitempty"`
}

func (x *GroupIntroReq) Reset() {
	*x = GroupIntroReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupIntroReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupIntroReq) ProtoMessage() {}

func (x *GroupIntroReq) ProtoReflect() protoreflect.Message {
	mi := &file_intro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupIntroReq.ProtoReflect.Descriptor instead.
func (*GroupIntroReq) Descriptor() ([]byte, []int) {
	return file_intro_proto_rawDescGZIP(), []int{0}
}

func (x *GroupIntroReq) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

type GroupIntroResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intro string `protobuf:"bytes,1,opt,name=Intro,proto3" json:"Intro,omitempty"`
}

func (x *GroupIntroResp) Reset() {
	*x = GroupIntroResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupIntroResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupIntroResp) ProtoMessage() {}

func (x *GroupIntroResp) ProtoReflect() protoreflect.Message {
	mi := &file_intro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupIntroResp.ProtoReflect.Descriptor instead.
func (*GroupIntroResp) Descriptor() ([]byte, []int) {
	return file_intro_proto_rawDescGZIP(), []int{1}
}

func (x *GroupIntroResp) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

// 返回招募信息
type RecruitInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *RecruitInfoReq) Reset() {
	*x = RecruitInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intro_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecruitInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecruitInfoReq) ProtoMessage() {}

func (x *RecruitInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_intro_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecruitInfoReq.ProtoReflect.Descriptor instead.
func (*RecruitInfoReq) Descriptor() ([]byte, []int) {
	return file_intro_proto_rawDescGZIP(), []int{2}
}

func (x *RecruitInfoReq) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type RecruitInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *RecruitInfoResp) Reset() {
	*x = RecruitInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intro_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecruitInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecruitInfoResp) ProtoMessage() {}

func (x *RecruitInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_intro_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecruitInfoResp.ProtoReflect.Descriptor instead.
func (*RecruitInfoResp) Descriptor() ([]byte, []int) {
	return file_intro_proto_rawDescGZIP(), []int{3}
}

func (x *RecruitInfoResp) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

var File_intro_proto protoreflect.FileDescriptor

var file_intro_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69,
	0x6e, 0x74, 0x72, 0x6f, 0x22, 0x2d, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x74,
	0x72, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x74, 0x72,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x22, 0x22, 0x0a, 0x0e, 0x52,
	0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22,
	0x23, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x32, 0x8c, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x52, 0x65, 0x63,
	0x72, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_intro_proto_rawDescOnce sync.Once
	file_intro_proto_rawDescData = file_intro_proto_rawDesc
)

func file_intro_proto_rawDescGZIP() []byte {
	file_intro_proto_rawDescOnce.Do(func() {
		file_intro_proto_rawDescData = protoimpl.X.CompressGZIP(file_intro_proto_rawDescData)
	})
	return file_intro_proto_rawDescData
}

var file_intro_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_intro_proto_goTypes = []interface{}{
	(*GroupIntroReq)(nil),   // 0: intro.GroupIntroReq
	(*GroupIntroResp)(nil),  // 1: intro.GroupIntroResp
	(*RecruitInfoReq)(nil),  // 2: intro.RecruitInfoReq
	(*RecruitInfoResp)(nil), // 3: intro.RecruitInfoResp
}
var file_intro_proto_depIdxs = []int32{
	0, // 0: intro.IntroClient.GetGroupIntro:input_type -> intro.GroupIntroReq
	2, // 1: intro.IntroClient.GetRecruitInfo:input_type -> intro.RecruitInfoReq
	1, // 2: intro.IntroClient.GetGroupIntro:output_type -> intro.GroupIntroResp
	3, // 3: intro.IntroClient.GetRecruitInfo:output_type -> intro.RecruitInfoResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_intro_proto_init() }
func file_intro_proto_init() {
	if File_intro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_intro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupIntroReq); i {
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
		file_intro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupIntroResp); i {
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
		file_intro_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecruitInfoReq); i {
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
		file_intro_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecruitInfoResp); i {
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
			RawDescriptor: file_intro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_intro_proto_goTypes,
		DependencyIndexes: file_intro_proto_depIdxs,
		MessageInfos:      file_intro_proto_msgTypes,
	}.Build()
	File_intro_proto = out.File
	file_intro_proto_rawDesc = nil
	file_intro_proto_goTypes = nil
	file_intro_proto_depIdxs = nil
}