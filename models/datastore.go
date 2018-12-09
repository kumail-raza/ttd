package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//DataStore DataStore
type DataStore interface {
	GetPerfumes() ([]Perfume, error)
}

type dataSource struct {
	ConnectionString string
}

//NewDataStore NewDataStore
func NewDataStore(connString string) DataStore {
	return &dataSource{ConnectionString: connString}
}
func (ds *dataSource) GetPerfumes() ([]Perfume, error) {

	session, err := mgo.Dial(ds.ConnectionString)
	if err != nil {
		return nil, err
	}
	var p []Perfume
	defer session.Close()
	return p, session.DB("db").C("perfumes").Find(bson.M{}).All(&p)

}
