env GOOS=darwin GOARCH=amd64 go build -v CSSimulator.go
mv -f CSSimulator _Exe/mac
rm _Exe/CSSimulator_mac.zip
zip -r _Exe/CSSimulator_mac.zip _Exe/mac/*

env GOOS=linux GOARCH=amd64 go build -v CSSimulator.go
mv -f CSSimulator _Exe/linux
rm _Exe/CSSimulator_linux.zip
zip -r _Exe/CSSimulator_linux.zip _Exe/linux/*

env GOOS=windows GOARCH=amd64 go build -v CSSimulator.go
mv -f CSSimulator.exe _Exe/win
rm _Exe/CSSimulator_win.zip
zip -r _Exe/CSSimulator_win.zip _Exe/win/*
