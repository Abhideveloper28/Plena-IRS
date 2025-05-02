package repository

import (
	"access-key-service/internal/model"

	"gorm.io/gorm"
)

type AccessKeyRepository struct {
	db *gorm.DB
}

func NewAccessKeyRepository(db *gorm.DB) *AccessKeyRepository {
	return &AccessKeyRepository{db: db}
}

func (r *AccessKeyRepository) Create(key *model.AccessKey) error {
	return r.db.Create(key).Error
}

func (r *AccessKeyRepository) GetAll() ([]model.AccessKey, error) {
	var keys []model.AccessKey
	err := r.db.Find(&keys).Error
	return keys, err
}

func (r *AccessKeyRepository) Get(key string) (*model.AccessKey, error) {
	var k model.AccessKey
	err := r.db.First(&k, "key = ?", key).Error
	return &k, err
}

func (r *AccessKeyRepository) Update(key *model.AccessKey) error {
	return r.db.Save(key).Error
}

func (r *AccessKeyRepository) Delete(key string) error {
	return r.db.Delete(&model.AccessKey{}, "key = ?", key).Error
}
