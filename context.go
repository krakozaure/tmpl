package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

var ctx = make(map[string]interface{})

func loadContext() {
	for _, file := range FilesList {
		for k, v := range getFileVariables(file) {
			ctx[k] = v
		}
	}

	for k, v := range getCliVariables() {
		ctx[k] = v
	}
}

func getFileVariables(file string) map[string]interface{} {
	vars := make(map[string]interface{})

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file\n%v\n", err)
	}

	if strings.HasSuffix(file, ".json") {
		err = json.Unmarshal(bytes, &vars)
	} else if strings.HasSuffix(file, ".toml") {
		err = toml.Unmarshal(bytes, &vars)
	} else if strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml") {
		err = yaml.Unmarshal(bytes, &vars)
	} else {
		err = fmt.Errorf("bad file type: %s", file)
	}
	if err != nil {
		log.Fatalf("unable to load data\n%v\n", err)
	}
	return vars
}

func getCliVariables() map[string]string {
	vars := make(map[string]string)
	for _, pair := range VarsList {
		kv := strings.SplitN(pair, "=", 2)

		v := kv[1]
		if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
			v = v[1 : len(v)-1]
		} else if strings.HasPrefix(v, "'") && strings.HasSuffix(v, "'") {
			v = v[1 : len(v)-1]
		}

		vars[kv[0]] = v
	}
	return vars
}
