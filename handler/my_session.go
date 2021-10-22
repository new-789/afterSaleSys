package handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type sess struct {
	Session sessions.Session
}
var Store cookie.Store

func Init(c *gin.Context) *sess {
	option := sessions.Options{MaxAge: 3600}
	se := &sess{
		Session: sessions.Default(c),
	}
	se.Session.Options(option)
	return se
}

func (s *sess) SetSe(key, val string) bool {
	s.Session.Set(key, val)
	err := s.Session.Save()
	if err != nil {
		fmt.Println("se set failed", err)
		return false
	}
	seBool := s.SaveSe()
	if seBool {
		return true
	}
	return false
}

func (s *sess) GetSe(key string) string {
	res := s.Session.Get(key)
	if val, ok := res.(string); ok {
		return val
	}
	return ""
}

func (s *sess) DelSe(key string) bool {
	s.Session.Delete(key)
	seBool := s.SaveSe()
	if seBool {
		return true
	}
	return false
}

func (s *sess) SaveSe() bool {
	err := s.Session.Save()
	if err != nil {
		return false
	}
	return true
}
