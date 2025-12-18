package main

import (
	"github.com/dulisoft/spirit/tools/unique"
	"gorm.io/gorm"
)

type GenIDMethod struct {
	ID uint64
}

func (m *GenIDMethod) BeforeCreate(_ *gorm.DB) error {
	if m == nil || m.ID > 0 {
		return nil
	}
	m.ID = unique.NewSnowflakeID()
	return nil
}
