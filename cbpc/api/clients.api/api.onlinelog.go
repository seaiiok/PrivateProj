package api

// OnlineLog online log ...
type OnlineLog struct {
	producerInfo
}

// getConsumerPing ...
func (me *OnlineLog) getConsumerPing() (err error) {
	 err = me.Init()
	return
}

// getConsumerKeys ...
func (me *OnlineLog) getConsumerKeys() (err error) {
	me.Cols, err = me.Query(me.SqlSelectCols)
	return
}

// setConsumerRows ...
func (me *OnlineLog) setConsumerRows() (err error) {
	me.Rows,err= me.Querys(me.SqlSelectRows)
	return
}
