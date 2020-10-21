--1--
cd src/github.com/aleksaan/diskusage/tests

--2--
go test -coverprofile=cover.out -coverpkg "github.com/aleksaan/diskusage/diskusage"

--3--
go tool cover -html=cover.out -o cover.html

--PROFILING

cd tests

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

go tool pprof cpu.prof