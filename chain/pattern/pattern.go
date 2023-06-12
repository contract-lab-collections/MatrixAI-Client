package pattern

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

// DOT is "." character
const DOT = "."

// Pallets
const (
	// OSS is a module about DeOSS
	OSS = "Oss"

	// HASHRATE_MARKET is a module about DeOSS
	HASHRATE_MARKET = "HashrateMarket"

	// SYSTEM is a module about the system
	SYSTEM = "System"
)

// Chain state
const (
	// SYSTEM
	ACCOUNT = "Account"
	EVENTS  = "Events"
)

// Extrinsic
const (
	// OSS
	TX_OSS_REGISTER = OSS + DOT + "authorize"

	// TX_HASHRATE_MARKET_REGISTER
	TX_HASHRATE_MARKET_REGISTER = HASHRATE_MARKET + DOT + "add_machine"
)

type FileHash [64]types.U8

type MachineUUID [16]types.U8

const (
	ERR_Failed  = "failed"
	ERR_Timeout = "timeout"
	ERR_Empty   = "empty"
)
