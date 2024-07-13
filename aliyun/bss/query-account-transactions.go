package bss

import (
	openapi "github.com/alibabacloud-go/bssopenapi-20171214/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
	"seatimc/backend/utils"
	"strconv"
	"time"
)

func QueryAccountTransactions(pagenum int32, pagesize int32) ([]aliyun.Transaction, *errors.CustomErr) {
	client, customErr := aliyun.CreateBssClient()

	var bills []aliyun.Transaction

	if customErr != nil {
		return bills, customErr
	}

	now := time.Now().Format("2006-01-02T15:04:05Z07:00")
	startTime := "2024-01-01T00:00:00Z8:00"

	res, err := client.QueryAccountTransactions(&openapi.QueryAccountTransactionsRequest{
		PageNum:         tea.Int32(pagenum),
		PageSize:        tea.Int32(pagesize),
		CreateTimeEnd:   tea.String(now),
		CreateTimeStart: tea.String(startTime),
	})

	if err != nil {
		return bills, errors.AliyunError(err)
	}

	for _, item := range res.Body.Data.AccountTransactionsList.AccountTransactionsList {
		rawTime := tea.StringValue(item.TransactionTime)

		parsedTime, err := utils.ParseTimeATBZ(rawTime)

		if err != nil {
			return nil, errors.ServerError(err)
		}

		rawAmount := tea.StringValue(item.Amount)

		parsedAmount, err := strconv.ParseFloat(rawAmount, 64)

		if err != nil {
			return nil, errors.ServerError(err)
		}

		bills = append(bills, aliyun.Transaction{
			BillingCycle:    tea.StringValue(item.BillingCycle),
			TransactionTime: parsedTime,
			Amount:          parsedAmount,
			IsIncome:        tea.StringValue(item.TransactionFlow) == "Income",
		})
	}

	return bills, nil
}
