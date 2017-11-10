package codewars

/**
 * Этот класс описывает игровой мир. Содержит также описания всех игроков, игровых объектов (<<юнитов>>) и сооружений.
 */
type World struct {
	/**
	 * Номер текущего тика.
	 */
	TickIndex int
	/**
	 * Базовую длительность игры в тиках. Реальная длительность может отличаться от этого значения в
	 * меньшую сторону. Эквивалентно {@code game.tickCount}.
	 */
	TickCount int
	/**
	 * Ширину мира.
	 */
	Width float64
	/**
	 * Высоту мира.
	 */
	Height float64
	/**
	 * Список игроков (в случайном порядке).
	 * В зависимости от реализации, объекты, задающие игроков, могут пересоздаваться после каждого тика.
	 */
	Players []*Player

	/**
	 * Список техники, о которой у стратегии не было информации в предыдущий игровой тик. В этот
	 * список попадает как только что произведённая техника, так и уже существующая, но находящаяся вне зоны видимости
	 * до этого момента.
	 */
	NewVehicles []*Vehicle

	/** Значения изменяемых полей для каждой видимой техники, если хотя бы одно поле этой техники
	 * изменилось. Нулевая прочность означает, что техника была уничтожена либо ушла из зоны видимости.
	 */
	VehicleUpdates []*VehicleUpdate

	TerrainByCellXY [][]Terrain
	WeatherByCellXY [][]Weather

	/**
	 * Список сооружений (в случайном порядке).
	 * В зависимости от реализации, объекты, задающие сооружения, могут пересоздаваться после каждого тика.
	 */
	Facilities []*Facility
}

/**
 * @return Возвращает вашего игрока.
 */
func (w *World) MyPlayer() *Player {
	for _, p := range w.Players {
		if p.Me {
			return p
		}
	}
	return nil
}

/**
 * @return Возвращает игрока, соревнующегося с вами.
 */
func (w *World) OpponentPlayer() *Player {
	for _, p := range w.Players {
		if !p.Me {
			return p
		}
	}

	return nil
}
