package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	openai "github.com/sashabaranov/go-openai"
)

type Config struct {
	AppID             string
	AppSecret         string
	EncryptKey        string
	VerificationToken string
	BotName           string
	GinMode           string
	Port              string
	HttpPort          string
	OpenAIKey         string
}

var client *lark.Client

func main() {
	// read the configuration options from environment variables
	config := &Config{
		AppID:             os.Getenv("APP_ID"),
		AppSecret:         os.Getenv("APP_SECRET"),
		EncryptKey:        os.Getenv("ENCRYPT_KEY"),
		VerificationToken: os.Getenv("VERIFICATION_TOKEN"),
		BotName:           os.Getenv("BOT_NAME"),
		GinMode:           os.Getenv("GIN_MODE"),
		Port:              os.Getenv("PORT"),
		HttpPort:          os.Getenv("HTTP_PORT"),
		OpenAIKey:         os.Getenv("OPENAI_API_KEY"),
	}

	// initialize the OpenAI API client
	openaiClient := openai.NewClient(config.OpenAIKey)

	// create a new Lark IM client
	client = lark.NewClient(config.AppID, config.AppSecret)

	// create an event dispatcher
	eventDispatcher := dispatcher.NewEventDispatcher(config.VerificationToken, config.EncryptKey).
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			log.Printf("Received P2 message receive event: %v", larkcore.Prettify(event))

			// generate a response using the OpenAI API
			completion, err := generateResponse(ctx, *event.Event.Message.Content, openaiClient)
			if err != nil {
				log.Printf("Error generating response: %v", err)
				return nil
			}

			log.Printf("Received Openai completion: %v", larkcore.Prettify(completion))

			respMsg := larkim.NewTextMsgBuilder().Text(completion).Build()

			chatId := event.Event.Message.ChatId
			resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
				ReceiveIdType(larkim.ReceiveIdTypeChatId).
				Body(larkim.NewCreateMessageReqBodyBuilder().
					MsgType(larkim.MsgTypeText).
					ReceiveId(*chatId).
					Content(respMsg).
					Build()).
				Build())

			if err != nil {
				log.Printf("Error Send Completion to Lark: %v", err)
				return nil
			}
			if !resp.Success() {
				log.Printf("Error sending message: %v", resp.Code)
				log.Printf("Error message: %s", resp.Msg)
				return nil
			}

			return nil
		}).
		OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
			log.Printf("Received P2 message read event: %v", larkcore.Prettify(event))
			return nil
		}).
		OnP2UserCreatedV3(func(ctx context.Context, event *larkcontact.P2UserCreatedV3) error {
			log.Printf("Received P2 user created event: %v", larkcore.Prettify(event))
			return nil
		})

	// create a card action handler
	cardActionHandler := larkcard.NewCardActionHandler(config.VerificationToken, config.EncryptKey, func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		log.Printf("Received card action: %v", larkcore.Prettify(cardAction))

		// TODO: handle the card action and return a response
		return nil, nil
	})

	// set up the Gin router
	router := gin.Default()

	// register the event and card action handlers
	router.POST("/webhook/event", sdkginext.NewEventHandlerFunc(eventDispatcher))
	router.POST("/webhook/card", sdkginext.NewCardActionHandlerFunc(cardActionHandler))

	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	err := router.Run(":" + config.Port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

// generateResponse uses the OpenAI API to generate a response to the user's message.
func generateResponse(ctx context.Context, message string, client *openai.Client) (string, error) {
	// set up the prompt and parameters
	completionRequest := &openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo0301,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
	}

	// generate the completion using the OpenAI API
	response, err := client.CreateChatCompletion(ctx, *completionRequest)
	if err != nil {
		return "", err
	}

	// return the generated text
	return response.Choices[0].Message.Content, nil
}
