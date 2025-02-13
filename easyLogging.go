package main

import (
	"fmt"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"update record",
	"delete record",
}

type User struct {
	id    int
	email string
	logs  []logItem
}
type logItem struct {
	action string
	time   time.Time
}

func (u User) GetAction() string {
	out := fmt.Sprintf("ID: %d | Email: %s\n Activity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.time)

	}
	return out
}
func main() {
	u := User{
		id:    1,
		email: "john@gmail.com",
		logs: []logItem{
			{actions[0], time.Now()},
			{actions[1], time.Now()},
			{actions[2], time.Now()},
			{actions[3], time.Now()},
			{actions[4], time.Now()},
		},
	}
	fmt.Println(u.GetAction())
}
