echo --------- Building MAC files....
env GOOS=darwin GOARCH=amd64 go build -v CSSimulator.go
echo Copying JSON files...
cp config.json _Exe/mac
cp data.json _Exe/mac
cp CSSimulator /Users/james/Dropbox/Work/CodeForSanJose/Open311/_test/CitySourced/cfg1/CitySourcedAPI
cp CSSimulator /Users/james/Dropbox/Work/CodeForSanJose/Open311/_test/CitySourced/cfg2/CitySourcedAPI
cp CSSimulator /Users/james/Dropbox/Work/CodeForSanJose/Open311/_test/CitySourced/cfg3/CitySourcedAPI
mv -f CSSimulator _Exe/mac
rm _Exe/CSSimulator_mac.zip
zip -r _Exe/CSSimulator_mac.zip _Exe/mac/*

echo --------- Building Linux files....
env GOOS=linux GOARCH=amd64 go build -v CSSimulator.go
cp config.json _Exe/linux
cp data.json _Exe/linux
mv -f CSSimulator _Exe/linux
rm _Exe/CSSimulator_linux.zip
zip -r _Exe/CSSimulator_linux.zip _Exe/linux/*

echo --------- Building Windows files....
env GOOS=windows GOARCH=amd64 go build -v CSSimulator.go
cp config.json _Exe/win
cp data.json _Exe/win
mv -f CSSimulator.exe _Exe/win
rm _Exe/CSSimulator_win.zip
zip -r _Exe/CSSimulator_win.zip _Exe/win/*
