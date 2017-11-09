package geo

import "github.com/dhconnelly/rtreego"

func Distance2LatAndLon(distance float64) (float64, float64) {
	var lonRatio, latRatio, lonRadius, latRadius float64
	lonRatio = 85375.5033525
	latRatio = 111319.490821
	lonRadius = distance / lonRatio
	latRadius = distance / latRatio

	return latRadius, lonRadius
}

func GetRectByPointAndDistance(point rtreego.Point, distance float64) *rtreego.Rect {
	latRadius, lonRadius := Distance2LatAndLon(distance)
	a := rtreego.Point{point[0] - latRadius, point[1] - lonRadius}
	b := []float64{2 * latRadius, 2 * lonRadius}
	rect, _ := rtreego.NewRect(a, b)

	return rect
}
