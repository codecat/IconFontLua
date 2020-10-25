package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/codecat/go-libs/log"
	"gopkg.in/yaml.v2"
)

type icon struct {
	ID      string
	Name    string
	Unicode string
}

type response struct {
	Icons []icon
}

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/ForkAwesome/Fork-Awesome/master/src/icons/icons.yml")
	if err != nil {
		log.Error("Unable to get icons.yaml: %s", err.Error())
		return
	}

	bytesIcons, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("Unable to read icons.yaml stream: %s", err.Error())
		return
	}

	resIcons := response{
		Icons: make([]icon, 0),
	}
	yaml.Unmarshal(bytesIcons, &resIcons)

	f, err := os.Create("fork-awesome.lua")
	if err != nil {
		log.Error("Unable to create output file: %s", err.Error())
		return
	}

	f.WriteString("-- Generated using https://github.com/codecat/IconFontLua\n")
	f.WriteString("local utf8 = require('utf8')\n")
	f.WriteString("return {\n")

	for _, icon := range resIcons.Icons {
		id := strings.Replace(icon.ID, "-", "_", -1)
		fmt.Fprintf(f, "\t['%s'] = utf8.char(0x%s), -- %s\n", id, icon.Unicode, icon.Name)
	}

	f.WriteString("}\n")
}
