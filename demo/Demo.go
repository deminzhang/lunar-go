package main

import (
	"fmt"
	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/6tail/lunar-go/calendar"
	"time"
)

func main() {
	// 法定假日
	holiday := HolidayUtil.GetHoliday("2020-01-01")
	fmt.Println(holiday)

	holiday = HolidayUtil.GetHolidayByYmd(2020, 5, 1)
	fmt.Println(holiday)

	// 阳历
	solar := calendar.NewSolarFromDate(time.Now())
	fmt.Println(solar.ToFullString())

	// 儒略日
	solar = calendar.NewSolarFromJulianDay(2459045.5)
	fmt.Println(solar.ToYmdHms())

	solar = calendar.NewSolarFromYmd(2020, 7, 15)
	fmt.Println(solar.GetJulianDay())

	// 节气表
	solar = calendar.NewSolarFromYmd(2022, 7, 15)
	lunar := solar.GetLunar()
	jieQi := lunar.GetJieQiTable()
	for i := lunar.GetJieQiList().Front(); i != nil; i = i.Next() {
		name := i.Value.(string)
		fmt.Println(name, jieQi[name].ToYmdHms())
	}

	// 阳历往后推一天
	solar = solar.Next(1)
	fmt.Println(solar.ToFullString())

	// 法定假日
	holidays := HolidayUtil.GetHolidaysByTargetYmd(2019, 1, 1)
	for i := holidays.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(*calendar.Holiday))
	}

	// 阴历
	lunar = calendar.NewLunarFromYmd(1986, 4, 21)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetSolar().ToFullString())

	// 八字
	for _, v := range lunar.GetBaZi() {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()

	// 八字五行
	for _, v := range lunar.GetBaZiWuXing() {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()

	// 八字天干十神
	for _, v := range lunar.GetBaZiShiShenGan() {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()

	// 八字地支十神
	for _, v := range lunar.GetBaZiShiShenZhi() {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()

	// 日宜
	lunar = calendar.NewLunarFromDate(time.Now())
	for i := lunar.GetDayYi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 日忌
	for i := lunar.GetDayJi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 日吉神宜趋
	for i := lunar.GetDayJiShen().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 日凶煞宜忌
	for i := lunar.GetDayXiongSha().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 时辰宜
	for i := lunar.GetTimeYi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 时辰忌
	for i := lunar.GetTimeJi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 月相
	fmt.Println(lunar.GetYueXiang())

	// 值年九星
	jx := lunar.GetYearNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值月九星
	jx = lunar.GetMonthNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值日九星
	jx = lunar.GetDayNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值时九星
	jx = lunar.GetTimeNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 八字转阳历
	for i := calendar.ListSolarFromBaZi("庚子", "戊子", "己卯", "庚午").Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(*calendar.Solar).ToFullString())
	}
	fmt.Println()
}