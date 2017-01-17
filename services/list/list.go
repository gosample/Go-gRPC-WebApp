package services

 import (
   "fmt"
   "github.com/golang/glog"
   "stars-app/variables"
   "stars-app/messages/user"
   "gopkg.in/mgo.v2"
   "gopkg.in/mgo.v2/bson"
   "golang.org/x/net/context"
   "net/http"
   "errors"
   "io/ioutil"
   "stars-app/utils"
   "encoding/json"
   ghResponse "stars-app/messages/ghResponse"
 )

type GitHubServices struct{}

type Item struct {
	Repo  string `json:"name"`
	Stars int    `json:"stargazers_count"`
}

 func (m *GitHubServices) ListStars(c context.Context, s *ghResponse.List) (*ghResponse.List, error) {


   if s.Token == "" {
     //Token not found!
     s.Token="";
     return s, nil
   }


   session, err := mgo.Dial(variables.MongoAddr)
   if err != nil {
     glog.Error("Mongo Connection Failed.");
     return nil, errors.New("Server Error!");
   }
   defer session.Close()

   conn := session.DB("mongo").C("users")
   var us user.User;
   err = conn.Find(bson.M{"token":s.Token}).One(&us)
   if err != nil {
     //Token not found!
        s.Token="";
        return s, nil
   }



   req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/repos", s.GhUser), nil)
   req.SetBasicAuth(utils.GITHUB_USERNAME, utils.GITHUB_API_KEY)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
     glog.Error("GitHub Connection Error.");
     return nil, errors.New("Server Error!")
   }

   respBody, err := ioutil.ReadAll(resp.Body)
   if err != nil {
     glog.Error("GitHub Response Error.");
     return nil, errors.New("Server Error!")
   }

    var tempList []*ghResponse.ListGhList;
    err = json.Unmarshal(respBody, &tempList)
    if err != nil {
      //GH User not found!
      return s, nil
    }

    s.List=tempList;

 	return s, nil
 }
