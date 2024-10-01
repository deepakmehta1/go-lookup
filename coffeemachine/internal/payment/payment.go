package payment

import (
	"errors"
	"sync"
)

// PaymentProcessor handles cash transactions.
type PaymentProcessor struct {
	balance float64
	mu      sync.Mutex
}

// NewPaymentProcessor creates a new payment processor instance.
func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{}
}

// AcceptCash accepts cash input from the user.
func (pp *PaymentProcessor) AcceptCash(amount float64) {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	pp.balance += amount
}

// VerifyAmount checks if the inserted cash is sufficient.
func (pp *PaymentProcessor) VerifyAmount(price float64) error {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	if pp.balance < price {
		return errors.New("insufficient cash")
	}
	pp.balance -= price
	return nil
}

// Refund refunds the remaining balance.
func (pp *PaymentProcessor) Refund() float64 {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	refund := pp.balance
	pp.balance = 0
	return refund
}
