package codewars

/**
 * Тип местности.
 */
type Terrain byte

const (
	/**
	 * Равнина.
	 */
	Terrain_Plain Terrain = iota

	/**
	 * Топь.
	 */
	Terrain_Swamp

	/**
	 * Лес.
	 */
	Terrain_Forest
)