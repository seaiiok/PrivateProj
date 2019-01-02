package api

// D3Weight door 3 car weighting system
type D3Weight struct {
	consumerInfo
}

// getConsumerPing ...
func (me *D3Weight) getConsumerPing() (err error) {
	err = me.Init()
   return
}

// getConsumerKeys ...
func (me *D3Weight) getConsumerKeys() (err error) {
	me.Cols, err = me.Query(me.SqlSelectCols)
	return
}

// setConsumerRows ...
func (me *D3Weight) setConsumerRows() (err error) {
	err = me.Inserts(me.SqlInsertsRows, me.Rows)
	return
}
