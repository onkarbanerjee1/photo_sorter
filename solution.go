package solution

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type photo struct {
	name,
	newName,
	location string
	pos     int
	takenAt time.Time
}

type byTime []*photo

type byID []*photo

// Solution provides the solution
func Solution(S string) string {
	// write your code in Go 1.4
	lines := strings.Split(S, "\n")
	photos, locations := []*photo{}, map[string][]*photo{}
	for i := 0; i < len(lines); i++ {
		phot, err := newPhoto(lines[i], i)
		if err != nil {
			fmt.Println("Exiting on error", err)
			os.Exit(1)
		}
		photos = append(photos, phot)

		if _, ok := locations[phot.location]; !ok {
			locations[phot.location] = []*photo{phot}
		} else {
			locations[phot.location] = append(locations[phot.location], phot)
		}
	}

	for loc, phots := range locations {
		byTimes := byTime(phots)
		sort.Sort(byTimes)

		max := len(byTimes)
		for i := 0; i < max; i++ {

			parts := strings.Split(byTimes[i].name, ".")
			if len(parts) != 2 {
				fmt.Println("Invalid format, can't get extension")
				os.Exit(1)
			}
			ext := parts[1]

			number := getNumber(i+1, max)
			byTimes[i].newName = strings.Join([]string{loc, number, ".", ext}, "")

		}
	}

	vals := []string{}
	byIDs := byID(photos)
	sort.Sort(byIDs)
	for _, each := range byIDs {
		vals = append(vals, each.newName)
	}
	return strings.Join(vals, "\n")
}

// newPhoto returns a new Photo given from a photo name record
func newPhoto(s string, pos int) (*photo, error) {
	parts := strings.Split(s, ",")
	if len(parts) < 3 {
		return nil, fmt.Errorf("Invalid photo format %s", s)
	}
	t, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(parts[2]))
	if err != nil {
		return nil, fmt.Errorf("Invalid time in photo format %s", err)
	}
	return &photo{name: strings.TrimSpace(parts[0]),
		location: strings.TrimSpace(parts[1]),
		takenAt:  t,
		pos:      pos}, nil
}

func getNumber(i, max int) string {
	maxStr := strconv.Itoa(max)
	iStr := strconv.Itoa(i)
	maxLength := len(maxStr)
	numberOfZeros := maxLength - len(iStr)
	var result string
	for j := 0; j < numberOfZeros; j++ {
		result = result + "0"
	}
	return result + iStr
}

func (a byTime) Len() int           { return len(a) }
func (a byTime) Less(i, j int) bool { return a[i].takenAt.Before(a[j].takenAt) }
func (a byTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (b byID) Len() int           { return len(b) }
func (b byID) Less(i, j int) bool { return b[i].pos < b[j].pos }
func (b byID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
