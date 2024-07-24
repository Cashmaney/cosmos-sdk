package simulation

import (
	"context"
	"math/rand"

	coreaddress "cosmossdk.io/core/address"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// OpWeightSubmitTextProposal app params key for text proposal
const OpWeightSubmitTextProposal = "op_weight_submit_text_proposal"

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs() []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsgX(
			OpWeightSubmitTextProposal,
			DefaultWeightTextProposal,
			SimulateTextProposal,
		),
	}
}

// SimulateTextProposal returns a random text proposal content.
// A text proposal is a proposal that contains no msgs.
func SimulateTextProposal(_ context.Context, r *rand.Rand, _ []simtypes.Account, _ coreaddress.Codec) (sdk.Msg, error) {
	return nil, nil
}

// ProposalContents defines the module weighted proposals' contents
//
//nolint:staticcheck // used for legacy testing
func ProposalContents() []simtypes.WeightedProposalContent {
	return []simtypes.WeightedProposalContent{
		simulation.NewWeightedProposalContent(
			OpWeightMsgDeposit,
			DefaultWeightTextProposal,
			SimulateLegacyTextProposalContent,
		),
	}
}

// SimulateTextProposalContent returns a random text proposal content.
//
//nolint:staticcheck // used for legacy testing
func SimulateLegacyTextProposalContent(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) simtypes.Content {
	return v1beta1.NewTextProposal(
		simtypes.RandStringOfLength(r, 140),
		simtypes.RandStringOfLength(r, 5000),
	)
}
