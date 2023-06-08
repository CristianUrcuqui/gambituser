git add .
git commit -m "Ultimo Commit"
git push

export GOOS=linux 
export GOARCH=amd64 

go build -o main
rm -f main.zip
zip main.zip main