// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.27.1
// source: cel/expr/conformance/test/suite.proto

package test

import (
	expr "cel.dev/expr"
	conformance "cel.dev/expr/conformance"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestSuite struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Sections      []*TestSection         `protobuf:"bytes,3,rep,name=sections,proto3" json:"sections,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestSuite) Reset() {
	*x = TestSuite{}
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSuite) ProtoMessage() {}

func (x *TestSuite) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSuite.ProtoReflect.Descriptor instead.
func (*TestSuite) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_suite_proto_rawDescGZIP(), []int{0}
}

func (x *TestSuite) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestSuite) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TestSuite) GetSections() []*TestSection {
	if x != nil {
		return x.Sections
	}
	return nil
}

type TestSection struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tests         []*TestCase            `protobuf:"bytes,3,rep,name=tests,proto3" json:"tests,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestSection) Reset() {
	*x = TestSection{}
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestSection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSection) ProtoMessage() {}

func (x *TestSection) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSection.ProtoReflect.Descriptor instead.
func (*TestSection) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_suite_proto_rawDescGZIP(), []int{1}
}

func (x *TestSection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestSection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TestSection) GetTests() []*TestCase {
	if x != nil {
		return x.Tests
	}
	return nil
}

type TestCase struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Name          string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Expr          string                   `protobuf:"bytes,3,opt,name=expr,proto3" json:"expr,omitempty"`
	Env           *conformance.Environment `protobuf:"bytes,4,opt,name=env,proto3" json:"env,omitempty"`
	InputBindings map[string]*InputValue   `protobuf:"bytes,5,rep,name=input_bindings,json=inputBindings,proto3" json:"input_bindings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	InputContext  *InputContext            `protobuf:"bytes,6,opt,name=input_context,json=inputContext,proto3" json:"input_context,omitempty"`
	Output        *TestOutput              `protobuf:"bytes,7,opt,name=output,proto3" json:"output,omitempty"`
	DeducedType   *expr.Type               `protobuf:"bytes,8,opt,name=deduced_type,json=deducedType,proto3" json:"deduced_type,omitempty"`
	DisableCheck  bool                     `protobuf:"varint,9,opt,name=disable_check,json=disableCheck,proto3" json:"disable_check,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestCase) Reset() {
	*x = TestCase{}
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCase) ProtoMessage() {}

func (x *TestCase) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCase.ProtoReflect.Descriptor instead.
func (*TestCase) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_suite_proto_rawDescGZIP(), []int{2}
}

func (x *TestCase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestCase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TestCase) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *TestCase) GetEnv() *conformance.Environment {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *TestCase) GetInputBindings() map[string]*InputValue {
	if x != nil {
		return x.InputBindings
	}
	return nil
}

func (x *TestCase) GetInputContext() *InputContext {
	if x != nil {
		return x.InputContext
	}
	return nil
}

func (x *TestCase) GetOutput() *TestOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *TestCase) GetDeducedType() *expr.Type {
	if x != nil {
		return x.DeducedType
	}
	return nil
}

func (x *TestCase) GetDisableCheck() bool {
	if x != nil {
		return x.DisableCheck
	}
	return false
}

type InputContext struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to InputContextKind:
	//
	//	*InputContext_ContextMessage
	//	*InputContext_ContextExpr
	InputContextKind isInputContext_InputContextKind `protobuf_oneof:"input_context_kind"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *InputContext) Reset() {
	*x = InputContext{}
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputContext) ProtoMessage() {}

func (x *InputContext) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputContext.ProtoReflect.Descriptor instead.
func (*InputContext) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_suite_proto_rawDescGZIP(), []int{3}
}

func (x *InputContext) GetInputContextKind() isInputContext_InputContextKind {
	if x != nil {
		return x.InputContextKind
	}
	return nil
}

func (x *InputContext) GetContextMessage() *anypb.Any {
	if x != nil {
		if x, ok := x.InputContextKind.(*InputContext_ContextMessage); ok {
			return x.ContextMessage
		}
	}
	return nil
}

func (x *InputContext) GetContextExpr() string {
	if x != nil {
		if x, ok := x.InputContextKind.(*InputContext_ContextExpr); ok {
			return x.ContextExpr
		}
	}
	return ""
}

type isInputContext_InputContextKind interface {
	isInputContext_InputContextKind()
}

type InputContext_ContextMessage struct {
	ContextMessage *anypb.Any `protobuf:"bytes,2,opt,name=context_message,json=contextMessage,proto3,oneof"`
}

type InputContext_ContextExpr struct {
	ContextExpr string `protobuf:"bytes,3,opt,name=context_expr,json=contextExpr,proto3,oneof"`
}

func (*InputContext_ContextMessage) isInputContext_InputContextKind() {}

func (*InputContext_ContextExpr) isInputContext_InputContextKind() {}

type InputValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Kind:
	//
	//	*InputValue_Value
	//	*InputValue_Expr
	Kind          isInputValue_Kind `protobuf_oneof:"kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputValue) Reset() {
	*x = InputValue{}
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputValue) ProtoMessage() {}

func (x *InputValue) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputValue.ProtoReflect.Descriptor instead.
func (*InputValue) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_suite_proto_rawDescGZIP(), []int{4}
}

func (x *InputValue) GetKind() isInputValue_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *InputValue) GetValue() *expr.Value {
	if x != nil {
		if x, ok := x.Kind.(*InputValue_Value); ok {
			return x.Value
		}
	}
	return nil
}

func (x *InputValue) GetExpr() string {
	if x != nil {
		if x, ok := x.Kind.(*InputValue_Expr); ok {
			return x.Expr
		}
	}
	return ""
}

type isInputValue_Kind interface {
	isInputValue_Kind()
}

type InputValue_Value struct {
	Value *expr.Value `protobuf:"bytes,1,opt,name=value,proto3,oneof"`
}

type InputValue_Expr struct {
	Expr string `protobuf:"bytes,2,opt,name=expr,proto3,oneof"`
}

func (*InputValue_Value) isInputValue_Kind() {}

func (*InputValue_Expr) isInputValue_Kind() {}

type TestOutput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to ResultKind:
	//
	//	*TestOutput_ResultValue
	//	*TestOutput_ResultExpr
	//	*TestOutput_EvalError
	//	*TestOutput_Unknown
	ResultKind    isTestOutput_ResultKind `protobuf_oneof:"result_kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestOutput) Reset() {
	*x = TestOutput{}
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOutput) ProtoMessage() {}

func (x *TestOutput) ProtoReflect() protoreflect.Message {
	mi := &file_cel_expr_conformance_test_suite_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestOutput.ProtoReflect.Descriptor instead.
func (*TestOutput) Descriptor() ([]byte, []int) {
	return file_cel_expr_conformance_test_suite_proto_rawDescGZIP(), []int{5}
}

func (x *TestOutput) GetResultKind() isTestOutput_ResultKind {
	if x != nil {
		return x.ResultKind
	}
	return nil
}

func (x *TestOutput) GetResultValue() *expr.Value {
	if x != nil {
		if x, ok := x.ResultKind.(*TestOutput_ResultValue); ok {
			return x.ResultValue
		}
	}
	return nil
}

func (x *TestOutput) GetResultExpr() string {
	if x != nil {
		if x, ok := x.ResultKind.(*TestOutput_ResultExpr); ok {
			return x.ResultExpr
		}
	}
	return ""
}

func (x *TestOutput) GetEvalError() *expr.ErrorSet {
	if x != nil {
		if x, ok := x.ResultKind.(*TestOutput_EvalError); ok {
			return x.EvalError
		}
	}
	return nil
}

func (x *TestOutput) GetUnknown() *expr.UnknownSet {
	if x != nil {
		if x, ok := x.ResultKind.(*TestOutput_Unknown); ok {
			return x.Unknown
		}
	}
	return nil
}

type isTestOutput_ResultKind interface {
	isTestOutput_ResultKind()
}

type TestOutput_ResultValue struct {
	ResultValue *expr.Value `protobuf:"bytes,8,opt,name=result_value,json=resultValue,proto3,oneof"`
}

type TestOutput_ResultExpr struct {
	ResultExpr string `protobuf:"bytes,9,opt,name=result_expr,json=resultExpr,proto3,oneof"`
}

type TestOutput_EvalError struct {
	EvalError *expr.ErrorSet `protobuf:"bytes,10,opt,name=eval_error,json=evalError,proto3,oneof"`
}

type TestOutput_Unknown struct {
	Unknown *expr.UnknownSet `protobuf:"bytes,11,opt,name=unknown,proto3,oneof"`
}

func (*TestOutput_ResultValue) isTestOutput_ResultKind() {}

func (*TestOutput_ResultExpr) isTestOutput_ResultKind() {}

func (*TestOutput_EvalError) isTestOutput_ResultKind() {}

func (*TestOutput_Unknown) isTestOutput_ResultKind() {}

var File_cel_expr_conformance_test_suite_proto protoreflect.FileDescriptor

var file_cel_expr_conformance_test_suite_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x69, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70,
	0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x65, 0x6c, 0x2f,
	0x65, 0x78, 0x70, 0x72, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x08, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x7e, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x22,
	0xb6, 0x04, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12, 0x33, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x5d, 0x0a, 0x0e, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x64, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x64,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a,
	0x67, 0x0a, 0x12, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70,
	0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3f, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x70, 0x72, 0x42,
	0x14, 0x0a, 0x12, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x53, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x04,
	0x65, 0x78, 0x70, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x65, 0x78,
	0x70, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x21, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x78,
	0x70, 0x72, 0x12, 0x33, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70,
	0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x09, 0x65, 0x76,
	0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x4f, 0x0a, 0x1d, 0x64, 0x65, 0x76, 0x2e,
	0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x42, 0x0a, 0x53, 0x75, 0x69, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x63, 0x65, 0x6c, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cel_expr_conformance_test_suite_proto_rawDescOnce sync.Once
	file_cel_expr_conformance_test_suite_proto_rawDescData = file_cel_expr_conformance_test_suite_proto_rawDesc
)

func file_cel_expr_conformance_test_suite_proto_rawDescGZIP() []byte {
	file_cel_expr_conformance_test_suite_proto_rawDescOnce.Do(func() {
		file_cel_expr_conformance_test_suite_proto_rawDescData = protoimpl.X.CompressGZIP(file_cel_expr_conformance_test_suite_proto_rawDescData)
	})
	return file_cel_expr_conformance_test_suite_proto_rawDescData
}

var file_cel_expr_conformance_test_suite_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cel_expr_conformance_test_suite_proto_goTypes = []any{
	(*TestSuite)(nil),               // 0: cel.expr.conformance.test.TestSuite
	(*TestSection)(nil),             // 1: cel.expr.conformance.test.TestSection
	(*TestCase)(nil),                // 2: cel.expr.conformance.test.TestCase
	(*InputContext)(nil),            // 3: cel.expr.conformance.test.InputContext
	(*InputValue)(nil),              // 4: cel.expr.conformance.test.InputValue
	(*TestOutput)(nil),              // 5: cel.expr.conformance.test.TestOutput
	nil,                             // 6: cel.expr.conformance.test.TestCase.InputBindingsEntry
	(*conformance.Environment)(nil), // 7: cel.expr.conformance.Environment
	(*expr.Type)(nil),               // 8: cel.expr.Type
	(*anypb.Any)(nil),               // 9: google.protobuf.Any
	(*expr.Value)(nil),              // 10: cel.expr.Value
	(*expr.ErrorSet)(nil),           // 11: cel.expr.ErrorSet
	(*expr.UnknownSet)(nil),         // 12: cel.expr.UnknownSet
}
var file_cel_expr_conformance_test_suite_proto_depIdxs = []int32{
	1,  // 0: cel.expr.conformance.test.TestSuite.sections:type_name -> cel.expr.conformance.test.TestSection
	2,  // 1: cel.expr.conformance.test.TestSection.tests:type_name -> cel.expr.conformance.test.TestCase
	7,  // 2: cel.expr.conformance.test.TestCase.env:type_name -> cel.expr.conformance.Environment
	6,  // 3: cel.expr.conformance.test.TestCase.input_bindings:type_name -> cel.expr.conformance.test.TestCase.InputBindingsEntry
	3,  // 4: cel.expr.conformance.test.TestCase.input_context:type_name -> cel.expr.conformance.test.InputContext
	5,  // 5: cel.expr.conformance.test.TestCase.output:type_name -> cel.expr.conformance.test.TestOutput
	8,  // 6: cel.expr.conformance.test.TestCase.deduced_type:type_name -> cel.expr.Type
	9,  // 7: cel.expr.conformance.test.InputContext.context_message:type_name -> google.protobuf.Any
	10, // 8: cel.expr.conformance.test.InputValue.value:type_name -> cel.expr.Value
	10, // 9: cel.expr.conformance.test.TestOutput.result_value:type_name -> cel.expr.Value
	11, // 10: cel.expr.conformance.test.TestOutput.eval_error:type_name -> cel.expr.ErrorSet
	12, // 11: cel.expr.conformance.test.TestOutput.unknown:type_name -> cel.expr.UnknownSet
	4,  // 12: cel.expr.conformance.test.TestCase.InputBindingsEntry.value:type_name -> cel.expr.conformance.test.InputValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cel_expr_conformance_test_suite_proto_init() }
func file_cel_expr_conformance_test_suite_proto_init() {
	if File_cel_expr_conformance_test_suite_proto != nil {
		return
	}
	file_cel_expr_conformance_test_suite_proto_msgTypes[3].OneofWrappers = []any{
		(*InputContext_ContextMessage)(nil),
		(*InputContext_ContextExpr)(nil),
	}
	file_cel_expr_conformance_test_suite_proto_msgTypes[4].OneofWrappers = []any{
		(*InputValue_Value)(nil),
		(*InputValue_Expr)(nil),
	}
	file_cel_expr_conformance_test_suite_proto_msgTypes[5].OneofWrappers = []any{
		(*TestOutput_ResultValue)(nil),
		(*TestOutput_ResultExpr)(nil),
		(*TestOutput_EvalError)(nil),
		(*TestOutput_Unknown)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cel_expr_conformance_test_suite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cel_expr_conformance_test_suite_proto_goTypes,
		DependencyIndexes: file_cel_expr_conformance_test_suite_proto_depIdxs,
		MessageInfos:      file_cel_expr_conformance_test_suite_proto_msgTypes,
	}.Build()
	File_cel_expr_conformance_test_suite_proto = out.File
	file_cel_expr_conformance_test_suite_proto_rawDesc = nil
	file_cel_expr_conformance_test_suite_proto_goTypes = nil
	file_cel_expr_conformance_test_suite_proto_depIdxs = nil
}
