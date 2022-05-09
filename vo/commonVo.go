package vo

type PageVo struct {
	PageNum   int64 `json:"pageNum"`
	PageSize  int64 `json:"pageSize"`
	TotalSize int64 `json:"totalSize"`
}

type ResponsePageVo struct {
	PageNum   int64       `json:"pageNum"`
	PageSize  int64       `json:"pageSize"`
	TotalSize int64       `json:"totalSize"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type ResponseVo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MoralisVo struct {
	Total    int64  `json:"total"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	Cursor   string `json:"cursor"`
	Result   []struct {
		TokenAddress      string      `json:"token_address"`
		TokenId           string      `json:"token_id"`
		BlockNumberMinted string      `json:"block_number_minted"`
		OwnerOf           string      `json:"owner_of"`
		BlockNumber       string      `json:"block_number"`
		Amount            string      `json:"amount"`
		ContractType      string      `json:"contract_type"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		TokenURI          interface{} `json:"token_uri"`
		Metadata          interface{} `json:"metadata"`
		SyncedAt          interface{} `json:"synced_at"`
		IsValid           int         `json:"is_valid"`
		Syncing           int         `json:"syncing"`
		Frozen            int         `json:"frozen"`
	} `json:"result"`
	Status string `json:"status"`
}

type Price struct {
	Price string `json:"price"`
}

type BinancePrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
