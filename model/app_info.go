package model

// APP详情信息，所有的路由都是基于APP进行发布的
type AppInfo struct {
	Id            uint   `gorm:"column:id" json:"id" form:"id"`
	RouterFix     string `gorm:"column:router_fix" json:"router_fix" form:"router_fix"`             //路由前缀
	Name          string `gorm:"column:name" json:"name" form:"name"`                               //APP名称
	Info          string `gorm:"column:info" json:"info" form:"info"`                               //APP简介
	RouterTimeout uint   `gorm:"column:router_timeout" json:"router_timeout" form:"router_timeout"` //路由超时时间
	KeepTimeout   uint   `gorm:"column:keep_timeout" json:"keep_timeout" form:"keep_timeout"`       //长连接超时时间
	TlsOut        int    `gorm:"column:tls_out" json:"tls_out" form:"tls_out"`                      //TLS握手超时时间
	Autos         uint8  `gorm:"column:autos" json:"autos" form:"autos"`                            //登录认证方式 0-无认证 1-JWT认证
	Keys          uint8  `gorm:"column:keys" json:"keys" form:"keys"`                               //是否启用KEY认证
	AutosJwtUid   string `gorm:"column:autos_jwtUid" json:"autos_jwtUid" form:"autos_jwtUid"`       // JWT中的用户ID标识
	CreatedAt     string `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt     string `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
}

// APP路由表，用来存储路由的转发信息及是否需要授权信息
type AppRouter struct {
	Id        uint64 `gorm:"column:id" json:"id" form:"id"`
	AppId     uint   `gorm:"column:app_id" json:"app_id" form:"app_id"`       //APP_ID
	Router    string `gorm:"column:router" json:"router" form:"router"`       //路由地址，需拼接所属APP的前缀
	Targets   string `gorm:"column:targets" json:"targets" form:"targets"`    //转发地址，[]string{}格式的JSON字符串
	IsLogin   uint8  `gorm:"column:is_login" json:"is_login" form:"is_login"` //是否需要登录授权【APP中的JWT类型TOKEN】
	IsKey     uint8  `gorm:"column:is_key" json:"is_key" form:"is_key"`       //是否需要KEY授权（权限鉴定）
	KeyType   uint8  `gorm:"column:key_type" json:"key_type" form:"key_type"` //KEY授权类型 0-无需授权 1-通用授权 2-专用授权
	KeyUse    int    `gorm:"column:key_use" json:"key_use" form:"key_use"`    //专用授权的授权值
	CreatedAt string `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
}

// 转发路由命中详情表
type AppRouterLog struct {
	Id            uint64 `gorm:"column:id" json:"id" form:"id"`
	Rid           uint64 `gorm:"column:rid" json:"rid" form:"rid"`                                  //路由ID
	RUri          string `gorm:"column:r_uri" json:"r_uri" form:"r_uri"`                            //请求URL地址，从targets中读取的某个地址列表
	Code          uint   `gorm:"column:code" json:"code" form:"code"`                               //响应的HTTPCODE值
	ResponeLength uint   `gorm:"column:respone_length" json:"respone_length" form:"respone_length"` //响应的字节大小，单位：B
	ResponeTime   uint   `gorm:"column:respone_time" json:"respone_time" form:"respone_time"`       //响应耗时 单位：毫秒
	GatwatTime    uint   `gorm:"column:gatwat_time" json:"gatwat_time" form:"gatwat_time"`          // 网关处理耗时 单位：毫秒
	ClientTime    string `gorm:"column:client_time" json:"client_time" form:"client_time"`          //请求时间
}
