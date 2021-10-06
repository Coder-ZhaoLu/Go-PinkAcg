package mysql

import (
	"database/sql"
	"web_app/models"
)

func GetStarPostsById(pid int64, uid int64) (stars []*models.Star, err error) {
	sqlStr := `select user_id, post_id from stars where user_id=? and post_id=?`
	err = db.Select(&stars, sqlStr, uid, pid)
	return
}

func GetStarUserById(uid int64) (follows []*models.Star, err error) {
	sqlStr := `select user_id, post_id from stars where user_id=?`
	err = db.Select(&follows, sqlStr, uid)
	return
}

// StarPost 喜欢文章
func StarPost(pid int64, uid int64) (res sql.Result, err error) {
	// 获取喜欢的文章数量
	var starsNum []*models.Star
	sqlStr2 := `select user_id, post_id from stars where post_id=?`
	err = db.Select(&starsNum, sqlStr2, pid)

	// 开启事务
	tx, err := db.Begin()
	sqlStr := `insert into stars (user_id, post_id) values(?,?)`
	res, err = tx.Exec(sqlStr, uid, pid)
	if err != nil {
		err = tx.Rollback()
	}

	sqlStr3 := `update post set favorite=? where post_id=?`
	res, err = tx.Exec(sqlStr3, len(starsNum)+1, pid)
	if err != nil {
		err = tx.Rollback()
	}

	// 提交事务
	err = tx.Commit()
	return
}

// UnStarPost 喜欢文章
// 0代表无状态，1代表喜欢，2代表不喜欢
func UnStarPost(pid int64, uid int64) (res sql.Result, err error) {
	// 获取喜欢的文章数量
	var unStarsNum []*models.Star
	sqlStr2 := `select user_id, post_id from stars where post_id=?`
	err = db.Select(&unStarsNum, sqlStr2, pid)

	// 开启事务
	tx, err := db.Begin()
	sqlStr := `delete from stars where post_id=? and user_id=?`
	res, err = tx.Exec(sqlStr, pid, uid)
	if err != nil {
		err = tx.Rollback()
	}

	sqlStr3 := `update post set favorite=? where post_id=?`
	res, err = tx.Exec(sqlStr3, len(unStarsNum)-1, pid)
	if err != nil {
		err = tx.Rollback()
	}

	// 提交事务
	err = tx.Commit()
	return
}
