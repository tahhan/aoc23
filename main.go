package main

import (
	day01_1 "aoc23/day01-1"
	day01_2 "aoc23/day01-2"
	day02_1 "aoc23/day02-1"
	day02_2 "aoc23/day02-2"
	day03_1 "aoc23/day03-1"
	day03_2 "aoc23/day03-2"
	day04_1 "aoc23/day04-1"
	day04_2 "aoc23/day04-2"
	day05_1 "aoc23/day05-1"
	day06_1 "aoc23/day06-1"
	day06_2 "aoc23/day06-2"
	day07_1 "aoc23/day07-1"
	day07_2 "aoc23/day07-2"
	day08_1 "aoc23/day08-1"
	day08_2 "aoc23/day08-2"
	day09_1 "aoc23/day09-1"
	"fmt"
)

func main() {
	fmt.Printf("day 1-1: %d\n", day01_1.Trebuchet())
	fmt.Printf("day 1-2: %d\n", day01_2.Trebuchet2())
	fmt.Printf("day 2-1: %d\n", day02_1.CubeConundrum())
	fmt.Printf("day 2-2: %d\n", day02_2.CubeConundrum2())
	fmt.Printf("day 3-1: %d\n", day03_1.GearRatios())
	fmt.Printf("day 3-2: %d\n", day03_2.GearRatios2())
	fmt.Printf("day 4-1: %d\n", day04_1.Scratchcards())
	fmt.Printf("day 4-2: %d\n", day04_2.Scratchcards2())
	fmt.Printf("day 5-1: %d\n", day05_1.Seeds())
	//fmt.Printf("day 5-2: %d\n", day05_2.Seeds2()) // takes too long :(
	fmt.Printf("day 6-1: %d\n", day06_1.Boats())
	fmt.Printf("day 6-2: %d\n", day06_2.Boats2())
	fmt.Printf("day 7-1: %d\n", day07_1.CamelCards())
	fmt.Printf("day 7-2: %d\n", day07_2.CamelCards2())
	fmt.Printf("day 8-1: %d\n", day08_1.HauntedWasteland())
	fmt.Printf("day 8-2: %d\n", day08_2.HauntedWasteland2())
	fmt.Printf("day 9-1: %d\n", day09_1.MirageMaintenance())
}
