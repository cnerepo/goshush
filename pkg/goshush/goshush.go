package goshush

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Service represents a connection to GoShush
type Service struct {
	session *mgo.Session
}

// Connect opens a connection to the GoShush Service
func Connect() (*Service, error) {

	info := &mgo.DialInfo{
		Addrs:    []string{"33.33.33.11"},
		Timeout:  60 * time.Second,
		Database: "goshush",
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return &Service{session}, nil
}

// Authenticate opens an authenticated session to AWS.
func Authenticate() error {
	_, err := session.NewSession()
	if err != nil {
		return err
	}

	return nil
}

// Add stores a new secret to the master data-store.
func (s *Service) Add(n string, p string) error {
	a := Account{
		Name:     strings.Trim(n, "\n"),
		Password: strings.Trim(p, "\n"),
	}
	c := s.session.DB("goshush").C("accounts")
	defer s.session.Close()

	if err := c.Insert(a); err != nil {
		fmt.Println("Ooops something went wrong saving to the database...")
		return err
	}
	return nil
}

// Search retrieves a secret that is an exact match of needle.
func (s *Service) Search(needle string) ([]Account, error) {
	needle = strings.Trim(needle, "\n")
	c := s.session.DB("goshush").C("accounts")
	defer s.session.Close()

	var results []Account
	err := c.Find(bson.M{"name": needle}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Delete removes a secret from the master data-store.
func (s *Service) Delete(needle string) error {
	needle = strings.Trim(needle, "\n")
	c := s.session.DB("goshush").C("accounts")
	defer s.session.Close()

	err := c.Remove(bson.M{"name": needle})
	if err != nil {
		return err
	}
	return nil
}
