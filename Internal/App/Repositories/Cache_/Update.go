package Cache_

import (
	"SimpleServer/Internal/App/Models"
	"errors"
)

func (c *Cache) Update(key string, age int, sex bool) error {
	c.Lock()
	defer c.Unlock()
	value, ok := c.items[key]
	// updates info about user
	if !ok {
		c.Unlock()
		return errors.New("couldn't find the user")
		// if user was not found, return error
	} else {
		if person, ok := value.Value.(*Models.Person); ok {
			person.Age = age
			person.Sex = sex
		}
	}

	return nil
}
