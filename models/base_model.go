package models

import (
	"mortred/utils"
	"os"

	"github.com/avelino/slugify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type BaseModel struct {
	ErrorsMap map[string]interface{}
}

// Struct validation
func (b *BaseModel) validateStruct(err error) error {

	errMap := make(map[string]interface{})

	if err != nil {
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					errMap[err.Field()] = err.Field() + " is required"
				case "email":
					errMap[err.Field()] = err.Field() + " not valid"
				case "gte":
					errMap[err.Field()] = err.Field() + " value must be greater than " + err.Param()
				case "lte":
					errMap[err.Field()] = err.Field() + " value must be lower than " + err.Param()
				}
			}
		}
	}

	// Assign error collection into struct
	b.ErrorsMap = errMap

	return err
}

// IsHasStructError Check is current struct validation has an error
func (b *BaseModel) IsHasStructError() bool {

	if len(b.ErrorsMap) > 0 {
		return true
	}

	return false
}

func (b *BaseModel) GenerateSlug(t string, n string, scope *gorm.Scope) {
	var total int
	slug := slugify.Slugify(n)
	GetDB().Table(t).Where("slug LIKE ?", slug+"%").Count(&total)

	if total > 1 && !scope.PrimaryKeyZero() {
		slug = slug + "-" + cast.ToString(total-1)
	} else if scope.PrimaryKeyZero() && total >= 1 {
		slug = slug + "-" + cast.ToString(total)
	}

	scope.SetColumn("slug", slug)
}

func (b *BaseModel) GetConfig(name string) interface{} {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$GOPATH/src/" + os.Getenv("APP_NAME") + "/config")

	err := viper.ReadInConfig()
	if err != nil {
		utils.Log("error", err)
	}

	return viper.Get(name)
}
