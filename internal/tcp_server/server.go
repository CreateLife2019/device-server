package tcp_server

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"net"
)

type Client struct {
	conn   net.Conn
	Server *Server
}

// TCP server
type Server struct {
	address                  string // Address to open connection: localhost:9999
	config                   *tls.Config
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message []byte)
}

// Read client data from channel
func (c *Client) listen() {
	c.Server.onNewClientCallback(c)
	for {
		message := make([]byte, 1024)
		n, err := c.conn.Read(message)
		if err != nil {
			break
		}
		logrus.Infof("read count:%d", n)
		if err != nil {
			err = c.conn.Close()
			if err != nil {
				logrus.Warnf("close conn failed")
			}
			c.Server.onClientConnectionClosed(c, err)
			return
		}

		c.Server.onNewMessage(c, message[:n])
	}
}

// Send text message to client
func (c *Client) Send(message string) error {
	return c.SendBytes([]byte(message))
}

// SendBytes Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	_, err := c.conn.Write(b)
	if err != nil {
		err = c.conn.Close()
		if err != nil {
			logrus.Warnf("conn close failed:%s", err.Error())
		}
		c.Server.onClientConnectionClosed(c, err)
	}
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// OnNewClient Called right after server starts listening new client
func (s *Server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// OnClientConnectionClosed Called right after connection closed
func (s *Server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

// OnNewMessage Called when Client receives new message
func (s *Server) OnNewMessage(callback func(c *Client, message []byte)) {
	s.onNewMessage = callback
}

// Listen starts network server
func (s *Server) Listen() {
	var listener net.Listener
	var err error
	if s.config == nil {
		listener, err = net.Listen("tcp", s.address)
	} else {
		listener, err = tls.Listen("tcp", s.address, s.config)
	}
	if err != nil {
		logrus.Fatal("Error starting TCP server.\r\n", err)
	} else {
		logrus.Infof("start tcp server suc,address:%s", s.address)
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			logrus.Error("Error starting TCP server.\r\n", err)
		}
	}()

	for {
		conn, _ := listener.Accept()
		client := &Client{
			conn:   conn,
			Server: s,
		}
		go client.listen()
	}
}

// New Creates new tcp server instance
func New(address string) *Server {
	logrus.Debugln("Creating server with address", address)
	s := &Server{
		address: address,
	}

	s.OnNewClient(func(c *Client) {})
	s.OnNewMessage(func(c *Client, message []byte) {})
	s.OnClientConnectionClosed(func(c *Client, err error) {})

	return s
}

func NewWithTLS(address, certFile, keyFile string) *Server {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		logrus.Fatal("Error loading certificate files. Unable to create TCP server with TLS functionality.\r\n", err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	s := New(address)
	s.config = config
	return s
}
