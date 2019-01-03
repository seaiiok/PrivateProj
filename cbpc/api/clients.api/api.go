package api

// Producer server database interface,
// getProducerKeys get database key col, 
// getProducerRows insert database client data rows,
type Producer interface {
	getProducerPing() error
	getProducerKeys() error
	getProducerRows() error
}
