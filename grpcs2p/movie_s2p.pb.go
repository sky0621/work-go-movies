// Code generated by protoc-gen-go.
// source: grpcs2p/movie_s2p.proto
// DO NOT EDIT!

/*
Package moviess2p is a generated protocol buffer package.

It is generated from these files:
	grpcs2p/movie_s2p.proto

It has these top-level messages:
	Movie
	Movies
	MovieSkey
*/
package moviess2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Movie struct {
	Skey          string `protobuf:"bytes,1,opt,name=skey" json:"skey,omitempty"`
	Filename      string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
	Title         string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Playtime      string `protobuf:"bytes,4,opt,name=playtime" json:"playtime,omitempty"`
	Photodatetime string `protobuf:"bytes,5,opt,name=photodatetime" json:"photodatetime,omitempty"`
}

func (m *Movie) Reset()                    { *m = Movie{} }
func (m *Movie) String() string            { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()               {}
func (*Movie) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Movies struct {
	Movies []*Movie `protobuf:"bytes,1,rep,name=movies" json:"movies,omitempty"`
}

func (m *Movies) Reset()                    { *m = Movies{} }
func (m *Movies) String() string            { return proto.CompactTextString(m) }
func (*Movies) ProtoMessage()               {}
func (*Movies) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Movies) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

type MovieSkey struct {
	Skey string `protobuf:"bytes,1,opt,name=skey" json:"skey,omitempty"`
}

func (m *MovieSkey) Reset()                    { *m = MovieSkey{} }
func (m *MovieSkey) String() string            { return proto.CompactTextString(m) }
func (*MovieSkey) ProtoMessage()               {}
func (*MovieSkey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Movie)(nil), "moviess2p.Movie")
	proto.RegisterType((*Movies)(nil), "moviess2p.Movies")
	proto.RegisterType((*MovieSkey)(nil), "moviess2p.MovieSkey")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MovieS2PService service

type MovieS2PServiceClient interface {
	GetMovie(ctx context.Context, in *MovieSkey, opts ...grpc.CallOption) (*Movie, error)
	GetMovies(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Movies, error)
}

type movieS2PServiceClient struct {
	cc *grpc.ClientConn
}

func NewMovieS2PServiceClient(cc *grpc.ClientConn) MovieS2PServiceClient {
	return &movieS2PServiceClient{cc}
}

func (c *movieS2PServiceClient) GetMovie(ctx context.Context, in *MovieSkey, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := grpc.Invoke(ctx, "/moviess2p.MovieS2PService/GetMovie", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieS2PServiceClient) GetMovies(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Movies, error) {
	out := new(Movies)
	err := grpc.Invoke(ctx, "/moviess2p.MovieS2PService/GetMovies", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieS2PService service

type MovieS2PServiceServer interface {
	GetMovie(context.Context, *MovieSkey) (*Movie, error)
	GetMovies(context.Context, *Movie) (*Movies, error)
}

func RegisterMovieS2PServiceServer(s *grpc.Server, srv MovieS2PServiceServer) {
	s.RegisterService(&_MovieS2PService_serviceDesc, srv)
}

func _MovieS2PService_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieSkey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieS2PServiceServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moviess2p.MovieS2PService/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieS2PServiceServer).GetMovie(ctx, req.(*MovieSkey))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieS2PService_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieS2PServiceServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moviess2p.MovieS2PService/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieS2PServiceServer).GetMovies(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieS2PService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moviess2p.MovieS2PService",
	HandlerType: (*MovieS2PServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _MovieS2PService_GetMovie_Handler,
		},
		{
			MethodName: "GetMovies",
			Handler:    _MovieS2PService_GetMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("grpcs2p/movie_s2p.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0xb6, 0x09, 0xcd, 0x88, 0xa8, 0x43, 0xc1, 0x25, 0x17, 0x4b, 0xf0, 0x90, 0x53, 0x94,
	0xf1, 0x21, 0x3c, 0x09, 0xd2, 0x3e, 0x80, 0xc4, 0x3a, 0xea, 0x62, 0xe2, 0x2e, 0x99, 0xa5, 0xd2,
	0xa7, 0xf0, 0x95, 0x25, 0xb3, 0xb6, 0x60, 0xd3, 0xdb, 0x7e, 0x7f, 0xb3, 0xdf, 0x0c, 0x5c, 0xbd,
	0xf7, 0x7e, 0x2d, 0xe4, 0x6f, 0x3b, 0xb7, 0xb1, 0xfc, 0x2c, 0xe4, 0x6b, 0xdf, 0xbb, 0xe0, 0x30,
	0x57, 0x42, 0x84, 0x7c, 0xf9, 0x93, 0x40, 0xfa, 0x38, 0x20, 0x44, 0x98, 0xca, 0x27, 0x6f, 0x4d,
	0xb2, 0x48, 0xaa, 0x7c, 0xa9, 0x6f, 0x2c, 0x60, 0xf6, 0x66, 0x5b, 0xfe, 0x6a, 0x3a, 0x36, 0x27,
	0xca, 0xef, 0x31, 0xce, 0x21, 0x0d, 0x36, 0xb4, 0x6c, 0x26, 0x2a, 0x44, 0x30, 0x24, 0x7c, 0xdb,
	0x6c, 0x83, 0xed, 0xd8, 0x4c, 0x63, 0x62, 0x87, 0xf1, 0x06, 0xce, 0xfc, 0x87, 0x0b, 0xee, 0xb5,
	0x09, 0xac, 0x86, 0x54, 0x0d, 0xff, 0xc9, 0x92, 0x20, 0xd3, 0x42, 0x82, 0x15, 0x64, 0xb1, 0xa8,
	0x49, 0x16, 0x93, 0xea, 0x94, 0x2e, 0xea, 0x7d, 0xef, 0x5a, 0x2d, 0xcb, 0x3f, 0xbd, 0xbc, 0x86,
	0x5c, 0x89, 0xd5, 0x50, 0xfa, 0xc8, 0x22, 0xf4, 0x0d, 0xe7, 0xd1, 0x40, 0x4f, 0x2b, 0xee, 0x37,
	0x76, 0xcd, 0x48, 0x30, 0x7b, 0xe0, 0x10, 0x77, 0x9f, 0x1f, 0x4e, 0x1e, 0x06, 0x15, 0xa3, 0xff,
	0xf0, 0x0e, 0xf2, 0x5d, 0x46, 0x70, 0x24, 0x17, 0x97, 0x87, 0x8c, 0xbc, 0x64, 0x7a, 0xf1, 0xfb,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x58, 0xda, 0x2f, 0x8c, 0x01, 0x00, 0x00,
}
