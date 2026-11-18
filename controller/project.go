package controller

import (
	"fmt"
	"strconv"

	"github.com/TwiN/logr"
	"github.com/flaglogr/flaglogr/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Gin handler to get all projects
//
// @Summary		get all projects
// @Description	Retrieve a list of all projects of all groups of the current user
// @Tags		projects
// @Produce		json
// @Success		200	{object}		[]model.Project
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		500	{object}		nil
// @Router		/api/v1/projects/get [get]
func (m *Controller) GetProjects(c *gin.Context) {
	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	project, err := m.db.GetAllProjects(uid)
	if err != nil {
		c.AbortWithStatus(404)
	}
	c.JSON(200, project)
}

// Gin handler to create a new project
//
// @Summary		Create a new project
// @Description	Create a new project for a given group
// @Tags		projects
// @Produce		json
// @Accept      json
// @Param		payload body model.CreateProjectRequest true "The project to create"
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		500	{object}		nil
// @Router		/api/v1/projects/create [put]
func (m *Controller) CreateProject(c *gin.Context) {
	var req model.CreateProjectRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to create asset, error: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	id, ok := m.db.CreateProject(req, uid)
	if ok {
		c.JSON(201, CreatedResponse{
			Id: id,
		})
	} else {
		c.AbortWithStatus(500)
	}
}

// Gin handler to update project
//
// @Summary		Update a project
// @Description	Update the specified project
// @Tags		projects
// @Produce		json
// @Accept      json
// @Param		payload body model.UpdateProjectRequest true "Details to update for the project"
// @Success		200	{object}		nil "Project updated successfully"
// @Failure		400	{object}		nil
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		500	{object}		nil
// @Router		/api/v1/projects/update/{id} [patch]
func (m *Controller) UpdateProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	var req model.UpdateProjectRequest
	err = c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to update project, error: ", err))
		c.Abort()
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	if m.db.UpdateProject(uint(id), req, uid) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}

// DeleteSeller gin handler
//
// @Summary		Delete given project
// @Description	Deletes the specified project permanently, no questions asked, including all flags, evaluations and results
// @Tags		projects
// @Produce		json
// @Param		id path string true "The project ID as supplied by the get endpoint"
// @Success		200	{object}		nil
// @Failure		400	{object}		nil
// @Failure		401	{object}		nil
// @Failure		403	{object}		nil
// @Failure		404	{object}		nil "Project was not found"
// @Failure		500	{object}		nil
// @Router		/api/v1/projects/delete/{id} [delete]
func (m *Controller) DeleteProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	if m.db.DeleteProject(uint(id), uid) {
		c.JSON(204, nil)
	} else {
		c.JSON(404, nil)
	}
}
