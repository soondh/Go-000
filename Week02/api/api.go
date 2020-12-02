package api

import (
	"Week002/model"
	"Week002/service"

)

//查询用户资料
func QueryVideoInfoById(id int) (*model.VideoList, error){
	return service.QueryVideoInfoById(id)
}