# Loan-Process Rest Service
## Description
This is a REST api service for simulating loan processing that supports three operations(loaninitiation, loanpayment, loanbalance).
This service contains three handlers:
1) LoanInitiate
2) Payment
3) GetBalance

## Features
1) Can initiate a loan on any date with specified amount.
2) Can add payments in any order.
3) Request balance on any date.
4) Interest calculation should be based on the principal balance and exclude already added interest balance.
5) The interest added for a day is defined as: annual interest rate / 100 / 365 * principal balance.
6) The balance returned is based on all payments and the interest added up to the requested date.
7) Service does not handle multiple loans, i.e. the state will be cleared when a new loan is initiated.




## Repository Structure:
-----------------------
### Directory tree
    . 
    ├── cmd
    |      |__server
    |             |__server.go             # main function of application i.e; loanprocess-rest-service starts here
    |                         
    ├── internal                           # handlers for loan initiation, add payment and get balance information 
    |      |__handlers
    |              |__handlers.go
    |              |__loaninitiate_handler.go
    |              |__loanpayment_handler.go
    |              |__loanbalance_handler.go
    |                
    ├── test                                # unit tests for handlers     
    |      |__loaninitiate_handler_test.go                             
    |      |__loanpayment_handler_test.go
    |      |__loanbalance_handler_test.go
    |
    |
    ├── docs                                 # contains screenshots of run results for user reference 
    |      |__images                        
    |                  
    ├── vendor                               # contains application dependencies
    └── README.md

[`cmd`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/cmd "API documentation") package:
------------------------------------------------------------------------------------------------------------------

The `cmd` package is the starting point of our application. This has a server folder which contains server.go file where the main function of our application.


[`internal`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/internal "API documentation") package:
----------------------------------------------------------------------------------------------------------------------------

The `internal` package contains the source code which is internal to our application. 
This has `handlers` folder where our application code exists.

`handlers` folder contains the below handlers:
LoanInitiate handler handles the requests of initiation of loan (`loaninitiate_handler.go`)

Payment handler handles the loan repayment requests (`loanpayment_handler.go`)

GetBalance handler handles the requests to retrieve the remaining balance  on a specified date (`loanbalance_handler.go`)


[`test`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/test "API documentation") package:
--------------------------------------------------------------------------------------------------------------------

`test` package contains the test cases covered for all the application code functionalities


[`vendor`](https://github.com/saikiranambati942/loanprocess-rest-service/tree/master/vendor "API documentation") package:
------------------------------------------------------------------------------------------------------------------------

The `vendor` folder contains the application dependencies. All the packages needed to support builds and tests of application are included in this folder


## Execution:
After cloning the repository (https://github.com/saikiranambati942/loanprocess-rest-service.git),  run the below command

```
go run cmd/server/server.go
```

This will start the http server on localhost on port 8080.

To initiate the loan, trigger the "/loaninitiate" endpoint with the below data:
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
Now to add a payment on a particular date, trigger the "/payment" endpoint with the below data:

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
To fetch remaining balance on a particular date, trigger the "/getbalance"  endpoint with the below data:

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

## Sample Execution Result
#### Step1: Start server using the below command

```
go run cmd/server/server.go
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





## TODO



