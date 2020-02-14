package controllers

import (
	"github.com/beaconsoftwarellc/gadget/errors"
	"github.com/beaconsoftwarellc/quimby/example/models"
	qhttp "github.com/beaconsoftwarellc/quimby/http"
)

func (controller *widgetsController) doGet(context *qhttp.Context, username string, password string) (*models.WidgetCollection, errors.TracerError) {
	resp := &models.WidgetCollection{
		Items: controller.Specification.Storage.List(),
	}
	return resp, nil
}

func (controller *widgetsController) doPost(context *qhttp.Context, username string, password string, request *models.WidgetRequest) (*models.Widget, errors.TracerError) {
	err := request.Valid()
	if nil != err {
		return nil, err
	}
	widget := controller.Specification.Storage.Create(request)
	return widget, nil
}

func (controller *widgetController) doGet(context *qhttp.Context, username string, password string, widgetID string) (*models.Widget, errors.TracerError) {
	widget, err := controller.Specification.Storage.Get(widgetID)
	if nil != err {
		return nil, err
	}
	return widget, nil
}

func (controller *widgetController) doPut(context *qhttp.Context, username string, password string, widgetID string, request *models.WidgetRequest) (*models.Widget, errors.TracerError) {
	err := request.Valid()
	if nil != err {
		return nil, err
	}
	widget, err := controller.Specification.Storage.Get(widgetID)
	if nil != err {
		return nil, err
	}
	widget.Description = request.Description
	widget.SerialNumber = request.SerialNumber
	err = controller.Specification.Storage.Update(widget)
	if nil != err {
		return nil, err
	}
	return widget, nil
}

func (controller *widgetController) doPatch(context *qhttp.Context, username string, password string, widgetID string, request *models.WidgetPatch) (*models.Widget, errors.TracerError) {
	err := request.Valid()
	if nil != err {
		return nil, err
	}
	widget, err := controller.Specification.Storage.Get(widgetID)
	if nil != err {
		return nil, err
	}
	widget.Description = request.Description
	err = controller.Specification.Storage.Update(widget)
	if nil != err {
		return nil, err
	}
	return widget, nil
}

func (controller *widgetController) doDelete(context *qhttp.Context, username string, password string, widgetID string) (*models.Widget, errors.TracerError) {
	widget, err := controller.Specification.Storage.Get(widgetID)
	if nil != err {
		return nil, err
	}
	controller.Specification.Storage.Delete(widget.ID)
	return widget, nil
}
