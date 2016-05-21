package main

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func generateIssueBody(report Report) string {
	var body string
	for _, element := range report.Outdated.Dependencies {
		body += fmt.Sprintf("* [ ] `%v:%v:%v`\n", element.Group, element.Name, element.Available.Release)
	}
	return body
}

func reportToGithub(report Report, githubAccessToken, userName, repositoryName string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	currentTime := time.Now()
	title := "dependency-updates-" + currentTime.Format("20060102150405")
	body := generateIssueBody(report)
	issueRequest := &github.IssueRequest{Title: &title, Body: &body}

	_, _, err := client.Issues.Create(userName, repositoryName, issueRequest)
	if err != nil {
		return err
	}
	return nil
}
