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
func GetSingleProject(filterProject SingleProjectModel) SingleProjectModel {
	Collection := Connect().Collection(collection_project)
	var selectedProject SingleProjectModel
	err := Collection.FindOne(todoCTX, filterProject).Decode(&selectedProject)
	check(err)
	return selectedProject
}
func GetFilteredProjects(filterProject SingleProjectModel) []SingleProjectModel {
	Collection := Connect().Collection(collection_project)
	cursor, err := Collection.Find(todoCTX, filterProject)
	check(err)
	var selectedProjects []SingleProjectModel
	err = cursor.All(todoCTX, &selectedProjects)
	check(err)
	return selectedProjects
}
func SetUpdateSingleProject(filterProject SingleProjectModel, updateTo SingleProjectModel) bool {
	Collection := Connect().Collection(collection_project)
	_, err := Collection.UpdateOne(todoCTX, filterProject, bson.D{{"$set", updateTo}})
	check(err)
	return true
}
func CreateSingleUser(newUser SingleUserModel) bool {
	Collection := Connect().Collection(collection_user)
	_, err := Collection.InsertOne(todoCTX, newUser)
	check(err)
	return true
}
func GetSingleUser(filterUser SingleUserModel) (SingleUserModel, bool) {
	Collection := Connect().Collection(collection_user)
	var selectedUser SingleUserModel
	err := Collection.FindOne(todoCTX, filterUser).Decode(&selectedUser)
	return selectedUser, check(err)
}
func GetFilteredUsers(filterUser SingleUserModel) []SingleUserModel {
	Collection := Connect().Collection(collection_user)
	cursor, err := Collection.Find(todoCTX, filterUser)
	check(err)
	var selectedUsers []SingleUserModel
	err = cursor.All(context.TODO(), &selectedUsers)
	check(err)
	return selectedUsers
}
func SetUpdateSingleUser(filterUser SingleUserModel, updateTo SingleUserModel) bool {
	Collection := Connect().Collection(collection_user)
	_, err := Collection.UpdateOne(todoCTX, filterUser, bson.D{{"$set", updateTo}})
	check(err)
	return true
}
