package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adelapazborrero/gophisher/service/server"
	"github.com/adelapazborrero/gophisher/types"
)

type CampaignHandler struct {
	campaignService server.CampaignService
}

func NewCampaignHandler(campaignService server.CampaignService) *CampaignHandler {
	return &CampaignHandler{campaignService: campaignService}
}

func (ch *CampaignHandler) GetAllCampaigns(c *gin.Context) {
	campaigns, err := ch.campaignService.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not fetch campaigns"})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

func (ch *CampaignHandler) GetCampaignByID(c *gin.Context) {
	idStr := c.Param("id")

	campaign, err := ch.campaignService.Get(types.ID(idStr))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	c.JSON(http.StatusOK, campaign)
}
