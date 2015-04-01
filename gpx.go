/*
Copyright 2015 Kent Ryhorchuk
*/

package gogpx

//<trkpt lon="-122.20216054469347" lat="37.3474077321589">
//  <ele>365.3999938964844</ele>
//  <time>2015-03-08T22:05:57.000Z</time>
//  <extensions>
//    <gpxtpx:TrackPointExtension>
//      <gpxtpx:atemp>25.0</gpxtpx:atemp>
//      <gpxtpx:hr>79</gpxtpx:hr>
//    </gpxtpx:TrackPointExtension>
//  </extensions>
//</trkpt>

type TrackPointExtension struct {
  Atemp float32 `xml:"atemp"`
  Hr float32 `xml:"hr"`
}

type Extensions struct {
  Gpx TrackPointExtension `xml:"TrackPointExtension"`
}

type Trkpt struct {
  Lon float64 `xml:"lon,attr"`
  Lat float64 `xml:"lat,attr"`
  Time string `xml:"time"`
  Ele float32 `xml:"ele"`
  Ext Extensions `xml:"extensions"`
}
