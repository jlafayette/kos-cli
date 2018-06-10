// WAIT FOR TRANSFER
copypath("0:/f_tgt.ks", "1:/"). runpath("f_tgt.ks").
copypath("0:/f_remap.ks", "1:/"). runpath("f_remap.ks").

clearscreen.
print "Waiting for transfer...".
until close_enough(tgt_angle(Mun), 112, 2) {
    print "DIFF:  "+round(tgt_angle(Mun),2)+"    " at (5, 5).
    wait 1.
} set warp to 0. wait until kuniverse:timewarp:issettled.

// TRANSFER BURN
clearscreen.
lock steering to ship:prograde.
from {local i is 15.} until i = 0 step {set i to i - 1.} do {
    print "Tranfer burn in " + i.
    wait 1.
}
set orig_mun_diff to body("Mun"):apoapsis - ship:obt:apoapsis.
until 0 {
    local diff to body("Mun"):apoapsis - ship:obt:apoapsis.

    // If 0, the ship might be on collision course with the Mun
    local cutoff is {{.Cutoff}}. // 100_000
    
    if diff < cutoff {
        break.
    }
    set tval to 1.
    if diff < 4_000_000 {
        set tval to remap(diff, cutoff, 4_000_000, .05, 1). // input inputLow inputHigh outputLow outputHigh
    }
    autostage().
    lock throttle to tval.
    wait 0.01.
}
set tval to 0.
unlock steering.
