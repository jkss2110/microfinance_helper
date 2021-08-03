package controller

import (
	"encoding/json"
	"io/ioutil"
	"microfinance/models"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func getAllLoans(w http.ResponseWriter, r *http.Request) {

	filter := bson.D{}
	results, err := models.GetLoanDetailDB(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	loans := formatResult(results)
	models.LoanDetails = loans
	loanJSON, err := json.Marshal(loans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(loanJSON)
	w.WriteHeader(http.StatusAccepted)
}

func createLoanDetails(w http.ResponseWriter, r *http.Request) {
	var newLoan models.LoanDetail
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &newLoan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, erro := models.InsertOneLoanDetailDB(newLoan)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusCreated)
}

func HandleHttp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllLoans(w, r)
	case http.MethodPost:
		createLoanDetails(w, r)
	default:
		panic("Problem")
	}
}

func HandleSingleHttp(w http.ResponseWriter, r *http.Request) {
	sUrlParameters := strings.Split(r.URL.Path, "loandetails/")
	borrowerID, err := strconv.Atoi(sUrlParameters[len(sUrlParameters)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getSingleLoan(borrowerID, w, r)
	default:
		panic("Wrong request")
	}
}

func getSingleLoan(borrowerID int, w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{"Borrower", borrowerID}}
	results, err := models.GetLoanDetailDB(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	loanInfo := formatResult(results)

	w.Header().Set("Content-Type", "application/json")
	loanJSON, err := json.Marshal(loanInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(loanJSON)
	w.WriteHeader(http.StatusFound)
}

func formatResult(result []bson.D) []models.LoanDetail {
	loanInfo := models.LoanDetail{}
	loanDetails := []models.LoanDetail{}
	for _, doc := range result {
		bsonBytes, _ := bson.Marshal(doc)
		bson.Unmarshal(bsonBytes, &loanInfo)
		loanDetails = append(loanDetails, loanInfo)
	}
	return loanDetails
}
