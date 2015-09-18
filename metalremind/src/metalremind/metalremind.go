package metalremind

import (
	"bytes"
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	"github.com/google/go-github/github"
)

func init() {
	http.HandleFunc("/remind", handlerRemind)
	http.HandleFunc("/markdown", handlerMarkdown)
}

func handlerRemind(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	uc := urlfetch.Client(c)
	client := github.NewClient(uc)

	opt := &github.IssueListByRepoOptions{
		State:       "open",
		Sort:        "updated",
		Direction:   "desc",
		ListOptions: github.ListOptions{Page: 1, PerPage: 100},
	}
	issues, _, err := client.Issues.ListByRepo("sinmetal", "slide", opt)
	if err != nil {
		log.Errorf(c, "Issues.List returned error: %v", err)
		return
	}
	mc, err := GetMetalConfig(c)
	if err != nil {
		log.Errorf(c, "Metal Config get error: %v", err)
		return
	}

	for _, issue := range issues {
		sm := SlackMessage{}
		sm.Set(issue)
		_, err = postToSlack(c, mc.SlackPostURL, sm)
		if err != nil {
			log.Errorf(c, "%Slack Post error: %v", err)
		}
	}
	log.Infof(c, "%v", issues)
}

func handlerMarkdown(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	uc := urlfetch.Client(c)
	client := github.NewClient(uc)

	input := "# heading #\nLink to issue #1\n"
	md, _, err := client.Markdown(input, &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"})
	if err != nil {
		log.Errorf(c, "error: %v\n\n", err)
	}

	log.Infof(c, "converted markdown:\n%v\n", md)
}

type SlackMessage struct {
	UserName    string            `json:"username"`
	IconUrl     string            `json:"icon_url"`
	Text        string            `json:"text"`
	Attachments []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color      string       `json:"color"`
	AuthorName string       `json:"author_name"`
	AuthorLink string       `json:"author_link"`
	AuthorIcon string       `json:"author_icon"`
	Title      string       `json:"title"`
	TitleLink  string       `json:"title_link"`
	Fields     []SlackField `json:"fields"`
}

type SlackField struct {
	Title string `json:"title"`
}

func (sm *SlackMessage) Set(issue github.Issue) {
	fields := make([]SlackField, 0)
	for _, label := range issue.Labels {
		fields = append(fields, SlackField{
			Title: *label.Name,
		})
	}

	sa := SlackAttachment{
		Color:     "#36a64f",
		Title:     *issue.Title,
		TitleLink: *issue.HTMLURL,
		Fields:    fields,
	}
	if issue.User != nil {
		sa.AuthorName = *issue.User.Login
		sa.AuthorIcon = *issue.User.AvatarURL
	}

	sm.UserName = "metalremind"
	sm.IconUrl = "http://lh3.googleusercontent.com/IeLqEYxiOtooPD-2nRIzMA0DfaLWcsGz03kxSMdfXWyJLSJOa2v7n5Ec-EAMpc1isbACJg_OMGJj_9NC-HQYqQdToQ"
	sm.Text = *issue.Body
	sm.Attachments = []SlackAttachment{sa}
}

func postToSlack(c context.Context, url string, message SlackMessage) (resp *http.Response, err error) {
	client := urlfetch.Client(c)

	body, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	log.Infof(c, "%s", string(body))
	return client.Post(
		url,
		"application/json",
		bytes.NewReader(body))
}
