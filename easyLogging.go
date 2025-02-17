package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
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
	t := time.Now()
	wg := &sync.WaitGroup{}
	users := make(chan User)
	go GenerateUsers(1000, users)

	for user := range users {
		wg.Add(1)
		go SaveUserInfo(user, wg)

	}
	wg.Wait()
	fmt.Println("Time elapsed: ", time.Since(t).String())
}

func GenerateUsers(count int, users chan User) {

	for i := 0; i < count; i++ {
		users <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	close(users)
}
func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := 0; i < count; i++ {
		logs[i] = logItem{actions[rand.Intn(len(actions)-1)], time.Now()}
	}
	return logs
}
func SaveUserInfo(user User, wg *sync.WaitGroup) error {
	fmt.Printf("Saving user info: %d\n", user.id)
	filename := fmt.Sprintf("logs/user%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.GetAction())
	if err != nil {
		return err
	}
	wg.Done()
	return nil
}
