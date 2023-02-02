type customerCheckin struct {
	stationName string
	t           int
}

type UndergroundSystem struct {
	checkIn   map[int]*customerCheckin
	checkedIn map[string]int
	totalTime map[string]int
}

func Constructor() UndergroundSystem {
	checkIn := make(map[int]*customerCheckin)
	checkedIn := make(map[string]int)
	totalTime := make(map[string]int)
	return UndergroundSystem{checkIn, checkedIn, totalTime}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if this.checkIn[id] != nil {
		return
	}
	this.checkIn[id] = &customerCheckin{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if this.checkIn[id] == nil {
		return
	}
	checkInRecord := this.checkIn[id]
	inOutName := fmt.Sprintf("%s-%s", checkInRecord.stationName, stationName)
	this.checkedIn[inOutName]++
	this.totalTime[inOutName] += t - checkInRecord.t
	this.checkIn[id] = nil
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	inOutName := fmt.Sprintf("%s-%s", startStation, endStation)
	return float64(this.totalTime[inOutName]) / float64(this.checkedIn[inOutName])
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
