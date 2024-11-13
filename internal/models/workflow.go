package models

import (
  "gorm.io/gorm"
)

type Workflows struct {
  gorm.Model
  Name string `gorm:"size:255"`
  FileName string `gorm: `
}