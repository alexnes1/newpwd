package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGetFlags(t *testing.T) {
	tests := []struct {
		args   []string
		config pwdConfig
	}{
		{[]string{"-l", "5"}, pwdConfig{length: 5}},
		{[]string{"-length", "5"}, pwdConfig{length: 5}},
		{[]string{"--length", "5"}, pwdConfig{length: 5}},

		{[]string{"--no-lower"}, pwdConfig{noLower: true, length: 10}},
		{[]string{"-no-lower"}, pwdConfig{noLower: true, length: 10}},
		{[]string{"-w"}, pwdConfig{noLower: true, length: 10}},

		{[]string{"--no-upper"}, pwdConfig{noUpper: true, length: 10}},
		{[]string{"-no-upper"}, pwdConfig{noUpper: true, length: 10}},
		{[]string{"-u"}, pwdConfig{noUpper: true, length: 10}},

		{[]string{"--no-digits"}, pwdConfig{noDigits: true, length: 10}},
		{[]string{"-no-digits"}, pwdConfig{noDigits: true, length: 10}},
		{[]string{"-d"}, pwdConfig{noDigits: true, length: 10}},

		{[]string{"--no-punc"}, pwdConfig{noPunc: true, length: 10}},
		{[]string{"-no-punc"}, pwdConfig{noPunc: true, length: 10}},
		{[]string{"-p"}, pwdConfig{noPunc: true, length: 10}},

		{[]string{"--version"}, pwdConfig{showVersion: true, length: 10}},
		{[]string{"-version"}, pwdConfig{showVersion: true, length: 10}},
		{[]string{"-v"}, pwdConfig{showVersion: true, length: 10}},

		{[]string{"-l", "42", "-w", "-u", "-d", "-p", "-v"},
			pwdConfig{
				length:      42,
				noLower:     true,
				noUpper:     true,
				noDigits:    true,
				noPunc:      true,
				showVersion: true}},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, ", "), func(t *testing.T) {
			config, out, err := getFlags(tt.args)
			if out != "" {
				t.Errorf("expect empty out, got %q", out)
			}
			if err != nil {
				t.Errorf("expect empty err, got %s", err)
			}
			if !reflect.DeepEqual(config, tt.config) {
				t.Errorf("config: expected %+v, got %+v", tt.config, config)
			}
		})
	}

	t.Run("wrong flags", func(t *testing.T) {
		_, out, err := getFlags([]string{"--unknown-flag"})
		if out == "" {
			t.Errorf("expect output, got %q", out)
		}

		if err == nil {
			t.Errorf("expected err, got nil")
		}
	})
}

func TestRun(t *testing.T) {
	t.Run("version", func(t *testing.T) {
		out := &strings.Builder{}
		version := "v0.0.0-test"
		exitCode := run(out, []string{"--version"}, version)
		expectedOut := fmt.Sprintf("newpwd %s", version)
		if !strings.HasPrefix(out.String(), expectedOut) {
			t.Errorf("expected %q, got %q", expectedOut, out.String())
		}

		if exitCode != 0 {
			t.Errorf("expected exit code of 0, got %d", exitCode)
		}
	})

	t.Run("help", func(t *testing.T) {
		out := &strings.Builder{}
		version := "v0.0.0-test"
		exitCode := run(out, []string{"-h"}, version)
		expectedOut := "Usage of newpwd"
		if !strings.HasPrefix(out.String(), expectedOut) {
			t.Errorf("expected %q, got %q", expectedOut, out.String())
		}

		if exitCode != 2 {
			t.Errorf("expected exit code of 2, got %d", exitCode)
		}
	})

	t.Run("unknown flag", func(t *testing.T) {
		out := &strings.Builder{}
		version := "v0.0.0-test"
		exitCode := run(out, []string{"--unknown-flag"}, version)
		expectedOut := "flag provided but not defined:"
		if !strings.HasPrefix(out.String(), expectedOut) {
			t.Errorf("expected %q, got %q", expectedOut, out.String())
		}

		if exitCode != 1 {
			t.Errorf("expected exit code of 2, got %d", exitCode)
		}
	})

	t.Run("normal run", func(t *testing.T) {
		out := &strings.Builder{}
		version := "v0.0.0-test"
		exitCode := run(out, []string{"-l", "20"}, version)
		if len(out.String()) != 21 { // + '\n'
			t.Errorf("expected result of length 20, got %d", len(out.String()))
		}

		if exitCode != 0 {
			t.Errorf("expected exit code of 0, got %d", exitCode)
		}
	})
}
