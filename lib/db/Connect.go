package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoHost = "mongodb://localhost:27017"

func Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoHost))
	if err != nil {
		panic(err)
	}
	return client
}

/*

var finalResult1 bson.M
	var finalResult2 bson.M

	usersCollection := client.Database("get-projects--bot").Collection("users")
	user := bson.D{{"_id", "123212e2"}, {"fullName", "User 1"}, {"age", 30}} // {"fullName", "User 1"}, {"age", 30}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	usersCollection.FindOne(context.TODO(), bson.D{{"_id", "12"}}).Decode(&finalResult1)
	usersCollection.FindOne(context.TODO(), bson.D{{"_id", "23"}}).Decode(&finalResult2)

	// display the id of the newly inserted object
	fmt.Println(finalResult1)
	fmt.Println(finalResult2)
	fmt.Println(result.InsertedID)


*/
