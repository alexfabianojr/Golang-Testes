package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // o teste pode ser executado em paralelo

	if runtime.GOARCH == "amd64" {
		t.Skip("Nao funciona em arquitetura amd64")
	}
	t.Fail()
}
