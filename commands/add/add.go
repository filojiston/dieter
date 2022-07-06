package add

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const DATE_FORMAT string = "01-02-2006"
const FILE_PATH string = ".data"

type Food struct {
	name     string
	quantity int
	calories int
	date     string
}

func FoodToString(food *Food) string {
	return fmt.Sprintf("food: %s, quantity: %s, calories: %s, date: %s\n", food.name, strconv.Itoa(food.quantity), strconv.Itoa(food.calories), food.date)
}

func Handle(flagSet *flag.FlagSet, args []string) {
	food := new(Food)
	flagSet.StringVar(&food.name, "food", "", "name of the food")
	flagSet.IntVar(&food.quantity, "qty", 0, "quantity of the food")
	flagSet.IntVar(&food.calories, "cal", 0, "how much calories is it?")
	flagSet.StringVar(&food.date, "date", time.Now().Format(DATE_FORMAT), "when did you eat?")

	flagSet.Parse(args)

	if len(food.name) == 0 || food.quantity == 0 || food.calories == 0 {
		fmt.Println("you must use -food, -qty and -cal for adding a food.")
		return
	}

	f, _ := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_APPEND, 0755)
	f.WriteString(FoodToString(food))
}
