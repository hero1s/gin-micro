package model

// 推广渠道表
type Channel struct {
	Cid            int    `json:"cid" desc:"渠道编号"`
	Cname          string `json:"cname" desc:"渠道名称"`
	CompanyName    string `json:"company_name" desc:"公司名"`
	CreateIdentity string `json:"create_identity" desc:"创建者身份证"`
	CreateName     string `json:"create_name" desc:"创建者姓名"`
	CreateUid      int    `json:"create_uid" desc:"创建用户ID"`
	Phone          string `json:"phone" desc:"手机号"`
	Ptime          int    `json:"ptime" desc:"操作时间"`
	Ratio          int    `json:"ratio" desc:"分成比例"`
	Status         int8   `json:"status" desc:"状态(0正常1禁用)"`
}

func (c *Channel) TableName() string {
	return "channel"
}
func (c *Channel) DbName() string {
	return "cherry"
}
func (c *Channel) FullTableName() string {
	return "cherry.channel"
}

var TbChannel = &Channel{}
