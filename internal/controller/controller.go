package controller

import (
	"SimpleServer/internal/models"
	"SimpleServer/internal/providers/cache"
	"SimpleServer/pkg/userService"
)

type Controller struct {
	cache *cache.Cache
}

func New(cache *cache.Cache) *Controller {
	return &Controller{cache: cache}
}

func (c *Controller) GetEmployeeByEmail(email string) (*userService.Employee, error) {
	// checking if user exists
	info, err := c.cache.GetEmployeeInfoByEmail(email)
	if err != nil {
		return nil, err
	}

	return models.ConvertToEmployeeResponse(info), nil
}

func (c *Controller) DeleteEmployeeByEmail(email string) error {
	err := c.cache.DeleteByEmail(email)
	return err

}

func (c *Controller) AddEmployeeInCache(employee *userService.Employee) error {
	err := c.cache.Set(models.ConvertToEmployeeInfo(employee))
	return err
}
