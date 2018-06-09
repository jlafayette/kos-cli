// APPROACH
copypath("0:/approach.ks", "1:/").
runpath("approach.ks", "{{.TgtVessel}}", {{.TgtRange}}, {{.OffsetVector}}, {{.Agility}}).
deletepath("1:/approach.ks").
