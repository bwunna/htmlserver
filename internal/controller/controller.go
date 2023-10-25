package controller

import (
	"SimpleServer/internal/models"
	"SimpleServer/internal/providers/cache"
	"fmt"
	"time"
)

type Controller struct {
	cache *cache.Cache
}

func NewController(cache *cache.Cache) *Controller {
	// constructor for controller
	return &Controller{cache: cache}
}

func (c *Controller) UpdateSalary(name string) error {
	// updating salary for user
	err := c.cache.AskForPromotion(name)
	if err != nil {
		return err
	}

	return nil

}

func (c *Controller) GetUser(name string) (string, error) {
	// checking if user exists
	userData, err := c.cache.Get(name)
	if err != nil {
		return "", err
	}

	salaryData, err := c.cache.GetSalaryData(name)
	if err != nil {
		return "", err
	}
	info := fmt.Sprint(*userData, salaryData)

	return info, nil
}

func (c *Controller) DeleteUser(name string) error {
	// checking for valid key
	err := c.cache.Delete(name)
	// checking for successful deleting
	if err != nil {
		return err
	}
	return nil

}

func (c *Controller) AddUser(user *models.User) error {
	// trying to add user to cache or update info about him
	err := c.cache.Set(user, time.Minute*10)
	if err != nil {
		// checking for errors while adding the user
		if err = c.cache.Update(user); err != nil {
			return err
		}
	}
	return nil

}
