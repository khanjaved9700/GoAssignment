package costhandlres

import (
	datehandlers "awscostapi/dateHandlers"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func GetAwsCost() [][]string {
	sess, _ := session.NewSession()
	svc := costexplorer.New(sess)
	now := datehandlers.GetDates()

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: aws.String("DAILY"),
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(*now.Start),
			End:   aws.String(*now.End),
		},
		Metrics: []*string{
			aws.String("BLENDED_COST"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
		},
	}

	req, resp := svc.GetCostAndUsageRequest(input)

	err := req.Send()
	if err != nil {
		fmt.Println(err)
	}
	var resultsCosts [][]string

	for _, results := range resp.ResultsByTime {
		startDate := *results.TimePeriod.Start
		for _, groups := range results.Groups {
			for _, metrics := range groups.Metrics {
				info := []string{startDate, *groups.Keys[0], *metrics.Amount}
				resultsCosts = append(resultsCosts, info)

			}
		}
	}
	return resultsCosts
}
