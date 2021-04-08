package main

import (
	"flag"
	"fmt"
	"github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"os"
	user "os/user"
	tea "github.com/charmbracelet/bubbletea"
	"path/filepath"
)

const (
	projectDir = "/pclt/projects/"
)

type pclt struct{}

type saveArgs struct {
	cmd *flag.FlagSet
	user *user.User

	env  bool
	name string
	path string
}

func (s *saveArgs) init() {
	env := s.cmd.Bool("e", false, "Decides saving .env file to project template.")
	s.cmd.String("name", "", "Decides will be saved project name")

	_ = s.cmd.Parse(os.Args[2:])

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
	user *user.User

	project string
	path    string
}

func (c *createArgs) defaultCreate() {
	err := copy.Copy(c.user.HomeDir+projectDir+c.project, c.path+c.project)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (c *createArgs) springCreate() {
	var m model
	m.initMenu()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	m.downloadFile()
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

type listArgs struct {
	user *user.User
}

func (l *listArgs) projectList() {
	files, err := ioutil.ReadDir(l.user.HomeDir + projectDir)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("NAME                         SAVED_DATE         SIZE\n")

	for _, f := range files {
		y, m, d := f.ModTime().Date()
		fmt.Printf("%-28s %s           %d\n", f.Name(), fmt.Sprintf("%d/%d/%d", y, m, d), f.Size())
	}
}

func (p *pclt) init() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	saveCmd := flag.NewFlagSet("save", flag.ExitOnError)

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error: %v", currentUser)
	}

	if len(os.Args) < 2 {
		log.Fatal("expected subcommands.")
	}

	switch os.Args[1] {
	case "create":
		var create createArgs
		create.cmd = createCmd
		create.user = currentUser
		create.init()

		switch create.project {
		case "spring-init":
			create.springCreate()
		default:
			create.defaultCreate()
		}

	case "save":
		var save saveArgs
		save.cmd = saveCmd
		save.user = currentUser
		save.init()
		save.defaultSave()

	case "list":
		var list listArgs
		list.user = currentUser
		list.projectList()

	case "-help":
		help()

	default:
		log.Fatalln("Error: No such subcommand does not exist")
	}
}

func main() {
	var p pclt
	p.init()
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
