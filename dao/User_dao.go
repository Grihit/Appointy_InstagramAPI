package dao

import (
	"log"

	. "github.com/grihit/Appointy_InstagramAPI/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "Users"
)
const (
	PCOLLECTION = "Posts"
)

// Establish a connection to database
func (m *UserDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of users
func (m *UserDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

//Find list of all posts
func (m *UserDAO) FindAllPosts() ([]Post, error) {
	var posts []Post
	err := db.C(PCOLLECTION).Find(bson.M{}).All(&posts)
	return posts, err
}

//Find list of all posts by user
func (m *UserDAO) FindAllUserPosts(id string) ([]Post, error) {
	var posts []Post
	err := db.C(PCOLLECTION).Find(bson.M{"Userid": id}).All(&posts)
	return posts, err
}

// Find a user by its id
func (m *UserDAO) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Find posts by post id
func (m *UserDAO) FindPostById(id string) (Post, error) {
	var post Post
	err := db.C(PCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&post)
	return post, err
}

// Insert a user into database
func (m *UserDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

//Insert a post into database
func (m *UserDAO) InsertPost(post Post) error {
	err := db.C(PCOLLECTION).Insert(&post)
	return err
}


