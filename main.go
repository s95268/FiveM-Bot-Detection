package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)


type Player struct {
	Identifiers []string `json:"identifiers"`
}

type ServerData struct {
	Data struct {
		Players []Player `json:"players"`
	} `json:"Data"`
}

func hexToSteam64(hexID string) string {
	steam64id, _ := strconv.ParseInt(hexID, 16, 64)
	return strconv.FormatInt(steam64id, 10)
}

func getSteamProfile(steam64ID string, apiKey string) (string, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", apiKey, steam64ID)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	playersData, ok := data["response"].(map[string]interface{})["players"].([]interface{})
	if !ok || len(playersData) == 0 {
		return "", fmt.Errorf("Any player found")
	}

	player, ok := playersData[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Invalid format of player data")
	}

	personaName, ok := player["personaname"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid format of persona name")
	}

	return personaName, nil
}


func main() {
	cfxCode := flag.String("code", "", "CFX Code to check")
	flag.Parse()

	if *cfxCode == "" {
		fmt.Println("Please provide a CFX Code to check.")
		os.Exit(1)
	}

	response, err := http.Get("https://servers-frontend.fivem.net/api/servers/single/" + *cfxCode)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var serverData ServerData
	err = json.Unmarshal(body, &serverData)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	botScore := 0
	for _, player := range serverData.Data.Players {
		for _, identifier := range player.Identifiers {
			if strings.HasPrefix(identifier, "steam:") {
				steamHex := strings.Split(identifier, ":")[1]
				personaName, err := getSteamProfile(hexToSteam64(steamHex), "STEAM_API_KEY")
				if err != nil || personaName == "" {
					botScore++
				}
			}
		}
	}

	fmt.Println("Bot's Score:", botScore)
}
