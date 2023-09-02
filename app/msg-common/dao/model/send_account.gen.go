// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSendAccount = "send_account"

// SendAccount mapped from table <send_account>
type SendAccount struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SendChanel string `gorm:"column:send_chanel;not null;comment:发送渠道" json:"send_chanel"` // 发送渠道
	Config     string `gorm:"column:config;not null;comment:账户配置" json:"config"`           // 账户配置
	Title      string `gorm:"column:title;not null;comment:账号名称" json:"title"`             // 账号名称
}

// TableName SendAccount's table name
func (*SendAccount) TableName() string {
	return TableNameSendAccount
}