package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
)

func connect(uri string) error {
	var err error
	Ctx, Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	return err
}

func Close() {

	defer Cancel()
	defer func() {
		if err := Client.Disconnect(Ctx); err != nil {
			fmt.Println("Error close")
		}
	}()
}

func DBConnection() {
	err := connect("mongodb+srv://jk:jk@microservicehelper.gzriw.mongodb.net/MicroFinance?retryWrites=true&w=majority")
	if err != nil {
		panic("Error")
	}
}

// query method returns a cursor and error.
func QueryDB(ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := Client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}
func QueryTopDB(ctx context.Context, dataBase, col string, query, sorter interface{}, limit int64) (result *mongo.Cursor, err error) {
	collection := Client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetSort(sorter).SetLimit(limit))
	return
}
func InsertOneDB(ctx context.Context, dataBase, col string, loanInfo interface{}) (result *mongo.InsertOneResult, err error) {
	collection := Client.Database(dataBase).Collection(col)
	result, err = collection.InsertOne(ctx, loanInfo)
	return result, err
}
