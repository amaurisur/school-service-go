package controllers

import (
	"github.com/astaxie/beego"
	"github.com/school/utils"
	"github.com/school/viewmodels"
)

// Controller _
type Controller struct {
	beego.Controller
}

func (c *Controller) composeMeta(statusCode int) viewmodels.Meta {
	return viewmodels.Meta{Code: statusCode, Version: utils.GetAppConfig("version", "")}
}

// ComposeResponse returns a data structure to be rendered
func (c *Controller) ComposeResponse(statusCode int, body interface{}) viewmodels.Response {
	return viewmodels.Response{Meta: c.composeMeta(statusCode), Data: body}
}

// ComposeResponseError sends a error structure to be rendered
func (c *Controller) ComposeResponseError(statusCode int, descriptionCode string, description string) viewmodels.Response {
	var errors []viewmodels.Error
	errors = append(errors, viewmodels.Error{Description: description, Code: descriptionCode})
	return viewmodels.Response{Meta: c.composeMeta(statusCode), Errors: errors}
}

// ServeResponse _
func (c *Controller) ServeResponse(response viewmodels.Response) {
	c.Ctx.Output.SetStatus(response.Meta.Code)
	c.Data["json"] = response
	c.ServeJSON()
}

// ComposeResponseErrorCollection sends an errors list structure to be rendered
func (c *Controller) ComposeResponseErrorCollection(statusCode int, errors []viewmodels.Error) viewmodels.Response {
	return viewmodels.Response{Meta: c.composeMeta(statusCode), Errors: errors}
}
