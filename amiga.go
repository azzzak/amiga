package amiga

import (
	"context"
	"errors"
	"math/rand"
	"net/textproto"
	"sync"
	"time"
)

// ErrNilConn error.
var ErrNilConn = errors.New("connection has been lost")

// Amiga struct.
type Amiga struct {
	conn   *textproto.Conn
	ctx    context.Context
	cancel context.CancelFunc

	set    chan setHandler
	get    chan getHandler
	remove chan string

	sync.Mutex
	actions map[string]chan map[string]string

	onConnect func(string)
	onError   func(ErrResp)
	onEvent   handler
}

// Options of connection.
type Options struct {
	Host     string
	Port     int
	Username string
	Secret   string
}

// ErrResp struct.
type ErrResp struct {
	ActionID,
	Message string
}

type handler func(map[string]string)

type setHandler struct {
	event string
	h     handler
}

type getHandler struct {
	event string
	c     chan handler
}

// OnConnect create handler for catching connection message.
func OnConnect(h func(string)) func(ami *Amiga) {
	return func(ami *Amiga) {
		ami.onConnect = h
	}
}

// OnError create handler for catching errors.
func OnError(h func(ErrResp)) func(ami *Amiga) {
	return func(ami *Amiga) {
		ami.onError = h
	}
}

// OnEvent create handler for catching all events.
func OnEvent(h handler) func(ami *Amiga) {
	return func(ami *Amiga) {
		ami.onEvent = h
	}
}

// New amiga constructor.
func New(opts ...func(*Amiga)) *Amiga {
	ami := &Amiga{
		set:     make(chan setHandler, 1),
		get:     make(chan getHandler, 1),
		remove:  make(chan string, 1),
		actions: make(map[string]chan map[string]string),
	}

	ami.ctx, ami.cancel = context.WithCancel(context.Background())

	for _, opt := range opts {
		opt(ami)
	}

	rand.Seed(time.Now().UnixNano())
	go handlers(ami)

	return ami
}

func defaultOptions(o *Options) {
	if o.Host == "" {
		o.Host = "localhost"
	}

	if o.Port == 0 {
		o.Port = 5038
	}
}

// Connect to AMI.
func (ami *Amiga) Connect(o Options) error {
	defaultOptions(&o)

	conn, err := ami.connect(o)
	if err != nil {
		return err
	}

	go dispatcher(ami, conn)

	return nil
}

// Action execute an AMI action. If action successfully executed, method returns a map with Response fields. Otherwise error contains the Error message.
func (ami *Amiga) Action(a string, cmd map[string]string) (map[string]string, error) {
	if ami.conn == nil {
		return nil, ErrNilConn
	}

	c := make(chan map[string]string, 1)
	defer close(c)

	if cmd["ActionID"] == "" {
		if cmd == nil {
			cmd = make(map[string]string)
		}

		cmd["ActionID"] = randID()
	}

	ami.Lock()
	ami.actions[cmd["ActionID"]] = c
	ami.Unlock()

	_, err := ami.conn.Cmd(payload(a, cmd))
	if err != nil {
		return nil, err
	}

	res := <-c

	if res["Response"] == "Error" {
		return nil, errors.New(res["Message"])
	}

	return res, nil
}

// RegisterHandler for specific event.
func (ami *Amiga) RegisterHandler(event string, f handler) {
	ami.set <- setHandler{
		event: event,
		h:     f,
	}
}

// UnregisterHandler for specific event.
func (ami *Amiga) UnregisterHandler(event string) {
	ami.remove <- event
}

// Close connection to AMI. After connection closed internal goroutines are stopped to prevent leaking. If you want to reconnect to AMI you have to create new amiga instance.
func (ami *Amiga) Close() error {
	if ami.conn == nil {
		return ErrNilConn
	}

	var err error

	_, err = ami.Action("Logoff", nil)
	if err != nil {
		if f := ami.onError; f != nil {
			f(ErrResp{
				Message: err.Error(),
			})
		}
	}

	err = ami.conn.Close()

	ami.Lock()
	ami.actions = nil
	ami.Unlock()

	return err
}
