package constants

import "os"

type ConstantsCode string

const (
	envName      string = "ENVIRONMENT"
	EnvSecretkey string = "SECRETKEY"
	DbName       string = "DB_NAME"
	DbUser       string = "DB_USER"
	DbPass       string = "DB_PASS"
	DbHost       string = "DB_HOST"
	MaxFileSize  int64  = int64(3 * 1024 * 1024)
	ImgurUrl     string = "https://api.imgur.com/3/upload"
)

const (
	ContextDbConnection    ConstantsCode = "contextDbConnection"
	UnexpectedErrorMessage ConstantsCode = "Unexpected error"
)

var Environment ConstantsCode

func init() {
	Environment = ConstantsCode(os.Getenv(envName))
	if Environment == "" {
		Environment = ConstantsCode(Production)
	}

}
