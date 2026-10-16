package booking
import "fmt"
 
import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
      fmt.Println(date)
	t, _ := time.Parse(layout, date)
    fmt.Println(t)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
		layout := "January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)
      fmt.Println(t)
    nowTime := time.Now()
    fmt.Println(t.Before(nowTime))
return t.Before(nowTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
			layout := "Monday, January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)
    if (t.Hour() >= 12 && t.Hour() <= 18) {
        return true
    } else {
        return false;
    }  
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date);
   d  := fmt.Sprintf("%d", t.Day())
    y := fmt.Sprintf("%d", t.Year())
     h := fmt.Sprintf("%d", t.Hour())
       m := fmt.Sprintf("%d", t.Minute())
    k := "you have an appointment on " + t.Weekday().String() + ", " + t.Month().String() + " " + d + ", " + y + ", at " + h + ":" + m + "."
    fmt.Println(k)
     return "You have an appointment on " + t.Weekday().String() + ", " + t.Month().String() + " " + d + ", " + y + ", at " + h + ":" + m + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	nowTime := time.Now()
    uni := time.Date(nowTime.Year(),9, 15,0,0,0,0,time.UTC)
    return uni
}
