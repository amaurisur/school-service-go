package controllers

// ErrorController operations for Error
type ErrorController struct {
	Controller
}

// Error404 Path not found
func (e *ErrorController) Error404() {
	e.ServeResponse(e.ComposeResponseError(404, "err_not_found", "Path not found"))
}

// Error500 Internal server error
func (e *ErrorController) Error500() {
	e.ServeResponse(e.ComposeResponseError(500, "err_internal_server_error", "Internal server error"))
}
