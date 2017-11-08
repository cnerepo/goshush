package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
	"strings"
	"time"
)

type account struct {
	Name     string
	Password string
}

func initMongo() *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    []string{"33.33.33.11"},
		Timeout:  60 * time.Second,
		Database: "goshush",
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		fmt.Println("Could not connect to database...")
		os.Exit(2)
	}

	session.SetMode(mgo.Monotonic, true)
	return session
}

func addAccount(n string, p string, s *mgo.Session) bool {
	n = strings.Trim(n, "\n")
	p = strings.Trim(p, "\n")
	c := s.DB("goshush").C("accounts")
	err := c.Insert(&account{Name: n, Password: p})

	if err != nil {
		fmt.Println("Ooops something went wrong saving to the database...")
		return false
	} else {
		s.Close()
		return true
	}
}
