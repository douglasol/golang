package main

import (
	"fmt"
	"time"
)

func main() {

	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// tm := time.Date(2020, time.April, 34, 25, 72, 01, 0, time.UTC)
	//2/000, 2, 1, 12, 30, 0, 0, time.UTC
	//agora := time.Now()
	//ano, mes, dia := agora.Date()
	//fmt.Printf("ano = %v\n", ano)
	//fmt.Printf("mes = %v\n", mes)
	//fmt.Printf("dia = %v\n", dia)

	//currentTime := time.Now()
	//fmt.Println("The time is", currentTime)
	// 2022-05-14 16:45:41.0949515 -0300 -03 m=+0.002155801

	//fmt.Println(currentTime.Date())
	//2022 May 14

	//fmt.Println(currentTime.Zone())
	// -03 -10800

	//func (Time) Location ¶
	//fmt.Printf("location = %v\n", agora.Location())

	agora := time.Now()
	fmt.Println("Location : ", agora.Location(), " Time : ", agora) // local time

	/*location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("NY: ", location, " Time : ", agora.In(location))
	*/
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SP: ", location, " Time : ", agora.In(location))
	/*
		location, err = time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("TOKIO: ", location, " Time : ", agora.In(location))
	*/

	//t,_ := time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
	//fmt.Println(agora.YearDay()) // 315
	//fmt.Println(agora.Weekday()) // Wednesday

	t, _ := time.Parse("2022 01 02 15 04")
	a, m, d := t.Date()
	fmt.Printf("Data %v %v %v", a, m, d)

	//t,_ = time.Parse("2006 01 02 15 04", "2011 01 01 0 00")
	//fmt.Println(t.YearDay())
	//fmt.Println(t.Weekday())

	/*



	   loc, _ := time.LoadLocation("Asia/Shanghai")
	   now := time.Now().In(loc)
	   fmt.Println("Location : ", loc, " Time : ", now) // Asia/Shanghai

	*/

	//fmt.Println(currentTime.Clock())
	// 16 48 3 (h,m,s)

	//fmt.Println("The year is", currentTime.Year())
	//fmt.Println("The month is", currentTime.Month())
	//fmt.Println("The day is", currentTime.Day())
	//fmt.Println("The hour is", currentTime.Hour())
	//fmt.Println("The minute is", currentTime.Hour())
	//fmt.Println("The second is", currentTime.Second())
	//The year is 2022
	//The month is May
	//The day is 14
	//The hour is 16
	//The minute is 16
	//The second is 17

	//fmt.Printf("%d-%d-%d %d:%d:%d\n",
	//	currentTime.Year(),
	//	currentTime.Month(),
	//	currentTime.Day(),
	//	currentTime.Hour(),
	//	currentTime.Hour(),
	//	currentTime.Second())

	// 2022-5-14 16:16:21

	//theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	//fmt.Println("The time is", theTime)
	//The time is 2021-08-15 14:30:45.0000001 -0300 -03

	//theTime := time.Local
	//fmt.Println("Local:", theTime)
	//Local: Local

	//
	//https://www.digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go
	//https://stackoverflow.com/questions/49945445/inserting-datetime-value-into-ms-sql-using-golang

}
