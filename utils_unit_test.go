package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestLog_UnsupportedLevel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			expMsg := "Unsupported log level: unsupported"
			if r.(string) != expMsg {
				t.Fatal(fmt.Sprintf("Expected panic message: %v, but was: %v", expMsg, r.(string)))
			}
		}
	}()
	Log("unsupported", "abc")
	t.Fatal("Expected panic, but no panic")
}

func TestLog_InfoShownByDefault(t *testing.T) {
	// set custom Log output target
	var str bytes.Buffer
	log.SetOutput(&str)

	Log("info", "hello")
	output := strings.TrimSuffix(str.String(), "\n")
	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "hello")
}

func TestLog_DebugShownByDefaultAtStart(t *testing.T) {
	// set custom Log output target
	var str bytes.Buffer
	log.SetOutput(&str)

	Log("debug", "my debug msg")
	output := strings.TrimSuffix(str.String(), "\n")
	assert.Contains(t, output, "DEBUG")
	assert.Contains(t, output, "my debug msg")
}

func TestLog_MixLevel_DebugUnset(t *testing.T) {
	// set custom Log output target
	var str bytes.Buffer
	log.SetOutput(&str)
	SetLogLevel("info")

	Log("info", "hello")
	Log("debug", "my debug msg")
	output := strings.TrimSuffix(str.String(), "\n")
	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "hello")
	assert.NotContains(t, output, "DEBUG")
	assert.NotContains(t, output, "my debug msg")
}

func TestLog_MixLevel_DebugSet(t *testing.T) {
	// set custom Log output target
	var str bytes.Buffer
	log.SetOutput(&str)

	SetLogLevel("debug")

	Log("info", "hello")
	Log("debug", "my debug msg")
	output := strings.TrimSuffix(str.String(), "\n")
	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "hello")
	assert.Contains(t, output, "DEBUG")
	assert.Contains(t, output, "my debug msg")
}

func logSth() {
	Log("debug", "logging sth")
}

// When you run this test separately, you see output.
func TestLog_MixLevel_ForHuman(t *testing.T) {
	SetLogLevel("debug")

	Log("info", "hello")
	Log("debug", "my debug msg")
	logSth()
}