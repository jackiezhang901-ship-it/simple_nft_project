package base

import "time"

type User struct {
	Id          int64     `gorm:"column:user_id;AUTO_INCREMENT;primary_key" json:"user_id"` // 主键
	Address     string    `gorm:"column:address;NOT NULL" json:"address"`                   // 用户地址
	IsAllowed   bool      `gorm:"column:is_allowed;default:0;NOT NULL" json:"is_allowed"`   // 是否允许用户访问
	IsSigned    bool      `gorm:"column:is_signed;default:0" json:"is_signed"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;type:timestamp;autoCreateTime:milli;comment:创建时间"`        // 创建时间
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time;type:timestamp;autoUpdateTime:milli;comment:更新时间"`        // 更新时间
	LastLoginAt time.Time `json:"last_login_at" gorm:"column:last_login_at;type:timestamp;autoLastLoginAt:milli;comment:最后登录时间"` // 最后登录时间

	Username     string `gorm:"column:username;NOT NULL" json:"username"`           // 用户地址
	Email        string `gorm:"column:email;NOT NULL" json:"email"`                 // 用户地址
	PasswordHash string `gorm:"column:password_hash;NOT NULL" json:"password_hash"` // 用户地址
}

func UserTableName() string {
	return "users"
}
