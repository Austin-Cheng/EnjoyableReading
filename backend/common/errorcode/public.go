package errorcode

import "github.com/dulisoft/spirit/core/errorx"

var publicModule = errorx.New("BasicService.Public.")

var (
	PublicInvalidParameter     = publicModule.Description("PublicInvalidParameter", "内部错误")
	PublicDatabaseError        = publicModule.Description("PublicDatabaseError", "数据库异常")
	PublicInvalidParameterJson = publicModule.Description("PublicInvalidParameterJson", "参数值校验不通过：json格式错误")
)
