package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ghonzo/advent2022/common"
)

type member struct {
	name string
	day  map[int]stars
}

type stars struct {
	part1, part2 time.Time
}

func main() {
	var endpoint, session string
	var day int
	flag.StringVar(&endpoint, "endpoint", os.Getenv("LEADERBOARD_URL"), "URL of the leaderboard JSON endpoint. Can also set the LEADERBOARD_URL env variable.")
	flag.StringVar(&session, "session", os.Getenv("LEADERBOARD_SESSION"), "session cookie value. Can also set the LEADERBOARD_SESSION env variable.")
	flag.IntVar(&day, "day", 0, "day to display, or most recent if not provided")
	flag.Parse()
	if len(endpoint) == 0 {
		fmt.Println("No endpoint provided")
		flag.Usage()
		os.Exit(1)
	}
	if len(session) == 0 {
		fmt.Println("No session provided")
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println("Calling", endpoint)
	var body []byte
	var err error
	if strings.HasPrefix(endpoint, "file://") {
		body, err = callFileEndpoint(endpoint)
	} else {
		body, err = callHttpEndpoint(endpoint, session)
	}
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	fmt.Println(renderResults(body, day))
}

func callHttpEndpoint(endpoint, session string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "session="+session)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func callFileEndpoint(endpoint string) ([]byte, error) {
	input, err := os.Open(strings.TrimPrefix(endpoint, "file://"))
	if err != nil {
		return []byte{}, err
	}
	return io.ReadAll(input)
}

var zeroTime = time.Time{}

func renderResults(body []byte, day int) string {
	var result map[string]any
	var members []member
	json.Unmarshal(body, &result)
	membersJson := result["members"].(map[string]any)
	for _, v := range membersJson {
		attrMap := v.(map[string]any)
		m := member{name: attrMap["name"].(string)}
		m.day = make(map[int]stars)
		completionDayLevelMap := attrMap["completion_day_level"].(map[string]any)
		for dayStr, v2 := range completionDayLevelMap {
			starsMap := v2.(map[string]any)
			part1Map := starsMap["1"].(map[string]any)
			s := stars{part1: time.Unix(int64(part1Map["get_star_ts"].(float64)), 0)}
			if part2Map, ok := starsMap["2"].(map[string]any); ok {
				s.part2 = time.Unix(int64(part2Map["get_star_ts"].(float64)), 0)
			}
			m.day[common.Atoi(dayStr)] = s
		}
		members = append(members, m)
	}
	// Figure out the day. If zero, then the maximum day
	if day == 0 {
		for _, m := range members {
			for k := range m.day {
				if k > day {
					day = k
				}
			}
		}
	}
	eastCoastLocation, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic("bad time zone")
	}
	dayStart := time.Date(2022, time.December, day, 0, 0, 0, 0, eastCoastLocation)
	// Just return the names that have at least one star on the day
	var sb strings.Builder
	const fmtString = "%-20s | %11v | %11v"
	sb.WriteString(fmt.Sprintf(fmtString, "Day "+fmt.Sprint(day), "Part 1", "Part 2"))
	sb.WriteString("\n------------------------------------------------\n")
	for _, m := range membersForDay(members, day) {
		if stars, ok := m.day[day]; ok {
			part1Duration := stars.part1.Sub(dayStart)
			if stars.part2 == zeroTime {
				sb.WriteString(fmt.Sprintf(fmtString, m.name, part1Duration, "-"))
			} else {
				sb.WriteString(fmt.Sprintf(fmtString, m.name, part1Duration, stars.part2.Sub(dayStart)))
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

// All the members that have at least one star for the day, sorted with fastest first
func membersForDay(members []member, day int) []member {
	var selectedMembers []member
	for _, m := range members {
		if _, ok := m.day[day]; ok {
			selectedMembers = append(selectedMembers, m)
		}
	}
	sort.Slice(selectedMembers, func(i, j int) bool {
		if selectedMembers[i].day[day].part2 == zeroTime {
			if selectedMembers[j].day[day].part2 == zeroTime {
				return selectedMembers[i].day[day].part1.Before(selectedMembers[j].day[day].part1)
			}
			return false
		}
		if selectedMembers[j].day[day].part2 == zeroTime {
			return true
		}
		return selectedMembers[i].day[day].part2.Before(selectedMembers[j].day[day].part2)
	})
	return selectedMembers
}
