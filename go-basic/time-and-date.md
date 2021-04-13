# Time and Date

## Format a time or date \[complete guide\]

### Basic example <a id="basic-example"></a>

Go doesn’t use yyyy-mm-dd layout to format a time. Instead, you format a special **layout parameter**

`Mon Jan 2 15:04:05 MST 2006`

the same way as the time or date should be formatted. \(This date is easier to remember when written as `01/02 03:04:05PM ‘06 -0700`.\)

```text
const (
    layoutISO = "2006-01-02"
    layoutUS  = "January 2, 2006"
)
date := "1999-12-31"
t, _ := time.Parse(layoutISO, date)
fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
fmt.Println(t.Format(layoutUS)) // December 31, 1999
```

The function

* [`time.Parse`](https://golang.org/pkg/time/#Parse) parses a date string, and
* [`Format`](https://golang.org/pkg/time/#Time.Format) formats a [`time.Time`](https://golang.org/pkg/time/#Time).

They have the following signatures:

```text
func Parse(layout, value string) (Time, error)
func (t Time) Format(layout string) string
```

### Standard time and date formats <a id="standard-time-and-date-formats"></a>

| Go layout | Note |
| :--- | :--- |
| `January 2, 2006` | Date |
| `01/02/06` |  |
| `Jan-02-06` |  |
| `15:04:05` | Time |
| `3:04:05 PM` |  |
| `Jan _2 15:04:05` | Timestamp |
| `Jan _2 15:04:05.000000` | with microseconds |
| `2006-01-02T15:04:05-0700` | [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) \([RFC 3339](https://www.ietf.org/rfc/rfc3339.txt)\) |
| `2006-01-02` |  |
| `15:04:05` |  |
| `02 Jan 06 15:04 MST` | [RFC 822](https://www.ietf.org/rfc/rfc822.txt) |
| `02 Jan 06 15:04 -0700` | with numeric zone |
| `Mon, 02 Jan 2006 15:04:05 MST` | [RFC 1123](https://www.ietf.org/rfc/rfc1123.txt) |
| `Mon, 02 Jan 2006 15:04:05 -0700` | with numeric zone |

The following predefined date and timestamp [format constants](https://golang.org/pkg/time/#pkg-constants) are also available.

```text
ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700"
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
// Handy time stamps.
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"
```

### Layout options <a id="layout-options"></a>

| Type | Options |
| :--- | :--- |
| Year | `06`   `2006` |
| Month | `01`   `1`   `Jan`   `January` |
| Day | `02`   `2`   `_2`   \(width two, right justified\) |
| Weekday | `Mon`   `Monday` |
| Hours | `03`   `3`   `15` |
| Minutes | `04`   `4` |
| Seconds | `05`   `5` |
| ms μs ns | `.000`   `.000000`   `.000000000` |
| ms μs ns | `.999`   `.999999`   `.999999999`   \(trailing zeros removed\) |
| am/pm | `PM`   `pm` |
| Timezone | `MST` |
| Offset | `-0700`   `-07`   `-07:00`   `Z0700`   `Z07:00` |

### Corner cases <a id="corner-cases"></a>

It’s not possible to specify that an hour should be rendered without a leading zero in a 24-hour time format.

It’s not possible to specify midnight as `24:00` instead of `00:00`. A typical usage for this would be giving opening hours ending at midnight, such as `07:00-24:00`.

It’s not possible to specify a time containing a leap second: `23:59:60`. In fact, the time package assumes a Gregorian calendar without leap seconds.

## Time zones



Each [`Time`](https://golang.org/pkg/time/#Time) has an associated [`Location`](https://golang.org/pkg/time/#Location), which is used for display purposes.

The method [`In`](https://golang.org/pkg/time/#Time.In) returns a time with a specific location. Changing the location in this way changes only the presentation; it does not change the instant in time.

Here is a convenience function that changes the location associated with a time.

```text
// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
    loc, err := time.LoadLocation(name)
    if err == nil {
        t = t.In(loc)
    }
    return t, err
}
```

In use:

```text
for _, name := range []string{
	"",
	"Local",
	"Asia/Shanghai",
	"America/Metropolis",
} {
	t, err := TimeIn(time.Now(), name)
	if err == nil {
		fmt.Println(t.Location(), t.Format("15:04"))
	} else {
		fmt.Println(name, "<time unknown>")
	}
}
```

```text
UTC 19:32
Local 20:32
Asia/Shanghai 03:32
America/Metropolis <time unknown>
```

> **Warning:** A daylight savings time transition skips or repeats times. For example, in the United States, March 13, 2011 2:15am never occurred, while November 6, 2011 1:15am occurred twice. In such cases, the choice of time zone, and therefore the time, is not well-defined. Date returns a time that is correct in one of the two zones involved in the transition, but it does not guarantee which.[Package time: Date](https://golang.org/pkg/time/#Date)

## How to get current timestamp

Use [`time.Now`](https://golang.org/pkg/time/#Now) and one of [`time.Unix`](https://golang.org/pkg/time/#Time.Unix) or [`time.UnixNano`](https://golang.org/pkg/time/#Time.UnixNano) to get a timestamp.

```text
now := time.Now()      // current local time
sec := now.Unix()      // number of seconds since January 1, 1970 UTC
nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

fmt.Println(now)  // time.Time
fmt.Println(sec)  // int64
fmt.Println(nsec) // int64
```

```text
2009-11-10 23:00:00 +0000 UTC m=+0.000000000
1257894000
1257894000000000000
```

## Get year, month, day from time

he [`Date`](https://golang.org/pkg/time/#Time.Date) function returns the year, month and day of a [`time.Time`](https://golang.org/pkg/time/#Time).

```text
func (t Time) Date() (year int, month Month, day int)
```

In use:

```text
year, month, day := time.Now().Date()
fmt.Println(year, month, day)      // For example 2009 November 10
fmt.Println(year, int(month), day) // For example 2009 11 10
```

You can also extract the information with seperate calls:

```text
t := time.Now()
year := t.Year()   // type int
month := t.Month() // type time.Month
day := t.Day()     // type int
```

The [`time.Month`](https://golang.org/pkg/time/#Month) type specifies a month of the year \(January = 1, …\).

```text
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```

## How to find the day of week

The [`Weekday`](https://golang.org/pkg/time/#Time.Weekday) function returns returns the day of the week of a [`time.Time`](https://golang.org/pkg/time/#Time).

```text
func (t Time) Weekday() Weekday
```

In use:

```text
weekday := time.Now().Weekday()
fmt.Println(weekday)      // "Tuesday"
fmt.Println(int(weekday)) // "2"
```

### Type Weekday <a id="type-weekday"></a>

The [`time.Weekday`](https://golang.org/pkg/time/#Weekday) type specifies a day of the week \(Sunday = 0, …\).

```text
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```

## Days between two dates

```text
func main() {
    // The leap year 2016 had 366 days.
    t1 := Date(2016, 1, 1)
    t2 := Date(2017, 1, 1)
    days := t2.Sub(t1).Hours() / 24
    fmt.Println(days) // 366
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
```

## Days in a month

To compute the last day of a month, you can use the fact that [`time.Date`](https://golang.org/pkg/time/#Date) accepts values outside their usual ranges – the values are normalized during the conversion.

To compute the number of days in February, look at the day before March 1.

```text
func main() {
    t := Date(2000, 3, 0) // the day before 2000-03-01
    fmt.Println(t)        // 2000-02-29 00:00:00 +0000 UTC
    fmt.Println(t.Day())  // 29
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
```

[`AddDate`](https://golang.org/pkg/time/#Time.AddDate) normalizes its result in the same way. For example, adding one month to October 31 yields December 1, the normalized form of November 31.

```text
t = Date(2000, 10, 31).AddDate(0, 1, 0) // a month after October 31
fmt.Println(t)                          // 2000-12-01 00:00:00 +0000 UTC
```

## Measure execution time

### Measure a piece of code <a id="measure-a-piece-of-code"></a>

```text
start := time.Now()
// Code to measure
duration := time.Since(start)

// Formatted string, such as "2h3m0.5s" or "4.503μs"
fmt.Println(duration)

// Nanoseconds as int64
fmt.Println(duration.Nanoseconds())
```

### Measure a function call <a id="measure-a-function-call"></a>

You can track the execution time of a complete function call with this one-liner, which logs the result to the standard error stream.

```text
func foo() {
    defer duration(track("foo"))
    // Code to measure
}
```

```text
func track(msg string) (string, time.Time) {
    return msg, time.Now()
}

func duration(msg string, start time.Time) {
    log.Printf("%v: %v\n", msg, time.Since(start))
}
```

### Benchmarks <a id="benchmarks"></a>

The [`testing`](https://golang.org/pkg/testing/) package has support for benchmarking that can be used to examine the performance of your code.

