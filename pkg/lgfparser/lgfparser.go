package lgfparser

import (
	"bufio"
	"os"
	"strings"
)

type GeneralInfoType struct {
	Users      map[string]string
	Computers  map[string]string
	Apps       map[string]string
	Events     map[string]string
	Metadata   map[string]string
	Servers    map[string]string
	Main_ports map[string]string
	Sec_ports  map[string]string
}

func ParseLgf(lgfPath string) GeneralInfoType {

	var a GeneralInfoType

	a.Users = make(map[string]string)
	a.Computers = make(map[string]string)
	a.Apps = make(map[string]string)
	a.Events = make(map[string]string)
	a.Metadata = make(map[string]string)
	a.Servers = make(map[string]string)
	a.Main_ports = make(map[string]string)
	a.Sec_ports = make(map[string]string)

	lgf_file, err := os.Open(lgfPath)
	if err != nil {
		panic(err)
	}
	defer lgf_file.Close()

	lgf_scanner := bufio.NewScanner(lgf_file)
	for lgf_scanner.Scan() {

		NowString := lgf_scanner.Text()

		if len(NowString) > 2 {

			split_string := strings.Split(NowString[1:len(NowString)-2], ",")

			switch split_string[0] {
			case "1":
				a.Users[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			case "2":
				a.Computers[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			case "3":
				a.Apps[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			case "4":
				a.Events[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			case "5":
				a.Metadata[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			case "6":
				a.Servers[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			case "7":
				a.Main_ports[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			case "8":
				a.Sec_ports[split_string[len(split_string)-1]] = split_string[len(split_string)-2]
			}
		}
	}

	//if cfg.Config.Debug.PrintLgfMaps {
	//	fmt.Println(a)
	//}

	return a
}

func add_array_item(a, array_item []string) map[string]string {

	users := map[string]string{}

	if len(array_item) > 1 {
		switch array_item[0] {
		case "1":
			users[array_item[len(array_item)-1]] = array_item[len(array_item)-2]
		}
	}

	return users
}
