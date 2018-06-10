// No patched conics needed

// Get time to eta:apoapsis so that it goes negative when ship has moved past it.
function eta_error {
    local err is eta:apoapsis - ship:obt:period.
    if eta:apoapsis < ship:obt:period/2 {
        err is eta:apoapsis.
    }
    return err
}

// raise apoapsis to TgtAlt
lock steering to ship:prograde.
set orig_diff to {{.TgtAlt}} - ship:obt:apoapsis.
until 0 {
    local diff to {{.TgtAlt}} - ship:obt:apoapsis.
    if diff < 0 {
        break.
    }
    local threshhold is {{.TgtAlt}} / 4.  // Distance away from tgt to start throttling down.
    if diff < threshhold {
        set tval to remap(diff, 0, threshhold, .01, 1). // input inputLow inputHigh outputLow outputHigh
    } else {
        set tval to 1.
    }
    autostage().
    lock throttle to tval.
    wait 0.01.
}
unlock steering.
// WAIT FOR APOAPSIS
clearscreen. print "Waiting for APOAPSIS...".
local buffer to 30.
warpto(time:seconds + (eta:apoapsis - buffer)).
set warp to 0. wait until kuniverse:timewarp:issettled.
lock steering to ship:prograde.

{{if .TgtVessel}}
// Add calculation to space out orbits.
{{else}}// If no vessel
// raise the periapsis to match apoapsis
clearscreen. print "Raising periapsis".

until eta:apoapsis < 7 {
    print "eta:apoapsis: " + round(eta_error(),2) at (5, 2).
}
set orig_diff to {{.TgtAlt}} - ship:obt:periapsis.
lock throttle to tval.
print "before thrust loop...".
until 0 {
    local diff to {{.TgtAlt}} - ship:obt:periapsis.
    print "eta:apoapsis: " + round(eta_error(),2) at (5, 2).
    print "tgt altitude: " + round({{.TgtAlt}},2) at (5, 2).
    print "        diff: " + round(diff,2) at (5, 2).

    if diff < 0 {
        break.
    }
    if (eta:periapsis < (ship:obt:period/2 - ship:obt:period/4)) or 
       (eta:periapsis > (ship:obt:period/2 + ship:obt:period/4)) {
        break.
    }
    local threshhold is {{.TgtAlt}} / 2.
    if diff < threshhold {
        set tval to remap(diff, 0, threshhold, .01, 1). // input inputLow inputHigh outputLow outputHigh
    } else {
        set tval to 1.
    }
    autostage().
    wait 0.01.
}
print "Done!".
{{end}}
unlock steering.
set tval to 0.
set throttle to 0.
unlock throttle.
set ship:control:pilotmainthrottle to 0.
