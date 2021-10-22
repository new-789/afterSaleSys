package db

import (
	"fmt"
)

// QuertRetAdminDate 查询所有退货信息
func QuertRetAdminDate(sType string, data *Data) (goods []*ReturnGoods, err error) {
	switch sType {
	case "index":
		sqlStr := `select 
       					ldsnum,ldsname,matename,reason,renum,paid,method,stat,crtdate,dealdate,sendnum,ssid,ccid 
				   from returnGoods;`
		err = Db.Select(&goods, sqlStr)
		if err != nil {
			fmt.Println("=============== 退货后台首页查询所有数据错误 ===============", err)
			return nil, err
		}
	case "recv":
		sqlStr := `select 
						ldsnum,ldsname,matename,reason,renum,paid,method,stat,crtdate,dealdate,sendnum,ssid,ccid 
   				   from 
						returnGoods
				   where ldsnum=? and matename=?;`
		err = Db.Select(&goods, sqlStr, data.LdsNum, data.MateName)
		if err != nil {
			fmt.Println("=============== 退货后台接收时查询数据错误 ===============", err)
			return nil, err
		}
	case "send":
		sqlStr := `select 
						ldsnum,ldsname,matename,reason,renum,paid,method,stat,crtdate,dealdate,sendnum,ssid,ccid 
   				   from 
						returnGoods
				   where ldsnum=? and method=?;`
		err = Db.Select(&goods, sqlStr, data.LdsNum, data.Method)
		if err != nil {
			fmt.Println("=============== 退货后台发货时查询已换数量数据错误 ===============", err)
			return nil, err
		}
	}
	return
}

// InsertGoods 客户提交退货信息插入数据到数据库
func InsertGoods(data *Data) bool {
	sqlStr := `insert into returnGoods(ldsnum,ldsname,matename,reason,renum,stat,crtdate,ssid,ccid) values(?,?,?,?,?,?,?,?,?);`
	_, err := Db.Exec(sqlStr, data.LdsNum, data.LdsName, data.MateName, data.Reason, data.Renum, data.Stat, data.CrtDate, data.Ssid, data.Ccid)
	if err != nil {
		fmt.Println("============= 客户提交退货时插入数据库失败 ===============", err)
		return false
	}
	return true
}

// UpdateGoods 退货后台操作更新数据库操作
func UpdateGoods(Type, oldLdsNum string, data *Data) bool {
	switch Type {
	case "recv": // 退货后台收货更新实收数量和状态为1
		sqlStr := `update returnGoods set paid=?,stat=? where ldsnum=? and matename=?;`
		_, err := Db.Exec(sqlStr, data.ReturnGoods.Paid, data.Stat, data.LdsNum, data.MateName)
		if err != nil {
			fmt.Println("=========== 退货后台仓库收货更新实收数量和状态数据库失败 ==========", err)
			return false
		}
	case "maintain":
		sqlStr := `update returnGoods set method=?,stat=?,dealdate=? where ldsnum=? and matename=? and ccid=?;`
		_, err := Db.Exec(sqlStr, data.Method, data.Stat, data.DealDate, data.LdsNum, data.MateName, data.Ccid)
		if err != nil {
			fmt.Println("=========== 退货后台销售处理更新数据处理方法及状态处理时间库失败 ==========", err)
			return false
		}
	case "send":
		sqlStr := `update returnGoods set ldsnum=?,stat=?,sendnum=? where ldsnum=? and method=?`
		_, err := Db.Exec(sqlStr, data.LdsNum, data.Stat, data.SendNum, oldLdsNum, data.Method)
		if err != nil {
			fmt.Println("=========== 退货后台仓库发货更新数据处理方法及状态处理时间库失败 ==========", err)
			return false
		}
	}
	return true
}

// SelectReturnData 搜索功能数据库操作
func SelectReturnData(ldssnum string) (returnList []*ReturnGoods, err error) {
	//sqlStr := `select date,ldssnum,lgss,mateInfo,amount,bad_rea,status,c_id from ldss where ldssnum=?;`
	sqlStr := `select 
       				ldsnum,ldsname,matename,reason,renum,paid,method,stat,crtdate,dealdate,sendnum,ssid,ccid
			   from returnGoods 
			   where 
			         ldsnum=? or ldsnum like ? or matename like ? or reason like ? or ssid=? or ccid=?;`
	err = Db.Select(&returnList, sqlStr, ldssnum, "%"+ldssnum, "%"+ldssnum+"%", "%"+ldssnum+"%", ldssnum, ldssnum)
	if err != nil {
		fmt.Println("========== 搜索时按照快递单号查询数据失败 ==========", err)
		return nil, err
	}
	return
}

// SelectDateFilterData 按时间过滤数据数据库操作
func SelectDateFilterData(startDate, enDate, stat string) (returnList []*ReturnGoods, err error) {
	sqlStr := ``
	if stat == "4" { // 仅按时间进行过滤
		sqlStr = `select 
       				ldsnum,ldsname,matename,reason,renum,paid,method,stat,crtdate,dealdate,sendnum,ssid,ccid
			   from 
					returnGoods 
			   where 
			         crtdate between ? and ?  or dealdate between ? and ?;`
		err := Db.Select(&returnList, sqlStr, startDate, enDate, startDate, enDate)
		if err != nil {
			fmt.Println("================= 仅按时间过滤查询数据库失败 ==================", err)
			return nil, err
		}
		// 仅按状态进行过滤
	} else if (startDate == "" || enDate == "") && (stat == "0" || stat == "1" || stat == "2" || stat == "3") {
		sqlStr = `select 
       				ldsnum,ldsname,matename,reason,renum,paid,method,stat,crtdate,dealdate,sendnum,ssid,ccid
			   from 
					returnGoods 
			   where 
					stat=? ;`
		err = Db.Select(&returnList, sqlStr, stat)
		if err != nil {
			fmt.Println("================= 仅按状态过滤查询数据库失败 ==================", err)
			return nil, err
		}
	} else { // 安状态和时间同时进行过滤
		sqlStr = `select 
       				ldsnum,ldsname,matename,reason,renum,paid,method,stat,crtdate,dealdate,sendnum,ssid,ccid
			   from 
					returnGoods 
			   where 
			         stat=? and crtdate between ? and ?  or stat=? and dealdate between ? and ?;`
		err = Db.Select(&returnList, sqlStr, stat, startDate, enDate, stat, startDate, enDate)
		if err != nil {
			fmt.Println("================= 仅按状态和时间同时过滤查询数据库失败 ==================", err)
			return nil, err
		}
	}
	return
}
