package repository

import (
	"testing"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"fmt"
	"github.com/goinggo/tracelog"
)

func TestConnectToDB(t *testing.T){

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	db, err := ConnectToDB()

	if err != nil  {
		t.Error("Wrong DB or error in connection: ", db)
	}

	tracelog.Completed("testDB", "TestConnectToDB")
}

func TestAddDoc(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//fakeToken := pb_account.Token{"fff"}
	fakeStatus := pb_account.Account_Status(pb_account.Account_DISABLED)

	fakeAccount := pb_account.Account{
		"1234",
		"foo",
		"bar",
		nil,
		fakeStatus,
		"Account",
		"2018-01-11",
		"2028-01-10",
	}

	token, err :=  AddDoc(fakeAccount)

	if err != nil {
		t.Error("non possible to store the document into the DB")
	}

	fmt.Println("ok, token generated: ", token)
}

func TestIsPresent(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	uuid := "1234"

	p := IsPresent(uuid)

	if !p {
		t.Error("Element not present: ", uuid)
	}
}

func TestGetAccountByCredentials(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	fakeCredentials := pb_account.Credentials{}

	fakeCredentials.Username = "foo"
	fakeCredentials.Password = "bar"

	token := pb_account.Token{}

	to := GenerateToken(fakeCredentials.Username, fakeCredentials.Password)

	token.Token = to
	fakeCredentials.Token = &token

	account, err := GetAccountByCredentials(fakeCredentials)

	if err != nil {
		t.Error("Error in getting the informations", err)
	}

	if account == nil {
		t.Errorf("Error, no account found")
	}
}
