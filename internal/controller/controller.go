package controller

import (
	"SimpleServer/internal/models"
	"SimpleServer/internal/providers/cache"
	"SimpleServer/pkg/userService"
)

type Controller struct {
	cache *cache.Cache
}

// controller constructor

func NewController(cache *cache.Cache) *Controller {
	return &Controller{cache: cache}
}

// sending request to cache to get info about employee

func (c *Controller) GetEmployeeByEmail(email string) (*userService.Employee, error) {
	// checking if user exists
	info, err := c.cache.GetEmployeeInfoByEmail(email)
	if err != nil {
		return nil, err
	}

	return models.ConvertToEmployeeResponse(info), nil
}

// sending request to cache to delete the employee

func (c *Controller) DeleteEmployeeByEmail(email string) error {
	// checking for valid key
	err := c.cache.DeleteByEmail(email)
	// checking for successful deleting
	if err != nil {
		return err
	}
	return nil

}

// sending request to cache to add employee info

func (c *Controller) AddEmployeeInCache(employee *userService.Employee) error {
	err := c.cache.Set(models.ConvertToEmployeeInfo(employee))
	if err != nil {
		return err
	}
	return nil
}
