package main

import "log"

func PrintCurrentDayHeader() {
	PrintDayHeader(GetAOCDay(), GetAOCYear())
}

func PrintDayHeader(day int, year int) {
	log.Printf("🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄\n")
	log.Printf("🎄 🎄 🎄 AOC Day: %02d (%v)🎄 🎄 🎄 🎄\n", day, year)
	log.Printf("🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄 🎄\n")
}

func main() {
	Day01()
}
