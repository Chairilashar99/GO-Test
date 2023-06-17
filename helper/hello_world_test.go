package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Benchmark Table (recommendations for use)
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Heril",
			request: "Heril",
		},
		{
			name:    "Ashar",
			request: "Ashar",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

//Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Heril", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Heril")
		}
	})
	b.Run("Ashar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ashar")
		}
	})
}

//Benchmark Test
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Heril")
	}

}

func BenchmarkHelloWorldAshar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ashar")
	}

}

// Table Test (recommendations for use)
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Heril",
			request:  "Heril",
			expected: "Hello Heril",
		},
		{
			name:     "Ashar",
			request:  "Ashar",
			expected: "Hello Ashar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Heril", func(t *testing.T) {
		result := HelloWorld("Heril")
		require.Equal(t, "Hello Heril", result, "Result must be 'Hello Heril'")
	})
	t.Run("Ashar", func(t *testing.T) {
		result := HelloWorld("Ashar")
		require.Equal(t, "Hello Ashar", result, "Result must be 'Hello Ashar'")
	})
}

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
