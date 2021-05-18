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
	appEnrollFieldNames          = builderx.RawFieldNames(&AppEnroll{})
	appEnrollRows                = strings.Join(appEnrollFieldNames, ",")
	appEnrollRowsExpectAutoSet   = strings.Join(stringx.Remove(appEnrollFieldNames, "`aeid`", "`create_time`", "`update_time`"), ",")
	appEnrollRowsWithPlaceHolder = strings.Join(stringx.Remove(appEnrollFieldNames, "`aeid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppEnrollAeidPrefix = "cache#appEnroll#aeid#"
)

type (
	AppEnrollModel interface {
		Insert(data AppEnroll) (sql.Result, error)
		FindOne(aeid int64) (*AppEnroll, error)
		Update(data AppEnroll) error
		Delete(aeid int64) error
	}

	defaultAppEnrollModel struct {
		sqlc.CachedConn
		table string
	}

	AppEnroll struct {
		Aeid       int64          `db:"aeid"`
		Beid       int64          `db:"beid"`      // 对应的平台
		Ptyid      int64          `db:"ptyid"`     // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Uid        int64          `db:"uid"`       // 中台表用户的id
		Auid       int64          `db:"auid"`      // 中台表appuser的id
		Actid      int64          `db:"actid"`     // 投票活动的id
		Name       sql.NullString `db:"name"`      // 名字
		Address    sql.NullString `db:"address"`   // 地址
		Images     sql.NullString `db:"images"`    // 图片
		Descr      sql.NullString `db:"descr"`     // 介绍
		Votecount  int64          `db:"votecount"` // 投票数
		Viewcount  int64          `db:"viewcount"` // 浏览数
		Status     int64          `db:"status"`    // 0,未审核，1.审核通过，2.删除
		UpdateBy   int64          `db:"update_by"`
		CreateBy   int64          `db:"create_by"`
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"`
		DeletedAt  sql.NullTime   `db:"deleted_at"`
	}
)

func NewAppEnrollModel(conn sqlx.SqlConn, c cache.CacheConf) AppEnrollModel {
	return &defaultAppEnrollModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_enroll`",
	}
}

func (m *defaultAppEnrollModel) Insert(data AppEnroll) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appEnrollRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Beid, data.Ptyid, data.Uid, data.Auid, data.Actid, data.Name, data.Address, data.Images, data.Descr, data.Votecount, data.Viewcount, data.Status, data.UpdateBy, data.CreateBy, data.DeletedAt)

	return ret, err
}

func (m *defaultAppEnrollModel) FindOne(aeid int64) (*AppEnroll, error) {
	appEnrollAeidKey := fmt.Sprintf("%s%v", cacheAppEnrollAeidPrefix, aeid)
	var resp AppEnroll
	err := m.QueryRow(&resp, appEnrollAeidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `aeid` = ? limit 1", appEnrollRows, m.table)
		return conn.QueryRow(v, query, aeid)
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

func (m *defaultAppEnrollModel) Update(data AppEnroll) error {
	appEnrollAeidKey := fmt.Sprintf("%s%v", cacheAppEnrollAeidPrefix, data.Aeid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `aeid` = ?", m.table, appEnrollRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Uid, data.Auid, data.Actid, data.Name, data.Address, data.Images, data.Descr, data.Votecount, data.Viewcount, data.Status, data.UpdateBy, data.CreateBy, data.DeletedAt, data.Aeid)
	}, appEnrollAeidKey)
	return err
}

func (m *defaultAppEnrollModel) Delete(aeid int64) error {

	appEnrollAeidKey := fmt.Sprintf("%s%v", cacheAppEnrollAeidPrefix, aeid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `aeid` = ?", m.table)
		return conn.Exec(query, aeid)
	}, appEnrollAeidKey)
	return err
}

func (m *defaultAppEnrollModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppEnrollAeidPrefix, primary)
}

func (m *defaultAppEnrollModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `aeid` = ? limit 1", appEnrollRows, m.table)
	return conn.QueryRow(v, query, primary)
}
