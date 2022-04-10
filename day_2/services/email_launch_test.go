package services_test

import (
	"fmt"
	"testing"

	"git.wndv.co/thaneat.s/app2/models"
	"git.wndv.co/thaneat.s/app2/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFormatEmail(t *testing.T) {
	type testCase struct {
		email        models.Email
		bodyTemplate string
		expectedBody string
	}

	cases := []testCase{
		{
			email: models.Email{
				Email:     "test@example.com",
				FirstName: "Thaneat",
				LastName:  "Saithong",
			},
			bodyTemplate: "Test {{firstName}}",
			expectedBody: "Test Thaneat",
		},
		{
			email: models.Email{
				Email:     "test2@example.com",
				FirstName: "Thaneat",
				LastName:  "Saithong",
			},
			bodyTemplate: "Test {{lastName}}",
			expectedBody: "Test Saithong",
		},
		{
			email: models.Email{
				Email:     "test3@example.com",
				FirstName: "Thaneat",
				LastName:  "Saithong",
			},
			bodyTemplate: "Test {{firstName}} {{lastName}}",
			expectedBody: "Test Thaneat Saithong",
		},
	}

	for i, c := range cases {
		t.Run("Test format email # "+fmt.Sprint(i+1), func(t *testing.T) {
			campaign := models.Campaign{
				Subject:      "Test campaign",
				BodyTemplate: c.bodyTemplate,
			}

			output := services.FormatEmail(c.email, &campaign)

			if output.BodyTemplate != c.expectedBody {
				t.Errorf("Body is not matched: expected=%v got=%v", c.expectedBody, output.BodyTemplate)
			}
		})
	}
}

func BenchmarkFormatEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.FormatEmail(models.Email{
			Email:     "test@example.com",
			FirstName: "Thaneat",
			LastName:  "Saithong",
		}, &(models.Campaign{
			ID:           primitive.NewObjectID(),
			Subject:      "test",
			BodyTemplate: "{{firstName}} {{lastName}}",
		}))
	}
}
