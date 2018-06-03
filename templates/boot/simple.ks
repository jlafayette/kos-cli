{{/*This file is formatted to work with golang text/template package
Get the mission building tool here: https://github.com/jlafayette/kos-cli 
*/}}clearscreen.
print "Booting up...".
core:doevent("Open Terminal").
set ship:control:pilotmainthrottle to 0. wait 1.
if ship:altitude < 500 and ship:obt:body = Kerbin and ship:airspeed < 1 {
    print "Initializing mission sequence...". wait 1.
    copypath("0:/missions/{{.Filename}}", "1:/").
    runpath("1:/{{.Filename}}"{{if .Args}}{{range .Args}}, {{.}}{{end}}{{end}}).
    deletepath("1:/{{.Filename}}").
}
