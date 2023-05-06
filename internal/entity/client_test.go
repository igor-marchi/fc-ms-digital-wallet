package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatedNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenInvalidInput(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Error(t, err, "name is required")
	assert.Error(t, err, "email is required")
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("John Doe updated", "j@jupdated.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe updated", client.Name)
	assert.Equal(t, "j@jupdated.com", client.Email)
}

func TestUpdateClientWithInvalidInput(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "")
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestAddAccountToCLient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
