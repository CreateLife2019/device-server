package constants

const (
	MessageSuc = "请求成功"
)
const (
	MessageFailedConflictAccount       = "账号已存在不能新增"
	MessageFailedNotSamePassword       = "两次密码不一致"
	MessageFailedSamePassword          = "新旧密码不能相同"
	MessageFailedNotFound              = "未找到该账号"
	MessageFailedWrongPassword         = "旧密码错误"
	MessageFailedNoProxy               = "无代理信息"
	MessageFailedServer                = "服务器错误"
	MessageFailedConflictGroup         = "分组已存在不能新增或更新"
	MessageFailedNotAllowedDeleteGroup = "分组下有用户不能删除"
	MessageFailedGroupNotFound         = "未找到该分组"
	MessageFailedForbiddenDeleteGroup  = "系统分组不能删除"
)

const (
	Status200 = "200"
	Status500 = "500"
	Status400 = "400"
)
