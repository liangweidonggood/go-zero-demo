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
	appVotesActivityFieldNames          = builderx.RawFieldNames(&AppVotesActivity{})
	appVotesActivityRows                = strings.Join(appVotesActivityFieldNames, ",")
	appVotesActivityRowsExpectAutoSet   = strings.Join(stringx.Remove(appVotesActivityFieldNames, "`actid`", "`create_time`", "`update_time`"), ",")
	appVotesActivityRowsWithPlaceHolder = strings.Join(stringx.Remove(appVotesActivityFieldNames, "`actid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppVotesActivityActidPrefix = "cache#appVotesActivity#actid#"
)

type (
	AppVotesActivityModel interface {
		Insert(data AppVotesActivity) (sql.Result, error)
		FindOne(actid int64) (*AppVotesActivity, error)
		Update(data AppVotesActivity) error
		Delete(actid int64) error
	}

	defaultAppVotesActivityModel struct {
		sqlc.CachedConn
		table string
	}

	AppVotesActivity struct {
		Actid       int64          `db:"actid"`       // 投票活动的id
		Beid        int64          `db:"beid"`        // 对应的平台
		Ptyid       int64          `db:"ptyid"`       // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Title       sql.NullString `db:"title"`       // 投票活动名称
		Descr       sql.NullString `db:"descr"`       // 投票活动描述
		Status      int64          `db:"status"`      // 0无效，1.是有效
		StartDate   int64          `db:"start_date"`  // 投票活动开始时间
		EnrollDate  int64          `db:"enroll_date"` // 开始投票时间
		EndDate     int64          `db:"end_date"`    // 投票活动结束时间
		Votecount   int64          `db:"votecount"`   // 投票活动的总票数
		Enrollcount int64          `db:"enrollcount"` // 报名人数
		Viewcount   int64          `db:"viewcount"`   // 投票活动的总浏览量
		CreateTime  sql.NullTime   `db:"create_time"`
		CreateBy    int64          `db:"create_by"`
		UpdateBy    int64          `db:"update_by"`
		UpdateTime  sql.NullTime   `db:"update_time"`
		Type        sql.NullInt64  `db:"type"` // 投票的方式:1一次性，2.按天来
		Num         sql.NullInt64  `db:"num"`  // 单位
		DeletedAt   sql.NullTime   `db:"deleted_at"`
	}
)

func NewAppVotesActivityModel(conn sqlx.SqlConn, c cache.CacheConf) AppVotesActivityModel {
	return &defaultAppVotesActivityModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_votes_activity`",
	}
}

func (m *defaultAppVotesActivityModel) Insert(data AppVotesActivity) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appVotesActivityRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Beid, data.Ptyid, data.Title, data.Descr, data.Status, data.StartDate, data.EnrollDate, data.EndDate, data.Votecount, data.Enrollcount, data.Viewcount, data.CreateBy, data.UpdateBy, data.Type, data.Num, data.DeletedAt)

	return ret, err
}

func (m *defaultAppVotesActivityModel) FindOne(actid int64) (*AppVotesActivity, error) {
	appVotesActivityActidKey := fmt.Sprintf("%s%v", cacheAppVotesActivityActidPrefix, actid)
	var resp AppVotesActivity
	err := m.QueryRow(&resp, appVotesActivityActidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `actid` = ? limit 1", appVotesActivityRows, m.table)
		return conn.QueryRow(v, query, actid)
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

func (m *defaultAppVotesActivityModel) Update(data AppVotesActivity) error {
	appVotesActivityActidKey := fmt.Sprintf("%s%v", cacheAppVotesActivityActidPrefix, data.Actid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `actid` = ?", m.table, appVotesActivityRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Title, data.Descr, data.Status, data.StartDate, data.EnrollDate, data.EndDate, data.Votecount, data.Enrollcount, data.Viewcount, data.CreateBy, data.UpdateBy, data.Type, data.Num, data.DeletedAt, data.Actid)
	}, appVotesActivityActidKey)
	return err
}

func (m *defaultAppVotesActivityModel) Delete(actid int64) error {

	appVotesActivityActidKey := fmt.Sprintf("%s%v", cacheAppVotesActivityActidPrefix, actid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `actid` = ?", m.table)
		return conn.Exec(query, actid)
	}, appVotesActivityActidKey)
	return err
}

func (m *defaultAppVotesActivityModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppVotesActivityActidPrefix, primary)
}

func (m *defaultAppVotesActivityModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `actid` = ? limit 1", appVotesActivityRows, m.table)
	return conn.QueryRow(v, query, primary)
}
