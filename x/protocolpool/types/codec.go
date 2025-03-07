package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(ir codectypes.InterfaceRegistry) {
	ir.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgFundCommunityPool{},
		&MsgCommunityPoolSpend{},
		&MsgSubmitBudgetProposal{},
		&MsgClaimBudget{},
		&MsgCreateContinuousFund{},
		&MsgCancelContinuousFund{},
		&MsgWithdrawContinuousFund{},
	)

	msgservice.RegisterMsgServiceDesc(ir, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers the necessary x/protocolpool interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
//
// NOTE amino msg name paths are shorted due to the 40-character limit for amino.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgFundCommunityPool{}, "cosmos-sdk/pp/MsgFundCommunityPool")
	legacy.RegisterAminoMsg(cdc, &MsgCommunityPoolSpend{}, "cosmos-sdk/pp/MsgCommunityPoolSpend")
	legacy.RegisterAminoMsg(cdc, &MsgSubmitBudgetProposal{}, "cosmos-sdk/pp/MsgSubmitBudgetProposal")
	legacy.RegisterAminoMsg(cdc, &MsgClaimBudget{}, "cosmos-sdk/pp/MsgClaimBudget")
	legacy.RegisterAminoMsg(cdc, &MsgCreateContinuousFund{}, "cosmos-sdk/pp/MsgCreateContinuousFund")
	legacy.RegisterAminoMsg(cdc, &MsgCancelContinuousFund{}, "cosmos-sdk/pp/MsgCancelContinuousFund")
	legacy.RegisterAminoMsg(cdc, &MsgWithdrawContinuousFund{}, "cosmos-sdk/pp/MsgWithdrawContinuousFund")

	cdc.RegisterConcrete(&Budget{}, "cosmos-sdk/pp/Budget", nil)
	cdc.RegisterConcrete(&ContinuousFund{}, "cosmos-sdk/pp/ContinuousFund", nil)
	cdc.RegisterConcrete(&DistributionAmount{}, "cosmos-sdk/pp/DistributionAmount", nil)
	cdc.RegisterConcrete(&Params{}, "cosmos-sdk/pp/Params", nil)
}
