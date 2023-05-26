package frodao

import (
	"database/sql"
	"time"
)

type optionData struct {
	connMaxIdleTime time.Duration
	connMaxLifetime time.Duration
	maxIdleConns    int
	maxOpenConns    int
}

type Option func(o *optionData)

func ConnMaxIdleTime(value time.Duration) Option {
	return func(o *optionData) {
		o.connMaxIdleTime = value
	}
}

func ConnMaxLifetime(value time.Duration) Option {
	return func(o *optionData) {
		o.connMaxLifetime = value
	}
}

func MaxIdleConns(value int) Option {
	return func(o *optionData) {
		o.maxIdleConns = value
	}
}

func MaxOpenConns(value int) Option {
	return func(o *optionData) {
		o.maxOpenConns = value
	}
}

func SetOptions(dbSession *sql.DB, options []Option) {
	o := &optionData{
		connMaxIdleTime: 0,
		connMaxLifetime: 0,
		maxIdleConns:    2,
		maxOpenConns:    0,
	}

	for _, opt := range options {
		opt(o)
	}

	dbSession.SetConnMaxIdleTime(o.connMaxIdleTime)
	dbSession.SetConnMaxLifetime(o.connMaxLifetime)
	dbSession.SetMaxIdleConns(o.maxIdleConns)
	dbSession.SetMaxOpenConns(o.maxOpenConns)
}
