package errorcode

const (
	// system
	None = iota
	SystemError

	// database - common
	DatabaseConnectFailed  = 200001
	DatabaseInsertFailed   = 200002
	DatabaseQueryFailed    = 200003
	DatabaseUpdateFailed   = 200004
	DatabaseDeleteFailed   = 200005
	DatabaseRecordNotFound = 200006
	DatabaseDuplicateKey   = 200007

	// database - mongo
	DatabaseConnectClientError = 300001
	DatabaseConnectPoolFailed  = 300002
	DatabaseCollectionFailed   = 300003
	DatabaseParseDataError     = 300004
)
