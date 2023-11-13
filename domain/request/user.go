package request

type UserListRequest struct {
	PageInfo
}

// 增加用户代理
type AgentRequest struct {
	UserId int64 `json:"UserId"`
}
