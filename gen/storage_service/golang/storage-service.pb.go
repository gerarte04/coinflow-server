// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: storage-service.proto

package golang

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Transaction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Target        string                 `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Type          string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Category      string                 `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Cost          float64                `protobuf:"fixed64,7,opt,name=cost,proto3" json:"cost,omitempty"`
	Timestamp     string                 `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_storage_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_storage_service_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Transaction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Transaction) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Transaction) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Transaction) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type GetTransactionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxId          string                 `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionRequest) Reset() {
	*x = GetTransactionRequest{}
	mi := &file_storage_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRequest) ProtoMessage() {}

func (x *GetTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return file_storage_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTransactionRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

type GetTransactionsInPeriodRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Begin         string                 `protobuf:"bytes,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End           string                 `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionsInPeriodRequest) Reset() {
	*x = GetTransactionsInPeriodRequest{}
	mi := &file_storage_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionsInPeriodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsInPeriodRequest) ProtoMessage() {}

func (x *GetTransactionsInPeriodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsInPeriodRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsInPeriodRequest) Descriptor() ([]byte, []int) {
	return file_storage_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransactionsInPeriodRequest) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *GetTransactionsInPeriodRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type PostTransactionRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Tx               *Transaction           `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	WithAutoCategory bool                   `protobuf:"varint,2,opt,name=with_auto_category,json=withAutoCategory,proto3" json:"with_auto_category,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *PostTransactionRequest) Reset() {
	*x = PostTransactionRequest{}
	mi := &file_storage_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTransactionRequest) ProtoMessage() {}

func (x *PostTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTransactionRequest.ProtoReflect.Descriptor instead.
func (*PostTransactionRequest) Descriptor() ([]byte, []int) {
	return file_storage_service_proto_rawDescGZIP(), []int{3}
}

func (x *PostTransactionRequest) GetTx() *Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *PostTransactionRequest) GetWithAutoCategory() bool {
	if x != nil {
		return x.WithAutoCategory
	}
	return false
}

type GetTransactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tx            *Transaction           `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionResponse) Reset() {
	*x = GetTransactionResponse{}
	mi := &file_storage_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResponse) ProtoMessage() {}

func (x *GetTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionResponse) Descriptor() ([]byte, []int) {
	return file_storage_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransactionResponse) GetTx() *Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

type GetTransactionsInPeriodResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Txs           []*Transaction         `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransactionsInPeriodResponse) Reset() {
	*x = GetTransactionsInPeriodResponse{}
	mi := &file_storage_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionsInPeriodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsInPeriodResponse) ProtoMessage() {}

func (x *GetTransactionsInPeriodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsInPeriodResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionsInPeriodResponse) Descriptor() ([]byte, []int) {
	return file_storage_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetTransactionsInPeriodResponse) GetTxs() []*Transaction {
	if x != nil {
		return x.Txs
	}
	return nil
}

type PostTransactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxId          string                 `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostTransactionResponse) Reset() {
	*x = PostTransactionResponse{}
	mi := &file_storage_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTransactionResponse) ProtoMessage() {}

func (x *PostTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTransactionResponse.ProtoReflect.Descriptor instead.
func (*PostTransactionResponse) Descriptor() ([]byte, []int) {
	return file_storage_service_proto_rawDescGZIP(), []int{6}
}

func (x *PostTransactionResponse) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

var File_storage_service_proto protoreflect.FileDescriptor

const file_storage_service_proto_rawDesc = "" +
	"\n" +
	"\x15storage-service.proto\x12\x0fstorage_service\x1a\x1cgoogle/api/annotations.proto\"\xb9\x01\n" +
	"\vTransaction\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06target\x18\x03 \x01(\tR\x06target\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x12\n" +
	"\x04type\x18\x05 \x01(\tR\x04type\x12\x1a\n" +
	"\bcategory\x18\x06 \x01(\tR\bcategory\x12\x12\n" +
	"\x04cost\x18\a \x01(\x01R\x04cost\x12\x1c\n" +
	"\ttimestamp\x18\b \x01(\tR\ttimestamp\",\n" +
	"\x15GetTransactionRequest\x12\x13\n" +
	"\x05tx_id\x18\x01 \x01(\tR\x04txId\"H\n" +
	"\x1eGetTransactionsInPeriodRequest\x12\x14\n" +
	"\x05begin\x18\x01 \x01(\tR\x05begin\x12\x10\n" +
	"\x03end\x18\x02 \x01(\tR\x03end\"t\n" +
	"\x16PostTransactionRequest\x12,\n" +
	"\x02tx\x18\x01 \x01(\v2\x1c.storage_service.TransactionR\x02tx\x12,\n" +
	"\x12with_auto_category\x18\x02 \x01(\bR\x10withAutoCategory\"F\n" +
	"\x16GetTransactionResponse\x12,\n" +
	"\x02tx\x18\x01 \x01(\v2\x1c.storage_service.TransactionR\x02tx\"Q\n" +
	"\x1fGetTransactionsInPeriodResponse\x12.\n" +
	"\x03txs\x18\x01 \x03(\v2\x1c.storage_service.TransactionR\x03txs\".\n" +
	"\x17PostTransactionResponse\x12\x13\n" +
	"\x05tx_id\x18\x01 \x01(\tR\x04txId2\xaa\x03\n" +
	"\aStorage\x12\x85\x01\n" +
	"\x0eGetTransaction\x12&.storage_service.GetTransactionRequest\x1a'.storage_service.GetTransactionResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/transaction/id/{tx_id}\x12\x9c\x01\n" +
	"\x17GetTransactionsInPeriod\x12/.storage_service.GetTransactionsInPeriodRequest\x1a0.storage_service.GetTransactionsInPeriodResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x16/v1/transaction/period\x12x\n" +
	"\x0fPostTransaction\x12'.storage_service.PostTransactionRequest\x1a(.storage_service.PostTransactionResponse\"\x12\x82\xd3\xe4\x93\x02\f\"\n" +
	"/v1/commitB\x1cZ\x1agen/storage_service/golangb\x06proto3"

var (
	file_storage_service_proto_rawDescOnce sync.Once
	file_storage_service_proto_rawDescData []byte
)

func file_storage_service_proto_rawDescGZIP() []byte {
	file_storage_service_proto_rawDescOnce.Do(func() {
		file_storage_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_service_proto_rawDesc), len(file_storage_service_proto_rawDesc)))
	})
	return file_storage_service_proto_rawDescData
}

var file_storage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_storage_service_proto_goTypes = []any{
	(*Transaction)(nil),                     // 0: storage_service.Transaction
	(*GetTransactionRequest)(nil),           // 1: storage_service.GetTransactionRequest
	(*GetTransactionsInPeriodRequest)(nil),  // 2: storage_service.GetTransactionsInPeriodRequest
	(*PostTransactionRequest)(nil),          // 3: storage_service.PostTransactionRequest
	(*GetTransactionResponse)(nil),          // 4: storage_service.GetTransactionResponse
	(*GetTransactionsInPeriodResponse)(nil), // 5: storage_service.GetTransactionsInPeriodResponse
	(*PostTransactionResponse)(nil),         // 6: storage_service.PostTransactionResponse
}
var file_storage_service_proto_depIdxs = []int32{
	0, // 0: storage_service.PostTransactionRequest.tx:type_name -> storage_service.Transaction
	0, // 1: storage_service.GetTransactionResponse.tx:type_name -> storage_service.Transaction
	0, // 2: storage_service.GetTransactionsInPeriodResponse.txs:type_name -> storage_service.Transaction
	1, // 3: storage_service.Storage.GetTransaction:input_type -> storage_service.GetTransactionRequest
	2, // 4: storage_service.Storage.GetTransactionsInPeriod:input_type -> storage_service.GetTransactionsInPeriodRequest
	3, // 5: storage_service.Storage.PostTransaction:input_type -> storage_service.PostTransactionRequest
	4, // 6: storage_service.Storage.GetTransaction:output_type -> storage_service.GetTransactionResponse
	5, // 7: storage_service.Storage.GetTransactionsInPeriod:output_type -> storage_service.GetTransactionsInPeriodResponse
	6, // 8: storage_service.Storage.PostTransaction:output_type -> storage_service.PostTransactionResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_service_proto_init() }
func file_storage_service_proto_init() {
	if File_storage_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_service_proto_rawDesc), len(file_storage_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_service_proto_goTypes,
		DependencyIndexes: file_storage_service_proto_depIdxs,
		MessageInfos:      file_storage_service_proto_msgTypes,
	}.Build()
	File_storage_service_proto = out.File
	file_storage_service_proto_goTypes = nil
	file_storage_service_proto_depIdxs = nil
}
