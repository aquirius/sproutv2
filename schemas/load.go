package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

var dbName string

func init() {
	flag.StringVar(&dbName, "db", "sprout", "sprout")
	flag.Parse()
}

func run(name string, cmds []string, stdin []byte) error {
	cmd := exec.Command(name, cmds...)

	// capture outputs
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	assembleErr := func(message string, e error) error {
		return fmt.Errorf("%s: %s\n%s\n%s", message, stdout.String(), stderr.String(), e)
	}

	// fwd data to execution?
	if stdin != nil {
		cmd.Stdin = bytes.NewReader(stdin)
	}

	err := cmd.Start()
	if err != nil {
		return assembleErr("Failed to start command", err)
	}

	if err := cmd.Wait(); err != nil {
		return assembleErr("Failed to wait for command", err)
	}
	return nil
}

func main() {
	users, err := filepath.Abs("./schemas/user.sql")
	if err != nil {
		panic(err.Error())
	}

	plants, err := filepath.Abs("./schemas/plant.sql")
	if err != nil {
		panic(err.Error())
	}

	baskets, err := filepath.Abs("./schemas/basket.sql")
	if err != nil {
		panic(err.Error())
	}

	genetics, err := filepath.Abs("./schemas/genetic.sql")
	if err != nil {
		panic(err.Error())
	}

	if len(baskets) <= 0 {
		fmt.Errorf("no basket schema found")
	}

	if len(genetics) <= 0 {
		fmt.Errorf("no genetic schema found")
	}

	if len(users) <= 0 {
		fmt.Errorf("no user schema found")
	}

	if len(plants) <= 0 {
		fmt.Errorf("no plant schema found")
	}

	out := []string{}

	c, err := ioutil.ReadFile(users)
	if err != nil {
		panic(err.Error())
	}
	b, err := ioutil.ReadFile(baskets)
	if err != nil {
		panic(err.Error())
	}
	p, err := ioutil.ReadFile(plants)
	if err != nil {
		panic(err.Error())
	}
	g, err := ioutil.ReadFile(genetics)
	if err != nil {
		panic(err.Error())
	}
	out = append(out, string(c), string(p), string(g), string(b))

	s := ""
	for _, v := range out {
		s += v
		if !strings.HasSuffix(v, ";") {
			s += ";"
		}
		s += "\n\n"
	}
	s = `
	SET FOREIGN_KEY_CHECKS=0;
	` + s + `
	SET FOREIGN_KEY_CHECKS=1;
	`

	if err := run("docker", []string{
		"exec",
		"-i",
		"sqlsprout-backend",
		"mysql",
		"-uroot",
		"-psprout",
		"sprout",
	}, []byte(s)); err != nil {
		fmt.Errorf("failed to load sprout sql schemas: %w", err)
	}

}
