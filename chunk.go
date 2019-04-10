package chunk

import (
	"time"
)

const (
	numberOfAccount            = 100
	numberOfDebitSourceAccount = 10
	maxJob                     = 10
)

var (
	chunkedDataTempPerDebitSourceAccount       = [][]autoCreditHistory{}
	chunkedDataTempPerDebitSourceAccountLength = 0
	chunkedDataTemp                            = [][][]autoCreditHistory{}
	chunkedDataWrapper                         = [][][]autoCreditHistory{}
	loopCount                                  int
	lastLoop                                   bool
)

type (
	accountGroup struct {
		DebitSourceAccount string              `json:"debit_source_account"`
		AutoCreditDate     time.Time           `json:"auto_credit_date"`
		Account            []autoCreditHistory `json:"account"`
	}
	autoCreditHistory struct {
		AccountNumber      string    `json:"account_number" orm:"column(account_number);pk"`
		AutoCreditDate     time.Time `json:"auto_credit_date" orm:"column(auto_credit_date)"`
		DebitSourceAccount string    `json:"debit_source_account" orm:"column(debit_source_account)"`
		TargetAmount       float64   `json:"target_amount" orm:"column(target_amount)"`
		TransactionAmount  float64   `json:"transaction_amount" orm:"column(transaction_amount)"`
		TransactionStatus  string    `json:"transaction_status" orm:"column(transaction_status)"`
		AvailableBalance   float64   `json:"available_balance" orm:"column(available_balance)"`
		EndingBalance      float64   `json:"ending_balance" orm:"column(ending_balance)"`
		NextAutoCreditDate time.Time `json:"next_auto_credit_date" orm:"column(next_auto_credit_date)"`
		AutoCreditOption   string    `json:"auto_credit_option" orm:"column(auto_credit_option)"`
		AccountOpeningDate string    `json:"account_opening_date" orm:"column(account_opening_date)"`
	}
)

// chunk ...
func chunk(ungroupedAccountList []autoCreditHistory) (chunkedDataWrapper [][][]autoCreditHistory) {
	accountListGrouped := make(map[string][]autoCreditHistory)
	for _, val := range ungroupedAccountList {
		accountListGrouped[val.DebitSourceAccount] = append(accountListGrouped[val.DebitSourceAccount], val)
	}

	for _, val := range accountListGrouped {
		if loopCount == len(accountListGrouped)-1 {
			lastLoop = true
		}
		// newAccountGroup := accountGroup{
		// 	DebitSourceAccount: key,
		// 	AutoCreditDate:     time.Now(),
		// 	Account:            val,
		// }
		// dataToStore = append(dataToStore, newAccountGroup) //data to store to db later as history

		// fmt.Println("(len(val) + chunkedDataTempPerDebitSourceAccountLength)")
		// fmt.Println((len(val) + chunkedDataTempPerDebitSourceAccountLength))
		if ((len(val) + chunkedDataTempPerDebitSourceAccountLength) > maxJob) || lastLoop {
			if lastLoop {
				if len(val)+len(chunkedDataTempPerDebitSourceAccount) < maxJob {
					chunkedDataTempPerDebitSourceAccount = append(chunkedDataTempPerDebitSourceAccount, val)
				} else {
					chunkedDataTemp = append(chunkedDataTemp, [][]autoCreditHistory{val})
				}
			}
			if chunkedDataTempPerDebitSourceAccountLength == maxJob {
				chunkedDataWrapper = append(chunkedDataWrapper, chunkedDataTempPerDebitSourceAccount)
				// fmt.Println("chunkedDataWrapperrrrrrrrrrrrrrr")
				// fmt.Println(chunkedDataWrapper)
			} else {
				// fmt.Println("asdasdasd")
				if len(chunkedDataTemp) == 0 || (len(val)+chunkedDataTempPerDebitSourceAccountLength) != maxJob {
					if len(chunkedDataTempPerDebitSourceAccount) != 0 {
						chunkedDataTemp = append(chunkedDataTemp, chunkedDataTempPerDebitSourceAccount)
					}
				} else {
					for key, val := range chunkedDataTemp {
						if len(val) == maxJob {
							chunkedDataWrapper = append(chunkedDataWrapper, val)
							continue
						}
						if (len(val) + chunkedDataTempPerDebitSourceAccountLength) == maxJob {
							chunkedDataTemp[key] = append(chunkedDataTemp[key], chunkedDataTempPerDebitSourceAccount...)
							// fmt.Println("chunkedDataTemp[key]")
							// fmt.Println(chunkedDataTemp[key])
							var accountGroupPerDebitSourceTemp [][]autoCreditHistory
							accountGroupPerDebitSourceTemp = chunkedDataTemp[key]
							chunkedDataWrapper = append(chunkedDataWrapper, accountGroupPerDebitSourceTemp)
							// chunkedDataWrapper = append(chunkedDataWrapper, chunkedDataTemp[key])
							chunkedDataTemp[key] = nil
							break
						}
					}
				}
			}
			chunkedDataTempPerDebitSourceAccount = nil
			chunkedDataTempPerDebitSourceAccountLength = 0
		}
		if lastLoop {
			// fmt.Printf("length : %d\n", len(val))
			// fmt.Println("chunkedDataTemp")
			// chunkedDataTempWrapperByte, _ := json.Marshal(chunkedDataTemp)
			// fmt.Println(string(chunkedDataTempWrapperByte))
			for _, val := range chunkedDataTemp {
				chunkedDataWrapper = append(chunkedDataWrapper, val)
			}
			break
		}

		chunkedDataTempPerDebitSourceAccount = append(chunkedDataTempPerDebitSourceAccount, val) //chunked data to process auto credit
		chunkedDataTempPerDebitSourceAccountLength += len(val)
		loopCount++
		// fmt.Printf("length : %d\n", len(val))
		// fmt.Println("chunkedDataTempPerDebitSourceAccount")
		// fmt.Println(chunkedDataTempPerDebitSourceAccount)
	}
	return
}
