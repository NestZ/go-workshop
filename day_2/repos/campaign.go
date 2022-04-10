package repos

import (
	"context"

	"git.wndv.co/thaneat.s/app2/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignRepository struct {
	col *mongo.Collection
}

func (c *CampaignRepository) InsertCampaign(ctx context.Context, campaign models.Campaign) error {
	_, err := c.col.InsertOne(ctx, campaign)
	if err != nil {
		return err
	}
	return nil
}
