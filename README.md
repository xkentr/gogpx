# gogpx
GPX file manipulation in Go

Currently this is a very simple tool to transform Garmin generated gpx files to csv files for analysis in
spreadsheets or other tools.

## To Use

gpx2csv -gpxfile in.gpx -csvfile out.csv

## Example Output
```
Time,Lat,Lon,Ele,Temp,HR
2015-03-17T07:03:52-07:00,37.347424,-122.202185,370.4,0,83
2015-03-17T07:03:53-07:00,37.347424,-122.202185,370.4,17,82
2015-03-17T07:03:55-07:00,37.347426,-122.202204,370.4,17,82
2015-03-17T07:04:00-07:00,37.347458,-122.202337,370.2,17,83
2015-03-17T07:04:13-07:00,37.347687,-122.202714,369.6,17,105
2015-03-17T07:04:23-07:00,37.347787,-122.202972,372.2,17,117
```

## Simplified Build Instructions (Linux - Debian flavored)
1. Install go: sudo apt-get install golang
2. Make a directory to work in and cd into it
3. Set your GOPATH: export GOPATH=\`pwd\`
4. Build: go get github.com/xkentr/gogpx/gpx2csv

The gpx2csv executable will be in $GOPATH/bin . It does not require any other prerequisites.
