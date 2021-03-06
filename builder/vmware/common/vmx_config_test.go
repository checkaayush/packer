package common

import (
	"testing"

	"github.com/hashicorp/packer/template/interpolate"
)

func TestVMXConfigPrepare(t *testing.T) {
	c := new(VMXConfig)
	c.VMXData = map[string]string{
		"one": "foo",
		"two": "bar",
	}

	errs := c.Prepare(interpolate.NewContext())
	if len(errs) > 0 {
		t.Fatalf("bad: %#v", errs)
	}

	if len(c.VMXData) != 2 {
		t.Fatal("should have two items in VMXData")
	}
}
