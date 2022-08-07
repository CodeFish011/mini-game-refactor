package server

import (
	"log"
	"net"
	"sync"
	"time"
)

var connectionMgr *ConnectionManager

type ConnectionManager struct {
	svr              *Server
	mu               sync.RWMutex
	OnlineConnection map[uint64]*Connection // 记录当前在线的所有链接
}

func init() {
	connectionMgr = &ConnectionManager{
		OnlineConnection: make(map[uint64]*Connection),
	}
}

func NewConnectionManager(s *Server) *ConnectionManager {
	connectionMgr.svr = s
	return connectionMgr
}

func GetConnectionManager() *ConnectionManager {
	return connectionMgr
}

func (m *ConnectionManager) HandleConnection(conn net.Conn) {
	c := NewConnection(conn, m.svr.Config)

	log.Printf("有新的链接接入，地址: %s\n", c.Conn.RemoteAddr().String())
	// 生成 Connection 的 ID
	c.ID = uint64(time.Now().Unix())

	// 将链接加入到管理
	m.mu.Lock()
	m.OnlineConnection[c.ID] = c
	m.mu.Unlock()

	c.Start()
}

func (m *ConnectionManager) GetConnectionByID(id uint64) (*Connection, bool) {
	m.mu.Lock()
	c, ok := m.OnlineConnection[id]
	m.mu.Unlock()
	if !ok {
		return nil, false
	}
	return c, true
}

func (m *ConnectionManager) DeleteConnectionByID(id uint64) bool {
	m.mu.Lock()
	_, ok := m.OnlineConnection[id]
	m.mu.Unlock()
	// OnlineConnection 中存在对应的 ID
	if ok {
		delete(m.OnlineConnection, id)
		return true
	}
	return true
}
