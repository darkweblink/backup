package model

// 授权key配置表
type AuthorityConfig struct {
	Id        uint   `gorm:"column:id" json:"id" form:"id"`
	UKey      string `gorm:"column:u_key" json:"u_key" form:"u_key"`                //key
	Source    string `gorm:"column:source" json:"source" form:"source"`             //key来源
	StartTime string `gorm:"column:start_time" json:"start_time" form:"start_time"` //开始时间
	EndTime   string `gorm:"column:end_time" json:"end_time" form:"end_time"`       //结束时间
	Frequency int    `gorm:"column:frequency" json:"frequency" form:"frequency"`    //次数
	IsDeleted int8   `gorm:"column:is_deleted" json:"is_deleted" form:"is_deleted"` //是否删除 0-正常 1-删除
	CreatedAt string `gorm:"column:created_at" json:"created_at" form:"created_at"` //创建时间
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at" form:"updated_at"` //修改时间
}

// 特殊配置表
type AuthorityExclusive struct {
	Id                uint   `gorm:"column:id" json:"id" form:"id"`
	AuthorityConfigId int    `gorm:"column:authority_config_id" json:"authority_config_id" form:"authority_config_id"`
	UseType           int    `gorm:"column:use_type" json:"use_type" form:"use_type"`
	Frequency         int    `gorm:"column:frequency" json:"frequency" form:"frequency"`    //专用次数
	StartTime         string `gorm:"column:start_time" json:"start_time" form:"start_time"` //专用有效开始时间
	EndTime           string `gorm:"column:end_time" json:"end_time" form:"end_time"`       //专用有效结束时间
	IsDeleted         int8   `gorm:"column:is_deleted" json:"is_deleted" form:"is_deleted"` //是否删除 0-正常 1-删除
	CreatedAt         string `gorm:"column:created_at" json:"created_at" form:"created_at"` //创建时间
	UpdatedAt         string `gorm:"column:updated_at" json:"updated_at" form:"updated_at"` //修改时间
}

// 请求记录表
type AuthorityLog struct {
	Id                uint   `gorm:"column:id" json:"id" form:"id"`
	AuthorityConfigId int    `gorm:"column:authority_config_id" json:"authority_config_id" form:"authority_config_id"` //authority_config_id
	Type              int8   `gorm:"column:type" json:"type" form:"type"`                                              //类型 0-无 1-专用 2-通用
	UseType           int    `gorm:"column:use_type" json:"use_type" form:"use_type"`
	Desc              string `gorm:"column:desc" json:"desc" form:"desc"`                   //描述
	CreatedAt         string `gorm:"column:created_at" json:"created_at" form:"created_at"` //创建时间
}
