package yookassa

import (
	"encoding/json"
	"fmt"
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"go.uber.org/zap"
	"net/http"
)

type PayoutsSvc struct {
	*Client
	log *zap.SugaredLogger
}

func NewPayoutsService(client *Client, log *zap.SugaredLogger) *PayoutsSvc {
	return &PayoutsSvc{
		client, log,
	}
}

func (h *PayoutsSvc) MakePayout(payout *yoomodel.Payout, idempotencyKey string) (*http.Response, error) {
	jsonData, err := json.Marshal(payout)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payout: %s", err)
	}

	res, err := h.makeRequest(http.MethodPost, PayoutEndpoint, "", jsonData, nil, idempotencyKey)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *PayoutsSvc) GetPayoutInfo(id string) (*http.Response, error) {
	res, err := h.makeRequest(http.MethodGet, PayoutEndpoint, id, nil, nil, "")
	if err != nil {
		return nil, err
	}

	return res, nil
}
