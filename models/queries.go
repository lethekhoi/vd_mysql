package models

import "vd_mysql/entities"

type Queries interface {
	FindAll() (product []entities.Product, err error)
}
