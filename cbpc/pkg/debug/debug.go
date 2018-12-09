// Package orm is just a simple orm.
// Created by simplejia [9/2016]
package main

import (
	"fmt"
	"ii.sea/db"
	_ "ii.sea/libs/odbc"
	"database/sql"
)


func main(){
insertstr:=`insert into T_Standard1 (F_ID,F_StdNo,F_CarNo,F_ProName,F_Model,F_UnitFrom,F_Load,F_UnitTo,F_UnLoad,F_Transport,F_Driver,F_Memo,F_Number,F_Type,F_BanCi,F_Gross,F_Tare,F_Net,F_Deduct,F_Fact,F_Deduct2,F_Fact2,F_SettWgt,F_Price,F_Money,F_Price2,F_Money2,F_SumMoney,F_CardID,F_Unit,F_GrossTime,F_TareTime,F_BeginTime,F_EndTime,F_CNTime,F_opManGross,F_opManTare,F_opMan,F_edMan,F_edTime,F_RePrint,F_PrTimes,F_PrMan,F_PrTime,F_PrScaleName,F_ScaleNameGross,F_ScaleNameTare,F_ScaleNameNet,F_IsCancel,F_IsFinish,F_CanUpdate,F_IsUpdate,F_DisEdit,F_MacNo,F_Status,F_SerialNo,F_SerialNoDay,F_SerialNoMonth,F_SerialNoYear,F_WgtType,F_YLStr01,F_YLStr02,F_YLStr03,F_YLStr04,F_YLStr05,F_YLStr06,F_YLStr07,F_YLStr08,F_YLStr09,F_YLStr10,F_YLFlt01,F_YLFlt02,F_YLFlt03,F_YLFlt04,F_YLFlt05,F_YLFlt06,F_YLFlt07,F_YLFlt08,F_YLFlt09,F_YLFlt10,F_YLInt01,F_YLInt02,F_YLInt03,F_YLInt04,F_YLInt05,F_OrderNo) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
err:=db.DBInit("odbc","driver={SQL Server};Server=192.168.100.197;uid=sa;pwd=kmtSoft12345678;port=1433;Database=CMTWeight")
if err != nil {
	fmt.Println(err)
}
err=db.DBPing()
if err != nil {
	fmt.Println(err)
}

arr,err:=db.DBExecSelectCols("select * from T_Standard ")
if err != nil {
	fmt.Println(err)
}
for i := 0; i < len(arr); i++ {
	db.DBExecInsertCols(insertstr,arr[i])
}

// arr1,err:=db.DBExecSelectCol("select top 1 * from T_Standard")
// if err != nil {
// 	fmt.Println(err)
// }
// for i := 0; i < len(arr1); i++ {
// 	fmt.Println(arr1[i])
// }

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
