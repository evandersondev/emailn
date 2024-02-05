package endpoint

import (
	"net/http"

	"github.com/evandersondev/emailn/internal/dtos"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var campaign dtos.NewCampaignDto
	render.DecodeJSON(r.Body, &campaign)
	id, err := h.CampaignService.Create(campaign)

	return map[string]string{"id": id}, 301, err
}
