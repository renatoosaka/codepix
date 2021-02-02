package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account `json:"accounts" gorm:"ForeignKey:BankID" valid:"-"`
}

func (bank *Bank) isValid() error {
	if _, err := govalidator.ValidateStruct(bank); err != nil {
		return err
	}

	return nil
}

func NewBank(code, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	bank.UpdatedAt = time.Now()

	if err := bank.isValid(); err != nil {
		return nil, err
	}

	return &bank, nil
}
