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
	return &Profile{
		Name:    name,
		Email:   email,
		Locale:  locale,
		Picture: picture,
	}
}

func (p *Profile) Init(session *mgo.Session) {
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

func (p *Profile) GetProfiles(session *mgo.Session) []Profile {
	var profiles []Profile
	c := session.DB(config.MongoDatabase).C("profiles")
	c.Find(bson.M{}).All(&profiles)

	return profiles
}

func (p *Profile) UpdateActivity(session *mgo.Session) {
	p.LastLoginTs = time.Now()
	p.save(session)
}

func (p *Profile) LastLogin() string {
	return p.LastLoginTs.Format(time.RFC822)
}

func (p *Profile) save(session *mgo.Session) {
	c := session.DB(config.MongoDatabase).C("profiles")
	_, err := c.UpsertId(p.Email, p)
	if err != nil {
		log.Fatal(err)
	}
}
