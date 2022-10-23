package blm

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Array map[string]string

type Reader struct {
	file       string
	content    string
	definition []string
	pathInfo   Array
	headers    Array
	data       Array
}

func NewReader(filepath string) (*Reader, error) {
	var err error

	r := new(Reader)
	if err = r.setFile(filepath); err != nil {
		return nil, err
	}

	if err = r.setPathInfo(); err != nil {
		return nil, err
	}

	if err = r.setContent(); err != nil {
		return nil, err
	}

	r.setHeader()
	r.setDefinitions()
	r.setData()

	return r, nil
}

func (r *Reader) GetPathInfo() Array {
	return r.pathInfo
}

func (r *Reader) GetHeaders() Array {
	return r.headers
}

func (r *Reader) GetDefinitions() []string {
	return r.definition
}

func (r *Reader) GetData() Array {
	return r.data
}

func (r *Reader) setFile(file string) error {
	if _, err := os.Stat(file); err != nil {
		return err
	}

	r.file = file

	return nil
}

func (r *Reader) setContent() error {
	content, err := os.ReadFile(r.file)
	if err != nil {
		return err
	}

	r.content = string(content)
	//src/Reader.php:64

	return nil
}

func (r *Reader) setPathInfo() error {
	r.pathInfo = make(map[string]string)
	r.pathInfo["extension"] = filepath.Ext(r.file)
	r.pathInfo["base"] = filepath.Base(r.file)
	r.pathInfo["dir"] = filepath.Dir(r.file)

	if strings.ToLower(r.pathInfo["extension"]) != ".blm" {
		return errors.New("Given file is not a .BLM")
	}

	return nil
}

func (r *Reader) setHeader() {
	match, _ := regexp.Compile("/#HEADER#(.*?)#/sm")
	params := make(map[string]string)
	for _, line := range match.FindStringSubmatch(r.content) {
		if strings.TrimSpace(line) != "" {
			parts := strings.Split(line, " : ")
			regex := regexp.MustCompile("/(^[\\'\"]|[\\'\"]$)/")
			params[strings.TrimSpace(parts[0])] = regex.ReplaceAllString(strings.TrimSpace(parts[1]), "")
		}

		fmt.Println(line)
	}
}

func (r *Reader) setDefinitions() {

}

func (r *Reader) setData() {

}
