package service

import (
	"errors"
	"web3-token-info-service/internal/model"
	"web3-token-info-service/internal/pubsub"
)

func FetchTokenInfo(key string, tokenName string) (*model.TokenInfo, error) {
	if key == "" {
		return nil, errors.New("key is required")
	}

	token, ok := model.MockTokenDB[tokenName]
	if !ok {
		return nil, errors.New("token not found")
	}

	return &token, nil
}

func StartSubscriber() {
	go pubsub.StartSubscriber()
}
