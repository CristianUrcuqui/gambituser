git add .
git commit -m "Ultimo Commit"
git push

GOOS=linux GOARCH=amd64 

go build main.go && rm -f main.zip && zip main.zip main