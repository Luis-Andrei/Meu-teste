package db

import (
	"errors"
	"sync"

	"bank-server/models"
)

var (
	accounts = make(map[int]*models.Account)
	mutex    sync.RWMutex
	nextID   = 1
)

// Connect simula uma conexão com banco de dados
func Connect() error {
	return nil
}

// CreateAccount cria uma nova conta
func CreateAccount(name string, balance float64) (*models.Account, error) {
	mutex.Lock()
	defer mutex.Unlock()

	acc := &models.Account{
		ID:      nextID,
		Name:    name,
		Balance: balance,
	}
	accounts[nextID] = acc
	nextID++
	return acc, nil
}

// GetAccount retorna uma conta pelo ID
func GetAccount(id int) (*models.Account, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	acc, exists := accounts[id]
	if !exists {
		return nil, errors.New("conta não encontrada")
	}
	return acc, nil
}

// UpdateBalance atualiza o saldo de uma conta
func UpdateBalance(id int, amount float64) (*models.Account, error) {
	mutex.Lock()
	defer mutex.Unlock()

	acc, exists := accounts[id]
	if !exists {
		return nil, errors.New("conta não encontrada")
	}

	if amount < 0 && acc.Balance+amount < 0 {
		return nil, errors.New("saldo insuficiente")
	}

	acc.Balance += amount
	return acc, nil
}
