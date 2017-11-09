package codewars

/**
 * Класс, определяющий сооружение --- прямоугольную область на карте.
 */
type Facility struct {
	/**
	 * Уникальный идентификатор сооружения.
	 */
	Id int64
	/**
	 * Тип сооружения.
	 */
	FacilityType FacilityType
	/**
	 * Идентификатор игрока, захватившего сооружение, или {@code -1}, если сооружение никем не
	 * контролируется.
	 */
	OwnerPlayerId int64
	/**
	 * Абсцисса левой границы сооружения.
	 */
	Left float64
	/**
	 * Ордината верхней границы сооружения.
	 */
	Top float64
	/**
	 * Индикатор захвата сооружения в интервале от {@code -game.MaxFacilityCapturePoints} до
	 * {@code game.MaxFacilityCapturePoints}. Если индикатор находится в положительной зоне, очки захвата принадлежат
	 * вам, иначе вашему противнику.
	 */
	CapturePoints float64
	/**
	 * Тип техники, производящейся в данном сооружении, или {@code null}. Применимо только к заводу
	 * ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	VehicleType VehicleType
	/**
	 * Неотрицательное число --- прогресс производства техники. Применимо только к заводу
	 * ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	ProductionProgress int
}
