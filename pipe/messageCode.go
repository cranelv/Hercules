package pipe
const (
	//Running space messages
	RunningSpace_Create = iota + 0x1000
	RunningSpace_TX_Call
	RunningSpace_TX_Commit
	RunningSpace_TX_LocalCall

	//txpool messages
	TXPOOL_AppendTx = iota + 0x2000
	TXPOOL_GenerateBlock

	//block headerMessage
	HEADER_GenerateHeader = iota + 0x3000
	//gateway message
	GATEWAY_BroadcastMessage = iota + 0x4000
)
