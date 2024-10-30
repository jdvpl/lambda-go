git add .
git commit -m "updating data"
git push
set GOOS=linux
set GOARCH=amd64
go build main.go
rm main.zip
zip main.zip main