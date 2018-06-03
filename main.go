package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jlafayette/kos-cli/build"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: kos <command> [<args>]")
		fmt.Println("Commands: ")
		fmt.Println(" build    Create customized mission from templates")
		fmt.Println(" deploy   Copy files from development folder to KSP Ships/Script path")
		fmt.Println(" env      Display environment variables used by kos command")
		return
	}

	deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)
	verboseFlag := deployCommand.Bool("v", false, "Display verbose output")

	envCommand := flag.NewFlagSet("env", flag.ExitOnError)
	// // Setting environment variables doesn't change environment where the script was called from
	// // Either figure out how to do this or just have the user set them before running deploy.
	// setKspScriptFlag := envCommand.String("kspscript", "", "Path to deploy to, usually Ships/Script under KSP install")
	// setKspSourceFlag := envCommand.String("kspsource", "", "Development root path")

	buildCommand := flag.NewFlagSet("build", flag.ExitOnError)
	missionFlag := buildCommand.String("mission", "a_mission.ks", "Name of the mission")
	argsFlag := buildCommand.String("args", "", "Comma separated args for the mission script")

	kspscript := os.Getenv("KSPSCRIPT")
	kspsrc := os.Getenv("KSPSRC")

	switch os.Args[1] {
	case "deploy":
		deployCommand.Parse(os.Args[2:])
	case "env":
		envCommand.Parse(os.Args[2:])
	case "build":
		buildCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
	if deployCommand.Parsed() {
		err := checkEnv()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = deploy(kspsrc, kspscript, *verboseFlag)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if envCommand.Parsed() {
		fmt.Println("KSPSCRIPT:", kspscript)
		fmt.Println("KSPSRC:", kspsrc)
	}
	if buildCommand.Parsed() {
		err := checkEnv()
		if err != nil {
			fmt.Println(err)
			return
		}
		args := strings.Split(*argsFlag, ",")
		if len(args) == 1 {
			if args[0] == "" {
				// If only an empty string, convert to nil.
				// This is because to make the template condition work and not
				// add add a ', ' to the end.
				args = nil
			}
		}
		mission := build.Boot{Filename: *missionFlag, Args: args}
		build.Test(os.Stdout, kspsrc, &mission)
	}
}

func checkEnv() error {
	kspscript := os.Getenv("KSPSCRIPT")
	kspsrc := os.Getenv("KSPSRC")
	if kspscript == "" {
		return errors.New("Missing required environment variable: KSPSCRIPT")
	}
	if kspsrc == "" {
		return errors.New("Missing required environment variable: KSPSRC")
	}
	return nil
}
