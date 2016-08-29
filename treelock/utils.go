package treelock

import (
	"errors"
	"strings"
)

//ParseLockPath parse a lock path
//lock path format : for/bar/z
func ParseLockPath(path string) ([]string, error) {
	ps := strings.Split(path, "/")
	var bs []string
	for _, v := range ps {
		if len(v) < 1 {
			return nil, errors.New("parse path failed: has empty element.")
		} else if len(v) > 128 {
			return nil, errors.New("parse path failed:element is too large.max support 128 bit.")
		}
		bs = append(bs, strings.ToLower(v))
	}
	return bs, nil
}
