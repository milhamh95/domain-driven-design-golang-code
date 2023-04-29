package chapter2

import (
	"errors"
	"time"
)

// out internal model for campaign
type Campaign struct {
	ID      string
	Title   string
	Goal    string
	EndDate time.Time
}

/*
// published language from another team

	{
		"id":"4cdd4ba9-7c04-4a3d-ac52-71f37ba75d7f",
		"metadata":{
			"name":"some campaign",
			"category":"growth",
			"endDate":"2023-04-12"
		}
	}
*/
type MarketingCampaignModel struct {
	Id       string `json:"ID"`
	Metadata struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		EndDate  string `json:"endDate"`
	} `json:"metadata"`
}

// we translate MarketingCampaignModel
// into our campaign domain model
func (m *MarketingCampaignModel) ToCampaign() (*Campaign, error) {
	if m.Id == "" {
		return nil, errors.New("campaign ID cannot be empty")
	}

	formattedDate, err := time.Parse("2006-01-02", m.Metadata.EndDate)
	if err != nil {
		return nil, errors.New("endDate was not in a parsable format")
	}

	return &Campaign{
		ID:      m.Id,
		Title:   m.Metadata.Name,
		Goal:    m.Metadata.Category,
		EndDate: formattedDate,
	}, nil
}
