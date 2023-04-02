package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type Config struct {
	APP_ID                 string
	APP_SECRET             string
	APP_ENCRYPT_KEY        string
	APP_VERIFICATION_TOKEN string
	BOT_NAME               string
	GIN_MODE               string
	HTTP_PORT              string
	PORT                   string
	OPENAI_KEY             string
}

func main() {
	// create a new config struct
	config := Config{}

	// get environment variables and set them in config struct
	config.APP_ID = os.Getenv("APP_ID")
	config.APP_SECRET = os.Getenv("APP_SECRET")
	config.APP_ENCRYPT_KEY = os.Getenv("APP_ENCRYPT_KEY")
	config.APP_VERIFICATION_TOKEN = os.Getenv("APP_VERIFICATION_TOKEN")
	config.BOT_NAME = os.Getenv("BOT_NAME")
	config.GIN_MODE = os.Getenv("GIN_MODE")
	config.PORT = os.Getenv("PORT")
	config.HTTP_PORT = os.Getenv("HTTP_PORT")
	config.OPENAI_KEY = os.Getenv("OPENAI_KEY")

	fmt.Printf("%+v\n", config)

	// 创建注册消息处理器
	handler := dispatcher.NewEventDispatcher(config.APP_VERIFICATION_TOKEN, config.APP_ENCRYPT_KEY).OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	}).OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	}).OnP2UserCreatedV3(func(ctx context.Context, event *larkcontact.P2UserCreatedV3) error {
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	})

	// 创建卡片行为处理器
	cardHandler := larkcard.NewCardActionHandler(config.APP_VERIFICATION_TOKEN, config.APP_ENCRYPT_KEY, func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		fmt.Println(larkcore.Prettify(cardAction))

		// 返回卡片消息
		//return getCard(), nil

		//custom resp
		//return getCustomResp(),nil

		// 无返回值
		return nil, nil
	})

	// 注册处理器
	g := gin.Default()
	g.POST("/webhook/event", sdkginext.NewEventHandlerFunc(handler))
	g.POST("/webhook/card", sdkginext.NewCardActionHandlerFunc(cardHandler))

	// 启动服务
	err := g.Run(":9999")
	if err != nil {
		panic(err)
	}
}
