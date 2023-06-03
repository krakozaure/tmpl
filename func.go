package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v3"
)

func getFuncMap() template.FuncMap {
	f := sprig.GenericFuncMap()

	f["include"] = includeTemplate

	f["toBool"] = toBool
	f["toToml"] = toToml
	f["toYaml"] = toYaml

	f["absPath"] = absPath
	f["isFile"] = isFile
	f["isDir"] = isDir
	f["fileExists"] = fileExists
	f["fileMode"] = fileMode
	f["fileSize"] = fileSize
	f["fileMTime"] = fileMTime
	f["fileRead"] = fileRead

	return f
}

func includeTemplate(input string) string {
	if !filepath.IsAbs(input) {
		srcFileDir, ok := ctx["__includeDir__"].(string)
		if ok {
			input = filepath.Join(srcFileDir, input)
		}
	}

	outputString, err := templateExecute(input)
	if err != nil {
		if Strict {
			panic(fmt.Errorf("unable to render included template\n%v\n", err))
		}
		return ""
	}
	return outputString
}

func toBool(value string) bool {
	// 0/1, f/t, F/T, FALSE/TRUE, False/True, false/true
	result, err := strconv.ParseBool(value)
	if err != nil {
		if Strict {
			panic(err.Error())
		}
		return false
	}
	return result
}

func toYaml(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		if Strict {
			panic(err.Error())
		}
		return ""
	}
	return string(data)
}

func toToml(v interface{}) string {
	buf := bytes.NewBuffer(nil)
	enc := toml.NewEncoder(buf)
	err := enc.Encode(v)
	if err != nil {
		if Strict {
			panic(err.Error())
		}
		return ""
	}
	return buf.String()
}

func absPath(file string) string {
	new_file, err := filepath.Abs(file)
	if err != nil {
		if Strict {
			panic(err.Error())
		}
		return ""
	}
	return new_file
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		if Strict {
			panic(err.Error())
		}
		return false
	}
	return info.IsDir()
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		if Strict {
			panic(err.Error())
		}
		return false
	}
	return info.Mode().IsRegular()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if Strict {
			panic(err.Error())
		}
		return false
	}
	return true
}

func fileMode(path string) string {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		if Strict {
			panic(err.Error())
		}
		return ""
	}
	return info.Mode().String()
}

func fileSize(file string) int64 {
	info, err := os.Stat(file)
	if err != nil {
		if Strict {
			panic(err.Error())
		}
		return 0
	}
	return info.Size()
}

func fileMTime(file string) string {
	info, err := os.Stat(file)
	if err != nil {
		if Strict {
			panic(err.Error())
		}
		return ""
	}
	return info.ModTime().String()
}

func fileRead(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		if Strict {
			panic(err.Error())
		}
		return ""
	}
	return string(data)
}
