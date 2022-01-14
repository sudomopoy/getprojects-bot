package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// User Keyboards
var mainPage_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_enter_project),
	),
	// tgbotapi.NewKeyboardButtonRow(
	// tgbotapi.NewKeyboardButton(label_help),
	// ),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_support),
		tgbotapi.NewKeyboardButton(label_bot_designer),
	),
)
var i_want_to_cancel_enter_project_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_cancel_entring_project_proccess),
	),
)
var budgets_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("10 تا 100 هزار تومان"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("100 تا 500 هزار تومان"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("500 هزار تومان تا یک و نیم میلیون تومان"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("یک و نیم تا 5 میلیون تومان"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("توافقی"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_cancel_entring_project_proccess),
	),
)
var i_accept_bot_usage_roles_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_accept_bot_usage_roles),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_cancel_entring_project_proccess),
	),
)

var send_phone_number_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonContact("📱 شماره موبایلم را ارسال کن"),
	),
)

// Admin Keyboards

var ADMIN_mainPage_Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(label_users_list),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("لاگ"),
		tgbotapi.NewKeyboardButton("پیام ها"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("لیست پروژه ها"),
	),
)
