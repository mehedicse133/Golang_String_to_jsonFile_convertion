package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func redText() []string {
	str, err := ioutil.ReadFile("raw.txt")
	fmt.Println(err)
	var all_str string
	all_str = string(str)
	split_str_list := strings.Split(all_str, "\n")
	return split_str_list
}

func main() {
	str := redText()

	var tech []string
	var registry []string
	var registrar []string
	var registrant []string
	var admin []string
	var domain []string
	var date []string
	var nameServer []string
	var dnssec []string

	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i], "Tech") {
			tech = append(tech, str[i])
		}
		if strings.HasPrefix(str[i], "Registry") {
			registry = append(registry, str[i])
		}

		if strings.HasPrefix(str[i], "Registrar") {
			registrar = append(registrar, str[i])
		}

		if strings.HasPrefix(str[i], "Registrant") {
			registrant = append(registrant, str[i])

		}
		if strings.HasPrefix(str[i], "Admin") {
			admin = append(admin, str[i])
		}

		if strings.HasPrefix(str[i], "Domain") {
			domain = append(domain, str[i])
		}

		if strings.HasPrefix(str[i], "Updated Date") {
			date = append(date, str[i])
		}

		if strings.HasPrefix(str[i], "Creation Date") {
			date = append(date, str[i])
		}

		if strings.Contains(str[i], "Name Server") {
			nameServer = append(nameServer, str[i])

		}
		if strings.HasPrefix(str[i], "DNSSEC") {
			dnssec = append(dnssec, str[i])
		}

	}

	var name_Server = make(map[string]string)
	for i := 0; i < len(nameServer); i++ {

		splitString := strings.SplitN(nameServer[i], ":", 2)

		key0 := strings.ReplaceAll(splitString[0], " ", "")
		key := strings.ToLower(key0)
		value := strings.ReplaceAll(splitString[1], "\r", "")
		name_Server[key+strconv.Itoa(1)] = value

		if i == 1 {
			splitString := strings.SplitN(nameServer[i], ":", 2)

			key0 := strings.ReplaceAll(splitString[0], " ", "")
			key := strings.ToLower(key0)
			value := strings.ReplaceAll(splitString[1], "\r", "")
			name_Server[key+strconv.Itoa(2)] = value
		}

	}

	var domainInfo = make(map[string]string)
	for i := 0; i < len(domain); i++ {
		splitString := strings.SplitN(domain[i], ":", 2)
		takeKey := strings.Split(splitString[0], " ")
		key := strings.ToLower(takeKey[1])

		value := strings.ReplaceAll(splitString[1], "\r", "")
		domainInfo[key] = value
	}

	for i := 0; i < len(date); i++ {
		splitString := strings.SplitN(date[i], ":", 2)

		key := strings.ReplaceAll(splitString[0], " ", "")
		value := strings.ReplaceAll(splitString[1], "\r", "")
		domainInfo[key] = value
	}

	var registrarInfo = make(map[string]string)
	for i := 0; i < len(registrar); i++ {
		splitString := strings.SplitN(registrar[i], ":", 2)

		key := strings.ReplaceAll(splitString[0], " ", "")
		value := strings.ReplaceAll(splitString[1], "\r", "")
		registrarInfo[key] = value
	}

	var registrantInfo = make(map[string]string)
	for i := 0; i < len(registrant); i++ {
		splitString := strings.Split(registrant[i], ":")
		takeKey := strings.Split(splitString[0], " ")

		key0 := strings.Join(takeKey[1:], "")
		key := strings.ToLower(key0)

		value := strings.ReplaceAll(splitString[1], "\r", "")
		registrantInfo[key] = value
	}

	var adminInfo = make(map[string]string)
	for i := 0; i < len(admin); i++ {
		splitString := strings.Split(admin[i], ":")
		takeKey := strings.Split(splitString[0], " ")

		key0 := strings.Join(takeKey[1:], "")
		key := strings.ToLower(key0)
		value := strings.ReplaceAll(splitString[1], "\r", "")
		adminInfo[key] = value
	}

	var techInfo = make(map[string]string)
	for i := 0; i < len(tech); i++ {
		splitString := strings.Split(tech[i], ":")
		fmt.Println(splitString[1])
		takeKey := strings.Split(splitString[0], " ")

		key0 := strings.Join(takeKey[1:], "")
		key := strings.ToLower(key0)

		value := strings.ReplaceAll(splitString[1], "\r", "")
		techInfo[key] = value
	}
	techInfo["rawText"] = strings.Join(tech, ",")

	var registryInfo = map[string]string{}
	for i := 0; i < len(registry); i++ {
		splitString := strings.Split(registry[i], ":")
		takeKey := strings.Split(splitString[0], " ")

		key1 := takeKey[1]
		key2 := takeKey[2]
		value := strings.ReplaceAll(splitString[1], "\r", "")
		key0 := key1 + "" + strings.Title(key2)
		key := strings.ToLower(key0)

		registryInfo[key] = value

	}

	var dnsse = map[string]string{}
	for i := 0; i < len(dnssec); i++ {
		splitString := strings.Split(dnssec[i], ":")

		key0 := strings.ToLower(splitString[0])
		key := strings.ToLower(key0)

		value := strings.ReplaceAll(splitString[1], "\r", "")
		dnsse[key] = value
	}

	var data = map[string]map[string]string{}
	data["domain"] = domainInfo
	data["adminInformatioin"] = adminInfo
	data["techInformatioin"] = techInfo
	data["registryInformatioin"] = registryInfo
	data["registrerInformatioin"] = registrarInfo
	data["registrantInformatioin"] = registrantInfo
	data["nameServer"] = name_Server
	data["dnssec"] = dnsse

	byteData, err := json.Marshal(data)
	fmt.Println(err)
	jsondata := string(byteData)
	fmt.Println(string(jsondata))

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)

}
