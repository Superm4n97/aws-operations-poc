package resource

import (
	"errors"
	"fmt"
	"time"
)

func WaitForState(retry, timeout time.Duration, getStatus func() (bool, error)) error {
	for t := time.Second * 0; t <= timeout; t += retry {
		fmt.Println("getting state")
		res, err := getStatus()
		if err != nil {
			return err
		}
		if res == true {
			return nil
		}
		fmt.Println("retrying")
		time.Sleep(retry)
	}
	return errors.New(fmt.Sprintf("failed to get desired status"))
}
