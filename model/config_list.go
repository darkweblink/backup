package model

import (
	"gitee.com/xiaoyutab/xgotool/xgodb/xgodbconfig"
)

// 配置项列表，初始化时会将此处数据插入到数据库中
var ConfigList []xgodbconfig.Configure = []xgodbconfig.Configure{}
