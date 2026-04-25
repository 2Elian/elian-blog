package model

// SiteConfig 网站配置表
type SiteConfig struct {
	Model
	Key   string `json:"key" gorm:"uniqueIndex;size:100;not null"`
	Value string `json:"value" gorm:"type:longtext"`
}

func (SiteConfig) TableName() string { return "site_config" }

// OperationLog 操作日志表
type OperationLog struct {
	Model
	UserID uint   `json:"user_id" gorm:"index"`
	Module string `json:"module" gorm:"size:50"`
	Action string `json:"action" gorm:"size:50"`
	Method string `json:"method" gorm:"size:10"`
	URL    string `json:"url" gorm:"size:500"`
	IP     string `json:"ip" gorm:"size:50"`
	Desc   string `json:"desc" gorm:"size:500"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
}

func (OperationLog) TableName() string { return "operation_log" }
