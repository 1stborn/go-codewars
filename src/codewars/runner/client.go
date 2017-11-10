package runner

import (
	. "codewars"
	"errors"
	"os"
)

type MessageType byte

const (
	Message_GameOver MessageType = iota + 1
	Message_AuthenticationToken
	Message_TeamSize
	Message_ProtocolVersion
	Message_GameContext
	Message_PlayerContext
	Message_Move
)

const Version int = 2

var (
	ErrGameOver  = errors.New("game over")
	ErrWrongType = errors.New("wrong message type")
)

/**
 * Стратегия --- интерфейс, содержащий описание методов искусственного интеллекта армии.
 * Каждая пользовательская стратегия должна реализовывать этот интерфейс.
 * Может отсутствовать в некоторых языковых пакетах, если язык не поддерживает интерфейсы.
 */
type Strategy interface {
	/**
		 * Основной метод стратегии, осуществляющий управление армией. Вызывается каждый тик.
		 *
	     * me    Информация о вашем игроке.
	     * world Текущее состояние мира.
	     * game  Различные игровые константы.
	     * move  Результатом работы метода является изменение полей данного объекта.
	*/
	Move(*Player, *World, *Game, *Move)
}

type CodeWars struct {
	*AiCup

	players    map[int64]*Player
	facilities map[int64]*Facility
}

func Start(s Strategy) {
	var host, port, token string

	if len(os.Args) == 4 {
		host, port, token = os.Args[1], os.Args[2], os.Args[3]
	} else {
		host, port, token = "127.0.0.1", "31001", "0000000000000000"
	}

	cli := &CodeWars{
		AiCup:      new(AiCup),
		players:    make(map[int64]*Player),
		facilities: make(map[int64]*Facility),
	}

	if err := cli.Dial(host, port); err == nil {
		defer cli.Close()

		cli.writeToken(token)
		cli.writeProtoVersion(Version)
		cli.ReadTeamSize()

		g := cli.readGame()

		p := new(Player)
		w := new(World)

		for cli.readContext(p, w) != ErrGameOver {
			m := &Move{Type: Vehicle_None, Action:Action_None, Factor: 1.0}

			s.Move(p, w, g, m)

			cli.writeMove(m)
		}

		cli.Close()

		return
	} else {
		panic(err)
	}
}

func (c *CodeWars) readGame() *Game {
	c.ensureMessageType(Message_GameContext)

	if c.readBool() {
		return &Game{
			RandomSeed:                             c.readInt64(),
			TickCount:                              c.readInt(),
			WorldWidth:                             c.readFloat64(),
			WorldHeight:                            c.readFloat64(),
			FogOfWarEnabled:                        c.readBool(),
			VictoryScore:                           c.readInt(),
			FacilityCaptureScore:                   c.readInt(),
			VehicleEliminationScore:                c.readInt(),
			ActionDetectionInterval:                c.readInt(),
			BaseActionCount:                        c.readInt(),
			AdditionalActionCountPerControlCenter:  c.readInt(),
			MaxUnitGroup:                           c.readInt(),
			TerrainWeatherMapColumnCount:           c.readInt(),
			TerrainWeatherMapRowCount:              c.readInt(),
			PlainTerrainVisionFactor:               c.readFloat64(),
			PlainTerrainStealthFactor:              c.readFloat64(),
			PlainTerrainSpeedFactor:                c.readFloat64(),
			SwampTerrainVisionFactor:               c.readFloat64(),
			SwampTerrainStealthFactor:              c.readFloat64(),
			SwampTerrainSpeedFactor:                c.readFloat64(),
			ForestTerrainVisionFactor:              c.readFloat64(),
			ForestTerrainStealthFactor:             c.readFloat64(),
			ForestTerrainSpeedFactor:               c.readFloat64(),
			ClearWeatherVisionFactor:               c.readFloat64(),
			ClearWeatherStealthFactor:              c.readFloat64(),
			ClearWeatherSpeedFactor:                c.readFloat64(),
			CloudWeatherVisionFactor:               c.readFloat64(),
			CloudWeatherStealthFactor:              c.readFloat64(),
			CloudWeatherSpeedFactor:                c.readFloat64(),
			RainWeatherVisionFactor:                c.readFloat64(),
			RainWeatherStealthFactor:               c.readFloat64(),
			RainWeatherSpeedFactor:                 c.readFloat64(),
			VehicleRadius:                          c.readFloat64(),
			TankDurability:                         c.readInt(),
			TankSpeed:                              c.readFloat64(),
			TankVisionRange:                        c.readFloat64(),
			TankGroundAttackRange:                  c.readFloat64(),
			TankAerialAttackRange:                  c.readFloat64(),
			TankGroundDamage:                       c.readInt(),
			TankAerialDamage:                       c.readInt(),
			TankGroundDefence:                      c.readInt(),
			TankAerialDefence:                      c.readInt(),
			TankAttackCooldownTicks:                c.readInt(),
			TankProductionCost:                     c.readInt(),
			IFVDurability:                          c.readInt(),
			IFVSpeed:                               c.readFloat64(),
			IFVVisionRange:                         c.readFloat64(),
			IFVGroundAttackRange:                   c.readFloat64(),
			IFVAerialAttackRange:                   c.readFloat64(),
			IFVGroundDamage:                        c.readInt(),
			IFVAerialDamage:                        c.readInt(),
			IFVGroundDefence:                       c.readInt(),
			IFVAerialDefence:                       c.readInt(),
			IFVAttackCooldownTicks:                 c.readInt(),
			IFVProductionCost:                      c.readInt(),
			ARRVDurability:                         c.readInt(),
			ARRVSpeed:                              c.readFloat64(),
			ARRVVisionRange:                        c.readFloat64(),
			ARRVGroundDefence:                      c.readInt(),
			ARRVAerialDefence:                      c.readInt(),
			ARRVProductionCost:                     c.readInt(),
			ARRVRepairRange:                        c.readFloat64(),
			ARRVRepairSpeed:                        c.readFloat64(),
			HelicopterDurability:                   c.readInt(),
			HelicopterSpeed:                        c.readFloat64(),
			HelicopterVisionRange:                  c.readFloat64(),
			HelicopterGroundAttackRange:            c.readFloat64(),
			HelicopterAerialAttackRange:            c.readFloat64(),
			HelicopterGroundDamage:                 c.readInt(),
			HelicopterAerialDamage:                 c.readInt(),
			HelicopterGroundDefence:                c.readInt(),
			HelicopterAerialDefence:                c.readInt(),
			HelicopterAttackCooldownTicks:          c.readInt(),
			HelicopterProductionCost:               c.readInt(),
			FighterDurability:                      c.readInt(),
			FighterSpeed:                           c.readFloat64(),
			FighterVisionRange:                     c.readFloat64(),
			FighterGroundAttackRange:               c.readFloat64(),
			FighterAerialAttackRange:               c.readFloat64(),
			FighterGroundDamage:                    c.readInt(),
			FighterAerialDamage:                    c.readInt(),
			FighterGroundDefence:                   c.readInt(),
			FighterAerialDefence:                   c.readInt(),
			FighterAttackCooldownTicks:             c.readInt(),
			FighterProductionCost:                  c.readInt(),
			MaxFacilityCapturePoints:               c.readFloat64(),
			FacilityCapturePointsPerVehiclePerTick: c.readFloat64(),
			FacilityWidth:                          c.readFloat64(),
			FacilityHeight:                         c.readFloat64(),
		}
	}

	return nil
}

func (c *CodeWars) readContext(p *Player, w *World) error {
	switch c.readOpcode() {
	case Message_GameOver:
		return ErrGameOver
	case Message_PlayerContext:
		if c.readBool() {
			*p = *c.readPlayer()
			c.readWorld(w)
		}
		return nil
	default:
		return ErrWrongType
	}
}

func (c *CodeWars) readPlayer() *Player {
	switch c.readByte() {
	case 0:
		return nil
	case 127:
		return c.players[c.readInt64()]
	default:
		p := new(Player)
		p.Id = c.readInt64()
		p.Me = c.readBool()
		p.StrategyCrashed = c.readBool()
		p.Score = c.readInt()
		p.RemainingActionCooldownTicks = c.readInt()

		c.players[p.Id] = p

		return p
	}
}

func (c *CodeWars) readWorld(w *World) {
	if c.readBool() {
		w.TickIndex = c.readInt()
		w.TickCount = c.readInt()
		w.Width = c.readFloat64()
		w.Height = c.readFloat64()
		w.Players = c.readPlayers()
		w.NewVehicles = c.readVehicles()
		w.VehicleUpdates = c.readVehiclesUpdate()

		if w.TickIndex == 0 {
			w.TerrainByCellXY = c.readTerrains()
			w.WeatherByCellXY = c.readWeather()
		}

		w.Facilities = c.readFacilities()
	}
}

func (c *CodeWars) writeMove(m *Move) {
	c.writeOpcode(Message_Move)

	if m == nil {
		c.writeBool(false)
	} else {
		c.writeBool(true)

		c.writeByte(byte(m.Action))
		c.writeInt(m.Group)
		c.writeFloat64(m.Left)
		c.writeFloat64(m.Top)
		c.writeFloat64(m.Right)
		c.writeFloat64(m.Bottom)
		c.writeFloat64(m.X)
		c.writeFloat64(m.Y)
		c.writeFloat64(m.Angle)
		c.writeFloat64(m.MaxSpeed)
		c.writeFloat64(m.MaxAngularSpeed)
		c.writeByte(byte(m.Type))
		c.writeInt64(m.FacilityId)
	}

	c.flush()
}

func (c *CodeWars) readWeather() (weather [][]Weather) {
	for i := c.readInt(); i > 0; i-- {
		var slice []Weather
		for j := c.readInt(); j > 0; j-- {
			slice = append(slice, Weather(c.readByte()))
		}
		weather = append(weather, slice)
	}

	return
}

func (c *CodeWars) readTerrains() (terrain [][]Terrain) {
	for i := c.readInt(); i > 0; i-- {
		var slice []Terrain
		for j := c.readInt(); j > 0; j-- {
			slice = append(slice, Terrain(c.readByte()))
		}
		terrain = append(terrain, slice)
	}
	return
}

func (c *CodeWars) readFacility() *Facility {
	switch c.readByte() {
	case 0:
		return nil
	case 127:
		return c.facilities[c.readInt64()]
	default:
		f := new(Facility)
		f.Id = c.readInt64()
		f.FacilityType = FacilityType(c.readByte())
		f.OwnerPlayerId = c.readInt64()
		f.Left = c.readFloat64()
		f.Top = c.readFloat64()
		f.CapturePoints = c.readFloat64()
		f.VehicleType = VehicleType(c.readByte())
		f.ProductionProgress = c.readInt()

		c.facilities[f.Id] = f

		return f
	}
}

func (c *CodeWars) readVehicleUpdate() *VehicleUpdate {
	if c.readBool() {
		v := new(VehicleUpdate)
		v.Id = c.readInt64()
		v.X = c.readFloat64()
		v.Y = c.readFloat64()
		v.Durability = c.readInt()
		v.RemainingAttackCooldownTicks = c.readInt()
		v.Selected = c.readBool()
		v.Groups = c.readIntArray()

		return v
	}

	return nil
}

func (c *CodeWars) readNewVehicle() *Vehicle {
	if c.readBool() {
		v := new(Vehicle)
		v.Id = c.readInt64()
		v.X = c.readFloat64()
		v.Y = c.readFloat64()
		v.Radius = c.readFloat64()
		v.PlayerId = c.readInt64()
		v.Durability = c.readInt()
		v.MaxDurability = c.readInt()
		v.MaxSpeed = c.readFloat64()
		v.VisionRange = c.readFloat64()
		v.SquaredVisionRange = c.readFloat64()
		v.GroundAttackRange = c.readFloat64()
		v.SquaredGroundAttackRange = c.readFloat64()
		v.AerialAttackRange = c.readFloat64()
		v.SquaredAerialAttackRange = c.readFloat64()
		v.GroundDamage = c.readInt()
		v.AerialDamage = c.readInt()
		v.GroundDefence = c.readInt()
		v.AerialDefence = c.readInt()
		v.AttackCooldownTicks = c.readInt()
		v.RemainingAttackCooldownTicks = c.readInt()
		v.Type = VehicleType(c.readByte())
		v.Aerial = c.readBool()
		v.Selected = c.readBool()
		v.Groups = c.readIntArray()

		return v
	}

	return nil
}

func (c *CodeWars) readVehiclesUpdate() (updates []*VehicleUpdate) {
	for l := c.readInt(); l > 0; l-- {
		if v := c.readVehicleUpdate(); v != nil {
			updates = append(updates, v)
		}
	}
	return
}

func (c *CodeWars) readFacilities() (facilities []*Facility) {
	for l := c.readInt(); l > 0; l-- {
		if f := c.readFacility(); f != nil {
			facilities = append(facilities, f)
		}
	}
	return
}

func (c *CodeWars) readVehicles() (vehicles []*Vehicle) {
	for l := c.readInt(); l > 0; l-- {
		if v := c.readNewVehicle(); v != nil {
			vehicles = append(vehicles, v)
		}
	}
	return
}

func (c *CodeWars) readPlayers() (players []*Player) {
	for l := c.readInt(); l > 0; l-- {
		if p := c.readPlayer(); p != nil {
			players = append(players, p)
		}
	}
	return
}
