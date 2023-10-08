package models

import (
	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Category struct {
	BaseModel
    gorm.Model
    Name        string `json:"name" gorm:"unique_index" validate:"required"`
    Description string `json:"description" validate:"required"`
}

func (category *Category) Create() bool {
	//validate fields, it uses tags in struct fields
	var validate = validator.New()
	if err := validate.Struct(category); err != nil {
        return false
    }

	err := GetDB().Create(category).Error

	if err != nil {
		return false
	}

	return true
}

func (m *Category) GetById(id int) error {
	return GetDB().First(m, id).Error
}

func (m *Category) GetAll(limit int) ([]Category, error) {
	var models []Category
	err := GetDB().Limit(limit).Find(&models).Error
	return models, err
}
