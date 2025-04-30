package model

type FeatureFlag struct {
	Key   string `json:"key"`
	Value bool   `json:"enabled"`
}
