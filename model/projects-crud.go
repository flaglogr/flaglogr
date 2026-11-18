package model

import (
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetProject(id uint, uid uint) (Project, error) {
	return gorm.G[Project](c.DB).Where("id = ? AND user_id = ?", id, uid).First(ctx)
}

func (c *DBConnector) UpdateProject(id uint, req UpdateProjectRequest, uid uint) bool {
	project, err := c.GetProject(id, uid)
	if err != nil {
		return false
	}

	project.Name = dereferenceOrDefault(req.Name, project.Name)
	project.UpdatedAt = time.Now()

	if c.DB.Save(&project).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateProject(req CreateProjectRequest, userid uint) (uint, bool) {
	// TODO: Get user and associate it to that user
	project := Project{
		Name:      req.Name,
		UpdatedAt: time.Now(),
	}

	result := gorm.WithResult()
	err := gorm.G[Project](c.DB, result).Create(ctx, &project)
	if err != nil {
		return 0, false
	}
	return project.ID, true
}

func (c *DBConnector) DeleteProject(id uint, uid uint) bool {
	project, err := c.GetProject(id, uid)
	if err != nil {
		logr.Error(err.Error())
		return false
	}

	logr.Debugf("%s", project)

	c.DB.Exec("DELETE FROM `subtopics_projects` WHERE project_id = ?; DELETE FROM `subjects_assets` WHERE asset_id = ?", id, id)
	if rows, err := gorm.G[Project](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}

	return true
}
