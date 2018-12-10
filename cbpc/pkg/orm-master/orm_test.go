package orm

import (
	"database/sql"

	_ "ii.sea/libs/odbc"
"encoding/json"
	"testing"

)


func getDb() *sql.DB {
	db, err := sql.Open("odbc", "driver={SQL Server};Server=192.168.1.18;uid=sa;pwd=kmtSoft12345678;port=1433;Database=CMTWeight121702")
	if err != nil {
		panic(err)
	}
	return db
}

func TestRows2Strus(t *testing.T) {
	db := getDb()

	rows, err := db.Query("select * from T_Standard where F_ID=? or F_ID=?", 191, 192)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	var d []*Demo
	err = Rows2Strus(rows, &d)
	if err != nil {
		t.Fatal(err)
	}
	for p, v := range d {
		t.Log(p, v)
	}
}

func TestRows2Stru(t *testing.T) {
	db := getDb()

	rows, err := db.Query("select * from T_Standard where F_ID=? or F_ID=?", 191, 192)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	var d *Demo
	err = Rows2Stru(rows, &d)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(d)

	b, err := json.Marshal(d)
    if err != nil {
		t.Log(b,err)
	}
	arr:=make([]string,0)
	for _, v := range d {
		arr=append(arr,v)
	}
	return b
}

func TestRows2Cnts(t *testing.T) {
	db := getDb()

	rows, err := db.Query("select F_ID from T_Standard")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	var d []sql.NullString // 当返回值有可能为Null时(比如select max时)，应使用[]sql.NullInt64
	err = Rows2Cnts(rows, &d)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(d)
}

func TestRows2Cnt(t *testing.T) {
	db := getDb()

	rows, err := db.Query("select F_ID from T_Standard where F_ID=?", 191)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	var d sql.NullString
	err = Rows2Cnt(rows, &d)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(d)

}



//3#门车辆称重实体
type Demo1 struct {
	F_ID string 		
	F_StdNo string		
	F_CarNo string		 
	F_ProName string		
	F_Model string		
	F_UnitFrom string	
	F_Load string		
	F_UnitTo string		
	F_UnLoad string		
	F_Transport string		
	F_Driver string		
	F_Memo string
	F_Number string		
	F_Type string		
	F_BanCi string		
	F_Gross string	
	F_Tare string		
	F_Net string		
	F_Deduct string		
	F_Fact string		 
	F_Deduct2 string		
	F_Fact2 string		
	F_SettWgt string		
	F_Price string		
	F_Money string		
	F_Price2 string		
	F_Money2 string		
	F_SumMoney string		
	F_CardID string		
	F_Unit string		
	F_GrossTime string		 
	F_TareTime string		
	F_BeginTime string		
	F_EndTime string		 
	F_CNTime string		
	F_opManGross string		
	F_opManTare string		
	F_opMan string		
	F_edMan string		
	F_edTime string		 
	F_RePrint string		
	F_PrTimes string	
	F_PrMan string		
	F_PrTime string		
	F_PrScaleName string		
	F_ScaleNameGross string	
	F_ScaleNameTare string		
	F_ScaleNameNet string	
	F_IsCancel string		
	F_IsFinish string	
	F_CanUpdate string	
	F_IsUpdate string		
	F_DisEdit string		
	F_MacNo string		
	F_Status string		
	F_SerialNo string		 
	F_SerialNoDay string		
	F_SerialNoMonth string		
	F_SerialNoYear string		
	F_WgtType string		
	F_YLStr01 string	
	F_YLStr02 string	
	F_YLStr03 string		
	F_YLStr04 string		
	F_YLStr05 string	
	F_YLStr06 string		
	F_YLStr07 string		
	F_YLStr08 string	
	F_YLStr09 string		
	F_YLStr10 string	
	F_YLFlt01 string		
	F_YLFlt02 string		
	F_YLFlt03 string		
	F_YLFlt04 string		
	F_YLFlt05 string		
	F_YLFlt06 string	
	F_YLFlt07 string		
	F_YLFlt08 string		
	F_YLFlt09 string		
	F_YLFlt10 string		
	F_YLInt01 string		
	F_YLInt02 string	
	F_YLInt03 string		
	F_YLInt04 string		
	F_YLInt05 string		
	F_OrderNo string		 
}

//3#门车辆称重实体
type Demo struct {
	F_ID sql.NullString 		
	F_StdNo sql.NullString		
	F_CarNo sql.NullString		 
	F_ProName sql.NullString		
	F_Model sql.NullString		
	F_UnitFrom sql.NullString	
	F_Load sql.NullString		
	F_UnitTo sql.NullString		
	F_UnLoad sql.NullString		
	F_Transport sql.NullString		
	F_Driver sql.NullString		
	F_Memo sql.NullString
	F_Number sql.NullString		
	F_Type sql.NullString		
	F_BanCi sql.NullString		
	F_Gross sql.NullString	
	F_Tare sql.NullString		
	F_Net sql.NullString		
	F_Deduct sql.NullString		
	F_Fact sql.NullString		 
	F_Deduct2 sql.NullString		
	F_Fact2 sql.NullString		
	F_SettWgt sql.NullString		
	F_Price sql.NullString		
	F_Money sql.NullString		
	F_Price2 sql.NullString		
	F_Money2 sql.NullString		
	F_SumMoney sql.NullString		
	F_CardID sql.NullString		
	F_Unit sql.NullString		
	F_GrossTime sql.NullString		 
	F_TareTime sql.NullString		
	F_BeginTime sql.NullString		
	F_EndTime sql.NullString		 
	F_CNTime sql.NullString		
	F_opManGross sql.NullString		
	F_opManTare sql.NullString		
	F_opMan sql.NullString		
	F_edMan sql.NullString		
	F_edTime sql.NullString		 
	F_RePrint sql.NullString		
	F_PrTimes sql.NullString	
	F_PrMan sql.NullString		
	F_PrTime sql.NullString		
	F_PrScaleName sql.NullString		
	F_ScaleNameGross sql.NullString	
	F_ScaleNameTare sql.NullString		
	F_ScaleNameNet sql.NullString	
	F_IsCancel sql.NullString		
	F_IsFinish sql.NullString	
	F_CanUpdate sql.NullString	
	F_IsUpdate sql.NullString		
	F_DisEdit sql.NullString		
	F_MacNo sql.NullString		
	F_Status sql.NullString		
	F_SerialNo sql.NullString		 
	F_SerialNoDay sql.NullString		
	F_SerialNoMonth sql.NullString		
	F_SerialNoYear sql.NullString		
	F_WgtType sql.NullString		
	F_YLStr01 sql.NullString	
	F_YLStr02 sql.NullString	
	F_YLStr03 sql.NullString		
	F_YLStr04 sql.NullString		
	F_YLStr05 sql.NullString	
	F_YLStr06 sql.NullString		
	F_YLStr07 sql.NullString		
	F_YLStr08 sql.NullString	
	F_YLStr09 sql.NullString		
	F_YLStr10 sql.NullString	
	F_YLFlt01 sql.NullString		
	F_YLFlt02 sql.NullString		
	F_YLFlt03 sql.NullString		
	F_YLFlt04 sql.NullString		
	F_YLFlt05 sql.NullString		
	F_YLFlt06 sql.NullString	
	F_YLFlt07 sql.NullString		
	F_YLFlt08 sql.NullString		
	F_YLFlt09 sql.NullString		
	F_YLFlt10 sql.NullString		
	F_YLInt01 sql.NullString		
	F_YLInt02 sql.NullString	
	F_YLInt03 sql.NullString		
	F_YLInt04 sql.NullString		
	F_YLInt05 sql.NullString		
	F_OrderNo sql.NullString	
}