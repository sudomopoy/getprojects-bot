package main

import (
	"fmt"

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
	fmt.Println(result)
	return int(result["userId"].(int64)), result["title"].(string), result["description"].(string), result["username"].(string)
}
func GetAdmins() []int {
	return GetAdminsIds()

}
func GetUsersInfo() {
	users := GetFilterUser(bson.D{{"role", "user"}})
	for i := 0; i < len(users); i++ {
		//users[i]["_id"]
	}
	return
}
