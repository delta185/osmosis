// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/bridge/v1beta1/bridge.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AssetStatus int32

const (
	AssetStatus_ASSET_STATUS_UNSPECIFIED      AssetStatus = 0
	AssetStatus_ASSET_STATUS_OK               AssetStatus = 1
	AssetStatus_ASSET_STATUS_BLOCKED_INBOUND  AssetStatus = 2
	AssetStatus_ASSET_STATUS_BLOCKED_OUTBOUND AssetStatus = 3
	AssetStatus_ASSET_STATUS_BLOCKED_BOTH     AssetStatus = 4
)

var AssetStatus_name = map[int32]string{
	0: "ASSET_STATUS_UNSPECIFIED",
	1: "ASSET_STATUS_OK",
	2: "ASSET_STATUS_BLOCKED_INBOUND",
	3: "ASSET_STATUS_BLOCKED_OUTBOUND",
	4: "ASSET_STATUS_BLOCKED_BOTH",
}

var AssetStatus_value = map[string]int32{
	"ASSET_STATUS_UNSPECIFIED":      0,
	"ASSET_STATUS_OK":               1,
	"ASSET_STATUS_BLOCKED_INBOUND":  2,
	"ASSET_STATUS_BLOCKED_OUTBOUND": 3,
	"ASSET_STATUS_BLOCKED_BOTH":     4,
}

func (x AssetStatus) String() string {
	return proto.EnumName(AssetStatus_name, int32(x))
}

func (AssetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f999ddf08452f1f3, []int{0}
}

// Params defines params for x/bridge module.
type Params struct {
	// Signers used to sign inbound and release outbound transactions
	Signers []string `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty" yaml:"signers"`
	// Assets is a list used to create tokenfactory denoms
	// for corresponding trading pairs
	Assets []Asset `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets" yaml:"assets"`
	// VotesNeeded marks how many signers out of the list of signers need
	// to sign until a tx can be considered finalized
	VotesNeeded uint64 `protobuf:"varint,3,opt,name=votes_needed,json=votesNeeded,proto3" json:"votes_needed,omitempty" yaml:"votes_needed"`
	// Fee defines a param for fee that go towards the validator set
	// signing the incoming/outgoing txs. The fee is measured as a ratio,
	// so its value lies between 0 and 1.
	Fee cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=fee,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fee" yaml:"fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f999ddf08452f1f3, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *Params) GetAssets() []Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *Params) GetVotesNeeded() uint64 {
	if m != nil {
		return m.VotesNeeded
	}
	return 0
}

// AssetID defines a pair of the source chain name and its Osmosis
// representation denoted by denom. AssetID is a primary key for Asset.
type AssetID struct {
	// SourceChain is a source chain name
	SourceChain string `protobuf:"bytes,1,opt,name=source_chain,json=sourceChain,proto3" json:"source_chain,omitempty" yaml:"source_chain"`
	// Denom is the Osmosis representation of the SourceChain
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *AssetID) Reset()         { *m = AssetID{} }
func (m *AssetID) String() string { return proto.CompactTextString(m) }
func (*AssetID) ProtoMessage()    {}
func (*AssetID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f999ddf08452f1f3, []int{1}
}
func (m *AssetID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetID.Merge(m, src)
}
func (m *AssetID) XXX_Size() int {
	return m.Size()
}
func (m *AssetID) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetID.DiscardUnknown(m)
}

var xxx_messageInfo_AssetID proto.InternalMessageInfo

func (m *AssetID) GetSourceChain() string {
	if m != nil {
		return m.SourceChain
	}
	return ""
}

func (m *AssetID) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// Asset is a representation of the asset.
type Asset struct {
	// ID is the asset's primary key
	Id AssetID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	// Status is a current status of the asset
	Status AssetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=osmosis.bridge.v1beta1.AssetStatus" json:"status,omitempty" yaml:"status"`
	// Exponent represents the power of 10 used for coin representation
	Exponent uint64 `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty" yaml:"exponent"`
	// ExternalConfirmations is a number of the confirmations on the external
	// chain needed to consider the transfer confirmed
	ExternalConfirmations uint64 `protobuf:"varint,4,opt,name=external_confirmations,json=externalConfirmations,proto3" json:"external_confirmations,omitempty" yaml:"external_confirmations"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f999ddf08452f1f3, []int{2}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetId() AssetID {
	if m != nil {
		return m.Id
	}
	return AssetID{}
}

func (m *Asset) GetStatus() AssetStatus {
	if m != nil {
		return m.Status
	}
	return AssetStatus_ASSET_STATUS_UNSPECIFIED
}

func (m *Asset) GetExponent() uint64 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *Asset) GetExternalConfirmations() uint64 {
	if m != nil {
		return m.ExternalConfirmations
	}
	return 0
}

// InboundTransfer is a representation of the inbound transfer.
type InboundTransfer struct {
	// ExternalId is a unique ID of the transfer coming from outside.
	// Serves the purpose of uniquely identifying the transfer in another chain
	// (e.g., this might be the BTC tx hash).
	ExternalId string `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty" yaml:"external_id"`
	// ExternalHeight is the height at which the transfer occurred in the external
	// chain
	ExternalHeight uint64 `protobuf:"varint,2,opt,name=external_height,json=externalHeight,proto3" json:"external_height,omitempty" yaml:"external_height"`
	// DestAddr is a destination Osmosis address
	DestAddr string `protobuf:"bytes,3,opt,name=dest_addr,json=destAddr,proto3" json:"dest_addr,omitempty" yaml:"dest_addr"`
	// AssetID is the ID of the asset being transferred
	AssetId AssetID `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id" yaml:"asset_id"`
	// Amount of coins to transfer
	Amount cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount" yaml:"amount"`
	// Voters is a list of validators signed this transfer
	Voters []string `protobuf:"bytes,6,rep,name=voters,proto3" json:"voters,omitempty" yaml:"voters"`
	// Finalized indicates whether the transfer needs more votes or has
	// already accumulated a sufficient number. The finalised flag is set
	// to true as soon as length(voters) is greater than or equal to
	// the module's param votes_needed.
	Finalized bool `protobuf:"varint,7,opt,name=finalized,proto3" json:"finalized,omitempty" yaml:"finalized"`
}

func (m *InboundTransfer) Reset()         { *m = InboundTransfer{} }
func (m *InboundTransfer) String() string { return proto.CompactTextString(m) }
func (*InboundTransfer) ProtoMessage()    {}
func (*InboundTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f999ddf08452f1f3, []int{3}
}
func (m *InboundTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboundTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboundTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InboundTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboundTransfer.Merge(m, src)
}
func (m *InboundTransfer) XXX_Size() int {
	return m.Size()
}
func (m *InboundTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_InboundTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_InboundTransfer proto.InternalMessageInfo

func (m *InboundTransfer) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *InboundTransfer) GetExternalHeight() uint64 {
	if m != nil {
		return m.ExternalHeight
	}
	return 0
}

func (m *InboundTransfer) GetDestAddr() string {
	if m != nil {
		return m.DestAddr
	}
	return ""
}

func (m *InboundTransfer) GetAssetId() AssetID {
	if m != nil {
		return m.AssetId
	}
	return AssetID{}
}

func (m *InboundTransfer) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *InboundTransfer) GetFinalized() bool {
	if m != nil {
		return m.Finalized
	}
	return false
}

func init() {
	proto.RegisterEnum("osmosis.bridge.v1beta1.AssetStatus", AssetStatus_name, AssetStatus_value)
	proto.RegisterType((*Params)(nil), "osmosis.bridge.v1beta1.Params")
	proto.RegisterType((*AssetID)(nil), "osmosis.bridge.v1beta1.AssetID")
	proto.RegisterType((*Asset)(nil), "osmosis.bridge.v1beta1.Asset")
	proto.RegisterType((*InboundTransfer)(nil), "osmosis.bridge.v1beta1.InboundTransfer")
}

func init() {
	proto.RegisterFile("osmosis/bridge/v1beta1/bridge.proto", fileDescriptor_f999ddf08452f1f3)
}

var fileDescriptor_f999ddf08452f1f3 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x8f, 0xda, 0x46,
	0x14, 0xc7, 0xc0, 0xb2, 0xcb, 0x90, 0x2e, 0x64, 0x36, 0xbb, 0x71, 0xd3, 0x80, 0xc9, 0x44, 0xaa,
	0x68, 0xd5, 0xda, 0x82, 0x56, 0xaa, 0x94, 0x1b, 0x06, 0xa2, 0x58, 0x41, 0x10, 0x19, 0x90, 0xaa,
	0x5e, 0xd0, 0xc0, 0x0c, 0x60, 0x15, 0x7b, 0x56, 0x9e, 0x61, 0xb5, 0xdb, 0x7b, 0xef, 0xfd, 0x10,
	0x55, 0x3f, 0x46, 0xcf, 0x39, 0xe6, 0x58, 0xf5, 0x60, 0x55, 0xbb, 0xdf, 0xc0, 0x9f, 0xa0, 0xf2,
	0xd8, 0x26, 0xde, 0x14, 0xad, 0x7a, 0x9b, 0xf7, 0x7e, 0xbf, 0xdf, 0x7b, 0xcf, 0xef, 0x8f, 0xc1,
	0x4b, 0xc6, 0x5d, 0xc6, 0x1d, 0x6e, 0x2c, 0x7c, 0x87, 0xac, 0xa9, 0x71, 0xd5, 0x5e, 0x50, 0x81,
	0xdb, 0x89, 0xa9, 0x5f, 0xfa, 0x4c, 0x30, 0x78, 0x91, 0x90, 0xf4, 0xc4, 0x9b, 0x90, 0x9e, 0x3d,
	0x59, 0xb3, 0x35, 0x93, 0x14, 0x23, 0x7a, 0xc5, 0x6c, 0xf4, 0x6b, 0x1e, 0x94, 0xde, 0x61, 0x1f,
	0xbb, 0x1c, 0x7e, 0x03, 0x8e, 0xb9, 0xb3, 0xf6, 0xa8, 0xcf, 0x55, 0xa5, 0x59, 0x68, 0x95, 0x4d,
	0x18, 0x06, 0xda, 0xe9, 0x0d, 0x76, 0xb7, 0xaf, 0x50, 0x02, 0x20, 0x3b, 0xa5, 0xc0, 0x21, 0x28,
	0x61, 0xce, 0xa9, 0xe0, 0x6a, 0xbe, 0x59, 0x68, 0x55, 0x3a, 0x75, 0xfd, 0x70, 0x5e, 0xbd, 0x1b,
	0xb1, 0xcc, 0xf3, 0xf7, 0x81, 0x96, 0x0b, 0x03, 0xed, 0xb3, 0x38, 0x5e, 0x2c, 0x45, 0x76, 0x12,
	0x03, 0xbe, 0x02, 0x8f, 0xae, 0x98, 0xa0, 0x7c, 0xee, 0x51, 0x4a, 0x28, 0x51, 0x0b, 0x4d, 0xa5,
	0x55, 0x34, 0x9f, 0x86, 0x81, 0x76, 0x16, 0x0b, 0xb2, 0x28, 0xb2, 0x2b, 0xd2, 0x1c, 0x49, 0x0b,
	0xf6, 0x40, 0x61, 0x45, 0xa9, 0x5a, 0x6c, 0x2a, 0xad, 0xb2, 0xd9, 0x8e, 0xf2, 0xfc, 0x1d, 0x68,
	0x5f, 0x2c, 0x65, 0x39, 0x9c, 0xfc, 0xac, 0x3b, 0xcc, 0x70, 0xb1, 0xd8, 0xe8, 0x43, 0xba, 0xc6,
	0xcb, 0x9b, 0x3e, 0x5d, 0x86, 0x81, 0x06, 0xe2, 0xa8, 0x2b, 0x4a, 0x91, 0x1d, 0xa9, 0x91, 0x0b,
	0x8e, 0x65, 0xa1, 0x56, 0x3f, 0xaa, 0x85, 0xb3, 0x9d, 0xbf, 0xa4, 0xf3, 0xe5, 0x06, 0x3b, 0x9e,
	0xaa, 0xc8, 0xc0, 0x99, 0x5a, 0xb2, 0x28, 0xb2, 0x2b, 0xb1, 0xd9, 0x8b, 0x2c, 0xf8, 0x25, 0x38,
	0x22, 0xd4, 0x63, 0xae, 0x9a, 0x97, 0xa2, 0x5a, 0x18, 0x68, 0x8f, 0x62, 0x91, 0x74, 0x23, 0x3b,
	0x86, 0xd1, 0x1f, 0x79, 0x70, 0x24, 0xf3, 0x41, 0x13, 0xe4, 0x1d, 0x22, 0x73, 0x54, 0x3a, 0xda,
	0x83, 0x3d, 0xb4, 0xfa, 0xe6, 0xe3, 0xa4, 0x8b, 0xe5, 0x38, 0xa6, 0x43, 0x90, 0x9d, 0x77, 0x08,
	0x1c, 0x81, 0x12, 0x17, 0x58, 0xec, 0xb8, 0x4c, 0x7b, 0xda, 0x79, 0xf9, 0x60, 0x9c, 0x89, 0xa4,
	0x9a, 0x8f, 0x3f, 0x4e, 0x23, 0x16, 0x23, 0x3b, 0x89, 0x02, 0x0d, 0x70, 0x42, 0xaf, 0x2f, 0x99,
	0x47, 0x3d, 0x91, 0x4c, 0xe2, 0x2c, 0x0c, 0xb4, 0x6a, 0x4c, 0x4e, 0x11, 0x64, 0xef, 0x49, 0xf0,
	0x47, 0x70, 0x41, 0xaf, 0x05, 0xf5, 0x3d, 0xbc, 0x9d, 0x2f, 0x99, 0xb7, 0x72, 0x7c, 0x17, 0x0b,
	0x87, 0x79, 0x5c, 0x4e, 0xa5, 0x68, 0xbe, 0x08, 0x03, 0xad, 0x9e, 0xca, 0x0f, 0xf1, 0x90, 0x7d,
	0x9e, 0x02, 0xbd, 0x7b, 0xfe, 0x3f, 0x0b, 0xa0, 0x6a, 0x79, 0x0b, 0xb6, 0xf3, 0xc8, 0xd4, 0xc7,
	0x1e, 0x5f, 0x51, 0x1f, 0xfe, 0x00, 0x2a, 0xfb, 0x28, 0x49, 0xef, 0xca, 0xe6, 0x45, 0x18, 0x68,
	0xf0, 0x93, 0x14, 0x51, 0x7f, 0x40, 0x6a, 0x59, 0xd1, 0xa6, 0x54, 0xf7, 0xd8, 0x86, 0x3a, 0xeb,
	0x8d, 0x90, 0x0d, 0x2b, 0x9a, 0xcf, 0xc2, 0x40, 0xbb, 0xf8, 0x44, 0x1c, 0x13, 0x90, 0x7d, 0x9a,
	0x7a, 0xde, 0x48, 0x07, 0x6c, 0x83, 0x32, 0xa1, 0x5c, 0xcc, 0x31, 0x21, 0xbe, 0xec, 0x4e, 0xd9,
	0x7c, 0x12, 0x06, 0x5a, 0x2d, 0x1d, 0x73, 0x02, 0x21, 0xfb, 0x24, 0x7a, 0x77, 0x09, 0xf1, 0xe1,
	0x04, 0x9c, 0xc8, 0x3d, 0x8f, 0xaa, 0x2d, 0xfe, 0xbf, 0x49, 0x3f, 0x4d, 0x26, 0x5d, 0xcd, 0xdc,
	0x8b, 0xfc, 0x9e, 0x63, 0xf9, 0xb4, 0x08, 0x7c, 0x0d, 0x4a, 0xd8, 0x65, 0x3b, 0x4f, 0xa8, 0x47,
	0xb2, 0x08, 0x3d, 0xd9, 0xfc, 0xf3, 0xff, 0x6e, 0xbe, 0xe5, 0x89, 0xcc, 0xe9, 0x49, 0x51, 0x74,
	0x7a, 0xf2, 0x01, 0xbf, 0x02, 0xa5, 0xe8, 0x9a, 0x7c, 0xae, 0x96, 0xe4, 0xd5, 0x67, 0xf6, 0x22,
	0xf6, 0x23, 0x3b, 0x21, 0xc0, 0x0e, 0x28, 0xaf, 0x1c, 0x0f, 0x6f, 0x9d, 0x5f, 0x28, 0x51, 0x8f,
	0x9b, 0x4a, 0xeb, 0x24, 0xfb, 0xe9, 0x7b, 0x08, 0xd9, 0x1f, 0x69, 0x5f, 0xff, 0xae, 0x80, 0x4a,
	0x66, 0xed, 0xe0, 0x73, 0xa0, 0x76, 0x27, 0x93, 0xc1, 0x74, 0x3e, 0x99, 0x76, 0xa7, 0xb3, 0xc9,
	0x7c, 0x36, 0x9a, 0xbc, 0x1b, 0xf4, 0xac, 0xd7, 0xd6, 0xa0, 0x5f, 0xcb, 0xc1, 0x33, 0x50, 0xbd,
	0x87, 0x8e, 0xdf, 0xd6, 0x14, 0xd8, 0x04, 0xcf, 0xef, 0x39, 0xcd, 0xe1, 0xb8, 0xf7, 0x76, 0xd0,
	0x9f, 0x5b, 0x23, 0x73, 0x3c, 0x1b, 0xf5, 0x6b, 0x79, 0xf8, 0x02, 0xd4, 0x0f, 0x32, 0xc6, 0xb3,
	0x69, 0x4c, 0x29, 0xc0, 0x3a, 0xf8, 0xfc, 0x20, 0xc5, 0x1c, 0x4f, 0xdf, 0xd4, 0x8a, 0xe6, 0xf0,
	0xfd, 0x6d, 0x43, 0xf9, 0x70, 0xdb, 0x50, 0xfe, 0xb9, 0x6d, 0x28, 0xbf, 0xdd, 0x35, 0x72, 0x1f,
	0xee, 0x1a, 0xb9, 0xbf, 0xee, 0x1a, 0xb9, 0x9f, 0x3a, 0x6b, 0x47, 0x6c, 0x76, 0x0b, 0x7d, 0xc9,
	0x5c, 0x23, 0x19, 0xda, 0xb7, 0x5b, 0xbc, 0xe0, 0xa9, 0x61, 0x5c, 0x75, 0xbe, 0x37, 0xae, 0xd3,
	0x5f, 0xb2, 0xb8, 0xb9, 0xa4, 0x7c, 0x51, 0x92, 0x3f, 0xd7, 0xef, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0x25, 0x1a, 0x9f, 0xb1, 0x05, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBridge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.VotesNeeded != 0 {
		i = encodeVarintBridge(dAtA, i, uint64(m.VotesNeeded))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBridge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintBridge(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AssetID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourceChain) > 0 {
		i -= len(m.SourceChain)
		copy(dAtA[i:], m.SourceChain)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.SourceChain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExternalConfirmations != 0 {
		i = encodeVarintBridge(dAtA, i, uint64(m.ExternalConfirmations))
		i--
		dAtA[i] = 0x20
	}
	if m.Exponent != 0 {
		i = encodeVarintBridge(dAtA, i, uint64(m.Exponent))
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintBridge(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBridge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *InboundTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboundTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InboundTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Finalized {
		i--
		if m.Finalized {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintBridge(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBridge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.AssetId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBridge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ExternalHeight != 0 {
		i = encodeVarintBridge(dAtA, i, uint64(m.ExternalHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ExternalId) > 0 {
		i -= len(m.ExternalId)
		copy(dAtA[i:], m.ExternalId)
		i = encodeVarintBridge(dAtA, i, uint64(len(m.ExternalId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBridge(dAtA []byte, offset int, v uint64) int {
	offset -= sovBridge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovBridge(uint64(l))
		}
	}
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovBridge(uint64(l))
		}
	}
	if m.VotesNeeded != 0 {
		n += 1 + sovBridge(uint64(m.VotesNeeded))
	}
	l = m.Fee.Size()
	n += 1 + l + sovBridge(uint64(l))
	return n
}

func (m *AssetID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceChain)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	return n
}

func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Id.Size()
	n += 1 + l + sovBridge(uint64(l))
	if m.Status != 0 {
		n += 1 + sovBridge(uint64(m.Status))
	}
	if m.Exponent != 0 {
		n += 1 + sovBridge(uint64(m.Exponent))
	}
	if m.ExternalConfirmations != 0 {
		n += 1 + sovBridge(uint64(m.ExternalConfirmations))
	}
	return n
}

func (m *InboundTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExternalId)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	if m.ExternalHeight != 0 {
		n += 1 + sovBridge(uint64(m.ExternalHeight))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovBridge(uint64(l))
	}
	l = m.AssetId.Size()
	n += 1 + l + sovBridge(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovBridge(uint64(l))
	if len(m.Voters) > 0 {
		for _, s := range m.Voters {
			l = len(s)
			n += 1 + l + sovBridge(uint64(l))
		}
	}
	if m.Finalized {
		n += 2
	}
	return n
}

func sovBridge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBridge(x uint64) (n int) {
	return sovBridge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, Asset{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotesNeeded", wireType)
			}
			m.VotesNeeded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotesNeeded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssetID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssetID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= AssetStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			m.Exponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exponent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalConfirmations", wireType)
			}
			m.ExternalConfirmations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExternalConfirmations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InboundTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InboundTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboundTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalHeight", wireType)
			}
			m.ExternalHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExternalHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AssetId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalized", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Finalized = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBridge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBridge
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBridge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBridge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBridge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBridge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBridge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBridge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBridge = fmt.Errorf("proto: unexpected end of group")
)