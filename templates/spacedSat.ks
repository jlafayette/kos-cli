{{/*
.TgtAlt
.TgtVessel
.Offset
*/}}

copypath("0:/f_remap.ks", "1:/"). runpath("1:/f_remap.ks").
copypath("0:/f_tgt.ks", "1:/"). runpath("1:/f_tgt.ks").
copypath("0:/f_eta.ks", "1:/"). runpath("1:/f_eta.ks").
copypath("0:/f_autostage.ks", "1:/"). runpath("1:/f_autostage.ks").

// change_obt_period // function even though it doesn't start with f_
copypath("0:/sync_orbits.ks", "1:/"). runpath("1:/sync_orbits.ks").

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
local buffer to 20.
warpto(time:seconds + (eta:apoapsis - buffer)).
until 0 {
    if eta_error() < buffer { break. }
    wait 10.
}
set warp to 0. wait until kuniverse:timewarp:issettled.
lock steering to ship:prograde.

{{if .TgtVessel}}
// Add calculation to space out orbits.
clearscreen. print "Calculating new orbital period...".
set tgt_ves to vessel("{{.TgtVessel}}").
set angle1 to tgt_angle(tgt_ves).

set desired_angle to {{.Offset}}.

// EXAMPLE
// If Offset is 90, then tgt should be 90 degrees ahead.
// all we can do is go faster by having a tighter orbit, so if tgt is 89 ahead, then error is 359
// 180 -> 90
// 91  -> 1
// 90  -> 0
// 89  -> 359
// 0   -> 270

set angle_diff to mod((angle1 + 360) - desired_angle, 360). // 0 - 360
set new_period to tgt_ves:obt:period * (1 - (angle_diff/360)).
set orb_count to 1.
until new_period > ship:obt:period {
    set orb_count to orb_count + 1.
    set new_period to tgt_ves:obt:period * (1 - ((angle_diff/360) / orb_count)).
}

wait until eta_error() < 13.
change_obt_period(new_period, 0.005).

// The full wait time of ship:obt:period * orb_count doesn't land perfectly at
// apoapsis for final burn, so subtract half an orbit and then track eta:apoapsis
clearscreen. print "Waiting for orbital sync.".
set waittime to time:seconds + ship:obt:period * orb_count - (ship:obt:period/2).
warpto(waittime).
until time:seconds > waittime {
    print " Orbital change in: " + round(time:seconds - waittime,0) at (5, 3).
    wait 1.
}
set warp to 0. wait until kuniverse:timewarp:issettled.
warpto(time:seconds + eta:apoapsis - 20).
set warp to 0. wait until kuniverse:timewarp:issettled.

wait until eta_error() < 11.  // to burn at apoapsis.
change_obt_period(tgt_ves:obt:period, 0.001).   // make this very precise.

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
{{end}}
unlock steering.
set tval to 0.
set throttle to 0.
unlock throttle.
set ship:control:pilotmainthrottle to 0.

clearscreen. print "Done!".
// DIAGNOSTICS
set angle1 to tgt_angle(tgt_ves).
if angle1 < 180 {
    print " Ahead by: " + round(angle1,2) at (5, 3).
} else { 
    print " Behind by: " + round((360 - angle1),2) at (5, 3).
}
print " ship:obt:period: " + round(ship:obt:period,4) at (5, 4).
print "  tgt:obt:period: " + round(tgt_ves:obt:period,4) at (5, 5).
wait 10.
