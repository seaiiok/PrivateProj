package api

// Empty is a empty interface
type Empty struct {
	consumerInfo
}

// getConsumerPing ...
func (me *Empty) getConsumerPing() (err error) {
	err = me.Init()
   return
}

// getConsumerKeys ...
func (me *Empty) getConsumerKeys() (err error) {
	return nil
}

// setConsumerRows ...
func (me *Empty) setConsumerRows() (err error) {
	return nil
}
