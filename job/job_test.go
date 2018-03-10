package job

import (
	"testing"
	"github.com/goinggo/tracelog"
	"github.com/spf13/viper"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/onezerobinary/db-box/repository"
	"time"
)

func TestCheckAccountStatus(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	viper.SetConfigName("config")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "database_test", "StartConfig", "Error reading config file")
	}

	// Create an account
	username := "fake@fake.com"
	password := "fake"
	token := repository.GenerateToken(username, password)
	faketoken := pb_account.Token{token}

	fakeStatus := pb_account.Status{pb_account.Status_NOTSET}

	created := time.Now()
	expiration := created.Add(time.Duration(24*time.Hour))

	// Set the layout that are needed into the DB
	layout := "2006-01-02T15:04:05.000Z"
	c := string(created.Format(layout))
	e := string(expiration.Format(layout))

	fakeAccount := pb_account.Account{
		token,
		username,
		password,
		&faketoken,
		&fakeStatus,
		"Account",
		c,
		e,
	}

	backToken, err :=  repository.AddDoc(fakeAccount)

	if err != nil {
		t.Error("non possible to store the document into the DB")
	}

	if token != backToken {
		t.Error("token do not match")
	}

	DailyCheckNewAccounts()

	if fakeAccount.Status.Status != pb_account.Status_NOTSET {
		t.Error("Errors in checking the accounts during the cronjob.")
	}

	// Remove the created test account
	err = repository.RemoveDoc(faketoken)

	if err != nil {
		t.Error("Account not removed")
	}
}