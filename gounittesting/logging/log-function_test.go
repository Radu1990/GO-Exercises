package logging

import (
	"testing"
)

func TestLogStuff (t *testing.T) {
	LogStuff("testlogfile", "updates 1")
	a := "123 again"
	LogStuff("testlogfile", a)
	LogStuff("testlogfile_2", "these updates are written in a separate log")
}