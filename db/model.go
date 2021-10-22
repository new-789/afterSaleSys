package db

// 数据库对应结构体关系

type ClientAddr struct {
	Cid      int     `db:"cid" json:"cid"`
	UserName string  `db:"username" json:"username"`
	Phone    string  `db:"phone" json:"phone"`
	Address  *string `db:"address" json:"address"`
}

type Sale struct {
	Sid      int     `db:"sid" json:"sid"`
	SaleName string  `db:"saleName" json:"saleName"`
	Phone    *string `db:"phone" json:"phone"`
}

type Ldss struct {
	Lid         int    `db:"lid" json:"lid"`
	Status      int    `db:"status" json:"status"`   // 状态信息：状态1为未返2为以返
	Amount      int    `db:"amount" json:"amount"`   // 件数
	Paid        int    `db:"paid" json:"paid"`       // 实收数量
	Pstate      int    `db:"pstate" json:"pstate"`   // 实收数量状态,1为返回到客户,2为客户确定，并需要更新 status 状态,3为客户打回需电话确认
	SaleId      int    `db:"sale_id" json:"sale_id"` // 关联的销售人员
	ClientId    int    `db:"c_id" json:"c_id"`       // 关联的客户人员
	SentNum     int    `db:"sentnum" json:"sentnum"` // 已发货数量
	ShipNum     *int   `db:"shipNum" json:"shipNum"` // 发货的次数信息 TODO retInfo 表有数据后将 *int 改为 int
	Date        string `db:"date" json:"date"`
	LdssNum     string `db:"ldssnum" json:"ldssnum"`           // 物流单号
	LgssName    string `db:"lgss" json:"lgss"`                 // 物流名称
	MateInfo    string `db:"mateInfo" json:"mateInfo"`         // 物料信息
	BadRea      string `db:"bad_rea" json:"bad_rea"`           // 不良原因
	Remark      string `db:"remark" json:"remark"`             // 备注信息，用来记录收货详细信息
	RetDate     string `db:"ret_date" json:"ret_date"`         // 返回日期
	RetLdss     string `db:"ret_ldss" json:"ret_ldss"`         // 返回快递单号
	RetLdssName string `db:"ret_ldssname" json:"ret_ldssname"` // 返回快递公司名称
	RetRemark   string `db:"ret_remark" json:"ret_remark"`     // 返回备注信息
	ProRet      string `db:"pro_ret" json:"pro_ret"`           // 处理结果
}

type User struct {
	Uid      int    `db:"uid" json:"uid"`
	UserName string `db:"username" json:"username"`
	Phone    string `db:"phone" json:"phone"`
	Passwd   string `db:"passwd" json:"passwd"`
	Utype    int    `db:"utype" json:"utype"`
}

type ReturnGoods struct {
	Rid      int     `db:"rid" json:"rid"`           // id
	LdsNum   string  `db:"ldsnum" json:"ldsnum"`     // 快递单号
	LdsName  string  `db:"ldsname" json:"ldsname"`   // 快递公司
	MateName string  `db:"matename" json:"matename"` // 商品名称
	Reason   string  `db:"reason" json:"reason"`     // 退货原因
	Renum    *int    `db:"renum" json:"renum"`       // 退货数量
	Paid     int     `db:"paid" json:"paid"`         // 实收数量
	Method   *string `db:"method" json:"method"`     // 处理方式
	Stat     int     `db:"stat" json:"stat"`         // 状态
	CrtDate  string  `db:"crtdate" json:"crtdate"`   // 客户创建单据时间
	DealDate *string `db:"dealdate" json:"dealdate"` // 处理问题时间
	SendNum  *int    `db:"sendnum" json:"sendnum"`   // 返回数量
	Ssid     int     `db:"ssid" json:"ssid"`         // 关联的销售员
	Ccid     int     `db:"ccid" json:"ccid"`         // 关联的客户id
}

// RetInfo 发货信息记录表，用来记录发货次数超过一次的数据
type RetInfo struct {
	RetId      int    `db:"retId" json:"retId"`           // 发货表id
	RetAmount  int    `db:"retAmount" json:"retAmount"`   // 发货数量
	LdssId     int    `db:"ldss_id" json:"ldss_id"`       // 关联的 ldss 表 lid，用于根据 lid 查询数据
	ReturnDate string `db:"retDate" json:"retDate"`       // 发货时间
	RetLdsNum  string `db:"retLdsNum" json:"retLdsNum"`   // 发货快递单号
	RetLdsName string `db:"retLdsName" json:"retLdsName"` // 发货快递名称
	RetMate    string `db:"retMate" json:"retMate"`       // 发货商品名称
	RetMark    string `db:"retMark" json:"retMark"`       // 发货备注信息
}

type Data struct {
	*Sale
	*ClientAddr
	*Ldss
	*User
	*ReturnGoods
	*RetInfo
}
