package utils

import(
"reflect"
"strings"
"database/sql"
)


func Rows2Strus(rows *sql.Rows, strus interface{}) (err error) {
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	strusRV := reflect.Indirect(reflect.ValueOf(strus))
	elemRT := strusRV.Type().Elem()
	fieldInfo := getFieldInfo(elemRT)
	for rows.Next() {
		var struRV reflect.Value
		var struField reflect.Value
		if elemRT.Kind() == reflect.Ptr {
			struRV = reflect.New(elemRT.Elem())
			struField = reflect.Indirect(struRV)
		} else {
			struRV = reflect.Indirect(reflect.New(elemRT))
			struField = struRV
		}
		var values []interface{}
		for _, column := range columns {
			idx, ok := fieldInfo[strings.ToLower(column)]
			var v interface{}
			if !ok {
				var i interface{}
				v = &i
			} else {
				v = struField.FieldByIndex(idx).Addr().Interface()
			}
			values = append(values, v)
		 }

		err = rows.Scan(values...)
		if err != nil {
			return
		}

	
		strusRV = reflect.Append(strusRV, struRV)
	}
	if err = rows.Err(); err != nil {
		return
	}
	reflect.Indirect(reflect.ValueOf(strus)).Set(strusRV)

	return
}

var TagName = "orm"
func getFieldInfo(typ reflect.Type) map[string][]int {
	if typ.Kind() == reflect.Ptr {		
		typ = typ.Elem()
	}

	finfo := make(map[string][]int)

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		tag := f.Tag.Get(TagName)

		// Skip unexported fields or fields marked with "-"
		if f.PkgPath != "" || tag == "-" {
			continue
		}

		// Handle embedded structs
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			for k, v := range getFieldInfo(f.Type) {
				finfo[k] = append(f.Index, v...)
			}
			continue
		}

		// Use field name for untagged fields
		if tag == "" {
			tag = f.Name
		}

		tag = strings.ToLower(tag)

		finfo[tag] = f.Index
	}
	return finfo
}

//严格的 数组转结构体  var stru struct
func Array2Struct(arr []string,stru interface{}) {
	fs:=reflect.ValueOf(stru).Elem()
	for i := 0; i < fs.NumField(); i++ {
		fs.Field(i).SetString(arr[i])	
	}
}
