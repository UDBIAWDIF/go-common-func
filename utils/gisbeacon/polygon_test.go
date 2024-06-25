package gisbeacon

import (
	"testing"
)

// go test -v -run="TestInPolygon"
func TestInPolygon(t *testing.T) {
	polygonList := [][2]float64{}
	polygonList = append(polygonList, [2]float64{119.505387, 26.070804})
	polygonList = append(polygonList, [2]float64{119.504708, 26.071116})
	polygonList = append(polygonList, [2]float64{119.507193, 26.074442})
	polygonList = append(polygonList, [2]float64{119.508779, 26.076140})
	polygonList = append(polygonList, [2]float64{119.509407, 26.07569})
	polygon := NewPolygon(polygonList)
	isIn := polygon.Contains([2]float64{119.505775, 26.071719})
	t.Log(isIn)
	isIn = polygon.Contains([2]float64{119.505684, 26.071791})
	t.Log(isIn)
}
