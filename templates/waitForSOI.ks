// WAIT UNTIL {{.Body}} SOI
clearscreen.
print "Waiting to enter {{.Body}} SOI...".
until 0 {
    if ship:body = {{.Body}} { break. }
    wait 10.
}
set warp to 0.
wait until kuniverse:timewarp:issettled.
clearscreen.
print "You are now flying by {{.Body}}!". wait 3.
