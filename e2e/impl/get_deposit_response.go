package impl

import (
	"encoding/json"
	"errors"
	"math/big"
	"regexp"
	"time"
)

type DepositResponse struct {
	Deposits []Deposit
}

type Deposit struct {
	AvailableAt *ChainTime
	Value       *big.Int
}

type ChainTime time.Time

func (g *ChainTime) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	t, err := parseDate(string(data))
	if err != nil {
		return err
	}

	*g = ChainTime(t)

	return nil
}

func (g *ChainTime) Time() time.Time {
	return time.Time(*g)
}

func fixUnquotedJSON(data string) []byte {
	re := regexp.MustCompile("(\\w+)\\:(.*)")
	data = re.ReplaceAllString(data, "\"$1\":$2")
	return []byte(data)
}

func unmarshalDepositResponse(value string) (*DepositResponse, error) {
	resp := &DepositResponse{}
	valueBytes := fixUnquotedJSON(value)
	err := json.Unmarshal(valueBytes, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func parseDepositResponse(value string) (Deposit, error) {
	resp, err := unmarshalDepositResponse(value)
	if err != nil {
		return Deposit{}, err
	}

	if len(resp.Deposits) == 0 {
		return Deposit{}, errors.New("empty response")
	}

	if len(resp.Deposits) > 1 {
		return Deposit{}, errors.New("ge multiple deposits in the response")
	}

	deposit := resp.Deposits[0]

	if deposit.AvailableAt == nil {
		return Deposit{}, errors.New("empty AvailableAt field")
	}

	if deposit.Value == nil {
		return Deposit{}, errors.New("empty Value field")
	}

	return deposit, nil
}

func parseDate(date string) (time.Time, error) {
	const longForm = "\"2006-01-02 15:04:05 -0700 MST\""

	return time.Parse(longForm, date)
}
