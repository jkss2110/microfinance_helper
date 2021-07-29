package models

type LoanDetail struct {
	Lender       int
	Borrower     int
	LenderName   string
	BorrowerName string
	LoanAmt      int
	LoanEMI      int
}

var LoanDetails = []LoanDetail{
	{
		Lender:       1,
		Borrower:     2,
		LenderName:   "JK",
		BorrowerName: "KJ",
		LoanAmt:      1000,
		LoanEMI:      500,
	},
	{
		Lender:       2,
		Borrower:     3,
		LenderName:   "JK1",
		BorrowerName: "KJ1",
		LoanAmt:      1500,
		LoanEMI:      500,
	},
	{
		Lender:       3,
		Borrower:     4,
		LenderName:   "JK3",
		BorrowerName: "KJ4",
		LoanAmt:      100,
		LoanEMI:      50,
	},
}
