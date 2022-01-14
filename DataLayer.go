package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var todoCTX = context.TODO()

func CreateSingleProject(newProject SingleProjectModel) bool {
	Collection := Connect().Collection(collection_project)
	_, err := Collection.InsertOne(todoCTX, newProject)
	check(err)
	return true
}
func GetSingleProject(filterProject bson.M) SingleProjectModel {
	Collection := Connect().Collection(collection_project)
	var selectedProject SingleProjectModel
	Collection.FindOne(todoCTX, filterProject).Decode(&selectedProject)
	return selectedProject
}
func GetFilteredProjects(filterProject SingleProjectModel) []SingleProjectModel {
	Collection := Connect().Collection(collection_project)
	cursor, _ := Collection.Find(todoCTX, filterProject)
	var selectedProjects []SingleProjectModel
	err := cursor.All(todoCTX, &selectedProjects)
	check(err)
	return selectedProjects
}
func SetUpdateSingleProject(filterProject bson.M, updateTo SingleProjectModel) bool {
	Collection := Connect().Collection(collection_project)
	_, err := Collection.UpdateOne(todoCTX, filterProject, bson.D{{Key: "$set", Value: updateTo}})
	check(err)
	return true
}
func CreateSingleUser(newUser SingleUserModel) bool {
	Collection := Connect().Collection(collection_user)
	_, err := Collection.InsertOne(todoCTX, &newUser)
	check(err)
	return true
}
func GetSingleUser(filterUser bson.M) SingleUserModel {
	Collection := Connect().Collection(collection_user)
	var selectedUser SingleUserModel
	Collection.FindOne(todoCTX, filterUser).Decode(&selectedUser)
	return selectedUser
}
func GetFilteredUsers(filterUser bson.M) []SingleUserModel {
	Collection := Connect().Collection(collection_user)
	cursor, err := Collection.Find(todoCTX, filterUser)
	var selectedUsers []SingleUserModel
	err = cursor.All(todoCTX, &selectedUsers)
	check(err)
	return selectedUsers
}
func SetUpdateSingleUser(filterUser bson.M, updateTo SingleUserModel) bool {
	Collection := Connect().Collection(collection_user)
	_, err := Collection.UpdateOne(todoCTX, filterUser, bson.D{{Key: "$set", Value: updateTo}})
	return check(err)
}
