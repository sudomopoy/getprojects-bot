package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// User Keyboards
var mainPage_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_enter_project),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💡 راهنما "),
		tgbotapi.NewKeyboardButton("☎️ پشتیبانی"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💻 طراح ربات"),
	),
)
var i_want_to_cancel_enter_project_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("❌ از ثبت آگهی منصرف شدم"),
	),
)

// Admin Keyboards

var ADMIN_mainPage_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("لیست کاربران"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("لاگ"),
		tgbotapi.NewKeyboardButton("پیام ها"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("لیست پروژه ها"),
	),
)

