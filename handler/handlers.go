package handler

import (
	"afterSaleSys/db"
	"afterSaleSys/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	R *gin.Engine
)

func Routers() {
	R.GET("/login", login)
	R.POST("/login", login)
	R.GET("/", MyMiddleware, index)
	R.POST("/", MyMiddleware, index)
	R.PUT("/", MyMiddleware, index)
	R.GET("/add_info", addInfo)
	R.GET("/getWorkOrder", getWorkOrder)
	R.POST("/add_info", addInfo)

	R.GET("/admin", BackMiddleware, backAdmin)
	R.PUT("/admin", BackMiddleware, backAdmin)
	R.POST("/admin", BackMiddleware, backAdmin)
	R.POST("/admin/upldss", BackMiddleware, upLdss)
	R.POST("/adminLogin", adminLogin)
	R.GET("/adminLogin", adminLogin)
	R.GET("/logout", logOut)
	R.POST("/admin/paid_up", BackMiddleware, paidUpLdss)
	R.POST("/admin/paid_confirm", BackMiddleware, confirmPaid)
	R.POST("/admin/cli_info", selectClient)
	R.POST("/admin/dateFilterDate", BackMiddleware, dateFilterData)
	R.POST("/admin/upMaintain", BackMiddleware, upMaintainStat)
	R.GET("/admin/clients", BackMiddleware, showClients)
	R.GET("/admin/saleInfo", BackMiddleware, showSaleInfo)
	R.GET("/admin/search", BackMiddleware, searchInfo)
	R.GET("/admin/error", BackMiddleware, showError)
	R.POST("/notice", BackMiddleware, noticeFunc)
	R.GET("/notice", BackMiddleware, noticeFunc)
	R.GET("/admin/getRetInfo", BackMiddleware, GetRetInfo)
}

var wg sync.WaitGroup

// MyMiddleware 客户登录中间件
func MyMiddleware(c *gin.Context) {
	phone := Init(c).GetSe("phone")
	if phone == "" {
		c.Request.URL.Path = "/login"
		R.HandleContext(c)
	}
}

// BackMiddleware 后台操作中间件
func BackMiddleware(c *gin.Context) {
	uid := Init(c).GetSe("uid")
	if uid == "" {
		c.Request.URL.Path = "/adminLogin"
		R.HandleContext(c)
	}
	uType := db.QueryUserType(uid)
	if utype, err := strconv.Atoi(uType); err != nil {
		fmt.Println("================= 后台登录中间件中转换用户已类型错误 =================", err)
		return
	} else if utype == 2 {
		c.Redirect(301, "/return/admin")
		return
	}
	return
}

// login 登录方法
func login(c *gin.Context) {
	if c.Request.Method == "POST" {
		phone := c.PostForm("phone")
		phoneReg := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		reg := regexp.MustCompile(phoneReg)
		if !reg.MatchString(phone) {
			c.JSON(200, map[string]string{"errForMsg": "手机号码格式不正确，请检查后重试!"})
			return
		}

		data, err := db.QueryPhone(phone)
		if err != nil || len(data) < 1 {
			c.JSON(200, map[string]string{"errMsg": "此号码不存在，请先注册!"})
			return
		}
		if data[0].Phone == phone {
			setBool := Init(c).SetSe("phone", phone)
			if !setBool {
				c.JSON(200, map[string]string{"sqlErrMsg": "后台服务错误，请稍后重试....."})
				return
			}
			setBool = Init(c).SetSe("clientId", strconv.Itoa(data[0].Cid))
			if !setBool {
				c.JSON(200, map[string]string{"sqlErrMsg": "后台服务错误，请稍后重试....."})
				return
			}
			//c.Redirect(301, "/")
			c.JSON(200, map[string]string{"success": "登录成功"})
			//c.HTML(301, "template/index.html",map[string]string{"success": "登录成功"})
			return
		} else {
			c.JSON(200, map[string]string{"error": "后台服务错误，请稍后再试!"})
			return
		}
	}
	c.HTML(200, "template/login.html", gin.H{"title": "链力电子"})
}

// 退出登录
func logOut(c *gin.Context) {
	bol := Init(c).DelSe("uid")
	bol = Init(c).DelSe("utype")
	if !bol {
		c.JSON(200, map[string]string{"errMsg": "退出失败！"})
		return
	}
}

// addInfo 注册方法
func addInfo(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		if username == "" {
			c.JSON(200, map[string]string{"usernameErrMsg": "姓名不允许为空"})
			return
		}
		phone := c.PostForm("phone")
		phoneReg := "^((13[0-9])|(14[3,5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		reg := regexp.MustCompile(phoneReg)
		if !reg.MatchString(phone) {
			c.JSON(200, map[string]string{"phoneErrMsg": "手机号码格式不正确，请检查后重试!"})
			return
		}
		dataSli, err := db.QueryPhone(phone)
		if err != nil {
			c.JSON(200, map[string]string{"phoneErrMsg": "后台数据服务错误，请联系相关人员处理后重试，给您带来的不便，敬请谅解！"})
			return
		}
		if len(dataSli) > 0 && dataSli[0].Phone != "" {
			c.JSON(200, map[string]string{"phoneErrMsg": "此号码已存在，请换一个号码进行注册"})
			return
		}
		// 开始将注册信息写入数据库，完成注册操作
		data := &db.Data{ClientAddr: &db.ClientAddr{UserName: username, Phone: phone}}
		bl, err := db.InsertToClientTable(data)
		if err != nil || !bl {
			c.JSON(200, map[string]string{"sqlErrMsg": "后台错误，请联系相关人员并稍后重试....."})
			return
		}
		setBool := Init(c).SetSe("phone", phone)
		setBool = Init(c).SetSe("name", username)
		if !setBool {
			c.JSON(200, map[string]string{"sqlErrMsg": "后台错误，请联系相关人员并稍后重试....."})
			return
		}
		c.JSON(200, map[string]string{"success": "添加成功"})
		return
	}
	c.HTML(200, "template/addInfo.html", gin.H{"title": "链力电子"})
}

// 客户根据工单状态获取工单信息
func getWorkOrder(c *gin.Context) {
	ret := c.Query("type")
	cid := Init(c).GetSe("clientId")
	stat, err := strconv.Atoi(ret)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "数据类型错误，请联系相关人员处理后重试，给您带来的不便敬请谅解！"})
		return
	}
	c_id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "数据类型错误，请联系相关人员处理后重试，给您带来的不便敬请谅解！"})
		return
	}
	dataSli, err := db.GetWorkOrderInfo(stat, c_id)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "后台数据服务操作错误，请联系相关人员处理后重试，给您带来的不便敬请谅解！"})
		return
	}
	if len(dataSli) > 0 {
		c.JSON(200, map[string]interface{}{"success": dataSli})
	} else {
		c.JSON(200, map[string]interface{}{"success": "null"})
	}
}

// index 主页实现逻辑
func index(c *gin.Context) {
	phone := Init(c).GetSe("phone")
	clients, err := db.QueryClient(phone)
	if err != nil {
		c.JSON(200, gin.H{})
		return
	}
	var cid int
	for _, val := range clients {
		if val.Phone == phone {
			cid = val.Cid
		}
	}
	sales, err := db.QuerySale()
	if err != nil {
		c.JSON(200, gin.H{})
		return
	}
	// 查询 ldss 中我的工单信息数据
	ldsInfo, err := db.QueryLdssInfo(cid)
	if err != nil {
		c.JSON(200, gin.H{})
		return
	}
	// 查询实际接收获取信息数据
	var pstate string
	ldsPaidData, err := db.QueryPaidStat(cid)
	if err != nil {
		c.JSON(200, gin.H{})
		return
	}
	for _, val := range ldsPaidData {
		pstate = strconv.Itoa(val.Pstate)
	}
	// 客户确认收货更新 status 字段信息
	if c.Request.Method == "PUT" {
		ldssNum := c.PostForm("ldssNum")
		data := &db.Data{Ldss: &db.Ldss{
			LdssNum:  ldssNum,
			ClientId: cid,
			Status:   3,
		}}
		err := db.UpLdssData("upStatus", data)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "确认收货失败，建议联系售后人员处理后重试"})
			return
		}
		c.JSON(200, map[string]string{"success": "确认收货成功..."})
		return
	}
	// 提交物料快递及客户信息
	if c.Request.Method == "POST" {
		timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 03:04:05")
		lgsNum := c.PostForm("lgsNum")
		lgsName := c.PostForm("lgsName")
		mateName := c.PostForm("mateName")
		mateNum := c.PostForm("mateNum")
		amount, err := strconv.Atoi(mateNum)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "输入的物品数量数据类型不匹配，请检查后重试"})
			return
		}
		badRea := c.PostForm("badRea")
		saleId := c.PostForm("saleId")
		s_id, err := strconv.Atoi(saleId)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "选择的销售员数据类型不匹配，请检查后重试"})
			return
		}
		returnId := c.PostForm("clientId")
		c_id, err := strconv.Atoi(returnId)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "数据类型不匹配，请检查后重试"})
			return
		}
		returnAddr := c.PostForm("clientAddr")
		ldsData := &db.Data{Ldss: &db.Ldss{
			Amount:   amount,   // 物料数量
			Date:     timeStr,  // 时间信息
			LdssNum:  lgsNum,   // 物流单号
			LgssName: lgsName,  // 物流名称
			MateInfo: mateName, // 物料信息
			BadRea:   badRea,   // 不良原因
			SaleId:   s_id,     // 关联的销售人员
			ClientId: c_id,     // 关联的客户信息
		}}
		// 插入数据
		err = db.InsertToLdsTable(ldsData)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "后台数据服务错误，带来的不便敬请谅解，建议稍后重试"})
			return
		}

		clientData := &db.Data{
			ClientAddr: &db.ClientAddr{Cid: c_id, Address: &returnAddr},
		}
		err = db.UpLdssData("client", clientData)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "后台数据服务错误，带来的不便敬请谅解，建议稍后重试"})
			return
		}
		c.JSON(200, map[string]string{"success": "提交成功,可在菜单栏选择进入我的工单查看该进展信息"})
		return
	}
	c.HTML(200, "template/index.html", gin.H{"title": "链力电子",
		"sales": sales, "client": clients, "ldss": ldsInfo, "pstate": pstate, "paidInfo": ldsPaidData})
}

// backAdmin 后台管理页面逻辑
func backAdmin(c *gin.Context) {
	uid := Init(c).GetSe("utype")
	var clientList []*db.ClientAddr
	var statList []interface{}
	// 后台管理页面查询 ldss 表中所有数据
	ldssArry, err := db.QueryLdssInfo(0)
	if err != nil {
		c.JSON(200, map[string]string{})
		fmt.Println("查询ldss数据库错误")
		return
	}
	// 循环所有工单，根据状态查找出未完成的工单，用来实现通知功能
	for _, val := range ldssArry {
		clientList, err = db.QueryClientName(val.ClientId)
		if err != nil {
			fmt.Println("查询clientAddr数据库错误", err)
			return
		}
		if val.Status == 0 || val.Status == 1 || val.Status == 2 {
			statList = append(statList, val.Status)
		}
	}
	// ======================= 分页功能实现 ======================
	//pageSize := 15                                          // 每页显示的条数
	//pageCount := float64(len(ldssArry)) / float64(pageSize) // 总页数
	c.HTML(200, "template/admin.html", gin.H{"title": "链力电子", "ldssInfo": ldssArry,
		"cname": clientList, "status": len(statList), "uid": uid})
}

// GetRetInfo 根据发货详情获取发货数据信息
func GetRetInfo(c *gin.Context) {
	ldssId := c.Query("ldssId")
	id, err := strconv.Atoi(ldssId)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "抱歉!程序故障，转换数据类型出错或数据格式不正确，请联系技术解决后重试！"})
		return
	}
	retInfo, err := db.GetRetInfo(id)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "抱歉!程序故障，后台操作数据库错误，请联系技术解决后重试！"})
		return
	}
	if len(retInfo) > 0 {
		c.JSON(200, map[string]interface{}{"success": retInfo})
	} else {
		c.JSON(200, map[string]interface{}{"success": "retInfo"})
	}
}

// searchInfo 搜索方法功能实现
func searchInfo(c *gin.Context) {
	ldssnum, bol := c.GetQuery("ldssnum")
	if !bol {
		c.JSON(200, map[string]string{"errMsg": "获取参数错误"})
		return
	}
	ldssArry, err := db.SelectLdssData(ldssnum)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "后台数据服务错误，请联系技术人员解决!"})
		return
	}
	c.HTML(200, "template/search.html", gin.H{"title": "链力电子", "ldss": ldssArry})
}

// noticeFunc 展示通知信息方法
func noticeFunc(c *gin.Context) {
	cid, b := c.GetQuery("cid")
	if !b {
		c.JSON(200, map[string]string{"errMsg": "客户不存在"})
		return
	}
	c_id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(200, map[string]string{})
	}
	ldssList, err := db.QueryLdssInfo(c_id)
	if err != nil {
		c.JSON(200, map[string]string{})
	}
	c.HTML(200, "template/notice.html", gin.H{"title": "链力电子", "ldss": ldssList})
}

// upLdss 发货后更新 ldss 表 status 字段以及将数据插入到 retInfo 表
func upLdss(c *gin.Context) {
	var status, shipnum int
	oldldssnum := c.PostForm("oldldssnum")
	newldssnum := c.PostForm("newldssnum")
	ldssname := c.PostForm("ldssname")
	sentNum := c.PostForm("sentNum")
	ldssTableId := c.PostForm("ldssTableId")
	ldssId, err := strconv.Atoi(ldssTableId)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "数据转换错误。。。。。。"})
		return
	}
	snum, err := strconv.Atoi(sentNum)
	if err != nil {
		c.JSON(200, map[string]string{})
		return
	}
	remark := c.PostForm("remark")
	remateInfo := c.PostForm("remateinfo")
	// todo 更新之前先查找一遍防止相同的快递单号重复插入
	retSli, err := db.GetRetInfo(ldssId)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "后台查找 retInfo 表数据错误，请联系技术人员进行处理后再试"})
		return
	}
	for _, value := range retSli {
		if value.RetLdsNum == newldssnum {
			c.HTML(301, "template/error.html", gin.H{"errMsg": "重复！此次发货的快递单号以存在发货数据，请检查后重试!"})
			return
		}
	}

	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 03:04:05")
	//发货后现将发货信息插入到 retInfo 表中，然后通过 ldssTableId 相关数据，并统计数量，并将统计的数量插入到 ldss 表中的 shipNum 字段
	// 插入发货信息到 retInfo 表
	shipData := &db.Data{
		RetInfo: &db.RetInfo{
			RetLdsNum:  newldssnum,
			RetLdsName: ldssname,
			ReturnDate: timeStr,
			RetMate:    remateInfo,
			RetAmount:  snum,
			RetMark:    remark,
			LdssId:     ldssId,
		},
	}
	err = db.InsertInfoToShipTable(shipData)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "后台插入数据到 retInfo 表错误，请联系技术人员进行处理后再试"})
		return
	}
	// 获取发货信息表数据，然后对其求长度，最后将长度更新到 ldss 表的 shipNum 字段上表示本工单发了几次货
	retInfoSli, err := db.GetRetInfo(ldssId)
	shipnum = len(retInfoSli)
	// 查询实收数量
	numMap, err := db.QueryNumber(oldldssnum, remateInfo)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "后台查询数据错误，请联系技术人员进行处理后再试"})
		return
	}
	// 如发出的数量与实收的数量不相等则先查询已返数量，并进行相加后更新到以返回数量上
	if numMap["paid"] != snum {
		snum += numMap["sentnum"]
		if snum > numMap["paid"] {
			c.HTML(301, "template/error.html", gin.H{"errMsg": "发出的数量不能大于实收数量，请检查后重试"})
			return
		} else if snum == numMap["amount"] { // 如发出数量等于客户客户发出数量则更新状态为 3
			status = 3
		}
	}
	// 根据发货数量与实收数量更新status的值
	if (numMap["paid"] - snum) == 0 {
		status = 3
	} else if (numMap["paid"] - snum) < 0 {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "非法操作，发出的数量大于实收数量"})
		return
	} else {
		status = 2
	}
	data := &db.Data{Ldss: &db.Ldss{
		RetDate:     timeStr,
		RetLdss:     newldssnum,
		RetLdssName: ldssname,
		RetRemark:   remark,
		SentNum:     snum,
		Status:      status,
		MateInfo:    remateInfo,
		ShipNum:     &shipnum,
	}}
	bl := db.UpdateStatus("ldss", oldldssnum, data)
	if !bl {
		c.JSON(200, map[string]string{})
		return
	}
	// 推送微信前,根据快递单号查询客户信息
	dataSli, err := db.QueryClientInfo(oldldssnum, remateInfo)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "后台数据库操作错误,请联系技术解决后重试"})
		return
	}
	msg := "**Hi,我是小链！您有一份新的售后工单已发出返回给您的客户，返回信息如下：**\n" +
		">" +
		"工单编号: <font color='info'>" + ldssTableId + "</font>\n" +
		"发货单号: <font color='info'>" + newldssnum + "</font>\n" +
		"发出商品: <font color='info'>" + remateInfo + "</font>\n" +
		"发货备注: <font color='info'>" + remark + "</font>\n" +
		"发出数量: <font color='info'>" + sentNum + "</font>\n" +
		"\n" +
		"如有疑问可通过下面的客户信息联系客户:\n" +
		">" +
		"客户姓名: <font color='info'>" + dataSli[0].UserName + "</font>\n" +
		"联系方式: <font color='info'>" + dataSli[0].Phone + "</font>\n" +
		"客户地址: <font color='info'>" + *dataSli[0].Address + "</font>\n" +
		"\n" +
		"**请您实时留意快递信息有疑问及时联系客户，公司的发展离不开您，感谢您的理解和付出，谢谢！**"
	wg.Add(1)
	go sendToWechat(oldldssnum, remateInfo, &msg)
	wg.Wait()
	c.Redirect(301, "/admin")
}

// sendToWechat 推动消息到企业微信方法
func sendToWechat(ldssnum, mate string, msg *string) {
	defer wg.Done()
	err := pushMsgToWechat(ldssnum, mate, msg)
	if err != nil {
		fmt.Println("售后后台接收操作调用 PushMsgToWechat 方法推送微信消息失败：", err)
		return
	}
}

// paidUpLdss 接收操作更新 ldss 表 pstate、status 方法
func paidUpLdss(c *gin.Context) {
	var status int
	ldssnum := c.PostForm("ldssnum")
	mateinfo := c.PostForm("mateinfo")
	amount := c.PostForm("amount")
	beadInfo := c.PostForm("bead_info")
	TableId := c.PostForm("TableId")
	anum, err := strconv.Atoi(amount)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "提交的实收数量数据格式有误,或有必填项未填写，检查后重试"})
		return
	}
	paid := c.PostForm("paid")
	remark := c.PostForm("remark")
	pad, err := strconv.Atoi(paid)
	if ldssnum == "" || mateinfo == "" || paid == "" || err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "提交的实收数据有误,或有必填项未填写，检查后重试"})
		return
	}
	// TODO 接收前先查询该单的状态是否以接收过一次，接收过则返回提示信息
	state, err := db.GetLdssStatus(ldssnum, mateinfo)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "数据查找错误，请联系技术解决后重试!"})
		return
	}
	if state == "1" {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "错误，请勿重复接收！"})
		return
	}
	// TODO  ====================================
	// 对接收的数量与客发的数量作对比，大于或小于客发数量则返回错误提示
	if pad > anum {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "对不起，接收的数量不允许大于客户返回的数量!"})
		return
	} else if pad == anum {
		status = 1
	} else {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "接收的数量小于客户返回的数量，请联系客户确认后填写客返数量并在备注中写明原因!"})
		return
	}
	data := &db.Data{
		Ldss: &db.Ldss{
			LdssNum:  ldssnum,
			MateInfo: mateinfo,
			Paid:     pad,
			Pstate:   2,
			Status:   status,
			Remark:   remark,
		},
	}
	resBool := db.UpdateStatus("paid", ldssnum, data)
	if !resBool {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "后台数据库操作错误,请联系技术解决后重试"})
		return
	}
	// 推送微信前,根据快递单号查询客户信息
	dataSli, err := db.QueryClientInfo(ldssnum, mateinfo)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "后台数据库操作错误,请联系技术解决后重试"})
		return
	}
	// TODO 更新数据库成功后推送微信消息给销售，已完成
	// 初始化消息数据
	msg := "**Hi,我是小链！您有一份新的售后工单已到售后部门，需要您进行跟踪处理。信息如下：**\n" +
		">" +
		"工单编号: <font color='info'>" + TableId + "</font>\n" +
		"快递单号: <font color='info'>" + ldssnum + "</font>\n" +
		"商品信息: <font color='info'>" + mateinfo + "</font>\n" +
		"问题原因: <font color='info'>" + beadInfo + "</font>\n" +
		"返回数量: <font color='info'>" + amount + "</font>\n" +
		"\n" +
		"如有疑问可通过下面的客户信息联系客户:\n" +
		">" +
		"客户姓名: <font color='info'>" + dataSli[0].UserName + "</font>\n" +
		"联系方式: <font color='info'>" + dataSli[0].Phone + "</font>\n" +
		"客户地址: <font color='info'>" + *dataSli[0].Address + "</font>\n" +
		"\n" +
		"**还请您百忙之中抽出一点时间尽快处理，公司的发展离不开您，感谢您的理解和付出，谢谢！**"
	wg.Add(1)
	go sendToWechat(ldssnum, mateinfo, &msg)
	wg.Wait()
	c.Redirect(301, "/admin")
}

// confirmPaid 客户确认实收数量后更新实收数量状态及 status 状态 TODO 可以删除该功能
func confirmPaid(c *gin.Context) {
	clientId := c.PostForm("clientId")
	cid, err := strconv.Atoi(clientId)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "系统错误，建议稍后重试"})
		return
	}
	data := &db.Data{
		Ldss: &db.Ldss{
			ClientId: cid,
			Pstate:   2,
			Status:   1,
		},
	}
	bol := db.UpPaidStatus(data)
	if !bol {
		c.JSON(200, map[string]string{"errMsg": "系统错误，建议稍后重试"})
		return
	}
	c.Redirect(301, "/")
}

// 根据客户id 查询客户信息
func selectClient(c *gin.Context) {
	clientId := c.PostForm("clientId")
	cid, err := strconv.Atoi(clientId)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "提交的客户ID数据类型不匹配,检查后重试"})
		return
	}
	clientList, err := db.SelectClient(cid)
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "后台数据服务错误,请联系相关技术解决后再试"})
		return
	}
	var clientMap = make(map[string]interface{}, 0)
	for _, val := range clientList {
		clientMap["cid"] = val.Cid
		clientMap["username"] = val.UserName
		clientMap["phone"] = val.Phone
		clientMap["address"] = val.Address
	}
	c.JSON(200, map[string]interface{}{"success": clientMap})
	return
}

// 后台管理登录页面
func adminLogin(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		passwd := c.PostForm("passwd")
		hashPasswd := util.HashPasswd(passwd)
		userInfo, err := db.BackLoginSelect(username)
		if err != nil {
			c.JSON(200, map[string]string{"errMsg": "后台数据服务错误，稍后重试"})
			return
		}
		for _, val := range userInfo {
			if hashPasswd != val.Passwd {
				c.JSON(200, map[string]string{"errMsg": "密码错误，请检查后重试重试"})
				return
			}
			Init(c).SetSe("uid", strconv.Itoa(val.Uid))
			Init(c).SetSe("utype", strconv.Itoa(val.Utype))
		}
		c.JSON(200, map[string]string{"success": "登录成功"})
		return
	}
	c.HTML(200, "template/adminLogin.html", gin.H{"title": "链力电子"})
}

// 按时间过滤后台功能逻辑实现
func dateFilterData(c *gin.Context) {
	// 按时间过滤功能实现.........
	if c.Request.Method == "POST" {
		startDate := c.PostForm("startDate")
		enDate := c.PostForm("endDate")
		choiceStat := c.PostForm("choiceStat")
		stat, err := strconv.Atoi(choiceStat)
		if err != nil {
			c.HTML(301, "template/error.html", gin.H{"errMsg": "状态数据类型错误，请稍后后重试"})
			return
		}
		ldssLst, err := db.SelectDateLdss(startDate, enDate, stat)
		if err != nil {
			c.HTML(30, "template/error.html", gin.H{"errMsg": "后台数据服务错误，请联系技术进行处理"})
			return
		}
		c.HTML(200, "template/date_filter_ldss.html", gin.H{"ldssLst": ldssLst, "title": "链力电子"})
		return
	}
}

// upMaintainStat 通过维修按钮修改status字段信息
func upMaintainStat(c *gin.Context) {
	ldssnum := c.PostForm("ldssnum")
	badRea := c.PostForm("badRea")
	clientId := c.PostForm("clientId")
	mateName := c.PostForm("mateName")
	cid, err := strconv.Atoi(clientId)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "客户ID数据类型不匹配，检查后重试"})
		return
	}
	// TODO 接收前先查询次单的状态是否以接收过一次，接收过则返回提示信息
	state, err := db.GetLdssStatus(ldssnum, mateName)
	if err != nil {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "数据查找错误，请联系技术解决后重试!"})
		return
	}
	if state == "2" {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "错误，请勿重复提交！"})
		return
	}
	// TODO  ====================================
	data := &db.Data{
		Ldss: &db.Ldss{
			LdssNum:  ldssnum,
			ClientId: cid,
			BadRea:   badRea,
			Status:   2,
		},
	}
	bl := db.UpdateLdssBadRad(data)
	if !bl {
		c.HTML(301, "template/error.html", gin.H{"errMsg": "后台服务错误，请联系相关技术进行处理"})
		return
	}
	c.Redirect(301, "/admin")
}

// 查看所有客户信息
func showClients(c *gin.Context) {
	clients, err := db.SelectAllClientInfo()
	if err != nil {
		c.JSON(200, map[string]string{"errMsg": "502！后台数据服务错误，请联系技术人员解决后重试"})
		return
	}
	c.HTML(200, "template/showClients.html", gin.H{"title": "链力电子", "clients": clients})
}

// showSaleInfo 根据销售Id查看销售员信息
func showSaleInfo(c *gin.Context) {
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

// showError 错误页面handler
func showError(c *gin.Context) {
	c.HTML(200, "template/error.html", gin.H{})
}
