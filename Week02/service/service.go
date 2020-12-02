package service

import (
	"Week002/model"
	"Week002/dao"
)

func QueryVideoInfoById(id int) (*model.VideoList, error) {
	return dao.GetVideoById(id)
}