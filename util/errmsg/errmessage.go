package errmsg

// 把自定义的错误code 翻译成错误信息，抛给前端。就是一个字典

const (
	SUCCESS = 1
	ERROR   = 500

	// 1000 <= code < 2000 用户模块User的错误
	ERROR_USERNAME_USED    = 1001 // 用户名已经存在！
	ERROR_PASSWORD         = 1002 // 密码错误！
	ERROR_USER_NOT_EXIST   = 1003 // 用户名不存在！
	ERROR_TOKEN_NOT_EXIST  = 1004 // 用户携带的TOKEN不存在！
	ERROR_TOKEN_TIMEOUT    = 1005 // 用户携带的TOKEN过期了！
	ERROR_TOKEN_WRONG      = 1006 // 用户携带的TOKEN错误、虚假的！
	ERROR_TOKEN_TYPE_WRONG = 1007 // 用户携带的TOKEN格式错误！
	ERROR_USER_NO_RIGHT    = 1008 // 用户没有管理权限！

	// 2000 <= code < 3000 分类模块的错误
	ERROR_CATEGRORYNAME_USED  = 2001 // 分类名已经存在！
	ERROR_CATEGRORY_NOT_EXIST = 2002 //  分类不存在！

	// 3000 <= code < 4000 文章模块的错误
	ERROR_ARTICLE_NOT_EXIST = 3001 // 文章不存在！
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已经存在！",
	ERROR_PASSWORD:         "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户名不存在！",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在！",
	ERROR_TOKEN_TIMEOUT:    "TOKEN过期了！",
	ERROR_TOKEN_WRONG:      "TOKEN错误！",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误！",
	ERROR_USER_NO_RIGHT:    "用户没有管理权限！",

	ERROR_CATEGRORYNAME_USED:  "分类已经存在！",
	ERROR_CATEGRORY_NOT_EXIST: "分类不存在！",

	ERROR_ARTICLE_NOT_EXIST: "文章不存在！",
} // int=code ， string=返回的错误文字描述

func GetErrMsg(code int) string {
	return codeMsg[code]
}
