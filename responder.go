package gallop

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"sync"
)

var responderList []Responder
var once_resp_list sync.Once

func get_responder_list() []Responder {
	once_resp_list.Do(func() {
		responderList = []Responder{(StringResponder)(nil),
			(JsonResponder)(nil),
			(XMLResponder)(nil),
		}
	})
	return responderList
}
func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range get_responder_list() {
		r_ref := reflect.TypeOf(r)
		if h_ref.Type().ConvertibleTo(r_ref) {
			return h_ref.Convert(r_ref).Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

type StringResponder func(*Context) string

func (s StringResponder) RespondTo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, s(&Context{c}))
	}
}

type Json interface{}
type JsonResponder func(*Context) Json

func (j JsonResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, j(&Context{context}))
	}
}

type XML interface {}

type XMLResponder func(*Context) XML

func (s XMLResponder) RespondTo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.XML(200,s(&Context{c}))
	}
}