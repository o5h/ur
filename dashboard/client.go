package dashboard

import (
	"bufio"
	"net"
	"strings"
)

const (
	//Port - default dashboard port
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

	_, err = c.io.ReadString('\n') // skip welcome message
	return err
}

//Close connection
func (c *Client) Close() (err error) {

	_, err = c.sendString("quit")
	if err != nil {
		return err
	}
	return c.conn.Close()
}

//PowerOn robot
func (c *Client) PowerOn() (bool, error) {
	return c.send("power on", "Powering on")
}

//PowerOff robot
func (c *Client) PowerOff() (bool, error) {
	return c.send("power off", "Powering off")
}

//BrakeRelease - releases the brakes
func (c *Client) BrakeRelease() (bool, error) {
	return c.send("brake release", "Brake releasing")
}

//LoadProgram file
func (c *Client) LoadProgram(urpFilePath string) (bool, error) {
	return c.sendHasPrefix("load "+urpFilePath, "loading program:")
}

//GetProgram - returns program path
func (c *Client) GetProgram() (string, bool, error) {
	return c.sendHasPrefixString("get loaded program", "loaded program:")
}

//LoadInstallation file
func (c *Client) LoadInstallation(installationFilePath string) (bool, error) {
	return c.sendHasPrefix("load installation "+installationFilePath, "loading installation:")
}

//Play program
func (c *Client) Play() (bool, error) {
	return c.sendHasPrefix("play", "starting program")
}

//Stop program
func (c *Client) Stop() (bool, error) {
	return c.sendHasPrefix("stop", "stopped")
}

//Pause program
func (c *Client) Pause() (bool, error) {
	return c.sendHasPrefix("pause", "pausing program")
}

//Shutdown program
func (c *Client) Shutdown() (bool, error) {
	return c.sendHasPrefix("shutdown", "shutting down")
}

//Running returns true if in running state
func (c *Client) Running() (bool, error) {
	return c.sendHasPrefix("running", "program running: true")
}

//Mode of robot
func (c *Client) Mode() (Mode, error) {
	response, err := c.sendString("robotmode")
	return modeOf(response[len("Robotmode: "):]), err
}

//ShowPopup message
func (c *Client) ShowPopup(msg string) (bool, error) {
	return c.sendHasPrefix("popup "+msg, "showing popup")

}

//ClosePopup - closes popup dialog
func (c *Client) ClosePopup() (bool, error) {
	return c.sendHasPrefix("close popup", "closing popup")

}

//Log message
func (c *Client) Log(msg string) (bool, error) {
	return c.sendHasPrefix("addToLog "+msg, "added log message")
}

//IsProgramSaved -checks if program saved
func (c *Client) IsProgramSaved() (bool, error) {
	return c.sendHasPrefix("isProgramSaved", "true")
}

//GetProgramState - returns current program state
func (c *Client) GetProgramState() (ProgramState, error) {
	result, err := c.sendString("programState")
	splited := strings.Split(result, " ")
	if err != nil {
		return ProgramStateUndefined, err
	}
	return programState(splited[0]), nil
}

//GetPolyscopeVersion - returns polyscope version
func (c *Client) GetPolyscopeVersion() (string, error) {
	return c.sendString("PolyscopeVersion")
}

func (c *Client) send(cmd string, expecting string) (bool, error) {
	response, err := c.sendString(cmd)
	return expecting == response, err

}

func (c *Client) sendString(cmd string) (response string, err error) {
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
	return str, err

}

func (c *Client) sendHasPrefix(cmd string, prefix string) (b bool, err error) {
	response, err := c.sendString(cmd)
	if err != nil {
		return false, err
	}
	lower := strings.ToLower(response)
	return strings.HasPrefix(lower, prefix), err
}

func (c *Client) sendHasPrefixString(cmd string, prefix string) (s string, b bool, err error) {
	var response string
	response, err = c.sendString(cmd)
	if err != nil {
		return
	}
	lower := strings.ToLower(response)
	if strings.HasPrefix(lower, prefix) {
		return response[len(prefix):], true, nil
	}
	return
}
