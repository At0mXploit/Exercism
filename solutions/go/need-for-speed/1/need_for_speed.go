package speed

// Car struct represents a remote controlled car
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100, // full battery
		distance:     0,   // starting at 0 meters
	}
}

// Track struct represents a race track
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

// CanFinish checks if a car can finish the track with its current battery
func CanFinish(car Car, track Track) bool {
	drivesNeeded := track.distance / car.speed
	batteryRequired := drivesNeeded * car.batteryDrain
	return car.battery >= batteryRequired
}
