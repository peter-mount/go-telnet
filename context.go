package telnet

import "net"

type Context interface {
  Logger() Logger                        // Logger
  InjectLogger(Logger, net.Conn) Context // Inject logger
  LocalAddr() net.Addr                   // Local address
  RemoteAddr() net.Addr                  // Remote address
  UserData() *map[string]interface{}     // userData
}

type internalContext struct {
  logger   Logger
  con      net.Conn
  userData map[string]interface{}
}

func NewContext() Context {
  ctx := internalContext{
    userData: make(map[string]interface{}),
  }

  return &ctx
}

func (ctx *internalContext) UserData() *map[string]interface{} {
  return &ctx.userData
}

func (ctx *internalContext) Logger() Logger {
  return ctx.logger
}

func (ctx *internalContext) InjectLogger(logger Logger, con net.Conn) Context {
  ctx.logger = logger
  ctx.con = con
  return ctx
}

func (ctx *internalContext) LocalAddr() net.Addr {
  if ctx.con == nil {
    return nil
  }
  return ctx.con.LocalAddr()
}
func (ctx *internalContext) RemoteAddr() net.Addr {
  if ctx.con == nil {
    return nil
  }
  return ctx.con.RemoteAddr()
}
