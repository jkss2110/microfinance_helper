package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microfinance/models"
	"net/http"
	"strconv"
	"strings"
)

func getAllLoans(w http.ResponseWriter, r *http.Request) {
	loans := &models.LoanDetails
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
	nextID := getNextID()
	newLoan.Lender = nextID
	nextID++
	newLoan.Borrower = nextID + 1
	models.LoanDetails = append(models.LoanDetails, newLoan)
	w.WriteHeader(http.StatusCreated)
}

func getNextID() int {
	highestID := -1
	for _, loan := range models.LoanDetails {
		if highestID <= loan.Borrower {
			highestID = loan.Borrower
		}
	}
	return highestID + 1
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

func findLoanDetail(ID int) (models.LoanDetail, error) {
	for _, loan := range models.LoanDetails {
		if loan.Borrower == ID {
			return loan, nil
		}
	}
	return models.LoanDetail{}, fmt.Errorf("not found")
}

func getSingleLoan(borrowerID int, w http.ResponseWriter, r *http.Request) {
	loan, err := findLoanDetail(borrowerID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	loanJSON, err := json.Marshal(loan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(loanJSON)
	w.WriteHeader(http.StatusFound)
}
