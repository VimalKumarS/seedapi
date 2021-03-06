package service

import (
	"seedapi/model"
	"seedapi/util"
	repo "seedapi/repository"
	"time"
)

// SearchMapperSQL  define a config contain SQL mapper connection string
type SearchMapperSQL struct {
	config *model.AppSetting
}

// SearchServiceInstance instance
func SearchServiceInstance(config *model.AppSetting) *SearchMapperSQL {
	return &SearchMapperSQL{
		config: config,
	}
}

// GetAll search item
func (m *SearchMapperSQL) GetAll() ([]model.Search, error) {
	var search []model.Search
	// Todo: do some db operation
	//m.config.Connection
	repo.Data.QueryRow(`SELECT id, email, firstname, lastname FROM public.user WHERE email = $1   AND password = $2`, "email", "pwd")
	util.Logger.Info("GetAll method started")
	search = []model.Search{model.Search{ID: 1, Name: "name1", Date: time.Now()},
		model.Search{ID: 1, Name: "name1", Date: time.Now()},
	}

	return search, nil
}


func Add(nums ...int) int {
	var result int
	for _, i := range nums {
		result += i
	}

	return result
}
