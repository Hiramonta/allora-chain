package types

import "cosmossdk.io/errors"

var (
	ErrTopicReputerStakeDoesNotExist                     = errors.Register(ModuleName, 1, "topic reputer stake does not exist")
	ErrIntegerUnderflowTopicReputerStake                 = errors.Register(ModuleName, 2, "integer underflow for topic reputer stake")
	ErrIntegerUnderflowTarget                            = errors.Register(ModuleName, 3, "integer underflow for target")
	ErrIntegerUnderflowTopicStake                        = errors.Register(ModuleName, 4, "integer underflow for topic stake")
	ErrIntegerUnderflowAllTopicStakeSum                  = errors.Register(ModuleName, 5, "integer underflow for all topic stake sum")
	ErrIntegerUnderflowTotalStake                        = errors.Register(ModuleName, 6, "integer underflow for total stake")
	ErrIterationLengthDoesNotMatch                       = errors.Register(ModuleName, 7, "iteration length does not match")
	ErrInvalidTopicId                                    = errors.Register(ModuleName, 8, "invalid topic ID")
	ErrReputerAlreadyRegisteredInTopic                   = errors.Register(ModuleName, 9, "reputer already registered in topic")
	ErrWorkerAlreadyRegisteredInTopic                    = errors.Register(ModuleName, 10, "worker already registered in topic")
	ErrAddressAlreadyRegisteredInATopic                  = errors.Register(ModuleName, 11, "address already registered in a topic")
	ErrAddressIsNotRegisteredInAnyTopic                  = errors.Register(ModuleName, 12, "address is not registered in any topic")
	ErrAddressIsNotRegisteredInThisTopic                 = errors.Register(ModuleName, 13, "address is not registered in this topic")
	ErrInsufficientStakeToRegister                       = errors.Register(ModuleName, 14, "insufficient stake to register")
	ErrLibP2PKeyRequired                                 = errors.Register(ModuleName, 15, "libp2p key required")
	ErrAddressNotRegistered                              = errors.Register(ModuleName, 16, "address not registered")
	ErrStakeTargetNotRegistered                          = errors.Register(ModuleName, 17, "stake target not registered")
	ErrTopicIdOfStakerAndTargetDoNotMatch                = errors.Register(ModuleName, 18, "topic ID of staker and target do not match")
	ErrInsufficientStakeToRemove                         = errors.Register(ModuleName, 19, "insufficient stake to remove")
	ErrNoStakeToRemove                                   = errors.Register(ModuleName, 20, "no stake to remove")
	ErrDoNotSetMapValueToZero                            = errors.Register(ModuleName, 21, "do not set map value to zero")
	ErrBlockHeightNegative                               = errors.Register(ModuleName, 22, "block height negative")
	ErrBlockHeightLessThanPrevious                       = errors.Register(ModuleName, 23, "block height less than previous")
	ErrConfirmRemoveStakeNoRemovalStarted                = errors.Register(ModuleName, 24, "confirm remove stake no removal started")
	ErrConfirmRemoveStakeTooEarly                        = errors.Register(ModuleName, 25, "confirm remove stake too early")
	ErrConfirmRemoveStakeTooLate                         = errors.Register(ModuleName, 26, "confirm remove stake too late")
	ErrScalarMultiplyNegative                            = errors.Register(ModuleName, 27, "scalar multiply negative")
	ErrDivideMapValuesByZero                             = errors.Register(ModuleName, 28, "divide map values by zero")
	ErrTopicIdListValueDecodeInvalidLength               = errors.Register(ModuleName, 29, "topic ID list value decode invalid length")
	ErrTopicIdListValueDecodeJsonInvalidLength           = errors.Register(ModuleName, 30, "topic ID list value decode JSON invalid length")
	ErrTopicIdListValueDecodeJsonInvalidFormat           = errors.Register(ModuleName, 31, "topic ID list value decode JSON invalid format")
	ErrTopicDoesNotExist                                 = errors.Register(ModuleName, 32, "topic does not exist")
	ErrCannotRemoveMoreStakeThanStakedInTopic            = errors.Register(ModuleName, 33, "cannot remove more stake than staked in topic")
	ErrInferenceRequestAlreadyInMempool                  = errors.Register(ModuleName, 34, "inference request already in mempool")
	ErrInferenceRequestBidAmountLessThanPrice            = errors.Register(ModuleName, 35, "inference request bid amount less than price")
	ErrInferenceRequestTimestampValidUntilInPast         = errors.Register(ModuleName, 36, "inference request timestamp valid until in past")
	ErrInferenceRequestTimestampValidUntilTooFarInFuture = errors.Register(ModuleName, 37, "inference request timestamp valid until too far in future")
	ErrInferenceRequestCadenceTooFast                    = errors.Register(ModuleName, 38, "inference request cadence too fast")
	ErrInferenceRequestCadenceTooSlow                    = errors.Register(ModuleName, 39, "inference request cadence too slow")
	ErrInferenceRequestWillNeverBeScheduled              = errors.Register(ModuleName, 40, "inference request will never be scheduled")
	ErrOwnerCannotBeEmpty                                = errors.Register(ModuleName, 41, "owner cannot be empty")
	ErrInsufficientStakeAfterRemoval                     = errors.Register(ModuleName, 42, "insufficient stake after removal")
	ErrInferenceRequestBidAmountTooLow                   = errors.Register(ModuleName, 43, "inference request bid amount too low")
	ErrIntegerUnderflowUnmetDemand                       = errors.Register(ModuleName, 44, "integer underflow for unmet demand")
	ErrInferenceCadenceBelowMinimum                      = errors.Register(ModuleName, 45, "inference cadence must be at least 60 seconds (1 minute)")
	ErrLossCadenceBelowMinimum                           = errors.Register(ModuleName, 46, "loss cadence must be at least 10800 seconds (3 hours)")
	ErrNotWhitelistAdmin                                 = errors.Register(ModuleName, 47, "not whitelist admin")
	ErrNotInTopicCreationWhitelist                       = errors.Register(ModuleName, 48, "not in topic creation whitelist")
	ErrNotInReputerWhitelist                             = errors.Register(ModuleName, 49, "not in reputer whitelist")
	ErrTopicNotEnoughDemand                              = errors.Register(ModuleName, 50, "topic not enough demand")
	ErrInvalidRequestId                                  = errors.Register(ModuleName, 51, "invalid request ID")
	ErrInferenceRequestNotInMempool                      = errors.Register(ModuleName, 52, "inference request not in mempool")
	ErrIntegerUnderflowStakeFromDelegator                = errors.Register(ModuleName, 53, "integer underflow for stake from delegator")
	ErrIntegerUnderflowDelegatedStakePlacement           = errors.Register(ModuleName, 54, "integer underflow for delegated stake placement")
	ErrIntegerUnderflowDelegatedStakeUponReputer         = errors.Register(ModuleName, 55, "integer underflow for delegated stake upon reputer")
	ErrEToTheXExponentiationIsInfinity                   = errors.Register(ModuleName, 56, "exponentiation overflow")
	ErrNaturalLogarithmIsInfinity                        = errors.Register(ModuleName, 57, "natural logarithm overflow")
	ErrLnToThePExponentiationIsInfinity                  = errors.Register(ModuleName, 58, "exponentiation overflow")
	ErrPhiInvalidInput                                   = errors.Register(ModuleName, 59, "phi: invalid input")
	ErrPhiResultIsNaN                                    = errors.Register(ModuleName, 60, "phi: result is NaN")
	ErrAdjustedStakeInvalidSliceLength                   = errors.Register(ModuleName, 61, "adjusted stake: invalid slice length")
	ErrAdjustedStakeIsInfinity                           = errors.Register(ModuleName, 62, "adjusted stake: is infinity")
	ErrAdjustedStakeIsNaN                                = errors.Register(ModuleName, 63, "adjusted stake: is NaN")
	ErrExponentialMovingAverageIsNaN                     = errors.Register(ModuleName, 64, "exponential moving average: is NaN")
	ErrExponentialMovingAverageIsInfinity                = errors.Register(ModuleName, 65, "exponential moving average: is infinity")
	ErrExponentialMovingAverageInvalidInput              = errors.Register(ModuleName, 66, "exponential moving average: invalid input")
	ErrFractionDivideByZero                              = errors.Register(ModuleName, 67, "fraction: divide by zero")
	ErrFractionInvalidInput                              = errors.Register(ModuleName, 68, "fraction: invalid input")
	ErrFractionIsNaN                                     = errors.Register(ModuleName, 69, "fraction: is NaN")
	ErrFractionIsInfinity                                = errors.Register(ModuleName, 70, "fraction: is infinity")
	ErrFractionInvalidSliceLength                        = errors.Register(ModuleName, 71, "fraction: invalid slice length")
	ErrEntropyInvalidInput                               = errors.Register(ModuleName, 72, "entropy: invalid input")
	ErrEntropyIsNaN                                      = errors.Register(ModuleName, 73, "entropy: is NaN")
	ErrEntropyIsInfinity                                 = errors.Register(ModuleName, 74, "entropy: is infinity")
	ErrNumberRatioDivideByZero                           = errors.Register(ModuleName, 75, "number ratio: divide by zero")
	ErrNumberRatioInvalidInput                           = errors.Register(ModuleName, 76, "number ratio: invalid input")
	ErrNumberRatioIsNaN                                  = errors.Register(ModuleName, 77, "number ratio: is NaN")
	ErrNumberRatioIsInfinity                             = errors.Register(ModuleName, 78, "number ratio: is infinity")
	ErrNumberRatioInvalidSliceLength                     = errors.Register(ModuleName, 79, "number ratio: invalid slice length")
	ErrInferenceRewardsInvalidInput                      = errors.Register(ModuleName, 80, "inference rewards: invalid input")
	ErrInferenceRewardsIsNaN                             = errors.Register(ModuleName, 81, "inference rewards: is NaN")
	ErrInferenceRewardsIsInfinity                        = errors.Register(ModuleName, 82, "inference rewards: is infinity")
	ErrForecastingRewardsInvalidInput                    = errors.Register(ModuleName, 83, "forecasting rewards: invalid input")
	ErrForecastingRewardsIsNaN                           = errors.Register(ModuleName, 84, "forecasting rewards: is NaN")
	ErrForecastingRewardsIsInfinity                      = errors.Register(ModuleName, 85, "forecasting rewards: is infinity")
	ErrReputerRewardsInvalidInput                        = errors.Register(ModuleName, 86, "reputer rewards: invalid input")
	ErrReputerRewardsIsNaN                               = errors.Register(ModuleName, 87, "reputer rewards: is NaN")
	ErrReputerRewardsIsInfinity                          = errors.Register(ModuleName, 88, "reputer rewards: is infinity")
	ErrForecastingPerformanceScoreInvalidInput           = errors.Register(ModuleName, 89, "forecasting performance score: invalid input")
	ErrForecastingPerformanceScoreIsNaN                  = errors.Register(ModuleName, 90, "forecasting performance score: is NaN")
	ErrForecastingPerformanceScoreIsInfinity             = errors.Register(ModuleName, 91, "forecasting performance score: is infinity")
	ErrSigmoidInvalidInput                               = errors.Register(ModuleName, 92, "sigmoid: invalid input")
	ErrSigmoidIsNaN                                      = errors.Register(ModuleName, 93, "sigmoid: is NaN")
	ErrSigmoidIsInfinity                                 = errors.Register(ModuleName, 94, "sigmoid: is infinity")
	ErrForecastingUtilityInvalidInput                    = errors.Register(ModuleName, 95, "forecasting utility: invalid input")
	ErrForecastingUtilityIsNaN                           = errors.Register(ModuleName, 96, "forecasting utility: is NaN")
	ErrForecastingUtilityIsInfinity                      = errors.Register(ModuleName, 97, "forecasting utility: is infinity")
	ErrNormalizationFactorInvalidInput                   = errors.Register(ModuleName, 98, "normalization factor: invalid input")
	ErrNormalizationFactorIsNaN                          = errors.Register(ModuleName, 99, "normalization factor: is NaN")
	ErrNormalizationFactorIsInfinity                     = errors.Register(ModuleName, 100, "normalization factor: is infinity")
	ErrInvalidSliceLength                                = errors.Register(ModuleName, 101, "invalid slice length")
	ErrTopicCadenceBelowMinimum                          = errors.Register(ModuleName, 102, "topic cadence must be at least 60 seconds (1 minute)")
	ErrPhiCannotBeZero                                   = errors.Register(ModuleName, 103, "phi: cannot be zero")
	ErrInferenceRequestBlockValidUntilInPast             = errors.Register(ModuleName, 104, "inference request block valid until in past")
	ErrInferenceRequestBlockValidUntilTooFarInFuture     = errors.Register(ModuleName, 105, "inference request block valid until too far in future")
	ErrSumWeightsLessThanEta                             = errors.Register(ModuleName, 106, "sum weights less than eta")
	ErrSliceLengthMismatch                               = errors.Register(ModuleName, 107, "slice length mismatch")
)
