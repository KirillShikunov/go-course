package command

import (
	"12_keychain/internal/store"
	"flag"
	"fmt"
)

type Command interface {
	Execute() error
}

type list struct {
	cmd *flag.FlagSet
}

func (l *list) Execute() error {
	if err := l.cmd.Parse(flag.Args()[1:]); err != nil {
		return fmt.Errorf("error parsing list command: %w", err)
	}

	passwords := store.ListPasswords()
	for _, name := range passwords {
		fmt.Println(name)
	}

	return nil
}

type save struct {
	cmd *flag.FlagSet
}

func (s *save) Execute() error {
	if err := s.cmd.Parse(flag.Args()[1:]); err != nil {
		return fmt.Errorf("error parsing save command: %w", err)
	}
	if s.cmd.NArg() != 1 {
		return fmt.Errorf("provide a name for the password")
	}

	name := s.cmd.Arg(0)
	fmt.Println("Enter password:")

	var password string
	if _, err := fmt.Scan(&password); err != nil {
		return err
	}

	if err := store.SavePassword(name, password); err != nil {
		return err
	}

	fmt.Println("Password saved")

	return nil
}

type get struct {
	cmd *flag.FlagSet
}

func (g *get) Execute() error {
	if err := g.cmd.Parse(flag.Args()[1:]); err != nil {
		return fmt.Errorf("error parsing get command: %w", err)
	}

	if g.cmd.NArg() != 1 {
		return fmt.Errorf("provide a name to retrieve the password")
	}

	name := g.cmd.Arg(0)
	password, err := store.GetPassword(name)

	if err != nil {
		return err
	}

	fmt.Println("Password:", password)

	return nil
}

func Create(s string, cmd *flag.FlagSet) (Command, error) {
	switch s {
	case "list":
		return &list{cmd}, nil
	case "save":
		return &save{cmd}, nil
	case "get":
		return &get{cmd}, nil
	default:
		return nil, fmt.Errorf("command `%s` not found", s)
	}
}
