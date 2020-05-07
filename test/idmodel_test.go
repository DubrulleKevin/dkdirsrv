package test

import (
	"testing"

	"github.com/DubrulleKevin/dkdirsrv/model"
)

func TestEquals(t *testing.T) {
	id1 := model.NewIDModel(0)
	id2 := model.NewIDModel(1)
	id3 := model.NewIDModel(0)

	if id1.Equals(id2) {
		t.Error("ID detected equals but are not")
	}

	if !id1.Equals(id3) {
		t.Error("ID detected not equals but are")
	}

}
