/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `This command fetches a random dadjoke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")

		if jokeTerm != "" {
			getSpecificRandomJoke(jokeTerm)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	randomCmd.Flags().String("term", "", "A search term for a dad joke")
}

// Joke struct
type Joke struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	err := json.Unmarshal(responseBytes, &joke)
	if err != nil {
		log.Println("Could not unmarshal reponseBytes. %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getSpecificRandomJoke(jokeTerm string) {
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", jokeTerm)
	responseBytes := getJokeData(url)
	fmt.Println(jokeTerm)
	fmt.Println(string(responseBytes))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Println("Could not request a dadjoke. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/monsteronfire/dadjoke-cli)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("Did not get a response. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Could not read response body. %v", err)
	}

	return responseBytes
}

// TODO
// Including a search term returns an array of jokes. Write a function to get a random selection from this array
