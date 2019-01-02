package main

import (
	"ifix.cbpc/cbpc/internal/app/clients.app"
)

func main() {
	app.ServiceStart()
	select {}
}
