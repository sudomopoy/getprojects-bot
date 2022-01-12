package main

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func RecordedUserIfNotExists() {
	Connect()
}
func SetAdmin(id int) string {
	if DBSetAdmin(id) {
		return "ادمین اضافه شد."
	}
	return "ادمین از قبل وجود دارد."
}
func AddNewProject(title string, description string, userId int, username string) (string, string) {
	added, pjId := AddNewProjectBaseInfo(title, description, userId, username)
	if added {
		return label_project_entered, pjId
	} else {
		return "ثبت پروژه موفقیت آمیز نبود", "-1"
	}
}
func UpdateSingleProject(pjId string, NewArgs bson.D) (int, string, string, string) {
	var result bson.M = DBUpdateSingleProject(pjId, NewArgs)
	var userId int
	if reflect.TypeOf(result["userId"]).Kind() == reflect.Int32 {
		userId = int(result["userId"].(int32))
	} else {
		userId = int(result["userId"].(int64))
	}
	return userId, result["title"].(string), result["description"].(string), result["username"].(string)
}
func GetAdmins() []int {
	return GetAdminsIds()

}
func GetAllUsersInfo(pre string) []bson.M {
	users := GetFilterUser(bson.D{{"role", pre}})
	return users
}
func isPhoneNumberVerified(id int) bool {
	user := GetFilterUser(bson.D{{"_id", id}})
	return user[0]["phone"] == "NotSet"
}
func GetSingleProjectStatus(id string) (string, string) {
	project := GetSingleProject(id)
	return project["status"].(string), project["title"].(string)
}
func setPhoneNumber(id int, phone string) bool {
	return updateSingleUser(id, bson.D{{"phone", phone}})
}
