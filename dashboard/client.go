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

	_, err = c.request("quit")
	if err != nil {
		return err
	}
	return c.conn.Close()
}

//PowerOn robot
func (c *Client) PowerOn() (bool, error) {
	return c.requestExpecting("power on", "powering on")
}

//PowerOff robot
func (c *Client) PowerOff() (bool, error) {
	return c.requestExpecting("power off", "powering off")
}

//BrakeRelease - releases the brakes
func (c *Client) BrakeRelease() (bool, error) {
	return c.requestExpecting("brake release", "brake releasing")
}

//LoadProgram file
func (c *Client) LoadProgram(urpFilePath string) (bool, error) {
	return c.requestExpectingPrefix("load "+urpFilePath, "loading program: ")
}

//GetProgram - returns program path
func (c *Client) GetProgram() (string, bool, error) {
	return c.requestExpectingPrefixedString("get loaded program", "loaded program: ")
}

//LoadInstallation file
func (c *Client) LoadInstallation(installationFilePath string) (bool, error) {
	return c.requestExpectingPrefix("load installation "+installationFilePath, "loading installation: ")
}

//Play program
func (c *Client) Play() (bool, error) {
	return c.requestExpectingPrefix("play", "starting program")
}

//Stop program
func (c *Client) Stop() (bool, error) {
	return c.requestExpectingPrefix("stop", "stopped")
}

//Pause program
func (c *Client) Pause() (bool, error) {
	return c.requestExpectingPrefix("pause", "pausing program")
}

//IsProgramSaved -checks if program saved
func (c *Client) IsProgramSaved() (bool, error) {
	return c.requestExpectingPrefix("isProgramSaved", "true")
}

//GetProgramState - returns current program state
func (c *Client) GetProgramState() (ProgramState, error) {
	result, err := c.request("programState")
	if err != nil {
		return ProgramStateUndefined, err
	}
	splited := strings.Split(result, " ")
	return programState(splited[0]), nil
}

//GetPolyscopeVersion - returns polyscope version
func (c *Client) GetPolyscopeVersion() (string, error) {
	return c.request("PolyscopeVersion")
}

//Shutdown program
func (c *Client) Shutdown() (bool, error) {
	return c.requestExpectingPrefix("shutdown", "shutting down")
}

//IsRunning returns true if in running state
func (c *Client) IsRunning() (bool, error) {
	return c.requestExpectingPrefix("running", "program running: true")
}

//Mode of robot
func (c *Client) Mode() (Mode, error) {
	response, _, err := c.requestExpectingPrefixedString("robotmode", "robotmode: ")
	return modeOf(response), err
}

//ShowPopup message
func (c *Client) ShowPopup(msg string) (bool, error) {
	return c.requestExpectingPrefix("popup "+msg, "showing popup")
}

//ClosePopup - closes popup dialog
func (c *Client) ClosePopup() (bool, error) {
	return c.requestExpectingPrefix("close popup", "closing popup")
}

//SafetyClosePopup - closes safety popup dialog
func (c *Client) SafetyClosePopup() (bool, error) {
	return c.requestExpectingPrefix("close safety popup", "closing safety popup")
}

//SafetyRestart - restarts robot safety subsystem
func (c *Client) SafetyRestart() (bool, error) {
	return c.requestExpectingPrefix("restart safety", "restarting safety")
}

//SafetyUnlockProtectiveStop - unlocks protective stop
func (c *Client) SafetyUnlockProtectiveStop() (bool, error) {
	return c.requestExpectingPrefix("unlock protective stop", "protective stop releasing")
}

//SafetyMode - returns safety mode
func (c *Client) SafetyMode() (SafetyMode, error) {
	result, _, err := c.requestExpectingPrefixedString("safetymode", "safetymode: ")
	if err != nil {
		return SafetyModeUndefined, err
	}
	return safetyMode(result), nil
}

//Log message
func (c *Client) Log(msg string) (bool, error) {
	return c.requestExpectingPrefix("addToLog "+msg, "added log message")
}

func (c *Client) request(cmd string) (response string, err error) {
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

func (c *Client) requestExpecting(cmd string, expecting string) (bool, error) {
	response, err := c.request(cmd)
	return expecting == strings.ToLower(response), err

}

func (c *Client) requestExpectingPrefix(cmd string, prefix string) (b bool, err error) {
	response, err := c.request(cmd)
	if err != nil {
		return false, err
	}
	lower := strings.ToLower(response)
	return strings.HasPrefix(lower, prefix), err
}

func (c *Client) requestExpectingPrefixedString(cmd string, prefix string) (s string, b bool, err error) {
	var response string
	response, err = c.request(cmd)
	if err != nil {
		return
	}
	lower := strings.ToLower(response)
	if strings.HasPrefix(lower, prefix) {
		return response[len(prefix):], true, nil
	}
	return
}
