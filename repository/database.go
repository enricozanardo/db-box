package repository

import (
	"log"
	"github.com/hyperledger/fabric/core/ledger/util/couchdb"
	"fmt"
	"os"
	"time"
	"github.com/spf13/viper"
	"crypto/sha1"
	"encoding/hex"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"encoding/json"
	"io"
	"github.com/goinggo/tracelog"
	"errors"
)

type DBSettings struct {
	username string
	password string
	url string
	name string
}

func fetchError(message string, error error){
	log.Fatalf(message, error)
}

func writeLog(message string, v ... interface{}){
	log.Printf(message, v)
}

func GetSettings() (dbSettings DBSettings){

	DBuser := os.Getenv("COUCHDB_USER")
	DBpassword := os.Getenv("COUCHDB_PASSWORD")
	DBaddress := os.Getenv("COUCHDB_PORT_5984_TCP_ADDR") //  localhost
	DBport := os.Getenv("COUCHDB_PORT_5984_TCP_PORT") //  5984
	DBname := os.Getenv("DBNAME")
	DBurl := DBaddress + ":" + DBport

	// if no production settings are found use the local settings
	if len(DBaddress) == 0 {

		//development environment
		viper.SetConfigName("config")
		// Internal tests
		//viper.AddConfigPath("../.")

		// Remote tests
		viper.AddConfigPath("./")

		if err := viper.ReadInConfig(); err != nil {
			fetchError("Error reading config file %s", err)
		}
		// Confirm which config file is used
		writeLog("Using config: %s\n", viper.ConfigFileUsed())
		fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

		// Set the variables reading from the config.yaml file
		DBuser = viper.GetString("database.username")
		DBpassword = viper.GetString("database.password")
		DBurl = viper.GetString("database.address")
		DBname = viper.GetString("database.name")

		tracelog.Warning("database", "GetSettings", "####### Development #########")
	}

	dbSettings.username = DBuser
	dbSettings.password = DBpassword
	dbSettings.url = DBurl
	dbSettings.name = DBname

	return
}

func SetDBInstance(appEnv DBSettings) (cDBInstance couchdb.CouchInstance, dbName string) {

	var dbSetting = couchdb.CouchDBDef {
		URL: appEnv.url,
		Username: appEnv.username,
		Password: appEnv.password,
		MaxRetries: 3,
		MaxRetriesOnStartup: 10,
		RequestTimeout: time.Second*35,
	}

	couchDBInstance, err := couchdb.CreateCouchInstance(dbSetting.URL, dbSetting.Username, dbSetting.Password,
		dbSetting.MaxRetries, dbSetting.MaxRetriesOnStartup, dbSetting.RequestTimeout)

	if err != nil {
		tracelog.Errorf(err, "database", "SetDBInstance", "Unexpected error")
		os.Exit(1)
	}

	return *couchDBInstance, appEnv.name
}

func ConnectToDB() (db couchdb.CouchDatabase, err error) {

	appEnv := GetSettings()

	// Create a new connection
	couchDBInstance, dbName := SetDBInstance(appEnv)

	db = couchdb.CouchDatabase{CouchInstance: couchDBInstance, DBName: dbName}

	_, err = db.CreateDatabaseIfNotExist()

	if err != nil {
		tracelog.Errorf(err, "database", "ConnectToDB", "Error to create the DB")
		return db, err
	}

	tracelog.Trace("database","ConnectToDB","Connected to the DB")

	return db, nil
}

func IsPresent(uuid string) bool {

	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "IsPresent", "Error to connect to the DB")
		os.Exit(1)
	}

	queryString := "{\"selector\":{\"uuid\":{\"$eq\":\""+ uuid +"\"}}}"

	queryResults, err := db.QueryDocuments(queryString)

	if err != nil {
		tracelog.Errorf(err, "database", "IsPresent", "Error to search doc to the DB")
		os.Exit(1)
	}

	present := false

	for k, _ := range *queryResults {
		if k > 0 {
			tracelog.Errorf(err, "database", "IsPresent", "Error more then one entry found in the DB")
			os.Exit(1)
		}

		tracelog.Trace("database", "IsPresent", "Account is present")
		present = true
	}

	tracelog.Completed("database","IsPresent")
	return present
}

func AddDoc(account pb_account.Account) (token string, dbErr error) {

	token = GenerateToken(account.Username, account.Password)

	//add the missing info to the account
	pb_token := pb_account.Token{token}
	account.Token = &pb_token

	// Marshal the document in Json
	jsonDoc, _ := json.Marshal(account)

	rev := ""

	db, _ := ConnectToDB()

	// Store the document into the DB
	_, err := db.SaveDoc(account.Uuid, rev, & couchdb.CouchDoc{JSONValue: jsonDoc, Attachments: nil})

	if err != nil {
		tracelog.Errorf(err, "database", "AddDoc", "Error to add doc to the DB")
	}

	tracelog.Completed("database","AddDoc")
	return token, nil
}

func GenerateToken(username string, password string) (token string){
	// Create the hash sha1 of the username and password
	h1 := sha1.New()
	io.WriteString(h1, username)
	io.WriteString(h1, password)

	token = hex.EncodeToString(h1.Sum(nil))

	tracelog.Completed("database","GenerateToken")
	return token
}

func GetAccountByCredentials(credentials pb_account.Credentials) (account *pb_account.Account, dberr error){

	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountByCredentials", "Error to connect to the DB")
		os.Exit(1)
	}

	token := GenerateToken(credentials.Username, credentials.Password)

	queryString := "{\"selector\":{\"token.token\":{\"$eq\":\""+ token +"\"}}}"

	queryResults, err := db.QueryDocuments(queryString)

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountByCredentials", "Error to search doc to the DB")
	}

	for k, v := range *queryResults {
		if k > 0 {
			tracelog.Errorf(err, "database", "GetAccountByCredentials", "Error more then one entry found in the DB")
			os.Exit(1)
		}

		//account found!
		value := v.Value
		err := json.Unmarshal(value[:], &account)

		if err != nil {
			tracelog.Errorf(err, "database", "GetAccountByCredentials", "Error to get the doc from the DB")
			os.Exit(1)
		}

		tracelog.Trace("database","GetAccountByCredentials","Account found")
		return account, nil
	}
	//account not found
	tracelog.Warning("database", "GetAccountByCredentials", "Account not found, return empty account")

	fakeAccount := genFakeAccount()
	return fakeAccount, err
}

func GetAccountByToken(token pb_account.Token) (account *pb_account.Account, dberr error) {

	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountByToken", "Error to connect to the DB")
		os.Exit(1)
	}

	stringToken := token.Token

	queryString := "{\"selector\":{\"token.token\":{\"$eq\":\""+ stringToken +"\"}}}"

	queryResults, err := db.QueryDocuments(queryString)

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountByToken", "Error to search doc to the DB")
	}

	for k, v := range *queryResults {
		if k > 0 {
			tracelog.Errorf(err, "database", "GetAccountByToken", "Error more then one entry found in the DB")
			os.Exit(1)
		}

		//account found!
		value := v.Value
		err := json.Unmarshal(value[:], &account)

		if err != nil {
			tracelog.Errorf(err, "database", "GetAccountByToken", "Error to get the doc from the DB")
			os.Exit(1)
		}

		tracelog.Trace("database","GetAccountByToken","Account found")
		return account, nil
	}
	//account not found
	tracelog.Warning("database", "GetAccountByToken", "Account not found, return empty account")
	//Return a fake account
	fakeAccount := genFakeAccount()
	return fakeAccount, err
}

func genFakeAccount() (fakeAccount *pb_account.Account) {
	token := GenerateToken("fake", "fake")
	fakeToken := pb_account.Token{token}
	fakeStatus := pb_account.Status{pb_account.Status_NOTSET}

	fakeAccount = &pb_account.Account{}
	fakeAccount.Token = &fakeToken
	fakeAccount.Status = &fakeStatus

	return fakeAccount
}


func RemoveDoc(token pb_account.Token) (err error){

	db, err := ConnectToDB()

	_, revision, _ :=  db.ReadDoc(token.Token)

	err = db.DeleteDoc(token.Token, revision)

	if err != nil {
		tracelog.Errorf(err, "database", "RemoveDoc", "Error to delete doc from the DB")
		os.Exit(1)
	}

	return
}

func UpdateDoc(account pb_account.Account) (err error){

	db, err := ConnectToDB()
	// Get revision
	_, revision, _ :=  db.ReadDoc(account.Uuid)

	// Marshal the document in Json
	jsonDoc, _ := json.Marshal(account)

	// Store the document into the DB
	text, err := db.SaveDoc(account.Uuid, revision, & couchdb.CouchDoc{JSONValue: jsonDoc, Attachments: nil})

	tracelog.Trace("database", "UpdateDoc", text)

	if err != nil {
		tracelog.Errorf(err, "database", "RemoveDoc", "Error to update account to the DB")
		os.Exit(1)
		return
	}

	return
}

func CheckEmail(email pb_account.Email) (token *pb_account.Token, dberr error){

	account := pb_account.Account{}

	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "CheckEmail", "Error to connect to the DB")
		os.Exit(1)
	}

	stringEmail := email.Email

	queryString := "{\"selector\":{\"username\":{\"$eq\":\""+ stringEmail +"\"}}}"

	queryResults, err := db.QueryDocuments(queryString)

	if err != nil {
		tracelog.Errorf(err, "database", "CheckEmail", "Error to search email to the DB")
	}

	for k, v := range *queryResults {
		if k > 0 {
			tracelog.Errorf(err, "database", "CheckEmail", "Error more then one entry found in the DB")
			os.Exit(1)
		}

		//account found!
		value := v.Value
		err := json.Unmarshal(value[:], &account)

		if err != nil {
			tracelog.Errorf(err, "database", "CheckEmail", "Error to get the doc from the DB")
			os.Exit(1)
		}

		tracelog.Trace("database","CheckEmail","Email found")
		return account.Token, nil
	}
	//account not found
	tracelog.Warning("database", "CheckEmail", "Email not found, return nil")
	dberr = err
	return nil, dberr
}

func GetAccountStatus(token pb_account.Token) (accountStatus *pb_account.Status, dberr error){
	account := pb_account.Account{}

	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountStatus", "Error to connect to the DB")
		os.Exit(1)
	}

	stringToken := token.Token

	queryString := "{\"selector\":{\"token.token\":{\"$eq\":\""+ stringToken +"\"}}}"

	queryResults, err := db.QueryDocuments(queryString)

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountStatus", "Error to search account status into the DB")
	}

	for k, v := range *queryResults {
		if k > 0 {
			tracelog.Errorf(err, "database", "GetAccountStatus", "Error more then one entry found in the DB")
			os.Exit(1)
		}

		//account found!
		value := v.Value
		err := json.Unmarshal(value[:], &account)

		if err != nil {
			tracelog.Errorf(err, "database", "GetAccountStatus", "Error to get the doc from the DB")
			os.Exit(1)
		}

		tracelog.Trace("database","GetAccountStatus","Email found")

		return account.Status, nil
	}
	//account not found
	tracelog.Warning("database", "GetAccountStatus", "Email not found, return nil")
	dberr = err
	return nil, dberr
}

func SetAccountStatus(updateStatus pb_account.UpdateStatus)(dberr error){
	account := pb_account.Account{}

	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "SetAccountStatus", "Error to connect to the DB")
		os.Exit(1)
	}

	stringToken := updateStatus.Token.Token

	queryString := "{\"selector\":{\"token.token\":{\"$eq\":\""+ stringToken +"\"}}}"

	queryResults, err := db.QueryDocuments(queryString)

	if err != nil {
		tracelog.Errorf(err, "database", "SetAccountStatus", "Error to search account status into the DB")
	}

	for k, v := range *queryResults {
		if k > 0 {
			tracelog.Errorf(err, "database", "SetAccountStatus", "Error more then one entry found in the DB")
			os.Exit(1)
		}

		//account found!
		value := v.Value
		err := json.Unmarshal(value[:], &account)

		if err != nil {
			tracelog.Errorf(err, "database", "SetAccountStatus", "Error to get the doc from the DB")
			os.Exit(1)
		}

		tracelog.Trace("database","SetAccountStatus","Account found")

		account.Status = updateStatus.Status

		// Update the the document
		err = UpdateDoc(account)

		if err != nil {
			tracelog.Errorf(err, "database", "SetAccountStatus", "Error to update Status into the DB")
			os.Exit(1)
		}

		return  nil
	}

	tracelog.Errorf(err, "database", "SetAccountStatus", "Error to update Status into the DB")
	dberr = errors.New("Error to update Status into the DB")
	return dberr
}

func GetAccountsByStatus(status pb_account.Status) (accounts *pb_account.Accounts, err error){
	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountsByStatus", "Error to connect to the DB")
		os.Exit(1)
	}

	s := status.Status.String()

	value := 0

	switch  {
		case s == "NOTSET":
			value = 0
		case s == "ENABLED":
			value = 1
		case s == "DISABLED":
			value = 2
		case s == "SUSPENDED":
			value = 3
		case s == "REVOKED":
			value = 4
		default:
			value = 0
	}

	output := fmt.Sprintf("{\"selector\":{\"status.status\":{\"$eq\": %v }}}", value)

	queryResults, err := db.QueryDocuments(output)

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccountsByStatus", "Error to search account status into the DB")
	}

	var accountList pb_account.Accounts

	for _, v := range *queryResults {

		account := pb_account.Account{}

		//account found!
		value := v.Value
		err := json.Unmarshal(value[:], &account)

		if err != nil {
			tracelog.Errorf(err, "database", "GetAccountsByStatus", "Error to get the doc from the DB")
			os.Exit(1)
		}

		tracelog.Trace("database","GetAccountsByStatus","account with this status found")

		accountList.Accounts = append(accountList.Accounts, &account)
	}

	tracelog.Trace("database", "GetAccountsByStatus", "Done, return (if any) all the accounts based on a given status")

	return &accountList, nil
}

func GetAccounts() (accounts *pb_account.Accounts, err error){
	db, err := ConnectToDB()

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccounts", "Error to connect to the DB")
		os.Exit(1)
	}

	value := "Account"

	output := fmt.Sprintf("{\"selector\":{\"type\":{\"$eq\":\"%v\"}}}", value)

	queryResults, err := db.QueryDocuments(output)

	if err != nil {
		tracelog.Errorf(err, "database", "GetAccounts", "Error to search accounts into the DB")
	}

	var accountList pb_account.Accounts

	for _, v := range *queryResults {

		account := pb_account.Account{}

		//account found!
		value := v.Value
		err := json.Unmarshal(value[:], &account)

		if err != nil {
			tracelog.Errorf(err, "database", "GetAccounts", "Error to get the doc from the DB")
			os.Exit(1)
		}

		tracelog.Trace("database","GetAccounts","account with this status found")

		accountList.Accounts = append(accountList.Accounts, &account)
	}

	tracelog.Trace("database", "GetAccounts", "Done, return all the Accounts")

	return &accountList, nil
}

