package local_invoke

import (
	"agent/app/common/enums/llm_enums"
	"agent/app/utils"
	"context"
	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/schema"
	"testing"
	"time"
)

func TestLocalInvokeLLM(t *testing.T) {
	ctx := context.Background()
	model, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		// 基础配置
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Timeout: 30 * time.Second,         // 请求超时时间

		// 模型配置
		Model: llm_enums.QWen3_14b.String(), // 模型名称
		//Format: json.RawMessage("json"),           // 输出格式（可选）

	})

	// 准备消息
	messages := []*schema.Message{
		schema.SystemMessage("你是一个助手"),
		schema.UserMessage("介绍一下 Ollama"),
	}

	response, err := model.Generate(ctx, messages)
	if err != nil {
		t.Fatal(err)
		return
	}

	// 处理回复
	t.Logf("response content: %s", response.Content)
	t.Logf("response: %s", utils.ToJsonString(response))
}

func TestLocalInvokeLLMGenerate(t *testing.T) {
	llm, _ := NewMyChatModel(new(MyChatModelConfig))
	llm = llm.Model(llm_enums.QWen3_14b.String())
	messages := []*schema.Message{
		schema.SystemMessage("你是一个助手"),
		schema.UserMessage("介绍一下 Ollama"),
	}

	response, err := llm.Generate(context.Background(), messages)
	if err != nil {
		t.Fatal(err)
		return
	}
	// 处理回复
	t.Logf("response content: %s", response.Content)
	t.Logf("response: %s", utils.ToJsonString(response))
}
