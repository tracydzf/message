// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsRecord = "sms_record"

// SmsRecord mapped from table <sms_record>
type SmsRecord struct {
	ID                int64  `gorm:"column:id;primaryKey" json:"id"`
	MessageTemplateID int64  `gorm:"column:message_template_id;not null;comment:消息模板ID" json:"message_template_id"` // 消息模板ID
	Phone             int64  `gorm:"column:phone;not null;comment:手机号" json:"phone"`                                // 手机号
	SupplierID        int32  `gorm:"column:supplier_id;not null;comment:发送短信渠道商的ID" json:"supplier_id"`             // 发送短信渠道商的ID
	SupplierName      string `gorm:"column:supplier_name;not null;comment:发送短信渠道商的名称" json:"supplier_name"`         // 发送短信渠道商的名称
	MsgContent        string `gorm:"column:msg_content;not null;comment:短信发送的内容" json:"msg_content"`                // 短信发送的内容
	SeriesID          string `gorm:"column:series_id;not null;comment:下发批次的ID" json:"series_id"`                    // 下发批次的ID
	ChargingNum       int32  `gorm:"column:charging_num;not null;comment:计费条数" json:"charging_num"`                 // 计费条数
	ReportContent     string `gorm:"column:report_content;not null;comment:回执内容" json:"report_content"`             // 回执内容
	Status            int32  `gorm:"column:status;not null;comment:短信状态： 10.发送 20.成功 30.失败" json:"status"`          // 短信状态： 10.发送 20.成功 30.失败
	SendDate          int32  `gorm:"column:send_date;not null;comment:发送日期：20211112" json:"send_date"`              // 发送日期：20211112
	Created           int32  `gorm:"column:created;not null;comment:创建时间" json:"created"`                           // 创建时间
	Updated           int32  `gorm:"column:updated;not null;comment:更新时间" json:"updated"`                           // 更新时间
}

// TableName SmsRecord's table name
func (*SmsRecord) TableName() string {
	return TableNameSmsRecord
}
