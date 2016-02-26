package domain

import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/dfernandez/geb/config"
)

type Profile struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Locale  string `json:"locale"`
	Picture string `json:"picture"`
}

func (p *Profile) UpdateActivity() {
	p.save()
}

func (p *Profile) save() {
	session, err := mgo.Dial(config.MongoServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDatabase).C("profile")
	_, err = c.UpsertId(p.Email, p)
	if err != nil {
		log.Fatal(err)
	}
}
