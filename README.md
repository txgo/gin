---
title: TaskBot
description: Taskbot is a chatbot that uses OpenAI's GPT-3 language model to generate responses to user messages. It is built using the Lark API and runs on the Gin web framework.
tags:
  - gin
  - golang
  - lark
  - bot
  - openai
  - ChatGPT
---

# Deploy to railway 

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/dTvvSf)

## ✨ Features

- Gin
- Go

Configuration
Taskbot is configured using environment variables:

- APP_I-D: The App ID for the Lark API.
- APP_SECRET: The App Secret for the Lark API.
- ENCRYPT_KEY: The Encrypt Key for the Lark API.
- VERIFICATION_TOKEN: The Verification Token for the Lark API.
- BOT_NAME: The name of the chatbot.
- GIN_MODE: The Gin mode (debug or release).
- PORT: The port on which to listen for incoming requests.
- HTTP_PORT: The HTTP port on which to listen for incoming requests.
- OPENAI_API_KEY: The API Key for the OpenAI API.

# Usage
To use Taskbot, add it as a contact in your Lark workspace and start a one-on-one chat with it. Taskbot will respond to your messages using OpenAI's GPT-3 language model.

Implementation
Taskbot is built using the following technologies:

- Go: The programming language used to write Taskbot.
- Lark API: The API used to interact with Lark.
- Gin: The web framework used to handle incoming requests.
- OpenAI API: The API used to generate responses to user messages.

## License
Taskbot is licensed under the MIT License.

## 💁‍♀️ Start to develop

- Connect to your Railway project `railway link`
- Start the development server `railway run go run main.go`

## 📝 Notes


## Taskbot（中文版）
Taskbot是一个聊天机器人，使用OpenAI的GPT-3语言模型生成用户消息的响应。它使用Lark API构建，并在Gin Web框架上运行。

Taskbot在一对一聊天中监听消息，并使用OpenAI API生成响应。生成的响应将发送回聊天中的用户。

## 配置
Taskbot使用环境变量进行配置：

- APP_ID：Lark API的App ID。
- APP_SECRET：Lark API的App Secret。
- ENCRYPT_KEY：Lark API的Encrypt Key。
- VERIFICATION_TOKEN：Lark API的Verification Token。
- BOT_NAME：聊天机器人的名称。
- GIN_MODE：Gin模式（debug或release）。
- PORT：监听传入请求的端口。
- HTTP_PORT：监听传入请求的HTTP端口。
- OPENAI_API_KEY：OpenAI API的API密钥。
# 用法
要使用Taskbot，请将其添加为Lark工作区中的联系人，并开始与其进行一对一聊天。 Taskbot将使用OpenAI的GPT-3语言模型响应您的消息。

## 实现
Taskbot使用以下技术构建：

- Go：编写Taskbot使用的编程语言。
- Lark API：用于与Lark进行交互的API。
- Gin：用于处理传入请求的Web框架。
- OpenAI API：用于生成用户消息响应的API。

## 许可证
Taskbot采用MIT许可证授权。