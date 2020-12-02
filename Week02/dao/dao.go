package dao

import (
	"Week002/model"
	"database/sql"

	pkgerr "github.com/pkg/errors"
)

var (
	db *sql.DB
)

func GetVideoById(id int) (*model.VideoList, error) {

	//模拟db query发生ErrNoRows异常，其他异常同理
	if id == 0 {
		return nil, pkgerr.Wrap(sql.ErrNoRows, "[FAILED] video id not found in db query")

		//如果跟下面一样不wrap，堆栈信息就丢了
		//return nil, sql.ErrNoRows
	} else {
		video := model.VideoList{
			1,
			"马保国",
		}
		return &video, nil
	}
}