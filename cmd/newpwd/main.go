package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/alexnes1/newpwd"
)

var version string

const defaultLength = 10
const progname = "newpwd"

type pwdConfig struct {
	length      int
	noLower     bool
	noUpper     bool
	noDigits    bool
	noPunc      bool
	showVersion bool
}

func getFlags(args []string) (config pwdConfig, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	lengthUsage := "password length"
	flags.IntVar(&config.length, "length", defaultLength, lengthUsage)
	flags.IntVar(&config.length, "l", defaultLength, lengthUsage+" (shorthand)")

	lowerUsage := "no lowercase letters"
	flags.BoolVar(&config.noLower, "no-lower", false, lowerUsage)
	flags.BoolVar(&config.noLower, "w", false, lowerUsage+" (shorthand)")

	upperUsage := "no uppercase letters"
	flags.BoolVar(&config.noUpper, "no-upper", false, upperUsage)
	flags.BoolVar(&config.noUpper, "u", false, upperUsage+" (shorthand)")

	digitsUsage := "no digits"
	flags.BoolVar(&config.noDigits, "no-digits", false, digitsUsage)
	flags.BoolVar(&config.noDigits, "d", false, digitsUsage+" (shorthand)")

	puncUsage := "no punctuation symbols"
	flags.BoolVar(&config.noPunc, "no-punc", false, puncUsage)
	flags.BoolVar(&config.noPunc, "p", false, puncUsage+" (shorthand)")

	versionUsage := "show version"
	flags.BoolVar(&config.showVersion, "version", false, versionUsage)
	flags.BoolVar(&config.showVersion, "v", false, versionUsage+" (shorthand)")

	err = flags.Parse(args)
	output = buf.String()

	return
}

func run(out io.Writer, args []string, version string) int {
	config, output, err := getFlags(args)

	if err == flag.ErrHelp {
		fmt.Fprintln(out, output)
		return 2
	} else if err != nil {
		fmt.Fprintf(out, "%s\n", err)
		return 1
	}

	if config.showVersion {
		fmt.Fprintf(out, "%s %s (%s)\n", progname, version, runtime.Version())
		return 0
	}

	fmt.Fprintln(out, newpwd.Make(config.length, !config.noLower, !config.noUpper,
		!config.noDigits, !config.noPunc))
	return 0
}

func main() {
	rand.Seed(time.Now().UnixNano())

	exitCode := run(os.Stdout, os.Args[1:], version)
	os.Exit(exitCode)
}
