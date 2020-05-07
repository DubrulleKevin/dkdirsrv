package test

import (
	"testing"

	"github.com/DubrulleKevin/dkdirsrv/interface/storage"
)

func TestGetNext(t *testing.T) {
	idr := storage.NewIDRepositoryImplMemory()

	id1, err := idr.GetNext()

	if err != nil {
		t.Error("GetNext method returned an error")
	}

	id2, err := idr.GetNext()

	if err != nil {
		t.Error("GetNext method returned an error")
	}

	if id1.Equals(id2) {
		t.Error("Repository generated 2 ID detected as equal")
	}

}
