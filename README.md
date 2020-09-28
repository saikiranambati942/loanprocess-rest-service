# Loan-Process Rest Service
## Description:
 This is a REST api service for simulating loan processing that supports three operations(loaninitiation, loanpayment, loanbalance).
This service contains three handlers:
1) LoanInitiate
2) Payment
3) GetBalance




## Repository Structure:
-----------------------
### Directory tree
    . 
    ├── cmd
    |      |__server
    |             |__server.go   
    |                         
    ├── internal
    |      |__handlers
    |              |__handlers.go
    |              |__loaninitiate_handler.go
    |              |__loanpayment_handler.go
    |              |__loanbalance_handler.go
    |                
    ├── test
    |      |__handlers.go
    |      |__loaninitiate_handler.go
    |      |__loanpayment_handler.go
    |      |__loanbalance_handler.go  
    |                  
    ├── vendor                   
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
Step1: Start server using the below command
```
go run cmd/server/server.go
```
Step2: Initiate the loan with input data



## TODO



