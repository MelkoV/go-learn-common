package dictionary

import (
	"fmt"
)

const User = "user"
const Admin = "admin"

type IStorage interface {
	Get(section string, key string, v ...interface{}) string
}

type storage struct {
	vars map[string]map[string]string
}

func (s *storage) Get(section string, key string, v ...interface{}) string {
	_, ok := s.vars[section]
	if !ok {
		return format(key, v...)
	}
	_, ok = s.vars[section][key]
	if !ok {
		return format(key, v...)
	}
	return format(s.vars[section][key], v...)
}

func format(message string, v ...interface{}) string {
	return fmt.Sprintf(message, v...)
}

func NewStorage(v map[string]map[string]string) *storage {
	s := &storage{
		vars: v,
	}
	return s
}
