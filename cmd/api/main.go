package main

import (
	"net/http"

	"github.com/evandersondev/emailn/internal/domain/campaign"
	"github.com/evandersondev/emailn/internal/endpoint"
	"github.com/evandersondev/emailn/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignServices := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoint.Handler{
		CampaignService: campaignServices,
	}

	r.Get("/campaigns", endpoint.HandlerError(handler.CampaignGetAll))
	r.Post("/campaigns", endpoint.HandlerError(handler.CampaignPost))

	http.ListenAndServe(":3000", r)
}
