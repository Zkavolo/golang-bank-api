package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var (
	data  Data
	mutex = &sync.Mutex{}
)

func initStorage() {
	data = Data{
		Users:    make(map[string]User),
		Payments: []Payment{},
	}
	loadData()
}

func loadData() {
	fileData, err := ioutil.ReadFile("data.json")
	if err == nil {
		json.Unmarshal(fileData, &data)
	}
}

func saveData() error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("data.json", jsonData, 0644)
}

func getUser(username string) (User, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	user, ok := data.Users[username]
	return user, ok
}

func addUser(user User) {
	mutex.Lock()
	defer mutex.Unlock()
	data.Users[user.Username] = user
	saveData()
}

func updateUser(user User) {
	mutex.Lock()
	defer mutex.Unlock()
	data.Users[user.Username] = user
	saveData()
}

func addPayment(payment Payment) {
	mutex.Lock()
	defer mutex.Unlock()
	data.Payments = append(data.Payments, payment)
	saveData()
}
