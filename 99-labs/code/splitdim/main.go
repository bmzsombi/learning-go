package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"splitdim/pkg/api"
	"splitdim/pkg/db/kvstore"
	"splitdim/pkg/db/local"
)

// KVStoreMode defines the data layer mode (local/redis/kvstore).
var KVStoreMode = "local"

// KVStoreAddr stores the key-value store address as a DNS domain name or IP address.
var KVStoreAddr = "localhost:8001"

// TransferHandler is a HTTP handler that implements the money transfer API.
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	var body api.Transfer
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json, err := json.Marshal(body)
	if err != nil {
		w.Write(json)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Transfer called")
	err = db.Transfer(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err2 := errors.New(err.Error())
		fmt.Fprintf(w, "Transfer request failed: %s", err2)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// AccountListHandler is a HTTP handler that returns the current balance of each registered user.
func AccountListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Printf("AccountList called")
	var accountList []api.Account
	accountList, err := db.AccountList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err2 := errors.New(err.Error())
		fmt.Fprintf(w, "AccountList called failed: %s", err2)
		return
	}
	json, err := json.Marshal(accountList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(json)
	w.WriteHeader(http.StatusOK)
}

// ClearHandler is a HTTP handler that returns a list of transfers to clear the balance of each user.
func ClearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Printf("Clear has been called")
	transfers, err := db.Clear()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(transfers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(json)
	w.WriteHeader(http.StatusOK)
}

// ResetHandler is a HTTP handler that allows to zero out all balances.
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Printf("Reset has been called")
	err := db.Reset()
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err2 := errors.New(err.Error())
		fmt.Fprintf(w, "API request failed: %s", err2)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

var db api.DataLayer

func main() {
	// Set the default logger to a fancier log format.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if os.Getenv("KVSTORE_MODE") != "" {
		KVStoreMode = os.Getenv("KVSTORE_MODE")
	}
	if os.Getenv("KVSTORE_ADDR") != "" {
		KVStoreAddr = os.Getenv("KVSTORE_ADDR")
	}

	switch KVStoreMode {
	case "kvstore":
		log.Printf("Using the kvstore datalayer using at %q", KVStoreAddr)
		db = kvstore.NewDataLayer(KVStoreAddr)
	case "local":
		fallthrough
	default:
		log.Println("Using the local datalayer")
		db = local.NewDataLayer()
	}

	//db = local.NewDataLayer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/api/transfer", TransferHandler)
	http.HandleFunc("/api/accounts", AccountListHandler)
	http.HandleFunc("/api/clear", ClearHandler)
	http.HandleFunc("/api/reset", ResetHandler)

	log.Println("Server listening on http://:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
