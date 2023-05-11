package main

import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "strconv"
        "github.com/fatih/color"
        "net/http"
	"io/ioutil"
        "strings"
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
	Fullname   string `json:"full_name"`
	URL        string `json:"html_url"`
	Fork       bool   `json:"fork"`
	Update     string `json:"pushed_at"`
	Lang       string `json:"language"`
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
	datalist := color.GreenString("\n"+data.Fullname)+"\t\t"+color.YellowString("☆Star:")+strconv.Itoa(data.Star)+color.HiMagentaString("\nDescription: ")+data.Description+color.HiCyanString("\nLanguage: ")+data.Lang+color.BlueString("\nLicense: ")+data.License.Name+color.MagentaString("\nURL: ")+data.URL
	fmt.Println(datalist)
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
func welcome(){
        welcometext := `         ____  _  _    _             _       ___          __        
        / ___|(_)| |_ | |__   _   _ | |__   |_ _| _ __   / _|  ___  
       | |  _ | || __|| '_ \ | | | || '_ \   | | | '_ \ | |_  / _ \ 
       | |_| || || |_ | | | || |_| || |_) |  | | | | | ||  _|| (_) |
        \____||_| \__||_| |_| \__,_||_.__/  |___||_| |_||_|   \___/ 
                                                                    `
        fmt.Println(welcometext)
}
func Helps(message ...string){
        if message != nil{
                fmt.Println(color.RedString("! Message: "),message)
        }
        fmt.Println("Hello")
}
func main() {
        color.HiGreen("Github Info")
        args := os.Args
	if len(args) != 1{if len(args) == 3{
        if args[2] == "--user"{
                getInfoUser(args[1])
                }else if args[2] == "--repo"{
			getInfoRepo(args[1])
			}
                }else if len(args) == 2{
                        if strings.Contains(args[1], "/"){
                                getInfoRepo(args[1])
                        }else{
                                getInfoUser(args[1])
                        }
                }
        }else{
                welcome()

        }
}
