package apiutils

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type IDefaultHttpHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type DefaultHttpHandler struct {
	IHttpBodyUnmarshaler
	IStructValidator
	IDefaultService
	IRespWriter
}

func (dh DefaultHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	resp := dh.Process(r.Context(), nil)
	dh.Write(w, 200, map[string]interface{}{"message": resp})
}

func GetDefaultHandler() DefaultHttpHandler {
	return DefaultHttpHandler{
		IHttpBodyUnmarshaler: &JsonBodyUnmarshaler{},
		IStructValidator:     &ReqStructValidator{},
		IDefaultService:      &DefaultService{},
		IRespWriter:          &JsonRespWriter{},
	}
}

type IHttpBodyUnmarshaler interface {
	Unmarshal(r *http.Request, req interface{}) (err error)
}

type JsonBodyUnmarshaler struct {
}

func (rf *JsonBodyUnmarshaler) Unmarshal(r *http.Request, req interface{}) (err error) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyByte, req)
	if err != nil {
		return
	}
	return
}

type IStructValidator interface {
	ValidateStruct(req interface{}) (err error)
}

type ReqStructValidator struct {
}

func (rsv *ReqStructValidator) ValidateStruct(req interface{}) (err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return
	}
	return
}

type IRespWriter interface {
	Write(w http.ResponseWriter, httpCode int, resp interface{})
}

type TextRespWriter struct {
}

func (rw *TextRespWriter) Write(w http.ResponseWriter, httpCode int, resp interface{}) {
	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(httpCode)
	rByte := []byte(resp.(string))
	w.Write(rByte)
}

type JsonRespWriter struct {
}

func (rw *JsonRespWriter) Write(w http.ResponseWriter, httpCode int, resp interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpCode)
	rByte, _ := json.Marshal(&resp)
	w.Write(rByte)
}

type XmlRespWriter struct {
}

func (rw *XmlRespWriter) Write(w http.ResponseWriter, httpCode int, resp interface{}) {
	w.Header().Add("Content-type", "application/xml")
	w.WriteHeader(httpCode)
	rByte, _ := xml.Marshal(&resp)
	w.Write(rByte)
}
