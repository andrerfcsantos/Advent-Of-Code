package day16_2020

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"sort"
	"strconv"
	"strings"
)

type BySetSize Columns

func (a BySetSize) Len() int           { return len(a) }
func (a BySetSize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySetSize) Less(i, j int) bool { return len(a[i].ValidFields.Set) < len(a[j].ValidFields.Set) }

type Columns []Column

type Column struct {
	Index       int
	ValidFields utils.StringSet
}

type Range struct {
	MinVal int
	MaxVal int
}

type Rule struct {
	Name       string
	LowerRange Range
	UpperRange Range
}

type RuleList []Rule

func (rl RuleList) ValidForAnyField(val int) bool {

	for _, rule := range rl {
		if (val >= rule.LowerRange.MinVal && val <= rule.LowerRange.MaxVal) ||
			(val >= rule.UpperRange.MinVal && val <= rule.UpperRange.MaxVal) {
			return true
		}
	}
	return false
}

func (rl RuleList) ValidFieldsForVal(val int) utils.StringSet {
	res := utils.NewStringSet()

	for _, rule := range rl {
		if (val >= rule.LowerRange.MinVal && val <= rule.LowerRange.MaxVal) ||
			(val >= rule.UpperRange.MinVal && val <= rule.UpperRange.MaxVal) {
			res.Add(rule.Name)
		}
	}

	return res
}

type Ticket []int

type Solver struct {
	Rules         RuleList
	MyTicket      Ticket
	NearbyTickets []Ticket
	ValidTickets  []Ticket
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines, err := utils.GroupByEmptyLines(input)
	if err != nil {
		return err
	}

	// Parse ranges
	for _, line := range lines[0] {
		classSep := strings.Split(line, ": ")
		ranges := strings.Split(classSep[1], " or ")
		lowRange := strings.Split(ranges[0], "-")
		upperRange := strings.Split(ranges[1], "-")

		s.Rules = append(s.Rules, Rule{
			Name: classSep[0],
			LowerRange: Range{
				MinVal: utils.MustAtoi(lowRange[0]),
				MaxVal: utils.MustAtoi(lowRange[1]),
			},
			UpperRange: Range{
				MinVal: utils.MustAtoi(upperRange[0]),
				MaxVal: utils.MustAtoi(upperRange[1]),
			},
		})
	}

	// Parse my ticket
	myTicketStr := lines[1][1]
	nums := strings.Split(myTicketStr, ",")

	for _, num := range nums {
		s.MyTicket = append(s.MyTicket, utils.MustAtoi(num))
	}

	// Parse nearby tickets
	for _, line := range lines[2][1:] {
		var ticketNums []int
		ticketNumsStr := strings.Split(line, ",")
		for _, num := range ticketNumsStr {
			ticketNums = append(ticketNums, utils.MustAtoi(num))
		}
		s.NearbyTickets = append(s.NearbyTickets, ticketNums)
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	var vals []int

	for _, ticket := range s.NearbyTickets {
		validTicket := true
		for _, num := range ticket {
			if !s.Rules.ValidForAnyField(num) {
				vals = append(vals, num)
				validTicket = false
				break
			}
		}
		if validTicket {
			s.ValidTickets = append(s.ValidTickets, ticket)
		}
	}

	sum := 0
	for _, val := range vals {
		sum += val
	}

	return strconv.Itoa(sum), nil
}

func (s *Solver) Part2() (string, error) {

	// Consider our ticket as a valid ticket
	s.ValidTickets = append(s.ValidTickets, s.MyTicket)

	nCols := len(s.MyTicket)
	nRows := len(s.ValidTickets)

	var cols Columns

	// Compute valid fields for each column. A field is valid for a column if all
	// values in the column satisfy the rules for the field.
	for j := 0; j < nCols; j++ {
		var colSet []utils.StringSet
		for i := 0; i < nRows; i++ {
			colSet = append(colSet, s.Rules.ValidFieldsForVal(s.ValidTickets[i][j]))
		}
		cols = append(cols, Column{
			Index:       j,
			ValidFields: utils.StringSetIntersection(colSet...),
		})
	}

	fieldAttributions := make(map[string]int)
	size := len(cols)

	// Attribute fields to columns
	for i := 0; i < size; i++ {
		// Sort the columns so the first column is the one with less valid fields
		sort.Sort(BySetSize(cols))

		if len(cols[0].ValidFields.Set) == 1 {
			// There's only 1 valid field for this column, assign the field to the column
			elems := cols[0].ValidFields.Elements()
			fieldAttributions[elems[0]] = cols[0].Index

			for _, col := range cols {
				// Remove this field from the valid fields of the other columns
				col.ValidFields.Remove(elems[0])
			}

			// Remove the column we just processed from the list
			if len(cols) > 1 {
				cols = cols[1:]
			} else {
				cols = []Column{}
			}
		}
	}

	// Compute final result
	res := 1

	for field, pos := range fieldAttributions {
		if strings.HasPrefix(field, "departure") {
			res *= s.MyTicket[pos]
		}
	}

	return strconv.Itoa(res), nil
}
