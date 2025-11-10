package model

import (
	"encoding/json"
	"strings"

	"gorm.io/gorm"
)

type MaterialProvider struct {
	Name string `json:"name" gorm:"primaryKey"`
	Keys string `json:"keys" gorm:"type:text"`
}

func (m *MaterialProvider) GetKeys() ([]string, error) {
	var keys []string
	// If Keys is empty, return an empty slice instead of JSON unmarshal error
	if strings.TrimSpace(m.Keys) == "" {
		return []string{}, nil
	}
	err := json.Unmarshal([]byte(m.Keys), &keys)
	return keys, err
}

func (m *MaterialProvider) SetKeys(keys []string) error {
	data, err := json.Marshal(keys)
	if err != nil {
		return err
	}
	m.Keys = string(data)
	return nil
}

func GetMaterialProvider(name string) (*MaterialProvider, error) {
	var provider MaterialProvider
	err := DB.Where("name = ?", name).First(&provider).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Return an empty provider (no error) so callers can handle absence gracefully
			return &MaterialProvider{Name: name}, nil
		}
		return nil, err
	}
	return &provider, nil
}

func UpdateMaterialProvider(name string, keys []string) error {
	var provider MaterialProvider
	err := DB.Where("name = ?", name).First(&provider).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			provider = MaterialProvider{Name: name}
		} else {
			return err
		}
	}
	if err := provider.SetKeys(keys); err != nil {
		return err
	}
	return DB.Save(&provider).Error
}

func DeleteMaterialProvider(name string) error {
	return DB.Where("name = ?", name).Delete(&MaterialProvider{}).Error
}

func AllMaterialProviders() ([]*MaterialProvider, error) {
	var providers []*MaterialProvider
	err := DB.Find(&providers).Error
	return providers, err
}

// EnsureMaterialProvider ensures a provider row exists; if not, create with empty keys
func EnsureMaterialProvider(name string) error {
	var count int64
	if err := DB.Model(&MaterialProvider{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		provider := MaterialProvider{Name: name}
		if setErr := provider.SetKeys([]string{}); setErr != nil {
			return setErr
		}
		return DB.Create(&provider).Error
	}
	return nil
}