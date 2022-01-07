//
// Copyright 2018 The Feast Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: feast/types/Value.proto

package types

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

type Null int32

const (
	Null_NULL Null = 0
)

// Enum value maps for Null.
var (
	Null_name = map[int32]string{
		0: "NULL",
	}
	Null_value = map[string]int32{
		"NULL": 0,
	}
)

func (x Null) Enum() *Null {
	p := new(Null)
	*p = x
	return p
}

func (x Null) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Null) Descriptor() protoreflect.EnumDescriptor {
	return file_feast_types_Value_proto_enumTypes[0].Descriptor()
}

func (Null) Type() protoreflect.EnumType {
	return &file_feast_types_Value_proto_enumTypes[0]
}

func (x Null) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Null.Descriptor instead.
func (Null) EnumDescriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{0}
}

type ValueType_Enum int32

const (
	ValueType_INVALID             ValueType_Enum = 0
	ValueType_BYTES               ValueType_Enum = 1
	ValueType_STRING              ValueType_Enum = 2
	ValueType_INT32               ValueType_Enum = 3
	ValueType_INT64               ValueType_Enum = 4
	ValueType_DOUBLE              ValueType_Enum = 5
	ValueType_FLOAT               ValueType_Enum = 6
	ValueType_BOOL                ValueType_Enum = 7
	ValueType_UNIX_TIMESTAMP      ValueType_Enum = 8
	ValueType_BYTES_LIST          ValueType_Enum = 11
	ValueType_STRING_LIST         ValueType_Enum = 12
	ValueType_INT32_LIST          ValueType_Enum = 13
	ValueType_INT64_LIST          ValueType_Enum = 14
	ValueType_DOUBLE_LIST         ValueType_Enum = 15
	ValueType_FLOAT_LIST          ValueType_Enum = 16
	ValueType_BOOL_LIST           ValueType_Enum = 17
	ValueType_UNIX_TIMESTAMP_LIST ValueType_Enum = 18
	ValueType_NULL                ValueType_Enum = 19
)

// Enum value maps for ValueType_Enum.
var (
	ValueType_Enum_name = map[int32]string{
		0:  "INVALID",
		1:  "BYTES",
		2:  "STRING",
		3:  "INT32",
		4:  "INT64",
		5:  "DOUBLE",
		6:  "FLOAT",
		7:  "BOOL",
		8:  "UNIX_TIMESTAMP",
		11: "BYTES_LIST",
		12: "STRING_LIST",
		13: "INT32_LIST",
		14: "INT64_LIST",
		15: "DOUBLE_LIST",
		16: "FLOAT_LIST",
		17: "BOOL_LIST",
		18: "UNIX_TIMESTAMP_LIST",
		19: "NULL",
	}
	ValueType_Enum_value = map[string]int32{
		"INVALID":             0,
		"BYTES":               1,
		"STRING":              2,
		"INT32":               3,
		"INT64":               4,
		"DOUBLE":              5,
		"FLOAT":               6,
		"BOOL":                7,
		"UNIX_TIMESTAMP":      8,
		"BYTES_LIST":          11,
		"STRING_LIST":         12,
		"INT32_LIST":          13,
		"INT64_LIST":          14,
		"DOUBLE_LIST":         15,
		"FLOAT_LIST":          16,
		"BOOL_LIST":           17,
		"UNIX_TIMESTAMP_LIST": 18,
		"NULL":                19,
	}
)

func (x ValueType_Enum) Enum() *ValueType_Enum {
	p := new(ValueType_Enum)
	*p = x
	return p
}

func (x ValueType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValueType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_feast_types_Value_proto_enumTypes[1].Descriptor()
}

func (ValueType_Enum) Type() protoreflect.EnumType {
	return &file_feast_types_Value_proto_enumTypes[1]
}

func (x ValueType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValueType_Enum.Descriptor instead.
func (ValueType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{0, 0}
}

type ValueType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValueType) Reset() {
	*x = ValueType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueType) ProtoMessage() {}

func (x *ValueType) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueType.ProtoReflect.Descriptor instead.
func (*ValueType) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{0}
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ValueType is referenced by the metadata types, FeatureInfo and EntityInfo.
	// The enum values do not have to match the oneof val field ids, but they should.
	// In JSON "*_val" field can be omitted
	//
	// Types that are assignable to Val:
	//	*Value_BytesVal
	//	*Value_StringVal
	//	*Value_Int32Val
	//	*Value_Int64Val
	//	*Value_DoubleVal
	//	*Value_FloatVal
	//	*Value_BoolVal
	//	*Value_UnixTimestampVal
	//	*Value_BytesListVal
	//	*Value_StringListVal
	//	*Value_Int32ListVal
	//	*Value_Int64ListVal
	//	*Value_DoubleListVal
	//	*Value_FloatListVal
	//	*Value_BoolListVal
	//	*Value_UnixTimestampListVal
	//	*Value_NullVal
	Val isValue_Val `protobuf_oneof:"val"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{1}
}

func (m *Value) GetVal() isValue_Val {
	if m != nil {
		return m.Val
	}
	return nil
}

func (x *Value) GetBytesVal() []byte {
	if x, ok := x.GetVal().(*Value_BytesVal); ok {
		return x.BytesVal
	}
	return nil
}

func (x *Value) GetStringVal() string {
	if x, ok := x.GetVal().(*Value_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (x *Value) GetInt32Val() int32 {
	if x, ok := x.GetVal().(*Value_Int32Val); ok {
		return x.Int32Val
	}
	return 0
}

func (x *Value) GetInt64Val() int64 {
	if x, ok := x.GetVal().(*Value_Int64Val); ok {
		return x.Int64Val
	}
	return 0
}

func (x *Value) GetDoubleVal() float64 {
	if x, ok := x.GetVal().(*Value_DoubleVal); ok {
		return x.DoubleVal
	}
	return 0
}

func (x *Value) GetFloatVal() float32 {
	if x, ok := x.GetVal().(*Value_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

func (x *Value) GetBoolVal() bool {
	if x, ok := x.GetVal().(*Value_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (x *Value) GetUnixTimestampVal() int64 {
	if x, ok := x.GetVal().(*Value_UnixTimestampVal); ok {
		return x.UnixTimestampVal
	}
	return 0
}

func (x *Value) GetBytesListVal() *BytesList {
	if x, ok := x.GetVal().(*Value_BytesListVal); ok {
		return x.BytesListVal
	}
	return nil
}

func (x *Value) GetStringListVal() *StringList {
	if x, ok := x.GetVal().(*Value_StringListVal); ok {
		return x.StringListVal
	}
	return nil
}

func (x *Value) GetInt32ListVal() *Int32List {
	if x, ok := x.GetVal().(*Value_Int32ListVal); ok {
		return x.Int32ListVal
	}
	return nil
}

func (x *Value) GetInt64ListVal() *Int64List {
	if x, ok := x.GetVal().(*Value_Int64ListVal); ok {
		return x.Int64ListVal
	}
	return nil
}

func (x *Value) GetDoubleListVal() *DoubleList {
	if x, ok := x.GetVal().(*Value_DoubleListVal); ok {
		return x.DoubleListVal
	}
	return nil
}

func (x *Value) GetFloatListVal() *FloatList {
	if x, ok := x.GetVal().(*Value_FloatListVal); ok {
		return x.FloatListVal
	}
	return nil
}

func (x *Value) GetBoolListVal() *BoolList {
	if x, ok := x.GetVal().(*Value_BoolListVal); ok {
		return x.BoolListVal
	}
	return nil
}

func (x *Value) GetUnixTimestampListVal() *Int64List {
	if x, ok := x.GetVal().(*Value_UnixTimestampListVal); ok {
		return x.UnixTimestampListVal
	}
	return nil
}

func (x *Value) GetNullVal() Null {
	if x, ok := x.GetVal().(*Value_NullVal); ok {
		return x.NullVal
	}
	return Null_NULL
}

type isValue_Val interface {
	isValue_Val()
}

type Value_BytesVal struct {
	BytesVal []byte `protobuf:"bytes,1,opt,name=bytes_val,json=bytesVal,proto3,oneof"`
}

type Value_StringVal struct {
	StringVal string `protobuf:"bytes,2,opt,name=string_val,json=stringVal,proto3,oneof"`
}

type Value_Int32Val struct {
	Int32Val int32 `protobuf:"varint,3,opt,name=int32_val,json=int32Val,proto3,oneof"`
}

type Value_Int64Val struct {
	Int64Val int64 `protobuf:"varint,4,opt,name=int64_val,json=int64Val,proto3,oneof"`
}

type Value_DoubleVal struct {
	DoubleVal float64 `protobuf:"fixed64,5,opt,name=double_val,json=doubleVal,proto3,oneof"`
}

type Value_FloatVal struct {
	FloatVal float32 `protobuf:"fixed32,6,opt,name=float_val,json=floatVal,proto3,oneof"`
}

type Value_BoolVal struct {
	BoolVal bool `protobuf:"varint,7,opt,name=bool_val,json=boolVal,proto3,oneof"`
}

type Value_UnixTimestampVal struct {
	UnixTimestampVal int64 `protobuf:"varint,8,opt,name=unix_timestamp_val,json=unixTimestampVal,proto3,oneof"`
}

type Value_BytesListVal struct {
	BytesListVal *BytesList `protobuf:"bytes,11,opt,name=bytes_list_val,json=bytesListVal,proto3,oneof"`
}

type Value_StringListVal struct {
	StringListVal *StringList `protobuf:"bytes,12,opt,name=string_list_val,json=stringListVal,proto3,oneof"`
}

type Value_Int32ListVal struct {
	Int32ListVal *Int32List `protobuf:"bytes,13,opt,name=int32_list_val,json=int32ListVal,proto3,oneof"`
}

type Value_Int64ListVal struct {
	Int64ListVal *Int64List `protobuf:"bytes,14,opt,name=int64_list_val,json=int64ListVal,proto3,oneof"`
}

type Value_DoubleListVal struct {
	DoubleListVal *DoubleList `protobuf:"bytes,15,opt,name=double_list_val,json=doubleListVal,proto3,oneof"`
}

type Value_FloatListVal struct {
	FloatListVal *FloatList `protobuf:"bytes,16,opt,name=float_list_val,json=floatListVal,proto3,oneof"`
}

type Value_BoolListVal struct {
	BoolListVal *BoolList `protobuf:"bytes,17,opt,name=bool_list_val,json=boolListVal,proto3,oneof"`
}

type Value_UnixTimestampListVal struct {
	UnixTimestampListVal *Int64List `protobuf:"bytes,18,opt,name=unix_timestamp_list_val,json=unixTimestampListVal,proto3,oneof"`
}

type Value_NullVal struct {
	NullVal Null `protobuf:"varint,19,opt,name=null_val,json=nullVal,proto3,enum=feast.types.Null,oneof"`
}

func (*Value_BytesVal) isValue_Val() {}

func (*Value_StringVal) isValue_Val() {}

func (*Value_Int32Val) isValue_Val() {}

func (*Value_Int64Val) isValue_Val() {}

func (*Value_DoubleVal) isValue_Val() {}

func (*Value_FloatVal) isValue_Val() {}

func (*Value_BoolVal) isValue_Val() {}

func (*Value_UnixTimestampVal) isValue_Val() {}

func (*Value_BytesListVal) isValue_Val() {}

func (*Value_StringListVal) isValue_Val() {}

func (*Value_Int32ListVal) isValue_Val() {}

func (*Value_Int64ListVal) isValue_Val() {}

func (*Value_DoubleListVal) isValue_Val() {}

func (*Value_FloatListVal) isValue_Val() {}

func (*Value_BoolListVal) isValue_Val() {}

func (*Value_UnixTimestampListVal) isValue_Val() {}

func (*Value_NullVal) isValue_Val() {}

type BytesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val [][]byte `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
}

func (x *BytesList) Reset() {
	*x = BytesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesList) ProtoMessage() {}

func (x *BytesList) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesList.ProtoReflect.Descriptor instead.
func (*BytesList) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{2}
}

func (x *BytesList) GetVal() [][]byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type StringList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
}

func (x *StringList) Reset() {
	*x = StringList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringList) ProtoMessage() {}

func (x *StringList) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringList.ProtoReflect.Descriptor instead.
func (*StringList) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{3}
}

func (x *StringList) GetVal() []string {
	if x != nil {
		return x.Val
	}
	return nil
}

type Int32List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []int32 `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *Int32List) Reset() {
	*x = Int32List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32List) ProtoMessage() {}

func (x *Int32List) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32List.ProtoReflect.Descriptor instead.
func (*Int32List) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{4}
}

func (x *Int32List) GetVal() []int32 {
	if x != nil {
		return x.Val
	}
	return nil
}

type Int64List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []int64 `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *Int64List) Reset() {
	*x = Int64List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64List) ProtoMessage() {}

func (x *Int64List) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64List.ProtoReflect.Descriptor instead.
func (*Int64List) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{5}
}

func (x *Int64List) GetVal() []int64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type DoubleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []float64 `protobuf:"fixed64,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *DoubleList) Reset() {
	*x = DoubleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleList) ProtoMessage() {}

func (x *DoubleList) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleList.ProtoReflect.Descriptor instead.
func (*DoubleList) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{6}
}

func (x *DoubleList) GetVal() []float64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type FloatList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []float32 `protobuf:"fixed32,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *FloatList) Reset() {
	*x = FloatList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatList) ProtoMessage() {}

func (x *FloatList) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatList.ProtoReflect.Descriptor instead.
func (*FloatList) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{7}
}

func (x *FloatList) GetVal() []float32 {
	if x != nil {
		return x.Val
	}
	return nil
}

type BoolList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []bool `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *BoolList) Reset() {
	*x = BoolList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolList) ProtoMessage() {}

func (x *BoolList) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolList.ProtoReflect.Descriptor instead.
func (*BoolList) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{8}
}

func (x *BoolList) GetVal() []bool {
	if x != nil {
		return x.Val
	}
	return nil
}

// This is to avoid an issue of being unable to specify `repeated value` in oneofs or maps
// In JSON "val" field can be omitted
type RepeatedValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []*Value `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
}

func (x *RepeatedValue) Reset() {
	*x = RepeatedValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_types_Value_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedValue) ProtoMessage() {}

func (x *RepeatedValue) ProtoReflect() protoreflect.Message {
	mi := &file_feast_types_Value_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedValue.ProtoReflect.Descriptor instead.
func (*RepeatedValue) Descriptor() ([]byte, []int) {
	return file_feast_types_Value_proto_rawDescGZIP(), []int{9}
}

func (x *RepeatedValue) GetVal() []*Value {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_feast_types_Value_proto protoreflect.FileDescriptor

var file_feast_types_Value_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x97, 0x02, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x89, 0x02, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x59,
	0x54, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c,
	0x45, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x06, 0x12, 0x08,
	0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x49, 0x58,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a,
	0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0c, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0d, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0e, 0x12, 0x0f, 0x0a,
	0x0b, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0f, 0x12, 0x0e,
	0x0a, 0x0a, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x10, 0x12, 0x0d,
	0x0a, 0x09, 0x42, 0x4f, 0x4f, 0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x11, 0x12, 0x17, 0x0a,
	0x13, 0x55, 0x4e, 0x49, 0x58, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f,
	0x4c, 0x49, 0x53, 0x54, 0x10, 0x12, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x13,
	0x22, 0xdd, 0x06, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x09,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x08,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x12, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x10, 0x75, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x56, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x41, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x41, 0x0a, 0x0f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x0e, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x0d, 0x62,
	0x6f, 0x6f, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x6f, 0x6f,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x4f, 0x0a, 0x17, 0x75, 0x6e, 0x69, 0x78,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x65, 0x61, 0x73,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x14, 0x75, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x08, 0x6e, 0x75, 0x6c,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x66, 0x65,
	0x61, 0x73, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x48, 0x00,
	0x52, 0x07, 0x6e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x42, 0x05, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x1d, 0x0a, 0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22,
	0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22,
	0x1d, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1d,
	0x0a, 0x09, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1e, 0x0a,
	0x0a, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1d, 0x0a,
	0x09, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1c, 0x0a, 0x08,
	0x42, 0x6f, 0x6f, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x35, 0x0a, 0x0d, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x2a, 0x10, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c,
	0x4c, 0x10, 0x00, 0x42, 0x55, 0x0a, 0x11, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66,
	0x65, 0x61, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_feast_types_Value_proto_rawDescOnce sync.Once
	file_feast_types_Value_proto_rawDescData = file_feast_types_Value_proto_rawDesc
)

func file_feast_types_Value_proto_rawDescGZIP() []byte {
	file_feast_types_Value_proto_rawDescOnce.Do(func() {
		file_feast_types_Value_proto_rawDescData = protoimpl.X.CompressGZIP(file_feast_types_Value_proto_rawDescData)
	})
	return file_feast_types_Value_proto_rawDescData
}

var file_feast_types_Value_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_feast_types_Value_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_feast_types_Value_proto_goTypes = []interface{}{
	(Null)(0),             // 0: feast.types.Null
	(ValueType_Enum)(0),   // 1: feast.types.ValueType.Enum
	(*ValueType)(nil),     // 2: feast.types.ValueType
	(*Value)(nil),         // 3: feast.types.Value
	(*BytesList)(nil),     // 4: feast.types.BytesList
	(*StringList)(nil),    // 5: feast.types.StringList
	(*Int32List)(nil),     // 6: feast.types.Int32List
	(*Int64List)(nil),     // 7: feast.types.Int64List
	(*DoubleList)(nil),    // 8: feast.types.DoubleList
	(*FloatList)(nil),     // 9: feast.types.FloatList
	(*BoolList)(nil),      // 10: feast.types.BoolList
	(*RepeatedValue)(nil), // 11: feast.types.RepeatedValue
}
var file_feast_types_Value_proto_depIdxs = []int32{
	4,  // 0: feast.types.Value.bytes_list_val:type_name -> feast.types.BytesList
	5,  // 1: feast.types.Value.string_list_val:type_name -> feast.types.StringList
	6,  // 2: feast.types.Value.int32_list_val:type_name -> feast.types.Int32List
	7,  // 3: feast.types.Value.int64_list_val:type_name -> feast.types.Int64List
	8,  // 4: feast.types.Value.double_list_val:type_name -> feast.types.DoubleList
	9,  // 5: feast.types.Value.float_list_val:type_name -> feast.types.FloatList
	10, // 6: feast.types.Value.bool_list_val:type_name -> feast.types.BoolList
	7,  // 7: feast.types.Value.unix_timestamp_list_val:type_name -> feast.types.Int64List
	0,  // 8: feast.types.Value.null_val:type_name -> feast.types.Null
	3,  // 9: feast.types.RepeatedValue.val:type_name -> feast.types.Value
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_feast_types_Value_proto_init() }
func file_feast_types_Value_proto_init() {
	if File_feast_types_Value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feast_types_Value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueType); i {
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
		file_feast_types_Value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_feast_types_Value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesList); i {
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
		file_feast_types_Value_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringList); i {
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
		file_feast_types_Value_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32List); i {
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
		file_feast_types_Value_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64List); i {
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
		file_feast_types_Value_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleList); i {
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
		file_feast_types_Value_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloatList); i {
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
		file_feast_types_Value_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolList); i {
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
		file_feast_types_Value_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedValue); i {
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
	file_feast_types_Value_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Value_BytesVal)(nil),
		(*Value_StringVal)(nil),
		(*Value_Int32Val)(nil),
		(*Value_Int64Val)(nil),
		(*Value_DoubleVal)(nil),
		(*Value_FloatVal)(nil),
		(*Value_BoolVal)(nil),
		(*Value_UnixTimestampVal)(nil),
		(*Value_BytesListVal)(nil),
		(*Value_StringListVal)(nil),
		(*Value_Int32ListVal)(nil),
		(*Value_Int64ListVal)(nil),
		(*Value_DoubleListVal)(nil),
		(*Value_FloatListVal)(nil),
		(*Value_BoolListVal)(nil),
		(*Value_UnixTimestampListVal)(nil),
		(*Value_NullVal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feast_types_Value_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feast_types_Value_proto_goTypes,
		DependencyIndexes: file_feast_types_Value_proto_depIdxs,
		EnumInfos:         file_feast_types_Value_proto_enumTypes,
		MessageInfos:      file_feast_types_Value_proto_msgTypes,
	}.Build()
	File_feast_types_Value_proto = out.File
	file_feast_types_Value_proto_rawDesc = nil
	file_feast_types_Value_proto_goTypes = nil
	file_feast_types_Value_proto_depIdxs = nil
}
