package formbind

import (
	"dataprocess_service/controllers/formbind/json"
	"net/http"
	"strings"
)

func Decode(req *http.Request, body []byte, receive interface{}) error {
	if strings.Contains(req.Header.Get("Content-Type"), "json") && len(body) > 0 {
		return json.Unmarshal(body, receive)
	}
	if req.MultipartForm != nil || req.Form != nil {
		return DecodeForm(req, receive)
	}
	return nil
}
