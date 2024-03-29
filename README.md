# Microfinance Helper

This is a microfinancing monitoring and helper application which helps to keep track and provide score (like CIBIL score) for the lender as well as the borrower. 

## What is Micro Finance ?
Microfinance is a banking service provided to unemployed or low-income individuals or groups who otherwise would have no other access to financial services. 

## Why is it needed?
Microfinance allows people to take on reasonable small business loans safely, and in a manner that is consistent with ethical lending practices.
The majority of microfinancing operations occur in developing nations.

## What are the challenges faced by current system?
- The process is informal.
- There exist no authoritative rating mechanism exisiting for the sector.
- Largely unmonitored.

## What issues are we trying to address?
- Monitor loans with does not involve a Bank or NBFC.
- Monitoring the flow of loan money 
- Rating the payment capability and trust worthiness for the borrower and the lender. 

## How rating works; ToDo - work more on this idea.
- On every defaulting of loan, the score will reduce for the borrower.
- Depending on the loan amount, the score of both borrower and lender will increase. 
- On every successful loan payment, the score of the borrower and lender will increase.
- IEEE paper to try out - 'Credit Scoring Model Implementation in a Microfinance Context'

## Integrate e-Rupi as a payment method -ToDo
- e-rupi doesnt involve any intermediate parties hence this should be incooperated.

## Activity Diagrams
<p align="center">
<img src="./docs/img/Applying for loan.png" alt="alt text" width="500px" height="400px">
</p>
<p align="center"><b>Activity flow for applying a loan</b></p>

<p align="center">
<img src="./docs/img/Loan_payment.png" alt="alt text" width="500px" height="400px">
</p>
<p align="center"><b>Activity flow for loan repayment</b></p>

<p align="center">
<img src="./docs/img/Missed_Payment.png" alt="alt text" width="500px" height="600px">
</p>
<p align="center"><b>Activity flow for missing payment</b></p>

## TODO Discuss with Legendary people :P

Learn more

## Technology stack
- React.js
- Golang
- MongoDB
- Ethereum Blockchain
