package responsebuilder

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BadRequest(context *gin.Context, err error) {
	fmt.Println(err.Error())
	var ve validator.ValidationErrors
	errors.As(err, &ve)
	apiError := make([]string, len(ve))
	for i, fe := range ve {
		apiError[i] = fe.StructNamespace() + " - failed on validation '" + fe.Tag() + "'"

	}
	context.JSON(http.StatusBadRequest, gin.H{
		"errors": apiError,
	})
}
