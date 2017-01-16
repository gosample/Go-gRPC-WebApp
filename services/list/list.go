package services

 import (
   "fmt"
   "golang.org/x/net/context"
   "encoding/base64"
   "net/http"
   "errors"
   "io/ioutil"
   "stars-app/utils"
   "strings"
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
     return nil, errors.New("Token error")
   }

   tokStr, err := base64.StdEncoding.DecodeString(s.Token)
   if err != nil {
     panic(err)
   }
   tokParts := strings.Split(string(tokStr), ":")
   usernameMatch := ""
   passwdMatch := ""

     if tokParts[0] == "admin" {
       usernameMatch = "admin"
     }
     if tokParts[1] == "password" {
       passwdMatch = "password"
     }

   if usernameMatch == "" {
     return nil, errors.New("User Not Found")
   }
   if passwdMatch == "" {
     return nil, errors.New("Password Not Found")
   }

   req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/repos", s.GhUser), nil)
   req.SetBasicAuth(utils.GITHUB_USERNAME, utils.GITHUB_API_KEY)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
     return nil, err
   }

   respBody, err := ioutil.ReadAll(resp.Body)
   if err != nil {
     return nil, err
   }
   var items []Item

   err = json.Unmarshal(respBody, &items)
   if err != nil {
     return nil, err
   }
    var tempList []*ghResponse.ListGhList;
    err = json.Unmarshal(respBody, &tempList)
    if err != nil {
      return nil, err
    }

    s.List=tempList;

 	return s, nil
 }
