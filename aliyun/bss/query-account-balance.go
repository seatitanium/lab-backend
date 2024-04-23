package bss

import (
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
	"strconv"
)

func QueryAccountBalance() (*aliyun.AvailableBalance, *errHandler.CustomErr) {
	client, customErr := aliyun.CreateBssClient()

	if customErr != nil {
		return nil, customErr
	}

	resp, err := client.QueryAccountBalance()

	if err != nil {
		return nil, errHandler.AliyunError(err)
	}

	var result *aliyun.AvailableBalance

	result.AvailableAmount, err = strconv.ParseFloat(tea.StringValue(resp.Body.Data.AvailableAmount), 32)

	if err != nil {
		return nil, errHandler.ServerError(err)
	}

	result.AvailableCashAmount, err = strconv.ParseFloat(tea.StringValue(resp.Body.Data.AvailableCashAmount), 32)

	if err != nil {
		return nil, errHandler.ServerError(err)
	}

	return result, nil
}
