package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

var ctx = make(map[string]interface{})

func loadContext(input string) {
	ctx["__includeDir__"] = getIncludeDir(input)

	if UseEnv {
		ctx["Env"] = getEnvVariables()
	}

	for _, file := range FilesList {
		for k, v := range getFileVariables(file) {
			ctx[k] = v
		}
	}

	for k, v := range getCliVariables() {
		ctx[k] = v
	}
}

func getIncludeDir(input string) string {
	// If input is stdin,
	// template paths are relative to the current working directory
	// else relative to the input directory
	if input == "-" {
		cwd, err := os.Getwd()
		if err != nil {
			return "."
		} else {
			return cwd
		}
	} else {
		return filepath.Dir(input)
	}
}

func getEnvVariables() map[string]string {
	vars := make(map[string]string)
	for _, env := range os.Environ() {
		kv := strings.SplitN(env, "=", 2)
		vars[kv[0]] = kv[1]
	}
	return vars
}

func getFileVariables(file string) map[string]interface{} {
	vars := make(map[string]interface{})

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Errorf("unable to read file\n%v\n", err))
		return vars
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
		fmt.Fprint(os.Stderr, fmt.Errorf("unable to load data\n%v\n", err))
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
