package services

 import (
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
                return nil, errors.New("Server Error");
        }
        defer session.Close()
  conn := session.DB("mongo").C("users")
  err = conn.Find(bson.M{"username": s.Username,"password":s.Password}).One(&s)
  if err != nil {
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
                return nil, errors.New("Server Error");
        }
        defer session.Close()
  conn := session.DB("mongo").C("users")
  var temp *user.User;
  err = conn.Find(bson.M{"username": s.Username}).One(&temp)
  if err != nil {
    tokStr := []byte(s.Username + ":" + s.Password)
    tokEnc := base64.StdEncoding.EncodeToString(tokStr)
    s.Token=tokEnc;
    err = conn.Insert(s);
    if err!=nil {
      s.Token="";
      return s, nil;
    }
    return s, nil;
  } else{
    s.Username="";
    s.Password="";
    return s,nil;
  }
}
