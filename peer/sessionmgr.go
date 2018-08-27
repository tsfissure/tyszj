package peer

import (
	"sync"
	"sync/atomic"
	"tyszj"
)

type ISessionManager interface {
	Add(ses tyszj.ISession)
	Remove(ses tyszj.ISession)
	Count() int32

	SetIDBase(base int64)
}

type SessionManager struct {
	sesById  sync.Map
	sesIDGen int64
	count    int32
}

func (self *SessionManager) Add(ses tyszj.ISession) {
	id := atomic.AddInt64(&self.sesIDGen, 1)
	self.count = atomic.AddInt32(&self.count, 1)
	ses.SetID(id)
	self.sesById.Store(id, ses)
}
func (self *SessionManager) Remove(ses tyszj.ISession) {
	self.sesById.Delete(ses.ID())
	self.count = atomic.AddInt32(&self.count, -1)
}
func (self *SessionManager) Count() int32 {
	return atomic.LoadInt32(&self.count)
}

func (self *SessionManager) SetIDBase(base int64) {
	atomic.StoreInt64(&self.sesIDGen, base)
}
func (self *SessionManager) VisitAllSession(f func(tyszj.ISession) bool) {
	self.sesById.Range(func(k, v interface{}) bool {
		return f(v.(tyszj.ISession))
	})
}
func (self *SessionManager) CloseAllSession() {
	self.VisitAllSession(func(ses tyszj.ISession) bool {
		ses.Close()
		return true
	})
}
