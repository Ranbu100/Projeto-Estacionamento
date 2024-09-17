package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Nome string
}
