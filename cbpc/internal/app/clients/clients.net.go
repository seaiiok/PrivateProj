package clients

import (
	"ifix.cbpc/cbpc/internal/pkg"
)

// ServiceStart . 
func ServiceStart() {
	clientsInit()
	pkg.ClientsHTTPStart()
}
