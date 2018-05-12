package repository

import (
	"testing"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"fmt"
	"github.com/goinggo/tracelog"
	"github.com/spf13/viper"
	pb_device "github.com/onezerobinary/db-box/proto/device"
)

func startConfig(){
	viper.SetConfigName("config")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "database_test", "StartConfig", "Error reading config file")
	}
}

func TestConnectToDB(t *testing.T){

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	db, err := ConnectToDB()

	if err != nil  {
		t.Error("Wrong DB or error in connection: ", db)
	}

	tracelog.Completed("testDB", "TestConnectToDB")
}

func TestAddDoc(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	//fakeToken := pb_account.Token{"fff"}
	fakeStatus := pb_account.Status{pb_account.Status_DISABLED}

	username := "enrico@enrico.com"
	password := "enrico"

	faketoken := GenerateToken(username, password)

	fakeAccount := pb_account.Account{
		faketoken,
		username,
		password,
		nil,
		&fakeStatus,
		"Account",
		"2018-01-11",
		"2028-01-10",
		nil,
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

	startConfig()

	fakeCredentials := pb_account.Credentials{}

	fakeCredentials.Username = "enrico@enrico.com"
	fakeCredentials.Password = "enrico"

	token := pb_account.Token{}

	to := GenerateToken(fakeCredentials.Username, fakeCredentials.Password)

	token.Token = to
	fakeCredentials.Token = &token

	account, err := GetAccountByCredentials(fakeCredentials)

	if err != nil {
		t.Error("Error in getting the informations", err)
	}

	p := IsPresent(account.Uuid)

	if !p {
		t.Error("Element not present: ", account.Uuid)
	}
}

func TestGetAccountByCredentials(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeCredentials := pb_account.Credentials{}

	fakeCredentials.Username = "enrico@enrico.com"
	fakeCredentials.Password = "enrico"

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

func TestGetAccountByToken(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeCredentials := pb_account.Credentials{}

	fakeCredentials.Username = "enrico@enrico.com"
	fakeCredentials.Password = "enrico"

	token := pb_account.Token{}

	to := GenerateToken(fakeCredentials.Username, fakeCredentials.Password)

	token.Token = to

	account, err := GetAccountByToken(token)

	if err != nil {
		t.Error("Error in getting the informations", err)
	}

	if account == nil {
		t.Errorf("Error, no account found")
	}
}

func TestUpdateDoc(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	// Create and add an account to the DB
	fakeStatus := pb_account.Status{pb_account.Status_ENABLED}

	username := "enrico@enrico.com"
	password := "enrico"

	faketoken := GenerateToken(username, password)

	fakeAccount := pb_account.Account{
		faketoken,
		username,
		password,
		nil,
		&fakeStatus,
		"Account",
		"2018-01-11",
		"2028-01-10",
		nil,
	}

	token, err :=  AddDoc(fakeAccount)

	tk := pb_account.Token{Token:token}

	account, err := GetAccountByToken(tk)

	account.Token.Token = token
	//change the something to the account
	account.Status = &pb_account.Status{pb_account.Status_SUSPENDED}

	err = UpdateDoc(*account)

	if err != nil {
		t.Errorf("Error to update the account")
	}
}

func TestCheckEmail(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	email := pb_account.Email{ "Mary"}

	token, err := CheckEmail(email)

	if err != nil {
		t.Errorf("Error to get the email back")
	}

	if token != nil {
		response := "Token: " + token.Token
		tracelog.Trace("", "", response)
	}
}

func TestSetAccountStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	token := pb_account.Token{Token:"2284fe70432bbef5a5354653c88d8e5cda2880dd"}
	account, err := GetAccountByToken(token)

	if err != nil {
		tracelog.Errorf(err,"database_test", "TestSetAccountStatus", "Account not found")
	}

	if account != nil {

		account.Status = &pb_account.Status{pb_account.Status_ENABLED}
		accountUpdateStatus := pb_account.UpdateStatus{&token, account.Status}

		err = SetAccountStatus(accountUpdateStatus)

		if err != nil {
			tracelog.Errorf(err,"database_test", "TestSetAccountStatus", "Account status not updated")
			t.Errorf("Status not updated")
		}
	}
}

func TestGetAccountStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	token := pb_account.Token{Token:"2284fe70432bbef5a5354653c88d8e5cda2880dd"}
	account, err := GetAccountByToken(token)

	if err != nil {
		tracelog.Errorf(err,"database_test", "TestGetAccountStatus", "Account not found")
	}

	if account != nil {
		accountStatus, err := GetAccountStatus(token)

		if err != nil {
			tracelog.Errorf(err,"database_test", "TestGetAccountStatus", "Account Status not found")
		}

		if accountStatus != nil {
			text := "Account Status retrieved: " + accountStatus.Status.String()
			tracelog.Trace("database_test", "TestGetAccountStatus" , text)
		} else {
			t.Errorf("Account status not retrieved")
		}
	}
}

func TestGetAccountsByStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	accountStatus := pb_account.Status{pb_account.Status_ENABLED}

	acconts, err := GetAccountsByStatus(accountStatus)

	if err != nil {
		tracelog.Errorf(err,"database_test", "TestGetAccountsByStatus", "Accounts not found")
	}

	if acconts != nil {
		text := "Acounts retrieved with status: " + accountStatus.Status.String()
		tracelog.Trace("database_test", "TestGetAccountsByStatus" , text)
	} else {
		t.Errorf("Accounts not retrieved")
	}

}

func TestGetAccounts(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	acconts, err := GetAccounts()

	if err != nil {
		tracelog.Errorf(err,"database_test", "TestGetAccounts", "Accounts not found")
	}

	if acconts != nil {
		text := "Accounts retrieved"
		tracelog.Trace("database_test", "TestGetAccounts" , text)
	} else {
		t.Errorf("Accounts not retrieved")
	}
}

func TestAddExpoPushToken(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	token := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}

	accountPushToken := pb_account.ExpoPushToken{}
	accountPushToken.Token = &token
	accountPushToken.Expotoken = "ExponentPushToken[uaT9EBFnTnEUW78GfiSVaO]"

	resp, err := AddExpoPushToken(&accountPushToken)

	if err != nil {
		t.Errorf("Error: It was not possible to add the expoToken. ", err)
	}

	if !resp.Response {
		t.Errorf("Response not valid! ExpoToken not added. ")
	}
}

func TestRemoveDoc(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeCredentials := pb_account.Credentials{}

	fakeCredentials.Username = "enrico@enrico.com"
	fakeCredentials.Password = "enrico"

	token := pb_account.Token{}

	to := GenerateToken(fakeCredentials.Username, fakeCredentials.Password)

	token.Token = to

	err := RemoveDoc(token)

	if err != nil {
		t.Errorf("Error to delete the account")
	}
}


// Test Device methods
func TestDeviceIsPresent(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeDeviceToken := pb_device.ExpoPushToken{}

	fakeDeviceToken.Expopushtoken = "abcd"

	p := DeviceIsPresent(fakeDeviceToken.Expopushtoken)

	if !p {
		t.Error("Element not present: ", fakeDeviceToken.Expopushtoken)
	}
}

func TestAddDevice(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	expoPushToken := pb_device.ExpoPushToken{ "ExponentPushToken[VqalPOCUT5DVmVUpf6Qq3A]"}

	fakeDevice := pb_device.Device{
		Expopushtoken: &expoPushToken,
		Type: "Device",
		Active: true,
		Latitude: 12.999,
		Longitude: 45.3456,
		Mobilenumber: "123123456789",
	}

	ok, _ := AddDevice(&fakeDevice)

	if (!ok.Response) {
		t.Error("It was not possible to store the document into the DB")
	}
}

func TestGetDeviceByExpoToken(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeExpoPushToken := pb_device.ExpoPushToken{ "ExponentPushToken[VqalPOCUT5DVmVUpf6Qq3A]"}

	dev, _ := GetDeviceByExpoToken(&fakeExpoPushToken)

	if (dev.Expopushtoken.Expopushtoken != fakeExpoPushToken.Expopushtoken) {
		t.Error("It was not possible to get the device back")
	} else {
		fmt.Println("Device info:", dev.Expopushtoken)
	}
}


func TestUpdateStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeExpoPushToken := pb_device.ExpoPushToken{ "ExponentPushToken[VqalPOCUT5DVmVUpf6Qq3A]"}
	status := pb_device.Status{}
	status.Expopushtoken = &fakeExpoPushToken
	status.Active = true

	updated, _ := UpdateStatus(&status)

	if (!updated.Response) {
		t.Error("It was not possible to updated the device")
	} else {
		fmt.Println("Device updated? ", updated.Response)
	}
}

func TestUpdatePosition(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeExpoPushToken := pb_device.ExpoPushToken{ "ExponentPushToken[VqalPOCUT5DVmVUpf6Qq3A]"}

	position := pb_device.Position{}
	position.Expopushtoken = &fakeExpoPushToken
	position.Latitude = 19.999
	position.Longitude = 07.1987

	updated, _ := UpdatePosition(&position)

	if !updated.Response {
		t.Error("It was not possible to updated the device")
	} else {
		fmt.Println("Device updated? ", updated.Response)
	}
}

func TestUpdateMobileNumber(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	startConfig()

	fakeExpoPushToken := pb_device.ExpoPushToken{ "ExponentPushToken[VqalPOCUT5DVmVUpf6Qq3A]"}

	mobile := pb_device.MobileNumber{}
	mobile.Expopushtoken = &fakeExpoPushToken
	mobile.Mobilenumber = "987654321"

	updated, _ := UpdateMobileNumber(&mobile)

	if !updated.Response {
		t.Error("It was not possible to updated the device")
	} else {
		fmt.Println("Device updated? ", updated.Response)
	}
}



