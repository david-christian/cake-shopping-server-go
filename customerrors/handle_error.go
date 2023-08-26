package customerrors

import (
	"errors"
	"fmt"
	"shopingCar_go/constants"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func DefinitionError(err error) (errorCode, message string) {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		mysqlErr = err.(*mysql.MySQLError)
		if mysqlErr.Number == 1062 {
			errorCode = string(constants.Duplicated)
			message = "Duplicate entry error"
		}
	}

	if err == gorm.ErrRecordNotFound {
		errorCode = string(constants.Invalid)
		message = "No rows found"
	}

	if errorCode == "" && message == "" {
		errorCode = string(constants.Unexpected)
		message = "unexpected error"
	}

	if constants.Environment == constants.Develop {
		message += fmt.Sprintf(" system error: %s", err.Error())
	}

	return errorCode, message

}
