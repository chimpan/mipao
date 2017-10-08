package constant

/*
错误码
 */
const (
	Err_Phone_nil        = iota + 40000 //手机号码是空
	Err_Password_nil                    //密码是空
	Err_Code_nil                        //验证码是空
	Err_Code_Invaild                    //验证码已使用
	Err_Code_Unmatch                    //验证码有误
	Err_Phone_Exit                      //手机号码已存在
	Err_Phone_NO_FOUND                  //手机号码未注册
	Err_User_Id_NO_FOUND                //用户未找到
	Err_Phone_Invaild                   //手机号码格式不对
	Err_Login_Info                      //账号或密码有误
	Err_Params                          //参数错误
	Err_Not_Pic_File                    //不是图片文件
)
