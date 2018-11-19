package models

import "fmt"

type BatchSendMessageReq struct {
	*Request
	QueueName    string   `name:"queueName"`
	MsgBody      []string `name:"msgBody"`
	DelaySeconds *int     `name:"delaySeconds"`
}

type BatchSendMessageResp struct {
	*Response
	MsgList []*SendMsgId `json:"msgList"`
}

type SendMsgId struct {
	MsgId string `json:"msgId"`
}

func (sm *SendMsgId) String() string {
	return fmt.Sprintf("%s", sm.MsgId)
}

func NewBatchSendMessageReq(queueName string) *BatchSendMessageReq {
	return &BatchSendMessageReq{
		QueueName: queueName,
		Request:   InitReq("BatchSendMessage", "queue"),
	}
}

func NewBatchSendMessageResp() *BatchSendMessageResp {
	return new(BatchSendMessageResp)
}
