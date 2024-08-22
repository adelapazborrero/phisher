package routes

import (
	"github.com/adelapazborrero/gophisher/config"
	"github.com/adelapazborrero/gophisher/handlers"
	"github.com/adelapazborrero/gophisher/service/server"
)

const (
	messageCurrentVersion = "/api/v1"
	apiMessages           = messageCurrentVersion + "/campaigns"
)

func CampaignRoutes(campaignService server.CampaignService) []config.Route {
	campaignHandler := handlers.NewCampaignHandler(campaignService)
	return []config.Route{
		{
			Method:  "GET",
			Path:    apiMessages,
			Handler: campaignHandler.GetAllCampaigns,
		},
		{
			Method:  "GET",
			Path:    apiMessages + "/:id",
			Handler: campaignHandler.GetCampaignByID,
		},
	}
}
