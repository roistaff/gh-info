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
        Email      string `json:"email"`
        Location   string `json:"location"`
        Blog       string `json:"blog"`
        Bio        string `json:"bio"`
        Twitter    string `json:"twitter_username"`
        Repos      int    `json:"public_repos"`
        Followers  int    `json:"followers"`
        Following  int    `json:"following"`
        Join       string `json:"created_at"`
        Username   string `json:"login"`
}

const url string = "https://api.github.com/"

func getAPI(url2 string) []byte {
        resp, _ := http.Get(url + url2)
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        return body
}

func getInfoUser(username string){
        var data UserResponse
        body := getAPI("users/" + username)
        if err := json.Unmarshal(body, &data); err != nil {
                log.Fatal(err)
        }
        fmt.Println("user:"+ username+" found.")
        datalist := color.YellowString("\nUsername")+": "+data.Username+color.HiMagentaString("\nProfile")+": "+data.Bio+color.HiBlueString("\nFollower")+": "+strconv.Itoa(data.Followers)+color.HiCyanString("\tFollowing")+": "+strconv.Itoa(data.Following)
        fmt.Println(datalist)

}
func main() {
        color.HiGreen("Github Info")
        args := os.Args[2:]
        if args[1] == "--user"{
                getInfoUser(args[0])
                }
}
