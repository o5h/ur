package dashboard

import (
	"bufio"
	"log"
	"net"
	"strings"
)

const (
	Port = "29999"
)

// Client - UR Dashboard Client
type Client struct {
	host string
	conn net.Conn
	io   *bufio.ReadWriter
}

//New - creates new Client
func New() *Client {
	return &Client{}
}

//Dial - connects to dashboard
func (c *Client) Dial(addr string) (err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return
	}

	c.conn, err = net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		return
	}

	c.io = bufio.NewReadWriter(
		bufio.NewReader(c.conn),
		bufio.NewWriter(c.conn))

	var str string
	str, err = c.io.ReadString('\n') // skip welcome message
	log.Println("Dashboard:", str)

	return err
}

//Close connection
func (c *Client) Close() {
	c.conn.Close()
}

//PowerOn robot
func (c *Client) PowerOn() (bool, error) {
	return c.send("power on", "Powering on")
}

//BrakeRelease - releases the brakes
func (c *Client) BrakeRelease() (bool, error) {
	return c.send("brake release", "Brake releasing")
}

//Load progra file
func (c *Client) Load(urpFilePath string) (bool, error) {
	return c.sendStartWith("load "+urpFilePath, "Loading program:")
}

func (c *Client) send(cmd string, expecting string) (b bool, err error) {
	_, err = c.io.WriteString(cmd)
	if err != nil {
		return
	}

	_, err = c.io.WriteString("\n")
	if err != nil {
		return
	}

	err = c.io.Flush()
	if err != nil {
		return
	}

	var str string
	str, err = c.io.ReadString('\n')
	if err != nil {
		return
	}
	str = str[:len(str)-1]
	return expecting == str, err

}

func (c *Client) sendStartWith(cmd string, start string) (b bool, err error) {
	_, err = c.io.WriteString(cmd)
	if err != nil {
		return
	}

	err = c.io.Flush()
	if err != nil {
		return
	}

	var line string
	line, err = c.io.ReadString('\n')
	if err != nil {
		return
	}
	return strings.HasPrefix(line, start), err

}
