package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type LoanDetail struct {
	Lender       int
	Borrower     int
	LenderName   string
	BorrowerName string
	LoanAmt      int
	LoanEMI      int
}

var LoanDetails = []LoanDetail{}

func GetAllLoanDetailDB() {
	var filter, option interface{}
	// filter  gets all document
	filter = bson.D{}
	//  option remove id field from all documents
	option = bson.D{{"_id", 0}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// This method returns momngo.cursor and error if any.
	cursor, err := QueryDB(ctx, "MicroFinance", "loandetails", filter, option)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err := cursor.All(Ctx, &results); err != nil {
		panic(err)
	}
	// Append the loan details
	formatResult(results)
}

func formatResult(result []bson.D) {
	loanInfo := LoanDetail{}
	LoanDetails = []LoanDetail{}
	for _, doc := range result {
		bsonBytes, _ := bson.Marshal(doc)
		bson.Unmarshal(bsonBytes, &loanInfo)
		LoanDetails = append(LoanDetails, loanInfo)
	}
}
