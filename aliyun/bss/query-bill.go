package bss

import (
	client2 "github.com/alibabacloud-go/bssopenapi-20171214/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
	"seatimc/backend/utils"
)

func getResp(client *client2.Client, productCode string, subscriptionType string, cycle string) ([]aliyun.Bill, *errors.CustomErr) {
	resp, err := client.QueryBill(&client2.QueryBillRequest{
		BillingCycle:     tea.String(cycle),
		IsHideZeroCharge: tea.Bool(true),
		PageNum:          tea.Int32(1),
		PageSize:         tea.Int32(300),
		ProductCode:      tea.String(productCode),
		SubscriptionType: tea.String(subscriptionType),
	})

	if err != nil {
		return nil, errors.AliyunError(err)
	}

	var result []aliyun.Bill

	for _, b := range resp.Body.Data.Items.Item {
		result = append(result, aliyun.Bill{
			Cycle:            cycle,
			CashAmount:       tea.Float32Value(b.CashAmount),
			ProductCode:      tea.StringValue(b.ProductCode),
			SubscriptionType: tea.StringValue(b.SubscriptionType),
		})
	}

	return result, nil
}

func QueryBill(startCycle string, endCycle string, productCode string, subscriptionType string) ([]aliyun.Bill, *errors.CustomErr) {
	client, customErr := aliyun.CreateBssClient()

	if customErr != nil {
		return nil, customErr
	}

	if startCycle == endCycle {
		bills, bErr := getResp(client, productCode, subscriptionType, startCycle)

		if bErr != nil {
			return nil, bErr
		}

		return bills, nil
	}

	startCycleDate, err := utils.ParseTimeRFC3339(startCycle + "-01T00:00:00Z")

	if err != nil {
		return nil, errors.WrongParam()
	}

	endCycleDate, err := utils.ParseTimeRFC3339(endCycle + "-01T00:00:00Z")

	if err != nil {
		return nil, errors.WrongParam()
	}

	if endCycleDate.Before(startCycleDate) {
		return nil, errors.WrongParam()
	}

	var result []aliyun.Bill

	for endCycleDate.After(startCycleDate) || endCycleDate.Equal(startCycleDate) {
		cycle := endCycleDate.Format("2006-01")

		bills, bErr := getResp(client, productCode, subscriptionType, cycle)

		if bErr != nil {
			return nil, bErr
		}

		result = append(result, bills...)

		startCycleDate.AddDate(0, 1, 0)
	}

	return result, nil
}
