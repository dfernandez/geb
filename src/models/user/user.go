package user

import (
    log "github.com/Sirupsen/logrus"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/dfernandez/geb/config"
    "time"
)

type User struct {
    Id             bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Name           string `json:"name"`
    Email          string `json:"email"`
    Locale         string `json:"locale"`
    Picture        string `json:"picture"`
    LastLoginTs    time.Time
    RegistrationTs time.Time
}

func Count(session *mgo.Session) int {
    c := session.DB(config.MongoDatabase).C("users")
    count, _ := c.Find(bson.M{}).Count()

    return count
}

func Users(session *mgo.Session) []User {
    var users []User
    c := session.DB(config.MongoDatabase).C("users")
    c.Find(bson.M{}).All(&users)

    return users
}

func NewUser(name string, email string, locale string, picture string) *User {
    return &User{
        Name:    name,
        Email:   email,
        Locale:  locale,
        Picture: picture,
    }
}

var IsAdmin = func(u interface{}) bool {
	for _, adm := range config.Administrators {
		if adm == u.(User).Email {
			return true
		}
	}

	return false
}

func (p *User) Init(session *mgo.Session) {
    var user User
    c := session.DB(config.MongoDatabase).C("users")
    err := c.Find(bson.M{"email": p.Email}).One(&user)

    // Non-existent user
    if err != nil {
        p.RegistrationTs = time.Now()
        return
    }

    p.Id             = user.Id
    p.LastLoginTs    = user.LastLoginTs
    p.RegistrationTs = user.RegistrationTs
}

func (p *User) UpdateActivity(session *mgo.Session) {
    p.LastLoginTs = time.Now()
    p.save(session)
}

func (p *User) LastLogin() string {
    if p.LastLoginTs.IsZero() {
        return ""
    }

    return p.LastLoginTs.Format(time.RFC822)
}

func (p *User) Registration() string {
    return p.RegistrationTs.Format(time.RFC822)
}

func (p *User) save(session *mgo.Session) {
    c := session.DB(config.MongoDatabase).C("users")
    _, err := c.Upsert(bson.M{"email": p.Email}, p)
    if err != nil {
        log.Error(err)
    }
}
