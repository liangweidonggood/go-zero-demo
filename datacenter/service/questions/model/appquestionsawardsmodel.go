package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	appQuestionsAwardsFieldNames          = builderx.RawFieldNames(&AppQuestionsAwards{})
	appQuestionsAwardsRows                = strings.Join(appQuestionsAwardsFieldNames, ",")
	appQuestionsAwardsRowsExpectAutoSet   = strings.Join(stringx.Remove(appQuestionsAwardsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	appQuestionsAwardsRowsWithPlaceHolder = strings.Join(stringx.Remove(appQuestionsAwardsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppQuestionsAwardsIdPrefix    = "cache#appQuestionsAwards#id#"
	cacheAppQuestionsAwardsTitlePrefix = "cache#appQuestionsAwards#title#"
)

type (
	AppQuestionsAwardsModel interface {
		Insert(data AppQuestionsAwards) (sql.Result, error)
		FindOne(id int64) (*AppQuestionsAwards, error)
		FindOneByTitle(title sql.NullString) (*AppQuestionsAwards, error)
		Update(data AppQuestionsAwards) error
		Delete(id int64) error
	}

	defaultAppQuestionsAwardsModel struct {
		sqlc.CachedConn
		table string
	}

	AppQuestionsAwards struct {
		Id               int64          `db:"id"`
		Beid             int64          `db:"beid"`  // 对应的平台
		Ptyid            int64          `db:"ptyid"` // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Title            sql.NullString `db:"title"` // 奖项名
		ActivityId       int64          `db:"activity_id"`
		StartProbability int64          `db:"start_probability"` // 开始概率
		EndProbability   int64          `db:"end_probability"`   // 结束概率
		Number           int64          `db:"number"`            // 中奖个数
		IsLottery        int64          `db:"is_lottery"`        // 是否属于中奖
		Header           string         `db:"header"`            // 分享标题
		Des              string         `db:"des"`               // 分享文本
		Image            sql.NullString `db:"image"`             // 分享图片
		CreatedAt        sql.NullTime   `db:"created_at"`
		UpdatedAt        sql.NullTime   `db:"updated_at"`
	}
)

func NewAppQuestionsAwardsModel(conn sqlx.SqlConn, c cache.CacheConf) AppQuestionsAwardsModel {
	return &defaultAppQuestionsAwardsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_questions_awards`",
	}
}

func (m *defaultAppQuestionsAwardsModel) Insert(data AppQuestionsAwards) (sql.Result, error) {
	appQuestionsAwardsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsAwardsTitlePrefix, data.Title)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appQuestionsAwardsRowsExpectAutoSet)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Title, data.ActivityId, data.StartProbability, data.EndProbability, data.Number, data.IsLottery, data.Header, data.Des, data.Image, data.CreatedAt, data.UpdatedAt)
	}, appQuestionsAwardsTitleKey)
	return ret, err
}

func (m *defaultAppQuestionsAwardsModel) FindOne(id int64) (*AppQuestionsAwards, error) {
	appQuestionsAwardsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsAwardsIdPrefix, id)
	var resp AppQuestionsAwards
	err := m.QueryRow(&resp, appQuestionsAwardsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsAwardsRows, m.table)
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

func (m *defaultAppQuestionsAwardsModel) FindOneByTitle(title sql.NullString) (*AppQuestionsAwards, error) {
	appQuestionsAwardsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsAwardsTitlePrefix, title)
	var resp AppQuestionsAwards
	err := m.QueryRowIndex(&resp, appQuestionsAwardsTitleKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `title` = ? limit 1", appQuestionsAwardsRows, m.table)
		if err := conn.QueryRow(&resp, query, title); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAppQuestionsAwardsModel) Update(data AppQuestionsAwards) error {
	appQuestionsAwardsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsAwardsIdPrefix, data.Id)
	appQuestionsAwardsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsAwardsTitlePrefix, data.Title)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appQuestionsAwardsRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Title, data.ActivityId, data.StartProbability, data.EndProbability, data.Number, data.IsLottery, data.Header, data.Des, data.Image, data.CreatedAt, data.UpdatedAt, data.Id)
	}, appQuestionsAwardsIdKey, appQuestionsAwardsTitleKey)
	return err
}

func (m *defaultAppQuestionsAwardsModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	appQuestionsAwardsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsAwardsTitlePrefix, data.Title)
	appQuestionsAwardsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsAwardsIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, appQuestionsAwardsIdKey, appQuestionsAwardsTitleKey)
	return err
}

func (m *defaultAppQuestionsAwardsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppQuestionsAwardsIdPrefix, primary)
}

func (m *defaultAppQuestionsAwardsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsAwardsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
