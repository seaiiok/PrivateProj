package api

// Empty is a empty interface
type Empty struct {
	producerInfo
}

// getProducerPing ...
func (me *Empty) getProducerPing() (err error) {
	err = me.Init()
   return
}

// getProducerKeys ...
func (me *Empty) getProducerKeys() (err error) {
	return nil
}

// setProducerRows ...
func (me *Empty) getProducerRows() (err error) {
	return nil
}
