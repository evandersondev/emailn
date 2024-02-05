package endpoint

import "github.com/evandersondev/emailn/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
