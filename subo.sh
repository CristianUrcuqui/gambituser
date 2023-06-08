git add .
git commit -m "Ultimo Commit"
git push

set GOOS=linux 
set GOARCH=amd64 

go build main.go
rm -f main.zip
zip main.zip main