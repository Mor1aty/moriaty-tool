package moconstant

const (

	/**
	 * 执行成功
	 */
	CodeExecSuccess = 200

	/**
	 * 获取成功
	 */
	CodeObtainSuccess = 201

	/**
	 * 获取数据为空
	 */
	CodeObtainNull = 202

	/**
	 * 执行成功
	 */
	CodeExecFailure = 400

	/**
	 * 获取失败
	 */
	CodeObtainFailure = 401

	/**
	 * 参数校验未通过
	 */
	CodeParamFailure = 402

	/**
	 * Token 校验未通过
	 */
	CodeTokenFailure = 403

	/**
	 * 权限校验未通过
	 */
	CodeLevelFailure = 404

	/**
	 * 请求方法错误
	 */
	CodeMethodFailure = 405

	/**
	 * 服务器异常
	 */
	CodeServerException = 500

	/**
	 * 返回错误信息
	 */
	MsgError = "服务器错误"

	/**
	 * 返回成功信息
	 */
	MsgSuccess = "执行成功"

	/**
	 * 返回验证未过信息
	 */
	MsgTokenFailed = "Token 过期"

	/**
	 * 返回权限未过信息
	 */
	MsgLevelFailed = "权限不足"

	/**
	 * 返回邮件成功信息
	 */
	MsgMailSuccess = "邮件发送成功"

	/**
	 * 返回邮件错误信息
	 */
	MsgMailError = "邮件发送错误"

	/**
	 * 返回无参数信息
	 */
	MsgParamEmpty = "缺少参数"

	/**
	 * 返回参数为空信息
	 */
	MsgParamNull = " 参数不能为空"

	/**
	 * 返回参数邮件信息
	 */
	MsgParamEmail = " 不是有效邮件格式"

	/**
	 * 返回参数数字信息
	 */
	MsgParamNumber = " 不是有效纯数字格式"

	/**
	 * 返回参数手机信息
	 */
	MsgParamPhone = " 不是有效手机格式"

	/**
	 * 日期格式：yyyy-MM-dd HH:mm:ss
	 */
	DateType01 = "2006-01-02 15:04:05"

	/**
	 * 日期格式：yyyy-MM-dd
	 */
	DateType02 = "2006-01-02"

	/**
	 * 日期格式：HH:mm:ss
	 */
	DateType03 = "15:04:05"

	/**
	 * 日期格式：yyyy-MM-dd HH:mm:ss:SSS
	 */
	DateType04 = "2006-01-02 15:04:05:000"

	/**
	 * 日期格式：HH:mm:ss:SSS
	 */
	DateType05 = "15:04:05:000"

	/**
	 * 邮箱正则
	 */
	RegularMail = "[\\w\\.\\-]+@([\\w\\-]+\\.)+[\\w\\-]+"

	/**
	 * 纯数字正则
	 */
	RegularNumber = "^(0|[1-9][0-9]*|-[1-9][0-9]*)$"

	/**
	 * 整数正则
	 */
	RegularInteger = "^[-\\+]?[\\d]*$"

	/**
	 * 浮点数正则
	 */
	RegularFloat = "^[-\\+]?[.\\d]*$"

	/**
	 * 手机正则
	 */
	RegularPhone = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(17[013678])|(18[0,5-9]))\\d{8}$"
)
