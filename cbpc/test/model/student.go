package model

import(
	"time"
	"net/http"
)
type Retriever struct {
	UserAgen string
	TimeOut time.Duration
}

func (r Retriever) Get(url string) string {
	resp,err:=http.Get(url)
	if err != nil {
		panic(err)
	}
}

