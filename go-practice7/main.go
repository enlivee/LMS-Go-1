package main

import "errors"

var (
	valueError = errors.New("Значение не может быть меньше 0!")
	depositError = errors.New("Значение должно быть больше 0!")
	balanceError = errors.New("Баланс должен быть больше 0!")
)

type Account struct {
	balance float64
	owner string
}

func NewAccount(owner string) *Account {
	return &Account{balance: 0, owner: owner}
}

func (a *Account) SetBalance(value float64) error{
	if value < 0 {
		return valueError
	}
	a.balance = value
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(value float64) error {
	if value <= 0 {
		return depositError
	}
	a.balance += value
	return nil
}

func (a *Account) Withdraw(value float64) error{
	if value <= 0 {
		return depositError
	}
	a.balance -= value
	if a.balance <= 0 {
		a.balance += value
		return balanceError
	}
	return nil
}