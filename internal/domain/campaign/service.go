package campaign

import (
	"github.com/evandersondev/emailn/internal/dtos"
	internalerrors "github.com/evandersondev/emailn/internal/internal_errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign dtos.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}
