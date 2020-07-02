// +build testground

package miner

import (
	abi "github.com/filecoin-project/specs-actors/actors/abi"
)

// The period over which all a miner's active sectors will be challenged.
const WPoStProvingPeriod = abi.ChainEpoch(240) // proving period set to 46*5 epochs ~ 240sec. instead of 24 hours (assuming block time delay of 1sec.)

// The duration of a deadline's challenge window, the period before a deadline when the challenge is available.
const WPoStChallengeWindow = abi.ChainEpoch(5) // challenge window set to 10 epochs ~ 10sec. instead of 40 minutes (assuming block time delay of 1sec.)

// The number of non-overlapping PoSt deadlines in each proving period.
const WPoStPeriodDeadlines = uint64(WPoStProvingPeriod / WPoStChallengeWindow) // 36 periods (one period every 5 epochs)

// The maximum age of a fault before the sector is terminated.
const FaultMaxAge = WPoStProvingPeriod*3 - 1 // not 14 days, but 3 days

// Number of epochs between publishing the precommit and when the challenge for interactive PoRep is drawn
// used to ensure it is not predictable by miner.
const PreCommitChallengeDelay = abi.ChainEpoch(10)
