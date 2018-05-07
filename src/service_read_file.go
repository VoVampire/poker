package src

import (
	"io/ioutil"
	"github.com/json-iterator/go"
)

type Matches struct {
	Matches []Match `json:"matches"`
}

type Match struct {
	PlayerA string `json:"alice"`
	PlayerB string `json:"bob"`
	Result  int    `json:"result"`
}

func GetMatchesFromMatchSamples(path string) (*Matches, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	matches := Matches{}
	err = jsoniter.Unmarshal(file, &matches)
	return &matches, err
}
