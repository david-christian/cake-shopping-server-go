package constants

type ErrorCode string

const (
	Invalid                  ErrorCode = "Invalid"
	Unauthorized             ErrorCode = "Unauthorized"
	Duplicated               ErrorCode = "Duplicated"
	Unexpected               ErrorCode = "Unexpected"
	TypeError                ErrorCode = "type_error"
	MaxFileError             ErrorCode = "max_file_error"
	NotAllowedExtensionError ErrorCode = "not_allowed_extension_error"
	FileConversionError      ErrorCode = "file conversion error"
)
