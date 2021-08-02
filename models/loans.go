package models

import (
	"context"
	"microfinance/shared"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type LoanDetail struct {
	Lender       int    `bson:"Lender"`
	Borrower     int    `bson:"Borrower"`
	LenderName   string `bson:"LenderName"`
	BorrowerName string `bson:"BorrowerName"`
	LoanAmt      int    `bson:"LoanAmt"`
	LoanEMI      int    `bson:"LoanEMI"`
}

var LoanDetails = []LoanDetail{}

func GetAllLoanDetailDB() (results []bson.D, err error) {
	var filter, option interface{}
	// filter  gets all document
	filter = bson.D{}
	//  option remove id field from all documents
	option = bson.D{{"_id", 0}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// This method returns momngo.cursor and error if any.
	cursor, err := QueryDB(ctx, shared.Database, shared.LoanDetailCollection, filter, option)
	if err != nil {
		panic(err)
	}
	if err := cursor.All(Ctx, &results); err != nil {
		panic(err)
	}
	return results, err
}

func InsertOneLoanDetailDB(loanInfo LoanDetail) (result interface{}, erro error) {
	var option interface{}
	option, err := bson.Marshal(loanInfo)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, erro = InsertOneDB(ctx, shared.Database, shared.LoanDetailCollection, option)
	return result, erro
}
