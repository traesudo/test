package formbind

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/imiskolee/dateparse"
	"github.com/imiskolee/form"
	"github.com/pkg/errors"
)

func DecodeForm(r *http.Request, receive interface{}) (err error) {
	defer (func() {
		if e := recover(); e != nil {
			log.Println(e)
			debug.PrintStack()
			err = errors.New(fmt.Sprint(err))
		}
	})()
	rawQuerys := []string{}
	if r.MultipartForm != nil {
		r.Form = r.MultipartForm.Value
	}
	for k, v := range r.Form {
		for _, vv := range v {
			rawQuerys = append(rawQuerys, fmt.Sprintf("%s=%s", k, url.PathEscape(vv)))
		}
	}
	rawQuery := strings.Join(rawQuerys, "&")
	form := form.NewForm(rawQuery)
	form.NeedQueryUnescape(false)
	maps, err := form.Decode()
	if err != nil {
		return err
	}
	rv := reflect.ValueOf(receive)
	nv := parse(maps, rv.Type().Elem())
	rv.Elem().Set(nv)
	return nil
}

func getFieldName(field reflect.StructField) string {
	fieldName := field.Name
	if val := field.Tag.Get("json"); val != "" {
		fieldName = val
	}
	return fieldName
}

func isNull(data interface{}) bool {
	if v, ok := data.(string); ok && v == "null" || v == "(null)" {
		return true
	}
	return false
}

func parse(data interface{}, rt reflect.Type) reflect.Value {
	rv := reflect.New(rt)
	if rv.Kind() == reflect.Ptr {
		if rv.IsValid() && rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}
	if canSet(rv) {
		setValue(data, rv)
	} else {
		switch rv.Kind() {
		case reflect.Struct:
			ret := mapToStruct(data.(map[string]interface{}), rv)
			rv.Set(*ret)
			break
		case reflect.Slice:
			ret := sliceToSlice(data, rv.Type())
			rv.Set(*ret)
			break
		case reflect.Map:
			ret := mapToMap(data.(map[string]interface{}), rv.Type())
			rv.Set(*ret)
			break
		case reflect.Ptr:
			return parse(data, rt.Elem())

		default:
			setValue(data, rv)
		}
	}
	return rv
}

func mapToMap(data map[string]interface{}, typ reflect.Type) *reflect.Value {
	if typ.Kind() != reflect.Map {
		panic("data must be map")
	}
	zeroRet := reflect.MakeMap(typ)
	for k, v := range data {
		if isNull(v) {
			zeroRet.SetMapIndex(reflect.ValueOf(k), reflect.Zero(typ.Elem()))
		} else {
			child := parse(v, typ.Elem())
			zeroRet.SetMapIndex(reflect.ValueOf(k), child)
		}
	}
	return &zeroRet
}

func mapToStruct(data map[string]interface{}, rv reflect.Value) *reflect.Value {

	if rv.Kind() != reflect.Struct {
		panic("data must be struct")
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldName := getFieldName(rv.Type().Field(i))
		index := strings.Index(fieldName, ",")
		if index > 0 {
			fieldName = strings.TrimSpace(fieldName[:index])
		}
		ft := field.Type()
		var child reflect.Value
		hasValue := false
		if rv.Type().Field(i).Anonymous {
			child = parse(data, ft)
			hasValue = true
		} else {
			//Skip field null value
			if v, ok := data[fieldName]; ok && !isNull(v) {
				child = parse(v, ft)
				hasValue = true
			}
		}
		if hasValue {
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.New(field.Type().Elem()))
				field.Elem().Set(child)
			} else {
				field.Set(child)
			}
		}
	}
	return &rv
}

func sliceToSlice(data interface{}, typ reflect.Type) *reflect.Value {
	if typ.Kind() != reflect.Slice && typ.Kind() != reflect.Array {
		panic("data must be slice")
	}
	lst := data.([]interface{})
	zeroRet := reflect.MakeSlice(typ, len(lst), len(lst))
	for i, v := range lst {
		if !isNull(v) {
			child := parse(v, typ.Elem())
			if typ.Elem().Kind() == reflect.Ptr {
				zeroRet.Index(i).Set(reflect.New(typ.Elem().Elem()))
				zeroRet.Index(i).Elem().Set(child)
			} else {
				zeroRet.Index(i).Set(child)
			}
		}
	}
	return &zeroRet
}

func stringToTime(data string) (*reflect.Value, error) {
	return nil, nil
}

func canSet(field reflect.Value) bool {
	switch field.Interface().(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	case bool:
		return true
	case string:
		return true
	case time.Time:
		return true
	}
	addrField := field
	if field.CanAddr() {
		addrField = field.Addr()
	}
	unmarshalerFunc := addrField.MethodByName("UnmarshalText")
	if unmarshalerFunc.IsValid() {
		return true
	}
	return false
}

func setValue(data interface{}, field reflect.Value) {
	formVal := ""
	switch v := data.(type) {
	case string:
		formVal = v
		break
	case []byte:
		formVal = string(v)
	default:
		formVal = fmt.Sprint(formVal)
	}
	formVal, _ = url.PathUnescape(formVal)
	//TODO: hack null value
	if isNull(formVal) {
		return
	}
	switch field.Interface().(type) {
	case int, int8, int16, int32, int64:
		if formVal == "" {
			return
		}
		num, err := strconv.ParseInt(fmt.Sprint(data), 10, 64)
		if err != nil {
			panic("can't parse int value:" + formVal + " field:" + field.String())
		}
		field.SetInt(num)
		return
	case uint, uint8, uint16, uint32, uint64:
		if formVal == "" {
			return
		}
		num, err := strconv.ParseUint(formVal, 10, 64)
		if err != nil {
			panic("can't parse uint value:" + formVal + " field:" + field.String())
		}
		field.SetUint(num)
		return
	case float32, float64:
		if formVal == "" {
			return
		}
		num, err := strconv.ParseFloat(formVal, 64)
		if err != nil {
			panic("can't parse float value:" + formVal + " field:" + field.String())
		}
		field.SetFloat(num)
		return
	case bool:
		if formVal == "" {
			return
		}
		boolVal, err := strconv.ParseBool(formVal)
		if err != nil {
			panic("can't parse bool value:" + formVal + " field:" + field.String())
		}
		field.SetBool(boolVal)
		return
	case string:
		field.SetString(formVal)
		return
	case time.Time:
		if formVal == "" {
			return
		}
		t, err := dateparse.ParseAny(formVal)
		if err != nil {
			panic("can't parse time.Time:" + formVal + " field:" + field.String())
		}
		field.Set(reflect.ValueOf(t))
		return
	}

	addrField := field
	if field.CanAddr() {
		addrField = field.Addr()
	}
	unmarshalerFunc := addrField.MethodByName("UnmarshalText")
	if unmarshalerFunc.IsValid() {
		output := unmarshalerFunc.Call([]reflect.Value{reflect.ValueOf([]byte(formVal))})
		data := output[0].Interface()
		if data != nil {
			panic("can't parse value with UnmarshalText:" + formVal)
		}
		return
	}
	if field.Kind() == reflect.Interface {
		field.Set(reflect.ValueOf(data))
	}
}
