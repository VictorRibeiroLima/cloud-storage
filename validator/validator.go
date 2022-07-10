package validator

import (
	"github.com/VictorRibeiroLima/cloud-storage/database"
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

var unique validator.Func = func(fl validator.FieldLevel) bool {
	db := database.DbConnection
	table := fl.Param()
	field := fl.FieldName()
	value := fl.Field().String()
	query := "SELECT count(*) FROM " + table + " WHERE " + field + " = ?"
	var count int
	db.Raw(query, value).Scan(&count)
	return count == 0
}

func BindValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique", unique)
	}
}
