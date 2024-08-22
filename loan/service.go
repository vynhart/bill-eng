package loan

// This service object is supposed to communicate with loan service 
// to get the loan data and translate it to Loan object
//
// It is supposed to be used in the main function that serve 
// HTTP, gRPC, background job or any other client code in conjunction with
// the main bill service.

// Not implemented yet
type Service struct {}

func (srv *Service) GetLoan(id string) Loan {
	return Loan{}
}