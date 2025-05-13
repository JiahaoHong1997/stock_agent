package llm_params

import (
	"errors"
	"net/http"
	"time"
)

type MyChatModel struct {
	Client     *http.Client
	ApiKey     string
	BaseURL    string
	Model      string
	Timeout    time.Duration
	RetryCount int
}

type MyChatModelConfig struct {
	APIKey string
}

func NewMyChatModel(config *MyChatModelConfig) (*MyChatModel, error) {
	if config.APIKey == "" {
		return nil, errors.New("api key is required")
	}

	return &MyChatModel{
		Client: &http.Client{},
		ApiKey: config.APIKey,
	}, nil
}
