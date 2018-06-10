{{if .LKO}}// DEORBIT
// This part assumes the ship is on the last stage before parachutes.
// Also assumes that ship is orbiting Kerbin and has enough fuel to lower 
// periapsis to 30000.
clearscreen.
print "Deorbiting...".
wait 10.
copypath("0:/deorbit.ks", "1:/").
runpath("deorbit.ks").
deletepath("1:/deorbit.ks").
{{else}}clearscreen. 
print "Lowering periapsis...". lock steering to retrograde. wait 10.
until ship:obt:periapsis < {{.TgtPeriapsis}} {
    set tval to remap(ship:obt:periapsis, {{.TgtPeriapsis}}, 250000, .05, 1).
    autostage().
    lock throttle to tval.
    wait 0.01.
} lock throttle to 0. unlock steering.
// WAIT
print "Waiting for reentry...".
wait until ship:altitude < 250000.
set warp to 2.
wait until ship:altitude < 100000.
set warp to 0. wait 5.
// BURN TIL PERIAPSIS < 30km or out of fuel
clearscreen. print "Burning to lower periapsis to < 30km".
lock steering to ship:retrograde. wait 5.
until 0 {
    if ship:obt:periapsis < 30000 { break. }
    if ship:liquidfuel < 1 { break. }
    lock throttle to 1.
    wait .01.
} lock throttle to 0. unlock steering.{{end}}
