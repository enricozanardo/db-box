package job

import (
	"testing"
	"time"
)

func TestCheckAccountStatus(t *testing.T) {

	go CheckAccountStatus()

	time.Sleep(10 * time.Second)
	//t.Errorf("Check fail")

}