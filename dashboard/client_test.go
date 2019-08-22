package dashboard_test

import (
	"testing"

	"github.com/o5h/ur/dashboard"
)

func TestClient(t *testing.T) {
	client := dashboard.New()
	err := client.Dial("192.168.234.129:29999")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	client.PowerOn()
	client.BrakeRelease()

}
