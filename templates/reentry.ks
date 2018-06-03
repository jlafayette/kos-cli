// REENTRY
clearscreen.
print "Preparing for re-entry.".
lock steering to ship:prograde. wait 4.
stage. wait 2.
print "Added drag chute trigger...".
when ((ship:airspeed < 420) and (alt:radar < {{.DragchuteAlt}})) then {
    print "Deploying drag chutes.".
    stage.
}
print "Added parachute trigger.".
when ((ship:airspeed < 250) and (alt:radar < {{.ParachuteAlt}})) then {
    print "Deploying parachutes.".
    stage.
}
until ship:airspeed < .5 {
    print "ALT:RADAR: " + round(alt:radar, 2) + "    " at (5, 5).
    lock steering to ship:srfretrograde.
    wait 0.1.
}
unlock steering.
clearscreen.
print "Finished mission script.".
