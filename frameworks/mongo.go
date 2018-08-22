package frameworks

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
)

type MongoDB struct {
	client mongo.Client
	uri string
	database mongo.Database
}

func NewMongoDB(username, password, url, databaseName string) (MongoDB, error) {
	uri := "mongodb://" + username + ":" + password + url
	client, err := mongo.NewClient(uri)
	database :=  client.Database(databaseName)
	return MongoDB{client, uri, database}, err
}

func (m *MongoDB) Open() error {
	return m.client.Connect(context.Background())
}

func (m *MongoDB) Close() error {
	return m.client.Disconnect(context.Background())
}

func (m *MongoDB) GetById(id int) (interface{}, error) {
	return interface{}{}, nil
}

func (m *MongoDB) GetAll() ([]interface{}, error) {
	return interface{}{}, nil
}

func (m *MongoDB) Insert(v interface{}) error {
	return nil
}

func (m *MongoDB) Update(v interface{}) error {
	return nil
}

func (m *MongoDB) GetCollection(name string) {
	return m.database.Collection(name)
}