package api

// D3Weight door 3 car weighting system
type D3Weight struct {
	producerInfo
}

// getConsumerPing ...
func (me *D3Weight) getProducerPing() (err error) {
	err = me.Init()
	return
}

// getConsumerKeys ...
func (me *D3Weight) getProducerKeys() (err error) {
	me.Cols, err = me.Query(me.SqlSelectCols)
	return
}

// setConsumerRows ...
func (me *D3Weight) getProducerRows() (err error) {
	me.Rows,err= me.Querys(me.SqlSelectRows)
	return
}
