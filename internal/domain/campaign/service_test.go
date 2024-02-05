package campaign

import (
	"errors"
	"testing"

	"github.com/evandersondev/emailn/internal/dtos"
	internalerrors "github.com/evandersondev/emailn/internal/internal_errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

var (
	newCampignDto = dtos.NewCampaignDto{
		Name:    "Test Y",
		Content: "Body Content",
		Emails:  []string{"email1@mail.com", "email2@mail.com"},
	}
	service = Service{}
)

func Test_Create_Campaign_Domain_Error(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(dtos.NewCampaignDto{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Save_Campaign(t *testing.T) {
	repository := new(repositoryMock)
	service.Repository = repository
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampignDto.Name ||
			campaign.Content != newCampignDto.Content ||
			len(campaign.Contacts) != len(newCampignDto.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampignDto)

	repository.AssertExpectations(t)
}

func Test_Error_Campaign(t *testing.T) {
	assert := assert.New(t)

	repository := new(repositoryMock)
	service.Repository = repository
	repository.On("Save", mock.Anything).Return(internalerrors.ErrInternal)

	_, err := service.Create(newCampignDto)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
