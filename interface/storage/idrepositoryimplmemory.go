package storage

import (
	"sync"

	"github.com/DubrulleKevin/dkdirsrv/model"
)

// IDRepositoryImplMemory structure
type IDRepositoryImplMemory struct {
	mutex sync.Mutex
	ids   []*model.IDModel
}

// NewIDRepositoryImplMemory function
func NewIDRepositoryImplMemory() *IDRepositoryImplMemory {
	return &IDRepositoryImplMemory{}
}

// GetNext method
func (idr *IDRepositoryImplMemory) GetNext() (*model.IDModel, error) {
	idr.mutex.Lock()
	defer idr.mutex.Unlock()

	next := int64(len(idr.ids))

	newID := model.NewIDModel(next)
	idr.ids = append(idr.ids, newID)

	return newID, nil
}
