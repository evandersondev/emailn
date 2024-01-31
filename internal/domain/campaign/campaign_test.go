package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "body content"
	emails  = []string{"email1@mail.com", "email1@mail.com"}
	fake    = faker.New()
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, emails)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}

func Test_NewCampaign_ID_Is_Not_Nil(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, emails)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_Date_Is_Greater_Than_Now(t *testing.T) {
	assert := assert.New(t)
	dateNow := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedAt, dateNow)
}

func Test_NewCampaign_Name_Validate_Min(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign("", content, emails)

	assert.Equal("name is required with min 5", error.Error())
}

func Test_NewCampaign_Name_Validate_Max(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(fake.Lorem().Text(25), content, emails)

	assert.Equal("name is required with max 24", error.Error())
}

func Test_NewCampaign_Content_Validate_Min(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, "", emails)

	assert.Equal("content is required with min 5", error.Error())
}

func Test_NewCampaign_Content_Validate_Max(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, fake.Lorem().Text(1040), emails)

	assert.Equal("content is required with max 1024", error.Error())
}

func Test_NewCampaign_Contacts_Validate_Email(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", error.Error())
}
