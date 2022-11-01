package apihelper

//
//import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	log "github.com/sirupsen/logrus"
//	"net/http"
//	"net/url"
//	"reflect"
//	"runtime/debug"
//	"strings"
//)
//
//func Decode(req *http.Request, body []byte, receive interface{}) error {
//	if strings.Contains(req.Header.Get("Content-Type"), "json") && len(body) > 0 {
//		return json.Unmarshal(body, receive)
//	}
//	if req.MultipartForm != nil || req.Form != nil {
//		return DecodeForm(req, receive)
//	}
//	return nil
//}
//func DecodeForm(r *http.Request, receive interface{}) (err error) {
//	defer (func() {
//		if e := recover(); e != nil {
//			log.Println(e)
//			debug.PrintStack()
//			err = errors.New(fmt.Sprint(err))
//		}
//	})()
//	rawQuerys := []string{}
//	if r.MultipartForm != nil {
//		r.Form = r.MultipartForm.Value
//	}
//	for k, v := range r.Form {
//		for _, vv := range v {
//			rawQuerys = append(rawQuerys, fmt.Sprintf("%s=%s", k, url.PathEscape(vv)))
//		}
//	}
//	rawQuery := strings.Join(rawQuerys, "&")
//	form := form.NewForm(rawQuery)
//	form.NeedQueryUnescape(false)
//	maps, err := form.Decode()
//	if err != nil {
//		return err
//	}
//	rv := reflect.ValueOf(receive)
//	nv := parse(maps, rv.Type().Elem())
//	rv.Elem().Set(nv)
//	return nil
//}
