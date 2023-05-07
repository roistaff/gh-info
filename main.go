package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "github.com/fatih/color"
        "strconv"
)
type UserResponse struct {
        Company    string `json:"company"`
        Location   string `json:"location"`
        Blog       string `json:"blog"`
        Bio        string `json:"bio"`
        Twitter    string `json:"twitter_username"`
        Repos      int    `json:"public_repos"`
        Followers  int    `json:"followers"`
        Following  int    `json:"following"`
        Join       string `json:"created_at"`
        Username   string `json:"login"`
	Message    string `json:"message"`
}
type LicenseInfo struct {
        Name string `json:"name"`
}
type OwnerInfo struct {
	Name string `json:"login"`
	}
type RepoResponse struct {
	Fullname   string `json:"fullname"`
	URL        string `json:"html_url"`
	Fork       bool   `json:"fork"`
	Update     string `json:"pushed_at"`
	Lang       string `json:"language"`
	Homepage   string `json:"homepage"`
	Message    string `json:"message"`
	License    LicenseInfo `json:"license"`
	Owner      OwnerInfo `json:"owner"`
	Pages      bool   `json:"has_pages"`
	Branch     string `json:"default_branch"`
	Description string `json:"description"`
	Star       int    `json:"stargazers_count"`
}
const url string = "https://api.github.com/"

func getAPI(url2 string) []byte {
        resp, _ := http.Get(url + url2)
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        return body
}
func getInfoRepo(fullname string){
	var data RepoResponse
	body := getAPI("repos/" +fullname)
	if err := json.Unmarshal(body,&data); err != nil {
		log.Fatal(err)
	}
	if data.Message == "Not Found"{
		color.Red("Repo '"+fullname+"' Not Found")
		os.Exit(1)
	}
	fmt.Println(data.Owner)
}
func getInfoUser(username string){
        var data UserResponse
        body := getAPI("users/" + username)
        if err := json.Unmarshal(body, &data); err != nil {
                log.Fatal(err)
        }
	if data.Message == "Not Found"{
		color.Red("user ' "+ username +"' Not Found.")
		os.Exit(1)
		}else{
        fmt.Println("user '"+ username+"' found.")
        datalist := color.YellowString("\nUsername: ")+data.Username+color.HiMagentaString("\nProfile: ")+data.Bio+color.HiBlueString("\nFollower: ")+strconv.Itoa(data.Followers)+color.HiCyanString("\tFollowing: ")+strconv.Itoa(data.Following)+color.HiGreenString("\nCompany: ")+data.Company+color.MagentaString("\nLocation: ")+data.Location+color.GreenString("\nBlog: ")+data.Blog+color.HiBlueString("\tTwitter: ")+data.Twitter+color.CyanString("\nRepos: ")+strconv.Itoa(data.Repos)
        fmt.Println(datalist)
	}
}

func main() {
        color.HiGreen("Github Info")
        args := os.Args
	if len(args) != 1{
        if args[2] == "--user"{
                getInfoUser(args[1])
                }
}
	getInfoRepo("roistaff/go-mem-api")
}
