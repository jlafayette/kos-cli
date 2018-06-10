// WAIT FOR {{.OrbPart}}
clearscreen. print "Waiting for {{.OrbPart}}...".
warpto(time:seconds + eta:{{.OrbPart}}).
set warp to 0. wait until kuniverse:timewarp:issettled.
