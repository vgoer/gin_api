package models

type Admin struct {
	ID
	Username string `json:"username" grom:"not null; commen:用户名"`
	Password string `json:"password" grom:"not null; default:''; commen:密码"`
	Mobile   string `json:"mobile" gorm:"not null; index; comment:用户手机号"`
	Token    string `json:"token" gorm:"comment:token"`
	Timestamps
	SoftDeletes
}
