package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bank-server/db"
	"bank-server/models"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newAcc, err := db.CreateAccount(acc.Name, acc.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAcc)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	acc, err := db.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var deposit struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acc, err := db.UpdateBalance(id, deposit.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var withdraw struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&withdraw); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acc, err := db.UpdateBalance(id, -withdraw.Amount)
	if err != nil {
		if err.Error() == "saldo insuficiente" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}
