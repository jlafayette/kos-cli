// delete all files from volume 1.
list files in allfiles.
for f in allfiles { deletepath("1:/" + f). }
{{if .Load}}copypath("{{.Load}}", "1:/").{{end}}
wait 1. clearscreen.
