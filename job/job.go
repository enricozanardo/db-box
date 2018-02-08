package job

import "github.com/claudiu/gocron"

// Every day at 23:59 check the status of the account
func CheckAccountStatus() {
	gocron.Start()
	gocron.Every(5).Seconds().Do(dailyCheckNewAccounts)
}

// If an account is in status suspened then set its status disabled
func dailyCheckNewAccounts(){
	println("Do something")
}
