package main

import (
	"io/ioutil"
	"testing"
)

func TestGenerateIssueBody(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("testdata/normal.json")
	report, _ := parse(jsonData, nil)

	issueBody := generateIssueBody(report)
	expected := "* [ ] `com.github.ben-manes:gradle-versions-plugin:0.13.0`([Url](https://github.com/ben-manes/gradle-versions-plugin))\n* [ ] `com.google.firebase:firebase-core:9.4.0`\n"

	if issueBody != expected {
		t.Errorf("Expect \n%v\n, but \n%v", expected, issueBody)
	}
}
