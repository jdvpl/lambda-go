git add .
git commit -m "update"
git push
set GOOS=linux
set GOARCH=amd64
go build main.go
rm main.zip
zip main.zip main
