package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

func DBSetAdmin(id int) bool {
	ClCTX := Connect().Collection(collection_user)
	user := bson.D{{"role", "admin"}, {"admin-token", HashGen(token)}} // {"fullName", "User 1"}, {"age", 30}
	ClCTX.UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", user}})
	return true
}

func IsAdmin(id int) bool {
	result := false
	if !CollectionIsEmpty(collection_user) {
		ClCTX := Connect().Collection(collection_user)
		var userM bson.M
		if err := ClCTX.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&userM); err != nil {
			log.Fatal(err)
		}
		result = userM["role"] == "admin"
	}
	return result
}
func SetUserBaseInfoIfNotExists(id int) bool {
	result := false
	ClCTX := Connect().Collection(collection_user)
	user := bson.D{{"_id", id}} // {"fullName", "User 1"}, {"age", 30}
	cu, err := ClCTX.CountDocuments(context.TODO(), user)
	if err != nil {
		return false
	}
	if cu == 0 {
		user := bson.D{{"_id", id}, {"role", "user"}, {"phone", "NotSet"}} // {"fullName", "User 1"}, {"age", 30}
		ClCTX.InsertOne(context.TODO(), user)
	}
	return result
}

func CollectionIsEmpty(clName string) bool {
	ClCTX := Connect().Collection(clName)
	Count, _ := ClCTX.CountDocuments(context.TODO(), bson.D{})
	return int(Count) == 0
}
func AddNewProjectBaseInfo(title string, description string, userId int, username string) (bool, string) {
	ClCTX := Connect().Collection(collection_projects)
	pjId := idGenarator()
	project := bson.D{{"_id", pjId}, {"userId", userId}, {"title", title}, {"description", description}, {"status", "pending"}, {"username", username}} // {"fullName", "User 1"}, {"age", 30}
	ClCTX.InsertOne(context.TODO(), project)
	return true, pjId
}
func GetAdminsIds() []int {
	ClCTX := Connect().Collection(collection_user)
	cursor, err := ClCTX.Find(context.TODO(), bson.D{{"role", "admin"}})
	if err != nil {
		log.Fatal(err)
	}
	var admins []bson.M
	if err = cursor.All(context.TODO(), &admins); err != nil {
		log.Fatal(err)
	}
	var res []int
	for i := 0; i < len(admins); i++ {
		if reflect.TypeOf(admins[i]["_id"]).Kind() == reflect.Int32 {
			res = append(res, int(admins[i]["_id"].(int32)))
		} else {
			res = append(res, int(admins[i]["_id"].(int64)))
		}
	}
	return res
}
func GetFilterUser(filter bson.D) []bson.M {
	ClCTX := Connect().Collection(collection_user)
	cursor, err := ClCTX.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}
	return users
}

func DBUpdateSingleProject(pjId string, NewArgs bson.D) bson.M {
	ClCTX := Connect().Collection(collection_projects)
	fmt.Println(pjId)
	ClCTX.UpdateOne(context.TODO(), bson.D{{"_id", pjId}}, bson.D{{"$set", NewArgs}})
	var prjRes bson.M
	if err := ClCTX.FindOne(context.TODO(), bson.D{{"_id", pjId}}).Decode(&prjRes); err != nil {
		log.Fatal(err)
	}
	return prjRes
}
