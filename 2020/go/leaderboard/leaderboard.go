package leaderboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Leaderboard struct {
	Members map[string]MemberInfo `json:"members"`
	OwnerId string                     `json:"owner_id"`
	Event   string                     `json:"event"`
}

type ByTimestamp []Star
func (s ByTimestamp) Len() int           { return len(s) }
func (s ByTimestamp) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByTimestamp) Less(i, j int) bool { return s[i].Timestamp.Before(s[j].Timestamp) }

func (l *Leaderboard) Stars() []Star {
	var stars []Star

	for _,member := range l.Members {
		for day, dayInfo := range member.CompletionDayLevel {
			for level, levelInfo := range dayInfo {
				stars = append(stars, Star{
					MemberId:   member.Id,
					MemberName: member.Name,
					Day:        day,
					Level:      level,
					Timestamp:  levelInfo.StarTimestamp,
				})
			}
		}
	}

	return stars
}


func (l *Leaderboard) StarsByMemberId(id string) []Star {
	var stars []Star

	mInfo, ok :=  l.Members[id]
	if !ok {
		return stars
	}

	for day, dayInfo := range mInfo.CompletionDayLevel {
		for level, levelInfo := range dayInfo {
			stars = append(stars, Star{
				MemberId:   mInfo.Id,
				MemberName: mInfo.Name,
				Day:        day,
				Level:      level,
				Timestamp:  levelInfo.StarTimestamp,
			})
		}
	}

	return stars
}

func (l *Leaderboard) StarsByDay(day int) []Star {
	var stars []Star
	for _,member := range l.Members {
		dayInfo, ok := member.CompletionDayLevel[day]
		if !ok {
			continue
		}
		for level, levelInfo := range dayInfo {
			stars = append(stars, Star{
				MemberId:   member.Id,
				MemberName: member.Name,
				Day:        day,
				Level:      level,
				Timestamp:  levelInfo.StarTimestamp,
			})
		}
	}

	return stars
}

type Star struct {
	MemberId string
	MemberName string
	Day int
	Level int
	Timestamp time.Time
}

type leaderboardReply struct {
	Members map[string]memberInfoReply `json:"members"`
	OwnerId string                     `json:"owner_id"`
	Event   string                     `json:"event"`
}

func (l *leaderboardReply) toLeaderboard() (*Leaderboard, error) {
	leaderboard :=  Leaderboard{
		Members: make(map[string]MemberInfo),
		OwnerId: l.OwnerId,
		Event:   l.Event,
	}

	for mId, info := range l.Members {
		mInfo, err  := info.toMemberInfo()
		if err != nil {
			return nil, fmt.Errorf("could not convert member info: %v", err)
		}
		leaderboard.Members[mId] = *mInfo
	}

	return &leaderboard, nil
}


type MemberInfo struct {
	LastStarTimestamp  *time.Time                `json:"last_star_ts"`
	GlobalScore        int                      `json:"global_score"`
	LocalScore         int                      `json:"local_score"`
	Id                 string                   `json:"id"`
	CompletionDayLevel map[int]map[int]StarInfo `json:"completion_day_level"`
	Stars              int                      `json:"starts"`
	Name               string                   `json:"name"`
}

type memberInfoReply struct {
	LastStarTimestamp  interface{}                   `json:"last_star_ts"`
	GlobalScore        int                           `json:"global_score"`
	LocalScore         int                           `json:"local_score"`
	Id                 string                        `json:"id"`
	CompletionDayLevel map[int]map[int]starInfoReply `json:"completion_day_level"`
	Stars              int                           `json:"starts"`
	Name               string                        `json:"name"`
}

func (m *memberInfoReply) toMemberInfo() (*MemberInfo, error) {
	res := MemberInfo {
		GlobalScore: m.GlobalScore,
		LocalScore: m.LocalScore,
		Id: m.Id,
		Stars: m.Stars,
		Name: m.Name,
		CompletionDayLevel: make(map[int]map[int]StarInfo),
	}

	if f, ok := m.LastStarTimestamp.(string); ok {

		ts, err := strconv.ParseInt(f,10,64)
		if err != nil {
			return &res, fmt.Errorf("could not parse %v as an int unix timestamp: %v", m.LastStarTimestamp, err)
		}
		t := time.Unix(ts, 0)
		res.LastStarTimestamp = &t

	} else if _, ok := m.LastStarTimestamp.(float64); ok {
		res.LastStarTimestamp = nil
	} else {
		return nil, fmt.Errorf("could not parse last_star_ts as either string or int")
	}

	for day, levels :=  range m.CompletionDayLevel {
		if _, ok := res.CompletionDayLevel[day]; !ok {
			res.CompletionDayLevel[day] = make(map[int]StarInfo)
		}

		for level, sInfo := range levels {
			ts, err := strconv.ParseInt(sInfo.StarTimestamp,10,64)
			if err != nil {
				return nil, fmt.Errorf("could not parse %v as an int unix timestamp: %v", sInfo.StarTimestamp, err)
			}
			res.CompletionDayLevel[day][level] = StarInfo{StarTimestamp: time.Unix(ts, 0)}
		}
	}

	return &res, nil
}

type StarInfo struct {
	StarTimestamp time.Time `json:"get_star_ts"`
}

type starInfoReply struct {
	StarTimestamp string `json:"get_star_ts"`
}

func (s* starInfoReply) toStarInfo() (StarInfo, error) {
	var res StarInfo

	ts, err := strconv.ParseInt(s.StarTimestamp,10,64)
	if err != nil {
		return res, fmt.Errorf("could not parse %v as an int unix timestamp: %v", s.StarTimestamp, err)
	}
	res.StarTimestamp = time.Unix(ts, 0)

	return res, nil

}

func FetchLeaderboard(session string, leaderboardId string, year int) (*Leaderboard, error) {
	url := fmt.Sprintf("https://adventofcode.com/%v/leaderboard/private/view/%v.json", year, leaderboardId)
	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GET request for AoC input: %w", err)
	}

	req.Header.Add("cookie", fmt.Sprintf("session=%s;", session))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing GET request for AoC input: %w", err)
	}
	defer resp.Body.Close()

	input, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body request of GET request for AoC input: %w", err)
	}

	return LeaderboardFromJson(input)
}


func LeaderboardFromJson(leaderboard []byte) (*Leaderboard, error) {
	var lReply leaderboardReply

	err := json.Unmarshal(leaderboard, &lReply)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling leaderboard data: %w\n%v", err, string(leaderboard))
	}

	l, err := lReply.toLeaderboard()
	if err != nil {
		return nil, fmt.Errorf("converting json into leaderboard: %w", err)
	}

	return l, nil
}
