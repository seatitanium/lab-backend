package bss

import (
	"fmt"
	openapi "github.com/alibabacloud-go/bssopenapi-20171214/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
)

func QueryBill(billingCycle string, pagenum int32, pagesize int32) ([]aliyun.Bill, *errHandler.CustomErr) {
	client, customErr := aliyun.CreateBssClient()

	var bills []aliyun.Bill

	if customErr != nil {
		return bills, customErr
	}

	res, err := client.QueryBill(&openapi.QueryBillRequest{
		BillingCycle: tea.String(billingCycle),
		PageNum:      tea.Int32(pagenum),
		PageSize:     tea.Int32(pagesize),
	})

	if err != nil {
		return bills, errHandler.AliyunError(err)
	}

	for _, item := range res.Body.Data.Items.Item {
		bills = append(bills, aliyun.Bill{
			ProductCode:   tea.StringValue(item.ProductCode),
			ProductName:   tea.StringValue(item.ProductName),
			PaymentAmount: tea.Float32Value(item.PaymentAmount),
			PretaxAmount:  tea.Float32Value(item.PretaxAmount),
		})

		fmt.Printf("%v\n", item)
	}

	return bills, nil
}
