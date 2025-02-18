package repository

import "avito_internship_backend/internal/entity"

type Account interface {
	Deposit(id, amount int) error
	GetBalance(id int) (*entity.Balance, error)
}
