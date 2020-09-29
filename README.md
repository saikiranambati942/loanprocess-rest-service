# Loan-Process Rest Service

## 🚩 Table of Contents

 - [Description](#description)
 - [Features](#features)
 - [Repository Structure](#repository-structure)
 - [How to Run and Test?](#how-to-run-and-test)
 - [Sample Execution Result](#sample-execution-result)
 - [TO DO](#to-do)

## Description
This is a REST api service for loan process simulation that supports three operations(loaninitiation, loanpayment, loanbalance), 
implemented using three handlers:
1) LoanInitiate
2) Payment
3) GetBalance

## Features
1) Can initiate a loan on any date with specified amount.
2) Can add payments in any order.
3) Request balance on any date.
4) Interest calculation is based on principal balance and excludes already added interest balance.
5) The interest added for a day is calculated using: (annual interest rate * principal balance)/(100 * 365).
6) The balance returned is based on all payments and the interest added up to the requested date.
#### Limitation: 
Service does not handle multiple loans, i.e. the state will be cleared when a new loan is initiated.



## Repository Structure

### Directory tree
    . 
    ├── cmd
    |      └──loanapiserver
    |             └──loanapiserver.go             # main function of application i.e; loanprocess-rest-service starts here
    |                         
    ├── internal                                  # handlers for loan initiation, add payment and get balance information 
    |      └──handlers
    |              └──handlers.go
    |              └──loaninitiate_handler.go
    |              └──loanpayment_handler.go
    |              └──loanbalance_handler.go
    |                
    ├── test                                      # unit tests for handlers     
    |      └──loaninitiate_handler_test.go                             
    |      └──loanpayment_handler_test.go
    |      └──loanbalance_handler_test.go
    |
    |
    ├── docs                                     # contains screenshots of run results for user reference 
    |      └──images                        
    |                  
    ├── vendor                                   # contains application dependencies
    └── README.md

[`cmd`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/cmd "API documentation") package:
------------------------------------------------------------------------------------------------------------------

 `cmd` package is the initial point of the application where `loanapiserver` is the placeholder for loanapiserver.go(starting point of application)


[`internal`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/internal "API documentation") package:
----------------------------------------------------------------------------------------------------------------------------

 `internal` package contains the private code internal to our application, has below `handlers`:

LoanInitiate handler handles the requests of initiation of loan (`loaninitiate_handler.go`)

Payment handler handles the loan repayment requests (`loanpayment_handler.go`)

GetBalance handler handles the requests to retrieve the remaining balance  on a specified date (`loanbalance_handler.go`)


[`test`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/test "API documentation") package:
--------------------------------------------------------------------------------------------------------------------

`test` package contains the unit test cases covered for all the application code functionalities


[`vendor`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/vendor "API documentation") package:
------------------------------------------------------------------------------------------------------------------------

`vendor` folder contains application dependencies, which includes all the packages needed to support builds and tests of application


## How to Run and Test?
After cloning the repository (https://github.com/saikiranambati942/loanprocess-rest-service.git), run the below command from the root directory to start the http server on localhost:8080

```
go run cmd/loanapiserver/loanapiserver.go
```


To initiate the loan, trigger the "/loaninitiate" endpoint with the below request format:
```
{
  "loanamount": float64,
  "interest":float64,
  "startdate":string (YYYY-MM-DD format)
  
}
```
For example:
```
{
  "loanamount": 5000,
  "interest":5,
  "startdate":"2020-02-02"
  
}
```
Now to add a payment on a particular date, trigger the "/payment" endpoint with the below request format:

```
{
     "amount": float64,
     "date": string (YYYY-MM-DD format)
}
   ```
For example

```
{
 "amount": 1000, 
 "date": "2020-02-20"
}
```
To fetch remaining balance on a particular date, trigger the "/getbalance"  endpoint with the below request format:

```
{
     "date": string (YYYY-MM-DD format)
}
   ```

For example   

```
{ 
 "date": "2020-02-27"
}
```

#### Please note that: Month, day values may be outside their usual ranges and will be normalized during the conversion(handled by builtin Date function in Go)
     For example, October 32 converts to November 1.

## Sample Execution Result
#### Step1: Start server using the below command from the root directory

```
go run cmd/loanapiserver/loanapiserver.go
```
#### Step2: Initiate the loan with input data

![](https://github.com/saikiranambati942/loanprocess-rest-service/blob/master/docs/images/loaninitiate.png)

#### Step3: Add a payment on any date

![](https://github.com/saikiranambati942/loanprocess-rest-service/blob/master/docs/images/loanpayment.png)

#### Step4: Check balance on any date

##### Checking balance on payment date 

![](https://github.com/saikiranambati942/loanprocess-rest-service/blob/master/docs/images/loanbalance_on_paymentdate.png)

##### Checking balance randomly before payment date 

![](https://github.com/saikiranambati942/loanprocess-rest-service/blob/master/docs/images/loanbalance_before_paymentdate.png)

##### Checking balance randomly after payment date 

![](https://github.com/saikiranambati942/loanprocess-rest-service/blob/master/docs/images/loanbalance_after_paymentdate.png)





## TO DO
1) Welcome page.
2) Adding authorization and authentication functionality for a user.
3) To handle multiple loan requests by integrating with a database. 
4) Dockerize the application.





