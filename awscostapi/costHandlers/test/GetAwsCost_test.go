package test

import (
	costhandlres "awscostapi/costHandlers"
	"fmt"
	"testing"
)

func TestGetAwsCosts(t *testing.T) {

	result := costhandlres.GetAwsCost()
	fmt.Println(result[0][0])

	if result[0][0] != "2022-10-01" {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}
}
