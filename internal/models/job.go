package models

import (
  "gorm.io/gorm"
)

type Job struct {
  gorm.Model
  ID uint 
  Name string `gorm:"type:varchar(100)" json:"name"`
  Duration uint `json:"name"`
  ErrorMessage string `gorm:"type:text" json:"errorMessage"`
  RunID uint 
  StatusID uint 
} 