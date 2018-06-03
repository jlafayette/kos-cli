// WAIT UNTIL CREW IS ABOARD
clearscreen.
until 0 {
    print "Waiting for crew to board." at (0, 1).
    if ship:crew():length > {{.TgtCrew}}-1 { break. }
    wait 1.
}
