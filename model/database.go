// Package model has configuration about
// database as well
package model

import mgo "gopkg.in/mgo.v2"

// Database is a wrapper of the connection
type Database struct {
	Conn *mgo.Session
}

// NewDatabase returns a wrapper of api
func NewDatabase() (*Database, error) {
	db, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &Database{Conn: db}, nil
}

// Close quit the database
func (d Database) Close() {
	d.Conn.Close()
}

// Clone returns the new instance
func (d Database) Clone() *Database {
	return &Database{Conn: d.Conn.Clone()}
}

// GetCollection returns the mongoDB collection
func (d Database) GetCollection(database string, collection string) *mgo.Collection {
	return d.Conn.DB(database).C(collection)
}
