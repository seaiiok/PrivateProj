/* 
Package api.

type Device
Device user protocol

type Device struct {
    DeviceName   string
    DeviceIP     string
    DeviceTask   string
    DeviceRouter string
}
type Err
Err user protocol

type Err struct {
    Err error
    // contains filtered or unexported fields
}
type Proto
Proto user protocol

type Proto struct {
    Err
    SQL
    Device
}
func NewProto
func NewProto() *Proto
NewProto init user proto

func (*Proto) GetData
func (p *Proto) GetData() [][]string
GetData return data

func (*Proto) GetFulQuery
func (p *Proto) GetFulQuery() string
GetFulQuery return query

func (*Proto) GetProcessTrace
func (p *Proto) GetProcessTrace() []string
GetProcessTrace add ProcessTrace

func (*Proto) GetQuery
func (p *Proto) GetQuery() string
GetQuery return query

func (*Proto) SetArgs
func (p *Proto) SetArgs(args []string)
SetArgs set args

func (*Proto) SetData
func (p *Proto) SetData(data [][]string)
SetData set data

func (*Proto) SetError
func (p *Proto) SetError(err error)
SetError set data

func (*Proto) SetProcessTrace
func (p *Proto) SetProcessTrace(str ...string)
SetProcessTrace add ProcessTrace

func (*Proto) SetRouter
func (p *Proto) SetRouter(router string)
SetRouter set data

type SQL
SQL user protocol

type SQL struct {
    DatabaseDriver string
    DatabaseSource string
    InsertSQL      string
    QuerySQL       string
    Args           []string
    Data           [][]string
}
*/
package api