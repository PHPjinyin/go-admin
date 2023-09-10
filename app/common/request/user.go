package request

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetMessages 自定义错误提示
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "用户名称不能为空",
		"Mobile.required":   "手机号码不能为空",
		"Password.required": "用户密码不能为空",
	}
}

// Login 用户登录校验
type Login struct {
	Username string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名不能为空",
		"password.required": "用户密码不能为空",
	}
}
