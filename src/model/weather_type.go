package model

/**
 * Тип погоды.
 */
type Weather byte

const (
	/**
	 * Ясно.
	 */
	Weather_Clear Weather = iota

	/**
	 * Плотные облака.
	 */
	Weather_Cloud

	/**
	 * Сильный дождь.
	 */
	Weather_Rain
)

