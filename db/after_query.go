package db

import (
	"errors"
	"fmt"
)

func QueryPhone(str string) (data []*ClientAddr, err error) {
	//phone := &ClientAddr{}
	dsn := `select cid,phone from clientAddr where phone=?;`
	err = Db.Select(&data, dsn, str)
	if err != nil {
		fmt.Println("=============== 注册时查询手机号码为空，开始进行注册 ==============", err)
		return nil, err
	}
	//if phone.Phone == "" {
	//	return nil, errors.New("您的号码不存在！")
	//}
	return
}

func QueryUserType(uid string) (utype string) {
	sqlStr := `select utype from user where uid=?`
	err := Db.Get(&utype, sqlStr, uid)
	if err != nil {
		fmt.Println()
		return ""
	}
	return
}

func InsertToClientTable(data *Data) (bool, error) {
	sqlStr := `insert into clientAddr(username,phone) values(?, ?);`
	_, err := Db.Exec(sqlStr, data.ClientAddr.UserName, data.ClientAddr.Phone)
	if err != nil {
		fmt.Println("==============插入数到 clientTable 表失败=============", err)
		return false, errors.New("插入数据库失败")
	}
	return true, nil
}

// QuerySale 下拉菜单数据库数据查询
func QuerySale() (sales []*Sale, err error) {
	sqlStr := `select sid,saleName,phone from sale;`
	err = Db.Select(&sales, sqlStr)
	if err != nil {
		return nil, err
	}
	return
}

func QueryClient(str string) (clients []*ClientAddr, err error) {
	sqlStr := `select cid,username,phone,address from clientAddr where phone=?;`
	err = Db.Select(&clients, sqlStr, str)
	if err != nil {
		return nil, err
	}
	return
}

// QueryLdssInfo 管理后台首页查询数据操作
func QueryLdssInfo(id int) (ldssAry []*Ldss, err error) {
	switch id {
	case 0: // 后台管理查询所有工单数据
		sqlStr := `select 
						lid,date,ldssnum,lgss,mateInfo,amount,paid,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id 
				   from 
				   		ldss 
				   order by 
						date
				   desc;`
		err = Db.Select(&ldssAry, sqlStr)
		if err != nil {
			fmt.Println("========== 后台查询 Lgss 数据表失败1 =========", err)
			return nil, err
		}
	default: // 我的工单数据查询操作,及按通知点击用户处理查询数据操作
		sqlStr := `select 
						date,ldssnum,lgss,mateInfo,amount,paid,pstate,sentnum,bad_rea,status,c_id 
				   from 
				   		ldss 
				   where 
				   		c_id=? and status=? or c_id=? and status=? or c_id=? and  status=?;`
		err = Db.Select(&ldssAry, sqlStr, id, 0, id, 1, id, 2)
		if err != nil {
			fmt.Println("========== 按客户ID查询 Lgss 数据表失败3 =========", err)
			return nil, err
		}
	}
	return
}

func QueryPaidStat(cid int) (ldss []*Ldss, err error) {
	// 接收返回通知数据库查询实际接受数量、实际接收状态、以及物流单号，按客户ID和pstate状态为1查询
	sqlStr := `select date,ldssnum,lgss,mateInfo,amount,paid,pstate,bad_rea,status,c_id from ldss where c_id=? and pstate=?;`
	err = Db.Select(&ldss, sqlStr, cid, 1)
	if err != nil {
		fmt.Println("========== 按客户ID及pstate查询 Lgss 数据表失败2 =========", err)
		return nil, err
	}
	return
}

func InsertToLdsTable(data *Data) error {
	sqlStr := `insert into ldss(date,ldssnum,lgss,mateInfo,amount,bad_rea,status,sale_id,c_id) values(?,?,?,?,?,?,?,?,?);`
	_, err := Db.Exec(sqlStr, data.Date, data.Ldss.LdssNum, data.LgssName, data.MateInfo, data.Amount, data.BadRea, 0, data.SaleId, data.ClientId)
	if err != nil {
		fmt.Println("==============插入数到 Lgss 表失败=============", err)
		return errors.New("插入数据失败")
	}
	return nil
}

func UpLdssData(inType string, data *Data) error {
	switch inType {
	case "ldss": // TODO 销售员更新工单状态以及快递信息
		sqlStr := `update ldss set ldssnum=?,ldss=?, status=? where ldssnum=?`
		_, err := Db.Exec(sqlStr, data.Ldss.LdssNum, data.LgssName, data.Status, data.Ldss.LdssNum)
		if err != nil {
			fmt.Println("==============更新 Lgss 表中状态值失败=============", err)
			return errors.New("更新 Lgss 表中状态数据失败")
		}
	case "client": // 更新客户地址信息
		sqlStr := `update clientAddr set address=? where cid=?`
		_, err := Db.Exec(sqlStr, data.ClientAddr.Address, data.ClientAddr.Cid)
		if err != nil {
			fmt.Println("==============更新 clientAddr 表中地址信息失败=============", err)
			return errors.New("更新 Lgss 表中状态数据失败")
		}
	case "upStatus": // 客户确认收货后更新 status 状态信息
		sqlStr := `update ldss set status=? where ldssnum=? and c_id=?`
		_, err := Db.Exec(sqlStr, data.Status, data.Ldss.LdssNum, data.ClientId)
		if err != nil {
			fmt.Println("============== 客户确认收货更新 ldss 表的 status 字段失败=============", err)
			return errors.New("更新 Lgss 表中状态数据失败")
		}
	default:
		return errors.New("抱歉暂不支持该操作")
	}
	return nil
}

func QuerySaleName(sid int) (saleList []*Sale, err error) {
	sqlStr := `select sid,saleName,phone from sale where sid=?`
	err = Db.Select(&saleList, sqlStr, sid)
	if err != nil {
		fmt.Println("================= 根据用户ID查询数据库失败 ==================", err)
		return nil, err
	}
	return
}

// QueryClientName 按照客户id查询客户名称
func QueryClientName(cid int) (clientName []*ClientAddr, err error) {
	//sqlStr := `select cid,username from clientAddr where cid=?`
	sqlStr := `select cid,username from clientAddr as t1 inner join (select c_id from ldss where status=? or status=? or status=?) as t2 on t1.cid=t2.c_id;`
	err = Db.Select(&clientName, sqlStr, 0, 1, 2)
	if err != nil {
		return nil, err
	}
	return
}

func UpdateStatus(changeType, oldlgssnum string, data *Data) bool {
	switch changeType {
	case "status": // 客户确认实收数量后更新状态方法
		sqlStr := `update ldss set status=? where ldssnum=?;`
		_, err := Db.Exec(sqlStr, 1, oldlgssnum)
		if err != nil {
			fmt.Println("---------- 按照快递单号更新 status 字段失败 ----------", err)
			return false
		}
	case "paid":
		// 更新实收数量状态，1为返回并通知客户，2为客户确认并需要更新status状态为2，3为客户不确认并返回给后台
		sqlStr := `update ldss set paid=?,pstate=?,status=?,remark=? where ldssnum=? and mateInfo=?`
		_, err := Db.Exec(sqlStr, data.Ldss.Paid, data.Pstate, data.Status, data.Remark, oldlgssnum, data.MateInfo)
		if err != nil {
			fmt.Println("---------- 按照快递单号更新 pstate,paid 字段失败 ----------", err)
			return false
		}
	case "ldss": // 发货更新 ldss 表相关子字段数据
		// 修改已发数量数据，以及如果实收数量和已发数量相等则更新 status 状态为3
		sqlStr := `update ldss set ret_date=?,ret_ldss=?,ret_ldssname=?,ret_remark=?,sentnum=?,shipNum=?,status=? where ldssnum=? and mateInfo=?;`
		_, err := Db.Exec(sqlStr, data.RetDate, data.RetLdss, data.RetLdssName, data.RetRemark, data.SentNum, data.ShipNum, data.Status, oldlgssnum, data.MateInfo)
		if err != nil {
			fmt.Println("---------- 按照快递单号更新快递字段及状态字段失败 ----------", err)
			return false
		}
	}
	return true
}

// SelectLdssData 搜索功能后台查询数据库操作
func SelectLdssData(ldssnum string) (ldssList []*Ldss, err error) {
	//sqlStr := `select date,ldssnum,lgss,mateInfo,amount,bad_rea,status,c_id from ldss where ldssnum=?;`
	sqlStr := `select 
       				lid,date,ldssnum,lgss,mateInfo,amount,paid,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id
			   from ldss where ldssnum=? or ldssnum like ? or mateInfo like ? or bad_rea like ? or sale_id=? or c_id=? or ret_ldss like ? or lid=?;` // TODO 新增了最后一个条件
	err = Db.Select(&ldssList, sqlStr, ldssnum, "%"+ldssnum, "%"+ldssnum+"%", "%"+ldssnum+"%", ldssnum, ldssnum, "%"+ldssnum+"%", ldssnum) // todo 新增了一个参数
	if err != nil {
		fmt.Println("========== 搜索时按照快递单号查询数据失败 ==========", err)
		return nil, err
	}
	return
}

// UpPaidStatus 由客户确认收货信息无误后，更新实收货物状态信息
func UpPaidStatus(data *Data) bool {
	sqlStr := `update ldss set pstate=?,status=? where c_id=? and status=?;`
	_, err := Db.Exec(sqlStr, data.Pstate, data.Status, data.ClientId, 0)
	if err != nil {
		fmt.Println("---------- 更新实收货物状态信息失败 ----------", err)
		return false
	}
	return true
}

// SelectClient 根据客户 id 查询客户所有信息
func SelectClient(cid int) (clientAry []*ClientAddr, err error) {
	sqlStr := `select cid,username,phone,address from clientAddr where cid=?;`
	err = Db.Select(&clientAry, sqlStr, cid)
	if err != nil {
		fmt.Println("===================== 根据客户id查询客户信息失败 ===================")
		return nil, err
	}
	return
}

// BackLoginSelect 后台登录数据库查询操作
func BackLoginSelect(username string) (saleInfo []*User, err error) {
	sqlStr := `select uid,username,phone,passwd,utype from user where username=?;`
	err = Db.Select(&saleInfo, sqlStr, username)
	if err != nil {
		fmt.Println("================== 后台登录查询数据库失败 ==================", err)
		return nil, err
	}
	return
}

// SelectDateLdss 按时间或状态过滤数据查询操作
func SelectDateLdss(startDate, endDate string, stat int) (ldss []*Ldss, err error) {
	sqlStr := ``
	if stat == 4 { // 仅按时间间隔过滤数据
		sqlStr = `select 
					lid,date,ldssnum,lgss,mateInfo,amount,paid,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id 
				  from 
				  	ldss 
				  where date between ? and ?;`
		err = Db.Select(&ldss, sqlStr, startDate, endDate)
		if err != nil {
			fmt.Println("============== 按照时间过滤查询数据库失败1 =====================", err)
			return nil, err
		}
	} else if (startDate == "" || endDate == "") && (stat == 0 || stat == 1 || stat == 2 || stat == 3) { // 仅按状态过滤数据
		sqlStr = `select 
					lid,date,ldssnum,lgss,mateInfo,amount,paid,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id 
				  from ldss where status=?;`
		err = Db.Select(&ldss, sqlStr, stat)
		if err != nil {
			fmt.Println("============== 按照时间过滤查询数据库失败2 =====================", err)
			return nil, err
		}
	} else {
		sqlStr = `select 
					lid,date,ldssnum,lgss,mateInfo,amount,paid,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id
				  from 
					ldss 
				  where status=? and date between ? and ?;`
		err = Db.Select(&ldss, sqlStr, stat, startDate, endDate)
		if err != nil {
			fmt.Println("============== 按照时间过滤查询数据库失败3 =====================", err)
			return nil, err
		}
	}
	return
}

// UpdateLdssBadRad 后台通过维修状态更新问题原因及status为 2
func UpdateLdssBadRad(data *Data) bool {
	sqlStr := `update ldss set pro_ret=?,status=? where ldssnum=? and c_id=?;`
	fmt.Println("=====>>>>>>>>>>>>>>>>>>>>>", data.BadRea, data.Status, data.Ldss.LdssNum, data.ClientId)
	_, err := Db.Exec(sqlStr, data.BadRea, data.Status, data.Ldss.LdssNum, data.ClientId)
	if err != nil {
		fmt.Println("================ 维修处理更改状态信息失败 ================", err)
		return false
	}
	return true
}

// QueryNumber 按照 ldssnum 获取实收数量和已发数量信息
func QueryNumber(ldssnum, remateinfo string) (numsMap map[string]int, err error) {
	var nums []*Ldss
	sqlStr := `select amount,paid,sentnum from ldss where ldssnum=? and mateInfo=?;`
	err = Db.Select(&nums, sqlStr, ldssnum, remateinfo) // TODO 多传了最后一个参数
	if err != nil {
		fmt.Println("============== 按照 ldssnum 查询实收数量失败 ==============", err)
		return nil, err
	}
	numsMap = make(map[string]int, 0)
	for _, val := range nums {
		numsMap["amount"] = val.Amount
		numsMap["paid"] = val.Paid
		numsMap["sentnum"] = val.SentNum
	}
	return
}

// SelectAllClientInfo 查询所有客户信息
func SelectAllClientInfo() (clients []*ClientAddr, err error) {
	sqlStr := `select cid,username,phone,address from clientAddr;`
	err = Db.Select(&clients, sqlStr)
	if err != nil {
		fmt.Println("============ 查询所有客户信息出错 ===============", err)
		return nil, err
	}
	return
}

// LdssSltSaleName 根据快递单号查询销售人员名称
func LdssSltSaleName(ldssnum, mate string) (saleName string, err error) {
	sqlStr := `select saleName from sale where sid = (select sale_id from ldss where ldssnum=? and mateInfo=?);`
	err = Db.Get(&saleName, sqlStr, ldssnum, mate)
	if err != nil {
		fmt.Println("========== 推送微信消息时根据快递单号查询销售或业务人员错误：===========", err)
		return "", err
	}
	return
}

// GetRetInfo 查询发货表信息数据  getType string,
func GetRetInfo(ldssId int) (data []*RetInfo, err error) {
	query := `select 
						retDate,retLdsNum,retLdsName,retMate,retAmount,retMark,ldss_id 
				  from 
						retInfo 
				  where 
						ldss_id=?;`
	err = Db.Select(&data, query, ldssId)
	if err != nil {
		fmt.Println("========== 管理后台根据 ldssId 查询发货信息表数据错误 ：===========", err)
		return nil, err
	}
	//}
	return data, nil
}

// InsertInfoToShipTable 后台发货操作插入数据到 retInfo 表
func InsertInfoToShipTable(data *Data) error {
	query := `insert into 
				retInfo(retDate,retLdsNum,retLdsName,retMate,retAmount,retMark,ldss_id) 
				values(?,?,?,?,?,?,?);`
	_, err := Db.Exec(query, data.ReturnDate, data.RetLdsNum, data.RetLdsName, data.RetMate, data.RetAmount, data.RetMark, data.LdssId)
	if err != nil {
		return errors.New("发货操作，插入发货信息到 retInfo 表失败")
	}
	return nil
}

// GetLdssStatus 按照客户返回快递单号查询接收状态信息  TODO  新增
func GetLdssStatus(ldssNum, mate string) (status string, err error) {
	queryStr := `select status from ldss where ldssnum=? and mateInfo=?;`
	err = Db.Get(&status, queryStr, ldssNum, mate)
	if err != nil {
		fmt.Println("================ 接收操作时按照快递单号查询数据错误 ================", err)
		return "", nil
	}
	return
}

// GetWorkOrderInfo ...
func GetWorkOrderInfo(stat, cid int) (data []*Ldss, err error) {
	queryStr := ``
	switch stat {
	case 6:
		queryStr = `select 
       					lid,date,ldssnum,lgss,mateInfo,amount,paid,pstate,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id
					from ldss where c_id=?;`
		err = Db.Select(&data, queryStr, cid)
		if err != nil {
			fmt.Println("=============== 客户根据工单状态 id 查询所有工单信息错误 =======================", err)
			return nil, err
		}
	case 0:
		queryStr = `select 
       					lid,date,ldssnum,lgss,mateInfo,amount,paid,pstate,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id
					from 
						ldss 
					where status=? and c_id=?;`
		err = Db.Select(&data, queryStr, stat, cid)
		if err != nil {
			fmt.Println("=============== 客户根据工单状态 id 查询已发工单信息错误 =======================", err)
			return nil, err
		}
	case 1:
		queryStr = `select 
       					lid,date,ldssnum,lgss,mateInfo,amount,paid,pstate,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id
					from 
						ldss 
					where status=? and c_id=?;`
		err = Db.Select(&data, queryStr, stat, cid)
		if err != nil {
			fmt.Println("=============== 客户根据工单状态 id 查询已收工单信息错误 =======================", err)
			return nil, err
		}
	case 3:
		queryStr = `select 
       					lid,date,ldssnum,lgss,mateInfo,amount,paid,pstate,sentnum,bad_rea,status,remark,pro_ret,ret_date,ret_ldss,ret_ldssname,ret_remark,shipNum,sale_id,c_id
					from 
						ldss 
					where status=? and c_id=?;`
		err = Db.Select(&data, queryStr, stat, cid)
		if err != nil {
			fmt.Println("=============== 客户根据工单状态 id 查询已发工单信息错误 =======================", err)
			return nil, err
		}
	default:
		return nil, errors.New("暂不支持该类型操作！")
	}
	return
}

func QueryClientInfo(ldssnum, mate string) (data []*ClientAddr, err error) {
	queryStr := `
				select 
					cid,username,phone,address 
				from 
					clientAddr 
				where 
					cid = (select c_id from ldss where ldssnum=? and mateInfo=?);
				`
	err = Db.Select(&data, queryStr, ldssnum, mate)
	if err != nil {
		fmt.Println("=============== 推送微信消息时根据快递单号和物料名称查询客户信息失败 ============", err)
		return nil, err
	}
	return
}
