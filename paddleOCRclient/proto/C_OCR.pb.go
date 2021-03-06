// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: proto/C_OCR.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageType  string `protobuf:"bytes,1,opt,name=ImageType,proto3" json:"ImageType,omitempty"`
	ImageBytes []byte `protobuf:"bytes,2,opt,name=ImageBytes,proto3" json:"ImageBytes,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_C_OCR_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_C_OCR_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_C_OCR_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

func (x *Request) GetImageBytes() []byte {
	if x != nil {
		return x.ImageBytes
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	ImageBytes []byte `protobuf:"bytes,2,opt,name=ImageBytes,proto3" json:"ImageBytes,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_C_OCR_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_C_OCR_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_C_OCR_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Response) GetImageBytes() []byte {
	if x != nil {
		return x.ImageBytes
	}
	return nil
}

var File_proto_C_OCR_proto protoreflect.FileDescriptor

var file_proto_C_OCR_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x5f, 0x4f, 0x43, 0x52, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x32, 0x25, 0x0a, 0x05,
	0x43, 0x5f, 0x4f, 0x43, 0x52, 0x12, 0x1c, 0x0a, 0x03, 0x4f, 0x63, 0x72, 0x12, 0x08, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_C_OCR_proto_rawDescOnce sync.Once
	file_proto_C_OCR_proto_rawDescData = file_proto_C_OCR_proto_rawDesc
)

func file_proto_C_OCR_proto_rawDescGZIP() []byte {
	file_proto_C_OCR_proto_rawDescOnce.Do(func() {
		file_proto_C_OCR_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_C_OCR_proto_rawDescData)
	})
	return file_proto_C_OCR_proto_rawDescData
}

var file_proto_C_OCR_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_C_OCR_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: Request
	(*Response)(nil), // 1: Response
}
var file_proto_C_OCR_proto_depIdxs = []int32{
	0, // 0: C_OCR.Ocr:input_type -> Request
	1, // 1: C_OCR.Ocr:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_C_OCR_proto_init() }
func file_proto_C_OCR_proto_init() {
	if File_proto_C_OCR_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_C_OCR_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_C_OCR_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_C_OCR_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_C_OCR_proto_goTypes,
		DependencyIndexes: file_proto_C_OCR_proto_depIdxs,
		MessageInfos:      file_proto_C_OCR_proto_msgTypes,
	}.Build()
	File_proto_C_OCR_proto = out.File
	file_proto_C_OCR_proto_rawDesc = nil
	file_proto_C_OCR_proto_goTypes = nil
	file_proto_C_OCR_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// C_OCRClient is the client API for C_OCR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type C_OCRClient interface {
	Ocr(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type c_OCRClient struct {
	cc grpc.ClientConnInterface
}

func NewC_OCRClient(cc grpc.ClientConnInterface) C_OCRClient {
	return &c_OCRClient{cc}
}

func (c *c_OCRClient) Ocr(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/C_OCR/Ocr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// C_OCRServer is the server API for C_OCR service.
type C_OCRServer interface {
	Ocr(context.Context, *Request) (*Response, error)
}

// UnimplementedC_OCRServer can be embedded to have forward compatible implementations.
type UnimplementedC_OCRServer struct {
}

func (*UnimplementedC_OCRServer) Ocr(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ocr not implemented")
}

func RegisterC_OCRServer(s *grpc.Server, srv C_OCRServer) {
	s.RegisterService(&_C_OCR_serviceDesc, srv)
}

func _C_OCR_Ocr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(C_OCRServer).Ocr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/C_OCR/Ocr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(C_OCRServer).Ocr(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _C_OCR_serviceDesc = grpc.ServiceDesc{
	ServiceName: "C_OCR",
	HandlerType: (*C_OCRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ocr",
			Handler:    _C_OCR_Ocr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/C_OCR.proto",
}
