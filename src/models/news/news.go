package news

import (
	log "github.com/Sirupsen/logrus"
	"time"
	"github.com/dfernandez/geb/src/models/user"
	"gopkg.in/mgo.v2"
	"github.com/dfernandez/geb/config"
	"gopkg.in/mgo.v2/bson"
)

type News struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title     string
	Body      string
	Tags      []string
	Author    user.User
	CreatedTs time.Time
	EditedTs  time.Time
}

func Count(session *mgo.Session) int {
	c := session.DB(config.MongoDatabase).C("news")
	count, _ := c.Find(bson.M{}).Count()

	return count
}

func Newss(session *mgo.Session) []News {
	var news []News
	c := session.DB(config.MongoDatabase).C("news")
	c.Find(bson.M{}).All(&news)

	return news
}

func NewNews(title string, body string, author user.User) *News {
	return &News{
		Title:  title,
		Body:   body,
		Author: author,
	}
}

func (n *News) Created() string {
	return n.CreatedTs.Format(time.RFC822)
}

func (n *News) Edited() string {
	if n.EditedTs.IsZero() {
		return ""
	}

	return n.EditedTs.Format(time.RFC822)
}

func (n *News) Load(session *mgo.Session) {
	c := session.DB(config.MongoDatabase).C("news")
	err := c.Find(bson.M{"_id": n.Id}).One(n)

	if err != nil {
		log.Error(err)
	}
}

func (n *News) Insert(session *mgo.Session) {
	n.CreatedTs = time.Now()
	n.save(session)
}

func (n *News) Delete(session *mgo.Session) {
	c := session.DB(config.MongoDatabase).C("news")
	err := c.Remove(bson.M{"_id": n.Id})

	if err != nil {
		log.Error(err)
	}
}

func (n *News) save(session *mgo.Session) {
	c := session.DB(config.MongoDatabase).C("news")

	var err error

	if n.Id.Valid() {
		err = c.UpdateId(bson.M{"_id": n.Id}, n)
	} else {
		err = c.Insert(n)
	}

	if err != nil {
		log.Error(err)
	}
}