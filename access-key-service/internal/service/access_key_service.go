package service

import (
	"access-key-service/internal/model"
	"access-key-service/internal/pubsub"
	"access-key-service/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AccessKeyService struct {
	repo      *repository.AccessKeyRepository
	publisher *pubsub.RedisPublisher
}

func NewAccessKeyService(r *repository.AccessKeyRepository, p *pubsub.RedisPublisher) *AccessKeyService {
	return &AccessKeyService{repo: r, publisher: p}
}

func (s *AccessKeyService) CreateKey(rateLimit int, expiryMinutes int) (*model.AccessKey, error) {
	key := &model.AccessKey{
		Key:       uuid.NewString(),
		RateLimit: rateLimit,
		ExpiresAt: time.Now().Add(time.Duration(expiryMinutes) * time.Minute),
	}
	if err := s.repo.Create(key); err != nil {
		return nil, err
	}
	s.publisher.PublishKeyUpdate(*key)
	return key, nil
}

func (s *AccessKeyService) ListKeys() ([]model.AccessKey, error) {
	return s.repo.GetAll()
}

func (s *AccessKeyService) UpdateKey(key string, rateLimit int, expiryMinutes int) error {
	k, err := s.repo.Get(key)
	if err != nil {
		return err
	}
	k.RateLimit = rateLimit
	k.ExpiresAt = time.Now().Add(time.Duration(expiryMinutes) * time.Minute)
	if err := s.repo.Update(k); err != nil {
		return err
	}
	s.publisher.PublishKeyUpdate(*k)
	return nil
}

func (s *AccessKeyService) DeleteKey(key string) error {
	err := s.repo.Delete(key)
	if err == nil {
		s.publisher.PublishKeyUpdate(model.AccessKey{Key: key})
	}
	return err
}

func (s *AccessKeyService) GetPlan(key string) (*model.AccessKey, error) {
	return s.repo.Get(key)
}

func (s *AccessKeyService) DisableKey(key string) error {
	k, err := s.repo.Get(key)
	if err != nil {
		return err
	}
	k.Disabled = true
	if err := s.repo.Update(k); err != nil {
		return err
	}
	s.publisher.PublishKeyUpdate(*k)
	return nil
}
