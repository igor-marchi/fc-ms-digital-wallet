package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, err
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}