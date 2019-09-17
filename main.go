package main

import (
	"fmt"
	"github.com/kullcrom/squanchy/attomdata/basicprofile"
	"github.com/kullcrom/squanchy/attomdata/building"
)

func main() {
	construction := building.NewConstruction("good")
	fmt.Println(construction.Condition)
	summary := new(building.Summary)
	fmt.Println(summary.Levels)
	building := new(building.Building)
	fmt.Println(building.Construction)
	profile := new(basicprofile.BasicProfile)
	fmt.Println(profile.Property)
}
