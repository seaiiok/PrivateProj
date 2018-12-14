package clients

import (
	"ifix.cbpc/cbpc/internal/pkg"
)

// ServiceStart . 
func ServiceStart() {
	serverInit()
	pkg.ClientsHTTPStart()
}
