package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cloudenterprise/goblog/accountservice/dbclient"
	"github.com/gorilla/mux"
)

// DBClient is the db instance
var DBClient dbclient.IBoltClient

// GetAccount handles find account route
func GetAccount(w http.ResponseWriter, r *http.Request) {

	// read the 'accountId' path parameter from the mux map
	var accountID = mux.Vars(r)["accountID"]

	// Read the account struct from boltdb
	account, err := DBClient.QueryAccount(accountID)

	// if err, return 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// if found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
