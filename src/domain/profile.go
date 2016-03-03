package domain

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/dfernandez/geb/config"
	"time"
)

type Profile struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Locale      string `json:"locale"`
	Picture     string `json:"picture"`
	LastLoginTs time.Time
}

func NewProfile(name string, email string, locale string, picture string) *Profile {
	p := &Profile{
		Name:    name,
		Email:   email,
		Locale:  locale,
		Picture: picture,
	}
	p.Init()

	return p
}

func (p *Profile) Init() {
	session, err := mgo.Dial(config.MongoServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	var profile Profile
	c := session.DB(config.MongoDatabase).C("profiles")
	c.FindId(p.Email).One(&profile)

	p.LastLoginTs = profile.LastLoginTs
}

func (p *Profile) IsAdmin() bool {
	for _, adm := range config.Administrators {
		if adm == p.Email {
			return true
		}
	}
	return false
}

func (p *Profile) GetProfiles() []Profile {
	session, err := mgo.Dial(config.MongoServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	var profiles []Profile
	c := session.DB(config.MongoDatabase).C("profiles")
	c.Find(bson.M{}).All(&profiles)

	return profiles
}

func (p *Profile) UpdateActivity() {
	p.LastLoginTs = time.Now()
	p.save()
}

func (p *Profile) LastLogin() string {
	return p.LastLoginTs.Format(time.RFC822)
}

func (p *Profile) save() {
	session, err := mgo.Dial(config.MongoServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDatabase).C("profiles")
	_, err = c.UpsertId(p.Email, p)
	if err != nil {
		log.Fatal(err)
	}
}
