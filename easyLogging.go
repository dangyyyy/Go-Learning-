package main

import (
	"fmt"
	"math/rand"
	"os"
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
	users := GenerateUsers(1000)
	for _, user := range users {
		SaveUserInfo(user)

	}
}

func GenerateUsers(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	return users
}
func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := 0; i < count; i++ {
		logs[i] = logItem{actions[rand.Intn(len(actions)-1)], time.Now()}
	}
	return logs
}
func SaveUserInfo(user User) error {
	fmt.Printf("Saving user info: %d\n", user.id)
	filename := fmt.Sprintf("logs/user%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.GetAction())
	return err
}
