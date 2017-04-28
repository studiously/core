// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "studiously": class Resource Client
//
// Command:
// $ goagen
// --design=github.com/studiously/core/design
// --out=$(GOPATH)/src/github.com/studiously/core
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListClassPath computes a request path to the list action of class.
func ListClassPath() string {

	return fmt.Sprintf("/classes")
}

// Get all classes a user is in
func (c *Client) ListClass(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListClassRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClassRequest create the request corresponding to the list action endpoint of the class resource.
func (c *Client) NewListClassRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowClassPath computes a request path to the show action of class.
func ShowClassPath(classID int) string {
	param0 := strconv.Itoa(classID)

	return fmt.Sprintf("/classes/%s", param0)
}

// Get class by ID
func (c *Client) ShowClass(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowClassRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowClassRequest create the request corresponding to the show action endpoint of the class resource.
func (c *Client) NewShowClassRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowMembersClassPath computes a request path to the show_members action of class.
func ShowMembersClassPath(classID int) string {
	param0 := strconv.Itoa(classID)

	return fmt.Sprintf("/classes/%s/members", param0)
}

// Get members of a class
func (c *Client) ShowMembersClass(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowMembersClassRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowMembersClassRequest create the request corresponding to the show_members action endpoint of the class resource.
func (c *Client) NewShowMembersClassRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowQuestionsClassPath computes a request path to the show_questions action of class.
func ShowQuestionsClassPath(classID int) string {
	param0 := strconv.Itoa(classID)

	return fmt.Sprintf("/classes/%s/questions", param0)
}

// Get questions for a class
func (c *Client) ShowQuestionsClass(ctx context.Context, path string, answered *bool, authorID *int, questionType *string, unitID *int) (*http.Response, error) {
	req, err := c.NewShowQuestionsClassRequest(ctx, path, answered, authorID, questionType, unitID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowQuestionsClassRequest create the request corresponding to the show_questions action endpoint of the class resource.
func (c *Client) NewShowQuestionsClassRequest(ctx context.Context, path string, answered *bool, authorID *int, questionType *string, unitID *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if answered != nil {
		tmp7 := strconv.FormatBool(*answered)
		values.Set("answered", tmp7)
	}
	if authorID != nil {
		tmp8 := strconv.Itoa(*authorID)
		values.Set("author_id", tmp8)
	}
	if questionType != nil {
		values.Set("question_type", *questionType)
	}
	if unitID != nil {
		tmp9 := strconv.Itoa(*unitID)
		values.Set("unit_id", tmp9)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
