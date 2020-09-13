package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RunMain(t *testing.T) {
	err := os.Setenv("LIMIT_JOBS", "10")
	if err != nil {
		t.Fatal("Unable to set Environment variable LIMIT_JOBS")
	}
	main()
}

func Test_RandomNumber(t *testing.T) {
	n := randomNumber(10)
	assert.IsType(t, 0, n)
	assert.GreaterOrEqual(t, n, 0)
	assert.LessOrEqual(t, n, 10)
}

func Test_getIntEnv(t *testing.T) {
	err := os.Setenv("TEST_INT_VARIABLE", "1")
	if err != nil {
		t.Fatal("Unable to set Environment variable TEST_VARIABLE")
	}
	testIntVariable := getIntEnv("TEST_INT_VARIABLE", 10)
	assert.Equal(t, 1, testIntVariable)
}

func Test_concurrentJobs(t *testing.T) {
	concurrentJobs(2, 10)
}
