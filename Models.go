package main

import (
	"time"
)

type SingleProjectModel struct {
	_id           string    `bson:"_id"`
	userId        string    `bson:"userId"`
	title         string    `bson:"title"`
	description   string    `bson:"description"`
	status        string    `bson:"status"`
	channelPostId int       `bson:"channelPostId"`
	created_at    time.Time `bson:"created_at"`
}
type SingleUserModel struct {
	_id         string    `bson:"_id"`
	firstName   string    `bson:"firstName"`
	lastName    string    `bson:"lastName"`
	bio         string    `bson:"bio"`
	phoneNumber string    `bson:"phoneNumber"`
	role        string    `bson:"role"`
	status      string    `bson:"status"`
	userName    string    `bson:"userName"`
	created_at  time.Time `bson:"created_at"`
}
