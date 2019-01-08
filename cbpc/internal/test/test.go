package main
 
import (
  "fmt"
  "reflect"
)
 
type Injector struct {
  mappers map[reflect.Type]reflect.Value
}
 
func (inj *Injector) Set(value interface{}) {
  inj.mappers[reflect.TypeOf(value)] = reflect.ValueOf(value)
}
 
func (inj *Injector) Get(t reflect.Type) reflect.Value {
  return inj.mappers[t]
}
 
func (inj *Injector) Invoke(i interface{}) []reflect.Value {
  t := reflect.TypeOf(i)
  if t.Kind() != reflect.Func {
	return nil
  }
  inValues := make([]reflect.Value, t.NumIn())
  for k := 0; k < t.NumIn(); k++ {
    inValues[k] = inj.Get(t.In(k))
  }
  ret := reflect.ValueOf(i).Call(inValues)
  return ret
}
 
func New() *Injector {
  return &Injector{make(map[reflect.Type]reflect.Value)}
}
 
func Container(f interface{}) { 
  inj.Invoke(f) 
}
 
func Dependency(b []string) {
  fmt.Println("exsql: ", b)
}

func HI(b string) {
	fmt.Println("heeo: ", b)
  }
 
var inj *Injector
 
func main() {
  inj = New()
  b:=[]string{"a","b"}
  inj.Set(b)

 
  d := Dependency
  Container(d)


  inj.Set("b1")
  Container(HI)

}
