package service

import "poker/model"

func ReadPokerFile(path string)(*model.PokerFile, error) {
	pf := model.PokerFile{}
	err := pf.GetFile(path)
	return &pf, err
}