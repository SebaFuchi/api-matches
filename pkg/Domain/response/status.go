package response

type Status int

const (
	InternalServerError Status = iota
	DBQueryError
	DBExecutionError
	DBRowsError
	DBLastRowIdError
	DBScanError

	SuccesfulCreation
	SuccesfulSearch
	CreationFailure
	MatchExists

	NotFound
	BadRequest
	Unknown
)

func (s Status) String() string {
	return [...]string{
		"InternalServerError",
		"DBQueryError",
		"DBExecutionError",
		"DBRowsError",
		"DBLastRowIdError",
		"DBScanError",

		"SuccesfulCreation",
		"SuccesfulSearch",
		"CreationFailure",
		"MatchExists",

		"NotFound",
		"BadRequest",
		"Unknown",
	}[s]
}

func (s Status) Index() int {
	return int(s)
}
