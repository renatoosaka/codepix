package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindKeyByKind(key, kind string) (*PixKey, error)
	FindAccount(id string) (*Account, error)
	FindBank(id string) (*Bank, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" gorm:"type:varchar(20)" valid:"notnull"`
	Key       string   `json:"key" gorm:"type:varchar(255)" valid:"notnull"`
	AccountID string   `json:"-" gorm:"column:account_id;type:uuid;not null" valid:"notnull"`
	Account   *Account `json:"account" valid:"-"`
	Status    string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	if _, err := govalidator.ValidateStruct(pixKey); err != nil {
		return err
	}

	if pixKey.Kind != "cpf" && pixKey.Kind != "email" {
		return errors.New("Invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("Invalid type of status")
	}

	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Account: account,
		Key:     key,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	pixKey.UpdatedAt = time.Now()

	if err := pixKey.isValid(); err != nil {
		return nil, err
	}

	return &pixKey, nil
}
