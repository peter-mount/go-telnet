package telnet

import "net"

type Context interface {
  Logger() Logger              // Logger
  InjectLogger(Logger) Context // Inject logger
  LocalAddr() net.Addr         // Local address
  RemoteAddr() net.Addr        // Remote address
}

type internalContext struct {
  logger Logger
  con    net.Conn
}

func NewContext() Context {
  ctx := internalContext{}

  return &ctx
}

func (ctx *internalContext) Logger() Logger {
  return ctx.logger
}

func (ctx *internalContext) InjectLogger(logger Logger) Context {
  ctx.logger = logger

  return ctx
}

func (ctx *internalContext) LocalAddr() net.Addr {
  return ctx.con.LocalAddr()
}
func (ctx *internalContext) RemoteAddr() net.Addr {
  return ctx.con.RemoteAddr()
}
