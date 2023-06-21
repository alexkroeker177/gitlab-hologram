// This file contains types inside the gitlab domain.

// Package app contains structs and implementations for this app.
package app

import (
	"github.com/Bitspark/go-bitnode/bitnode"
	"log"
)

// Issue description: A map.
type Issue struct {
	// Assignee description: A type reference.
	Assignee ProjectMember `json:"assignee"`

	// Confidential description: A boolean.
	Confidential bool `json:"confidential"`

	// Description description: A string.
	Description string `json:"description"`

	// Title description: A string.
	Title string `json:"title"`
}

// ProjectMember description: A map.
type ProjectMember struct {
	// Name description: A string.
	Name string `json:"name"`

	// UserId description: A string.
	UserId string `json:"userId"`

	// Username description: A string.
	Username string `json:"username"`
}

// MergeRequest description: A map.
type MergeRequest struct {
	// Requester description: A type reference.
	Requester ProjectMember `json:"requester"`

	// Status description: A string.
	Status string `json:"status"`

	// BranchName description: A string.
	BranchName string `json:"branchName"`

	// Issue description: A type reference.
	Issue Issue `json:"issue"`
}

// Project description: A map.
type Project struct {
	// CommitCount description: An integer.
	CommitCount int `json:"commitCount"`

	// Commits description: A list of strings.
	Commits []Commit `json:"commits"`

	// Name description: A string.
	Name string `json:"name"`

	// ProjectId description: A string.
	ProjectId string `json:"projectId"`
}

// Commit description: A map.
type Commit struct {
	// Author description: A type reference.
	Author ProjectMember `json:"author"`

	// CommitHash description: A string.
	CommitHash string `json:"commitHash"`

	// CommitMessage description: A string.
	CommitMessage string `json:"commitMessage"`

	// Date description: A type reference.
	Date Date `json:"date"`
}

// Date description: A map.
type Date struct {
	// Seconds description: An integer.
	Seconds int `json:"seconds"`

	// Year description: An integer.
	Year int `json:"year"`

	// Day description: An integer.
	Day int `json:"day"`

	// Hours description: An integer.
	Hours int `json:"hours"`

	// Milliseconds description: An integer.
	Milliseconds int `json:"milliseconds"`

	// Minutes description: An integer.
	Minutes int `json:"minutes"`

	// MonthIndex description: An integer.
	MonthIndex int `json:"monthIndex"`
}

// DOMAIN STRUCT

// Domain containing mainly wrappers for applications.
type Domain struct {
	Domain *bitnode.Domain
	Node   bitnode.Node
}

// NewGitlab creates a new Gitlab instance.
func (gitlab *Domain) NewGitlab() (*Gitlab, error) {
	// Get the Gitlab sparkable from the domain.
	gitlabSpark, err := gitlab.Domain.GetSparkable("hub.gitlab.Gitlab")
	if err != nil {
		log.Fatal(err)
	}

	// Remove docker implementation.
	delete(gitlabSpark.Implementation, "docker")

	// Prepare the Gitlab spark.
	gitlabSpk, err := gitlab.Node.PrepareSystem(bitnode.Credentials{}, *gitlabSpark)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Gitlab.
	gitlab := &Gitlab{
		System: gitlabSpk,
	}

	return gitlab, nil
}
