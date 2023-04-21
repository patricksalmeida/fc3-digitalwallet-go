package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Jhon Doe 1", "jhon1@doe.com")
	account1 := NewAccount(client1)

	client2, _ := NewClient("Jhon Doe 2", "jhon2@doe.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account1.Balance)
}

func TestCreateTransactionWithNoFunds(t *testing.T) {
	client1, _ := NewClient("Jhon Doe 1", "jhon1@doe.com")
	account1 := NewAccount(client1)

	client2, _ := NewClient("Jhon Doe 2", "jhon2@doe.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 2000)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "insuficient funds")
}
