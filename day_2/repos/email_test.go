package repos_test

import (
	"context"
	"testing"

	"git.wndv.co/thaneat.s/app2/models"
	"git.wndv.co/thaneat.s/app2/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestEmailRepository(t *testing.T) {
	emailRepo := repos.Init().Email
	err := emailRepo.InsertEmail(context.TODO(), models.Email{
		ID:        primitive.NewObjectID(),
		Email:     "test1@example.com",
		FirstName: "Thaneat",
		LastName:  "Saithong",
	})
	if err != nil {
		t.Error(err)
	}
}
