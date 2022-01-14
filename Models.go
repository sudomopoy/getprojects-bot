package main

import (
	"time"
)

type SingleProjectModel struct {
	ID            string    `bson:"_id"`
	UserId        string    `bson:"userId"`
	Title         string    `bson:"title"`
	Budget        string    `bson:"budget"`
	Description   string    `bson:"description"`
	Status        string    `bson:"status"`
	ChannelPostId int       `bson:"channelPostId"`
	Created_at    time.Time `bson:"created_at"`
}
type SingleUserModel struct {
	ID          string    `bson:"_id"`
	FirstName   string    `bson:"firstName"`
	LastName    string    `bson:"lastName"`
	Bio         string    `bson:"bio"`
	PhoneNumber string    `bson:"phoneNumber"`
	Role        string    `bson:"role"`
	Status      string    `bson:"status"`
	UserName    string    `bson:"userName"`
	Created_at  time.Time `bson:"created_at"`
}
