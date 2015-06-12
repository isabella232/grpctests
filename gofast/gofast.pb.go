// Code generated by protoc-gen-gogo.
// source: gofast.proto
// DO NOT EDIT!

/*
Package gofast is a generated protocol buffer package.

It is generated from these files:
	gofast.proto

It has these top-level messages:
	MyRequest
	MyResponse
*/
package gofast

import proto "github.com/golang/protobuf/proto"
import math "math"

import io "io"
import fmt "fmt"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MyRequest struct {
	Value            *int64 `protobuf:"varint,1,opt" json:"Value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MyRequest) Reset()         { *m = MyRequest{} }
func (m *MyRequest) String() string { return proto.CompactTextString(m) }
func (*MyRequest) ProtoMessage()    {}

func (m *MyRequest) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type MyResponse struct {
	Value            *int64 `protobuf:"varint,1,opt" json:"Value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MyResponse) Reset()         { *m = MyResponse{} }
func (m *MyResponse) String() string { return proto.CompactTextString(m) }
func (*MyResponse) ProtoMessage()    {}

func (m *MyResponse) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func init() {
}
func (m *MyRequest) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := skipGofast(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}

	return nil
}
func (m *MyResponse) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := skipGofast(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}

	return nil
}
func skipGofast(data []byte) (n int, err error) {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if index >= l {
					return 0, io.ErrUnexpectedEOF
				}
				index++
				if data[index-1] < 0x80 {
					break
				}
			}
			return index, nil
		case 1:
			index += 8
			return index, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			index += length
			return index, nil
		case 3:
			for {
				var wire uint64
				var start int = index
				for shift := uint(0); ; shift += 7 {
					if index >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[index]
					index++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				wireType := int(wire & 0x7)
				if wireType == 4 {
					break
				}
				next, err := skipGofast(data[start:])
				if err != nil {
					return 0, err
				}
				index = start + next
			}
			return index, nil
		case 4:
			return index, nil
		case 5:
			index += 4
			return index, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}
func (m *MyRequest) Size() (n int) {
	var l int
	_ = l
	if m.Value != nil {
		n += 1 + sovGofast(uint64(*m.Value))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MyResponse) Size() (n int) {
	var l int
	_ = l
	if m.Value != nil {
		n += 1 + sovGofast(uint64(*m.Value))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGofast(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGofast(x uint64) (n int) {
	return sovGofast(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MyRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MyRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGofast(data, i, uint64(*m.Value))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *MyResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MyResponse) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGofast(data, i, uint64(*m.Value))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Gofast(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Gofast(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGofast(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

// Client API for MyTest service

type MyTestClient interface {
	UnaryCall(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
}

type myTestClient struct {
	cc *grpc.ClientConn
}

func NewMyTestClient(cc *grpc.ClientConn) MyTestClient {
	return &myTestClient{cc}
}

func (c *myTestClient) UnaryCall(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := grpc.Invoke(ctx, "/gofast.MyTest/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyTest service

type MyTestServer interface {
	UnaryCall(context.Context, *MyRequest) (*MyResponse, error)
}

func RegisterMyTestServer(s *grpc.Server, srv MyTestServer) {
	s.RegisterService(&_MyTest_serviceDesc, srv)
}

func _MyTest_UnaryCall_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(MyRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(MyTestServer).UnaryCall(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _MyTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gofast.MyTest",
	HandlerType: (*MyTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _MyTest_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}