/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
    if len(intervals) <= 1{
        return true
    }
    // sort by start time
    sort.Slice(intervals, func(i, j int) bool{
        return intervals[i].start < intervals[j].start
    })
    // use a for loop to compare each interval
    for i := 1; i < len(intervals); i++{
        if intervals[i].start < intervals[i-1].end{
            return false
        }
    }
    return true
}
