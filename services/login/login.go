package services

 import (
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
                panic(err)
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
