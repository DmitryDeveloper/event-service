package models

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	BaseModel
	Title            string     `json:"title" validate:"required"`
	ShortDescription string     `json:"short_description" validate:"required"`
	Description      string     `json:"description" validate:"required"`
	UserId           int        `json:"user_id" validate:"required"`
	IsApproved       bool       `json:"is_approved"`
	Categories       []Category `gorm:"many2many:event_categories;" json:"categories" validate:"required"`
}

func (event *Event) Create() bool {
	//validate fields, it uses tags in struct fields
	var validate = validator.New()
	if err := validate.Struct(event); err != nil {
		return false
	}

	err := GetDB().Create(event).Error
	return err == nil
}

func (m *Event) GetById(id int) error {
	return GetDB().First(m, id).Error
}

func (m *Event) GetAll(limit int) ([]Event, error) {
	var models []Event
	err := GetDB().Limit(limit).Find(&models).Error
	return models, err
}

func (m *Event) GetByCategoryId(categoryId int) ([]Event, error) {
	var models []Event
	err := GetDB().Joins("JOIN event_categories ON event_categories.event_id = events.id").
		Where("event_categories.category_id = ?", categoryId).
		Find(&models).Error
	return models, err
}
