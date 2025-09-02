package training

import (
	"log"
	"time"
)

func main() {
	p1 := person{12, "ali", "mohamadi", "0988847",
	 time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC)}
	
	log.Println(p1)
}

type person struct {
	Id          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Birth       time.Time
}