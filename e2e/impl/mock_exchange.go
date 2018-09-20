package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kcoin/e2e/cluster"
	"github.com/kowala-tech/kcoin/mock-exchange/app"
)

type MockExchangeContext struct {
	globalCtx *Context
	mockAddr  string
}

func NewMockExchangeContext(parentCtx *Context) *MockExchangeContext {
	ctx := &MockExchangeContext{
		globalCtx: parentCtx,
	}
	return ctx
}

func (ctx *MockExchangeContext) Reset() {
}

func (ctx *MockExchangeContext) TheMockExchangeIsRunning() error {
	mockAddr, err := ctx.runMockExchange()
	if err != nil {
		return fmt.Errorf("error executing mock exchange: %s", err)
	}

	ctx.mockAddr = mockAddr

	return nil
}

func (ctx *MockExchangeContext) runMockExchange() (mockAddr string, err error) {
	mockSpec, err := cluster.MockExchangeSpec(ctx.globalCtx.nodeSuffix)
	if err != nil {
		return "", err
	}
	if err := ctx.globalCtx.nodeRunner.Run(mockSpec); err != nil {
		return "", err
	}
	mockIP, err := ctx.globalCtx.nodeRunner.IP(mockSpec.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v:8080", mockIP), nil
}

func (ctx *MockExchangeContext) IFetchTheExchangeWithMockData(accountsDataTable *gherkin.DataTable) error {
	request, err := getRateRequestFromTableData(accountsDataTable)
	if err != nil {
		return err
	}

	postData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to encode table data: %s", err)
	}

	response, err := http.Post(
		fmt.Sprintf("http://%s:9080/api/fetch", ctx.globalCtx.nodeRunner.HostIP()),
		"",
		bytes.NewReader(postData),
	)
	if err != nil {
		return fmt.Errorf("failed to send post request to exchange: %s", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("fetching the mocked exchange server returned status code %d", response.StatusCode)
	}

	return nil
}

func getRateRequestFromTableData(accountsDataTable *gherkin.DataTable) (*app.Request, error) {
	request := app.Request{}
	for i, row := range accountsDataTable.Rows {
		if i != 0 {
			typeVal := row.Cells[0].Value
			if typeVal == "buy" {
				amountVal, err := strconv.ParseFloat(row.Cells[1].Value, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid amount data in table for exchange mock server: %s", err)
				}

				rateVal, err := strconv.ParseFloat(row.Cells[2].Value, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid rate data in table for exchange mock server: %s", err)
				}

				request.Buy = append(request.Buy, app.RateValue{Amount: amountVal, Rate: rateVal})
			} else if typeVal == "sell" {
				amountVal, err := strconv.ParseFloat(row.Cells[1].Value, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid amount data in table for exchange mock server: %s", err)
				}

				rateVal, err := strconv.ParseFloat(row.Cells[2].Value, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid rate data in table for exchange mock server: %s", err)
				}

				request.Sell = append(request.Sell, app.RateValue{Amount: amountVal, Rate: rateVal})
			} else {
				continue
			}
		}
	}

	return &request, nil
}
