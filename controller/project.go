package controller

import (
	"log"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/gin-gonic/gin"
)

// GetAdminSellers gin handler
//
// @Summary		Get all sellers
// @Description	Retrieve a list of all currently registered sellers to modify
// @Tags		projects
// @Produce		json
// @Success		200	{object}		nil "TODO: Update"
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		500	{object}		nil
// @Router		/api/projects/get [get]
func (m *Controller) GetProjects(c *gin.Context) {
}

// GetAdminAllUsers gin handler
//
// @Summary		Get all users of this instance
// @Description	Retrieve a list of all currently registered users to create a new seller
// @Tags		admin
// @Produce		json
// @Accept      json
// @Param		payload body string true "TODO: Define"
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/api/projects/create [put]
func (m *Controller) CreateProject(c *gin.Context) {
	c.JSON(200, nil)
}

type UpdateProjectRequest struct {
	UserID   uint   `json:"uid" example:"1"`
	Website  string `json:"website" example:"https://janishutz.com"`
	PublicID string `json:"publicId" example:"com.janishutz"`
}

// CreateSeller gin handler
//
// @Summary		Create a new seller
// @Description	Create a new seller from a given user ID and other parametes
// @Tags		admin
// @Produce		json
// @Accept      json
// @Param		payload body UpdateProjectRequest true "Used to update or create the seller"
// @Success		200	{object}		nil "Seller created successfully"
// @Failure		400	{object}		nil
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		409	{object}		nil "Seller already exists"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/api/projects/update/{id} [post]
func (m *Controller) UpdateProject(c *gin.Context) {
	body := UpdateProjectRequest{}
	if err := c.BindJSON(&body); err != nil {
		log.Print("Error for JSON binding, ", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, nil)
}

// DeleteSeller gin handler
//
// @Summary		Delete given seller
// @Description	Deletes the specified seller permanently, no questions asked, including products
// @Tags		admin
// @Produce		json
// @Param		id path string true "The SellerID as supplied by the sellers endpoint"
// @Success		200	{object}		nil
// @Failure		400	{object}		nil
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		404	{object}		nil "User was not found"
// @Failure		500	{object}		httputil.HTTPError
// @Router		/api/projects/delete/{id} [delete]
func (m *Controller) DeleteProject(c *gin.Context) {
	idRaw, found := c.Params.Get("id")
	if !found {
		c.AbortWithStatus(400)
		return
	}
	id, err := strconv.ParseUint(idRaw, 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	logr.Debug(string(id))
	c.JSON(200, nil)
}
