## dieter
simple go cli application that you can use to track how many calories you're getting. it uses .data file in the project root for data storage.

subcommands:
* add
    * -food -> string, name of the food
    * -qty -> int, quantity of the food
    * -cal -> int, calorie value of the food
    * -date -> (optional) string, (mm-dd-yyyy format) 
* get
    * -type -> today for todays stats or week for weekly stats, with all food data and total calories.