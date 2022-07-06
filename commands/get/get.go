package get

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT string = "01-02-2006"
const FILE_PATH string = ".data"

type GetType string

const (
	Today GetType = "today"
	Week          = "week"
)

func ReadFoodsFromFile() []string {
	var buffer bytes.Buffer

	f, _ := os.Open(FILE_PATH)
	buffer.ReadFrom(f)
	result := strings.Split(buffer.String(), "\n")
	return result[:len(result)-1]
}

func FilterFoods(foods []string, getType GetType) ([]string, int) {
	var filteredFoods []string
	var totalCalories int

	for _, food := range foods {
		foodSplitted := strings.Split(food, ",")
		date, _ := time.Parse(DATE_FORMAT, foodSplitted[3][7:])
		diffAsDays := int(time.Now().Sub(date).Hours() / 24)
		if (diffAsDays == 0 && getType == Today) || (diffAsDays <= 7 && getType == Week) {
			filteredFoods = append(filteredFoods, food)
			calories, _ := strconv.Atoi(foodSplitted[2][11:])
			totalCalories += calories
		}
	}
	return filteredFoods, totalCalories
}

func Handle(flagSet *flag.FlagSet, args []string) {
	getType := flagSet.String("type", "", "today or week?")
	flagSet.Parse(args)

	if len(*getType) == 0 {
		fmt.Println("you must specify the type.")
		return
	}

	var filteredFoods []string
	var totalCalories int
	allFoods := ReadFoodsFromFile()
	switch GetType(*getType) {
	case Today:
		filteredFoods, totalCalories = FilterFoods(allFoods, Today)
		break
	case Week:
		filteredFoods, totalCalories = FilterFoods(allFoods, Week)
	}
	for _, food := range filteredFoods {
		fmt.Println(food)
	}
	fmt.Println("total calories:", totalCalories)
}
