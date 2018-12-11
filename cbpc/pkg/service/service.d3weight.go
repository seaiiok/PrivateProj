package service

import(
"database/sql"
)

// ServerDatebaseGetReady(req *Proto) *Proto
// ServerDatebaseGetData(req *Proto) *Proto
// ServerDatebaseSetData(req *Proto) *Proto

//3#门车辆称重实体
type D3WeightSql struct {
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

type D3Weight struct {
    ArrayD3Weight []D3WeightSql	
}

func (p *D3Weight)ServerDatebaseGetReady(req *Proto) *Proto {
	req.myBody.DataContent=append(req.myBody.DataContent,p.ArrayD3Weight)
	return req
}

func (p *D3Weight)ServerDatebaseGetData(req *Proto) *Proto {
	req.myBody.DataContent=append(req.myBody.DataContent,p.ArrayD3Weight)
	return req
}

func (p *D3Weight)ServerDatebaseSetData(req *Proto) *Proto {
	req.myBody.DataContent=append(req.myBody.DataContent,p.ArrayD3Weight)
	return req
}