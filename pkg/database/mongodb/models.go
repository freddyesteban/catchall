package mongodb

import (
	"github.com/go-playground/validator"
)

type event string

const (
	TypeBounced   event = "bounced"
	TypeDelivered event = "delivered"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Event struct {
	Type   event  `json:"type" bson:"type" validate:"required,alpha"`
	Domain string `json:"domain" bson:"domain" validate:"required,alpha"`
}

func (e *Event) Validate() error {
	if err := validate.Struct(e); err != nil {
		// Check if there wasn't an error with the validation process
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}
