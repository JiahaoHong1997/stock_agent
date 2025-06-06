package local_invoke

import (
	"agent/app/utils"
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"net/http"
	"time"
)

type MyChatModel struct {
	client     *http.Client
	apiKey     string
	baseURL    string
	model      string
	timeout    *time.Duration
	retryCount int
}

type MyChatModelConfig struct {
	APIKey string
}

func NewMyChatModel(config *MyChatModelConfig) (*MyChatModel, error) {
	return &MyChatModel{
		client: &http.Client{},
		apiKey: config.APIKey,
	}, nil
}

func (m *MyChatModel) BaseURL(baseURL string) *MyChatModel {
	m.baseURL = baseURL
	return m
}

func (m *MyChatModel) Model(modelName string) *MyChatModel {
	m.model = modelName
	return m
}

// 定义 Option 结构体
type MyChatModelOptions struct {
	Options    *model.Options
	RetryCount int
	Timeout    *time.Duration
}

// 定义 Option 函数
func WithRetryCount(count int) model.Option {
	return model.WrapImplSpecificOptFn(func(o *MyChatModelOptions) {
		o.RetryCount = count
	})
}

func WithTimeout(timeout *time.Duration) model.Option {
	return model.WrapImplSpecificOptFn(func(o *MyChatModelOptions) {
		o.Timeout = timeout
	})
}

func (m *MyChatModel) Generate(ctx context.Context, messages []*schema.Message, opts ...model.Option) (*schema.Message, error) {
	// 1. 处理选项
	options := &MyChatModelOptions{
		Options: &model.Options{
			Model: &m.model,
		},
		RetryCount: m.retryCount,
		Timeout:    m.timeout,
	}
	options.Options = model.GetCommonOptions(options.Options, opts...)
	options = model.GetImplSpecificOptions(options, opts...)

	// 2. 开始生成前的回调
	ctx = callbacks.OnStart(ctx, &model.CallbackInput{
		Messages: messages,
		Config: &model.Config{
			Model: *options.Options.Model,
		},
	})

	// 3. 执行生成逻辑
	response, err := m.doGenerate(ctx, messages, options)

	// 4. 处理错误和完成回调
	if err != nil {
		ctx = callbacks.OnError(ctx, err)
		return nil, err
	}

	ctx = callbacks.OnEnd(ctx, &model.CallbackOutput{
		Message: response,
	})

	return response, nil
}

func (m *MyChatModel) Stream(ctx context.Context, messages []*schema.Message, opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	// 1. 处理选项
	options := &MyChatModelOptions{
		Options: &model.Options{
			Model: &m.model,
		},
		RetryCount: m.retryCount,
		Timeout:    m.timeout,
	}
	options.Options = model.GetCommonOptions(options.Options, opts...)
	options = model.GetImplSpecificOptions(options, opts...)

	// 2. 开始流式生成前的回调
	ctx = callbacks.OnStart(ctx, &model.CallbackInput{
		Messages: messages,
		Config: &model.Config{
			Model: *options.Options.Model,
		},
	})

	// 3. 创建流式响应
	// Pipe产生一个StreamReader和一个StreamWrite，向StreamWrite中写入可以从StreamReader中读到，二者并发安全。
	// 实现中异步向StreamWrite中写入生成内容，返回StreamReader作为返回值
	// ***StreamReader是一个数据流，仅可读一次，组件自行实现Callback时，既需要通过OnEndWithCallbackOutput向callback传递数据流，也需要向返回一个数据流，需要对数据流进行一次拷贝
	// 考虑到此种情形总是需要拷贝数据流，OnEndWithCallbackOutput函数会在内部拷贝并返回一个未被读取的流
	// 以下代码演示了一种流处理方式，处理方式不唯一
	sr, sw := schema.Pipe[*model.CallbackOutput](1)

	// 4. 启动异步生成
	go func() {
		defer sw.Close()

		// 流式写入
		m.doStream(ctx, messages, options, sw)
	}()

	// 5. 完成回调
	_, nsr := callbacks.OnEndWithStreamOutput(ctx, sr)

	return schema.StreamReaderWithConvert(nsr, func(t *model.CallbackOutput) (*schema.Message, error) {
		return t.Message, nil
	}), nil
}

func (m *MyChatModel) WithTools(tools []*schema.ToolInfo) (model.ToolCallingChatModel, error) {
	// 实现工具绑定逻辑
	return nil, nil
}

func (m *MyChatModel) doGenerate(ctx context.Context, messages []*schema.Message, opts *MyChatModelOptions) (*schema.Message, error) {
	// 实现生成逻辑
	llm, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		// 基础配置
		BaseURL: utils.TernaryString(m.baseURL != "", m.baseURL, "http://localhost:11434"),        // 服务地址
		Timeout: utils.TernaryTime(m.timeout != nil, utils.ToDuration(m.timeout), 30*time.Second), // 请求超时时间

		// 模型配置
		Model: m.model, // 模型名称
		//Format: json.RawMessage("json"),           // 输出格式（可选）

	})
	if err != nil {
		return nil, err
	}

	llmResp, err := llm.Generate(ctx, messages)
	if err != nil {
		fmt.Println("llm generate error:", err)
		return nil, err
	}
	return llmResp, nil
}

func (m *MyChatModel) doStream(ctx context.Context, messages []*schema.Message, opts *MyChatModelOptions, sr *schema.StreamWriter[*model.CallbackOutput]) {
	// 流式生成文本写入sr中
	return
}
