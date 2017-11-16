package client

import (
	. "codewars"
	"errors"
	"net"
	"bufio"
	"encoding/binary"
	"os"
)

var ByteOrder = binary.LittleEndian

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

const Version int = 3

var (
	ErrGameOver  = errors.New("game over")
	ErrWrongType = errors.New("wrong message type")
)

type RemoteProcessClient struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer

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

	cli := &RemoteProcessClient{
		players:    make(map[int64]*Player),
		facilities: make(map[int64]*Facility),
	}

	if err := cli.Dial(host, port); err == nil {
		defer cli.Close()

		cli.writeToken(token)
		cli.writeProtoVersion(Version)
		cli.ReadTeamSize()

		g := cli.readGame()

		pc := &PlayerContext{Player: new(Player), World: new(World)}

		for cli.readContext(pc) != ErrGameOver {
			m := &Move{
				Type:       Vehicle_None,
				Action:     Action_None,
				Factor:     1,
				FacilityId: -1,
				VehicleId:  -1,
			}

			s.Move(pc.Player, pc.World, g, m)

			cli.writeMove(m)
		}

		cli.Close()

		return
	} else {
		panic(err)
	}
}

func (c *RemoteProcessClient) readGame() *Game {
	c.ensureMessageType(Message_GameContext)

	if c.readBool() {
		return &Game{
			RandomSeed:                                            c.readInt64(),
			TickCount:                                             c.readInt(),
			WorldWidth:                                            c.readFloat64(),
			WorldHeight:                                           c.readFloat64(),
			FogOfWarEnabled:                                       c.readBool(),
			VictoryScore:                                          c.readInt(),
			FacilityCaptureScore:                                  c.readInt(),
			VehicleEliminationScore:                               c.readInt(),
			ActionDetectionInterval:                               c.readInt(),
			BaseActionCount:                                       c.readInt(),
			AdditionalActionCountPerControlCenter:                 c.readInt(),
			MaxUnitGroup:                                          c.readInt(),
			TerrainWeatherMapColumnCount:                          c.readInt(),
			TerrainWeatherMapRowCount:                             c.readInt(),
			PlainTerrainVisionFactor:                              c.readFloat64(),
			PlainTerrainStealthFactor:                             c.readFloat64(),
			PlainTerrainSpeedFactor:                               c.readFloat64(),
			SwampTerrainVisionFactor:                              c.readFloat64(),
			SwampTerrainStealthFactor:                             c.readFloat64(),
			SwampTerrainSpeedFactor:                               c.readFloat64(),
			ForestTerrainVisionFactor:                             c.readFloat64(),
			ForestTerrainStealthFactor:                            c.readFloat64(),
			ForestTerrainSpeedFactor:                              c.readFloat64(),
			ClearWeatherVisionFactor:                              c.readFloat64(),
			ClearWeatherStealthFactor:                             c.readFloat64(),
			ClearWeatherSpeedFactor:                               c.readFloat64(),
			CloudWeatherVisionFactor:                              c.readFloat64(),
			CloudWeatherStealthFactor:                             c.readFloat64(),
			CloudWeatherSpeedFactor:                               c.readFloat64(),
			RainWeatherVisionFactor:                               c.readFloat64(),
			RainWeatherStealthFactor:                              c.readFloat64(),
			RainWeatherSpeedFactor:                                c.readFloat64(),
			VehicleRadius:                                         c.readFloat64(),
			TankDurability:                                        c.readInt(),
			TankSpeed:                                             c.readFloat64(),
			TankVisionRange:                                       c.readFloat64(),
			TankGroundAttackRange:                                 c.readFloat64(),
			TankAerialAttackRange:                                 c.readFloat64(),
			TankGroundDamage:                                      c.readInt(),
			TankAerialDamage:                                      c.readInt(),
			TankGroundDefence:                                     c.readInt(),
			TankAerialDefence:                                     c.readInt(),
			TankAttackCooldownTicks:                               c.readInt(),
			TankProductionCost:                                    c.readInt(),
			IFVDurability:                                         c.readInt(),
			IFVSpeed:                                              c.readFloat64(),
			IFVVisionRange:                                        c.readFloat64(),
			IFVGroundAttackRange:                                  c.readFloat64(),
			IFVAerialAttackRange:                                  c.readFloat64(),
			IFVGroundDamage:                                       c.readInt(),
			IFVAerialDamage:                                       c.readInt(),
			IFVGroundDefence:                                      c.readInt(),
			IFVAerialDefence:                                      c.readInt(),
			IFVAttackCooldownTicks:                                c.readInt(),
			IFVProductionCost:                                     c.readInt(),
			ARRVDurability:                                        c.readInt(),
			ARRVSpeed:                                             c.readFloat64(),
			ARRVVisionRange:                                       c.readFloat64(),
			ARRVGroundDefence:                                     c.readInt(),
			ARRVAerialDefence:                                     c.readInt(),
			ARRVProductionCost:                                    c.readInt(),
			ARRVRepairRange:                                       c.readFloat64(),
			ARRVRepairSpeed:                                       c.readFloat64(),
			HelicopterDurability:                                  c.readInt(),
			HelicopterSpeed:                                       c.readFloat64(),
			HelicopterVisionRange:                                 c.readFloat64(),
			HelicopterGroundAttackRange:                           c.readFloat64(),
			HelicopterAerialAttackRange:                           c.readFloat64(),
			HelicopterGroundDamage:                                c.readInt(),
			HelicopterAerialDamage:                                c.readInt(),
			HelicopterGroundDefence:                               c.readInt(),
			HelicopterAerialDefence:                               c.readInt(),
			HelicopterAttackCooldownTicks:                         c.readInt(),
			HelicopterProductionCost:                              c.readInt(),
			FighterDurability:                                     c.readInt(),
			FighterSpeed:                                          c.readFloat64(),
			FighterVisionRange:                                    c.readFloat64(),
			FighterGroundAttackRange:                              c.readFloat64(),
			FighterAerialAttackRange:                              c.readFloat64(),
			FighterGroundDamage:                                   c.readInt(),
			FighterAerialDamage:                                   c.readInt(),
			FighterGroundDefence:                                  c.readInt(),
			FighterAerialDefence:                                  c.readInt(),
			FighterAttackCooldownTicks:                            c.readInt(),
			FighterProductionCost:                                 c.readInt(),
			MaxFacilityCapturePoints:                              c.readFloat64(),
			FacilityCapturePointsPerVehiclePerTick:                c.readFloat64(),
			FacilityWidth:                                         c.readFloat64(),
			FacilityHeight:                                        c.readFloat64(),
			BaseTacticalNuclearStrikeCooldown:                     c.readInt(),
			TacticalNuclearStrikeCooldownDecreasePerControlCenter: c.readInt(),
			TacticalNuclearStrikeMaxDamage:                        c.readFloat64(),
			TacticalNuclearStrikeRadius:                           c.readFloat64(),
			TacticalNuclearStrikeDelay:                            c.readInt(),
		}
	}

	return nil
}

func (c *RemoteProcessClient) readContext(pc *PlayerContext) error {
	switch c.readOpcode() {
	case Message_GameOver:
		return ErrGameOver
	case Message_PlayerContext:
		if c.readBool() {
			if me := c.readPlayer(); me != nil {
				*pc.Player = *me
			}
			c.readWorld(pc.World)
		}
		return nil
	default:
		return ErrWrongType
	}
}

func (c *RemoteProcessClient) readPlayer() *Player {
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
		p.RemainingNuclearStrikeCooldownTicks = c.readInt()
		p.NextNuclearStrikeVehicleId = c.readInt64()
		p.NextNuclearStrikeTickIndex = c.readInt()
		p.NextNuclearStrikeX = c.readFloat64()
		p.NextNuclearStrikeY = c.readFloat64()

		c.players[p.Id] = p

		return p
	}
}

func (c *RemoteProcessClient) readWorld(w *World) {
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

func (c *RemoteProcessClient) writeMove(m *Move) {
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
		c.writeFloat64(m.Factor)
		c.writeFloat64(m.MaxSpeed)
		c.writeFloat64(m.MaxAngularSpeed)
		c.writeByte(byte(m.Type))
		c.writeInt64(m.FacilityId)
		c.writeInt64(m.VehicleId)
	}

	c.flush()
}

func (c *RemoteProcessClient) readWeather() (weather [][]Weather) {
	for i := c.readInt(); i > 0; i-- {
		var slice []Weather
		for j := c.readInt(); j > 0; j-- {
			slice = append(slice, Weather(c.readByte()))
		}
		weather = append(weather, slice)
	}

	return
}

func (c *RemoteProcessClient) readTerrains() (terrain [][]Terrain) {
	for i := c.readInt(); i > 0; i-- {
		var slice []Terrain
		for j := c.readInt(); j > 0; j-- {
			slice = append(slice, Terrain(c.readByte()))
		}
		terrain = append(terrain, slice)
	}
	return
}

func (c *RemoteProcessClient) readFacility() *Facility {
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

func (c *RemoteProcessClient) readVehicleUpdate() *VehicleUpdate {
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

func (c *RemoteProcessClient) readNewVehicle() *Vehicle {
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

func (c *RemoteProcessClient) readVehiclesUpdate() (updates []*VehicleUpdate) {
	for l := c.readInt(); l > 0; l-- {
		if v := c.readVehicleUpdate(); v != nil {
			updates = append(updates, v)
		}
	}
	return
}

func (c *RemoteProcessClient) readFacilities() (facilities []*Facility) {
	if l := c.readInt(); l > 0 {
		for ; l > 0; l-- {
			if f := c.readFacility(); f != nil {
				facilities = append(facilities, f)
			}
		}
	} else {
		for _, f := range c.facilities {
			facilities = append(facilities, f)
		}
	}

	return
}

func (c *RemoteProcessClient) readVehicles() (vehicles []*Vehicle) {
	for l := c.readInt(); l > 0; l-- {
		if v := c.readNewVehicle(); v != nil {
			vehicles = append(vehicles, v)
		}
	}
	return
}

func (c *RemoteProcessClient) readPlayers() (players []*Player) {
	if l := c.readInt(); l > 0 {
		for ; l > 0; l-- {
			if p := c.readPlayer(); p != nil {
				players = append(players, p)
			}
		}
	} else {
		for _, p := range c.players {
			players = append(players, p)
		}
	}

	return
}

func (c *RemoteProcessClient) Dial(host, port string) (err error) {
	if c.conn, err = net.Dial("tcp", host+":"+port); err == nil {
		c.reader = bufio.NewReader(c.conn)
		c.writer = bufio.NewWriter(c.conn)
	}

	return
}

func (c *RemoteProcessClient) writeToken(token string) {
	c.writeOpcode(Message_AuthenticationToken)
	c.writeString(token)
	c.flush()
}

func (c *RemoteProcessClient) writeProtoVersion(ver int) {
	c.writeOpcode(Message_ProtocolVersion)
	c.writeInt(ver)
	c.flush()
}

func (c *RemoteProcessClient) ReadTeamSize() int {
	c.ensureMessageType(Message_TeamSize)
	return c.readInt()
}

func (c *RemoteProcessClient) Close() error {
	return c.conn.Close()
}

func (c *RemoteProcessClient) readOpcode() MessageType {
	return MessageType(c.readByte())
}

func (c *RemoteProcessClient) readIntArray() []int {
	var arr []int
	if ln := c.readInt(); ln > 0 {
		for ; ln > 0; ln-- {
			arr = append(arr, c.readInt())
		}
	}
	return arr
}

func (c *RemoteProcessClient) readInt() int {
	var v int32
	if err := binary.Read(c.reader, ByteOrder, &v); err != nil {
		panic(err)
	}
	return int(v)
}

func (c *RemoteProcessClient) readInt64() int64 {
	var v int64
	if err := binary.Read(c.reader, ByteOrder, &v); err != nil {
		panic(err)
	}
	return v
}

func (c *RemoteProcessClient) readFloat64() float64 {
	var v float64
	if err := binary.Read(c.reader, ByteOrder, &v); err != nil {
		panic(err)
	}
	return v
}

func (c *RemoteProcessClient) writeBool(b bool) {
	if b {
		c.writeByte(1)
	} else {
		c.writeByte(0)
	}
}

func (c *RemoteProcessClient) readBool() bool {
	return c.readByte() != 0
}

func (c *RemoteProcessClient) readByte() byte {
	b, err := c.reader.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func (c *RemoteProcessClient) readString() string {
	return string(c.readBytes())
}

func (c *RemoteProcessClient) ensureMessageType(m MessageType) {
	if b, err := c.reader.ReadByte(); err != nil || b != byte(m) {
		panic(ErrWrongType)
	}
}

func (c *RemoteProcessClient) writeOpcode(m MessageType) {
	c.writeByte(byte(m))
}

func (c *RemoteProcessClient) writeInt(v int) {
	if err := binary.Write(c.writer, ByteOrder, int32(v)); err != nil {
		panic(err)
	}
}

func (c *RemoteProcessClient) writeFloat64(v float64) {
	if err := binary.Write(c.writer, ByteOrder, v); err != nil {
		panic(err)
	}
}

func (c *RemoteProcessClient) writeInt64(v int64) {
	if err := binary.Write(c.writer, ByteOrder, v); err != nil {
		panic(err)
	}
}

func (c *RemoteProcessClient) readBytes() []byte {
	l := c.readInt()
	r := make([]byte, l)
	for i := range r {
		r[i] = c.readByte()
	}
	return r
}

func (c *RemoteProcessClient) writeByte(v byte) {
	if err := c.writer.WriteByte(v); err != nil {
		panic(err)
	}
}

func (c *RemoteProcessClient) writeBytes(v []byte) {
	c.writeInt(len(v))
	if _, err := c.writer.Write(v); err != nil {
		panic(err)
	}
}

func (c *RemoteProcessClient) writeString(v string) {
	c.writeInt(len(v))
	if _, err := c.writer.WriteString(v); err != nil {
		panic(err)
	}
}

func (c *RemoteProcessClient) flush() {
	if err := c.writer.Flush(); err != nil {
		panic(err)
	}
}
