// DEPLOY SOLAR PANELS
panels on.
// DEPLOY ANTENNA FOR COMMUNICATION
copypath("0:/extend_antenna.ks", "1:/").
runpath("extend_antenna.ks", "{{.Antenna}}").
deletepath("1:/extend_antenna.ks").
