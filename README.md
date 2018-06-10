# kos-cli
CLI for helping with KSP-KOS development

### About

This command line tool helps with KerboScript development for the [Kerbal Space
Program](https://kerbalspaceprogram.com) game. [kOS: Kerbal Operating 
System](https://ksp-kos.github.io/KOS/index.html) is a mod that adds a
scripting language to automate missions.

### Build
```cmd
go build -o kos.exe -v .
```

### Commands

`kos env` : Display the environement variables that are used by the kos command.

`kos deploy` : Copies all .ks files in these subfolders of KSPSRC to the directory defined by KSPSCRIPT.

- $KSPSRC/library/*.ks -> $KSPSCRIPT/*.ks
- $KSPSRC/boot/*.ks -> $KSPSCRIPT/boot/*.ks
- $KSPSRC/missions/*.ks -> $missions/*.ks
- $KSPSRC/working/boot/*.ks -> $KSPSCRIPT/boot/*.ks
- $KSPSRC/working/missions/*.ks -> $KSPSCRIPT/missions/*.ks

`kos build` : Build runnable KerboScript missions (consist of a boot file and a mission file) from 
section templates. Try `kos build -help` for more information.

### Environment

These environment variables are used as the default arguments in many of the subcommands,
so you don't have to pass them in as flags every time.

**KSPSCRIPT** : The `Ships/Script` subfolder of the KSP install directory. This is where .ks files are deployed to.

**KSPSRC** : This is the directory of the development project that should be deployed to KSPSCRIPT.

**KSPTEMPLATE** : This is where the template .ks files are located that are used by the kos build subcommand.

Windows Example:

```cmd
set "KSPSCRIPT=G:\Games\Steam\steamapps\common\Kerbal Space Program\Ships\Script"
set "KSPSRC=G:\kerboscripting"
set "KSPTEMPLATE=G:\go\src\github.com\jlafayette\kos-cli\templates"
```
