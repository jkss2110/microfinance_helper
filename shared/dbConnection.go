package shared

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error close")
		}
	}()
}

// query method returns a cursor and error.
func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}

func DBConnection() {
	client, ctx, cancel, err := connect("mongodb+srv://jk:jk@microservicehelper.gzriw.mongodb.net/MicroFinance?retryWrites=true&w=majority")
	if err != nil {
		fmt.Println("Error")
	}
	// Free the resource when mainn dunction is  returned
	defer close(client, ctx, cancel)

	/*var filter, option interface{}
	// filter  gets all document
	filter = bson.D{}
	//  option remove id field from all documents
	option = bson.D{{"_id", 0}}
	// This method returns momngo.cursor and error if any.
	cursor, err := query(client, ctx, "MicroFinance", "loandetails", filter, option)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	// printing the result of query.
	fmt.Println("Query Result")
	for _, doc := range results {
		fmt.Println(doc)
	}*/
}
