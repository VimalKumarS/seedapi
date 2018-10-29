package serviceinterface

import "seedapi/model"

// SearchInterface define the base contract for mapper
type SearchInterface interface {
	GetAll() ([]model.Search, error)
}
