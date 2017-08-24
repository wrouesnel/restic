package hdfs

import (
	"net/url"
	"path"
	"strings"

	"restic/errors"
)

// Config contains all configuration necessary to connect to an HDFS cluster
type Config struct {
	Namenodes     []string
	KerberosUsername, Username, Password string
	Path		  string
}

const defaultPath = "/"

func ParseConfig(s string) (interface{}, error) {
	switch {
	case strings.HasPrefix(s, "hdfs://"):
		s = s[7:]
	case strings.HasPrefix(s, "hdfs:"):
		s = s[5:]
	default:
		return nil, errors.New("hdfs: invalid format")
	}

	// use the first entry as the potentially comma-separated list of namenodes and the
	// remainder as the path
	paths := strings.SplitN(s, "/", 2)

	hdfsConfig := &Config{}


	if len(paths) < 1 {
		return nil, errors.New("hdfs: invalid format, namenode not found")
	}
	hdfsConfig.Namenodes = strings.Split(paths[0], ",")

	if len(paths) == 2 {
		hdfsConfig.Path = paths[1]
	} else {
		hdfsConfig.Path = defaultPath
	}

	return hdfsConfig, nil
}