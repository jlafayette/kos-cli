// WAIT FOR TARGET
copypath("0:/f_tgt.ks", "1:/"). runpath("f_tgt.ks"). // lng_to_deg, tgt_angle, close_enough
set tgt_name to "{{.TgtVessel}}".
set tgt_ves to vessel(tgt_name).
set tgt_angle to {{.TgtAngle}}.
clearscreen.
until 0 {
    print "      ship:longitude: " + round(ship:longitude,2) +    "      " at (0, 2).
    print "   tgt_ves:longitude: " + round(tgt_ves:longitude,2) + "      " at (0, 3).
    if close_enough(ship:longitude, tgt_ves:longitude, tgt_angle) {
        break.
    } else if close_enough(ship:longitude, tgt_ves:longitude, tgt_angle+10) {
        if warp > 2 { set warp to 2. }
    }
}
set warp to 0. wait until kuniverse:timewarp:issettled.
deletepath("1:/f_tgt.ks").
