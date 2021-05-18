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
	baseMemberFieldNames          = builderx.RawFieldNames(&BaseMember{})
	baseMemberRows                = strings.Join(baseMemberFieldNames, ",")
	baseMemberRowsExpectAutoSet   = strings.Join(stringx.Remove(baseMemberFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	baseMemberRowsWithPlaceHolder = strings.Join(stringx.Remove(baseMemberFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBaseMemberIdPrefix     = "cache#baseMember#id#"
	cacheBaseMemberMobilePrefix = "cache#baseMember#mobile#"
)

type (
	BaseMemberModel interface {
		Insert(data BaseMember) (sql.Result, error)
		FindOne(id int64) (*BaseMember, error)
		FindOneByMobile(mobile string) (*BaseMember, error)
		Update(data BaseMember) error
		Delete(id int64) error
	}

	defaultBaseMemberModel struct {
		sqlc.CachedConn
		table string
	}

	BaseMember struct {
		Id         int64          `db:"id"`          // 用户id
		Username   string         `db:"username"`    // 帐号
		Password   string         `db:"password"`    // 密码
		Salt       string         `db:"salt"`        // 密码加盐
		Mobile     string         `db:"mobile"`      // 手机号
		Icard      sql.NullString `db:"icard"`       // 身份证号码
		Realname   string         `db:"realname"`    // 真实姓名
		Status     int64          `db:"status"`      // 状态
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"`
		DeletedAt  sql.NullTime   `db:"deleted_at"`
	}
)

func NewBaseMemberModel(conn sqlx.SqlConn, c cache.CacheConf) BaseMemberModel {
	return &defaultBaseMemberModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`base_member`",
	}
}

func (m *defaultBaseMemberModel) Insert(data BaseMember) (sql.Result, error) {
	baseMemberMobileKey := fmt.Sprintf("%s%v", cacheBaseMemberMobilePrefix, data.Mobile)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, baseMemberRowsExpectAutoSet)
		return conn.Exec(query, data.Username, data.Password, data.Salt, data.Mobile, data.Icard, data.Realname, data.Status, data.DeletedAt)
	}, baseMemberMobileKey)
	return ret, err
}

func (m *defaultBaseMemberModel) FindOne(id int64) (*BaseMember, error) {
	baseMemberIdKey := fmt.Sprintf("%s%v", cacheBaseMemberIdPrefix, id)
	var resp BaseMember
	err := m.QueryRow(&resp, baseMemberIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseMemberRows, m.table)
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

func (m *defaultBaseMemberModel) FindOneByMobile(mobile string) (*BaseMember, error) {
	baseMemberMobileKey := fmt.Sprintf("%s%v", cacheBaseMemberMobilePrefix, mobile)
	var resp BaseMember
	err := m.QueryRowIndex(&resp, baseMemberMobileKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", baseMemberRows, m.table)
		if err := conn.QueryRow(&resp, query, mobile); err != nil {
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

func (m *defaultBaseMemberModel) Update(data BaseMember) error {
	baseMemberIdKey := fmt.Sprintf("%s%v", cacheBaseMemberIdPrefix, data.Id)
	baseMemberMobileKey := fmt.Sprintf("%s%v", cacheBaseMemberMobilePrefix, data.Mobile)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, baseMemberRowsWithPlaceHolder)
		return conn.Exec(query, data.Username, data.Password, data.Salt, data.Mobile, data.Icard, data.Realname, data.Status, data.DeletedAt, data.Id)
	}, baseMemberIdKey, baseMemberMobileKey)
	return err
}

func (m *defaultBaseMemberModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	baseMemberIdKey := fmt.Sprintf("%s%v", cacheBaseMemberIdPrefix, id)
	baseMemberMobileKey := fmt.Sprintf("%s%v", cacheBaseMemberMobilePrefix, data.Mobile)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, baseMemberIdKey, baseMemberMobileKey)
	return err
}

func (m *defaultBaseMemberModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBaseMemberIdPrefix, primary)
}

func (m *defaultBaseMemberModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseMemberRows, m.table)
	return conn.QueryRow(v, query, primary)
}
