package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "test"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)
var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 180

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		if update.ChannelPost != nil {
			fmt.Println(update.ChannelPost.Chat.ID)
		}

		// Check if we've gotten a message update.
		if update.Message != nil {
			userId := int(update.Message.Chat.ID)

			// Construct a new message from the given chat ID and containing
			// the text that we received.
			fmt.Print("%s", update)
			msg := tgbotapi.NewMessage(int64(userId), "")
			SetUserBaseInfoIfNotExists(userId)
			if IsAdmin(int(userId)) {
				data := update.Message.Text
				msg.Text = "خانه"
				switch data {
				case "/start":
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				default:
					msg.Text = "دستور پیدا نشد! \n لطفا یکی از گزینه های زیر را انتخاب کنید."
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				}
			} else {
				cache, err := RedisClientGet(userId)
				if !err {
					data := strings.Split(cache, ":")
					case_key := data[0]
					metaData := data[1]
					switch case_key {
					case "ENTER_PROJECT_STEP+1":
						RedisClientSet(userId, "ENTER_PROJECT_STEP+2:"+update.Message.Text)
						msg.ReplyMarkup = i_want_to_cancel_enter_project_Keyboard
						msg.Text = "لطفا متن آگهی خود را وارد کنید."
					case "ENTER_PROJECT_STEP+2":
						res, pjId := AddNewProject(metaData, update.Message.Text, userId, update.Message.Chat.UserName)
						msg.Text = res
						msg.ReplyMarkup = mainPage_Keyboard
						admins := GetAdmins()
						for i := 0; i < len(admins); i++ {
							adminId := admins[i]
							messageTextForAdmin := "کاربر" + strconv.Itoa(userId) + " پروژه ای با عنوان << " + metaData + " >> \n و متن : \n" + update.Message.Text + "\n ثبت کرده است"
							messageForAdmin := tgbotapi.NewMessage(int64(adminId), messageTextForAdmin)
							var new_project_accept_or_denied = tgbotapi.NewInlineKeyboardMarkup(
								tgbotapi.NewInlineKeyboardRow(
									tgbotapi.NewInlineKeyboardButtonData("تایید پروژه", "accept_project:"+strconv.Itoa(pjId)),
									tgbotapi.NewInlineKeyboardButtonData("رد پروژه", "denied_project:"+strconv.Itoa(pjId)),
								),
							)
							messageForAdmin.ReplyMarkup = new_project_accept_or_denied
							bot.Send(messageForAdmin)
						}
					default:
						//TODO delete redis cache
						msg.ReplyMarkup = mainPage_Keyboard
					}
				} else {
					msg.Text = "صفحه اصلی 🏠"
					switch update.Message.Text {
					case "/start":
						msg.ReplyMarkup = mainPage_Keyboard
					case label_enter_project:
						RedisClientSet(userId, "ENTER_PROJECT_STEP+1:0")
						msg.ReplyMarkup = i_want_to_cancel_enter_project_Keyboard
						msg.Text = "لطفا عنوان آگهی خود را وارد کنید."
					case token + password:
						msg.Text = SetAdmin(int(userId))
						msg.ReplyMarkup = mainPage_Keyboard
					default:
						msg.Text = "دستور پیدا نشد! \n لطفا یکی از گزینه های زیر را انتخاب کنید."
						msg.ReplyMarkup = mainPage_Keyboard
					}
				}
			}
			msg.BaseChat.ReplyToMessageID = update.Message.MessageID

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
			}
			userId := int(update.CallbackQuery.Message.Chat.ID)
			fmt.Println(userId)

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			if IsAdmin(int(userId)) {
				data := update.CallbackQuery.Data
				metaData := ""
				if strings.Contains(data, ":") {
					tmp := strings.Split(data, ":")
					data = tmp[0]
					metaData = tmp[1]
				}
				msg.Text = "خانه"
				switch data {
				case "accept_project":
					pjId, err := strconv.Atoi(metaData)
					if err != nil {
						log.Fatalf("err")
					}
					userId, title, description, username := UpdateSingleProject(pjId, bson.D{{"status", "accept"}})
					msg.Text = "پروژه تایید شد✅"
					chanelMessageText := func() string {
						return fmt.Sprintf(`
• %v
%v

➖➖➖➖➖➖➖
برای ثبت آگهی به ربات @getprojectsbot مراجعه کنید.
t.me/getprojectsofficial	
									`, title, description)
					}()
					var project_in_chanel = tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonURL("پیام به کارفرما", "https://t.me/account/"+username),
						),
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonURL("ثبت پروژه مشابه", "https://t.me/account/getprojectsbot"),
						),
					)
					chanelMessage := tgbotapi.NewMessage(int64(-1001763684409), chanelMessageText)
					chanelMessage.ReplyMarkup = project_in_chanel
					userMessageText := "پروژه ی: " + title + "\n تایید شد و در کانال قرار گرفت.✅"
					bot.Send(chanelMessage)
					bot.Send(tgbotapi.NewMessage(int64(userId), userMessageText))
				case "denied_project":
					pjId, err := strconv.Atoi(metaData)
					if err != nil {
						log.Fatalf("err")
					}
					userId, title, _, _ := UpdateSingleProject(pjId, bson.D{{"status", "reject"}})
					msg.Text = "پروژه رد شد🚫"
					userMessageText := "پروژه ی: " + title + "\n رد شد .🚫"
					bot.Send(tgbotapi.NewMessage(int64(userId), userMessageText))
				default:
					msg.Text = "دستور پیدا نشد! \n لطفا یکی از گزینه های زیر را انتخاب کنید."
					msg.ReplyMarkup = ADMIN_mainPage_Keyboard
				}
			}
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
		// else if update.CallbackQuery != nil {
		// 	// Respond to the callback query, telling Telegram to show the user
		// 	// a message with the data received.
		// 	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		// 	if _, err := bot.Request(callback); err != nil {
		// 		panic(err)
		// 	}

		// 	// And finally, send a message containing the data received.
		// 	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
		// 	if _, err := bot.Send(msg); err != nil {
		// 		panic(err)
		// 	}
		// }
	}
}
