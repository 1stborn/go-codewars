package model

/**
 * Предоставляет доступ к различным игровым константам.
 */
type Game struct {
	/**
	 * Некоторое число, которое ваша стратегия может использовать для инициализации генератора
	 * случайных чисел. Данное значение имеет рекомендательный характер, однако позволит более точно воспроизводить
	 * прошедшие игры.
	 */
	RandomSeed int64
	/**
	 * Базовая длительность игры в тиках. Реальная длительность может отличаться от этого значения в
	 * меньшую сторону. Эквивалентно {@code world.TickCount}.
	 */
	TickCount int
	/**
	 * Ширина карты.
	 */
	WorldWidth float64
	/**
	 * Высота карты.
	 */
	WorldHeight float64

	/**
	 * {@code true}, если и только если в данной игре включен режим частичной видимости.
	 */
	FogOfWarEnabled bool
	/**
	 * Количество баллов, получаемое игроком в случае уничтожения всех юнитов противника.
	 */
	VictoryScore int
	/**
	 * Количество баллов за захват сооружения.
	 */
	FacilityCaptureScore int
	/**
	 * Количество баллов за уничтожение юнита противника.
	 */
	VehicleEliminationScore int
	/**
	 * Интервал, учитываемый в ограничении количества действий стратегии.
	 */
	ActionDetectionInterval int
	/**
	 * Базовое количество действий, которое может совершить стратегия за
	 * {@code ActionDetectionInterval} последовательных тиков.
	 */
	BaseActionCount int
	/**
	 * Дополнительное количество действий за каждый захваченный центр управления
	 * ({@code FacilityType.CONTROL_CENTER}).
	 */
	AdditionalActionCountPerControlCenter int
	/**
	 * Максимально возможный индекс группы юнитов.
	 */
	MaxUnitGroup int
	/**
	 * Количество столбцов в картах местности и погоды.
	 */
	TerrainWeatherMapColumnCount int
	/**
	 * Количество строк в картах местности и погоды.
	 */
	TerrainWeatherMapRowCount int
	/**
	 * Мультипликатор радиуса обзора наземной техники, находящейся на равнинной местности
	 * ({@code Terrain.PLAIN}).
	 */
	PlainTerrainVisionFactor float64
	/**
	 * Мультипликатор радиуса обзора любой техники при обнаружении наземной техники противника,
	 * находящейся на равнинной местности ({@code Terrain.PLAIN}).
	 */
	PlainTerrainStealthFactor float64
	/**
	 * Мультипликатор максимальной скорости наземной техники, находящейся на равнинной местности
	 * ({@code Terrain.PLAIN}).
	 */
	PlainTerrainSpeedFactor float64
	/**
	 * Мультипликатор радиуса обзора наземной техники, находящейся в болотистой местности
	 * ({@code Terrain.SWAMP}).
	 */
	SwampTerrainVisionFactor float64
	/**
	 * Мультипликатор радиуса обзора любой техники при обнаружении наземной техники противника,
	 * находящейся в болотистой местности ({@code Terrain.SWAMP}).
	 */
	SwampTerrainStealthFactor float64
	/**
	 * Мультипликатор максимальной скорости наземной техники, находящейся в болотистой местности
	 * ({@code Terrain.SWAMP}).
	 */
	SwampTerrainSpeedFactor float64
	/**
	 * Мультипликатор радиуса обзора наземной техники, находящейся в лесистой местности
	 * ({@code Terrain.FOREST}).
	 */
	ForestTerrainVisionFactor float64
	/**
	 * Мультипликатор радиуса обзора любой техники при обнаружении наземной техники противника,
	 * находящейся в лесистой местности ({@code Terrain.FOREST}).
	 */
	ForestTerrainStealthFactor float64
	/**
	 * Мультипликатор максимальной скорости наземной техники, находящейся в лесистой местности
	 * ({@code Terrain.FOREST}).
	 */
	ForestTerrainSpeedFactor float64
	/**
	 * Мультипликатор радиуса обзора воздушной техники, находящейся в области ясной погоды
	 * ({@code Weather.CLEAR}).
	 */
	ClearWeatherVisionFactor float64
	/**
	 * Мультипликатор радиуса обзора любой техники при обнаружении воздушной техники противника,
	 * находящейся в области ясной погоды ({@code Weather.CLEAR}).
	 */
	ClearWeatherStealthFactor float64
	/**
	 * Мультипликатор максимальной скорости воздушной техники, находящейся в области ясной погоды
	 * ({@code Weather.CLEAR}).
	 */
	ClearWeatherSpeedFactor float64
	/**
	 * Мультипликатор радиуса обзора воздушной техники, находящейся в плотных облаках
	 * ({@code Weather.CLOUD}).
	 */
	CloudWeatherVisionFactor float64
	/**
	 * Мультипликатор радиуса обзора любой техники при обнаружении воздушной техники противника,
	 * находящейся в плотных облаках ({@code Weather.CLOUD}).
	 */
	CloudWeatherStealthFactor float64
	/**
	 * Мультипликатор максимальной скорости воздушной техники, находящейся в плотных облаках
	 * ({@code Weather.CLOUD}).
	 */
	CloudWeatherSpeedFactor float64
	/**
	 * Мультипликатор радиуса обзора воздушной техники, находящейся в условиях сильного дождя
	 * ({@code Weather.RAIN}).
	 */
	RainWeatherVisionFactor float64
	/**
	 * Мультипликатор радиуса обзора любой техники при обнаружении воздушной техники противника,
	 * находящейся в условиях сильного дождя ({@code Weather.RAIN}).
	 */
	RainWeatherStealthFactor float64
	/**
	 * Мультипликатор максимальной скорости воздушной техники, находящейся в условиях сильного дождя
	 * ({@code Weather.RAIN}).
	 */
	RainWeatherSpeedFactor float64
	/**
	 * Радиус техники.
	 */
	VehicleRadius float64
	/**
	 * Максимальная прочность танка.
	 */
	TankDurability int
	/**
	 * Максимальная скорость танка.
	 */
	TankSpeed float64
	/**
	 * Базовый радиус обзора танка.
	 */
	TankVisionRange float64
	/**
	 * Дальность атаки танка по наземным целям.
	 */
	TankGroundAttackRange float64
	/**
	 * Дальность атаки танка по воздушным целям.
	 */
	TankAerialAttackRange float64
	/**
	 * Урон одной атаки танка по наземной технике.
	 */
	TankGroundDamage int
	/**
	 * Урон одной атаки танка по воздушной технике.
	 */
	TankAerialDamage int
	/**
	 * Защита танка от атак наземной техники.
	 */
	TankGroundDefence int
	/**
	 * Защита танка от атак воздушной техники.
	 */
	TankAerialDefence int
	/**
	 * Интервал в тиках между двумя последовательными атаками танка.
	 */
	TankAttackCooldownTicks int
	/**
	 * Количество тиков, необхожимое для производства одного танка на заводе
	 * ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	TankProductionCost int
	/**
	 * Максимальная прочность БМП.
	 */
	IFVDurability int
	/**
	 * Максимальная скорость БМП.
	 */
	IFVSpeed float64
	/**
	 * Базовый радиус обзора БМП.
	 */
	IFVVisionRange float64
	/**
	 * Дальность атаки БМП по наземным целям.
	 */
	IFVGroundAttackRange float64
	/**
	 * Дальность атаки БМП по воздушным целям.
	 */
	IFVAerialAttackRange float64
	/**
	 * Урон одной атаки БМП по наземной технике.
	 */
	IFVGroundDamage int
	/**
	 * Урон одной атаки БМП по воздушной технике.
	 */
	IFVAerialDamage int
	/**
	 * Защита БМП от атак наземной техники.
	 */
	IFVGroundDefence int
	/**
	 * Защита БМП от атак воздушной техники.
	 */
	IFVAerialDefence int
	/**
	 * Интервал в тиках между двумя последовательными атаками БМП.
	 */
	IFVAttackCooldownTicks int
	/**
	 * Количество тиков, необхожимое для производства одной БМП на заводе
	 * ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	IFVProductionCost int
	/**
	 * Максимальная прочность БРЭМ.
	 */
	ARRVDurability int
	/**
	 * Максимальная скорость БРЭМ.
	 */
	ARRVSpeed float64
	/**
	 * Базовый радиус обзора БРЭМ.
	 */
	ARRVVisionRange float64
	/**
	 * Защита БРЭМ от атак наземной техники.
	 */
	ARRVGroundDefence int
	/**
	 * Защита БРЭМ от атак воздушной техники.
	 */
	ARRVAerialDefence int
	/**
	 * Количество тиков, необхожимое для производства одной БРЭМ на заводе
	 * ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	ARRVProductionCost int
	/**
	 * Максимальное расстояние (от центра до центра), на котором БРЭМ может ремонтировать
	 * дружественная технику.
	 */
	ARRVRepairRange float64
	/**
	 * Максимальное количество прочности, которое БРЭМ может восстановить дружественной технике за
	 * один тик.
	 */
	ARRVRepairSpeed float64
	/**
	 * Максимальная прочность ударного вертолёта.
	 */
	HelicopterDurability int
	/**
	 * Максимальная скорость ударного вертолёта.
	 */
	HelicopterSpeed float64
	/**
	 * Базовый радиус обзора ударного вертолёта.
	 */
	HelicopterVisionRange float64
	/**
	 * Дальность атаки ударного вертолёта по наземным целям.
	 */
	HelicopterGroundAttackRange float64
	/**
	 * Дальность атаки ударного вертолёта по воздушным целям.
	 */
	HelicopterAerialAttackRange float64
	/**
	 * Урон одной атаки ударного вертолёта по наземной технике.
	 */
	HelicopterGroundDamage int
	/**
	 * Урон одной атаки ударного вертолёта по воздушной технике.
	 */
	HelicopterAerialDamage int
	/**
	 * Защита ударного вертолёта от атак наземной техники.
	 */
	HelicopterGroundDefence int
	/**
	 * Защита ударного вертолёта от атак воздушной техники.
	 */
	HelicopterAerialDefence int
	/**
	 * Интервал в тиках между двумя последовательными атаками ударного вертолёта.
	 */
	HelicopterAttackCooldownTicks int
	/**
	 * Количество тиков, необхожимое для производства одного ударного вертолёта на заводе
	 * ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	HelicopterProductionCost int
	/**
	 * Максимальная прочность истребителя.
	 */
	FighterDurability int
	/**
	 * Максимальная скорость истребителя.
	 */
	FighterSpeed float64
	/**
	 * Базовый радиус обзора истребителя.
	 */
	FighterVisionRange float64
	/**
	 * Дальность атаки истребителя по наземным целям.
	 */
	FighterGroundAttackRange float64
	/**
	 * Дальность атаки истребителя по воздушным целям.
	 */
	FighterAerialAttackRange float64
	/**
	 * Урон одной атаки истребителя по наземной технике.
	 */
	FighterGroundDamage int
	/**
	 * Урон одной атаки истребителя по воздушной технике.
	 */
	FighterAerialDamage int
	/**
	 * Защита истребителя от атак наземной техники.
	 */
	FighterGroundDefence int
	/**
	 * Защита истребителя от атак воздушной техники.
	 */
	FighterAerialDefence int
	/**
	 * Интервал в тиках между двумя последовательными атаками истребителя.
	 */
	FighterAttackCooldownTicks int
	/**
	 * Количество тиков, необхожимое для производства одного истребителя на заводе
	 * ({@code FacilityType.VEHICLE_FACTORY}).
	 */
	FighterProductionCost int
	/**
	 * Максимально возможная абсолютная величина индикатора захвата сооружения
	 * ({@code facility.CapturePoints}).
	 */
	MaxFacilityCapturePoints float64
	/**
	 * Скорость изменения индикатора захвата сооружения ({@code facility.CapturePoints}) за каждую
	 * единицу техники, центр которой находится внутри сооружения.
	 */
	FacilityCapturePointsPerVehiclePerTick float64
	/**
	 * Ширина сооружения.
	 */
	FacilityWidth float64
	/**
	 * Высота сооружения.
	 */
	FacilityHeight float64
	/**
	 * Минимально возможный интервал между двумя последовательными тактическими ядерными ударами
	 */
	BaseTacticalNuclearStrikeCooldown int
	/**
	 * Уменьшение интервала между тактическими ядерными ударами за каждый захваченный центр
	 */
	TacticalNuclearStrikeCooldownDecreasePerControlCenter int
	/**
	 * Урон тактического ядерного удара в центре взрыва.
	 */
	TacticalNuclearStrikeMaxDamage float64
	/**
	 *  Радиус взрыва тактического ядерного удара.
	 */
	TacticalNuclearStrikeRadius float64
	/**
	 * Задержка между запросом нанесения тактического ядерного удара и собственно самим нанесением.
	 */
	TacticalNuclearStrikeDelay int
}
