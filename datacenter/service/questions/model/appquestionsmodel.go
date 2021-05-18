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
	appQuestionsFieldNames          = builderx.RawFieldNames(&AppQuestions{})
	appQuestionsRows                = strings.Join(appQuestionsFieldNames, ",")
	appQuestionsRowsExpectAutoSet   = strings.Join(stringx.Remove(appQuestionsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	appQuestionsRowsWithPlaceHolder = strings.Join(stringx.Remove(appQuestionsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppQuestionsIdPrefix       = "cache#appQuestions#id#"
	cacheAppQuestionsQuestionPrefix = "cache#appQuestions#question#"
)

type (
	AppQuestionsModel interface {
		Insert(data AppQuestions) (sql.Result, error)
		FindOne(id int64) (*AppQuestions, error)
		FindOneByQuestion(question string) (*AppQuestions, error)
		Update(data AppQuestions) error
		Delete(id int64) error
	}

	defaultAppQuestionsModel struct {
		sqlc.CachedConn
		table string
	}

	AppQuestions struct {
		Id         int64        `db:"id"`
		Beid       int64        `db:"beid"`        // 对应的平台
		Ptyid      int64        `db:"ptyid"`       // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		ActivityId int64        `db:"activity_id"` // 活动的id
		Question   string       `db:"question"`    // 问题
		Options    string       `db:"options"`     // 选项
		Corrent    string       `db:"corrent"`     // 正确选项（ABCD）
		Status     int64        `db:"status"`      // 状态（01）
		CreatedAt  sql.NullTime `db:"created_at"`
		UpdatedAt  sql.NullTime `db:"updated_at"`
	}
)

func NewAppQuestionsModel(conn sqlx.SqlConn, c cache.CacheConf) AppQuestionsModel {
	return &defaultAppQuestionsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_questions`",
	}
}

func (m *defaultAppQuestionsModel) Insert(data AppQuestions) (sql.Result, error) {
	appQuestionsQuestionKey := fmt.Sprintf("%s%v", cacheAppQuestionsQuestionPrefix, data.Question)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appQuestionsRowsExpectAutoSet)
		return conn.Exec(query, data.Beid, data.Ptyid, data.ActivityId, data.Question, data.Options, data.Corrent, data.Status, data.CreatedAt, data.UpdatedAt)
	}, appQuestionsQuestionKey)
	return ret, err
}

func (m *defaultAppQuestionsModel) FindOne(id int64) (*AppQuestions, error) {
	appQuestionsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsIdPrefix, id)
	var resp AppQuestions
	err := m.QueryRow(&resp, appQuestionsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsRows, m.table)
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

func (m *defaultAppQuestionsModel) FindOneByQuestion(question string) (*AppQuestions, error) {
	appQuestionsQuestionKey := fmt.Sprintf("%s%v", cacheAppQuestionsQuestionPrefix, question)
	var resp AppQuestions
	err := m.QueryRowIndex(&resp, appQuestionsQuestionKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `question` = ? limit 1", appQuestionsRows, m.table)
		if err := conn.QueryRow(&resp, query, question); err != nil {
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

func (m *defaultAppQuestionsModel) Update(data AppQuestions) error {
	appQuestionsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsIdPrefix, data.Id)
	appQuestionsQuestionKey := fmt.Sprintf("%s%v", cacheAppQuestionsQuestionPrefix, data.Question)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appQuestionsRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.ActivityId, data.Question, data.Options, data.Corrent, data.Status, data.CreatedAt, data.UpdatedAt, data.Id)
	}, appQuestionsIdKey, appQuestionsQuestionKey)
	return err
}

func (m *defaultAppQuestionsModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	appQuestionsIdKey := fmt.Sprintf("%s%v", cacheAppQuestionsIdPrefix, id)
	appQuestionsQuestionKey := fmt.Sprintf("%s%v", cacheAppQuestionsQuestionPrefix, data.Question)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, appQuestionsIdKey, appQuestionsQuestionKey)
	return err
}

func (m *defaultAppQuestionsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppQuestionsIdPrefix, primary)
}

func (m *defaultAppQuestionsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appQuestionsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
