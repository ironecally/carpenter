package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var githubClient *github.Client

func initGithub(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	if githubClient == nil {

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: GetGithubToken(),
			},
		)

		tc := oauth2.NewClient(ctx, ts)

		githubClient = github.NewClient(tc)
	}
}

func getRepoInfo(ctx context.Context, repositoryOwner string, repositoryName string) error {
	repo, _, err := githubClient.Repositories.Get(ctx, repositoryOwner, repositoryName)
	if err != nil {
		return err
	}

	log.Printf("%+v", repo)
	return nil
}
