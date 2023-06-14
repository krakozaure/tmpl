package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

var inputsDir = "."

func getFuncMap() template.FuncMap {
	f := sprig.GenericFuncMap()

	f["fromInputDir"] = fromInputDir
	f["include"] = include

	f["toBool"] = toBool
	f["toToml"] = toToml
	f["toYaml"] = toYaml

	f["absPath"] = absPath
	f["fileExists"] = fileExists
	f["fileMode"] = fileMode
	f["fileMtime"] = fileMtime
	f["fileRead"] = fileRead
	f["fileSize"] = fileSize
	f["isDir"] = isDir
	f["isFile"] = isFile

	return f
}

// --- Includes ------------------------------------------------------------------------------------

func include(input string) (string, error) {
	var err error
	includeDir := inputsDir
	if !filepath.IsAbs(input) {
		includeDir, err = getIncludeDir(input)
		if err != nil {
			return "", err
		}
	}
	input = filepath.Join(includeDir, input)

	outputString, err := executeTemplateFile(input)
	if err != nil {
		return "", err
	}
	return outputString, nil
}

func fromInputDir(input string) (string, error) {
	dir, err := getIncludeDir(input)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, input), nil
}

func getIncludeDir(input string) (string, error) {
	if input == "-" {
		cwd, err := filepath.Abs(".")
		if err != nil {
			return "", err
		}
		return cwd, nil
	} else {
		return inputsDir, nil
	}
}

// --- Type conversion -----------------------------------------------------------------------------

func toBool(value string) (bool, error) {
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return result, nil
}

func toToml(v interface{}) (string, error) {
	buf := bytes.NewBuffer(nil)
	enc := toml.NewEncoder(buf)
	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func toYaml(v interface{}) (string, error) {
	data, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// --- Paths ---------------------------------------------------------------------------------------

func absPath(file string) (string, error) {
	new_file, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return new_file, nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func fileMode(path string) (string, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}
	return info.Mode().String(), nil
}

func fileMtime(file string) (string, error) {
	info, err := os.Stat(file)
	if err != nil {
		return "", err
	}
	return info.ModTime().String(), nil
}

func fileRead(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func fileSize(file string) (int64, error) {
	info, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func isDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	return info.IsDir(), nil
}

func isFile(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	return info.Mode().IsRegular(), nil
}
