package sip2

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

// The Handler inteface must be implemented by
type Handler interface {
	Handle(request Message) (response Message)
}

type Client struct {
	conn          net.Conn
	authenticated bool
	lastMessage   Message
	onlineSince   time.Time
}

// Server is a SIP2 server, listening for and accepting TCP-connections
// from clients. The transaction logic must be supplied by an implementation
// of the Handler interface.
type Server struct {
	ln net.Listener

	handler Handler

	// When useAuth is true, clients must authenticate with username and password
	// before being allowed to interact any further with the server.
	useAuth bool

	// When validateMsg is true, incoming and outgoing SIP messages are validated
	// according to the SIP2 specification, with any violation logged.
	validateMsg bool

	// Log all incoming requests and outgoing responses.
	Log bool
}

// NewServer returns a new SIP2 Server using the given handler. It will return an
// error if it fails to bind a listener to the given port.
func NewServer(h Handler, port int) (s *Server, err error) {
	s = &Server{
		handler: h,
	}

	s.ln, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return
	}

	return s, err
}

type Status struct {
	ClientsConnected int
	ClientIPs        []string
}

func (s *Server) Status() Status {
	return Status{}
}

func (s *Server) Run() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			//println(err.Error())
			continue
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(c net.Conn) {
	r := bufio.NewReader(c)
	for {
		b, err := r.ReadBytes('\r')
		if err != nil {
			if err != io.EOF {
				println(err.Error())
			}
			return
		}

		if s.Log {
			log.Printf("[%v] -> %s", c.RemoteAddr(), string(b))
		}
		req, err := Decode(b)
		if err != nil {
			println(err.Error())
			continue // TODO or return?
		}
		resp := s.handler.Handle(req)
		if err := resp.Encode(c); err != nil {
			if err != io.EOF {
				println(err.Error())
			}
			return
		}
		if s.Log {
			// TODO avoid encoding twice
			var b bytes.Buffer
			resp.Encode(&b)
			log.Printf("[%v] <- %s", c.RemoteAddr(), b.String())
		}

	}
}

func (s *Server) Close() {
	// TODO first close all connected clients?
	if s.ln != nil {
		s.ln.Close()
	}
}
