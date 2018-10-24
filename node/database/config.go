package database

import (
	"strconv"
)

type DatabaseConfig struct {
	MysqlHost    string
	MysqlPort    int
	MysqlSocket  string
	DatabaseName string
	DbUser       string
	DbPassword   string
	TablesPrefix string
}

func (dbc *DatabaseConfig) HasMinimum() bool {
	if (dbc.MysqlHost == "" || dbc.MysqlPort == 0) && dbc.MysqlSocket == "" || dbc.DatabaseName == "" {
		return false
	}
	return true
}

func (dbc *DatabaseConfig) GetServerAddress() string {
	return dbc.MysqlHost + ":" + strconv.Itoa(dbc.MysqlPort)
}

func (dbc *DatabaseConfig) GetMySQLConnString() string {
	if dbc.MysqlSocket != "" {
		return "unix(" + dbc.MysqlSocket + ")"
	}
	prefix := ""

	if dbc.DbUser != "" {
		prefix = dbc.DbUser + ":" + dbc.DbPassword + "@"
	}

	return prefix + "tcp(" + dbc.MysqlHost + ":" + strconv.Itoa(dbc.MysqlPort) + ")/" + dbc.DatabaseName
}
