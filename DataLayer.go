package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var todoCTX = context.TODO()

func CreateSingleProject(newProject SingleProjectModel) bool {
	Collection := Connect().C(collection_project)
	err := Collection.Insert(newProject)
	check(err)
	return true
}
func GetSingleProject(filterProject bson.M) SingleProjectModel {
	Collection := Connect().C(collection_project)
	var selectedProject SingleProjectModel
	Collection.Find(filterProject).One(&selectedProject)
	return selectedProject
}
func GetFilteredProjects(filterProject SingleProjectModel) []SingleProjectModel {
	Collection := Connect().C(collection_project)
	cursor := Collection.Find(filterProject)
	var selectedProjects []SingleProjectModel
	err := cursor.All(&selectedProjects)
	check(err)
	return selectedProjects
}
func SetUpdateSingleProject(filterProject bson.M, updateTo SingleProjectModel) bool {
	Collection := Connect().C(collection_project)
	err := Collection.Update(filterProject, updateTo)
	check(err)
	return true
}
func CreateSingleUser(newUser SingleUserModel) bool {
	Collection := Connect().C(collection_user)
	err := Collection.Insert(&newUser)
	check(err)
	return true
}
func GetSingleUser(filterUser bson.M) SingleUserModel {
	Collection := Connect().C(collection_user)
	var selectedUser SingleUserModel
	Collection.Find(filterUser).One(&selectedUser)
	return selectedUser
}
func GetFilteredUsers(filterUser bson.M) []SingleUserModel {
	Collection := Connect().C(collection_user)
	cursor := Collection.Find(filterUser)
	var selectedUsers []SingleUserModel
	err := cursor.All(&selectedUsers)
	check(err)
	return selectedUsers
}
func SetUpdateSingleUser(filterUser bson.M, updateTo SingleUserModel) bool {
	Collection := Connect().C(collection_user)
	err := Collection.Update(filterUser, updateTo)
	return check(err)
}
