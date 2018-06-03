// LAUNCH
clearscreen.
copypath("0:/launch.ks", "1:/").
runpath("launch.ks", {{.TgtAlt}}, {{.TgtDir}}).
deletepath("1:/launch.ks").
