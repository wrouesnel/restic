package hdfs

import (
	"restic"
	"github.com/colinmarc/hdfs"
)

type hdfsbackend struct {
	client *hdfs.Client

}

func Open(cfg Config) (restic.Backend, error) {

}