package data

import (
	"fmt"
	"strconv"
)

type Price float64

func (p Price) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%.2f $", p)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
