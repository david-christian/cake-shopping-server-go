package models

type SusessRes struct {
	Ok int8
}

type ErrorRes struct {
	Ok      int8
	Message string
	Status  string
}
