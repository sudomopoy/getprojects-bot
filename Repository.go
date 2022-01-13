package main

import "time"

func SetUserBaseInfoIfNotExists(userBaseInfo SingleUserModel) SingleUserModel {
	selectUserFilter := SingleUserModel{
		_id: userBaseInfo._id,
	}
	userInfo, hasErr := GetSingleUser(selectUserFilter)
	if hasErr {
		userInfo.firstName = userBaseInfo.firstName
		userInfo.lastName = userBaseInfo.lastName
		userInfo.bio = userBaseInfo.bio
		userInfo.role = USER_ROLE
		userInfo.status = USER_STATUS_BLOCKED
		userInfo.phoneNumber = USER_PHONENUMBER_STATE_NOT_SET
		userInfo.created_at = time.Now()
		hasErr = CreateSingleUser(userInfo)

		if !hasErr {
			_, hasErr = GetSingleUser(userInfo)
		}
	}
	return userInfo
}

func GetSingleUserRole(userBaseInfo SingleUserModel) string {
	selectUserFilter := SingleUserModel{
		_id: userBaseInfo._id,
	}
	userInfo, _ := GetSingleUser(selectUserFilter)
	return userInfo.role
}
func GetSingleUserUserName(userBaseInfo SingleUserModel) string {
	selectUserFilter := SingleUserModel{
		_id: userBaseInfo._id,
	}
	userInfo, _ := GetSingleUser(selectUserFilter)
	return userInfo.userName
}
func GetSingleUserPhoneNumber(userBaseInfo SingleUserModel) string {
	selectUserFilter := SingleUserModel{
		_id: userBaseInfo._id,
	}
	userInfo, _ := GetSingleUser(selectUserFilter)
	return userInfo.phoneNumber
}
func GetAllUsersInfo() []SingleUserModel {
	selectUserFilter := SingleUserModel{
		role: USER_ROLE,
	}
	usersList := GetFilteredUsers(selectUserFilter)
	return usersList
}
func GetAllAdminsInfo() []SingleUserModel {
	selectUserFilter := SingleUserModel{
		role: ADMIN_ROLE,
	}
	adminsList := GetFilteredUsers(selectUserFilter)
	return adminsList
}
func UpdateSingleUserInfo(userBaseInfo SingleUserModel) bool {
	selectUserFilter := SingleUserModel{
		_id: userBaseInfo._id,
	}
	hasErr := SetUpdateSingleUser(selectUserFilter, userBaseInfo)
	return hasErr
}
func CreateNewProjectBase(projectBaseInfo SingleProjectModel) (SingleProjectModel, bool) {
	projectBaseInfo.status = PROJECT_STATUS_PENDING
	projectBaseInfo._id = idGenarator()
	projectBaseInfo.channelPostId = -1
	projectBaseInfo.created_at = time.Now()
	hasErr := CreateSingleProject(projectBaseInfo)
	return projectBaseInfo, hasErr
}
func GetSingleProjectInfo(projectBaseInfo SingleProjectModel) SingleProjectModel {
	selectProjectFilter := SingleProjectModel{
		_id: projectBaseInfo._id,
	}
	selectProjectFilter = GetSingleProject(selectProjectFilter)
	return selectProjectFilter
}
func UpdateSingleProjectInfo(projectBaseInfo SingleProjectModel) bool {
	selectprojectFilter := SingleProjectModel{
		_id: projectBaseInfo._id,
	}
	hasErr := SetUpdateSingleProject(selectprojectFilter, projectBaseInfo)
	return hasErr
}
