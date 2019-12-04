package Handler

import (
	"encoding/json"
	"github.com/Hiroshii8/GoTrans/Entity/Json"
	translate "github.com/Hiroshii8/GoTrans/Translate"
	"net/http"
)

type Handler struct{}

func (h *Handler) TranslateHandler(w http.ResponseWriter, r *http.Request) {

	var (
		isSuccess  = true
		errMessage string
	)
	request := Json.TranslateJsonReq{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	trans := translate.New()

	// call Request method for translate the text
	transResp, err := trans.Request(request.From, request.To, request.Text)
	if err != nil {
		isSuccess = false
		errMessage = err.Error()
	}

	// prepare for response
	resp := Json.TranslateJsonResp{
		Success: isSuccess,
		TranslateData: Json.TranslateData{
			Text:       transResp,
			ErrMessage: errMessage,
		},
	}
	res, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}
