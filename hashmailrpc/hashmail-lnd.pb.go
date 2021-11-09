// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: hashmail-lnd.proto

package hashmailrpc

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

type LndAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LndAuth) Reset() {
	*x = LndAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LndAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LndAuth) ProtoMessage() {}

func (x *LndAuth) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LndAuth.ProtoReflect.Descriptor instead.
func (*LndAuth) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{0}
}

type CipherBoxAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A description of the stream one is attempting to initialize.
	Desc *CipherBoxDesc `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	// Types that are assignable to Auth:
	//	*CipherBoxAuth_LndAuth
	Auth isCipherBoxAuth_Auth `protobuf_oneof:"auth"`
}

func (x *CipherBoxAuth) Reset() {
	*x = CipherBoxAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherBoxAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherBoxAuth) ProtoMessage() {}

func (x *CipherBoxAuth) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherBoxAuth.ProtoReflect.Descriptor instead.
func (*CipherBoxAuth) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{1}
}

func (x *CipherBoxAuth) GetDesc() *CipherBoxDesc {
	if x != nil {
		return x.Desc
	}
	return nil
}

func (m *CipherBoxAuth) GetAuth() isCipherBoxAuth_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (x *CipherBoxAuth) GetLndAuth() *LndAuth {
	if x, ok := x.GetAuth().(*CipherBoxAuth_LndAuth); ok {
		return x.LndAuth
	}
	return nil
}

type isCipherBoxAuth_Auth interface {
	isCipherBoxAuth_Auth()
}

type CipherBoxAuth_LndAuth struct {
	LndAuth *LndAuth `protobuf:"bytes,2,opt,name=lnd_auth,json=lndAuth,proto3,oneof"`
}

func (*CipherBoxAuth_LndAuth) isCipherBoxAuth_Auth() {}

type DelCipherBoxResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelCipherBoxResp) Reset() {
	*x = DelCipherBoxResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelCipherBoxResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelCipherBoxResp) ProtoMessage() {}

func (x *DelCipherBoxResp) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelCipherBoxResp.ProtoReflect.Descriptor instead.
func (*DelCipherBoxResp) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{2}
}

type CipherChallenge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CipherChallenge) Reset() {
	*x = CipherChallenge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherChallenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherChallenge) ProtoMessage() {}

func (x *CipherChallenge) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherChallenge.ProtoReflect.Descriptor instead.
func (*CipherChallenge) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{3}
}

type CipherError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CipherError) Reset() {
	*x = CipherError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherError) ProtoMessage() {}

func (x *CipherError) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherError.ProtoReflect.Descriptor instead.
func (*CipherError) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{4}
}

type CipherSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc *CipherBoxDesc `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *CipherSuccess) Reset() {
	*x = CipherSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherSuccess) ProtoMessage() {}

func (x *CipherSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherSuccess.ProtoReflect.Descriptor instead.
func (*CipherSuccess) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{5}
}

func (x *CipherSuccess) GetDesc() *CipherBoxDesc {
	if x != nil {
		return x.Desc
	}
	return nil
}

type CipherInitResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resp:
	//	*CipherInitResp_Success
	//	*CipherInitResp_Challenge
	//	*CipherInitResp_Error
	Resp isCipherInitResp_Resp `protobuf_oneof:"resp"`
}

func (x *CipherInitResp) Reset() {
	*x = CipherInitResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherInitResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherInitResp) ProtoMessage() {}

func (x *CipherInitResp) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherInitResp.ProtoReflect.Descriptor instead.
func (*CipherInitResp) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{6}
}

func (m *CipherInitResp) GetResp() isCipherInitResp_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *CipherInitResp) GetSuccess() *CipherSuccess {
	if x, ok := x.GetResp().(*CipherInitResp_Success); ok {
		return x.Success
	}
	return nil
}

func (x *CipherInitResp) GetChallenge() *CipherChallenge {
	if x, ok := x.GetResp().(*CipherInitResp_Challenge); ok {
		return x.Challenge
	}
	return nil
}

func (x *CipherInitResp) GetError() *CipherError {
	if x, ok := x.GetResp().(*CipherInitResp_Error); ok {
		return x.Error
	}
	return nil
}

type isCipherInitResp_Resp interface {
	isCipherInitResp_Resp()
}

type CipherInitResp_Success struct {
	//
	//CipherSuccess is returned if the initialization of the cipher box was
	//successful.
	Success *CipherSuccess `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type CipherInitResp_Challenge struct {
	//
	//CipherChallenge is returned if the authentication mechanism was revoked
	//or needs to be refreshed.
	Challenge *CipherChallenge `protobuf:"bytes,2,opt,name=challenge,proto3,oneof"`
}

type CipherInitResp_Error struct {
	//
	//CipherError is returned if the authentication mechanism failed to
	//validate.
	Error *CipherError `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*CipherInitResp_Success) isCipherInitResp_Resp() {}

func (*CipherInitResp_Challenge) isCipherInitResp_Resp() {}

func (*CipherInitResp_Error) isCipherInitResp_Resp() {}

type CipherBoxDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId []byte `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
}

func (x *CipherBoxDesc) Reset() {
	*x = CipherBoxDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherBoxDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherBoxDesc) ProtoMessage() {}

func (x *CipherBoxDesc) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherBoxDesc.ProtoReflect.Descriptor instead.
func (*CipherBoxDesc) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{7}
}

func (x *CipherBoxDesc) GetStreamId() []byte {
	if x != nil {
		return x.StreamId
	}
	return nil
}

type CipherBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc *CipherBoxDesc `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Msg  []byte         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CipherBox) Reset() {
	*x = CipherBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashmail_lnd_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherBox) ProtoMessage() {}

func (x *CipherBox) ProtoReflect() protoreflect.Message {
	mi := &file_hashmail_lnd_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherBox.ProtoReflect.Descriptor instead.
func (*CipherBox) Descriptor() ([]byte, []int) {
	return file_hashmail_lnd_proto_rawDescGZIP(), []int{8}
}

func (x *CipherBox) GetDesc() *CipherBoxDesc {
	if x != nil {
		return x.Desc
	}
	return nil
}

func (x *CipherBox) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_hashmail_lnd_proto protoreflect.FileDescriptor

var file_hashmail_lnd_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x6c, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70,
	0x63, 0x22, 0x09, 0x0a, 0x07, 0x4c, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x22, 0x7a, 0x0a, 0x0d,
	0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x41, 0x75, 0x74, 0x68, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x42, 0x6f, 0x78, 0x44, 0x65, 0x73, 0x63, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x31, 0x0a,
	0x08, 0x6c, 0x6e, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6e,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68,
	0x42, 0x06, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x43,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0x0a, 0x0f,
	0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22,
	0x0d, 0x0a, 0x0b, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f,
	0x0a, 0x0d, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x42, 0x6f, 0x78, 0x44, 0x65, 0x73, 0x63, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22,
	0xc0, 0x01, 0x0a, 0x0e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48,
	0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61,
	0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x65,
	0x73, 0x70, 0x22, 0x2c, 0x0a, 0x0d, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x22, 0x4d, 0x0a, 0x09, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x42, 0x6f, 0x78, 0x44, 0x65, 0x73, 0x63, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32,
	0xa6, 0x02, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x47, 0x0a, 0x0c,
	0x4e, 0x65, 0x77, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x12, 0x1a, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x42, 0x6f, 0x78, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x1b, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d,
	0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x42, 0x6f, 0x78, 0x12, 0x1a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x41, 0x75, 0x74,
	0x68, 0x1a, 0x1d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x65, 0x6c, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x42, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x1a, 0x1a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69,
	0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x44, 0x65,
	0x73, 0x63, 0x28, 0x01, 0x12, 0x42, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x76, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x1a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x44, 0x65, 0x73, 0x63, 0x1a, 0x16,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2d, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x69, 0x6c, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hashmail_lnd_proto_rawDescOnce sync.Once
	file_hashmail_lnd_proto_rawDescData = file_hashmail_lnd_proto_rawDesc
)

func file_hashmail_lnd_proto_rawDescGZIP() []byte {
	file_hashmail_lnd_proto_rawDescOnce.Do(func() {
		file_hashmail_lnd_proto_rawDescData = protoimpl.X.CompressGZIP(file_hashmail_lnd_proto_rawDescData)
	})
	return file_hashmail_lnd_proto_rawDescData
}

var file_hashmail_lnd_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hashmail_lnd_proto_goTypes = []interface{}{
	(*LndAuth)(nil),          // 0: hashmailrpc.LndAuth
	(*CipherBoxAuth)(nil),    // 1: hashmailrpc.CipherBoxAuth
	(*DelCipherBoxResp)(nil), // 2: hashmailrpc.DelCipherBoxResp
	(*CipherChallenge)(nil),  // 3: hashmailrpc.CipherChallenge
	(*CipherError)(nil),      // 4: hashmailrpc.CipherError
	(*CipherSuccess)(nil),    // 5: hashmailrpc.CipherSuccess
	(*CipherInitResp)(nil),   // 6: hashmailrpc.CipherInitResp
	(*CipherBoxDesc)(nil),    // 7: hashmailrpc.CipherBoxDesc
	(*CipherBox)(nil),        // 8: hashmailrpc.CipherBox
}
var file_hashmail_lnd_proto_depIdxs = []int32{
	7,  // 0: hashmailrpc.CipherBoxAuth.desc:type_name -> hashmailrpc.CipherBoxDesc
	0,  // 1: hashmailrpc.CipherBoxAuth.lnd_auth:type_name -> hashmailrpc.LndAuth
	7,  // 2: hashmailrpc.CipherSuccess.desc:type_name -> hashmailrpc.CipherBoxDesc
	5,  // 3: hashmailrpc.CipherInitResp.success:type_name -> hashmailrpc.CipherSuccess
	3,  // 4: hashmailrpc.CipherInitResp.challenge:type_name -> hashmailrpc.CipherChallenge
	4,  // 5: hashmailrpc.CipherInitResp.error:type_name -> hashmailrpc.CipherError
	7,  // 6: hashmailrpc.CipherBox.desc:type_name -> hashmailrpc.CipherBoxDesc
	1,  // 7: hashmailrpc.HashMail.NewCipherBox:input_type -> hashmailrpc.CipherBoxAuth
	1,  // 8: hashmailrpc.HashMail.DelCipherBox:input_type -> hashmailrpc.CipherBoxAuth
	8,  // 9: hashmailrpc.HashMail.SendStream:input_type -> hashmailrpc.CipherBox
	7,  // 10: hashmailrpc.HashMail.RecvStream:input_type -> hashmailrpc.CipherBoxDesc
	6,  // 11: hashmailrpc.HashMail.NewCipherBox:output_type -> hashmailrpc.CipherInitResp
	2,  // 12: hashmailrpc.HashMail.DelCipherBox:output_type -> hashmailrpc.DelCipherBoxResp
	7,  // 13: hashmailrpc.HashMail.SendStream:output_type -> hashmailrpc.CipherBoxDesc
	8,  // 14: hashmailrpc.HashMail.RecvStream:output_type -> hashmailrpc.CipherBox
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_hashmail_lnd_proto_init() }
func file_hashmail_lnd_proto_init() {
	if File_hashmail_lnd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hashmail_lnd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LndAuth); i {
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
		file_hashmail_lnd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherBoxAuth); i {
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
		file_hashmail_lnd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelCipherBoxResp); i {
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
		file_hashmail_lnd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherChallenge); i {
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
		file_hashmail_lnd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherError); i {
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
		file_hashmail_lnd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherSuccess); i {
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
		file_hashmail_lnd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherInitResp); i {
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
		file_hashmail_lnd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherBoxDesc); i {
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
		file_hashmail_lnd_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherBox); i {
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
	file_hashmail_lnd_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CipherBoxAuth_LndAuth)(nil),
	}
	file_hashmail_lnd_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*CipherInitResp_Success)(nil),
		(*CipherInitResp_Challenge)(nil),
		(*CipherInitResp_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hashmail_lnd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hashmail_lnd_proto_goTypes,
		DependencyIndexes: file_hashmail_lnd_proto_depIdxs,
		MessageInfos:      file_hashmail_lnd_proto_msgTypes,
	}.Build()
	File_hashmail_lnd_proto = out.File
	file_hashmail_lnd_proto_rawDesc = nil
	file_hashmail_lnd_proto_goTypes = nil
	file_hashmail_lnd_proto_depIdxs = nil
}