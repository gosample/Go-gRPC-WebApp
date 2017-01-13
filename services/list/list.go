package services

 import (
   "fmt"
   "golang.org/x/net/context"
   ghResponse "stars-app/messages/ghResponse"
 )

type GitHubServices struct{}

 func (m *GitHubServices) ListStars(c context.Context, s *ghResponse.List) (*ghResponse.List, error) {
 	fmt.Printf("rpc request Echo\n")
 	return s, nil
 }
