package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/renatoosaka/codepix/domain/model"
)

type TransactionRepositoryDB struct {
	DB *gorm.DB
}

func (r *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	if err := r.DB.Create(transaction).Error; err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	if err := r.DB.Save(transaction).Error; err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.DB.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, errors.New("no transaction was found")
	}

	return &transaction, nil
}
