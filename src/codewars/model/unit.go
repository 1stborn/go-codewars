package model

import "math"

type CircularUnit struct {
	Unit
	/**
	* Радиус объекта.
	 */
	Radius float64
}

func (c *CircularUnit) GetRadius() float64 {
	return c.Radius
}

/**
 * Базовый класс для определения объектов (<<юнитов>>) на игровом поле.
 */
type Unit struct {
	/**
	 * Уникальный идентификатор объекта.
	 */
	Id int64
	/**
	 * X-координата центра объекта. Ось абсцисс направлена слева направо.
	 */
	X float64
	/**
	 * Y-координата центра объекта. Ось ординат направлена сверху вниз.
	 */
	Y float64
}

/**
 * X-координата точки.
 * Y-координата точки.
 * Возвращает расстояние до точки от центра данного объекта.
 */
func (u *Unit) GetDistanceTo(x, y float64) float64 {
	return math.Sqrt(math.Pow(x-u.X, 2) + math.Pow(y-u.Y, 2))
}

/**
 * Объект, до центра которого необходимо определить расстояние.
 * Возвращает расстояние от центра данного объекта до центра указанного объекта.
 */
func (u *Unit) GetDistanceUnit(u2 *Unit) float64 {
	return u.GetDistanceTo(u2.X, u2.Y)
}

/**
 * X-координата точки.
 * Y-координата точки.
 * Возвращает квадрат расстояния до точки от центра данного объекта.
 */
func (u *Unit) GetSquaredDistanceTo(x, y float64) float64 {
	dx := x - u.X
	dy := y - u.Y
	return dx*dx + dy*dy
}

/**
 * Объект, до центра которого необходимо определить квадрат расстояния.
 * Возвращает квадрат расстояния от центра данного объекта до центра указанного объекта.
 */
func (u *Unit) GetSquaredDistanceUnit(u2 *Unit) float64 {
	return u.GetSquaredDistanceTo(u2.X, u2.Y)
}
