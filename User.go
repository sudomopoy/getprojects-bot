package main

import "go.mongodb.org/mongo-driver/bson"




func RecordedUserIfNotExists() {
	Connect()
}
func SetAdmin(id int) string {
	if DBSetAdmin(id) {
		return "ادمین اضافه شد."
	}
	return "ادمین از قبل وجود دارد."
}
func AddNewProject(title string, description string, userId int, username string) (string, int) {
	added, pjId := AddNewProjectBaseInfo(title, description, userId, username)
	if added {
		return "پروژه ثبت گردید. پس از تایید انتشار میابد.", pjId
	} else {
		return "ثبت پروژه موفقیت آمیز نبود", -1
	}
}
func UpdateSingleProject(pjId int, NewArgs bson.D) (int, string, string, string) {
	var result bson.M = DBUpdateSingleProject(pjId, NewArgs)
	return int(result["userId"].(int64)), result["title"].(string), result["description"].(string), result["username"].(string)
}
func GetAdmins() []int {
	return GetAdminsIds()

}
