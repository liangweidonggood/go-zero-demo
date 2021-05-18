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
	appQuestionsTestsFieldNames          = builderx.RawFieldNames(&AppQuestionsTests{})
	appQuestionsTestsRows                = strings.Join(appQuestionsTestsFieldNames, ",")
	appQuestionsTestsRowsExpectAutoSet   = strings.Join(stringx.Remove(appQuestionsTestsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	appQuestionsTestsRowsWithPlaceHolder = strings.Join(stringx.Remove(appQuestionsTestsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppQuestionsTestsIdPrefix    = "cache#appQuestionsTests#id#"
	cacheAppQuestionsTestsTitlePrefix = "cache#appQuestionsTests#title#"
)

type (
	AppQuestionsTestsModel interface {
		Insert(data AppQuestionsTests) (sql.Result, error)
		FindOne(id int64) (*AppQuestionsTests, error)
		FindOneByTitle(title string) (*AppQuestionsTests, error)
		Update(data AppQuestionsTests) error
		Delete(id int64) error
	}

	defaultAppQuestionsTestsModel struct {
		sqlc.CachedConn
		table string
	}

	AppQuestionsTests struct {
		Id          int64        `db:"id"`
		Beid        int64        `db:"beid"`         // 对应的平台
		Ptyid       int64        `db:"ptyid"`        // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Title       string       `db:"title"`        // 题库类名
		QuestionIds string       `db:"question_ids"` // 题库题编号
		Status      int64        `db:"status"`       // 状态（0/1）
		CreatedAt   sql.NullTime `db:"created_at"`
		UpdatedAt   sql.NullTime `db:"updated_at"`
	}
)

func NewAppQuestionsTestsModel(conn sqlx.SqlConn, c cache.CacheConf) AppQuestionsTestsModel {
	return &defaultAppQuestionsTestsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_questions_tests`",
	}
}

func (m *defaultAppQuestionsTestsModel) Insert(data AppQuestionsTests) (sql.Result, error) {
	appQuestionsTestsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsTestsTitlePrefix, data.Title)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, appQuestionsTestsRowsExpectAutoSet)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Title, data.QuestionIds, data.Status, data.CreatedAt, data.UpdatedAt)
	}, appQuestionsTestsTitleKey)
	return ret, err
}

func (m *defaultAppQuestionsTestsModel) FindOne(id int64) (*AppQuestionsTests, error) {
	appQuestionsTestsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsTestsIdPrefix, id)
	var resp AppQuestionsTests
	err := m.QueryRow(&resp, appQuestionsTestsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsTestsRows, m.table)
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

func (m *defaultAppQuestionsTestsModel) FindOneByTitle(title string) (*AppQuestionsTests, error) {
	appQuestionsTestsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsTestsTitlePrefix, title)
	var resp AppQuestionsTests
	err := m.QueryRowIndex(&resp, appQuestionsTestsTitleKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `title` = ? limit 1", appQuestionsTestsRows, m.table)
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

func (m *defaultAppQuestionsTestsModel) Update(data AppQuestionsTests) error {
	appQuestionsTestsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsTestsIdPrefix, data.Id)
	appQuestionsTestsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsTestsTitlePrefix, data.Title)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appQuestionsTestsRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Title, data.QuestionIds, data.Status, data.CreatedAt, data.UpdatedAt, data.Id)
	}, appQuestionsTestsTitleKey, appQuestionsTestsIdKey)
	return err
}

func (m *defaultAppQuestionsTestsModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	appQuestionsTestsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsTestsIdPrefix, id)
	appQuestionsTestsTitleKey := fmt.Sprintf("%s%v", cacheAppQuestionsTestsTitlePrefix, data.Title)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, appQuestionsTestsIdKey, appQuestionsTestsTitleKey)
	return err
}

func (m *defaultAppQuestionsTestsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppQuestionsTestsIdPrefix, primary)
}

func (m *defaultAppQuestionsTestsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsTestsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
