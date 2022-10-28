package datehandlres

import (
	"time"

	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func SetHeading() []string {
	headings := []string{"DATE", "AWS SERVICES", "COST"}

	return headings
}

// / getDates returns a DateInterval for the last week
func GetDates() *costexplorer.DateInterval {
	now := time.Now()
	then := now.AddDate(0, 0, -30)
	dateRange := costexplorer.DateInterval{}
	dateRange.SetEnd(now.Format("2006-01-02"))
	dateRange.SetStart(then.Format("2006-01-02"))

	return &dateRange
}
