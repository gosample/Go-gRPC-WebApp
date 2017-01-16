package services

 import (
   "fmt"
   "golang.org/x/net/context"
   "encoding/base64"
   "errors"
   user "stars-app/messages/user"
 )

type AuthServices struct{}

func (m *AuthServices) Login(c context.Context, s *user.User) (*user.User, error) {
	
  if(s.Username == "admin" && s.Password == "password"){
    tokStr := []byte(s.Username + ":" + s.Password)
    tokEnc := base64.StdEncoding.EncodeToString(tokStr)
    s.Token=tokEnc;
    return s, nil
  } else{
    return nil, errors.New("User Not Found");
  }

}
