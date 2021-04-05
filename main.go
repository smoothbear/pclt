package main

import (
	"flag"
	"fmt"
	"github.com/otiai10/copy"
	"log"
	"os"
	user "os/user"
	"path/filepath"
)

const (
	projectDir = "/pclt/projects/"
)

type pclt struct{}

type saveArgs struct {
	cmd *flag.FlagSet

	env  bool
	name string
	path string
}

func (s *saveArgs) init() {
	env := s.cmd.Bool("e", false, "Decides saving .env file to project template.")
	s.cmd.String("name", "", "Decides will be saved project name")

	_ = s.cmd.Parse(os.Args[2:])

	s.path = s.cmd.Arg(0)
	if s.path == "" {
		path, _ := os.Getwd()
		s.path = path
	}

	if s.name == "" {
		path, _ := os.Getwd()
		s.name = filepath.Base(path)
	}

	s.env = *env
}

func (s *saveArgs) defaultSave() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error: %v", currentUser)
	}

	err = copy.Copy(s.path, currentUser.HomeDir+projectDir+s.name)
}

type createArgs struct {
	cmd *flag.FlagSet

	project string
	path    string
}

func (c *createArgs) defaultCreate() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error: %v", currentUser)
	}

	err = copy.Copy(currentUser.HomeDir+projectDir+c.project, c.path+c.project)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (c *createArgs) init() {
	project := c.cmd.String("pn", "default", "Set using project template.")

	_ = c.cmd.Parse(os.Args[2:])

	c.project = *project
	c.path = c.cmd.Arg(0)

	if c.project == "" || c.path == "" {
		log.Fatal("Error: Not enough to run this command.")
	}
}

func (p *pclt) init() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	saveCmd := flag.NewFlagSet("save", flag.ExitOnError)

	if len(os.Args) < 2 {
		log.Fatal("expected subcommands.")
	}

	switch os.Args[1] {
	case "create":
		var create createArgs
		create.cmd = createCmd
		create.init()

		switch create.project {
		default:
			create.defaultCreate()
		}

	case "save":
		var save saveArgs
		save.cmd = saveCmd
		save.init()
		save.defaultSave()

	case "-help":
		help()

	default:
		log.Fatalln("Error: No such subcommand does not exist")
	}
}

func help() {
	fmt.Println("List of commands")
	fmt.Println("\n- pclt create:")
	fmt.Println("\tpclt create -<project name> <path>")
	fmt.Println("\tIntroduce")
	fmt.Println("\t\t - creating projects by using saved project template or supported project package")
	fmt.Println("\tExample usage")
	fmt.Println("\t\t- pclt create -pn spring ./")
	fmt.Println("\n\n- pclt save:")
	fmt.Println("\tpclt save -<environment> <project name> <path>")
	fmt.Println("\tIntroduce")
	fmt.Println("\t\t - saving projects what on this directory.")
	fmt.Println("\tExample usage")
	fmt.Println("\t\t- pclt save -e -name elephant ./")
}

func main() {
	var p pclt
	p.init()
}
