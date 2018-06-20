package controllers

import (
	"gopkg.in/mgo.v2"

	"github.com/capt4ce/ecommerce-microservice/products/common"
)

// Close mgo.Session
func (c *mgo.Session) Close() {
	c.Close()
}

// Returns mgo.collection for the given name
func (c *mgo.Session) DbCollection(name string) *mgo.Collection {
	return c.DB(common.AppConfig.Database).C(name)
}

// Create a new Context object for each HTTP request
func NewContext() *mgo.Session {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
