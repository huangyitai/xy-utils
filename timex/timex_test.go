package timex

import (
	"fmt"
	"testing"
)

func TestStartOfWeek(t *testing.T) {
	fmt.Println(StartOfYesterday(), EndOfYesterday())
	fmt.Println(StartOfToday(), EndOfToday())
	fmt.Println(StartOfTomorrow(), EndOfTomorrow())

	fmt.Println(StartOfYesterday().Unix(), EndOfYesterday().Unix())
	fmt.Println(StartOfToday().Unix(), EndOfToday().Unix())
	fmt.Println(StartOfTomorrow().Unix(), EndOfTomorrow().Unix())

	fmt.Println(StartOfLastWeekNow(), EndOfLastWeekNow())
	fmt.Println(StartOfWeekNow(), EndOfWeekNow())
	fmt.Println(StartOfNextWeekNow(), EndOfNextWeekNow())

	fmt.Println(StartOfLastWeekNow().Unix(), EndOfLastWeekNow().Unix())
	fmt.Println(StartOfWeekNow().Unix(), EndOfWeekNow().Unix())
	fmt.Println(StartOfNextWeekNow().Unix(), EndOfNextWeekNow().Unix())

	fmt.Println(StartOfLastMonthNow(), EndOfLastMonthNow())
	fmt.Println(StartOfMonthNow(), EndOfMonthNow())
	fmt.Println(StartOfNextMonthNow(), EndOfNextMonthNow())

	fmt.Println(StartOfLastYearNow(), EndOfLastYearNow())
	fmt.Println(StartOfYearNow(), EndOfYearNow())
	fmt.Println(StartOfNextYearNow(), EndOfNextYearNow())

}
