package models

import (
  "gorm.io/gorm"
)

type History struct {
  gorm.Model
  Id uint
  WorkflowID uint 
  EventID uint "commit"
  UserID uint "alexditz"
  LaunchTime uint
}