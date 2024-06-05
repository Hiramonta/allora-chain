package integration_test

import (
	upgradetypes "cosmossdk.io/x/upgrade/types"
	testCommon "github.com/allora-network/allora-chain/test/common"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/require"
)

func getDepositRequired(m testCommon.TestConfig) sdktypes.Coin {
	queryGovParamsResponse, err := m.Client.QueryGov().Params(m.Ctx, &govtypesv1.QueryParamsRequest{})
	require.NoError(m.T, err)
	return queryGovParamsResponse.Params.ExpeditedMinDeposit[0]
}

func proposeUpgrade(m testCommon.TestConfig) uint64 {
	name := "vIntegration"
	msgSoftwareUpgrade := &upgradetypes.MsgSoftwareUpgrade{
		Authority: authtypes.NewModuleAddress("gov").String(),
		Plan: upgradetypes.Plan{
			Name:   name,
			Height: 1e6,
			Info:   "{}",
		},
	}
	txResp, err := m.Client.BroadcastTx(m.Ctx, m.AliceAcc, msgSoftwareUpgrade)
	require.NoError(m.T, err)
	_, err = m.Client.WaitForTx(m.Ctx, txResp.TxHash)
	require.NoError(m.T, err)
	msgSoftwareUpgradeResponse := &upgradetypes.MsgSoftwareUpgradeResponse{}
	err = txResp.Decode(msgSoftwareUpgradeResponse)
	require.NoError(m.T, err)
	require.NotNil(m.T, msgSoftwareUpgradeResponse)
	return 3

	// here's how to do a vanilla governance proposal - still not sure
	// if we have to propose this by hand ourselves or use the upgrade
	// module to do it for us
	//
	//	submitProposalMsg := &govtypesv1.MsgSubmitProposal{
	//		Title:    title,
	//		Summary:  summary,
	//		Proposer: m.AliceAddr,
	//		Metadata: fmt.Sprintf(
	//			"{title:\"%s\",summary:\"%s\"}", title, summary,
	//		), // metadata must match title and summary exactly
	//		Expedited: true,
	//		InitialDeposit: sdktypes.NewCoins(
	//			getDepositRequired(m),
	//		),
	//	}
	//	txResp, err := m.Client.BroadcastTx(m.Ctx, m.AliceAcc, submitProposalMsg)
	//	require.NoError(m.T, err)
	//	_, err = m.Client.WaitForTx(m.Ctx, txResp.TxHash)
	//	require.NoError(m.T, err)
	//	submitProposalMsgResponse := &govtypesv1.MsgSubmitProposalResponse{}
	//	err = txResp.Decode(submitProposalMsgResponse)
	//	require.NoError(m.T, err)
	//	require.NotNil(m.T, submitProposalMsgResponse.ProposalId)
	//	return submitProposalMsgResponse.ProposalId
}

func UpgradeChecks(m testCommon.TestConfig) {
	m.T.Log("--- Propose Upgrade to vIntegration software version from v0 ---")
	proposalId := proposeUpgrade(m)
	m.T.Logf("--- Vote on Upgrade Proposal %d ---", proposalId)
}