package models

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	EventId int `json:"event_id" validate:"required"`
	UserId  int `json:"user_id" validate:"required"`
}

func (s *Subscription) Create() bool {
	var validate = validator.New()
	if err := validate.Struct(s); err != nil {
		return false
	}

	err := GetDB().Create(s).Error

	return err == nil
}

func DeleteByEventIdAndUserId(eventId, userId int) bool {
	s := Subscription{}
	err := GetDB().Where("event_id = ? AND user_id = ?", eventId, userId).Delete(&s).Error
	return err == nil
}

func (s Subscription) String() string {
	return fmt.Sprintf("EventID: %d, UserID: %d", s.EventId, s.UserId)
}
