package Handler

import (
	"encoding/json"
	"net/http"

	jsonEntity "github.com/Hiroshii8/GoTrans/entity/json"
	translate "github.com/Hiroshii8/GoTrans/translate"
)

type Handler struct{}

func (h *Handler) TranslateHandler(w http.ResponseWriter, r *http.Request) {

	var (
		isSuccess  = true
		errMessage string
	)
	request := jsonEntity.TranslateJSONReq{}

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
	resp := jsonEntity.TranslateJSONResp{
		Success: isSuccess,
		TranslateData: jsonEntity.TranslateData{
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
