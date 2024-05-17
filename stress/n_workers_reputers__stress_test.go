package stress_test

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	cosmossdk_io_math "cosmossdk.io/math"
	"github.com/allora-network/allora-chain/app/params"
	alloraMath "github.com/allora-network/allora-chain/math"
	"github.com/allora-network/allora-chain/x/emissions/types"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/stretchr/testify/require"
)

const MAX_ITERATIONS = 10000

const defaultEpochLength = 10
const approximateBlockLengthSeconds = 5
const minWaitingNumberofEpochs = 3

func getNonZeroTopicEpochLastRan(ctx context.Context, query emissionstypes.QueryClient, topicID uint64, maxRetries int) (*emissionstypes.Topic, error) {
	sleepingTimeBlocks := defaultEpochLength
	// Retry loop for a maximum of 5 times
	for retries := 0; retries < maxRetries; retries++ {
		topicResponse, err := query.GetTopic(ctx, &emissionstypes.QueryTopicRequest{TopicId: topicID})
		if err == nil {
			storedTopic := topicResponse.Topic
			if storedTopic.EpochLastEnded != 0 {
				sleepingTimeSeconds := time.Duration(minWaitingNumberofEpochs*storedTopic.EpochLength*approximateBlockLengthSeconds) * time.Second
				fmt.Println(time.Now(), " Topic found, sleeping...", sleepingTimeSeconds)
				time.Sleep(sleepingTimeSeconds)
				fmt.Println(time.Now(), " Slept.")
				return topicResponse.Topic, nil
			}
			sleepingTimeBlocks = int(storedTopic.EpochLength)
		} else {
			fmt.Println("Error getting topic, retry...", err)
		}
		// Sleep for a while before retrying
		fmt.Println("Retrying sleeping for a default epoch, retry ", retries, " for sleeping time ", sleepingTimeBlocks)
		time.Sleep(time.Duration(sleepingTimeBlocks*approximateBlockLengthSeconds) * time.Second)
	}

	return nil, errors.New("topicEpochLastRan is still 0 after retrying")
}

func generateWorkerAttributedValueLosses(m TestMetadata, workerAddresses map[string]string, lowLimit, sum int) []*types.WorkerAttributedValue {
	values := make([]*types.WorkerAttributedValue, 0)
	for key := range workerAddresses {
		values = append(values, &types.WorkerAttributedValue{
			Worker: workerAddresses[key],
			Value:  alloraMath.NewDecFromInt64(int64(rand.Intn(lowLimit) + sum)),
		})
	}
	return values
}

func generateWithheldWorkerAttributedValueLosses(m TestMetadata, workerAddresses map[string]string, lowLimit, sum int) []*types.WithheldWorkerAttributedValue {
	values := make([]*types.WithheldWorkerAttributedValue, 0)
	for key := range workerAddresses {
		values = append(values, &types.WithheldWorkerAttributedValue{
			Worker: workerAddresses[key],
			Value:  alloraMath.NewDecFromInt64(int64(rand.Intn(lowLimit) + sum)),
		})
	}
	return values
}

func generateSingleWorkerBundle(m TestMetadata, topic *types.Topic, blockHeight int64,
	workerAddressName string, workerAddresses map[string]string) *types.WorkerDataBundle {
	// Nonce: calculate from EpochLastRan + EpochLength
	topicId := topic.Id
	// Iterate workerAddresses to get the worker address, and generate as many forecasts as there are workers
	forecastElements := make([]*types.ForecastElement, 0)
	for key := range workerAddresses {
		forecastElements = append(forecastElements, &types.ForecastElement{
			Inferer: workerAddresses[key],
			Value:   alloraMath.NewDecFromInt64(int64(rand.Intn(51) + 50)),
		})
	}
	infererAddress := workerAddresses[workerAddressName]
	infererValue := alloraMath.NewDecFromInt64(int64(rand.Intn(300) + 3000))

	// Create a MsgInsertBulkReputerPayload message
	workerDataBundle := &types.WorkerDataBundle{
		Worker: infererAddress,
		InferenceForecastsBundle: &types.InferenceForecastBundle{
			Inference: &types.Inference{
				TopicId:     topicId,
				BlockHeight: blockHeight,
				Inferer:     infererAddress,
				Value:       infererValue,
			},
			Forecast: &types.Forecast{
				TopicId:          topicId,
				BlockHeight:      blockHeight,
				Forecaster:       infererAddress,
				ForecastElements: forecastElements,
			},
		},
	}

	// Sign
	src := make([]byte, 0)
	src, err := workerDataBundle.InferenceForecastsBundle.XXX_Marshal(src, true)
	require.NoError(m.t, err, "Marshall reputer value bundle should not return an error")

	sig, pubKey, err := m.n.Client.Context().Keyring.Sign(workerAddressName, src, signing.SignMode_SIGN_MODE_DIRECT)
	require.NoError(m.t, err, "Sign should not return an error")
	workerPublicKeyBytes := pubKey.Bytes()
	workerDataBundle.InferencesForecastsBundleSignature = sig
	workerDataBundle.Pubkey = hex.EncodeToString(workerPublicKeyBytes)

	return workerDataBundle
}

// Generate the same valueBundle for a reputer
func generateValueBundle(m TestMetadata, topicId uint64, workerAddresses map[string]string, reputerNonce, workerNonce *types.Nonce) types.ValueBundle {
	return types.ValueBundle{
		TopicId:                topicId,
		CombinedValue:          alloraMath.NewDecFromInt64(100),
		InfererValues:          generateWorkerAttributedValueLosses(m, workerAddresses, 3000, 3500),
		ForecasterValues:       generateWorkerAttributedValueLosses(m, workerAddresses, 50, 50),
		NaiveValue:             alloraMath.NewDecFromInt64(100),
		OneOutInfererValues:    generateWithheldWorkerAttributedValueLosses(m, workerAddresses, 50, 50),
		OneOutForecasterValues: generateWithheldWorkerAttributedValueLosses(m, workerAddresses, 50, 50),
		OneInForecasterValues:  generateWorkerAttributedValueLosses(m, workerAddresses, 50, 50),
		ReputerRequestNonce: &types.ReputerRequestNonce{
			ReputerNonce: reputerNonce,
			WorkerNonce:  workerNonce,
		},
	}
}

// Inserts worker bulk, given a topic, blockHeight, and leader worker address (which should exist in the keyring)
func InsertLeaderWorkerBulk(
	m TestMetadata,
	topic *types.Topic,
	blockHeight int64,
	leaderWorkerAccountName, leaderWorkerAddress string,
	WorkerDataBundles []*types.WorkerDataBundle) error {

	topicId := topic.Id
	nonce := emissionstypes.Nonce{BlockHeight: blockHeight}

	// Create a MsgInsertBulkReputerPayload message
	workerMsg := &types.MsgInsertBulkWorkerPayload{
		Sender:            leaderWorkerAddress,
		Nonce:             &nonce,
		TopicId:           topicId,
		WorkerDataBundles: WorkerDataBundles,
	}
	// serialize workerMsg to json and print
	LeaderAcc, err := m.n.Client.AccountRegistry.GetByName(leaderWorkerAccountName)
	if err != nil {
		fmt.Println("Error getting leader worker account: ", leaderWorkerAccountName, " - ", err)
		return err
	}
	txResp, err := m.n.Client.BroadcastTx(m.ctx, LeaderAcc, workerMsg)
	if err != nil {
		fmt.Println("Error broadcasting worker bulk: ", err)
		return err
	}
	_, err = m.n.Client.WaitForTx(m.ctx, txResp.TxHash)
	if err != nil {
		fmt.Println("Error waiting for worker bulk: ", err)
		return err
	}
	return nil
}

// Worker Bob inserts bulk inference and forecast
func InsertWorkerBulk(m TestMetadata, topic *types.Topic, leaderWorkerAccountName string, workerAddresses map[string]string) (int64, int64) {
	// Get fresh topic to use its EpochLastRan
	topicResponse, err := m.n.QueryEmissions.GetTopic(m.ctx, &emissionstypes.QueryTopicRequest{TopicId: topic.Id})
	require.NoError(m.t, err)
	freshTopic := topicResponse.Topic
	blockHeight := freshTopic.EpochLastEnded + freshTopic.EpochLength
	// Get Bundles
	workerDataBundles := make([]*types.WorkerDataBundle, 0)
	for key := range workerAddresses {
		workerDataBundles = append(workerDataBundles, generateSingleWorkerBundle(m, topic, blockHeight, key, workerAddresses))
	}
	leaderWorkerAddress := workerAddresses[leaderWorkerAccountName]
	fmt.Println("Inserting worker bulk for blockHeight: ", blockHeight, "leader name: ", leaderWorkerAccountName, ", addr: ", leaderWorkerAddress, " len: ", len(workerDataBundles))
	InsertLeaderWorkerBulk(m, freshTopic, blockHeight, leaderWorkerAccountName, leaderWorkerAddress, workerDataBundles)
	return blockHeight, blockHeight - freshTopic.EpochLength
}

// Generate a ReputerValueBundle
func generateSingleReputerValueBundle(
	m TestMetadata,
	reputerAddressName, reputerAddress string,
	valueBundle types.ValueBundle) *types.ReputerValueBundle {

	valueBundle.Reputer = reputerAddress
	// Sign
	src := make([]byte, 0)
	src, err := valueBundle.XXX_Marshal(src, true)
	require.NoError(m.t, err, "Marshall reputer value bundle should not return an error")

	valueBundleSignature, pubKey, err := m.n.Client.Context().Keyring.Sign(reputerAddressName, src, signing.SignMode_SIGN_MODE_DIRECT)
	require.NoError(m.t, err, "Sign should not return an error")
	reputerPublicKeyBytes := pubKey.Bytes()

	// Create a MsgInsertBulkReputerPayload message
	reputerValueBundle := &types.ReputerValueBundle{
		ValueBundle: &valueBundle,
		Signature:   valueBundleSignature,
		Pubkey:      hex.EncodeToString(reputerPublicKeyBytes),
	}

	return reputerValueBundle
}

func generateReputerValueBundleMsg(m TestMetadata,
	topicId uint64,
	reputerValueBundles []*types.ReputerValueBundle,
	leaderReputerAccountName, leaderReputerAddress string,
	reputerNonce, workerNonce *emissionstypes.Nonce) *types.MsgInsertBulkReputerPayload {

	return &types.MsgInsertBulkReputerPayload{
		Sender:  leaderReputerAddress,
		TopicId: topicId,
		ReputerRequestNonce: &types.ReputerRequestNonce{
			ReputerNonce: reputerNonce,
			WorkerNonce:  workerNonce,
		},
		ReputerValueBundles: reputerValueBundles,
	}
}

func InsertReputerBulk(m TestMetadata,
	topic *types.Topic,
	leaderReputerAccountName string,
	reputerAddresses, workerAddresses map[string]string,
	BlockHeightCurrent, BlockHeightEval int64) error {
	// Nonce: calculate from EpochLastRan + EpochLength
	topicId := topic.Id
	// Nonces are last two blockHeights
	reputerNonce := &types.Nonce{
		BlockHeight: BlockHeightCurrent,
	}
	workerNonce := &types.Nonce{
		BlockHeight: BlockHeightEval,
	}
	leaderReputerAddress := reputerAddresses[leaderReputerAccountName]
	valueBundle := generateValueBundle(m, topicId, workerAddresses, reputerNonce, workerNonce)
	reputerValueBundles := make([]*types.ReputerValueBundle, 0)
	for reputerAddressName := range reputerAddresses {
		reputerAddress := reputerAddresses[reputerAddressName]
		reputerValueBundle := generateSingleReputerValueBundle(m, reputerAddressName, reputerAddress, valueBundle)
		reputerValueBundles = append(reputerValueBundles, reputerValueBundle)
	}

	reputerValueBundleMsg := generateReputerValueBundleMsg(m, topicId, reputerValueBundles, leaderReputerAccountName, leaderReputerAddress, reputerNonce, workerNonce)
	LeaderAcc, err := m.n.Client.AccountRegistry.GetByName(leaderReputerAccountName)
	if err != nil {
		fmt.Println("Error getting leader worker account: ", leaderReputerAccountName, " - ", err)
		return err
	}
	txResp, err := m.n.Client.BroadcastTx(m.ctx, LeaderAcc, reputerValueBundleMsg)
	if err != nil {
		fmt.Println("Error broadcasting reputer value bundle: ", err)
		return err
	}
	_, err = m.n.Client.WaitForTx(m.ctx, txResp.TxHash)
	return nil
}

// Register two actors and check their registrations went through
func WorkerReputerLoop(m TestMetadata, topicId uint64) {
	const stakeToAdd uint64 = 10000
	workerAddresses := make(map[string]string)
	reputerAddresses := make(map[string]string)

	// Make a loop, in each iteration
	// 1. generate a new bech32 reputer account and a bech32 worker account. Store them in a slice
	// 2. Pass worker slice in the call to insertWorkerBulk
	// 3. Pass reputer slice in the call to insertReputerBulk
	// 4. sleep one epoch, then repeat.

	// Get fresh topic
	topic, err := getNonZeroTopicEpochLastRan(m.ctx, m.n.QueryEmissions, 1, 5)
	if err != nil {
		m.t.Log("--- Failed getting a topic that was ran ---")
		require.NoError(m.t, err)
	}

	// Translate the epoch length into time
	epochTimeSeconds := time.Duration(topic.EpochLength*approximateBlockLengthSeconds) * time.Second
	for i := 0; i < MAX_ITERATIONS; i++ {
		start := time.Now()

		fmt.Println("iteration: ", i, " / ", MAX_ITERATIONS)
		// Generate new worker accounts
		workerAccountName := "stressWorker" + strconv.Itoa(i)
		workerAccount, _, err := m.n.Client.AccountRegistry.Create(workerAccountName)
		if err != nil {
			fmt.Println("Error creating worker address: ", workerAccountName, " - ", err)
			continue
		}
		workerAddress, err := workerAccount.Address(params.HumanCoinUnit)
		if err != nil {
			fmt.Println("Error getting worker address: ", workerAccountName, " - ", err)
			continue
		}
		fmt.Println("Worker address: ", workerAddress)
		err = fundAccount(m, m.n.FaucetAcc, m.n.FaucetAddr, workerAddress, 100000)
		if err != nil {
			fmt.Println("Error funding worker address: ", workerAddress, " - ", err)
			continue
		}
		err = RegisterWorkerForTopic(m, workerAddress, workerAccount, topicId)
		if err != nil {
			fmt.Println("Error registering worker address: ", workerAddress, " - ", err)
			continue
		}
		workerAddresses[workerAccountName] = workerAddress

		// Generate new reputer account
		reputerAccountName := "stressReputer" + strconv.Itoa(i)
		reputerAccount, _, err := m.n.Client.AccountRegistry.Create(reputerAccountName)
		if err != nil {
			fmt.Println("Error creating reputer address: ", reputerAccountName, " - ", err)
			continue
		}
		reputerAddress, err := reputerAccount.Address(params.HumanCoinUnit)
		if err != nil {
			fmt.Println("Error getting reputer address: ", reputerAccountName, " - ", err)
			continue
		}
		fmt.Println("Reputer address: ", reputerAddress)
		err = fundAccount(m, m.n.FaucetAcc, m.n.FaucetAddr, reputerAddress, 100000)
		if err != nil {
			fmt.Println("Error funding reputer address: ", reputerAddress, " - ", err)
			continue
		}
		err = RegisterReputerForTopic(m, reputerAddress, reputerAccount, topicId)
		if err != nil {
			fmt.Println("Error registering reputer address: ", reputerAddress, " - ", err)
			continue
		}
		err = StakeReputer(m, topicId, reputerAddress, reputerAccount, stakeToAdd)
		if err != nil {
			fmt.Println("Error staking reputer address: ", reputerAddress, " - ", err)
			continue
		}
		reputerAddresses[reputerAccountName] = reputerAddress

		// Choose one random leader from the worker accounts
		leaderWorkerAccountName, leaderWorkerAddress, err := GetRandomMapEntryValue(workerAddresses)
		if err != nil {
			fmt.Println("Error getting random worker address: ", err)
			continue
		}
		// Insert worker
		m.t.Log("--- Insert Worker Bulk --- with leader: ", leaderWorkerAccountName, " and worker address: ", leaderWorkerAddress)
		blockHeightCurrent, blockHeightEval := InsertWorkerBulk(m, topic, leaderWorkerAccountName, workerAddresses)

		// Insert reputer bulk
		// Choose one random leader from reputer accounts
		leaderReputerAccountName, _, err := GetRandomMapEntryValue(reputerAddresses)
		if err != nil {
			fmt.Println("Error getting random worker address: ", err)
			continue
		}

		startReputer := time.Now()
		InsertReputerBulk(m, topic, leaderReputerAccountName, reputerAddresses, workerAddresses, blockHeightCurrent, blockHeightEval)
		elapsedBulk := time.Since(startReputer)
		fmt.Println("Insert Reputer Elapsed time:", elapsedBulk)

		// Sleep for one epoch
		elapsed := time.Since(start)
		fmt.Println("Insert Worker Elapsed time:", elapsed, " BlockHeightCurrent: ", blockHeightCurrent, " BlockHeightEval: ", blockHeightEval)
		sleepingTimeSeconds := epochTimeSeconds - elapsed
		fmt.Println(time.Now(), " Sleeping...", sleepingTimeSeconds, ", epoch length: ", epochTimeSeconds)
		time.Sleep(sleepingTimeSeconds)
		fmt.Println(time.Now(), " Slept.")
	}

}

func GetRandomMapEntryValue(workerAddresses map[string]string) (string, string, error) {
	// Get the number of entries in the map
	numEntries := len(workerAddresses)
	if numEntries == 0 {
		return "", "", fmt.Errorf("map is empty")
	}

	// Generate a random index
	randomIndex := rand.Intn(numEntries)

	// Iterate over the map to find the entry at the random index
	var randomKey string
	var i int
	for key := range workerAddresses {
		if i == randomIndex {
			randomKey = key
			break
		}
		i++
	}

	// Return the value corresponding to the randomly selected key
	return randomKey, workerAddresses[randomKey], nil
}

func fundAccount(m TestMetadata, senderAccount cosmosaccount.Account, senderAddress, address string, amount int64) error {
	initialCoins := sdktypes.NewCoins(sdktypes.NewCoin(params.BaseCoinUnit, cosmossdk_io_math.NewInt(amount)))
	// Fund the account from faucet account
	sendMsg := &banktypes.MsgSend{
		FromAddress: senderAddress,
		ToAddress:   address,
		Amount:      initialCoins,
	}
	_, err := m.n.Client.BroadcastTx(m.ctx, senderAccount, sendMsg)
	if err != nil {
		fmt.Println("Error funding worker address: ", err)
		return err
	}
	return nil
}