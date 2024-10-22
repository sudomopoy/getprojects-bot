package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	//return
	if GetProccessMode() == "product" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: "https://63ddf52733e346ed9eed3f93267bffc6@sentry.hamravesh.com/228",
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err, " bot: ", token)
	}

	bot.Debug = GetProccessMode() == "development"

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 180

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.ChannelPost != nil {
			sentry.CaptureMessage(strconv.Itoa(int(update.ChannelPost.Chat.ID)))
		} else if update.Message != nil {
			var userBaseInfo SingleUserModel = SingleUserModel{
				ID:        strconv.FormatInt(update.Message.Chat.ID, 10),
				FirstName: update.Message.Chat.FirstName,
				LastName:  update.Message.Chat.LastName,
				Bio:       update.Message.Chat.Bio,
				UserName:  update.Message.Chat.UserName,
			}
			userId, err := strconv.Atoi(userBaseInfo.ID)
			check(err)
			msg := tgbotapi.NewMessage(int64(userId), "---")
			userBaseInfo, wasNewUser := SetUserBaseInfoIfNotExists(userBaseInfo)
			if !wasNewUser {
				updateUserData := SingleUserModel{
					ID:          strconv.FormatInt(update.Message.Chat.ID, 10),
					FirstName:   update.Message.Chat.FirstName,
					LastName:    update.Message.Chat.LastName,
					Bio:         update.Message.Chat.Bio,
					UserName:    update.Message.Chat.UserName,
					Role:        userBaseInfo.Role,
					PhoneNumber: userBaseInfo.PhoneNumber,
					Status:      userBaseInfo.Status,
					Created_at:  userBaseInfo.Created_at,
				}
				UpdateSingleUserInfo(updateUserData)
				userBaseInfo = GetSingleUserInfo(updateUserData)
			}
			if GetSingleUserRole(userBaseInfo) == ADMIN_ROLE { // !ADMIN AREA
				messageData := update.Message.Text
				msg.Text = label_home
				switch messageData {
				case "/start":
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				case label_users_list:
					allUsersList := GetAllUsersInfo()
					var finalUsers string = ""
					for i := 0; i < len(allUsersList); i++ {
						if allUsersList[i].PhoneNumber != USER_PHONENUMBER_STATE_NOT_SET {
							finalUsers += func() string {
								return fmt.Sprintf("کاربر : %v , شماره تلفن : %v \n", allUsersList[i].ID, allUsersList[i].PhoneNumber)
							}()
						}
					}
					if finalUsers == "" {
						finalUsers = label_user_not_found
					}
					msg.Text = finalUsers
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				default:
					msg.Text = description_command_not_found
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				}
			} else { // !USER
				cache, err := RedisClientGet(userId)
				log_excepts(func() string {
					return fmt.Sprintf(log_text,
						update.Message.Chat.UserName,
						userId,
						update.Message.Text,
						update.Message.Chat.FirstName,
						update.Message.Chat.LastName,
						update.Message.Chat.Bio,
						cache,
					)
				}())
				if update.Message.Text == "/start" {
					msg.Text = func() string {
						return fmt.Sprintf(description_welcome, update.Message.Chat.FirstName)
					}()
					RedisClientRemove(userId)
					msg.ReplyMarkup = mainPage_Keyboard
				} else if update.Message.Chat.UserName == "" {
					msg.Text = description_must_have_id
				} else if cache == "GETTING_PHONE_NUMBER:0" {
					var phoneNumber string = "---"
					if update.Message != nil {
						if update.Message.Contact != nil {
							phoneNumber = update.Message.Contact.PhoneNumber
							if userId == int(update.Message.Contact.UserID) && phoneNumber != "---" {
								if IranianPhoneValidate(phoneNumber) {
									userBaseInfo.PhoneNumber = phoneNumber
									if !UpdateSingleUserInfo(userBaseInfo) {
										msg.Text = "شماره تلفن تایید شد."
										RedisClientRemove(userId)
										msg.ReplyMarkup = mainPage_Keyboard
									} else {
										msg.Text = "مشکلی پیش آمده."
										userBaseInfo.PhoneNumber = "---"
									}
								} else {
									msg.Text = "شماره تلفن میبایست ایرانی باشد"
								}
							}
						}
					}
					if msg.Text == "---" {
						msg.Text = "لطفا شماره تلفن خود را به اشتراک بگذارید."
						userBaseInfo.PhoneNumber = "---"
					}
				} else if GetSingleUserPhoneNumber(userBaseInfo) == USER_PHONENUMBER_STATE_NOT_SET {
					msg.ReplyMarkup = send_phone_number_Keyboard
					msg.Text = description_should_have_phone_number
					RedisClientSet(userId, "GETTING_PHONE_NUMBER:0")
				} else if label_cancel_entring_project_proccess == update.Message.Text {
					RedisClientRemove(userId)
					msg.Text = description_project_entering_canceled
					msg.ReplyMarkup = mainPage_Keyboard
				} else if !err {
					data := strings.Split(cache, ":")
					case_key := data[0]
					metaData := data[1]
					switch case_key {
					case "ENTER_PROJECT_STEP+1":
						RedisClientSet(userId, "ENTER_PROJECT_STEP+2:"+update.Message.Text)
						msg.ReplyMarkup = budgets_Keyboard
						msg.Text = description_enter_budget
					case "ENTER_PROJECT_STEP+2":
						projectBudget := update.Message.Text
						projectDescription := metaData
						var projectBaseInfo SingleProjectModel = SingleProjectModel{
							UserId:      userBaseInfo.ID,
							Title:       titleDetector(projectDescription),
							Budget:      projectBudget,
							Description: projectDescription,
						}
						projectBaseInfo, hasErr := CreateNewProjectBase(projectBaseInfo)

						if hasErr {
							msg.Text = label_project_entered
						} else {
							msg.Text = label_has_problem
						}
						RedisClientRemove(userId)

						allAdmins := GetAllAdminsInfo()
						for i := 0; i < len(allAdmins); i++ {
							adminId, _ := strconv.Atoi(allAdmins[i].ID)
							messageTextForAdmin := func() string {
								return fmt.Sprintf(description_new_project_from_user, userId, metaData, update.Message.Text)
							}()
							messageForAdmin := tgbotapi.NewMessage(int64(adminId), messageTextForAdmin)
							var new_project_accept_or_denied = tgbotapi.NewInlineKeyboardMarkup(
								tgbotapi.NewInlineKeyboardRow(
									tgbotapi.NewInlineKeyboardButtonData(label_accept, "accept_project:"+projectBaseInfo.ID),
									tgbotapi.NewInlineKeyboardButtonData(label_reject, "denied_project:"+projectBaseInfo.ID),
								),
							)
							messageForAdmin.ReplyMarkup = new_project_accept_or_denied
							bot.Send(messageForAdmin)
							msg.ReplyMarkup = mainPage_Keyboard
							RedisClientRemove(userId)
						}
					default:
						RedisClientRemove(userId)
						msg.ReplyMarkup = mainPage_Keyboard
					}
				} else {
					msg.Text = label_home
					switch update.Message.Text {
					// case label_help:
					// 	msg.Text = description_help
					case label_support:
						msg.Text = description_support
					case label_bot_designer:
						msg.Text = description_bot_designer
					case "/start":
						RedisClientRemove(userId)
						msg.ReplyMarkup = mainPage_Keyboard
					case label_enter_project:
						msg.Text = description_bot_usage_roles
						msg.ReplyMarkup = i_accept_bot_usage_roles_Keyboard
					case label_accept_bot_usage_roles:
						RedisClientSet(userId, "ENTER_PROJECT_STEP+1:0")
						msg.ReplyMarkup = i_want_to_cancel_enter_project_Keyboard
						msg.Text = label_please_enter_poster_description
					case token + password:
						userBaseInfo.Role = ADMIN_ROLE
						hasErr := UpdateSingleUserInfo(userBaseInfo)
						if hasErr {
							msg.Text = label_has_problem

						} else {
							msg.Text = label_admin_added
						}
						RedisClientRemove(userId)
						msg.ReplyMarkup = mainPage_Keyboard
					default:
						msg.Text = description_command_not_found
						RedisClientRemove(userId)
						msg.ReplyMarkup = mainPage_Keyboard
					}
				}
			}
			msg.BaseChat.ReplyToMessageID = update.Message.MessageID

			if _, err = bot.Send(msg); err != nil {
				msg.Text = description_command_not_found
				msg.ReplyMarkup = mainPage_Keyboard
				RedisClientRemove(userId)
				bot.Send(msg)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
			}
			var userBaseInfo SingleUserModel = SingleUserModel{
				ID:        strconv.FormatInt(update.CallbackQuery.From.ID, 10),
				FirstName: update.CallbackQuery.From.FirstName,
				LastName:  update.CallbackQuery.From.LastName,
				UserName:  update.CallbackQuery.From.UserName,
			}
			userId, err := strconv.Atoi(userBaseInfo.ID)
			check(err)
			msg := tgbotapi.NewMessage(int64(userId), "---")
			userBaseInfo, _ = SetUserBaseInfoIfNotExists(userBaseInfo)

			//! ADMIN AREA
			if GetSingleUserRole(userBaseInfo) == ADMIN_ROLE {
				data := update.CallbackQuery.Data
				metaData := ""
				fmt.Println(data)

				if strings.Contains(data, ":") {
					tmp := strings.Split(data, ":")
					data = tmp[0]
					metaData = tmp[1]
				}
				msg.Text = label_home

				switch data {
				case "accept_project":
					projectId := metaData
					var projectBaseInfo SingleProjectModel = SingleProjectModel{
						ID: projectId,
					}
					projectBaseInfo = GetSingleProjectInfo(projectBaseInfo)
					fmt.Println(projectBaseInfo)
					if projectBaseInfo.Status == PROJECT_STATUS_PENDING {
						projectBaseInfo.Status = PROJECT_STATUS_ACCEPTED
						UpdateSingleProjectInfo(projectBaseInfo)
						msg.Text = label_project_accepted
						chanelMessageText := func() string {
							return fmt.Sprintf(description_project_poster, projectBaseInfo.Title, projectBaseInfo.Description, projectBaseInfo.Budget)
						}()
						userTempInfo := SingleUserModel{
							ID: projectBaseInfo.UserId,
						}
						var project_in_chanel = tgbotapi.NewInlineKeyboardMarkup(
							tgbotapi.NewInlineKeyboardRow(
								tgbotapi.NewInlineKeyboardButtonURL(label_message_to_owner, "t.me/"+GetSingleUserUserName(userTempInfo)),
							),
							tgbotapi.NewInlineKeyboardRow(
								tgbotapi.NewInlineKeyboardButtonURL(label_message_enter_same_project, "t.me/getprojectsbot"),
							),
						)
						chanelMessage := tgbotapi.NewMessage(masterChannelId, chanelMessageText)
						chanelMessage.ReplyMarkup = project_in_chanel
						userMessageText := func() string {
							return fmt.Sprintf(description_project_accepted, projectBaseInfo.Description)
						}()
						ownerUserId, _ := strconv.Atoi(projectBaseInfo.UserId)
						bot.Send(chanelMessage)
						bot.Send(tgbotapi.NewMessage(int64(ownerUserId), userMessageText))
					} else {
						msg.Text = func() string {
							return fmt.Sprintf(label_project_status_connot_edit, projectBaseInfo.Description)
						}()
					}
				case "denied_project":
					projectId := metaData
					check(err)
					var projectBaseInfo SingleProjectModel = SingleProjectModel{
						ID: projectId,
					}
					projectBaseInfo = GetSingleProjectInfo(projectBaseInfo)
					projectBaseInfo.Status = PROJECT_STATUS_REJECTED
					UpdateSingleProjectInfo(projectBaseInfo)
					msg.Text = label_project_rejected
					ownerUserId, _ := strconv.Atoi(projectBaseInfo.UserId)
					userMessageText := func() string {
						return fmt.Sprintf(description_project_rejected, projectBaseInfo.Description)
					}()
					bot.Send(tgbotapi.NewMessage(int64(ownerUserId), userMessageText))
				default:
					msg.Text = description_command_not_found
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
					RedisClientRemove(userId)
				}
			}
			if _, err := bot.Send(msg); err != nil {
				msg.Text = description_command_not_found
				msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				RedisClientRemove(userId)
				bot.Send(msg)
			}
		}
	}
}
