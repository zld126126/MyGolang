package util

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func PrintJson(any ...interface{}) {
	for _, obj := range any {
		b, err := json.Marshal(obj)
		fmt.Println(err, string(b))
	}
}

func ToJsonString(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(bs), nil
}
