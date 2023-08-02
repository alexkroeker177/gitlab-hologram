package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Issue struct {
	// Assignee description: A type reference.
	Assignee User `json:"assignee"`

	// Confidential description: A boolean.
	Confidential bool `json:"confidential"`

	// Description description: A string.
	Description string `json:"description"`

	// Title description: A string.
	Title string `json:"title"`
}

type User struct {
	// Name description: A string.
	Name string `json:"name"`

	// UserId description: A string.
	UserId string `json:"userId"`

	// Username description: A string.
	Username string `json:"username"`
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
	Author User `json:"author"`

	// CommitHash description: A string.
	CommitHash string `json:"commitHash"`

	// CommitMessage description: A string.
	CommitMessage string `json:"commitMessage"`

	// Date description: A type reference.
	Date Date `json:"date"`
}

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

type Gitlab struct {
	httpClient *http.Client
}

func (s *Gitlab) apiCall(method string, call string, params map[string]string, reqStruct any, respStruct any) error {

	var reqBody io.Reader
	if reqStruct != nil {
		reqBts, err := json.Marshal(reqStruct)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(reqBts)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("https://gitlab.bitspark.com/api/v4/%s?%s", call, encodeMapToURLString(params)), reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PRIVATE-TOKEN", "glpat-uyLgZJobFAu6ozXosShM")

	resp, err := s.httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	respBts, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Test %s", respBts)
	respStructWrapper := struct {
		Success bool            `json:"success"`
		Data    json.RawMessage `json:"data"`
		Error   struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	if err := json.Unmarshal(respBts, &respStructWrapper); err != nil {
		return err
	}

	if !respStructWrapper.Success {
		return fmt.Errorf("[%d] %s", respStructWrapper.Error.Code, respStructWrapper.Error.Message)
	}

	if err := json.Unmarshal(respStructWrapper.Data, respStruct); err != nil {
		return err
	}

	return nil
}

func encodeMapToURLString(data map[string]string) string {
	values := url.Values{}
	for key, value := range data {
		values.Add(key, value)
	}

	return values.Encode()
}

func main() {
	params := map[string]string{
		"description": "Testing the Gitlab API for hologram use",
	}
	var respStruct []struct {
		Type       string `json:"type"`
		Attributes map[string]struct {
			Type        string `json:"type"`
			Label       string `json:"label"`
			Value       any    `json:"value"`
			UniversalId string `json:"universal_id"`
		} `json:"attributes"`
	}

	s := Gitlab{}

	/*issue := Issue{
		Title:       "14.07 Issue creation with API test",
		Description: "TESTESTTESTTEST",
	}*/
	//project := Project{ProjectId: "62"}

	//s.createIssue(issue, project)
	s.httpClient = &http.Client{}

	if err := s.apiCall("GET", "projects/62", params, nil, &respStruct); err != nil {
		fmt.Errorf("error while making the api call")
	}

	// Create new issue with API

	/*if err := s.apiCall("PUT", "projects/62/issues/15", params, nil, &respStruct); err != nil {
		fmt.Errorf("error while making post request")
	}*/

}

func formatStringForPostRequest(x string) string {
	return strings.ReplaceAll(x, " ", "%20")
}

func (s *Gitlab) createIssue(issue Issue, project Project) {
	params := map[string]string{
		"description": issue.Description,
		"title":       issue.Title,
	}

	var respStruct []struct {
		Type       string `json:"type"`
		Attributes map[string]struct {
			Type        string `json:"type"`
			Label       string `json:"label"`
			Value       any    `json:"value"`
			UniversalId string `json:"universal_id"`
		} `json:"attributes"`
	}

	call := fmt.Sprintf("projects/%s/issues", project.ProjectId)
	s.httpClient = &http.Client{} // In Hologram die HTTP Client Erzeugung im Init Cycle stattfinden lassen

	if err := s.apiCall("POST", call, params, nil, &respStruct); err != nil {
		fmt.Errorf("Error while trying to create an issue...")
	}
	fmt.Printf("Hell Oooo%s", respStruct)

}

