package llm_enums

type LLM string

const (
	DeepSeekR1_14b LLM = "deepseek-r1:14b" // reasoning
	QWen3_14b      LLM = "qwen3:14b"       // reasoning、tools
)

func (l LLM) String() string {
	return string(l)
}
