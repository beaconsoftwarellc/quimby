package controllers

import (
	"fmt"

	"gitlab.com/beacon-software/quimby/http"
)

// Get returns the ID and subresource (if provided)
func (controller *resourceController) Get(context *http.Context) {
	id, _ := context.URIParameters["id"]
	subresource, ok := context.URIParameters["subresource"]
	context.Write(fmt.Sprintf("ID: %s\n", id))
	if ok {
		context.Write(fmt.Sprintf("Subresource: %s", subresource))
	}
}
