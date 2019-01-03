package cert

import (
	"testing"
)

func TestChengDuMES4OPC(t *testing.T) {

	baseinfo := CertInformation{
		Country:            []string{"CN"},
		Organization:       []string{"MES-shucai"},
		IsCA:               true,
		OrganizationalUnit: []string{"MES-Work"},
		EmailAddress:       []string{"seaii@qq.com"},
		Locality:           []string{"ChengDu"},
		Province:           []string{"SiChuang"},
		CommonName:         "MES-Work",
		CrtName:            "server.crt",
		KeyName:            "server.key",
	}

	err := CreateCRT(nil, nil, baseinfo)
	if err != nil {
		t.Log("Create crt error,Error info:", err)

	}

}
