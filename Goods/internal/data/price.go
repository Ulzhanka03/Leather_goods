package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidPriceFormat = errors.New("invalid price format")

type Price float64

func (p Price) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%.2f $", p)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
func (p *Price) UnmarshalJSON(jsonValue []byte) error {

	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidPriceFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "$" {
		return ErrInvalidPriceFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidPriceFormat
	}

	*p = Price(i)
	return nil
}
