package repository

import "github.com/DubrulleKevin/dkdirsrv/model"

// IDRepository interface
type IDRepository interface {
	GetNext() (*model.IDModel, error)
}
