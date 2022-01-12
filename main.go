package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
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
			fmt.Println(update.ChannelPost.Chat.ID)
		} else if update.Message != nil {
			userId := int(update.Message.Chat.ID)
			msg := tgbotapi.NewMessage(int64(userId), "")
			SetUserBaseInfoIfNotExists(userId)
			if IsAdmin(int(userId)) { // !ADMIN
				data := update.Message.Text
				msg.Text = label_home
				switch data {
				case "/start":
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				case label_users_list:
					users := GetAllUsersInfo("user")
					var finalUsers string = ""
					for i := 0; i < len(users); i++ {
						if users[i]["phone"] != "NotSet" {
							finalUsers += func() string {
								return fmt.Sprintf("کاربر : %v , شماره تلفن : %v \n", users[i]["_id"], users[i]["phone"])
							}()
						}
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
									if setPhoneNumber(userId, phoneNumber) {
										msg.Text = "شماره تلفن تایید شد."
										RedisClientRemove(userId)
										msg.ReplyMarkup = mainPage_Keyboard
									} else {
										msg.Text = "مشکلی پیش آمده."
									}
								} else {
									msg.Text = "شماره تلفن میبایست ایرانی باشد"
								}
							}
						}
					}
					if msg.Text == "" {
						msg.Text = "لطفا شماره تلفن خود را به اشتراک بگذارید."
					}
				} else if isPhoneNumberVerified(userId) {
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
						msg.ReplyMarkup = i_want_to_cancel_enter_project_Keyboard
						msg.Text = label_please_enter_poster_description
					case "ENTER_PROJECT_STEP+2":
						res, pjId := AddNewProject(metaData, update.Message.Text, userId, update.Message.Chat.UserName)
						msg.Text = res
						RedisClientRemove(userId)

						admins := GetAdmins()
						for i := 0; i < len(admins); i++ {
							adminId := admins[i]
							messageTextForAdmin := func() string {
								return fmt.Sprintf(description_new_project_from_user, userId, metaData, update.Message.Text)
							}()
							messageForAdmin := tgbotapi.NewMessage(int64(adminId), messageTextForAdmin)
							var new_project_accept_or_denied = tgbotapi.NewInlineKeyboardMarkup(
								tgbotapi.NewInlineKeyboardRow(
									tgbotapi.NewInlineKeyboardButtonData(label_accept, "accept_project:"+pjId),
									tgbotapi.NewInlineKeyboardButtonData(label_reject, "denied_project:"+pjId),
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
						msg.Text = label_please_enter_poster_title
					case token + password:
						msg.Text = SetAdmin(int(userId))
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
			userId := int(update.CallbackQuery.Message.Chat.ID)
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			//! ADMIN
			if IsAdmin(int(userId)) {
				data := update.CallbackQuery.Data
				metaData := ""
				if strings.Contains(data, ":") {
					tmp := strings.Split(data, ":")
					data = tmp[0]
					metaData = tmp[1]
				}
				msg.Text = label_home
				switch data {
				case "accept_project":
					pjId := metaData
					check(err)
					userId, title, description, username := UpdateSingleProject(pjId, bson.D{{"status", "accept"}})
					msg.Text = label_project_accepted
					chanelMessageText := func() string {
						return fmt.Sprintf(description_project_poster, title, description)
					}()
					var project_in_chanel = tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonURL(label_message_to_owner, "t.me/"+username),
						),
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonURL(label_message_enter_same_project, "t.me/getprojectsbot"),
						),
					)
					chanelMessage := tgbotapi.NewMessage(masterChannelId, chanelMessageText)
					chanelMessage.ReplyMarkup = project_in_chanel
					userMessageText := func() string {
						return fmt.Sprintf(description_project_accepted, title)
					}()
					bot.Send(chanelMessage)
					bot.Send(tgbotapi.NewMessage(int64(userId), userMessageText))
				case "denied_project":
					pjId := metaData
					check(err)
					userId, title, _, _ := UpdateSingleProject(pjId, bson.D{{"status", "reject"}})
					msg.Text = label_project_rejected
					userMessageText := func() string {
						return fmt.Sprintf(description_project_rejected, title)
					}()
					bot.Send(tgbotapi.NewMessage(int64(userId), userMessageText))
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
