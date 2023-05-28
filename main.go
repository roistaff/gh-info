package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/fatih/color"
)

type UserResponse struct {
	Company   string `json:"company"`
	Location  string `json:"location"`
	Blog      string `json:"blog"`
	Bio       string `json:"bio"`
	Twitter   string `json:"twitter_username"`
	Repos     int    `json:"public_repos"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Join      string `json:"created_at"`
	Username  string `json:"login"`
	Message   string `json:"message"`
}
type LicenseInfo struct {
	Name string `json:"name"`
}
type OwnerInfo struct {
	Name string `json:"login"`
}
type RepoResponse struct {
	Fullname    string      `json:"full_name"`
	URL         string      `json:"html_url"`
	Fork        bool        `json:"fork"`
	Update      string      `json:"pushed_at"`
	Lang        string      `json:"language"`
	Message     string      `json:"message"`
	License     LicenseInfo `json:"license"`
	Owner       OwnerInfo   `json:"owner"`
	Pages       bool        `json:"has_pages"`
	Branch      string      `json:"default_branch"`
	Description string      `json:"description"`
	Star        int         `json:"stargazers_count"`
}

func RenderTerm(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n")
	if string(body) == "404: Not Found" {
		color.Red("README.md is not found.")
		os.Exit(1)
	}
	result := markdown.Render(string(body), 80, 6)
	fmt.Println(string(result))
}
func getAPI(url2 string) []byte {
	const url string = "https://api.github.com/"
	client := http.Client{
		Timeout: 20 * time.Second,
	}
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		url+url2,
		nil,
	)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		color.Red("Error: Off-line.Please connect.")
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func getInfoRepo(fullname string) {
	var data RepoResponse
	body := getAPI("repos/" + fullname)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	if data.Message == "Not Found" {
		color.Red("Repo '" + fullname + "' Not Found")
		os.Exit(1)
	}
	datalist := color.GreenString("\n"+data.Fullname) + "\t\t" + color.YellowString("☆Star:") + strconv.Itoa(data.Star) + color.HiMagentaString("\nDescription: ") + data.Description + color.HiCyanString("\nLanguage: ") + data.Lang + color.BlueString("\nLicense: ") + data.License.Name + color.MagentaString("\nURL: ") + data.URL
	fmt.Println(datalist)
}
func getInfoUser(username string) {
	var data UserResponse
	body := getAPI("users/" + username)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	if data.Message == "Not Found" {
		color.Red("user ' " + username + "' Not Found.")
		os.Exit(1)
	} else {
		fmt.Println("user '" + username + "' found.")
		datalist := color.YellowString("\nUsername: ") + data.Username + color.HiMagentaString("\nProfile: ") + data.Bio + color.HiBlueString("\nFollower: ") + strconv.Itoa(data.Followers) + color.HiCyanString("\tFollowing: ") + strconv.Itoa(data.Following) + color.HiGreenString("\nCompany: ") + data.Company + color.MagentaString("\nLocation: ") + data.Location + color.GreenString("\nBlog: ") + data.Blog + color.HiBlueString("\tTwitter: ") + data.Twitter + color.CyanString("\nRepos: ") + strconv.Itoa(data.Repos)
		fmt.Println(datalist)
	}
}
func getLinkREADME(fullname string) string {
	var data RepoResponse
	body := getAPI("repos/" + fullname)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	if data.Message == "Not Found" {
		color.Red("Repo '" + fullname + "' Not Found")
		os.Exit(1)
	}
	branch := data.Branch
	url := "https://raw.githubusercontent.com/" + fullname + "/" + branch + "/README.md"
	return url
}
func welcome() {
	welcometext := ` ____  _  _    _             _       ___          __        
 / ___|(_)| |_ | |__   _   _ | |__   |_ _| _ __   / _|  ___  
| |  _ | || __|| '_ \ | | | || '_ \   | | | '_ \ | |_  / _ \ 
| |_| || || |_ | | | || |_| || |_) |  | | | | | ||  _|| (_) |
 \____||_| \__||_| |_| \__,_||_.__/  |___||_| |_||_|   \___/ 
        version:0.1.2`
	fmt.Println(welcometext)
}
func Helps(message ...string) {
	welcome()
	if message != nil {
		fmt.Println(color.RedString("! Message: "), message)
	}
	helptext := `
Usage:
        gh info user [username]` + color.HiBlackString(" To find user info") + `
        gh info repo [user/repo]` + color.HiBlackString(" To find repo info. For exsample: gh info roistaff/gh-info") + `
        gh info read [user/repo]` + color.HiBlackString(" To show repo's README.md") + `
        gh info help ` + color.HiBlackString(" To show helps") + `
Author:
        ` + color.HiYellowString("Staff Roi") + ` [roistaff1983@gmail.com]
        Please type "gh info user roistaff"
Futures:
        - add search command (search repos)
More about this:
        Please visit ` + color.HiGreenString("https://github.com/roistaff/gh-info")
	fmt.Println(helptext)
	os.Exit(1)
}
func main() {
	color.HiGreen("Github Info\n")
	args := os.Args
	if len(args) != 1 {
		if len(args) == 2 {
			if args[1] == "help" {
				Helps()
			} else if args[1] == "repo" || args[1] == "user" || args[1] == "read" {
				color.Red("× Error: Not enough arguments.")
				os.Exit(1)
			} else {
				Helps("Unknown command '" + args[1] + "'")
			}
		}
		if len(args) == 3 {
			if args[1] != "user" && args[1] != "repo" && args[1] != "help" && args[1] != "read" {
				Helps("Unknown command '" + args[1] + "'")
			} else if args[1] == "user" {
				getInfoUser(args[2])
			} else if args[1] == "repo" {
				getInfoRepo(args[2])
			} else if args[1] == "help" {
				Helps()
			} else if args[1] == "read" {
				url := getLinkREADME(args[2])
				RenderTerm(url)
			}
		}
	} else {
		Helps()
	}
}
