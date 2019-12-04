package Json

// TranslateJsonReq is request struct for http request with json
type TranslateJsonReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

// TranslateJsonResp is response struct for http request with json
type TranslateJsonResp struct {
	Success       bool          `json:"success"`
	TranslateData TranslateData `json:"data"`
}

// TranslateData will hold transaltion text include with error message
type TranslateData struct {
	Text       string `json:"text"`
	ErrMessage string `json:"error_message"`
}
