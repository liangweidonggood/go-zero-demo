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
	appVotesFieldNames          = builderx.RawFieldNames(&AppVotes{})
	appVotesRows                = strings.Join(appVotesFieldNames, ",")
	appVotesRowsExpectAutoSet   = strings.Join(stringx.Remove(appVotesFieldNames, "`avid`", "`create_time`", "`update_time`"), ",")
	appVotesRowsWithPlaceHolder = strings.Join(stringx.Remove(appVotesFieldNames, "`avid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppVotesAvidPrefix = "cache#appVotes#avid#"
)

type (
	AppVotesModel interface {
		Insert(data AppVotes) (sql.Result, error)
		FindOne(avid int64) (*AppVotes, error)
		Update(data AppVotes) error
		Delete(avid int64) error
	}

	defaultAppVotesModel struct {
		sqlc.CachedConn
		table string
	}

	AppVotes struct {
		Avid       int64        `db:"avid"`  // 投票人序号：自增
		Beid       int64        `db:"beid"`  // 对应的平台
		Ptyid      int64        `db:"ptyid"` // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Uid        int64        `db:"uid"`   // 中台表用户的id
		Auid       int64        `db:"auid"`  // 中台表appuser的id
		Actid      int64        `db:"actid"` // 投票活动的id
		Ip         string       `db:"ip"`    // 投票人IP
		Aeid       int64        `db:"aeid"`  // 投票的id
		CreateTime time.Time    `db:"create_time"`
		UpdateTime time.Time    `db:"update_time"`
		DeletedAt  sql.NullTime `db:"deleted_at"`
	}
)

func NewAppVotesModel(conn sqlx.SqlConn, c cache.CacheConf) AppVotesModel {
	return &defaultAppVotesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_votes`",
	}
}

func (m *defaultAppVotesModel) Insert(data AppVotes) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, appVotesRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Beid, data.Ptyid, data.Uid, data.Auid, data.Actid, data.Ip, data.Aeid, data.DeletedAt)

	return ret, err
}

func (m *defaultAppVotesModel) FindOne(avid int64) (*AppVotes, error) {
	appVotesAvidKey := fmt.Sprintf("%s%v", cacheAppVotesAvidPrefix, avid)
	var resp AppVotes
	err := m.QueryRow(&resp, appVotesAvidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `avid` = ? limit 1", appVotesRows, m.table)
		return conn.QueryRow(v, query, avid)
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

func (m *defaultAppVotesModel) Update(data AppVotes) error {
	appVotesAvidKey := fmt.Sprintf("%s%v", cacheAppVotesAvidPrefix, data.Avid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `avid` = ?", m.table, appVotesRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Uid, data.Auid, data.Actid, data.Ip, data.Aeid, data.DeletedAt, data.Avid)
	}, appVotesAvidKey)
	return err
}

func (m *defaultAppVotesModel) Delete(avid int64) error {

	appVotesAvidKey := fmt.Sprintf("%s%v", cacheAppVotesAvidPrefix, avid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `avid` = ?", m.table)
		return conn.Exec(query, avid)
	}, appVotesAvidKey)
	return err
}

func (m *defaultAppVotesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppVotesAvidPrefix, primary)
}

func (m *defaultAppVotesModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `avid` = ? limit 1", appVotesRows, m.table)
	return conn.QueryRow(v, query, primary)
}
