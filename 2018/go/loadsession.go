package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type AocSession struct {
	CookieValue string `json:"session"`
}

func GetSession() (*AocSession, error) {
	var session AocSession

	loaded := loadSessionFromEnv(&session)
	if loaded {
		return &session, nil
	}

	loaded = loadSessionFromJson(&session, "session.json")
	if loaded {
		return &session, nil
	}

	return nil, errors.New("Could not load session")
}

func loadSessionFromEnv(session *AocSession) bool {
	value, present := os.LookupEnv("AOC_SESSION")
	if present {
		session.CookieValue = value
		return true
	}
	return false
}

func loadSessionFromJson(session *AocSession, filename string) bool {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return false
	}

	json.Unmarshal(data, session)

	// fmt.Printf("Session: %v\n", session)
	return true
}
