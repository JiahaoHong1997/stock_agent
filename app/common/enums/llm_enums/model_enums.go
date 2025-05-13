package llm_enums

type LLM string

const (
	DeepSeekR1_14b LLM = "deepseek-r1:14b"
)

func (l LLM) String() string {
	return string(l)
}
