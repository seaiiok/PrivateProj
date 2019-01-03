package api

import (
	"ifix.cbpc/cbpc/pkg/db"
)

// producerInfo ...
type producerInfo struct {
	db.DB
	SqlSelectCols  string
	SqlSelectRows string
	Cols           []string
	Rows           [][]string
}

// GetProducerPing ping
func GetProducerPing(p Producer) error {
	return p.getProducerPing()
}

// GetProducerKeys get database keys
func GetProducerKeys(p Producer) error {
	return p.getProducerKeys()
}

// GetProducerRows insert data into database
func GetProducerRows(p Producer) error {
	return p.getProducerRows()

}
