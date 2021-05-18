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
	baseAppFieldNames          = builderx.RawFieldNames(&BaseApp{})
	baseAppRows                = strings.Join(baseAppFieldNames, ",")
	baseAppRowsExpectAutoSet   = strings.Join(stringx.Remove(baseAppFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	baseAppRowsWithPlaceHolder = strings.Join(stringx.Remove(baseAppFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBaseAppIdPrefix = "cache#baseApp#id#"
)

type (
	BaseAppModel interface {
		Insert(data BaseApp) (sql.Result, error)
		FindOne(id int64) (*BaseApp, error)
		Update(data BaseApp) error
		Delete(id int64) error
	}

	defaultBaseAppModel struct {
		sqlc.CachedConn
		table string
	}

	BaseApp struct {
		Id          int64          `db:"id"`
		Logo        sql.NullString `db:"logo"`        // 应用login
		Sname       string         `db:"sname"`       // 应用名称
		Isclose     int64          `db:"isclose"`     // 站点是否关闭
		Fullwebsite string         `db:"fullwebsite"` // 完整的域名
		Website     string         `db:"website"`     // 站点名称
		CreateBy    sql.NullString `db:"create_by"`
		UpdateBy    sql.NullString `db:"update_by"`
		CreatedAt   sql.NullTime   `db:"created_at"`
		UpdatedAt   sql.NullTime   `db:"updated_at"`
		DeletedAt   sql.NullTime   `db:"deleted_at"`
	}
)

func NewBaseAppModel(conn sqlx.SqlConn, c cache.CacheConf) BaseAppModel {
	return &defaultBaseAppModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`base_app`",
	}
}

func (m *defaultBaseAppModel) Insert(data BaseApp) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, baseAppRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Logo, data.Sname, data.Isclose, data.Fullwebsite, data.Website, data.CreateBy, data.UpdateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultBaseAppModel) FindOne(id int64) (*BaseApp, error) {
	baseAppIdKey := fmt.Sprintf("%s%v", cacheBaseAppIdPrefix, id)
	var resp BaseApp
	err := m.QueryRow(&resp, baseAppIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseAppRows, m.table)
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

func (m *defaultBaseAppModel) Update(data BaseApp) error {
	baseAppIdKey := fmt.Sprintf("%s%v", cacheBaseAppIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, baseAppRowsWithPlaceHolder)
		return conn.Exec(query, data.Logo, data.Sname, data.Isclose, data.Fullwebsite, data.Website, data.CreateBy, data.UpdateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, baseAppIdKey)
	return err
}

func (m *defaultBaseAppModel) Delete(id int64) error {

	baseAppIdKey := fmt.Sprintf("%s%v", cacheBaseAppIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, baseAppIdKey)
	return err
}

func (m *defaultBaseAppModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBaseAppIdPrefix, primary)
}

func (m *defaultBaseAppModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseAppRows, m.table)
	return conn.QueryRow(v, query, primary)
}
