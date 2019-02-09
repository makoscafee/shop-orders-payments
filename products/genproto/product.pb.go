// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package go_micro_srv_products

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Product struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Picture     string `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	PriceUsd    *Money `protobuf:"bytes,5,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	// Categories such as "vintage" or "gardening" that can be used to look up
	// other related products.
	Categories           []string `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *Product) GetPriceUsd() *Money {
	if m != nil {
		return m.PriceUsd
	}
	return nil
}

func (m *Product) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

// Represents an amount of money with its currency type.
type Money struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The whole units of the amount.
	// For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos                int32    `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Money.Unmarshal(m, b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Money.Marshal(b, m, deterministic)
}
func (m *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(m, src)
}
func (m *Money) XXX_Size() int {
	return xxx_messageInfo_Money.Size(m)
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Money) GetUnits() int64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *Money) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type CreateProductResponse struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Product              *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductResponse) Reset()         { *m = CreateProductResponse{} }
func (m *CreateProductResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProductResponse) ProtoMessage()    {}
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{3}
}

func (m *CreateProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductResponse.Unmarshal(m, b)
}
func (m *CreateProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductResponse.Marshal(b, m, deterministic)
}
func (m *CreateProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductResponse.Merge(m, src)
}
func (m *CreateProductResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProductResponse.Size(m)
}
func (m *CreateProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductResponse proto.InternalMessageInfo

func (m *CreateProductResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *CreateProductResponse) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type ListProductsResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProductsResponse) Reset()         { *m = ListProductsResponse{} }
func (m *ListProductsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProductsResponse) ProtoMessage()    {}
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{4}
}

func (m *ListProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductsResponse.Unmarshal(m, b)
}
func (m *ListProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductsResponse.Marshal(b, m, deterministic)
}
func (m *ListProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductsResponse.Merge(m, src)
}
func (m *ListProductsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProductsResponse.Size(m)
}
func (m *ListProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductsResponse proto.InternalMessageInfo

func (m *ListProductsResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type GetProductRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRequest) Reset()         { *m = GetProductRequest{} }
func (m *GetProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRequest) ProtoMessage()    {}
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{5}
}

func (m *GetProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRequest.Unmarshal(m, b)
}
func (m *GetProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRequest.Marshal(b, m, deterministic)
}
func (m *GetProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRequest.Merge(m, src)
}
func (m *GetProductRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductRequest.Size(m)
}
func (m *GetProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRequest proto.InternalMessageInfo

func (m *GetProductRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchProductsRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchProductsRequest) Reset()         { *m = SearchProductsRequest{} }
func (m *SearchProductsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchProductsRequest) ProtoMessage()    {}
func (*SearchProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{6}
}

func (m *SearchProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProductsRequest.Unmarshal(m, b)
}
func (m *SearchProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProductsRequest.Marshal(b, m, deterministic)
}
func (m *SearchProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProductsRequest.Merge(m, src)
}
func (m *SearchProductsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchProductsRequest.Size(m)
}
func (m *SearchProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProductsRequest proto.InternalMessageInfo

func (m *SearchProductsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchProductsResponse struct {
	Results              []*Product `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchProductsResponse) Reset()         { *m = SearchProductsResponse{} }
func (m *SearchProductsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchProductsResponse) ProtoMessage()    {}
func (*SearchProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{7}
}

func (m *SearchProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProductsResponse.Unmarshal(m, b)
}
func (m *SearchProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProductsResponse.Marshal(b, m, deterministic)
}
func (m *SearchProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProductsResponse.Merge(m, src)
}
func (m *SearchProductsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchProductsResponse.Size(m)
}
func (m *SearchProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProductsResponse proto.InternalMessageInfo

func (m *SearchProductsResponse) GetResults() []*Product {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "go.micro.srv.products.Empty")
	proto.RegisterType((*Product)(nil), "go.micro.srv.products.Product")
	proto.RegisterType((*Money)(nil), "go.micro.srv.products.Money")
	proto.RegisterType((*CreateProductResponse)(nil), "go.micro.srv.products.CreateProductResponse")
	proto.RegisterType((*ListProductsResponse)(nil), "go.micro.srv.products.ListProductsResponse")
	proto.RegisterType((*GetProductRequest)(nil), "go.micro.srv.products.GetProductRequest")
	proto.RegisterType((*SearchProductsRequest)(nil), "go.micro.srv.products.SearchProductsRequest")
	proto.RegisterType((*SearchProductsResponse)(nil), "go.micro.srv.products.SearchProductsResponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x8d, 0xac, 0x28, 0xb2, 0xc7, 0xb1, 0xa1, 0x83, 0x5d, 0x84, 0x29, 0x41, 0x6c, 0x2e, 0x86,
	0x36, 0x3a, 0xb8, 0x97, 0xb4, 0xd7, 0x50, 0x7a, 0x69, 0xa1, 0x6c, 0x28, 0x04, 0x4a, 0x09, 0xea,
	0x6a, 0x48, 0x97, 0xd6, 0x5a, 0x65, 0x77, 0x15, 0xf0, 0x97, 0xf4, 0x87, 0xfa, 0x61, 0x25, 0x2b,
	0xad, 0xe3, 0x34, 0x12, 0xf1, 0x4d, 0xf3, 0xde, 0xdb, 0x9d, 0x79, 0x6f, 0x24, 0xc1, 0xa4, 0xd2,
	0xaa, 0xa8, 0x85, 0xcd, 0x2a, 0xad, 0xac, 0xc2, 0xf9, 0x8d, 0xca, 0xd6, 0x52, 0x68, 0x95, 0x19,
	0x7d, 0x97, 0xb5, 0x9c, 0x61, 0x31, 0x44, 0x1f, 0xd6, 0x95, 0xdd, 0xb0, 0xbf, 0x01, 0xc4, 0x5f,
	0x1a, 0x14, 0xa7, 0x30, 0x90, 0x45, 0x12, 0xa4, 0xc1, 0x72, 0xc4, 0x07, 0xb2, 0x40, 0x84, 0xc3,
	0x32, 0x5f, 0x53, 0x32, 0x70, 0x88, 0x7b, 0xc6, 0x14, 0xc6, 0x05, 0x19, 0xa1, 0x65, 0x65, 0xa5,
	0x2a, 0x93, 0xd0, 0x51, 0xbb, 0x10, 0x26, 0x10, 0x57, 0x52, 0xd8, 0x5a, 0x53, 0x72, 0xe8, 0x58,
	0x5f, 0xe2, 0x3b, 0x18, 0x55, 0x5a, 0x0a, 0xba, 0xae, 0x4d, 0x91, 0x44, 0x69, 0xb0, 0x1c, 0xaf,
	0x5e, 0x65, 0x9d, 0xf3, 0x65, 0x9f, 0x55, 0x49, 0x1b, 0x3e, 0x74, 0xf2, 0xaf, 0xa6, 0xc0, 0x13,
	0x00, 0x91, 0x5b, 0xba, 0x51, 0x5a, 0x92, 0x49, 0x8e, 0xd2, 0x70, 0x39, 0xe2, 0x3b, 0x08, 0xbb,
	0x82, 0xc8, 0x1d, 0xc1, 0x53, 0x98, 0x88, 0x5a, 0x6b, 0x2a, 0xc5, 0xe6, 0x5a, 0xa8, 0x82, 0x5a,
	0x3b, 0xc7, 0x1e, 0xbc, 0x50, 0x05, 0xe1, 0x0c, 0xa2, 0xba, 0x94, 0xd6, 0x38, 0x67, 0x21, 0x6f,
	0x8a, 0x7b, 0xb4, 0xcc, 0x4b, 0x65, 0x9c, 0xa9, 0x88, 0x37, 0x05, 0xfb, 0x05, 0xf3, 0x0b, 0x4d,
	0xb9, 0xa5, 0x36, 0x25, 0x4e, 0xa6, 0x52, 0xa5, 0xa1, 0x7b, 0x9f, 0xc2, 0x11, 0x4d, 0x64, 0x43,
	0xee, 0x4b, 0x3c, 0x87, 0xb8, 0x35, 0xe2, 0x1a, 0x8c, 0x57, 0x27, 0x3d, 0x2e, 0xfd, 0x95, 0x5e,
	0xce, 0x38, 0xcc, 0x3e, 0x49, 0x63, 0x5b, 0xdc, 0x6c, 0x7b, 0xbd, 0x87, 0xa1, 0x3f, 0x94, 0x04,
	0x69, 0xb8, 0xc7, 0x95, 0x5b, 0x3d, 0x3b, 0x85, 0x17, 0x1f, 0xc9, 0x6e, 0xa7, 0xbf, 0xad, 0xc9,
	0x3c, 0x59, 0x35, 0x3b, 0x83, 0xf9, 0x25, 0xe5, 0x5a, 0xfc, 0x7c, 0x68, 0xdd, 0x08, 0x67, 0x10,
	0xdd, 0xd6, 0xa4, 0x37, 0xad, 0xb6, 0x29, 0x18, 0x87, 0x97, 0xff, 0xcb, 0xdb, 0x49, 0xcf, 0x21,
	0xd6, 0x64, 0xea, 0xdf, 0x7b, 0x0f, 0xea, 0xe5, 0xab, 0x3f, 0x21, 0x4c, 0x5b, 0xf0, 0x92, 0xf4,
	0x9d, 0x14, 0x84, 0xdf, 0x61, 0xf2, 0x28, 0x7b, 0x7c, 0xe6, 0xb2, 0xc5, 0x9b, 0x1e, 0xbe, 0x7b,
	0x83, 0xdf, 0xe0, 0x78, 0x37, 0x6d, 0xec, 0x7b, 0x19, 0xdd, 0x97, 0xb2, 0x78, 0xdd, 0xc3, 0x76,
	0x2d, 0x8c, 0x1d, 0xe0, 0x15, 0xc0, 0x43, 0xec, 0xb8, 0xec, 0x39, 0xfc, 0x64, 0x33, 0x8b, 0x67,
	0x2c, 0xb2, 0x03, 0x54, 0x30, 0x7d, 0x1c, 0x3e, 0xf6, 0xd9, 0xee, 0x5c, 0xe9, 0xe2, 0x6c, 0x4f,
	0xb5, 0xb7, 0xf2, 0xe3, 0xc8, 0xfd, 0x4a, 0xde, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x25, 0x9f,
	0xdb, 0x5a, 0x5b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*CreateProductResponse, error)
	ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/go.micro.srv.products.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, "/go.micro.srv.products.ProductService/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/go.micro.srv.products.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, "/go.micro.srv.products.ProductService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	CreateProduct(context.Context, *Product) (*CreateProductResponse, error)
	ListProducts(context.Context, *Empty) (*ListProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.products.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.products.ProductService/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.products.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.products.ProductService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.products.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _ProductService_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductService_SearchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
