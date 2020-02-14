package controllers

import (
	"fmt"
	"net/http"

	"github.com/beaconsoftwarellc/gadget/stringutil"
	qerror "github.com/beaconsoftwarellc/quimby/error"
	qhttp "github.com/beaconsoftwarellc/quimby/http"
)

func (controller *echoController) Get(context *qhttp.Context) {
	r := context.Request
	context.Write(fmt.Sprintf("Host: %s\n", r.Host))
	context.Write(fmt.Sprintf("RequestURI: %s\n", r.RequestURI))
	context.Write(fmt.Sprintf("Method: %s\n", r.Method))
	context.Write(fmt.Sprintf("RemoteAddr: %s\n", r.RemoteAddr))
	context.Write(fmt.Sprintf("Content Length: %d\n", r.ContentLength))
	context.Write("Headers:\n")
	for k, v := range context.Request.Header {
		context.Write(fmt.Sprintf("\t%s: %s\n", k, v))
	}

	if !stringutil.IsWhiteSpace(context.URIParameters["toEcho"]) {
		context.Write(fmt.Sprintf("URI:\n%s\n", context.URIParameters["toEcho"]))
	}

	if context.Request.ContentLength > 0 {
		context.Write("Body:\n")
		body, err := context.Read()
		if err != nil {
			context.SetError(qerror.NewRestError(qerror.SystemError, "", nil), http.StatusInternalServerError)
			return
		}
		context.Response.Write(body)
	}

}

// Post writes the information from the request to the body of the response.
func (controller *echoController) Post(context *qhttp.Context) {
	controller.Get(context)
}

// Put writes the information from the request to the body of the response.
func (controller *echoController) Put(context *qhttp.Context) {
	controller.Get(context)
}

// Delete writes the information from the request to the body of the response.
func (controller *echoController) Delete(context *qhttp.Context) {
	controller.Get(context)
}
