package config

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Val *validator.Validate
var Trans ut.Translator
