package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*AuthMail]()
}

// AuthMail 邮箱认证配置
type AuthMail struct {
	pkgmodel.Base

	Name          *string `json:"name"`           // 配置名称
	Host          *string `json:"host"`           // 邮件服务器地址
	ConnectString *string `json:"connect_string"` // 连接字符串
	IsActive      *bool   `json:"is_active"`      // 是否激活
}
