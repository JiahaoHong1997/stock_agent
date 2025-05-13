package main

//func main() {
//	config := ark.DefaultConfig("5dd4325f-db9d-486e-a7ba-0b9e959d7f21")
//	config.BaseURL = "https://ark.cn-beijing.volces.com/api/v3"
//	client := ark.NewClientWithConfig(config)
//
//	fmt.Println("----- standard request -----")
//	resp, err := client.CreateChatCompletion(
//		context.Background(),
//		ark.ChatCompletionRequest{
//			Model: "ep-20250507011721-mk5zw",
//			Messages: []ark.ChatCompletionMessage{
//				{
//					Role:    ark.ChatMessageRoleSystem,
//					Content: "你是人工智能助手",
//				},
//				{
//					Role:    ark.ChatMessageRoleUser,
//					Content: "常见的十字花科植物有哪些？",
//				},
//			},
//		},
//	)
//	if err != nil {
//		fmt.Printf("ChatCompletion error: %v\n", err)
//		return
//	}
//	fmt.Println(resp.Choices[0].Message.Content)
//
//	fmt.Println("----- streaming request -----")
//	stream, err := client.CreateChatCompletionStream(
//		context.Background(),
//		ark.ChatCompletionRequest{
//			Model: "ep-20250507011721-mk5zw",
//			Messages: []ark.ChatCompletionMessage{
//				{
//					Role:    ark.ChatMessageRoleSystem,
//					Content: "你是人工智能助手",
//				},
//				{
//					Role:    ark.ChatMessageRoleUser,
//					Content: "常见的十字花科植物有哪些？",
//				},
//			},
//		},
//	)
//	if err != nil {
//		fmt.Printf("stream chat error: %v\n", err)
//		return
//	}
//	defer stream.Close()
//
//	for {
//		recv, err := stream.Recv()
//		if err == io.EOF {
//			return
//		}
//		if err != nil {
//			fmt.Printf("Stream chat error: %v\n", err)
//			return
//		}
//
//		if len(recv.Choices) > 0 {
//			fmt.Print(recv.Choices[0].Delta.Content)
//		}
//	}
//}

//import (
//	"context"
//	"fmt"
//	"log"
//	"os"
//	"time"
//
//	"github.com/cloudwego/eino-ext/components/model/ark"
//	"github.com/cloudwego/eino/components/tool"
//	"github.com/cloudwego/eino/compose"
//	"github.com/cloudwego/eino/schema"
//)
//
//func main() {
//	// 初始化 tools
//	todoTools := []tool.BaseTool{
//		getAddTodoTool(), // NewTool 构建
//		updateTool,       // InferTool 构建
//		&ListTodoTool{},  // 实现Tool接口
//		searchTool,       // 官方封装的工具
//	}
//
//	// 创建并配置 ChatModel
//	timeout := 30 * time.Second
//	chatModel, err := ark.NewChatModel(context.Background(), &ark.ChatModelConfig{
//		APIKey:  "5dd4325f-db9d-486e-a7ba-0b9e959d7f21",
//		Region:  "cn-beijing",
//		Model:   "ep-20250507011721-mk5zw",
//		Timeout: &timeout,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 获取工具信息并绑定到 ChatModel
//	toolInfos := make([]*schema.ToolInfo, 0, len(todoTools))
//	for _, tool := range todoTools {
//		info, err := tool.Info(ctx)
//		if err != nil {
//			log.Fatal(err)
//		}
//		toolInfos = append(toolInfos, info)
//	}
//	err = chatModel.BindTools(toolInfos)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 创建 tools 节点
//	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
//		Tools: todoTools,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 构建完整的处理链
//	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
//	chain.
//		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
//		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))
//
//	// 编译并运行 chain
//	agent, err := chain.Compile(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 运行示例
//	resp, err := agent.Invoke(ctx, []*schema.Message{
//		{
//			Role:    schema.User,
//			Content: "添加一个学习 Eino 的 TODO，同时搜索一下 cloudwego/eino 的仓库地址",
//		},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 输出结果
//	for _, msg := range resp {
//		fmt.Println(msg.Content)
//	}
//}

import (
	"context"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

func main() {
	fmt.Println("server start")
	// 注册事件回调，OnP2MessageReceiveV1 为接收消息 v2.0；OnCustomizedEvent 内的 message 为接收消息 v1.0。
	eventHandler := dispatcher.NewEventDispatcher("jrRiD478dBgsoX0aN4uJLehdgaDaWLru", "UC1ZUJtLmALjSAu11B3GHctaseBqIbZM").
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			fmt.Printf("[ OnP2MessageReceiveV1 access ], data: %s\n", larkcore.Prettify(event))
			return nil
		})
	// 创建Client
	cli := larkws.NewClient("cli_a8aa598d02a1d01c", "wCMAHrc3yIyzVd5bAQEvHeg6RCoKGzCD",
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)
	// 启动客户端
	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
