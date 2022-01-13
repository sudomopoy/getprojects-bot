package main

import (
	"labix.org/v2/mgo"
)

// func Connect() *mongo.Database {

// 	var client *mongo.Client
// 	var err error
// 	if GetProccessMode() == "product" {
// 		credential := options.Credential{
// 			Username: mongoUsername,
// 			Password: mongoPassword,
// 		}
// 		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoHost).SetAuth(credential))
// 	} else {
// 		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoHost))
// 	}
// 	check(err)
// 	return client.Database(mongoDatabase)
// }
func Connect() *mgo.Database {
	var session *mgo.Session
	var err error
	if GetProccessMode() == "product" {
		session, err = mgo.Dial(`mongodb://root:NG43ubjnbXjsxdWW3me699QyQCu7XW48@get-projects-bot-api.mohsen8.svc:27017/?authSource=admin&authMechanism=SCRAM-SHA-256&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false`)
	} else {
		session, err = mgo.Dial(mongoHost)
	}

	check(err)
	//error check on every access
	session.SetSafe(&mgo.Safe{})
	return session.DB(mongoDatabase)
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
