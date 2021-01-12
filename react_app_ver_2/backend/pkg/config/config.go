package config

import (
	"os"
	"os/exec"
	"fmt"
)

//Commands is a map that hold all
//possible function with our CMP
var Commands map[string]CommandFactory

type CommandFactory func() Command

type Command interface {
	Run(args []string) int
}
type ProvisionContextCommand struct {
	Name string
}

func (c *ProvisionContextCommand) Run(args []string) int {
	cmd := exec.Command("cd")
	cmd.Run()
	fmt.Println("[Config]: cd Done")
	cmd = exec.Command("pwd")
	cmd.Run()
	fmt.Println("[Config]: pwd Done")
	path, _ := cmd.CombinedOutput()
	realPath := string(path) + "/Projects"
	if _, err := os.Stat(realPath); os.IsNotExist(err) {
		cmd = exec.Command("mkdir Projects")
		cmd.Run()
		fmt.Println("[Config]: mkdir Done")
	}
	cmd = exec.Command("cd Projects")
	cmd.Run()
	fmt.Println("[Config]: cd Done")
	cmd = exec.Command("touch config.tf")
	cmd.Run()
	fmt.Println("[Config]: touch Done")

	configString :=
		`terraform {
		required_providers {
		  aws = {
			source  = "hashicorp/aws"
			version = "~> 2.70"
		  }
		}
	}

	provider "aws" {
		profile = "default"
		region  = "us-west-2"
	}

	resource "aws_instance" "example" {
		ami           = "ami-830c94e3"
		instance_type = "t2.micro"
	}`

	f, _ := os.Create("config.tf")
	fmt.Println("[Config]: config Done")
	f.WriteString(configString)
	f.Close()
	cmd = exec.Command("terraform init")
	cmd.Run()
	fmt.Println("[Config]: terraform init Done")

	cmd = exec.Command("terraform apply")
	cmd.Run()
	fmt.Println("[Config]: terraform init apply Done")

	return 1
}

type DestroyContextCommand struct {
	Name string
}

func (c *DestroyContextCommand) Run(args []string) int {
	return 1
}

type ShowContextCommand struct {
	Name string
}

func (c *ShowContextCommand) Run(args []string) int {
	return 1
}

func InitCommands() {
	Commands = map[string]CommandFactory{
		"provision": func() Command {
			return &ProvisionContextCommand{
				Name: "provision",
			}
		},
		"destroy": func() Command {
			return &DestroyContextCommand{
				Name: "provision",
			}
		},
		"show": func() Command {
			return &ShowContextCommand{
				Name: "provision",
			}
		},
	}
}
