package server

import (
	"bytes"
	"io"
	"strings"
	"log"
	"net"
	"sync"
	"net/url"
)

//HandlerFunc for
type HandlerFunc func(req *Request)

//Server for
type Server struct {
	addr string

	mu sync.RWMutex

	handlers map[string]HandlerFunc
}


//Request for
type Request struct {
	Conn net.Conn
	QueryParams url.Values
}

//NewServer for
func NewServer(addr string) *Server {
	return &Server{addr: addr, handlers: make(map[string]HandlerFunc)}

}

//Register for
func (s *Server) Register(path string, handler HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handlers[path] = handler
}


//Start for
func (s *Server) Start() error {
	
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Print(err)
		return err
	}
	defer func() {
		if cerr := listener.Close(); cerr != nil {

			if err == nil {
				err = cerr
				return
			}
			log.Print(cerr)
		}
	}()
	

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		
		go s.handle(conn)
		

	}

	



}

func (s *Server) handle(conn net.Conn) {
	var err error
	
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				log.Print(err)
			}
			log.Print(err)
		}
	}()

	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err == io.EOF {
		log.Printf("%s", buf[:n])
		log.Print(err)
	}
	if err != nil {
		log.Print(err)
	}
	log.Printf("%s", buf[:n])

	data := buf[:n]
	requestLineDeLim := []byte{'\r', '\n'}
	requestLineEnd := bytes.Index(data, requestLineDeLim)
	if requestLineEnd == -1 {

	}

	requestLine := string(data[:requestLineEnd])
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {

	}

	method, path, version := parts[0], parts[1], parts[2]

	if method != "GET" {

	}

	if version != "HTTP/1.1" {

	}

	uri, err := url.ParseRequestURI(path)

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(uri.Path)
	log.Print(uri.Query())
	
	newRequest := &Request{Conn: conn, QueryParams: uri.Query()}
	
	
		s.mu.RLock()
		handler, ok := s.handlers[uri.Path]
		
		s.mu.RUnlock()
		if ok == true {
			log.Print("Ok printed ", handler)
			handler(newRequest)
		} else {

			return
		}
	


	
}