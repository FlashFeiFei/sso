package regster

//注册的基础信息
type UserRegisterBase struct {
	//密码
	Password string `form:"password" json:"password" binding:"required"`
	//验证码
	Code string `form:"code" json:"code" binding:"required"`
}

//手机号注册信息
type UserRegisterInfoByPhone struct {
	UserRegisterBase
	//手机号
	Phone string `form:"phone" json:"phone" binding:"required"`
}

//email注册信息
type UserRegisterInfoByEmail struct {
	UserRegisterBase
	//邮箱
	Email string `form:"email" json:"email" binding:"required"`
}
