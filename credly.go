package credly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CredlyBadges struct {
	Data []*CredlyBadge `json:"data"`
}

type CredlyBadge struct {
	ID            string              `json:"id"`
	Info          CredlyBadgeTemplate `json:"badge_template"`
	CreatedAt     string              `json:"created_at"`
	ExpiresAt     string              `json:"expires_at"`
	ExpiresAtDate *CustomTime         `json:"expires_at_date"`
	IssuedAt      string              `json:"issued_at"`
	IssuedAtDate  *CustomTime         `json:"issued_at_date"`
}

func (c *CredlyBadge) GetName() string {
	return c.Info.Name
}

func (c *CredlyBadge) GetDescription() string {
	return c.Info.Description
}

func (c *CredlyBadge) GetExpiredDate() *CustomTime {
	return c.ExpiresAtDate
}

func (c *CredlyBadge) GetIssueDate() *CustomTime {
	return c.IssuedAtDate
}

func (c *CredlyBadge) GetImageUrl() string {
	return c.Info.ImageUrl
}

func (c *CredlyBadge) IsExpired() bool {
	if c.ExpiresAtDate == nil {
		return false
	}
	return c.ExpiresAtDate.Before(time.Now())
}

type CredlyBadgeIssuer struct {
	Summary string `json:"summary"`
}

type CredlyBadgeTemplate struct {
	ID                string            `json:"id"`
	Description       string            `json:"description"`
	GlobalActivityUrl string            `json:"global_activity_url"`
	Level             string            `json:"level"`
	Name              string            `json:"name"`
	State             string            `json:"state"`
	Image             *CredlyBadgeImage `json:"image"`
	ImageUrl          string            `json:"image_url"`
	Url               string            `json:"url"`
}

type CredlyBadgeImage struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}

type CredlyBadgeOwner struct {
	Name      string
	VanityUrl string
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	c.Time = date
	return nil
}

type CredlyService struct{}

func (c *CredlyService) GetBadges(username string) (*CredlyBadges, error) {
	url := fmt.Sprintf("https://www.credly.com/users/%s/badges.json", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	badges := CredlyBadges{}
	if err := json.NewDecoder(resp.Body).Decode(&badges); err != nil {
		return nil, err
	}
	return &badges, err
}
