/**
 * 功能描述: 自定义错误信息code
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package errno

// 错误码定义
// 第1位: 服务级别     1(系统级错误)  2(普通错误)
// 第2-3位: 服务模块   01(用户)
// 第4-5位: 错误码	  01(具体错误代码)
var (
	// 通用错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
)
