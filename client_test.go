package cmq

import (
	"fmt"
	"github.com/baocaixiong/cmq-golang-sdk/models"
	"testing"
	"context"
	"time"
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
	if e := clt.Send(context.TODO(), msg, resp); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_ReceiveMessage(t *testing.T) {
	msg := models.NewReceiveMessageReq("queue")
	msg.PollingWaitSeconds = IntPtr(10)
	clt := NewClient(
		SetCredential(&Credential{
			SecretId:  "",
			SecretKey: "",
		}),
		Region("bj"),
		NetEnv("wan"),
	)
	resp := models.NewReceiveMessageResp()

	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	go func() {
		select {
		case <-ctx.Done():
			cancel()
			fmt.Println("cancel")
		}
	}()
	if e := clt.Send(ctx, msg, resp); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(resp.Message)
	}
}
