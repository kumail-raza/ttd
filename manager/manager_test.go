package manager

import (
	"testing"

	"github.com/minhajuddinhkhan/ttd/models"
)

type mockDataStore struct{}

func (md *mockDataStore) GetPerfumes() ([]models.Perfume, error) {
	return []models.Perfume{
		models.Perfume{Name: "hey", Type: "hey"},
	}, nil

}

func TestGetPerfumes(t *testing.T) {

	p, err := NewManager(&mockDataStore{}).GetPerfumes()
	if err != nil {
		t.Error(err)
	}
	if len(p) == 0 {
		t.Error("should be one")
	}
}
