package subpackage1_test

import (
	"testing"

	"github.com/sumitbopche01/gomonorepo/subpackage1"
)

func TestSubpackagePrint(t *testing.T) {
	subpackage1.PrintPackage1()
	t.Log("subpackage run success ")

	if 1 != 1 {
		t.Errorf("Got Error")
	}
}
