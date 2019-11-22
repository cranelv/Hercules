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

	//blockchain message
	BLOCKCHAIN_InsertBlock = iota + 0x5000
	BLOCKCHAIN_InsertHeader
	BLOCKCHAIN_InsertStateSet
	BLOCKCHAIN_DeleteBlock

	//database
	DATABASE_Insert = iota + 0x6000
	DATABASE_Update
	DATABASE_InsertUpdate
	DATABASE_Delete
)
