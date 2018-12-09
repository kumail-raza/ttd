package manager

import (
	"github.com/minhajuddinhkhan/ttd/models"
)

//Manager Manager
type Manager interface {
	GetPerfumes() ([]models.Perfume, error)
}

type manager struct {
	DataStore models.DataStore
}

func (m *manager) GetPerfumes() ([]models.Perfume, error) {
	return m.DataStore.GetPerfumes()
}

//NewManager creates a new manager
func NewManager(ds models.DataStore) Manager {
	return &manager{ds}
}
