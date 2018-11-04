package service

import (
	"seedapi/model"
	"time"
	"go.uber.org/zap"
)

// SearchMapperSQL  define a config contain SQL mapper connection string
type SearchMapperSQL struct {
	config *model.AppSetting
	logger *zap.Logger
}

// SearchServiceInstance instance
func SearchServiceInstance(config *model.AppSetting, Logger *zap.Logger) *SearchMapperSQL {
	return &SearchMapperSQL{
		config: config,
		logger:Logger,
	}
}

// GetAll search item
func (m *SearchMapperSQL) GetAll() ([]model.Search, error) {
	var search []model.Search
	// Todo: do some db operation
	//m.config.Connection

	search = []model.Search{model.Search{ID: 1, Name: "name1", Date: time.Now()},
		model.Search{ID: 1, Name: "name1", Date: time.Now()},
	}

	return search, nil
}
