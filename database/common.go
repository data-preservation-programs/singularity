package database

import (
	"net/url"
	"strings"

	"github.com/cockroachdb/errors"
)

func AddPragmaToSQLite(connString string) (string, error) {
	u, err := url.Parse(connString)
	if err != nil {
		return "", errors.WithStack(err)
	}

	qs := u.Query()
	qs.Add("_pragma", "busy_timeout(50000)")
	qs.Set("_pragma", "foreign_keys(1)")
	if strings.HasPrefix(connString, "file::memory:") {
		qs.Set("_pragma", "journal_mode(MEMORY)")
		qs.Set("mode", "memory")
		qs.Set("cache", "shared")
	} else {
		qs.Set("_pragma", "journal_mode(WAL)")
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
