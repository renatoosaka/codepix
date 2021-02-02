package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `json:"account_from" valid:"-"`
	AccountFromID     string   `json:"-" gorm:"column:account_from_id;type:uuid;not null" valid:"notnull"`
	Amount            float64  `json:"ammount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `json:"pix_key_to" valid:"-"`
	PixKeyIDTo        string   `json:"-" gorm:"column:pix_key_id;type:uuid;not null" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (transaction *Transaction) isValid() error {
	if _, err := govalidator.ValidateStruct(transaction); err != nil {
		return err
	}

	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater than 0 (zero)")
	}

	if transaction.Status != TransactionPending && transaction.Status != TransactionError &&
		transaction.Status != TransactionConfirmed && transaction.Status != TransactionCompleted {
		return errors.New("invalid status for the transaction")
	}

	if transaction.PixKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	return nil
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.UpdatedAt = time.Now()

	return transaction.isValid()
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()

	return transaction.isValid()
}

func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TransactionError
	transaction.UpdatedAt = time.Now()
	transaction.CancelDescription = description

	return transaction.isValid()
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Description: description,
		Status:      TransactionPending,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	if err := transaction.isValid(); err != nil {
		return nil, err
	}

	return &transaction, nil
}
