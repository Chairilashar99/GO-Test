package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Before After Test: tdk peduli unit testnya success dan gagal before & after tetap di eksekusi
func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

//Skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}

	result := HelloWorld("Heril")
	require.Equal(t, "Hello Heril", result, "Result must be 'Hello Heril'")
}

//Require
func TestHelloWordRequire(t *testing.T) {
	result := HelloWorld("Heril")
	require.Equal(t, "Hello Heril", result, "Result must be 'Hello Heril'")
	fmt.Println("TestHelloWorldRequire Done")
}

//Assert
func TestHelloWordAssert(t *testing.T) {
	result := HelloWorld("Heril")
	assert.Equal(t, "Hello Heril", result, "Result must be 'Hello Heril'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldHeril(t *testing.T) {
	result := HelloWorld("Heril")

	if result != "Hello Heril" {
		//error
		t.Error("Result must be 'Hello Heril'")
	}

	fmt.Println("TestHelloWorldHeril Done")
}

func TestHelloWorldAshar(t *testing.T) {
	result := HelloWorld("Ashar")

	if result != "Hello Ashar" {
		//error
		t.Fatal("Result must be 'Hello Ashar'")
	}
	fmt.Println("TestHelloWorldAshar Done")
}
