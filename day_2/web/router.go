package web

import (
	"context"

	"git.wndv.co/go/srv/errorx"
	"git.wndv.co/go/srv/gin"
	"git.wndv.co/thaneat.s/app2/models"
	"git.wndv.co/thaneat.s/app2/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Init(repositories *repos.Repos) *gin.Gin {
	r := gin.Default()
	r.Configurer().ProductionReady()

	r.GETw("/ping", func(ctx *gin.Context) (interface{}, error) {
		return gin.H{
			"message": "pong",
		}, nil
	})

	r.POSTw("/emails/register", func(ctx *gin.Context) (interface{}, error) {
		var emailReq models.EmailReq
		err := ctx.ShouldBindJSON(&emailReq)
		if err != nil {
			return nil, errorx.ErrBadRequest.Msg(err.Error())
		}

		err = repositories.Email.InsertEmail(context.TODO(), models.Email{
			ID:        primitive.NewObjectID(),
			Email:     emailReq.Email,
			FirstName: emailReq.FirstName,
			LastName:  emailReq.LastName,
		})
		if err != nil {
			return nil, errorx.ErrBadRequest.Msg(err.Error())
		}

		return gin.H{
			"message": "Successfully registered an email.",
		}, nil
	})

	r.POSTw("campaigns/prepare", func(ctx *gin.Context) (interface{}, error) {
		var campaignReq models.CampaignReq
		err := ctx.ShouldBindJSON(&campaignReq)
		if err != nil {
			return nil, errorx.ErrBadRequest.Msg(err.Error())
		}

		err = repositories.Campaign.InsertCampaign(context.TODO(), models.Campaign{
			ID:           primitive.NewObjectID(),
			Subject:      campaignReq.Subject,
			BodyTemplate: campaignReq.BodyTemplate,
		})
		if err != nil {
			return nil, errorx.ErrBadRequest.Msg(err.Error())
		}

		return gin.H{
			"message": "Successfully created a campaign.",
		}, nil
	})

	return r
}
