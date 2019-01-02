package net

import (
	"net/http"

	"ifix.cbpc/cbpc/pkg/protocol"
)

var client *http.Client

type ClientsProto struct {
	protocol.Protocol
}
