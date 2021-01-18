package amiga

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/textproto"
	"strings"
)

func handlers(ami *Amiga) {
	handlers := make(map[string]handler)

	for {
		select {
		case <-ami.ctx.Done():
			return
		case p := <-ami.set:
			handlers[p.event] = p.h
		case p := <-ami.get:
			p.c <- handlers[p.event]
		case p := <-ami.remove:
			delete(handlers, p)
		}
	}
}

func (ami *Amiga) connect(o Options) (*textproto.Conn, error) {
	var (
		conn *textproto.Conn
		err  error
	)

	connString := fmt.Sprintf("%s:%d", o.Host, o.Port)
	conn, err = textproto.Dial("tcp", connString)
	if err != nil {
		return nil, err
	}

	ami.conn = conn

	var welcome string
	welcome, err = conn.ReadLine()
	if err != nil {
		return nil, err
	}

	if f := ami.onConnect; f != nil {
		f(welcome)
	}

	_, err = ami.conn.Cmd(payload("Login", map[string]string{
		"Username": o.Username,
		"Secret":   o.Secret,
	}))
	if err != nil {
		return nil, err
	}

	var m map[string]string
	m, err = readFields(conn)
	if err != nil {
		return nil, err
	}

	if m["Response"] == "Error" {
		return nil, errors.New(m["Message"])
	}

	return conn, nil
}

func dispatcher(ami *Amiga, conn *textproto.Conn) {
	defer ami.cancel()

	for {
		m, err := readFields(conn)
		if err != nil {
			if f := ami.onError; f != nil {
				f(ErrResp{
					Message: err.Error(),
				})
			}
			return
		}

		event, ok := m["Event"]
		if ok {
			if f := ami.checkHandler(event); f != nil {
				f(m)
			}

			if f := ami.onEvent; f != nil {
				f(m)
			}
		}

		resp, ok := m["Response"]
		if ok {
			key := m["ActionID"]

			ami.Lock()
			if c, ok := ami.actions[key]; ok {
				c <- m
				delete(ami.actions, key)
			}
			ami.Unlock()

			if resp == "Goodbye" {
				return
			}

			if resp == "Error" {
				if f := ami.onError; f != nil {
					f(ErrResp{
						ActionID: key,
						Message:  m["Message"],
					})
				}
			}
		}
	}
}

func (ami *Amiga) checkHandler(event string) handler {
	c := make(chan handler)
	defer close(c)

	ami.get <- getHandler{
		event: event,
		c:     c,
	}

	return <-c
}

func payload(action string, m map[string]string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Action: %s\r\n", action))

	for k, v := range m {
		s := fmt.Sprintf("%s: %s\r\n", k, v)
		b.WriteString(s)
	}

	return b.String()
}

func readFields(conn *textproto.Conn) (map[string]string, error) {
	var m = make(map[string]string)
	var err error

	for {
		var status string
		status, err = conn.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			break
		}

		if status == "" {
			break
		}

		p := strings.SplitN(status, ": ", 2)
		if len(p) > 1 {
			m[p[0]] = p[1]
		}
	}

	return m, err
}

const letterBytes = "0123456789abcdef"
const (
	letterIdxBits = 4
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func randID() string {
	var b strings.Builder
	b.Grow(14)
	for i, cache := 11, rand.Int63(); i >= 0; {
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits

		if i == 3 || i == 7 {
			b.WriteByte('-')
		}
	}
	return b.String()
}
