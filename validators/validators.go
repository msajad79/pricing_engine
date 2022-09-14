package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/msajad79/pricing_engine/config"
)

func ValidateRegister() {
	config.Val.RegisterValidation("is_valid_routes", IsValidRoutes)
}

func IsValidRoutes(fl validator.FieldLevel) bool {
	fl.Field().String()
	return false
}
