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
		//viper.AddConfigPath("../.")
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
	tracelog.Warning("database", "GetAccountByCredentials", "Account not found, return nil")
	return nil, nil
}

func GetAccountByToken(token pb_account.Token) (account pb_account.Account, err error) {
	return
}

func RemoveDoc(token pb_account.Token) (err error){
	return
}

func UpdateDoc(account pb_account.Account) (err error){
	return
}

func CheckEmail(email pb_account.Email) (token pb_account.Token, err error){
	return
}

func GetAccountStatus(token pb_account.Token) (accountStatus pb_account.Status, err error){
	return
}

func SetAccountStatus(updateStatus pb_account.UpdateStatus)(err error){
	return
}

func GetAccountsByStatus(status pb_account.Status) (accounts pb_account.Accounts, err error){
	return
}

func GetAccounts() (accounts pb_account.Accounts, err error){
	return
}

