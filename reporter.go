package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type GithubReport struct {
	Title  string
	Report Report
}

func generateIssueBody(report Report) string {
	var body string
	var dependency Dependency

	for i := 0; i < len(report.Outdated.Dependencies); i++ {
		dependency = report.Outdated.Dependencies[i]
		body += fmt.Sprintf("* [ ] `%v:%v`", dependency.Pkg(), dependency.Available.Release)

		if len(dependency.ProjectUrl) > 0 {
			body += fmt.Sprintf("([Url](%v))\n", dependency.ProjectUrl)
		} else {
			body += "\n"
		}
	}
	return body
}

func reportToGithub(githubReport GithubReport, githubAccessToken, userName, repositoryName string) error {
	var pkgs string
	for _, dependency := range githubReport.Report.Outdated.Dependencies {
		pkgs += dependency.Pkg() + ","
	}
	pkgs = strings.TrimRight(pkgs, ",")

	body := generateIssueBody(githubReport.Report)
	fmt.Printf("%+v\n", body)
	if len(body) == 0 {
		// No libraries need to update
		return nil
	}

	title := githubReport.Title
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	issueRequest := &github.IssueRequest{Title: &title, Body: &body}

	ctx := context.Background()
	_, _, err := client.Issues.Create(ctx, userName, repositoryName, issueRequest)
	if err != nil {
		return errors.Wrap(err, "issue create failed")
	}
	return nil
}
