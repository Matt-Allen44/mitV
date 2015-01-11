package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 4 {
	} else {
		fmt.Println("Incorrect args - check usage - go run MITV OAUTH_KEY REPO(/AUTHOR/NAME) HOST")
		os.Exit(1)
	}

	go launchServer()

	for true {
		update()

		time.Sleep(time.Second * 15)
	}
}

var pageContent string = ""

func launchServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(os.Args[3], nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pageContent)
}

func update() {
	fmt.Println("tick")

	oauth := string(os.Args[1])
	repo := string(os.Args[2])
	wordLimit := 65

	repo = strings.Replace(repo, "]", "", -1)

	fmt.Println("Starting w/ ", oauth, " @ ", "https://api.github.com/repos/"+repo+"/commits")

	s := (getTableHeader(repo))

	page := httpGetPageAuth("https://api.github.com/repos/"+repo+"/commits", oauth, false)
	commitJSON := make(GITHUB_COMMITS, 1000)

	_ = json.Unmarshal([]byte(page), &commitJSON)

	for i := range commitJSON {

		id := commitJSON[i].Commit.Tree.Sha
		id = id[:3]

		comment := commitJSON[i].Commit.Message

		if len(comment) > wordLimit {
			comment = comment[:wordLimit] + "..."
		}

		s += formatRow(i%2 == 0, strings.ToUpper(string(id)), strings.ToUpper(strings.Split(repo, "/")[1]), commitJSON[i].Commit.Committer.Date, commitJSON[i].Commit.Committer.Name, comment, commitJSON[i].HtmlUrl)
		//	fmt.Println(commitJSON[i].Commit.Message)
	}

	s += (getTableFooter())
	pageContent = s

	//writeFile("index.html", s)
}

func writeFile(name string, contents string) {
	fmt.Println("Writing to file: ", name)
	_ = ioutil.WriteFile(name, []byte(contents), 0644)
}

//var colourOdd string =
//var colourEven string =

func getTableHeader(repo string) string {

	rowPadding := "8"
	headPadding := "2"

	colEven := "#000000"
	colOdd := "#191919"
	colBg := "#191919"
	colHead := "#000000"
	colText := "#ffffff"
	sizeHeader := "100%"

	return `<html> <center> <title>mitV - ` + repo + `</title> 
<body bgcolor = "` + colBg + `">

<h1 style="font-weight:normal;color:` + colText + `;background-color:` + colBg + `;letter-spacing:1pt;word-spacing:2pt;font-size:` + sizeHeader + `;text-align:center;font-family:lucida sans unicode, lucida grande, sans-serif;line-height:1;
"> WATCHING COMMITS FOR ` + repo + ` </h1>

<style>
/* unvisited link */
a:link {
    color: #ffffff;
}

/* visited link */
a:visited {
    color: #ffffff;
}

/* mouse over link */
a:hover {
    color: #ffffff;
}

/* selected link */
a:active {
    color: #ffffff;
}
</style>

<style type="text/css">
	.tg  {border-collapse:collapse;border-spacing:0;}
	.tg td{font-family:Arial, sans-serif;font-size:14px;padding:` + rowPadding + `px 20px;;overflow:hidden;word-break:normal;}
	.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:` + headPadding + `px 20px;overflow:hidden;word-break:normal;}
	.tg .tg-li8k{font-family:"Lucida Sans Unicode", "Lucida Grande", sans-serif !important;;background-color:` + colHead + `;color:` + colText + `}
	

	.tg .tg-yrsx{background-color:` + colOdd + `;color:` + colText + `}
	.tg .tg-uimw{background-color:` + colHead + `;color:` + colText + `}


	.tg .tg-zh8g{background-color:` + colEven + `;color:` + colText + `}
	.tg .tg-u7t1{background-color:` + colHead + `;color:` + colText + `}
	</style>
	<table class="tg">
	  <tr>
	    <th class="tg-u7t1">Change</th>
	    <th class="tg-uimw">Product</th>
	    <th class="tg-uimw">Time/Date</th>
	    <th class="tg-uimw">Developer</th>
	    <th class="tg-li8k">                              Description                             </th>
	  </tr>
	`

}

func formatRow(even bool, changeId string, product string, time string, developer string, description string, url string) string {
	s := "	<tr>"

	class := ""
	if even {
		class = "tg-yrsx"
	} else {
		class = "tg-zh8g"
	}

	s += "<td class=\"" + class + "\"> <a href=\"" + url + "\">" + changeId + "</a></td>"
	s += "<td class=\"" + class + "\">" + product + "</td>"
	s += "<td class=\"" + class + "\">" + time + "</td>"
	s += "<td class=\"" + class + "\">" + developer + "</td>"
	s += "<td class=\"" + class + "\"; width=\"50%\">" + description + "</td>"

	s += "</tr>"
	return s
}

func getTableFooter() string {
	return "</table>"
}

var oauth string = ""

func httpGetPageAuth(url string, oauth string, verbose bool) string {
	if url == "" {
		return ""
	}

	httpReq, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "ERROR"
	}

	httpReq.SetBasicAuth(oauth, "x-oauth-basic")
	httpClient := http.Client{}

	httpRes, _ := httpClient.Do(httpReq)

	verbose = true
	if verbose {
		fmt.Println("[HTTP CONNECTION] Connected to ", url, " recieved ", httpRes.Status)
	}

	page, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return "ERROR"
	}

	return (string(page))
}

//--------------------------------------------------------------------------------------//

type GITHUB_COMMITS []struct {
	Sha    string `json:"sha"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			Url string `json:"url"`
		} `json:"tree"`
		Url          string `json:"url"`
		CommentCount int    `json:"comment_count"`
	} `json:"commit"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	CommentsUrl string `json:"comments_url"`
	Author      struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		Sha     string `json:"sha"`
		Url     string `json:"url"`
		HtmlUrl string `json:"html_url"`
	} `json:"parents"`
}
