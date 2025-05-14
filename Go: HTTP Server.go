package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Lake struct for storing lake information
type Lake struct {
	Name string
	Area int32
}

func main() {
	// Initialize store
	store := make(map[string]Lake)

	scanner := bufio.NewScanner(os.Stdin)

	// Read number of actions
	var n int
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &n)
	}

	// Process each action
	for i := 0; i < n; i++ {
		// Read the action string
		var actionStr string
		if scanner.Scan() {
			actionStr = scanner.Text()
		}

		// Parse the action JSON
		var actionMap map[string]interface{}
		err := json.Unmarshal([]byte(actionStr), &actionMap)
		if err != nil {
			continue
		}

		// Get action type and payload
		actionType, _ := actionMap["type"].(string)
		payload := actionMap["payload"]

		switch actionType {
		case "post":
			// Handle post action
			payloadStr, ok := payload.(string)
			if !ok {
				// Try to handle if payload is already a map
				payloadMap, mapOk := payload.(map[string]interface{})
				if mapOk {
					id, _ := payloadMap["id"].(string)
					name, _ := payloadMap["name"].(string)
					areaFloat, _ := payloadMap["area"].(float64)

					store[id] = Lake{
						Name: name,
						Area: int32(areaFloat),
					}
				}
				continue
			}

			// If payload is backslash-escaped JSON string
			if strings.HasPrefix(payloadStr, "{\\\"") {
				payloadStr = strings.ReplaceAll(payloadStr, "\\\"", "\"")
				payloadStr = strings.ReplaceAll(payloadStr, "\\{", "{")
				payloadStr = strings.ReplaceAll(payloadStr, "\\}", "}")
			}

			var lakeData map[string]interface{}
			err = json.Unmarshal([]byte(payloadStr), &lakeData)
			if err != nil {
				continue
			}

			id, _ := lakeData["id"].(string)
			name, _ := lakeData["name"].(string)
			areaFloat, _ := lakeData["area"].(float64)

			store[id] = Lake{
				Name: name,
				Area: int32(areaFloat),
			}

		case "delete":
			// Handle delete action
			id, _ := payload.(string)

			_, exists := store[id]
			if !exists {
				fmt.Println("404 Not Found")
				continue
			}

			delete(store, id)

		case "get":
			// Handle get action
			id, _ := payload.(string)

			lake, exists := store[id]
			if !exists {
				fmt.Println("404 Not Found")
				continue
			}

			fmt.Println(lake.Name)
			fmt.Println(lake.Area)
		}
	}
}
