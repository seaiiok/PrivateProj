package api

// Consumer server database interface,
// getConsumerKeys get database key col, 
// setConsumerRows insert database client data rows,
type Consumer interface {
	getConsumerPing() error
	getConsumerKeys() error
	setConsumerRows() error
}
