package models

type FileUpload struct {
	Transactions []TransactionCreate `json:"transactions"`
}
