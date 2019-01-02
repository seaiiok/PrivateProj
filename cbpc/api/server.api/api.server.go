package api

import (
	"ifix.cbpc/cbpc/pkg/db"
)

// consumerInfo ...
type consumerInfo struct {
	db.DB
	SqlSelectCols  string
	SqlInsertsRows string
	Cols           []string
	Rows           [][]string
}

// GetConsumerPing ping
func GetConsumerPing(c Consumer) error {
	return c.getConsumerPing()
}

// GetConsumerKeys get database keys
func GetConsumerKeys(c Consumer) error {
	return c.getConsumerKeys()
}

// SetConsumerRows insert data into database
func SetConsumerRows(c Consumer) error {
	return c.setConsumerRows()

}
