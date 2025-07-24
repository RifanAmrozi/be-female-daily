package data

import "be-female-daily/model"

var Users = []model.User{
	{ID: 1, Username: "user1", Email: "user1@example.com", Password: "password1", Point: 100},
	{ID: 2, Username: "user2", Email: "user2@example.com", Password: "password2", Point: 200},
	{ID: 3, Username: "user3", Email: "user3@example.com", Password: "password3", Point: 300},
	// Add more users as needed
}
var Tickets = []model.Ticket{}
