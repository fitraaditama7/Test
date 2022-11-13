package entity

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	ID                uint           `gorm:"column:id;primaryKey"`
	Name              string         `gorm:"column:name"`
	Address           string         `gorm:"column:address"`
	Email             string         `gorm:"column:email"`
	Password          string         `gorm:"column:password"`
	Photos            pq.StringArray `gorm:"type:text[];column:photos"`
	CreditCardType    string         `gorm:"column:credit_card_type"`
	CreditCardNumber  string         `gorm:"column:credit_card_number"`
	CreditCardName    string         `gorm:"column:credit_card_name"`
	CreditCardExpired string         `gorm:"column:credit_card_expired"`
	CreditCardCVV     string         `gorm:"column:credit_card_cvv"`
	CreatedAt         time.Time      `gorm:"created_at"`
	UpdatedAt         *time.Time     `gorm:"updated_at"`
}
