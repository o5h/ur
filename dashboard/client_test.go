package dashboard_test

import (
	"log"
	"testing"

	"github.com/o5h/ur/dashboard"
)

type TT testing.T

func TestClient(tt *testing.T) {
	t := (*TT)(tt)

	client := dashboard.New()
	err := client.Dial("192.168.234.129:29999")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	t.testTrue(client.PowerOn())
	t.testTrue(client.BrakeRelease())
	ver, _ := client.GetPolyscopeVersion()
	log.Println("Version:= ", ver)

	t.testTrue(client.LoadProgram("w.urp"))
	prog, _, _ := client.GetProgram()
	t.testFalse(prog == "w.urp", err) // full path returned
	t.testTrue(client.LoadInstallation("default.installation"))
	t.testFalse(client.Play())
	mode, err := client.Mode()
	tt.Log("mode = ", mode)
	t.testTrue(mode == dashboard.ModePowerOff, err)
	t.testTrue(client.Stop())
	t.testFalse(client.IsRunning())

	t.testTrue(client.ShowPopup("Hello"))
	t.testTrue(client.ClosePopup())
	t.testTrue(client.SafetyClosePopup())

	t.testTrue(client.IsProgramSaved())
	tt.Log(client.GetProgramState())
	client.Log("Message")
	tt.Log(client.SafetyMode())
	t.testTrue(client.SafetyRestart())
	t.testTrue(client.SafetyUnlockProtectiveStop())

	//t.testTrue(client.PowerOff())

	//	client.Shutdown()
	//tt.Fail()
}

func (t *TT) testTrue(b bool, err error) {
	t.Helper()
	if !b {
		t.Fatalf("true expeced")
	}
	if err != nil {
		t.Fatal(err)
	}
}

func (t *TT) testFalse(b bool, err error) {
	t.Helper()
	if b {
		t.Fatalf("true expeced")
	}
	if err != nil {
		t.Fatal(err)
	}
}
