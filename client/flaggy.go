package client

import (
	"encoding/json"
	"fmt"
	"github.com/captechtimmy/flaggy/internal/model"
	"net/http"
	"time"
)

type FlaggyClient struct {
	baseEndpoint string
}

func NewFlaggyClient(baseEndpoint string) *FlaggyClient {
	return &FlaggyClient{baseEndpoint: baseEndpoint}
}

func (c *FlaggyClient) GetFlag(key string) (model.FeatureFlag, error) {
	client := http.Client{Timeout: time.Duration(10) * time.Duration(time.Second)}
	r, _ := http.NewRequest("GET", fmt.Sprintf("%s/featureflag/%s", c.baseEndpoint, key), nil)
	r.Header.Add("Content-Type", "application/json")
	res, err := client.Do(r)
	if err != nil {
		return model.FeatureFlag{}, err
	}
	enc := json.NewDecoder(res.Body)
	var f model.FeatureFlag
	err = enc.Decode(&f)
	if err != nil {
		return model.FeatureFlag{}, err
	}
	return f, nil
}
