package chunk

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getData(numberOfAccountData int, numberOfDebitSourceAccount int) (accountList []autoCreditHistory) {
	start := startCountProcessTime()
	for i := 0; i < numberOfAccountData; i++ {
		newAccount := autoCreditHistory{
			AccountNumber:      strconv.Itoa(i + 1),
			AccountOpeningDate: time.Now().Local().Add(time.Second * time.Duration(i)).Format("2006-01-02 15:04:05"),
			AutoCreditDate:     time.Now().Local().Add(time.Second * time.Duration(i)),
			TargetAmount:       0,
			TransactionAmount:  0,
			TransactionStatus:  "S",
			AvailableBalance:   1000,
			EndingBalance:      0,
			NextAutoCreditDate: time.Now().Local().Add(time.Second * time.Duration(i)),
			AutoCreditOption:   "daily",
		}
		newAccount.DebitSourceAccount = strconv.Itoa(randInt(1, numberOfDebitSourceAccount))
		accountList = append(accountList, newAccount)
	}
	timeElapsed := endCountProcessTime(start)
	timeElapsedSeconds := timeElapsed.Seconds()
	timeElapsedNanoSeconds := timeElapsed.Nanoseconds()
	fmt.Printf("Total time elapsed for getting random data => %f seconds => %d nanoseconds\n", timeElapsedSeconds, timeElapsedNanoSeconds)
	return
}

func randInt(min int, max int) (randomNumber int) {
	rand.Seed(time.Now().UnixNano())
	randomNumber = min + rand.Intn(max-min)
	return
}

func startCountProcessTime() time.Time {
	return time.Now()
}

func endCountProcessTime(start time.Time) time.Duration {
	return time.Since(start)
}

func printProcessTime(timeElapsed time.Duration, numberOfAccount int) {
	timeElapsedSeconds := timeElapsed.Seconds()
	timeElapsedNanoSeconds := timeElapsed.Nanoseconds()
	fmt.Printf("Total time elapsed for %d grouped data => %f seconds => %d nanoseconds\n", numberOfAccount, timeElapsedSeconds, timeElapsedNanoSeconds)
}
