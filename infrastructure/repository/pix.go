package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/renatoosaka/codepix/domain/model"
)

type PixKeyRepositoryDB struct {
	DB *gorm.DB
}

// RegisterKey(pixKey *PixKey) (*PixKey, error)
// AddBank(bank *Bank) error
// AddAccount(account *Account) error
// FindKeyByKind(key, kind string) (*PixKey, error)
// FindAccount(id string) (*Account, error)
// FindBank(id string) (*Bank, error)

func (r PixKeyRepositoryDB) AddBank(bank *model.Bank) error {
	err := r.DB.Create(bank).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDB) AddAccount(account *model.Account) error {
	err := r.DB.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.DB.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r PixKeyRepositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey
	r.DB.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pixKey, nil
}

func (r PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	r.DB.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}
	return &account, nil
}

func (r PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.DB.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}
	return &bank, nil
}
