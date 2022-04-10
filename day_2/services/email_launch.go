package services

import (
	"strings"

	"git.wndv.co/thaneat.s/app2/models"
)

func FormatEmail(email models.Email, campaign *models.Campaign) models.Campaign {
	bodyTemplate := campaign.BodyTemplate
	bodyTemplate = strings.Replace(bodyTemplate, "{{firstName}}", email.FirstName, 1)
	bodyTemplate = strings.Replace(bodyTemplate, "{{lastName}}", email.LastName, 1)
	campaign.BodyTemplate = bodyTemplate
	return *campaign
}
