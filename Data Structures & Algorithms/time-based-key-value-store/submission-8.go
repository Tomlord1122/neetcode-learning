type TimeValue struct{
	value string
	timestamp int
}
type TimeMap struct {
	store map[string][]TimeValue
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]TimeValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], TimeValue{
		value: value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	series, exist := this.store[key]
	if !exist || len(series) == 0{
		return ""
	}
	res := ""
	l, r := 0, len(series)-1
	for l <= r{
		m := (l+r) / 2
		if series[m].timestamp <= timestamp{
			res = series[m].value
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}
