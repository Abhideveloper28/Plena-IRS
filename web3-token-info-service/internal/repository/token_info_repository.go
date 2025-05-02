package repository

import (
	"fmt"
	"web3-token-info-service/internal/model"
)

func GetTokenInfo(tokenName string) (*model.TokenInfo, error) {
	token, ok := model.MockTokenDB[tokenName]
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	return &token, nil
}
