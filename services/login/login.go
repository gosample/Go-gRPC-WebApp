package services

 import (
   "log"
   "errors"
   "stars-app/variables"
   "golang.org/x/net/context"
   "encoding/base64"
   "gopkg.in/mgo.v2"
   "gopkg.in/mgo.v2/bson"
   user "stars-app/messages/user"
 )

type AuthServices struct{}

func (m *AuthServices) Login(c context.Context, s *user.User) (*user.User, error) {

  session, err := mgo.Dial(variables.MongoAddr)
  if err != nil {
    log.Error("Mongo Connection Failed.");
    return nil, errors.New("Server Error");
  }
  defer session.Close()

  conn := session.DB("mongo").C("users")
  err = conn.Find(bson.M{"username": s.Username,"password":s.Password}).One(&s)
  if err != nil {
    //User not found
    s.Username="";
    s.Password="";
    s.Token="";
    return s, nil;
  }

  tokStr := []byte(s.Username + ":" + s.Password)
  tokEnc := base64.StdEncoding.EncodeToString(tokStr)
  s.Token=tokEnc;

  err = conn.Update(bson.M{"username":s.Username},bson.M{ "$set": s})
  if err != nil {
    //User not found while updating the token
    s.Username="";
    s.Password="";
    s.Token="";
    return s, nil;
  }

  return s, nil

}

func (m *AuthServices) CreateUser(c context.Context, s *user.User) (*user.User, error) {

  session, err := mgo.Dial(variables.MongoAddr)
  if err != nil {
    log.Error("Mongo Connection Failed.");
    return nil, errors.New("Server Error");
  }
  defer session.Close()

  conn := session.DB("mongo").C("users")

  var temp *user.User;
  err = conn.Find(bson.M{"username": s.Username}).One(&temp)
  if err != nil {

    //New user not found in DB. So create..
    tokStr := []byte(s.Username + ":" + s.Password)
    tokEnc := base64.StdEncoding.EncodeToString(tokStr)
    s.Token=tokEnc;
    err = conn.Insert(s);
    if err!=nil {
      //Error while inserting value
      log.Error("Mongo Connection Failed.");
      s.Token="";
      return s, nil;
    }
    return s, nil;

  } else{
    //New user already exists.
    s.Username="";
    s.Password="";
    return s, nil;
  }
}
