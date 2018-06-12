{{/*
.TgtAlt
.TgtVessel
.Offset
*/}}

copypath("0:/f_remap.ks", "1:/"). runpath("1:/f_remap.ks").
copypath("0:/f_tgt.ks", "1:/"). runpath("1:/f_tgt.ks").
copypath("0:/f_eta.ks", "1:/"). runpath("1:/f_eta.ks").

// raise apoapsis to TgtAlt
clearscreen. print "Raising apoapsis to {{.TgtAlt}}".
lock steering to ship:prograde. wait 3.
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
    print "        diff: " + round(diff,2) at (5, 2).
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

set tl to tgt_ves:longitude. // -180 180
set sl to ship:longitude.    // -180 180

// At apoapsis, calculate the desired offset to the TgtVessel
set diff to 360 - tgt_angle(tgt_ves).
set new_period to tgt_ves:obt:period * (1 - (diff/360)).

// if diff is 0 done.
// if diff is 360, done.
// if diff is 

if tgt_angle(tgt_ves) < 180 { // target is in front of ship.
    set diff to -tgt_angle(tgt_ves).
} else { // target is behind ship.
    set diff to 360 - tgt_angle(tgt_ves).
    set new_period to tgt_ves:obt:period * (1 + (diff/360)).
}
    // This needs to be expressed in terms of Kerbin Longitude
    // Ship:longitude is 0, tgt:longitude is 70, meaning tgt offset is + 70...
    // Desired offset is 360*{{.Offset}}, for example 90.

// Get the current offset (in Longitude)
// Get the difference
// Raise the periapsis so that the orbital period = TgtVessel:obt:period - offset
// Wait for 1 orbital period
// Burn prograde until ship:obt:period = TgtVessel:obt:period

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
    print "tgt altitude: " + round({{.TgtAlt}},2) at (5, 3).
    print "        diff: " + round(diff,2) at (5, 4).
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
