package test

import (
	"testing"

	"github.com/DubrulleKevin/dkdirsrv/interface/storage"
	"github.com/DubrulleKevin/dkdirsrv/model"
)

func TestAddUser(t *testing.T) {
	ur := storage.NewUserRepositoryImplMemory()
	idr := storage.NewIDRepositoryImplMemory()

	ur.AddUser(idr.GetNext(), model.NewUserModel())
}
