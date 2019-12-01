package main

import "log"

func PrintDayHeader(year int, day int) {
	log.Printf("🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄\n")
	log.Printf("🎄 🎄 🎄 AOC Day: %02d (%v)🎄 🎄 🎄 🎄\n", day, year)
	log.Printf("🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄\n")
}

func main() {
	Day01()
}
