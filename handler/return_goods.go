package handler

import (
	"afterSaleSys/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func RetRouters() {
	r := R.Group("/return")
	r.GET("/admin", backMiddleware, returnIndex)
	r.POST("/retGoods", backMiddleware, returnGoods)
	r.GET("/admin/saleInfo", backMiddleware, selectSaleInfo)
	r.GET("/admin/search", backMiddleware, searchMsg)
	r.GET("admin/clients", backMiddleware, LookClientInfo)
	r.POST("/admin/paid_up", backMiddleware, recvUpData)
	r.POST("/admin/upMaintain", backMiddleware, upRetDealInfo)
	r.POST("/admin/upldss", backMiddleware, upSendDataInfo)
	r.POST("/admin/dateFilterDate", backMiddleware, dateFilterReturn)
}

// BackMiddleware 后台操作中间件
func backMiddleware(c *gin.Context) {
	uid := Init(c).GetSe("uid")
	if uid == "" {
		c.Request.URL.Path = "/adminLogin"
		R.HandleContext(c)
	}
	uType := db.QueryUserType(uid)
	if utype, err := strconv.Atoi(uType); err != nil {
		fmt.Println("================= 后台登录中间件中转换用户已类型错误 =================", err)
		return
	} else if utype == 1 {
		c.Redirect(301, "/admin")
		return
	}
	return
}

// 退货后台管理页面方法
func returnIndex(c *gin.Context) {
	//var saleSlice = make([]string, 0)
	goods, err := db.QuertRetAdminDate("index", nil)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "后台数据库查询错误，请联系技术解决"})
		return
	}
	uid := Init(c).GetSe("utype")
	//for _, val := range goods {
	//	if val.Stat == 2 {
	//		// TODO 1、通过 ssid 查询销售员，后将其添加到 saleMap map中,key 为 sid value 为 saleName
	//		saleAry, err := db.QuerySaleName(val.Ssid)
	//		if err != nil {
	//			c.JSON(200, map[string]string{"": ""})
	//			return
	//		}
	//		for _, vall := range saleAry {
	//			saleSlice = append(saleSlice, vall.SaleName)
	//		}
	//		// TODO 2. 循环遍历 saleMap 并使用多进程调用企业微信后台接口处理方法发送通知给销售员
	//	}
	//}
	c.HTML(200, "template/return_admin.html", gin.H{"title": "链力电子", "retGoods": goods, "uid": uid})
}

// 退货后台客户提交退货工单处理方法
func returnGoods(c *gin.Context) {
	fmt.Println("=======================================>")
	//if c.Request.Method == "POST" {
	ldssNum := c.PostForm("ldss_num")
	if len(ldssNum) > 16 {
		c.JSON(200, map[string]string{"errMsg": "快递单号超出长度限制，最长为16位，请检查后重试！"})
		return
	}
	ldssName := c.PostForm("ldss_name")
	mateName := c.PostForm("mate_name")
	reNum := c.PostForm("renum")
	renum, err := strconv.Atoi(reNum)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "退货数量数据类型不匹配，请检查后重试！"})
		return
	}
	phone := Init(c).GetSe("phone")
	// TODO 根据手机号查询客户 id
	clients, err := db.QueryClient(phone)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "后台操作数据服务器错误，请联系我司技术人员后再试！"})
		return
	}
	var cid int
	for _, val := range clients {
		if val.Phone == phone {
			cid = val.Cid
		}
	}
	badRea := c.PostForm("badrea")
	saleId := c.PostForm("sale_id")
	sid, err := strconv.Atoi(saleId)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "销售人员ID数据类型不匹配，请检查后重试！"})
		return
	}
	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 03:04:05")
	data := &db.Data{ReturnGoods: &db.ReturnGoods{
		LdsNum:   ldssNum,
		LdsName:  ldssName,
		MateName: mateName,
		Renum:    &renum,
		Reason:   badRea,
		Stat:     0,
		CrtDate:  timeStr,
		Ssid:     sid,
		Ccid:     cid,
	}}
	bol := db.InsertGoods(data)
	if !bol {
		c.JSON(200, map[string]string{"errMsg": "后台数据服务错误，请联系对接人员处理后重试！"})
		return
	}
	c.JSON(200, map[string]string{"success": "success"})
	//}
}

// 退货后台后台收货更新数据库操作
func recvUpData(c *gin.Context) {
	var state int
	ldssnum := c.PostForm("ldssnum")
	mateName := c.PostForm("matename")
	paid := c.PostForm("paid")
	paidNum, err := strconv.Atoi(paid)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "实收数量数据类型不对，件检查后重试!"})
		return
	}
	data := &db.Data{
		ReturnGoods: &db.ReturnGoods{
			LdsNum:   ldssnum,
			MateName: mateName,
		},
	}
	// TODO 先按照快递单号和物料名称查询出客发数量，然后将实收数量与客发数量进行对比
	// 如果接收的数量小于客发数量，第一次则提示非法操作，count 等于 2 表示第二次提交则放行更新到数据库
	if goods, err := db.QuertRetAdminDate("recv", data); err != nil {
		c.JSON(200, map[string]string{"errMsg": "查询数据库错误，请联系技术人员解决后再试!"})
		return
	} else {
		for _, val := range goods {
			if paidNum > *val.Renum {
				c.JSON(200, map[string]string{"errMsg": "非法操作，实收数量不允许大于客户返回数量"})
				state = 0
				return
			} else if paidNum < *val.Renum {
				paidNum += val.Paid
				if paidNum > *val.Renum {
					c.JSON(200, map[string]string{"errMsg": "非法操作，多次接收后的数量大于客户返回数量，请检查后重试!"})
					paidNum -= val.Paid
					return
				} else if paidNum < *val.Renum {
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>", paidNum)
					state = 0
				} else {
					fmt.Println("======================>", paidNum)
					state = 1
				}
			} else {
				fmt.Println(paidNum, "<======================")
				state = 1
			}
		}
	}
	data = &db.Data{
		ReturnGoods: &db.ReturnGoods{
			LdsNum:   ldssnum,
			MateName: mateName,
			Stat:     state,
			Paid:     paidNum,
		},
	}
	// TODO 更新到数据库
	bol := db.UpdateGoods("recv", "", data)
	if !bol {
		c.JSON(200, map[string]string{"errMsg": "后台数据库操作错误，请联系技术人员解决后再试!"})
		return
	}
	// todo 更新数据库成功后推送微信消息给售后
	c.JSON(200, map[string]string{"success": "success"})
}

// 退货后台点击销售id查询销售信息
func selectSaleInfo(c *gin.Context) {
	saleId := c.Query("saleId")
	sid, err := strconv.Atoi(saleId)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "数据类型不匹配，请检查后重试"})
		return
	}
	saleList, err := db.QuerySaleName(sid)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "后台数据错误，请联系技术人员进行处理"})
		return
	}
	var saleMap = make(map[string]interface{}, 0)
	for _, v := range saleList {
		saleMap["sid"] = v.Sid
		saleMap["salename"] = v.SaleName
		saleMap["phone"] = v.Phone
	}
	c.JSON(200, map[string]interface{}{"success": saleMap})
}

// 退货后台销售跟踪处理后更新处理信息
// 需要更新的信息为处理方法和处理时间，以及根据选择的是否换新情况更新状态信息，如不换新则直接更新状态为 3，换新则更新状态为2
func upRetDealInfo(c *gin.Context) {
	var state int
	if c.Request.Method == "POST" {
		ldsNum := c.PostForm("ldssnum")
		ldsName := c.PostForm("ldssName")
		mateName := c.PostForm("mateName")
		reNum := c.PostForm("amount")
		reCount, err := strconv.Atoi(reNum)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "返回数量数据类型不匹配，检查后重试!"})
			return
		}
		pNum := c.PostForm("pmount")
		paidCount, err := strconv.Atoi(pNum)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "实收数量数据类型不匹配，检查后重试!"})
			return
		}
		method := c.PostForm("badRea")
		clientId := c.PostForm("clientId")
		cid, err := strconv.Atoi(clientId)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "实收数量数据类型不匹配，检查后重试!"})
			return
		}
		choice := c.PostForm("choice")
		if choice == "2" {
			state = 2 // 等于 2 表示该退货物料需要换新，则显示发货按钮
		} else {
			state = 3 // 等于 3 表示该工单以完成
		}
		timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 03:04:05")
		data := &db.Data{
			ReturnGoods: &db.ReturnGoods{
				LdsNum:   ldsNum,
				LdsName:  ldsName,
				MateName: mateName,
				Renum:    &reCount,
				Paid:     paidCount,
				Method:   &method,
				Stat:     state,
				DealDate: &timeStr,
				Ccid:     cid,
			},
		}
		bol := db.UpdateGoods("maintain", "", data)
		if !bol {
			c.JSON(200, map[string]string{"errMsg": "后台数据库操作错误，请联系技术人员解决后再试！"})
			return
		}
		c.JSON(200, map[string]string{"success": "success"})
	}
}

// 退货后台点击发货按钮处理方法
func upSendDataInfo(c *gin.Context) {
	var state int
	// 发货后按照快递单号和处理方法更新发货数量和状态为 3
	oldldsNum := c.PostForm("oldldssnum")
	newLdsNum := c.PostForm("newldssnum")
	ldsName := c.PostForm("ldssname")
	ret := c.PostForm("result")
	sendNum := c.PostForm("sendNum")
	sNum, err := strconv.Atoi(sendNum)
	if err != nil {
		c.JSON(20, map[string]string{"errMsg": "发出数量数据类型错误，请检查后重试"})
		return
	}
	spread := c.PostForm("spread")
	// TODO 先查找已换数量，然后进行累加，如小于或等于实收数量则正常发货更改状态为 3，如大于实收数量则返回错误提示
	data := &db.Data{
		ReturnGoods: &db.ReturnGoods{
			LdsNum: oldldsNum,
			Method: &ret,
		},
	}
	if goods, err := db.QuertRetAdminDate("send", data); err != nil {
		c.JSON(20, map[string]string{"errMsg": "后台查询数据库错误，请联系技术人员解决后再试!"})
		return
	} else {
		for _, val := range goods {
			if sNum < val.Paid {
				sNum += *val.SendNum
				if sNum > val.Paid {
					c.JSON(20, map[string]string{"errMsg": "非法操作，换件的数量不允许大于实收数量!"})
					return
				} else if sNum < val.Paid && spread == "2" {
					state = 3 //TODO 是否允许换货商品的情况？,允许则允许少发更新状态为 3,否则更新为 2
				} else if sNum < val.Paid && spread == "1" {
					state = 2 // TODO 如果没有换货的情况则更新为2
				}
			} else if sNum > val.Paid {
				c.JSON(20, map[string]string{"errMsg": "非法操作，换件数量不允许大于实收数量!"})
				return
			} else {
				state = 3
			}
		}
	}
	data = &db.Data{ReturnGoods: &db.ReturnGoods{
		LdsNum:  newLdsNum,
		LdsName: ldsName,
		Method:  &ret,
		Stat:    state,
		SendNum: &sNum,
	}}
	bol := db.UpdateGoods("send", oldldsNum, data)
	if !bol {
		c.JSON(20, map[string]string{"errMsg": "后台更新数据库错误，请联系技术人员解决后再试!"})
		return
	}
	c.JSON(200, map[string]string{"success": "success"})
}

// 搜索内容后台实现
func searchMsg(c *gin.Context) {
	ldssnum, bol := c.GetQuery("ldssnum")
	if !bol {
		c.JSON(200, map[string]string{"errMsg": "获取参数错误"})
		return
	}
	ldssArry, err := db.SelectReturnData(ldssnum)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "后台数据服务错误，请联系技术人员解决!"})
		return
	}
	c.HTML(200, "template/return_search.html", gin.H{"title": "链力电子", "retGoods": ldssArry})
}

// 按时间过滤后台实现
func dateFilterReturn(c *gin.Context) {
	startDate := c.PostForm("startDate")
	enDate := c.PostForm("endDate")
	choiceStat := c.PostForm("choiceStat")
	//stat, err := strconv.Atoi(choiceStat)
	//if err != nil {
	//	c.HTML(301, "template/error.html", gin.H{"errMsg": "状态数据类型错误，请稍后后重试"})
	//	return
	//}
	retGoods, err := db.SelectDateFilterData(startDate, enDate, choiceStat)
	if err != nil {
		c.HTML(30, "template/error.html", gin.H{"errMsg": "后台数据服务错误，请联系技术进行处理"})
		return
	}
	c.HTML(200, "template/date_filter_return.html", gin.H{"retGoods": retGoods, "title": "链力电子"})
	return
}

// LookClientInfo 查看客户信息
func LookClientInfo(c *gin.Context) {
	clients, err := db.SelectAllClientInfo()
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "502！后台数据服务错误，请联系技术人员解决后重试"})
		return
	}
	c.HTML(200, "template/return_showClients.html", gin.H{"title": "链力电子", "clients": clients})
}
