package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/yasuoza/clean-architecture-go/server/usecase/repository"
	"gorm.io/gorm"
)

type dbRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) repository.DBRepository {
	return &dbRepository{db}
}

func (r *dbRepository) Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx := r.db.Begin()
	if !errors.Is(tx.Error, nil) {
		return nil, tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			log.Println("recover")
			tx.Rollback()
			panic(p)
		} else if !errors.Is(err, nil) {
			log.Println("rollback")
			fmt.Println(err)
			tx.Rollback()
			panic("error")
		} else {
			err = tx.Commit().Error
		}
	}()

	return txFunc(tx)
}
