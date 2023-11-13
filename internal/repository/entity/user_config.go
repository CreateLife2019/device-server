package entity

import "time"

type SessionInfo struct {
	// 开始时间
	StartTime time.Time `json:"startTime"`
	// 结束时间
	EndTime time.Time `json:"endTime"`
	// 类型 1 个人 2 群组
	SessionType int `json:"sessionType"`
	// 目标id
	DestId int64 `json:"destId"`
}
type Agent struct {
	Ip      string `json:"ip"`
	Timeout int    `json:"timeout"`
}
type UserConfig struct {
	Base
	MessageIntercept int           `json:"messageIntercept"` // 消息拦截
	SyncSessions     []SessionInfo `json:"syncMessage"`      // 会话同步
	UserId           int64         `json:"userId"`           // 用户id
	Agents           []Agent       `json:"agents"`           // 代理
}
