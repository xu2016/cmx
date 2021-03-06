package xjwd

/*通过构造水平向右的射线，判断点是否在多边形内部
1、原理：当射线与多边形各边相交的交点数为奇数是，点在多边形内，否则在多边形外
2、计算方法：射线与多边形的顶点相交，可以选择以下一种方法计算交点数（下面都是针对纵坐标，因为选择的是水平射线）：
    A、低闭高开法：当测试点的y属于[min(y1,y2),max(y1,y2))时才计算：
        1)首先验证是否满足高开(可排除与射线平行的边)
        2)验证低闭
        3)验证横坐标x范围
        4)排除水平边
        5)计算是否相交
    B、低开高闭法：当测试点的y属于(min(y1,y2),max(y1,y2)]时才计算：
        1)首先验证是否满足低开(可排除与射线平行的边)
        2)验证高闭
        3)验证横坐标x范围
        4)排除水平边
        5)计算是否相交
3、本函数选择"低开高闭法"。
*/
func insidepolygon(x, y float64) bool {

	return false
}
