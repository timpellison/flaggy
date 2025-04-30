package persistence

import (
	"errors"
	"github.com/timpellison/flaggy/internal/model"
)

type InMemoryStore struct {
	mapping map[string]*model.FeatureFlag
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{mapping: make(map[string]*model.FeatureFlag)}
}

func (s *InMemoryStore) GetFlag(key string) (*model.FeatureFlag, error) {
	if f, ok := s.mapping[key]; !ok {
		return nil, errors.New("feature flag does not exist")
	} else {
		return f, nil
	}
}

func (s *InMemoryStore) Add(featureFlag model.FeatureFlag) error {
	if s.mapping[featureFlag.Key] != nil {
		return errors.New("feature flag already exists")
	}

	s.mapping[featureFlag.Key] = &featureFlag
	return nil
}

func (s *InMemoryStore) Update(featureFlag model.FeatureFlag) error {
	if _, ok := s.mapping[featureFlag.Key]; !ok {
		return errors.New("feature flag does not exist")
	}

	s.mapping[featureFlag.Key] = &featureFlag
	return nil
}

func (s *InMemoryStore) Delete(featureFlag model.FeatureFlag) error {
	if _, ok := s.mapping[featureFlag.Key]; !ok {
		return errors.New("feature flag does not exist")
	}
	delete(s.mapping, featureFlag.Key)
	return nil
}
