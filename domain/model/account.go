package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `json:"bank" valid:"-"`
	BankID    string    `json:"-" gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*PixKey `json:"pix_keys" gorm:"ForeignKey:AccountID" valid:"-"`
}

func (account *Account) isValid() error {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		return err
	}

	return nil
}

func NewAccount(ownerName, number string, bank *Bank) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Number:    number,
		Bank:      bank,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	if err := account.isValid(); err != nil {
		return nil, err
	}

	return &account, nil
}
