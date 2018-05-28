package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: kos <command> [<args>]")
		fmt.Println("Commands: ")
		fmt.Println(" deploy   Copy files from development folder to KSP Ships/Script path")
		fmt.Println(" env      Display environment variables used by kos command")
		return
	}

	deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)

	envCommand := flag.NewFlagSet("env", flag.ExitOnError)
	// // Setting environment variables doesn't change environment where the script was called from
	// // Either figure out how to do this or just have the user set them before running deploy.
	// setKspScriptFlag := envCommand.String("kspscript", "", "Path to deploy to, usually Ships/Script under KSP install")
	// setKspSourceFlag := envCommand.String("kspsource", "", "Development root path")

	switch os.Args[1] {
	case "deploy":
		deployCommand.Parse(os.Args[2:])
	case "env":
		envCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
	if deployCommand.Parsed() {
		fmt.Println("Testing kos deploy command...")
		if os.Getenv("KSPSCRIPT") == "" {
			fmt.Println("Missing required environment variable: KSPSCRIPT")
			return
		}
		if os.Getenv("KSPSRC") == "" {
			fmt.Println("Missing required environment variable: KSPSRC")
			return
		}
		deploy(os.Getenv("KSPSRC"), os.Getenv("KSPSCRIPT"))
	}
	if envCommand.Parsed() {
		fmt.Println("KSPSCRIPT:", os.Getenv("KSPSCRIPT"))
		fmt.Println("KSPSRC:", os.Getenv("KSPSRC"))
	}
}
