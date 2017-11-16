package codewars

/**
 * Тип сооружения.
 */
type FacilityType byte

const (
	/**
	 * Центр управления. Увеличивает возможное количество действий игрока на
	 * {@code game.AdditionalActionCountPerControlCenter} за {@code game.actionDetectionInterval} игровых тиков.
	 */
	Facility_ControlCenter FacilityType = iota

	/**
	 * Завод. Может производить технику любого типа по выбору игрока.
	 */
	Facility_VehicleFactory
)
