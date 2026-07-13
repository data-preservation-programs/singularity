package testutil

// calibnet constants shared by integration tests

// CalibnetRPC defaults to the self-hosted forest gateway; glif rate-limited CI.
var CalibnetRPC = envOr("SINGULARITY_TEST_CALIBNET_RPC", "https://static.187.115.235.167.clients.your-server.de/calibnet/rpc/v1")

const (
	CalibnetChainID          = 314159
	CalibnetDDOContract      = "0x889fD50196BE300D06dc4b8F0F17fdB0af587095"
	CalibnetPaymentsContract = "0x09a0fDc2723fAd1A7b8e3e00eE5DF73841df55a0"
	CalibnetUSDFC            = "0xb3042734b608a1B16e9e86B374A3f3e389B4cDf0"

	// CalibnetDDOProviderActorID is the FF calibnet miner t0178773.
	CalibnetDDOProviderActorID uint64 = 178773
)
