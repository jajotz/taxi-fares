# TaxiFares
Program that gets taxi fares

## Specification
#### (i) Overview
1. The base fare is 400 yen for up to 1 km.
2. Up to 10 km, 40 yen is added every 400 meters.
3. Over 10 km, 40 yen is added every 350 meters.

This taxi equipped with the following two meters. Only one of the most recent real values is recorded on these meters.
- Distance Meter
- Fare Meter

#### (ii) Input Format
Distance meter records are sent line by line for standard input in the following format.
00:00:00.000 0.0 00:01:00.123 480.9 00:02:00.125 1141.2 00:03:00.100 1800.8

The specifications of the distance meter are as follows.
- Space-separated first half is elapsed time (Max 99:99:99.999), second half is mileage.(the unit is meters, Max: 99999999.9)
- It keeps only the latest values.
- It calculates and creates output of the mileage per minute, but an error of less than 1 second occurs.

#### (iii) Error Definition
Error occurs under the following conditions.
- Not under the format of, hh:mm:ss.fff<SPACE>xxxxxxxx.f<LF>, but under an improper format.
- Blank Line
- When the past time has been sent.
- The interval between records is more than 5 minutes apart.
- When there are less than two lines of data.
- When the total mileage is 0.0m.

#### (iv) Output
Display the current fare as an integer on the fare meter (standard output). 12345
Standard output displays nothing for incorrect inputs that do not meet specifications, the exit
code ends with a value other than 0.

### How to Run
#### Build
```
go build main.go
```
#### Run
```
go run main.go
```
Run from binary(Windows):
Build the apps then
```
./main.exe
```
Run from binary(Linux/Mac):
Build the apps then
```
./main
```

When the program run
Sample Input:
first line is n (number of data)
second line until EOF is the data
```
4
00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8
```
Sample Output:
```
Taxi Fare: 520
```


#### Stop the program
ctrl + c signal (Windows/Linux)
```
ctrl + c
```
cmd + c signal (Mac)
```
cmd + c
```

#### Run unit test
```
go test -v -coverprofile cover.out ./app
go test -v -coverprofile cover_util.out ./app
go tool cover -html=cover.out
go tool cover -html=cover_util.out
```
