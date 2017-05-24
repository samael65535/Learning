package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Member struct {
	Email     string
	Id        int
	Password  string
	FirstName string
}

type Session struct {
	Id        int
	MemberId  int
	SessionId string
}

func CreateSession(member Member) (Session, error) {
	r := Session{}
	r.MemberId = member.Id
	sessionId := sha256.Sum256([]byte(member.Email + time.Now().Format("12:00:00")))

	r.SessionId = hex.EncodeToString(sessionId[:])

	connect, err := getDBConnection()

	if err == nil {
		defer connect.Close()
		collection := connect.DB("webstore").C("session")
		colQuerier := bson.M{"member_id": member.Id, "session_id": r.SessionId}
		change := mgo.Change{
			Update:    bson.M{"$set": bson.M{"member_id": member.Id, "session_id": r.SessionId}},
			ReturnNew: false,
			Upsert:    true,
		}
		_, err := collection.Find(colQuerier).Apply(change, nil)
		if err == nil {
			return r, nil
		} else {
			return Session{}, errors.New("Unable to save session")
		}
	} else {
		return Session{}, errors.New("Unable to connect database")

	}

}
func GetMember(email string, password string) (Member, error) {
	connect, err := getDBConnection()

	if err == nil {
		defer connect.Close()
		collection := connect.DB("webstore").C("member")
		result := Member{}

		pwd := sha256.Sum256([]byte(password))
		colQuerier := bson.M{"email": email, "password": hex.EncodeToString(pwd[:])}
		err := collection.Find(colQuerier).One(&result)
		fmt.Println(hex.EncodeToString(pwd[:]))
		fmt.Println(password)
		fmt.Println(result)
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find Member with email: " + email)
		}
	} else {
		return Member{}, errors.New("Unable to connect to DB")
	}
}

func getDBConnection() (*mgo.Session, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
