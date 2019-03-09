package controllers

import (
	"net/http"

	"github.com/school/viewmodels"
)

// HealthController operations for Health
type HealthController struct {
	Controller
}

// Health check
// @Title Health
// @Description Check service health
// @Success 200 {object} viewmodels.Health
// @router /health [get]
func (h *HealthController) Health() {
	body := []viewmodels.Health{
		{
			Name:      "Events Service",
			Connected: true,
		},
	}
	h.Ctx.Output.SetStatus(http.StatusOK)
	h.Data["json"] = h.ComposeResponse(http.StatusOK, body)
	h.ServeJSON()
}
