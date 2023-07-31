package aws

import (
	"errors"
	"fmt"
	"time"
)

func WaitForState(retry, timeout time.Duration, getStatus func() (bool, error)) error {
	for t := time.Second * 0; t <= timeout; t += retry {
		res, err := getStatus()
		if err != nil {
			return err
		}
		if res {
			return nil
		}
		fmt.Println("retrying")
		time.Sleep(retry)
	}
	return errors.New(fmt.Sprintf("failed to get desired status"))
}
