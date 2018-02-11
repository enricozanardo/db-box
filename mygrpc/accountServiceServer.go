package mygrpc

import (
	"context"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"errors"
	"github.com/onezerobinary/db-box/repository"
)

/*
	200 -> OK
	201 -> Created
	302 -> Found
	400 -> Bad Request
	401 -> Unauthorized
	500 -> Internal Server Error
 */


type AccountServiceServer struct {

}

//Create a new Account
func (s *AccountServiceServer) CreateAccount (ctx context.Context, account *pb_account.Account) (*pb_account.Response, error) {

	response := pb_account.Response{}

	//Check if already present into the Database
	isAlreadyPresent := repository.IsPresent(account.Uuid)

	if isAlreadyPresent {
		response = pb_account.Response{ 302, account.Token}
		return &response, errors.New("Record already present into the system")
	}

	// Add a doc
	stringToken, err := repository.AddDoc(*account)

	// Set up response
	token := pb_account.Token{stringToken}

	if err != nil {
		response = pb_account.Response{ 400, &token}
		return &response, err
	}

	response = pb_account.Response{ 200, &token}
	// Send back the response
	return &response, nil
}

func (s *AccountServiceServer) GetAccountByCredentials(ctx context.Context, credentials *pb_account.Credentials) (*pb_account.Account, error) {

	account, err := repository.GetAccountByCredentials(*credentials)

	if err != nil {
		emptyAccount := pb_account.Account{}
		account = &emptyAccount
		return account, err
	}

	return account, nil
}

// Get an Account given the Token
func (s *AccountServiceServer) GetAccountByToken (ctx context.Context, token *pb_account.Token)  (*pb_account.Account, error) {

	account, err := repository.GetAccountByToken(*token)

	if err != nil {
		account = pb_account.Account{}
		return &account, err
	}

	return &account, nil
}

// Delete an Account given the Token
func (s *AccountServiceServer) DeleteAccount (ctx context.Context, token *pb_account.Token)  (*pb_account.Response, error) {

	response := pb_account.Response{}

	// Remove a specific doc
	err := repository.RemoveDoc(*token)

	if err != nil {
		response = pb_account.Response{ 400, token}
		return &response, err
	}

	response = pb_account.Response{ 200, token}

	return &response, nil
}

// Update an Account given the updated Account
func (s *AccountServiceServer) UpdateAccount (ctx context.Context, account *pb_account.Account) (*pb_account.Response, error) {

	response := pb_account.Response{}

	// Update an existing doc
	err := repository.UpdateDoc(*account)

	if err != nil {
		response = pb_account.Response{ 400, account.Token}
		return &response, err
	}

	response = pb_account.Response{ 200, account.Token}

	return &response, nil
}

// Check if an email address is already used
func (s *AccountServiceServer) CheckEmail (ctx context.Context, email *pb_account.Email)  (*pb_account.Response, error) {

	response := pb_account.Response{}

	// check if an email is already present into the system and then return the token
	token, err := repository.CheckEmail(*email)

	if err != nil {
		// No email found
		response = pb_account.Response{ 400, &token}
		return &response, err
	}

	// Email found
	response = pb_account.Response{ 200, &token}

	return &response, nil
}

// Get the Status of an account given the Token
func (s *AccountServiceServer) GetAccountStatus (ctx context.Context, token *pb_account.Token)  (*pb_account.Account_Status, error) {

	// Get the status of an account
	accountStatus, err := repository.GetAccountStatus(*token)

	if err != nil {
		// No email found
		accountStatus = pb_account.Account_Status(pb_account.Status_NOTSET)
		return &accountStatus, err
	}

	return &accountStatus, nil
}

// Set the Status of an account given the Updated Status
func (s *AccountServiceServer) SetAccountStatus (ctx context.Context, updateStatus *pb_account.UpdateStatus)  (*pb_account.Response, error) {

	response := pb_account.Response{}

	err := repository.SetAccountStatus(*updateStatus)

	if err != nil {
		// No email found
		response = pb_account.Response{ 400, updateStatus.Token}
		return &response, err
	}

	// Email found
	response = pb_account.Response{ 200, updateStatus.Token}

	return &response, nil

}

// Get all the accounts based on a specific Status
func (s *AccountServiceServer) GetAccountsByStatus (ctx context.Context, status *pb_account.Status)  (*pb_account.Accounts, error) {

	accounts, err := repository.GetAccountsByStatus(*status)

	if err != nil {
		//accounts = pb_account.Accounts{}
		return &accounts, err
	}

	return &accounts, nil
}

//Get the account collection
func (s *AccountServiceServer) GetAccounts (ctx context.Context) (*pb_account.Accounts, error) {

	accounts, err := repository.GetAccounts()

	if err != nil {
		accounts = pb_account.Accounts{}
		return &accounts, err
	}

	return &accounts, nil
}

