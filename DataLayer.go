package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

func RedisClientSet(id int, step string) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	err := rdb.Set(ctx, strconv.Itoa(id), step, 0).Err()
	if err != nil {
		return false
	}
	return true
}
func RedisClientGet(id int) (string, bool) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	val, err := rdb.Get(ctx, strconv.Itoa(id)).Result()
	if err == redis.Nil {
		return "", true
	} else if err != nil {
		return "", true
	} else {
		return val, false
	}
}
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
func AddNewProjectBaseInfo(title string, description string, userId int, username string) (bool, int) {
	ClCTX := Connect().Collection(collection_projects)
	pjId := rand.Intn(1000000000000000)
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
		res = append(res, int(admins[i]["_id"].(int32)))
	}
	return res
}

func DBUpdateSingleProject(pjId int, NewArgs bson.D) bson.M {
	ClCTX := Connect().Collection(collection_projects)
	fmt.Println(pjId)
	ClCTX.UpdateOne(context.TODO(), bson.D{{"_id", pjId}}, bson.D{{"$set", NewArgs}})
	var prjRes bson.M
	if err := ClCTX.FindOne(context.TODO(), bson.D{{"_id", pjId}}).Decode(&prjRes); err != nil {
		log.Fatal(err)
	}
	return prjRes
}
