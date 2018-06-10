// delete all files from volume 1.
list files in allfiles.
for f in allfiles { deletepath("1:/" + f). }
copypath("{{.Load}}", "1:/").
wait 1. clearscreen.