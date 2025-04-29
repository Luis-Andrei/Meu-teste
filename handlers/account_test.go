package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"bank-server/db"
	"bank-server/models"

	"github.com/gorilla/mux"
)

var createdAccountID int

func setupDB(t *testing.T) {
	err := db.Connect()
	if err != nil {
		t.Fatalf("Erro ao conectar no banco: %v", err)
	}
}

func cleanupDB(t *testing.T) {
	if createdAccountID != 0 {
		_, err := db.DB.Exec("DELETE FROM accounts WHERE id = $1", createdAccountID)
		if err != nil {
			t.Logf("Erro ao limpar banco: %v", err)
		}
	}
}

func TestFullAccountFlow(t *testing.T) {
	setupDB(t)
	defer cleanupDB(t)

	t.Run("Create Account", func(t *testing.T) {
		body := []byte(`{"name": "Cliente Teste", "balance": 500.0}`)
		req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		CreateAccount(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("esperado 201, obtido %d", w.Code)
		}

		var acc models.Account
		if err := json.NewDecoder(w.Body).Decode(&acc); err != nil {
			t.Fatalf("erro ao decodificar resposta: %v", err)
		}

		if acc.ID == 0 || acc.Balance != 500.0 {
			t.Errorf("criação falhou: %+v", acc)
		}

		createdAccountID = acc.ID
	})

	t.Run("Get Account", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/accounts/"+strconv.Itoa(createdAccountID), nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(createdAccountID)})
		w := httptest.NewRecorder()
		GetAccount(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("esperado 200, obtido %d", w.Code)
		}

		var acc models.Account
		if err := json.NewDecoder(w.Body).Decode(&acc); err != nil {
			t.Fatalf("erro ao decodificar resposta: %v", err)
		}

		if acc.ID != createdAccountID {
			t.Errorf("ID esperado %d, obtido %d", createdAccountID, acc.ID)
		}
	})

	t.Run("Deposit Funds", func(t *testing.T) {
		body := []byte(`{"amount": 200.0}`)
		req := httptest.NewRequest("POST", "/accounts/"+strconv.Itoa(createdAccountID)+"/deposit", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(createdAccountID)})
		w := httptest.NewRecorder()
		Deposit(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("esperado 200, obtido %d", w.Code)
		}

		var acc models.Account
		if err := json.NewDecoder(w.Body).Decode(&acc); err != nil {
			t.Fatalf("erro ao decodificar resposta: %v", err)
		}

		expected := 700.0
		if acc.Balance != expected {
			t.Errorf("saldo esperado %.2f, obtido %.2f", expected, acc.Balance)
		}
	})

	t.Run("Withdraw Funds", func(t *testing.T) {
		body := []byte(`{"amount": 100.0}`)
		req := httptest.NewRequest("POST", "/accounts/"+strconv.Itoa(createdAccountID)+"/withdraw", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(createdAccountID)})
		w := httptest.NewRecorder()
		Withdraw(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("esperado 200, obtido %d", w.Code)
		}

		var acc models.Account
		if err := json.NewDecoder(w.Body).Decode(&acc); err != nil {
			t.Fatalf("erro ao decodificar resposta: %v", err)
		}

		expected := 600.0
		if acc.Balance != expected {
			t.Errorf("saldo esperado %.2f, obtido %.2f", expected, acc.Balance)
		}
	})

	t.Run("Get Account After Withdraw", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/accounts/"+strconv.Itoa(createdAccountID), nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(createdAccountID)})
		w := httptest.NewRecorder()
		GetAccount(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("esperado 200, obtido %d", w.Code)
		}

		var acc models.Account
		if err := json.NewDecoder(w.Body).Decode(&acc); err != nil {
			t.Fatalf("erro ao decodificar resposta: %v", err)
		}

		if acc.Balance != 600.0 {
			t.Errorf("saldo esperado 600.0, obtido %.2f", acc.Balance)
		}
	})
}

func TestGetAccountInvalidID(t *testing.T) {
	setupDB(t)
	defer cleanupDB(t)

	req := httptest.NewRequest("GET", "/accounts/9999", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "9999"})
	w := httptest.NewRecorder()
	GetAccount(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("esperado 404, obtido %d", w.Code)
	}
}

func TestDepositInvalidAmount(t *testing.T) {
	setupDB(t)
	defer cleanupDB(t)

	body := []byte(`{"amount": -200.0}`)
	req := httptest.NewRequest("POST", "/accounts/1/deposit", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	Deposit(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperado 400, obtido %d", w.Code)
	}
}

func TestWithdrawInvalidAmount(t *testing.T) {
	setupDB(t)
	defer cleanupDB(t)

	body := []byte(`{"amount": -100.0}`)
	req := httptest.NewRequest("POST", "/accounts/1/withdraw", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	Withdraw(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperado 400, obtido %d", w.Code)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	setupDB(t)
	defer cleanupDB(t)

	// Primeiro cria uma conta com saldo 0
	body := []byte(`{"name": "Cliente Teste", "balance": 0.0}`)
	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	CreateAccount(w, req)

	var acc models.Account
	if err := json.NewDecoder(w.Body).Decode(&acc); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}

	// Tenta sacar mais do que tem
	body = []byte(`{"amount": 1000.0}`)
	req = httptest.NewRequest("POST", "/accounts/"+strconv.Itoa(acc.ID)+"/withdraw", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(acc.ID)})
	w = httptest.NewRecorder()
	Withdraw(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperado 400, obtido %d", w.Code)
	}
}

func intToStr(n int) string {
	return string(n)
}
