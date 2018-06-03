// CIRCULARIZE
copypath("0:/circularize.ks", "1:/").
runpath("circularize.ks", {{.TgtDir}}).
deletepath("1:/circularize.ks").
wait 1. clearscreen.
