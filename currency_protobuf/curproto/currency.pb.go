// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

/*
Package curproto is a generated protocol buffer package.

It is generated from these files:
	currency.proto

It has these top-level messages:
	Currency
	CurrencyList
	CurrencyRequest
*/
package curproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Currency struct {
	Code    string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Number  int32  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	Country string `protobuf:"bytes,4,opt,name=country" json:"country,omitempty"`
}

func (m *Currency) Reset()                    { *m = Currency{} }
func (m *Currency) String() string            { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()               {}
func (*Currency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Currency) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Currency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currency) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Currency) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type CurrencyList struct {
	Items []*Currency `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *CurrencyList) Reset()                    { *m = CurrencyList{} }
func (m *CurrencyList) String() string            { return proto.CompactTextString(m) }
func (*CurrencyList) ProtoMessage()               {}
func (*CurrencyList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CurrencyList) GetItems() []*Currency {
	if m != nil {
		return m.Items
	}
	return nil
}

type CurrencyRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CurrencyRequest) Reset()                    { *m = CurrencyRequest{} }
func (m *CurrencyRequest) String() string            { return proto.CompactTextString(m) }
func (*CurrencyRequest) ProtoMessage()               {}
func (*CurrencyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CurrencyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Currency)(nil), "curproto.Currency")
	proto.RegisterType((*CurrencyList)(nil), "curproto.CurrencyList")
	proto.RegisterType((*CurrencyRequest)(nil), "curproto.CurrencyRequest")
}

func init() { proto.RegisterFile("currency.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x3d, 0xcf, 0x82, 0x30,
	0x14, 0x85, 0xd3, 0x97, 0x8f, 0x17, 0xaf, 0x46, 0x93, 0x3b, 0x98, 0x8e, 0x84, 0xc4, 0xa4, 0x13,
	0x83, 0x2e, 0xee, 0xae, 0x4e, 0xfc, 0x03, 0x29, 0x77, 0x60, 0xa0, 0xd5, 0x4b, 0x3b, 0xf0, 0xef,
	0x0d, 0x85, 0xe2, 0x76, 0xce, 0xe9, 0xd3, 0xdc, 0x07, 0x8e, 0xda, 0x33, 0x93, 0xd1, 0x53, 0xfd,
	0x66, 0xeb, 0x2c, 0x16, 0xda, 0x73, 0x48, 0x55, 0x07, 0xc5, 0x63, 0x7d, 0x43, 0x84, 0x54, 0xdb,
	0x8e, 0xa4, 0x28, 0x85, 0xda, 0x35, 0x21, 0xcf, 0x9b, 0x79, 0x0d, 0x24, 0xff, 0x96, 0x6d, 0xce,
	0x78, 0x86, 0xdc, 0xf8, 0xa1, 0x25, 0x96, 0x49, 0x29, 0x54, 0xd6, 0xac, 0x0d, 0x25, 0xfc, 0x6b,
	0xeb, 0x8d, 0xe3, 0x49, 0xa6, 0x01, 0x8f, 0xb5, 0xba, 0xc3, 0x21, 0x5e, 0x79, 0xf6, 0xa3, 0x43,
	0x05, 0x59, 0xef, 0x68, 0x18, 0xa5, 0x28, 0x13, 0xb5, 0xbf, 0x62, 0x1d, 0x7d, 0xea, 0x88, 0x35,
	0x0b, 0x50, 0x5d, 0xe0, 0xb4, 0x4d, 0xf4, 0xf1, 0x34, 0xba, 0x4d, 0x49, 0xfc, 0x94, 0xda, 0x3c,
	0xfc, 0xbe, 0x7d, 0x03, 0x00, 0x00, 0xff, 0xff, 0x29, 0x83, 0xa2, 0xf2, 0xe9, 0x00, 0x00, 0x00,
}
