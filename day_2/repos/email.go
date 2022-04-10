package repos

import (
	"context"

	"git.wndv.co/thaneat.s/app2/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRepository struct {
	col *mongo.Collection
}

func (e *EmailRepository) InsertEmail(ctx context.Context, email models.Email) error {
	_, err := e.col.InsertOne(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
