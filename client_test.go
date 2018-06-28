package cmq

import (
	"fmt"
	"github.com/baocaixiong/cmq-golang-sdk/models"
	"testing"
)

func TestClient_Send(t *testing.T) {
	msg := models.NewPublishMessageReq("order", `{"order_id": "asfsadf"}`)
	clt := NewClient(
		SetCredential(&Credential{
			SecretId:  "123",
			SecretKey: "",
		}),
		Region("bj"),
		NetEnv("wan"),
	)
	msg.RoutingKey = StringPtr("order.create")

	resp := models.NewPublishMessageResp()
	if e := clt.Send(msg, resp); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_ReceiveMessage(t *testing.T) {
	msg := models.NewReceiveMessageReq("bid")
	clt := NewClient(
		SetCredential(&Credential{
			SecretId:  "123",
			SecretKey: "",
		}),
		Region("bj"),
		NetEnv("wan"),
	)
	resp := models.NewReceiveMessageResp()
	if e := clt.Send(msg, resp); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(resp)
	}
}
