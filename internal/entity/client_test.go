package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Jhon Doe", "jhon@doe.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Jhon Doe", client.Name)
	assert.Equal(t, "jhon@doe.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "jhon@doe.com")
	err := client.Update("Jhon Dooe", "jhon@dooe.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jhon Dooe", client.Name)
	assert.Equal(t, "jhon@dooe.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "jhon@doe.com")
	err := client.Update("", "jhon@dooe.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "jhon@doe.com")
	account := NewAccount(client)

	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
