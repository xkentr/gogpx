/*
Copyright 2015 Kent Ryhorchuk
*/

package main

import (
  //"bufio"
  "fmt"
  "os"
  "flag"
  "time"
  "github.com/xkentr/gogpx"
  "encoding/xml"
  //"strings"
  //"regexp"
  //"net/url"
)

var gpxfile = flag.String("gpxfile", "file.gpx", "Input file path")
var csvfile = flag.String("csvfile", "file.csv", "Output file path")

func main() {
  flag.Parse()

  gpx, err := os.Open(*gpxfile)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  defer gpx.Close()

  csv, err := os.Create(*csvfile)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  defer csv.Close()

  loc, _ := time.LoadLocation("America/Los_Angeles")

  fmt.Fprintf(csv, "Time,Lat,Lon,Ele,Temp,HR\n")
  decoder := xml.NewDecoder(gpx)
  for {
      // Read tokens from the XML document in a stream.
      t, _ := decoder.Token()
      if t == nil {
        break
      }

      // Inspect the type of the token just read.
      switch se := t.(type) {
      case xml.StartElement:
        // If we just read a StartElement token
        inElement := se.Name.Local
        //fmt.Printf("Element: %v\n", inElement)
        if inElement == "trkpt" {
          var pt gogpx.Trkpt
          // decode a whole chunk of following XML into the
          // variable p which is a Page (se above)
          decoder.DecodeElement(&pt, &se)
          t, _ := time.Parse(time.RFC3339, pt.Time)
          t = t.In(loc)
          fmt.Fprintf(csv, "%v,%v,%v,%v,%v,%v\n", t.Format(time.RFC3339), pt.Lat, pt.Lon, pt.Ele, pt.Ext.Gpx.Atemp, pt.Ext.Gpx.Hr)
        }
      default:
      }

    }

}
