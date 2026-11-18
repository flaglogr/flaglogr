package model

import (
	"time"

	"github.com/TwiN/logr"
	"gorm.io/gorm"
)

func (c *DBConnector) GetAllProjects(uid uint) ([]Project, error) {
	return gorm.G[Project](c.DB).Where("user_id = ?", uid).Find(ctx)
}

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
	if req.GroupID != nil {
		project.GroupID = req.GroupID
	}

	if c.DB.Save(&project).Error != nil {
		return false
	}
	return true
}

func (c *DBConnector) CreateProject(req CreateProjectRequest, userid uint) (uint, bool) {
	project := Project{
		Name:      req.Name,
		UserID:    userid,
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
	_, err := c.GetProject(id, uid)
	if err != nil {
		logr.Error(err.Error())
		return false
	}

	if rows, err := gorm.G[Project](c.DB).Where("id = ?", id).Delete(ctx); err != nil || rows == 0 {
		return false
	}

	return true
}
