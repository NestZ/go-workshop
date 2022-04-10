package models

type CampaignReq struct {
	Subject      string `bson: "subject"`
	BodyTemplate string `bson: "bodyTemplate`
}
