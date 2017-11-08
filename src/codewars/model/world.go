package model

import "sync"

const TileSize = 32

/**
 * Этот класс описывает игровой мир. Содержит также описания всех игроков, игровых объектов (<<юнитов>>) и сооружений.
 */
type World struct {
	sync.RWMutex

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
	Players map[int64]*Player

	/**
	 * Список техники, о которой у стратегии не было информации в предыдущий игровой тик. В этот
	 * список попадает как только что произведённая техника, так и уже существующая, но находящаяся вне зоны видимости
	 * до этого момента.
	 *
	 * Значения изменяемых полей для каждой видимой техники, если хотя бы одно поле этой техники
	 * изменилось. Нулевая прочность означает, что техника была уничтожена либо ушла из зоны видимости.
	 */
	Vehicles map[int64]*Vehicle

	LineSize int

	Land []LandType
	/**
	 * Список сооружений (в случайном порядке).
	 * В зависимости от реализации, объекты, задающие сооружения, могут пересоздаваться после каждого тика.
	 */
	Facilities map[int64]*Facility
}

/**
 * @return Возвращает вашего игрока.
 */
func (w *World) MyPlayer() *Player {
	w.RLock()
	defer w.RUnlock()

	for _, p := range w.Players {
		if p.Me {
			return p
		}
	}
	return nil
}

func (w *World) GetLand(x, y float64) LandType {
	return w.Land[w.LineSize*int(x/TileSize)+int(y/TileSize)]
}

/**
 * @return Возвращает игрока, соревнующегося с вами.
 */
func (w *World) OpponentPlayer() *Player {
	w.RLock()
	defer w.RUnlock()

	for _, p := range w.Players {
		if !p.Me {
			return p
		}
	}

	return nil
}

func (w *World) Player(id int64) *Player {
	w.RLock()
	if p, ok := w.Players[id]; ok {
		w.RUnlock()
		return p
	} else {
		w.RUnlock()
		w.Lock()
		p = &Player{Id: id}
		w.Players[id] = p
		w.Unlock()
		return p
	}
}

func (w *World) Vehicle(id int64) *Vehicle {
	w.RLock()
	if v, ok := w.Vehicles[id]; ok {
		w.RUnlock()
		return v
	} else {
		w.RUnlock()
		w.Lock()
		v := vehiclePool.Get().(*Vehicle)
		v.Id = id
		w.Vehicles[id] = v
		w.Unlock()
		return v
	}
}

func (w *World) Facility(id int64) *Facility {
	w.RLock()
	if f, ok := w.Facilities[id]; ok {
		w.RUnlock()
		return f
	} else {
		w.RUnlock()
		w.Lock()
		f := &Facility{Id: id}
		w.Facilities[id] = f
		w.Unlock()
		return f
	}
}
