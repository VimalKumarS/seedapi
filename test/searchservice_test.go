package test

import (
	"os"
	"seedapi/service"
	"testing"
)


func TestMain(m *testing.M) {
	println("Tests are about to run")
	result := m.Run()
	println("Tests done executing")

	os.Exit(result)
}

func TestCanSubtractNumbers(t *testing.T) {
	//t.Parallel()
	result := service.Add(1, 2)
	//time.Sleep(1 * time.Second)
	if result != 3 {
		t.Error("Failed to subtract two numbers properly")
	}
}