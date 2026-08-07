type timeValue struct{
    value string
    timestamp int
}
type TimeMap struct {
    store map[string][]timeValue   
}

func Constructor() TimeMap {
    return TimeMap{
        store: make(map[string][]timeValue),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.store[key] = append(this.store[key], timeValue{
        value:value,
        timestamp:timestamp,
        })
}

func (this *TimeMap) Get(key string, timestamp int) string {
    series, exist := this.store[key]
    result := ""
    if !exist || len(series) == 0{
        return result
    }

    left, right := 0, len(series)-1
    for left <= right{
        mid := left+(right-left)/2
        if series[mid].timestamp <= timestamp{
            result = series[mid].value
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return result
}
