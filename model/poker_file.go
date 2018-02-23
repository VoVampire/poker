package model

import (
	"io/ioutil"
	"fmt"
	"github.com/json-iterator/go"
)

type PokerFile struct {
	File map[string][]map[string]interface{}
}

func (pf *PokerFile) GetFile(path string) (err error) {
	if file, err := ioutil.ReadFile(path); err != nil {
		fmt.Println(err.Error())

	} else if err = jsoniter.Unmarshal(file, &pf.File); err != nil {
		fmt.Println(err.Error())
	}
	return
}
