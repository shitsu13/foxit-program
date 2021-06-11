package dao

import "go-scraper/src/dao/repository"

var (
	dao repository.Dao
)

func init() {
	dao = newSingleDao("dcard")
}

/*
	get instance
*/
func GetDao() repository.Dao {
	return dao
}
