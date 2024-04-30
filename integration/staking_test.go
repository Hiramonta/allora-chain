package integration_test

import (
	cosmosMath "cosmossdk.io/math"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

// register alice as a reputer in topic 1, then check success
func StakeAliceAsReputerTopic1(m TestMetadata) {
	addStake := &emissionstypes.MsgAddStake{
		Sender:  m.n.AliceAddr,
		TopicId: 1,
		Amount:  cosmosMath.NewUint(1000000),
	}

	txResp, err := m.n.Client.BroadcastTx(m.ctx, m.n.AliceAcc, addStake)
	require.NoError(m.t, err)
	_, err = m.n.Client.WaitForTx(m.ctx, txResp.TxHash)
	require.NoError(m.t, err)

	// Check Alice has stake on the topic
	aliceStaked, err := m.n.QueryEmissions.GetReputerStakeInTopic(
		m.ctx,
		&emissionstypes.QueryReputerStakeInTopicRequest{
			TopicId: 1,
			Address: m.n.AliceAddr,
		},
	)
	require.NoError(m.t, err)
	require.True(m.t, aliceStaked.Amount.Equal(cosmosMath.NewUint(1000000)))
}

// Register two actors and check their registrations went through
func StakingChecks(m TestMetadata) {
	m.t.Log("--- Staking Alice as Reputer ---")
	StakeAliceAsReputerTopic1(m)
	// m.t.Log("--- Registering Bob as Worker ---")
	// RegisterBobAsWorkerTopic1(m)
}