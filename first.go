package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic Income) {
	income := ic
	fmt.Println(income.source(), income.calculate())
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func main() {
	// project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	// project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	// project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	project4 := Advertisement{adName: "Project 4", CPC: 10, noOfClicks: 15}
	// incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(project4)
}
