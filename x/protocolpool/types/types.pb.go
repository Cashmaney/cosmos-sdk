// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/protocolpool/v1/types.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Budget defines the fields of a budget proposal.
type Budget struct {
	// recipient_address is the address of the recipient who can claim the budget.
	RecipientAddress string `protobuf:"bytes,1,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	// claimed_amount is the total amount claimed from the total budget amount requested.
	ClaimedAmount *types.Coin `protobuf:"bytes,2,opt,name=claimed_amount,json=claimedAmount,proto3" json:"claimed_amount,omitempty"`
	// last_claimed_at is the time when the budget was last successfully claimed or distributed.
	// It is used to track the next starting claim time for fund distribution.
	LastClaimedAt time.Time `protobuf:"bytes,3,opt,name=last_claimed_at,json=lastClaimedAt,proto3,stdtime" json:"last_claimed_at"`
	// tranches_left is the number of tranches left for the amount to be distributed.
	TranchesLeft uint64 `protobuf:"varint,4,opt,name=tranches_left,json=tranchesLeft,proto3" json:"tranches_left,omitempty"`
	// budget_per_tranche is the amount allocated per tranche.
	BudgetPerTranche types.Coin `protobuf:"bytes,5,opt,name=budget_per_tranche,json=budgetPerTranche,proto3" json:"budget_per_tranche"`
	// Period is the time interval(number of seconds) at which funds distribution should be performed.
	// For example, if a period is set to 3600, it represents an action that
	// should occur every hour (3600 seconds).
	Period time.Duration `protobuf:"bytes,6,opt,name=period,proto3,stdduration" json:"period"`
}

func (m *Budget) Reset()         { *m = Budget{} }
func (m *Budget) String() string { return proto.CompactTextString(m) }
func (*Budget) ProtoMessage()    {}
func (*Budget) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b7d0ea246d7f44, []int{0}
}
func (m *Budget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Budget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Budget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Budget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Budget.Merge(m, src)
}
func (m *Budget) XXX_Size() int {
	return m.Size()
}
func (m *Budget) XXX_DiscardUnknown() {
	xxx_messageInfo_Budget.DiscardUnknown(m)
}

var xxx_messageInfo_Budget proto.InternalMessageInfo

func (m *Budget) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *Budget) GetClaimedAmount() *types.Coin {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

func (m *Budget) GetLastClaimedAt() time.Time {
	if m != nil {
		return m.LastClaimedAt
	}
	return time.Time{}
}

func (m *Budget) GetTranchesLeft() uint64 {
	if m != nil {
		return m.TranchesLeft
	}
	return 0
}

func (m *Budget) GetBudgetPerTranche() types.Coin {
	if m != nil {
		return m.BudgetPerTranche
	}
	return types.Coin{}
}

func (m *Budget) GetPeriod() time.Duration {
	if m != nil {
		return m.Period
	}
	return 0
}

// ContinuousFund defines the fields of continuous fund proposal.
type ContinuousFund struct {
	// Recipient address of the account receiving funds.
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// Percentage is the percentage of funds to be allocated from Community pool.
	Percentage cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=percentage,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"percentage"`
	// Optional, if expiry is set, removes the state object when expired.
	Expiry *time.Time `protobuf:"bytes,3,opt,name=expiry,proto3,stdtime" json:"expiry,omitempty"`
}

func (m *ContinuousFund) Reset()         { *m = ContinuousFund{} }
func (m *ContinuousFund) String() string { return proto.CompactTextString(m) }
func (*ContinuousFund) ProtoMessage()    {}
func (*ContinuousFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b7d0ea246d7f44, []int{1}
}
func (m *ContinuousFund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContinuousFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContinuousFund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContinuousFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousFund.Merge(m, src)
}
func (m *ContinuousFund) XXX_Size() int {
	return m.Size()
}
func (m *ContinuousFund) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousFund.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousFund proto.InternalMessageInfo

func (m *ContinuousFund) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *ContinuousFund) GetExpiry() *time.Time {
	if m != nil {
		return m.Expiry
	}
	return nil
}

// DistributionAmount is used to store the coins of periodic distributions.
type DistributionAmount struct {
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *DistributionAmount) Reset()         { *m = DistributionAmount{} }
func (m *DistributionAmount) String() string { return proto.CompactTextString(m) }
func (*DistributionAmount) ProtoMessage()    {}
func (*DistributionAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b7d0ea246d7f44, []int{2}
}
func (m *DistributionAmount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionAmount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionAmount.Merge(m, src)
}
func (m *DistributionAmount) XXX_Size() int {
	return m.Size()
}
func (m *DistributionAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionAmount.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionAmount proto.InternalMessageInfo

func (m *DistributionAmount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

// Params defines the parameters for the protocolpool module.
type Params struct {
	// enabled_distribution_denoms lists the denoms that are allowed to be distributed.
	// This is to avoid spending time distributing undesired tokens to continuous funds and budgets.
	EnabledDistributionDenoms []string `protobuf:"bytes,1,rep,name=enabled_distribution_denoms,json=enabledDistributionDenoms,proto3" json:"enabled_distribution_denoms,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b7d0ea246d7f44, []int{3}
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

func (m *Params) GetEnabledDistributionDenoms() []string {
	if m != nil {
		return m.EnabledDistributionDenoms
	}
	return nil
}

func init() {
	proto.RegisterType((*Budget)(nil), "cosmos.protocolpool.v1.Budget")
	proto.RegisterType((*ContinuousFund)(nil), "cosmos.protocolpool.v1.ContinuousFund")
	proto.RegisterType((*DistributionAmount)(nil), "cosmos.protocolpool.v1.DistributionAmount")
	proto.RegisterType((*Params)(nil), "cosmos.protocolpool.v1.Params")
}

func init() {
	proto.RegisterFile("cosmos/protocolpool/v1/types.proto", fileDescriptor_c1b7d0ea246d7f44)
}

var fileDescriptor_c1b7d0ea246d7f44 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x4b, 0xdc, 0x40,
	0x14, 0xde, 0xd4, 0xed, 0xd2, 0x9d, 0xaa, 0xd5, 0x41, 0x4a, 0x5c, 0x21, 0xbb, 0x6c, 0x2f, 0x8b,
	0x60, 0xc2, 0xb6, 0x50, 0x0a, 0x85, 0x52, 0xd7, 0xad, 0x14, 0x6a, 0xc1, 0xa6, 0x9e, 0x7a, 0x09,
	0x93, 0xe4, 0x19, 0x07, 0x93, 0x99, 0x30, 0x33, 0x11, 0xf7, 0x57, 0xd4, 0x63, 0xe9, 0x2f, 0x28,
	0x3d, 0x79, 0xf0, 0x47, 0x78, 0x14, 0xa1, 0x50, 0x7a, 0xd0, 0xa2, 0x07, 0xff, 0x46, 0xc9, 0x64,
	0xd6, 0x6e, 0x3d, 0x68, 0x2f, 0xbb, 0xe1, 0xbd, 0xef, 0xfb, 0xe6, 0xfb, 0xde, 0xcb, 0x04, 0x75,
	0x23, 0x2e, 0x33, 0x2e, 0xbd, 0x5c, 0x70, 0xc5, 0x23, 0x9e, 0xe6, 0x9c, 0xa7, 0xde, 0x5e, 0xdf,
	0x53, 0xa3, 0x1c, 0xa4, 0xab, 0xab, 0xf8, 0x71, 0x85, 0x71, 0x27, 0x31, 0xee, 0x5e, 0xbf, 0xb5,
	0x90, 0xf0, 0x84, 0xeb, 0xa2, 0x57, 0x3e, 0x55, 0xfd, 0xd6, 0x62, 0x85, 0x0e, 0xaa, 0xc6, 0x24,
	0xb5, 0xe5, 0x98, 0xc3, 0x42, 0x22, 0xc1, 0xdb, 0xeb, 0x87, 0xa0, 0x48, 0xdf, 0x8b, 0x38, 0x65,
	0xa6, 0xdf, 0x4e, 0x38, 0x4f, 0x52, 0xa8, 0xcc, 0x84, 0xc5, 0xb6, 0xa7, 0x68, 0x06, 0x52, 0x91,
	0x2c, 0x1f, 0x0b, 0xdc, 0x04, 0xc4, 0x85, 0x20, 0x8a, 0xf2, 0xb1, 0xc0, 0x3c, 0xc9, 0x28, 0xe3,
	0x9e, 0xfe, 0xad, 0x4a, 0xdd, 0x83, 0x29, 0xd4, 0x18, 0x14, 0x71, 0x02, 0x0a, 0xbf, 0x41, 0xf3,
	0x02, 0x22, 0x9a, 0x53, 0x60, 0x2a, 0x20, 0x71, 0x2c, 0x40, 0x4a, 0xdb, 0xea, 0x58, 0xbd, 0xe6,
	0xc0, 0x3e, 0x3d, 0x5a, 0x59, 0x30, 0x5e, 0x57, 0xab, 0xce, 0x47, 0x25, 0x28, 0x4b, 0xfc, 0xb9,
	0x6b, 0x8a, 0xa9, 0xe3, 0xd7, 0x68, 0x36, 0x4a, 0x09, 0xcd, 0x20, 0x0e, 0x48, 0xc6, 0x0b, 0xa6,
	0xec, 0x7b, 0x1d, 0xab, 0xf7, 0xf0, 0xe9, 0xa2, 0x6b, 0x04, 0xca, 0x78, 0xae, 0x89, 0xe7, 0xae,
	0x71, 0xca, 0xfc, 0x19, 0x43, 0x58, 0xd5, 0x78, 0xbc, 0x81, 0x1e, 0xa5, 0x44, 0xaa, 0xe0, 0x5a,
	0x46, 0xd9, 0x53, 0x5a, 0xa2, 0xe5, 0x56, 0x01, 0xdd, 0x71, 0x40, 0x77, 0x6b, 0x3c, 0x81, 0xc1,
	0x83, 0xe3, 0xb3, 0x76, 0xed, 0xe0, 0xbc, 0x6d, 0xf9, 0x33, 0x25, 0x79, 0xcd, 0x28, 0x2a, 0xfc,
	0x04, 0xcd, 0x28, 0x41, 0x58, 0xb4, 0x03, 0x32, 0x48, 0x61, 0x5b, 0xd9, 0xf5, 0x8e, 0xd5, 0xab,
	0xfb, 0xd3, 0xe3, 0xe2, 0x06, 0x6c, 0x2b, 0xfc, 0x1e, 0xe1, 0x50, 0x4f, 0x21, 0xc8, 0x41, 0x04,
	0xa6, 0x65, 0xdf, 0xbf, 0xc3, 0xf8, 0xa0, 0x5e, 0x1e, 0xea, 0xcf, 0x55, 0xd4, 0x4d, 0x10, 0x5b,
	0x15, 0x11, 0xbf, 0x44, 0x8d, 0x1c, 0x04, 0xe5, 0xb1, 0xdd, 0x30, 0x12, 0x37, 0x8d, 0x0f, 0xcd,
	0x66, 0x2a, 0xdf, 0x5f, 0x4a, 0xdf, 0x86, 0xd2, 0xfd, 0x61, 0xa1, 0xd9, 0x35, 0xce, 0x14, 0x65,
	0x05, 0x2f, 0xe4, 0x7a, 0xc1, 0x62, 0xfc, 0x1c, 0x35, 0xaf, 0xe7, 0x7c, 0xe7, 0x4a, 0xfe, 0x42,
	0xf1, 0x07, 0x84, 0x72, 0x10, 0x11, 0x30, 0x45, 0x12, 0xd0, 0x7b, 0x68, 0x0e, 0xfa, 0xe5, 0x81,
	0xbf, 0xce, 0xda, 0x4b, 0x15, 0x59, 0xc6, 0xbb, 0x2e, 0xe5, 0x5e, 0x46, 0xd4, 0x8e, 0xbb, 0x01,
	0x09, 0x89, 0x46, 0x43, 0x88, 0x4e, 0x8f, 0x56, 0x90, 0xd1, 0x1e, 0x42, 0xe4, 0x4f, 0x88, 0xe0,
	0x17, 0xa8, 0x01, 0xfb, 0x39, 0x15, 0xa3, 0xff, 0xd8, 0x49, 0x5d, 0xef, 0xc3, 0xe0, 0xbb, 0x9f,
	0x2d, 0x84, 0x87, 0x54, 0x2a, 0x41, 0xc3, 0xa2, 0x8c, 0x6e, 0xb6, 0x3d, 0x42, 0x0d, 0xf3, 0x9e,
	0x58, 0x9d, 0xa9, 0xdb, 0xc7, 0xbd, 0x5e, 0x5a, 0xff, 0x7e, 0xde, 0xee, 0x25, 0x54, 0xed, 0x14,
	0xa1, 0x1b, 0xf1, 0xcc, 0xdc, 0x20, 0xf3, 0xb7, 0x22, 0xe3, 0x5d, 0x73, 0x37, 0x4b, 0x82, 0xfc,
	0x7a, 0x75, 0xb8, 0x3c, 0x9d, 0xea, 0x54, 0x41, 0x79, 0x91, 0xe4, 0xb7, 0xab, 0xc3, 0x65, 0xcb,
	0x37, 0x07, 0x76, 0xdf, 0xa2, 0xc6, 0x26, 0x11, 0x24, 0x93, 0xf8, 0x15, 0x5a, 0x02, 0x46, 0xc2,
	0x14, 0xe2, 0x20, 0x9e, 0xb0, 0x18, 0xc4, 0xc0, 0x78, 0x26, 0xb5, 0xb3, 0xa6, 0xbf, 0x68, 0x20,
	0x93, 0x21, 0x86, 0x1a, 0x30, 0x78, 0x77, 0x7c, 0xe1, 0x58, 0x27, 0x17, 0x8e, 0xf5, 0xfb, 0xc2,
	0xb1, 0x0e, 0x2e, 0x9d, 0xda, 0xc9, 0xa5, 0x53, 0xfb, 0x79, 0xe9, 0xd4, 0x3e, 0xf5, 0x6f, 0xf5,
	0xba, 0xff, 0xef, 0x97, 0x45, 0x5b, 0x0f, 0x1b, 0xba, 0xf6, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc1, 0x5f, 0x7c, 0x1e, 0x7d, 0x04, 0x00, 0x00,
}

func (m *Budget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Budget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Budget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Period, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Period):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.BudgetPerTranche.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.TranchesLeft != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TranchesLeft))
		i--
		dAtA[i] = 0x20
	}
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastClaimedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastClaimedAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintTypes(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	if m.ClaimedAmount != nil {
		{
			size, err := m.ClaimedAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContinuousFund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContinuousFund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContinuousFund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiry != nil {
		n5, err5 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Expiry, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiry):])
		if err5 != nil {
			return 0, err5
		}
		i -= n5
		i = encodeVarintTypes(dAtA, i, uint64(n5))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Percentage.Size()
		i -= size
		if _, err := m.Percentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DistributionAmount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionAmount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionAmount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
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
	if len(m.EnabledDistributionDenoms) > 0 {
		for iNdEx := len(m.EnabledDistributionDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EnabledDistributionDenoms[iNdEx])
			copy(dAtA[i:], m.EnabledDistributionDenoms[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.EnabledDistributionDenoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Budget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ClaimedAmount != nil {
		l = m.ClaimedAmount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastClaimedAt)
	n += 1 + l + sovTypes(uint64(l))
	if m.TranchesLeft != 0 {
		n += 1 + sovTypes(uint64(m.TranchesLeft))
	}
	l = m.BudgetPerTranche.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Period)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *ContinuousFund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Percentage.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Expiry != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiry)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *DistributionAmount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EnabledDistributionDenoms) > 0 {
		for _, s := range m.EnabledDistributionDenoms {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Budget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Budget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Budget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClaimedAmount == nil {
				m.ClaimedAmount = &types.Coin{}
			}
			if err := m.ClaimedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastClaimedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastClaimedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TranchesLeft", wireType)
			}
			m.TranchesLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TranchesLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BudgetPerTranche", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BudgetPerTranche.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Period, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ContinuousFund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ContinuousFund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContinuousFund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiry == nil {
				m.Expiry = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *DistributionAmount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: DistributionAmount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionAmount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
				return fmt.Errorf("proto: wrong wireType = %d for field EnabledDistributionDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnabledDistributionDenoms = append(m.EnabledDistributionDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
