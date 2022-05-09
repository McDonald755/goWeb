package config

var (
	// API root Group
	API = "/api"

	// V1 v1 Group
	V1           = "/v1"
	API_V1_GROUP = API + V1

	// XXX Group
	POOL = "/xxx"

	// COMMON common
	COMMON = "/common"

	SUCCESS        = 2000
	ERROR          = 5000
	INVALID_PARAMS = 4000
	UPLOADERROR    = 4001
	DBERROR        = 4002
	GETERROR       = 4003
)

var MsgFlags = map[int]string{
	SUCCESS:        "Success",
	ERROR:          "Fail",
	UPLOADERROR:    "Upload Picture Error",
	INVALID_PARAMS: "Request Parameter Error",
	DBERROR:        "Update Database Error",
	GETERROR:       "Get Price Error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
