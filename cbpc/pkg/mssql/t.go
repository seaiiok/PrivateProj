package main

import (
	"fmt"
"ii.sea/utils"
"strconv"
)

func main(){
	arr:=make([]string,86)
	for i := 0; i < len(arr)-10; i++ {
		arr[i]=strconv.Itoa(i)
	}
// stru:=new(Demo)
var stru Demo
utils.Array2Struct(arr,&stru)
fmt.Println(stru)
}


//3#门车辆称重实体
type Demo struct {
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
