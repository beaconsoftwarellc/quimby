package models

// THIS IS A GENERATED FILE. DO NOT MODIFY
// api_models.tmpl

import (
	"time"
)

// Widget represent simple Widget for the example
type Widget struct {
	ID           string    `json:"id"`
	SerialNumber string    `json:"serial_number"`
	Description  string    `json:"description"`
	CreatedOn    time.Time `json:"created_on"`
	UpdatedOn    time.Time `json:"updated_on"`
}

// WidgetRequest is the allowed input for a Widget (POST, PUT)
type WidgetRequest struct {
	SerialNumber string `json:"serial_number"`
	Description  string `json:"description"`
}

// WidgetPatch is the allowed input for a Widget (PATCH)
type WidgetPatch struct {
	Description string `json:"description"`
}

// Collection is the model for returning a Collection of objects in the API
type Collection struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	URI      string `json:"uri"`
}

// WidgetCollection is a paginated collection of Widget models
type WidgetCollection struct {
	Collection
	Items []*Widget `json:"items"`
}
