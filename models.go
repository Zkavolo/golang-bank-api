package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	LoggedIn bool   `json:"logged_in"`
}

type Payment struct {
	Username string  `json:"username"`
	Amount   float64 `json:"amount"`
}

type Data struct {
	Users    map[string]User `json:"users"`
	Payments []Payment       `json:"payments"`
}
