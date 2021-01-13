package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

func chkErr(e error) {
	if e != nil {
		fmt.Println("[Config][Terraform][Error]: ", e)
	}
}

func (c *ProvisionContextCommand) Run(args []string) int {
	fmt.Println("[Config][Terraform]: Provision Starting")
	directoryPath := filepath.Join("/home/dat-vu/Projects/TerraformTest")
	fmt.Println(directoryPath)
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		fmt.Println("[Config][Terraform][Info]: Directory does not exist! Creating a directory for the project now in ", directoryPath)
		err1 := os.MkdirAll(directoryPath, os.ModePerm)
		chkErr(err1)
		fmt.Println("[Config][Terraform][Info]: mkdir Done")
	}
	configFilePath := directoryPath + "/config.tf"
	fmt.Println("[Config][Terraform][Info]: Creating config file in ", configFilePath)
	configString := `terraform { 
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

	f, _ := os.Create(configFilePath)
	f.WriteString(configString)
	f.Close()
	fmt.Println("[Config][Terraform][Info]: Config file created!")
	cmd := exec.Command("terraform", "init")
	cmd.Dir = directoryPath
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
	}

	if err1 := cmd.Wait(); err1 != nil {
		log.Printf("Cmd returned error: %v", err1)
	} else {
		fmt.Println("[Config][Terraform][Info]: terraform init Done")
		cmd = exec.Command("terraform", "apply", "-auto-approve")
		cmd.Dir = directoryPath
		err2 := cmd.Run()
		chkErr(err2)
		fmt.Println("[Config][Terraform][Info]: terraform init apply Done")
	}
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
