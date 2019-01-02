package convert

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

//Struct2Reader ...
func Struct2Reader(t interface{}) (io.Reader, error) {
	b, err := json.Marshal(t)
	return bytes.NewReader(b), err
}

//Reader2Struct ...
func Reader2Struct(r io.Reader, t interface{}) error {
	b, err := ioutil.ReadAll(r)
	json.Unmarshal(b, &t)
	return err
}

//Struct2Arraybytes ...
func Struct2Arraybytes(t interface{}) ([]byte, error) {
	b, err := json.Marshal(t)
	return b, err
}

//Map2Struct ...
func Map2Struct(m []map[string]string, s interface{}) (res []interface{}) {
	for i := 0; i < len(m); i++ {
		jsonStr, _ := json.Marshal(m[i])
		json.Unmarshal(jsonStr, &s)
		res = append(res, s)
	}
	return
}
