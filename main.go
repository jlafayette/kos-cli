package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jlafayette/kos-cli/build"
	"github.com/jlafayette/kos-cli/deploy"
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

	// Environment variables are the defaults for many flags and displayed by env subcommand.
	kspscript := os.Getenv("KSPSCRIPT")
	kspsrc := os.Getenv("KSPSRC")
	ksptemplate := os.Getenv("KSPTEMPLATE")

	deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)
	verboseFlag := deployCommand.Bool("v", false, "Display verbose output")

	envCommand := flag.NewFlagSet("env", flag.ExitOnError)

	// kos build [-n|-name] LKO [-o|-out] "output\dir" [-p|-plan] "plan\file" [-t|-template] "templates\dir"
	// 			 "a_mission"	KSPSCRIPT									  KSPTEMPLATE
	buildCommand := flag.NewFlagSet("build", flag.ExitOnError)
	nameFlag := buildCommand.String("name", "", "Name of the mission")
	name2Flag := buildCommand.String("n", "a_mission.ks", "Alias for name")
	outFlag := buildCommand.String("out", "", "Directory to write mission and boot files to")
	out2Flag := buildCommand.String("o", kspscript, "Alias for out")
	planFlag := buildCommand.String("plan", "", "PLAN file to build mission from")
	plan2Flag := buildCommand.String("p", "", "Alias for plan")
	templateFlag := buildCommand.String("template", "", "Directory containing template files")
	template2Flag := buildCommand.String("t", ksptemplate, "Alias for template")

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
		err = deploy.Deploy(kspsrc, kspscript, *verboseFlag)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if envCommand.Parsed() {
		fmt.Println("KSPSCRIPT:", kspscript)
		fmt.Println("KSPSRC:", kspsrc)
		fmt.Println("KSPTEMPLATE:", ksptemplate)
	}

	// kos build [-n|-name] LKO [-o|-out] "output\dir" [-p|-plan] "plan\file" [-t|-template] "templates\dir"
	// 			 "a_mission"	KSPSCRIPT									  KSPTEMPLATE

	// Fall through -name --> -n --> -n(default)
	if buildCommand.Parsed() {

		var name string
		if *nameFlag != "" {
			name = *nameFlag
		} else {
			name = *name2Flag
		}
		var out string
		if *outFlag != "" {
			out = *outFlag
		} else {
			out = *out2Flag
		}
		var template string
		if *templateFlag != "" {
			template = *templateFlag
		} else {
			template = *template2Flag
		}
		var plan string
		if *planFlag != "" {
			plan = *planFlag
		} else if *plan2Flag != "" {
			plan = *plan2Flag
		} else {
			fmt.Println("missing required flag: -plan|-p")
			os.Exit(2)
			return // is this needed?
		}
		fmt.Printf("plan: %v\n", plan)

		// TODO: Validate arg values (paths should exist, name cannot contain quote char)

		bootTemplate := filepath.Join(template, "boot", "simple.ks")
		bf, err := os.Create(filepath.Join(out, "boot", name+".ks"))
		if err != nil {
			fmt.Printf("Error opening output boot file for writing, %s\n", err)
			return
		}
		defer bf.Close()
		err = build.Boot(bf, bootTemplate, name)
		if err != nil {
			fmt.Printf("Error writing boot file, %s\n", err)
			return
		}
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
