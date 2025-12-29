package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

const (
	Blue   = "\033[94m"
	Red    = "\033[91m"
	Cyan   = "\033[96m"
	White  = "\033[97m"
	Bold   = "\033[1m"
	Reset  = "\033[0m"
	Clear  = "\033[H\033[2J"
)

type LangStrings struct {
	EnterToken  string
	MenuClone   string
	MenuExit    string
	SourceID    string
	TargetID    string
	Starting    string
	Finished    string
	CreatedBy   string
	Roles       string
	Cats        string
	Channels    string
}

type Role struct {
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Hoist       bool   `json:"hoist"`
	Permissions int    `json:"permissions"`
	Mentionable bool   `json:"mentionable"`
}

type Channel struct {
	ID       string `json:"id"`
	Type     int    `json:"type"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	ParentID string `json:"parent_id"`
}

var messages = map[int]LangStrings{
	1: {
		EnterToken: "Enter your Token: ",
		MenuClone:  "Clone Server",
		MenuExit:   "Exit",
		SourceID:   "Source ID: ",
		TargetID:   "Target ID: ",
		Starting:   "Starting Replicator...",
		Finished:   "Cloning Finished Successfully!",
		CreatedBy:  "CREATED BY EMDAL",
		Roles:      "Replicating Roles...",
		Cats:       "Replicating Categories...",
		Channels:   "Replicating Channels...",
	},
	2: {
		EnterToken: "Ingresa tu Token: ",
		MenuClone:  "Clonar Servidor",
		MenuExit:   "Salir",
		SourceID:   "ID Origen: ",
		TargetID:   "ID Destino: ",
		Starting:   "Iniciando Replicador...",
		Finished:   "¡Clonación Finalizada con Éxito!",
		CreatedBy:  "CREADO POR EMDAL",
		Roles:      "Replicando Roles...",
		Cats:       "Replicando Categorías...",
		Channels:   "Replicando Canales...",
	},
	3: {
		EnterToken: "أدخل الرمز الخاص بك: ",
		MenuClone:  "استنساخ الخادم",
		MenuExit:   "خروج",
		SourceID:   "معرف المصدر: ",
		TargetID:   "معرف الهدف: ",
		Starting:   "بدء الاستنساخ...",
		Finished:   "تم الاستنساخ بنجاح!",
		CreatedBy:  "تم إنشاؤه بواسطة EMDAL",
		Roles:      "تكرار الأدوار...",
		Cats:       "تكرار الفئات...",
		Channels:   "تكرار القنوات...",
	},
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print(Clear)
	}
}

func main() {
	ClearScreen()
	
	fmt.Println("\n\n")
	fmt.Println(Red + Bold + "       ██████╗  ██████╗ ██╗      █████╗ ██╗███╗   ██╗ ██████╗ " + Reset)
	fmt.Println(Red + Bold + "      ██╔════╝ ██╔═══██╗██║     ██╔══██╗██║████╗  ██║██╔════╝ " + Reset)
	fmt.Println(Red + Bold + "      ██║  ███╗██║   ██║██║     ███████║██║██╔██╗ ██║██║  ███╗" + Reset)
	fmt.Println(Red + Bold + "      ██║   ██║██║   ██║██║     ██╔══██║██║██║╚██╗██║██║   ██║" + Reset)
	fmt.Println(Red + Bold + "      ╚██████╔╝╚██████╔╝███████╗██║  ██║██║██║ ╚████║╚██████╔╝" + Reset)
	fmt.Println(Red + Bold + "       ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ " + Reset)
	fmt.Println(Blue +Bold + "              ███████╗██████╗ ██████╗ ███████╗███╗   ██╗      " + Reset)
	fmt.Println(Blue +Bold + "              ██╔════╝██╔══██╗██╔══██╗██╔════╝████╗  ██║      " + Reset)
	fmt.Println(Blue +Bold + "              █████╗  ██████╔╝██║  ██║█████╗  ██╔██╗ ██║      " + Reset)
	fmt.Println(Blue +Bold + "              ██╔══╝  ██╔══██╗██║  ██║██╔══╝  ██║╚██╗██║      " + Reset)
	fmt.Println(Blue +Bold + "              ███████╗██║  ██║██████╔╝███████╗██║ ╚████║      " + Reset)
	fmt.Println(Blue +Bold + "              ╚══════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═══╝      " + Reset)
	fmt.Println("\n")

	fmt.Println(White + "              [ SELECT YOUR LANGUAGE ]" + Reset)
	fmt.Println(Cyan + "    ┌──────────────────────────────────────────────┐")
	fmt.Println("    │                                              │")
	fmt.Println("    │    " + Red + "1. English" + Cyan + "   " + Blue + "2. Español" + Cyan + "   " + White + "3. العربية" + Cyan + "    │")
	fmt.Println("    │                                              │")
	fmt.Println("    └──────────────────────────────────────────────┘" + Reset)
	
	var langID int
	fmt.Print("\n" + Red + "                     > " + Reset)
	fmt.Scanln(&langID)

	msg, ok := messages[langID]
	if !ok {
		msg = messages[1]
	}

	ClearScreen()
	
	fmt.Printf(Blue + Bold + `
   ______      __      _                                     __            
  / ____/___  / /___ _(_)___  ____ ____  ________  _________/ /__  ____  
 / / __/ __ \/ / __ ` + "`" + `/ / __ \/ __ ` + "`" + `/ _ \/ ___/ __ \/ ___/ __  / _ \/ __ \ 
/ /_/ / /_/ / / /_/ / / / / / /_/ /  __/ /  / /_/ / /  / /_/ /  __/ / / / 
\____/\____/_/\__,_/_/_/ /_/\__, /\___/_/   \____/_/   \__,_/\___/_/ /_/  
                           /____/                                         
` + Red + `
                     ` + msg.CreatedBy + `
` + Reset)

	var token string
	fmt.Printf("\n" + Blue + "[" + Red + "!" + Blue + "] " + msg.EnterToken + Reset)
	fmt.Scanln(&token)

	for {
		fmt.Printf("\n" + Blue + "[1] " + msg.MenuClone + "\n")
		fmt.Printf(Blue + "[2] " + msg.MenuExit + "\n" + Reset)
		
		var op int
		fmt.Printf(Red + "> " + Reset)
		fmt.Scanln(&op)

		if op == 2 {
			break
		}

		if op == 1 {
			var src, dst string
			fmt.Printf(Blue + msg.SourceID + Reset)
			fmt.Scanln(&src)
			fmt.Printf(Blue + msg.TargetID + Reset)
			fmt.Scanln(&dst)
			
			fmt.Printf("\n" + Red + "[*] " + msg.Starting + "\n" + Reset)
			cloneServer(token, src, dst, msg)
		}
	}
}

func cloneServer(token, source, target string, msg LangStrings) {
	roles := getRoles(token, source)
	channels := getChannels(token, source)

	var wg sync.WaitGroup

	for _, r := range roles {
		if r.Name != "@everyone" {
			wg.Add(1)
			go func(role Role) {
				defer wg.Done()
				createRole(token, target, role)
			}(r)
		}
	}
	wg.Wait()

	catMap := make(map[string]string)
	var catMutex sync.Mutex
	for _, c := range channels {
		if c.Type == 4 {
			newID := createChannel(token, target, c, "")
			catMutex.Lock()
			catMap[c.ID] = newID
			catMutex.Unlock()
		}
	}

	for _, c := range channels {
		if c.Type != 4 {
			wg.Add(1)
			go func(ch Channel) {
				defer wg.Done()
				parent := ""
				catMutex.Lock()
				if ch.ParentID != "" {
					parent = catMap[ch.ParentID]
				}
				catMutex.Unlock()
				createChannel(token, target, ch, parent)
			}(c)
		}
	}
	wg.Wait()

	fmt.Printf("\n" + Red + "[!] " + msg.Finished + "\n" + Reset)
}

func getRoles(token, guildID string) []Role {
	req, _ := http.NewRequest("GET", "https://discord.com/api/v9/guilds/"+guildID+"/roles", nil)
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return nil }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var roles []Role
	json.Unmarshal(body, &roles)
	return roles
}

func createRole(token, guildID string, r Role) {
	data, _ := json.Marshal(r)
	req, _ := http.NewRequest("POST", "https://discord.com/api/v9/guilds/"+guildID+"/roles", bytes.NewBuffer(data))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)
	fmt.Printf(Red + "[+] Role: " + Reset + r.Name + "\n")
}

func getChannels(token, guildID string) []Channel {
	req, _ := http.NewRequest("GET", "https://discord.com/api/v9/guilds/"+guildID+"/channels", nil)
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return nil }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var channels []Channel
	json.Unmarshal(body, &channels)
	return channels
}

func createChannel(token, guildID string, c Channel, parentID string) string {
	payload := map[string]interface{}{
		"name":     c.Name,
		"type":     c.Type,
		"position": c.Position,
	}
	if parentID != "" {
		payload["parent_id"] = parentID
	}
	data, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "https://discord.com/api/v9/guilds/"+guildID+"/channels", bytes.NewBuffer(data))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return "" }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	fmt.Printf(Blue + "[+] Channel: " + Reset + c.Name + "\n")
	if id, ok := result["id"].(string); ok {
		return id
	}
	return ""
}
