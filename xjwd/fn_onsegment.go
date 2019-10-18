package xjwd

import "math"

/*OnSegment 判断点Q(x,y)是否在点pi(x1,y1)和点pj(x2,y2)构成的直线段内
点pi和点pj构成的直线的斜率与点Q和点pi构成的斜率相同
且点Q在点pi和点pj的线段内
*/
func OnSegment(x1, y1, x2, y2, x, y float64) bool {
	b := false
	if (x-x1)*(y2-y1)-(x2-x1)*(y-y1) <= eps {
		b = true
	}
	if b && math.Min(x1, x2) <= x && x <= math.Max(x1, x2) && math.Min(y1, y2) <= y && y <= math.Max(y1, y2) {
		return true
	}
	return false
}
