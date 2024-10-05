package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if existingUser, ok := getUser(user.Username); ok {
		if existingUser.Password == user.Password {
			existingUser.LoggedIn = true
			updateUser(existingUser)
			fmt.Fprintf(w, "User %s logged in successfully", user.Username)
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	} else {
		user.LoggedIn = true
		addUser(user)
		fmt.Fprintf(w, "User %s registered and logged in successfully", user.Username)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if existingUser, ok := getUser(user.Username); ok {
		if existingUser.LoggedIn {
			existingUser.LoggedIn = false
			updateUser(existingUser)
			fmt.Fprintf(w, "User %s logged out successfully", user.Username)
		} else {
			http.Error(w, "User is not logged in", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payment Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user, ok := getUser(payment.Username); ok {
		if user.LoggedIn {
			addPayment(payment)
			fmt.Fprintf(w, "Payment of %.2f received from user %s", payment.Amount, payment.Username)
		} else {
			http.Error(w, "User is not logged in", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}
