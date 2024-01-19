package funcs

import (
	"math"
	"strconv"
	"strings"
)

type gisPoint struct {
	Lng float64 `json:"lng"` //经度(180°W-180°E) -180 - 0 - 180
	Lat float64 `json:"lat"` //纬度（90°S-90°N）-90 -0 -90
}

type gisArea struct {
	Points []gisPoint //多边形坐标 顺时针方向
}

// 实例化一个面
func NewGISArea(points []gisPoint) *gisArea {
	area := new(gisArea)
	area.Points = points
	return area
}

// 判断一个点是否在一个面内
// 如果点位于多边形的顶点或边上，也算做点在多边形内，直接返回true
// @param point  指定点坐标
func (a *gisArea) PointInArea(point gisPoint) bool {
	pointNum := len(a.Points) //点个数
	intersectCount := 0       //cross points count of x
	precision := 2e-10        //浮点类型计算时候与0比较时候的容差
	p1 := gisPoint{}          //neighbour bound vertices
	p2 := gisPoint{}
	p := point //测试点

	p1 = a.Points[0] //left vertex
	for i := 0; i < pointNum; i++ {
		if p.Lng == p1.Lng && p.Lat == p1.Lat {
			return true
		}
		p2 = a.Points[i%pointNum]
		if p.Lat < math.Min(p1.Lat, p2.Lat) || p.Lat > math.Max(p1.Lat, p2.Lat) {
			p1 = p2
			continue //next ray left point
		}

		if p.Lat > math.Min(p1.Lat, p2.Lat) && p.Lat < math.Max(p1.Lat, p2.Lat) {
			if p.Lng <= math.Max(p1.Lng, p2.Lng) { //x is before of ray
				if p1.Lat == p2.Lat && p.Lng >= math.Min(p1.Lng, p2.Lng) {
					return true
				}

				if p1.Lng == p2.Lng { //ray is vertical
					if p1.Lng == p.Lng { //overlies on a vertical ray
						return true
					} else { //before ray
						intersectCount++
					}
				} else { //cross point on the left side
					xinters := (p.Lat-p1.Lat)*(p2.Lng-p1.Lng)/(p2.Lat-p1.Lat) + p1.Lng
					if math.Abs(p.Lng-xinters) < precision {
						return true
					}

					if p.Lng < xinters { //before ray
						intersectCount++
					}
				}
			}
		} else { //special case when ray is crossing through the vertex
			if p.Lat == p2.Lat && p.Lng <= p2.Lng { //p crossing over p2
				p3 := a.Points[(i+1)%pointNum]
				if p.Lat >= math.Min(p1.Lat, p3.Lat) && p.Lat <= math.Max(p1.Lat, p3.Lat) {
					intersectCount++
				} else {
					intersectCount += 2
				}
			}
		}
		p1 = p2 //next ray left point
	}
	if intersectCount%2 == 0 { //偶数在多边形外
		return false
	} else { //奇数在多边形内
		return true
	}
}

func CheckGISPointInArea(areaString string, pointLng string, pointLat string) bool {
	// var poi string
	// poi = "118.183506,39.626009,118.190501,39.626009,118.190501,39.618406,118.183506,39.618406,118.152262,39.709224,118.16106,39.709224,118.16106,39.704288,118.152262,39.704288"
	// poi = areaString
	// strpoi := strings.Split(poi, ",")
	strpoi := strings.Split(areaString, ",")
	var pois []gisPoint
	for i := 0; i < len(strpoi); i = i + 2 {
		// fmt.Println("Lng:" + strpoi[i] + ",Lat:" + strpoi[i+1])
		//string到float64/32
		float64a, _ := strconv.ParseFloat(strpoi[i], 64)
		float64b, _ := strconv.ParseFloat(strpoi[i+1], 64)
		pois = append(pois, gisPoint{float64a, float64b}) //结构体数组赋值
		// fmt.Println(pois)
	}
	//points := []Point{{Lng:1,Lat:1},{Lng:-1,Lat:2},{Lng:-3,Lat:8},{Lng:0,Lat:4},{Lng:2.1,Lat:6.8},{Lng:9,Lat:-3.1}}
	// fmt.Println(pois)
	area := NewGISArea(pois)
	var positionLng float64
	var positionLat float64
	positionLng, _ = strconv.ParseFloat(pointLng, 64)
	positionLat, _ = strconv.ParseFloat(pointLat, 64)
	point := gisPoint{Lng: positionLng, Lat: positionLat}
	rt := area.PointInArea(point)
	// fmt.Print("点是否在多边形内:",rt)
	return rt
}

func IsPointInsidePolygon(point gisPoint, polygon []gisPoint) bool {
	inside := false
	for i, j := 0, len(polygon)-1; i < len(polygon); i++ {
		if (polygon[i].Lat > point.Lat) != (polygon[j].Lat > point.Lat) &&
			point.Lng < (polygon[j].Lng-polygon[i].Lng)*(point.Lat-polygon[i].Lat)/(polygon[j].Lat-polygon[i].Lat)+polygon[i].Lng {
			inside = !inside
		}
		j = i
	}
	return inside
}
