package codewars

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
