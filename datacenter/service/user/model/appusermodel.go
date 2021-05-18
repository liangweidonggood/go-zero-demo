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
	appUserFieldNames          = builderx.RawFieldNames(&AppUser{})
	appUserRows                = strings.Join(appUserFieldNames, ",")
	appUserRowsExpectAutoSet   = strings.Join(stringx.Remove(appUserFieldNames, "`auid`", "`create_time`", "`update_time`"), ",")
	appUserRowsWithPlaceHolder = strings.Join(stringx.Remove(appUserFieldNames, "`auid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppUserAuidPrefix = "cache#appUser#auid#"
)

type (
	AppUserModel interface {
		Insert(data AppUser) (sql.Result, error)
		FindOne(auid int64) (*AppUser, error)
		Update(data AppUser) error
		Delete(auid int64) error
	}

	defaultAppUserModel struct {
		sqlc.CachedConn
		table string
	}

	AppUser struct {
		Auid       int64          `db:"auid"`
		Beid       int64          `db:"beid"`     // 对应的平台
		Ptyid      int64          `db:"ptyid"`    // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Uid        int64          `db:"uid"`      // 对应中台表中的id
		Openid     sql.NullString `db:"openid"`   // 社交属性的openid
		Nickname   sql.NullString `db:"nickname"` // 昵称
		Avator     sql.NullString `db:"avator"`   // 头像
		Sex        int64          `db:"sex"`      // 性别
		City       sql.NullString `db:"city"`
		Province   sql.NullString `db:"province"`
		Country    sql.NullString `db:"country"`
		Privilege  sql.NullString `db:"privilege"`
		Unionid    sql.NullString `db:"unionid"`
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"`
		DeletedAt  sql.NullTime   `db:"deleted_at"`
	}
)

func NewAppUserModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserModel {
	return &defaultAppUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_user`",
	}
}

func (m *defaultAppUserModel) Insert(data AppUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appUserRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Beid, data.Ptyid, data.Uid, data.Openid, data.Nickname, data.Avator, data.Sex, data.City, data.Province, data.Country, data.Privilege, data.Unionid, data.DeletedAt)

	return ret, err
}

func (m *defaultAppUserModel) FindOne(auid int64) (*AppUser, error) {
	appUserAuidKey := fmt.Sprintf("%s%v", cacheAppUserAuidPrefix, auid)
	var resp AppUser
	err := m.QueryRow(&resp, appUserAuidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `auid` = ? limit 1", appUserRows, m.table)
		return conn.QueryRow(v, query, auid)
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

func (m *defaultAppUserModel) Update(data AppUser) error {
	appUserAuidKey := fmt.Sprintf("%s%v", cacheAppUserAuidPrefix, data.Auid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `auid` = ?", m.table, appUserRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Uid, data.Openid, data.Nickname, data.Avator, data.Sex, data.City, data.Province, data.Country, data.Privilege, data.Unionid, data.DeletedAt, data.Auid)
	}, appUserAuidKey)
	return err
}

func (m *defaultAppUserModel) Delete(auid int64) error {

	appUserAuidKey := fmt.Sprintf("%s%v", cacheAppUserAuidPrefix, auid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `auid` = ?", m.table)
		return conn.Exec(query, auid)
	}, appUserAuidKey)
	return err
}

func (m *defaultAppUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppUserAuidPrefix, primary)
}

func (m *defaultAppUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `auid` = ? limit 1", appUserRows, m.table)
	return conn.QueryRow(v, query, primary)
}
