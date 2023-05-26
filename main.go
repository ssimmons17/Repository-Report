package main 

import (

	"fmt"
	"log"
	"net/http"
	"io"
	"encoding/json"
	"encoding/csv"
	"os"
)



type Repository struct {

	Id int `json:"id"`
	Name string `json:"name"`
	HtmlUrl string `json:"html_url"`
	Description string `json:"description"`
}


func main () {


req, err := http.NewRequest("GET", "https://api.github.com/users/{USERNAME}/repos", nil)
if err != nil {

	log.Fatal(err)
}

req.SetBasicAuth(<USERNAME>, "<Personal Access Token",)

client := http.Client{}
res, err := client.Do(req)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Status code", res.StatusCode)

body, err := io.ReadAll(res.Body)
if err != nil {
	log.Fatal(err)
}


var repository []Repository

err1 := json.Unmarshal([]byte(body), &repository)
if err1 != nil {

	fmt.Println(err1)
}

fmt.Printf("%v\t %v\t %v \n", "ID", "Name", "Url")

for _, r := range repository {


	fmt.Printf("%d %v %v \n", r.Id, r.Name, r.HtmlUrl)

}

//csv headers
header := []string{"ID", "Name", "URL"}

csvFile, err := os.Create("GHRepo.csv")//create file at users download directory??
defer csvFile.Close()
if err != nil {
	
	fmt.Println(err)
}

w := csv.NewWriter(csvFile)
if err := w.Write(header); err != nil {
	fmt.Println(err)
}
w.Flush()

//write json to csv
for _, r := range repository {

	cells := []string{fmt.Sprint(r.Id), r.Name, r.HtmlUrl}
	err := w.Write(cells)
	if err != nil {

		fmt.Println(err)
	}

}

w.Flush()
if err := w.Error(); err != nil {

	fmt.Println(err)
}

fmt.Println("CSV creation complete!")

}

