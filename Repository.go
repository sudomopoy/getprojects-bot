package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SetUserBaseInfoIfNotExists(userBaseInfo SingleUserModel) (SingleUserModel, bool) {
	selectUserFilter := bson.M{"_id": userBaseInfo.ID}
	userInfo := GetSingleUser(selectUserFilter)

	if userInfo.ID == "" {
		userInfo.ID = userBaseInfo.ID
		userInfo.FirstName = userBaseInfo.FirstName
		userInfo.LastName = userBaseInfo.LastName
		userInfo.Role = USER_ROLE
		userInfo.UserName = userBaseInfo.UserName
		userInfo.Bio = userBaseInfo.Bio
		userInfo.Status = USER_STATUS_FREE
		userInfo.PhoneNumber = USER_PHONENUMBER_STATE_NOT_SET
		userInfo.Created_at = time.Now()
		CreateSingleUser(userInfo)
		return userInfo, true
	}
	return userInfo, false
}
func GetSingleUserInfo(userBaseInfo SingleUserModel) SingleUserModel {
	selectUserFilter := bson.M{"_id": userBaseInfo.ID}

	userInfo := GetSingleUser(selectUserFilter)
	return userInfo
}
func GetSingleUserRole(userBaseInfo SingleUserModel) string {
	selectUserFilter := bson.M{"_id": userBaseInfo.ID}

	userInfo := GetSingleUser(selectUserFilter)
	return userInfo.Role
}

func GetSingleUserUserName(userBaseInfo SingleUserModel) string {
	selectUserFilter := bson.M{"_id": userBaseInfo.ID}

	userInfo := GetSingleUser(selectUserFilter)
	return userInfo.UserName
}
func GetSingleUserPhoneNumber(userBaseInfo SingleUserModel) string {
	selectUserFilter := bson.M{"_id": userBaseInfo.ID}

	userInfo := GetSingleUser(selectUserFilter)
	return userInfo.PhoneNumber
}
func GetAllUsersInfo() []SingleUserModel {
	selectUserFilter := bson.M{"role": USER_ROLE}

	usersList := GetFilteredUsers(selectUserFilter)
	return usersList
}
func GetAllAdminsInfo() []SingleUserModel {
	selectUserFilter := bson.M{"role": ADMIN_ROLE}

	adminsList := GetFilteredUsers(selectUserFilter)
	return adminsList
}
func UpdateSingleUserInfo(userBaseInfo SingleUserModel) bool {
	selectUserFilter := bson.M{"_id": userBaseInfo.ID}
	hasErr := SetUpdateSingleUser(selectUserFilter, userBaseInfo)
	return hasErr
}
func CreateNewProjectBase(projectBaseInfo SingleProjectModel) (SingleProjectModel, bool) {
	projectBaseInfo.Status = PROJECT_STATUS_PENDING
	projectBaseInfo.ID = idGenarator()
	projectBaseInfo.ChannelPostId = -1
	projectBaseInfo.Created_at = time.Now()
	hasErr := CreateSingleProject(projectBaseInfo)
	return projectBaseInfo, hasErr
}
func GetSingleProjectInfo(projectBaseInfo SingleProjectModel) SingleProjectModel {
	selectProjectFilter := bson.M{
		"_id": projectBaseInfo.ID,
	}
	selectProject := GetSingleProject(selectProjectFilter)
	return selectProject
}
func UpdateSingleProjectInfo(projectBaseInfo SingleProjectModel) bool {
	selectProjectFilter := bson.M{
		"_id": projectBaseInfo.ID,
	}
	hasErr := SetUpdateSingleProject(selectProjectFilter, projectBaseInfo)
	return hasErr
}
