package json

// TranslateJSONReq is request struct for http request with json
type TranslateJSONReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

// TranslateJSONResp is response struct for http request with json
type TranslateJSONResp struct {
	Success       bool          `json:"success"`
	TranslateData TranslateData `json:"data"`
}

// TranslateData will hold translation text include with error message
type TranslateData struct {
	Text       string `json:"text"`
	ErrMessage string `json:"error_message"`
}
