package repository

import ()

type NoSqlRepository interface {
	Connector
	Reader
	Writer
}

type Connector interface {
	Open() error
	Close() error
}

type Reader interface {
	GetById(id int) (interface{}, error)
	GetAll() ([]interface{}, error)
}

type Writer interface {
	Insert(v interface{}) error
	Update(v interface{}) error
}