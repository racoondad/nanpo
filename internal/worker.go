package internal

import (
	"context"
	"math/rand"
	"time"
)

type Worker interface {
	// Returns current Logger.
	Logger() Logger

	// Set current Logger instead Logger.
	SetLogger(Logger)

	// Store Returns an address to current memstore.Store.
	//
	// memstore.Store is a collection of key/value entries. usually use to store metadata produced by freedom runtime.
	Store() *Memstore

	// Returns an address to current Bus.
	Bus() *Bus

	// Returns current context.
	Context() context.Context

	// Set current context instead Context.
	WithContext(context.Context)

	// Seturns a time since this Worker be created.
	StartTime() time.Time

	// DeferRecycle marks the resource won't be recycled immediately after
	// the request has ended.
	//
	// It is a bit hard to understand what is, here is a simply explain about
	// this.
	//
	// When an HTTP request is incoming, the program will probably serve a bunch
	// of business logic services, DB connections, transactions, Redis caches,
	// and so on. Once those procedures has done, the system should write
	// response and release those resource immediately. In other words, the
	// system should do some clean up procedures for this request. You might
	// thought it is a matter of course. But in special cases, such as goroutine
	// without synchronizing-signal. When all business procedures has completed
	// on business goroutine, and prepare to respond. GC system may be run before
	// the http handler goroutine to respond to the client. Once this opportunity
	// was met, the client will got an "Internal Server Error" or other wrong
	// result, because resource has been recycled by GC before to respond to client.
	DeferRecycle()

	// Indicates system need to wait a while for recycle resource.
	IsDeferRecycle() bool

	// Returns a rand.Rand act a random number seeder.
	Rand() *rand.Rand
}
