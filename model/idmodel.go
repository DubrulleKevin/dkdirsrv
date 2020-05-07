package model

// IDModel strcuture
type IDModel struct {
	id int64
}

// NewIDModel function
func NewIDModel(id int64) *IDModel {
	return &IDModel{
		id: id,
	}
}

// Equals method
func (im *IDModel) Equals(otherIDModel *IDModel) bool {
	return im.id == otherIDModel.id
}
