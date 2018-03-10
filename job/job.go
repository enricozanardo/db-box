package job

import (
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/onezerobinary/db-box/repository"
	"github.com/goinggo/tracelog"
	"time"
	"fmt"
	"github.com/jasonlvhit/gocron"
)


func DailyCheckNewAccounts(){

	statusNotSet := pb_account.Status{pb_account.Status_NOTSET}

	accounts, err := repository.GetAccountsByStatus(statusNotSet)

	if err != nil {
		tracelog.Errorf(err, "job", "dailyCheckNewAccounts", "It was not possible to retrieve the accounts")
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	// Get the current time
	mytime := time.Now()
	now := mytime.Format(layout)

	checkTime, err := time.Parse(layout, now)

	if err != nil {
		tracelog.Errorf(err, "job", "dailyCheckNewAccounts", "Error to convert the time")
		return
	}

	//Check the end time of those accounts
	for _, account := range accounts.Accounts {
		// Transform the string in time
		accountExpiration, err := time.Parse(layout, account.Expiration)

		if err != nil {
			tracelog.Errorf(err, "job", "dailyCheckNewAccounts", "Error to convert account expiration time")
			return
		}

		message := fmt.Sprintf("Account %v checked, expiration: %v -- checkedtime %v", account.Username, account.Expiration, checkTime)
		tracelog.Trace("job", "dailyCheckNewAccounts", message)

		if accountExpiration.Before(checkTime) {

			message := fmt.Sprintf("Account %v, suspended, expiration: %v -- checkedtime %v", account.Username, account.Expiration, checkTime)
			tracelog.Warning("job", "dailyCheckNewAccounts", message)

			account.Status.Status = pb_account.Status_DISABLED

			err := repository.UpdateDoc(*account)

			if err != nil {
				tracelog.Error(err,"job", "dailyCheckNewAccounts")
			}
		}
	}
}

// Every day at 23:59 check the status of the account
func CheckAccountStatus() {
	gocron.Start()
	gocron.Every(1).Day().At("23:59").Do(DailyCheckNewAccounts)
}