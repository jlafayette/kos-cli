// FIX COLLISION COURSE
clearscreen.
if ship:obt:periapsis < {{.TgtPeriapsis}} {
    print "Uh oh! You are on a collision course with {{.Body}}.". wait 3.
    print "Calculating alternate trajectory, please stand by...".
    lock normalV to vcrs(ship:velocity:orbit, -body:position).
    lock radialOutV to vcrs(ship:velocity:orbit, -normalV).
    lock steering to radialOutV. wait 5.
    set tval to 1.
    until 0 {
        if ship:obt:periapsis > {{.TgtPeriapsis}} {
            break.
        }
        wait 0.01.
    }
    set tval to 0.
    unlock steering.
    wait 1.
    clearscreen.
    print "Course correction completed.".
}