package main

type date struct {
  DateTime string
}

type CalendarItem struct {
    Summary string
    Description string
    Location string
    Start date
}

type Calendar struct {    
    Items []CalendarItem
}

type CombinedCalendarItem struct {
    Summary string
    Description string
    Location string
    Time string
}

type CombinedCalendar struct {
    Items []CombinedCalendarItem
}


