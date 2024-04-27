package domain

import (
	vo "github.com/axel-andrade/finance_planner_api/internal/core/domain/value_objects"
)

type Category struct {
	Base
	Name vo.Name `json:"name"`
}

func BuilCategory(name string) (*Category, error) {
	c := &Category{
		Name: vo.Name{Value: name},
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) validate() error {
	if err := c.Name.Validate(); err != nil {
		return err
	}

	return nil
}
