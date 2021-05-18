package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	appQuestionsLotteriesFieldNames          = builderx.RawFieldNames(&AppQuestionsLotteries{})
	appQuestionsLotteriesRows                = strings.Join(appQuestionsLotteriesFieldNames, ",")
	appQuestionsLotteriesRowsExpectAutoSet   = strings.Join(stringx.Remove(appQuestionsLotteriesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	appQuestionsLotteriesRowsWithPlaceHolder = strings.Join(stringx.Remove(appQuestionsLotteriesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppQuestionsLotteriesIdPrefix = "cache#appQuestionsLotteries#id#"
)

type (
	AppQuestionsLotteriesModel interface {
		Insert(data AppQuestionsLotteries) (sql.Result, error)
		FindOne(id int64) (*AppQuestionsLotteries, error)
		Update(data AppQuestionsLotteries) error
		Delete(id int64) error
	}

	defaultAppQuestionsLotteriesModel struct {
		sqlc.CachedConn
		table string
	}

	AppQuestionsLotteries struct {
		Id         int64         `db:"id"`
		Beid       int64         `db:"beid"`  // 对应的平台
		Ptyid      int64         `db:"ptyid"` // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Uid        int64         `db:"uid"`   // 中台表用户的id
		Auid       sql.NullInt64 `db:"auid"`
		ActivityId int64         `db:"activity_id"`
		AnswerId   int64         `db:"answer_id"`
		AwardId    int64         `db:"award_id"`
		IsWinning  int64         `db:"is_winning"` // 是否中奖（0/1）
		IsConvert  int64         `db:"is_convert"` // 兑奖名称
		CreateTime time.Time     `db:"create_time"`
		UpdateTime time.Time     `db:"update_time"`
	}
)

func NewAppQuestionsLotteriesModel(conn sqlx.SqlConn, c cache.CacheConf) AppQuestionsLotteriesModel {
	return &defaultAppQuestionsLotteriesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_questions_lotteries`",
	}
}

func (m *defaultAppQuestionsLotteriesModel) Insert(data AppQuestionsLotteries) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appQuestionsLotteriesRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Beid, data.Ptyid, data.Uid, data.Auid, data.ActivityId, data.AnswerId, data.AwardId, data.IsWinning, data.IsConvert)

	return ret, err
}

func (m *defaultAppQuestionsLotteriesModel) FindOne(id int64) (*AppQuestionsLotteries, error) {
	appQuestionsLotteriesIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsLotteriesIdPrefix, id)
	var resp AppQuestionsLotteries
	err := m.QueryRow(&resp, appQuestionsLotteriesIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsLotteriesRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAppQuestionsLotteriesModel) Update(data AppQuestionsLotteries) error {
	appQuestionsLotteriesIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsLotteriesIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appQuestionsLotteriesRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Uid, data.Auid, data.ActivityId, data.AnswerId, data.AwardId, data.IsWinning, data.IsConvert, data.Id)
	}, appQuestionsLotteriesIdKey)
	return err
}

func (m *defaultAppQuestionsLotteriesModel) Delete(id int64) error {

	appQuestionsLotteriesIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsLotteriesIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, appQuestionsLotteriesIdKey)
	return err
}

func (m *defaultAppQuestionsLotteriesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppQuestionsLotteriesIdPrefix, primary)
}

func (m *defaultAppQuestionsLotteriesModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsLotteriesRows, m.table)
	return conn.QueryRow(v, query, primary)
}
