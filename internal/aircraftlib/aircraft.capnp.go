package aircraftlib

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

// Constants defined in aircraft.capnp.
const (
	ConstEnum = Airport_jfk
)

// Constants defined in aircraft.capnp.
var (
	ConstDate = Zdate{Struct: capnp.MustUnmarshalRootPtr(x_832bcc6686a26d56[0:24]).Struct()}
	ConstList = Zdate_List{List: capnp.MustUnmarshalRootPtr(x_832bcc6686a26d56[24:64]).List()}
)

type Zdate struct{ capnp.Struct }

func NewZdate(s *capnp.Segment) (Zdate, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Zdate{st}, err
}

func NewRootZdate(s *capnp.Segment) (Zdate, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Zdate{st}, err
}

func ReadRootZdate(msg *capnp.Message) (Zdate, error) {
	root, err := msg.RootPtr()
	return Zdate{root.Struct()}, err
}

func (s Zdate) String() string {
	str, _ := text.Marshal(0xde50aebbad57549d, s.Struct)
	return str
}

func (s Zdate) Year() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s Zdate) SetYear(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s Zdate) Month() uint8 {
	return s.Struct.Uint8(2)
}

func (s Zdate) SetMonth(v uint8) {
	s.Struct.SetUint8(2, v)
}

func (s Zdate) Day() uint8 {
	return s.Struct.Uint8(3)
}

func (s Zdate) SetDay(v uint8) {
	s.Struct.SetUint8(3, v)
}

// Zdate_List is a list of Zdate.
type Zdate_List struct{ capnp.List }

// NewZdate creates a new list of Zdate.
func NewZdate_List(s *capnp.Segment, sz int32) (Zdate_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Zdate_List{l}, err
}

func (s Zdate_List) At(i int) Zdate { return Zdate{s.List.Struct(i)} }

func (s Zdate_List) Set(i int, v Zdate) error { return s.List.SetStruct(i, v.Struct) }

// Zdate_Promise is a wrapper for a Zdate promised by a client call.
type Zdate_Promise struct{ *capnp.Pipeline }

func (p Zdate_Promise) Struct() (Zdate, error) {
	s, err := p.Pipeline.Struct()
	return Zdate{s}, err
}

type Zdata struct{ capnp.Struct }

func NewZdata(s *capnp.Segment) (Zdata, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zdata{st}, err
}

func NewRootZdata(s *capnp.Segment) (Zdata, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zdata{st}, err
}

func ReadRootZdata(msg *capnp.Message) (Zdata, error) {
	root, err := msg.RootPtr()
	return Zdata{root.Struct()}, err
}

func (s Zdata) String() string {
	str, _ := text.Marshal(0xc7da65f9a2f20ba2, s.Struct)
	return str
}

func (s Zdata) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Zdata) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Zdata) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// Zdata_List is a list of Zdata.
type Zdata_List struct{ capnp.List }

// NewZdata creates a new list of Zdata.
func NewZdata_List(s *capnp.Segment, sz int32) (Zdata_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Zdata_List{l}, err
}

func (s Zdata_List) At(i int) Zdata { return Zdata{s.List.Struct(i)} }

func (s Zdata_List) Set(i int, v Zdata) error { return s.List.SetStruct(i, v.Struct) }

// Zdata_Promise is a wrapper for a Zdata promised by a client call.
type Zdata_Promise struct{ *capnp.Pipeline }

func (p Zdata_Promise) Struct() (Zdata, error) {
	s, err := p.Pipeline.Struct()
	return Zdata{s}, err
}

type Airport uint16

// Values of Airport.
const (
	Airport_none Airport = 0
	Airport_jfk  Airport = 1
	Airport_lax  Airport = 2
	Airport_sfo  Airport = 3
	Airport_luv  Airport = 4
	Airport_dfw  Airport = 5
	Airport_test Airport = 6
)

// String returns the enum's constant name.
func (c Airport) String() string {
	switch c {
	case Airport_none:
		return "none"
	case Airport_jfk:
		return "jfk"
	case Airport_lax:
		return "lax"
	case Airport_sfo:
		return "sfo"
	case Airport_luv:
		return "luv"
	case Airport_dfw:
		return "dfw"
	case Airport_test:
		return "test"

	default:
		return ""
	}
}

// AirportFromString returns the enum value with a name,
// or the zero value if there's no such value.
func AirportFromString(c string) Airport {
	switch c {
	case "none":
		return Airport_none
	case "jfk":
		return Airport_jfk
	case "lax":
		return Airport_lax
	case "sfo":
		return Airport_sfo
	case "luv":
		return Airport_luv
	case "dfw":
		return Airport_dfw
	case "test":
		return Airport_test

	default:
		return 0
	}
}

type Airport_List struct{ capnp.List }

func NewAirport_List(s *capnp.Segment, sz int32) (Airport_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Airport_List{l.List}, err
}

func (l Airport_List) At(i int) Airport {
	ul := capnp.UInt16List{List: l.List}
	return Airport(ul.At(i))
}

func (l Airport_List) Set(i int, v Airport) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type PlaneBase struct{ capnp.Struct }

func NewPlaneBase(s *capnp.Segment) (PlaneBase, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2})
	return PlaneBase{st}, err
}

func NewRootPlaneBase(s *capnp.Segment) (PlaneBase, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2})
	return PlaneBase{st}, err
}

func ReadRootPlaneBase(msg *capnp.Message) (PlaneBase, error) {
	root, err := msg.RootPtr()
	return PlaneBase{root.Struct()}, err
}

func (s PlaneBase) String() string {
	str, _ := text.Marshal(0xd8bccf6e60a73791, s.Struct)
	return str
}

func (s PlaneBase) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PlaneBase) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PlaneBase) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s PlaneBase) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s PlaneBase) Homes() (Airport_List, error) {
	p, err := s.Struct.Ptr(1)
	return Airport_List{List: p.List()}, err
}

func (s PlaneBase) HasHomes() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PlaneBase) SetHomes(v Airport_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewHomes sets the homes field to a newly
// allocated Airport_List, preferring placement in s's segment.
func (s PlaneBase) NewHomes(n int32) (Airport_List, error) {
	l, err := NewAirport_List(s.Struct.Segment(), n)
	if err != nil {
		return Airport_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s PlaneBase) Rating() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s PlaneBase) SetRating(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s PlaneBase) CanFly() bool {
	return s.Struct.Bit(64)
}

func (s PlaneBase) SetCanFly(v bool) {
	s.Struct.SetBit(64, v)
}

func (s PlaneBase) Capacity() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s PlaneBase) SetCapacity(v int64) {
	s.Struct.SetUint64(16, uint64(v))
}

func (s PlaneBase) MaxSpeed() float64 {
	return math.Float64frombits(s.Struct.Uint64(24))
}

func (s PlaneBase) SetMaxSpeed(v float64) {
	s.Struct.SetUint64(24, math.Float64bits(v))
}

// PlaneBase_List is a list of PlaneBase.
type PlaneBase_List struct{ capnp.List }

// NewPlaneBase creates a new list of PlaneBase.
func NewPlaneBase_List(s *capnp.Segment, sz int32) (PlaneBase_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2}, sz)
	return PlaneBase_List{l}, err
}

func (s PlaneBase_List) At(i int) PlaneBase { return PlaneBase{s.List.Struct(i)} }

func (s PlaneBase_List) Set(i int, v PlaneBase) error { return s.List.SetStruct(i, v.Struct) }

// PlaneBase_Promise is a wrapper for a PlaneBase promised by a client call.
type PlaneBase_Promise struct{ *capnp.Pipeline }

func (p PlaneBase_Promise) Struct() (PlaneBase, error) {
	s, err := p.Pipeline.Struct()
	return PlaneBase{s}, err
}

type B737 struct{ capnp.Struct }

func NewB737(s *capnp.Segment) (B737, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return B737{st}, err
}

func NewRootB737(s *capnp.Segment) (B737, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return B737{st}, err
}

func ReadRootB737(msg *capnp.Message) (B737, error) {
	root, err := msg.RootPtr()
	return B737{root.Struct()}, err
}

func (s B737) String() string {
	str, _ := text.Marshal(0xccb3b2e3603826e0, s.Struct)
	return str
}

func (s B737) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s B737) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s B737) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s B737) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// B737_List is a list of B737.
type B737_List struct{ capnp.List }

// NewB737 creates a new list of B737.
func NewB737_List(s *capnp.Segment, sz int32) (B737_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return B737_List{l}, err
}

func (s B737_List) At(i int) B737 { return B737{s.List.Struct(i)} }

func (s B737_List) Set(i int, v B737) error { return s.List.SetStruct(i, v.Struct) }

// B737_Promise is a wrapper for a B737 promised by a client call.
type B737_Promise struct{ *capnp.Pipeline }

func (p B737_Promise) Struct() (B737, error) {
	s, err := p.Pipeline.Struct()
	return B737{s}, err
}

func (p B737_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type A320 struct{ capnp.Struct }

func NewA320(s *capnp.Segment) (A320, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return A320{st}, err
}

func NewRootA320(s *capnp.Segment) (A320, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return A320{st}, err
}

func ReadRootA320(msg *capnp.Message) (A320, error) {
	root, err := msg.RootPtr()
	return A320{root.Struct()}, err
}

func (s A320) String() string {
	str, _ := text.Marshal(0xd98c608877d9cb8d, s.Struct)
	return str
}

func (s A320) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s A320) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s A320) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s A320) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// A320_List is a list of A320.
type A320_List struct{ capnp.List }

// NewA320 creates a new list of A320.
func NewA320_List(s *capnp.Segment, sz int32) (A320_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return A320_List{l}, err
}

func (s A320_List) At(i int) A320 { return A320{s.List.Struct(i)} }

func (s A320_List) Set(i int, v A320) error { return s.List.SetStruct(i, v.Struct) }

// A320_Promise is a wrapper for a A320 promised by a client call.
type A320_Promise struct{ *capnp.Pipeline }

func (p A320_Promise) Struct() (A320, error) {
	s, err := p.Pipeline.Struct()
	return A320{s}, err
}

func (p A320_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type F16 struct{ capnp.Struct }

func NewF16(s *capnp.Segment) (F16, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return F16{st}, err
}

func NewRootF16(s *capnp.Segment) (F16, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return F16{st}, err
}

func ReadRootF16(msg *capnp.Message) (F16, error) {
	root, err := msg.RootPtr()
	return F16{root.Struct()}, err
}

func (s F16) String() string {
	str, _ := text.Marshal(0xe1c9eac512335361, s.Struct)
	return str
}

func (s F16) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s F16) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s F16) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s F16) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// F16_List is a list of F16.
type F16_List struct{ capnp.List }

// NewF16 creates a new list of F16.
func NewF16_List(s *capnp.Segment, sz int32) (F16_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return F16_List{l}, err
}

func (s F16_List) At(i int) F16 { return F16{s.List.Struct(i)} }

func (s F16_List) Set(i int, v F16) error { return s.List.SetStruct(i, v.Struct) }

// F16_Promise is a wrapper for a F16 promised by a client call.
type F16_Promise struct{ *capnp.Pipeline }

func (p F16_Promise) Struct() (F16, error) {
	s, err := p.Pipeline.Struct()
	return F16{s}, err
}

func (p F16_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Regression struct{ capnp.Struct }

func NewRegression(s *capnp.Segment) (Regression, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Regression{st}, err
}

func NewRootRegression(s *capnp.Segment) (Regression, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Regression{st}, err
}

func ReadRootRegression(msg *capnp.Message) (Regression, error) {
	root, err := msg.RootPtr()
	return Regression{root.Struct()}, err
}

func (s Regression) String() string {
	str, _ := text.Marshal(0xb1f0385d845e367f, s.Struct)
	return str
}

func (s Regression) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s Regression) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Regression) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s Regression) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Regression) B0() float64 {
	return math.Float64frombits(s.Struct.Uint64(0))
}

func (s Regression) SetB0(v float64) {
	s.Struct.SetUint64(0, math.Float64bits(v))
}

func (s Regression) Beta() (capnp.Float64List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Float64List{List: p.List()}, err
}

func (s Regression) HasBeta() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Regression) SetBeta(v capnp.Float64List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewBeta sets the beta field to a newly
// allocated capnp.Float64List, preferring placement in s's segment.
func (s Regression) NewBeta(n int32) (capnp.Float64List, error) {
	l, err := capnp.NewFloat64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float64List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Regression) Planes() (Aircraft_List, error) {
	p, err := s.Struct.Ptr(2)
	return Aircraft_List{List: p.List()}, err
}

func (s Regression) HasPlanes() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Regression) SetPlanes(v Aircraft_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewPlanes sets the planes field to a newly
// allocated Aircraft_List, preferring placement in s's segment.
func (s Regression) NewPlanes(n int32) (Aircraft_List, error) {
	l, err := NewAircraft_List(s.Struct.Segment(), n)
	if err != nil {
		return Aircraft_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Regression) Ymu() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Regression) SetYmu(v float64) {
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Regression) Ysd() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s Regression) SetYsd(v float64) {
	s.Struct.SetUint64(16, math.Float64bits(v))
}

// Regression_List is a list of Regression.
type Regression_List struct{ capnp.List }

// NewRegression creates a new list of Regression.
func NewRegression_List(s *capnp.Segment, sz int32) (Regression_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	return Regression_List{l}, err
}

func (s Regression_List) At(i int) Regression { return Regression{s.List.Struct(i)} }

func (s Regression_List) Set(i int, v Regression) error { return s.List.SetStruct(i, v.Struct) }

// Regression_Promise is a wrapper for a Regression promised by a client call.
type Regression_Promise struct{ *capnp.Pipeline }

func (p Regression_Promise) Struct() (Regression, error) {
	s, err := p.Pipeline.Struct()
	return Regression{s}, err
}

func (p Regression_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Aircraft struct{ capnp.Struct }
type Aircraft_Which uint16

const (
	Aircraft_Which_void Aircraft_Which = 0
	Aircraft_Which_b737 Aircraft_Which = 1
	Aircraft_Which_a320 Aircraft_Which = 2
	Aircraft_Which_f16  Aircraft_Which = 3
)

func (w Aircraft_Which) String() string {
	const s = "voidb737a320f16"
	switch w {
	case Aircraft_Which_void:
		return s[0:4]
	case Aircraft_Which_b737:
		return s[4:8]
	case Aircraft_Which_a320:
		return s[8:12]
	case Aircraft_Which_f16:
		return s[12:15]

	}
	return "Aircraft_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewAircraft(s *capnp.Segment) (Aircraft, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Aircraft{st}, err
}

func NewRootAircraft(s *capnp.Segment) (Aircraft, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Aircraft{st}, err
}

func ReadRootAircraft(msg *capnp.Message) (Aircraft, error) {
	root, err := msg.RootPtr()
	return Aircraft{root.Struct()}, err
}

func (s Aircraft) String() string {
	str, _ := text.Marshal(0xe54e10aede55c7b1, s.Struct)
	return str
}

func (s Aircraft) Which() Aircraft_Which {
	return Aircraft_Which(s.Struct.Uint16(0))
}
func (s Aircraft) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Aircraft) B737() (B737, error) {
	p, err := s.Struct.Ptr(0)
	return B737{Struct: p.Struct()}, err
}

func (s Aircraft) HasB737() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Aircraft) SetB737(v B737) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewB737 sets the b737 field to a newly
// allocated B737 struct, preferring placement in s's segment.
func (s Aircraft) NewB737() (B737, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewB737(s.Struct.Segment())
	if err != nil {
		return B737{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Aircraft) A320() (A320, error) {
	p, err := s.Struct.Ptr(0)
	return A320{Struct: p.Struct()}, err
}

func (s Aircraft) HasA320() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Aircraft) SetA320(v A320) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewA320 sets the a320 field to a newly
// allocated A320 struct, preferring placement in s's segment.
func (s Aircraft) NewA320() (A320, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewA320(s.Struct.Segment())
	if err != nil {
		return A320{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Aircraft) F16() (F16, error) {
	p, err := s.Struct.Ptr(0)
	return F16{Struct: p.Struct()}, err
}

func (s Aircraft) HasF16() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Aircraft) SetF16(v F16) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewF16 sets the f16 field to a newly
// allocated F16 struct, preferring placement in s's segment.
func (s Aircraft) NewF16() (F16, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewF16(s.Struct.Segment())
	if err != nil {
		return F16{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Aircraft_List is a list of Aircraft.
type Aircraft_List struct{ capnp.List }

// NewAircraft creates a new list of Aircraft.
func NewAircraft_List(s *capnp.Segment, sz int32) (Aircraft_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Aircraft_List{l}, err
}

func (s Aircraft_List) At(i int) Aircraft { return Aircraft{s.List.Struct(i)} }

func (s Aircraft_List) Set(i int, v Aircraft) error { return s.List.SetStruct(i, v.Struct) }

// Aircraft_Promise is a wrapper for a Aircraft promised by a client call.
type Aircraft_Promise struct{ *capnp.Pipeline }

func (p Aircraft_Promise) Struct() (Aircraft, error) {
	s, err := p.Pipeline.Struct()
	return Aircraft{s}, err
}

func (p Aircraft_Promise) B737() B737_Promise {
	return B737_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Aircraft_Promise) A320() A320_Promise {
	return A320_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Aircraft_Promise) F16() F16_Promise {
	return F16_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Z struct{ capnp.Struct }
type Z_Which uint16

const (
	Z_Which_void        Z_Which = 0
	Z_Which_zz          Z_Which = 1
	Z_Which_f64         Z_Which = 2
	Z_Which_f32         Z_Which = 3
	Z_Which_i64         Z_Which = 4
	Z_Which_i32         Z_Which = 5
	Z_Which_i16         Z_Which = 6
	Z_Which_i8          Z_Which = 7
	Z_Which_u64         Z_Which = 8
	Z_Which_u32         Z_Which = 9
	Z_Which_u16         Z_Which = 10
	Z_Which_u8          Z_Which = 11
	Z_Which_bool        Z_Which = 12
	Z_Which_text        Z_Which = 13
	Z_Which_blob        Z_Which = 14
	Z_Which_f64vec      Z_Which = 15
	Z_Which_f32vec      Z_Which = 16
	Z_Which_i64vec      Z_Which = 17
	Z_Which_i32vec      Z_Which = 18
	Z_Which_i16vec      Z_Which = 19
	Z_Which_i8vec       Z_Which = 20
	Z_Which_u64vec      Z_Which = 21
	Z_Which_u32vec      Z_Which = 22
	Z_Which_u16vec      Z_Which = 23
	Z_Which_u8vec       Z_Which = 24
	Z_Which_zvec        Z_Which = 25
	Z_Which_zvecvec     Z_Which = 26
	Z_Which_zdate       Z_Which = 27
	Z_Which_zdata       Z_Which = 28
	Z_Which_aircraftvec Z_Which = 29
	Z_Which_aircraft    Z_Which = 30
	Z_Which_regression  Z_Which = 31
	Z_Which_planebase   Z_Which = 32
	Z_Which_airport     Z_Which = 33
	Z_Which_b737        Z_Which = 34
	Z_Which_a320        Z_Which = 35
	Z_Which_f16         Z_Which = 36
	Z_Which_zdatevec    Z_Which = 37
	Z_Which_zdatavec    Z_Which = 38
	Z_Which_boolvec     Z_Which = 39
)

func (w Z_Which) String() string {
	const s = "voidzzf64f32i64i32i16i8u64u32u16u8booltextblobf64vecf32veci64veci32veci16veci8vecu64vecu32vecu16vecu8veczveczvecveczdatezdataaircraftvecaircraftregressionplanebaseairportb737a320f16zdateveczdatavecboolvec"
	switch w {
	case Z_Which_void:
		return s[0:4]
	case Z_Which_zz:
		return s[4:6]
	case Z_Which_f64:
		return s[6:9]
	case Z_Which_f32:
		return s[9:12]
	case Z_Which_i64:
		return s[12:15]
	case Z_Which_i32:
		return s[15:18]
	case Z_Which_i16:
		return s[18:21]
	case Z_Which_i8:
		return s[21:23]
	case Z_Which_u64:
		return s[23:26]
	case Z_Which_u32:
		return s[26:29]
	case Z_Which_u16:
		return s[29:32]
	case Z_Which_u8:
		return s[32:34]
	case Z_Which_bool:
		return s[34:38]
	case Z_Which_text:
		return s[38:42]
	case Z_Which_blob:
		return s[42:46]
	case Z_Which_f64vec:
		return s[46:52]
	case Z_Which_f32vec:
		return s[52:58]
	case Z_Which_i64vec:
		return s[58:64]
	case Z_Which_i32vec:
		return s[64:70]
	case Z_Which_i16vec:
		return s[70:76]
	case Z_Which_i8vec:
		return s[76:81]
	case Z_Which_u64vec:
		return s[81:87]
	case Z_Which_u32vec:
		return s[87:93]
	case Z_Which_u16vec:
		return s[93:99]
	case Z_Which_u8vec:
		return s[99:104]
	case Z_Which_zvec:
		return s[104:108]
	case Z_Which_zvecvec:
		return s[108:115]
	case Z_Which_zdate:
		return s[115:120]
	case Z_Which_zdata:
		return s[120:125]
	case Z_Which_aircraftvec:
		return s[125:136]
	case Z_Which_aircraft:
		return s[136:144]
	case Z_Which_regression:
		return s[144:154]
	case Z_Which_planebase:
		return s[154:163]
	case Z_Which_airport:
		return s[163:170]
	case Z_Which_b737:
		return s[170:174]
	case Z_Which_a320:
		return s[174:178]
	case Z_Which_f16:
		return s[178:181]
	case Z_Which_zdatevec:
		return s[181:189]
	case Z_Which_zdatavec:
		return s[189:197]
	case Z_Which_boolvec:
		return s[197:204]

	}
	return "Z_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewZ(s *capnp.Segment) (Z, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Z{st}, err
}

func NewRootZ(s *capnp.Segment) (Z, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Z{st}, err
}

func ReadRootZ(msg *capnp.Message) (Z, error) {
	root, err := msg.RootPtr()
	return Z{root.Struct()}, err
}

func (s Z) String() string {
	str, _ := text.Marshal(0xea26e9973bd6a0d9, s.Struct)
	return str
}

func (s Z) Which() Z_Which {
	return Z_Which(s.Struct.Uint16(0))
}
func (s Z) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Z) Zz() (Z, error) {
	p, err := s.Struct.Ptr(0)
	return Z{Struct: p.Struct()}, err
}

func (s Z) HasZz() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZz(v Z) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewZz sets the zz field to a newly
// allocated Z struct, preferring placement in s's segment.
func (s Z) NewZz() (Z, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewZ(s.Struct.Segment())
	if err != nil {
		return Z{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) F64() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Z) SetF64(v float64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Z) F32() float32 {
	return math.Float32frombits(s.Struct.Uint32(8))
}

func (s Z) SetF32(v float32) {
	s.Struct.SetUint16(0, 3)
	s.Struct.SetUint32(8, math.Float32bits(v))
}

func (s Z) I64() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Z) SetI64(v int64) {
	s.Struct.SetUint16(0, 4)
	s.Struct.SetUint64(8, uint64(v))
}

func (s Z) I32() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s Z) SetI32(v int32) {
	s.Struct.SetUint16(0, 5)
	s.Struct.SetUint32(8, uint32(v))
}

func (s Z) I16() int16 {
	return int16(s.Struct.Uint16(8))
}

func (s Z) SetI16(v int16) {
	s.Struct.SetUint16(0, 6)
	s.Struct.SetUint16(8, uint16(v))
}

func (s Z) I8() int8 {
	return int8(s.Struct.Uint8(8))
}

func (s Z) SetI8(v int8) {
	s.Struct.SetUint16(0, 7)
	s.Struct.SetUint8(8, uint8(v))
}

func (s Z) U64() uint64 {
	return s.Struct.Uint64(8)
}

func (s Z) SetU64(v uint64) {
	s.Struct.SetUint16(0, 8)
	s.Struct.SetUint64(8, v)
}

func (s Z) U32() uint32 {
	return s.Struct.Uint32(8)
}

func (s Z) SetU32(v uint32) {
	s.Struct.SetUint16(0, 9)
	s.Struct.SetUint32(8, v)
}

func (s Z) U16() uint16 {
	return s.Struct.Uint16(8)
}

func (s Z) SetU16(v uint16) {
	s.Struct.SetUint16(0, 10)
	s.Struct.SetUint16(8, v)
}

func (s Z) U8() uint8 {
	return s.Struct.Uint8(8)
}

func (s Z) SetU8(v uint8) {
	s.Struct.SetUint16(0, 11)
	s.Struct.SetUint8(8, v)
}

func (s Z) Bool() bool {
	return s.Struct.Bit(64)
}

func (s Z) SetBool(v bool) {
	s.Struct.SetUint16(0, 12)
	s.Struct.SetBit(64, v)
}

func (s Z) Text() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Z) HasText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Z) SetText(v string) error {
	s.Struct.SetUint16(0, 13)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Z) Blob() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Z) HasBlob() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetBlob(v []byte) error {
	s.Struct.SetUint16(0, 14)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Z) F64vec() (capnp.Float64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Float64List{List: p.List()}, err
}

func (s Z) HasF64vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetF64vec(v capnp.Float64List) error {
	s.Struct.SetUint16(0, 15)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewF64vec sets the f64vec field to a newly
// allocated capnp.Float64List, preferring placement in s's segment.
func (s Z) NewF64vec(n int32) (capnp.Float64List, error) {
	s.Struct.SetUint16(0, 15)
	l, err := capnp.NewFloat64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) F32vec() (capnp.Float32List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Float32List{List: p.List()}, err
}

func (s Z) HasF32vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetF32vec(v capnp.Float32List) error {
	s.Struct.SetUint16(0, 16)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewF32vec sets the f32vec field to a newly
// allocated capnp.Float32List, preferring placement in s's segment.
func (s Z) NewF32vec(n int32) (capnp.Float32List, error) {
	s.Struct.SetUint16(0, 16)
	l, err := capnp.NewFloat32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float32List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I64vec() (capnp.Int64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int64List{List: p.List()}, err
}

func (s Z) HasI64vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI64vec(v capnp.Int64List) error {
	s.Struct.SetUint16(0, 17)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI64vec sets the i64vec field to a newly
// allocated capnp.Int64List, preferring placement in s's segment.
func (s Z) NewI64vec(n int32) (capnp.Int64List, error) {
	s.Struct.SetUint16(0, 17)
	l, err := capnp.NewInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I32vec() (capnp.Int32List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int32List{List: p.List()}, err
}

func (s Z) HasI32vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI32vec(v capnp.Int32List) error {
	s.Struct.SetUint16(0, 18)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI32vec sets the i32vec field to a newly
// allocated capnp.Int32List, preferring placement in s's segment.
func (s Z) NewI32vec(n int32) (capnp.Int32List, error) {
	s.Struct.SetUint16(0, 18)
	l, err := capnp.NewInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int32List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I16vec() (capnp.Int16List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int16List{List: p.List()}, err
}

func (s Z) HasI16vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI16vec(v capnp.Int16List) error {
	s.Struct.SetUint16(0, 19)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI16vec sets the i16vec field to a newly
// allocated capnp.Int16List, preferring placement in s's segment.
func (s Z) NewI16vec(n int32) (capnp.Int16List, error) {
	s.Struct.SetUint16(0, 19)
	l, err := capnp.NewInt16List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int16List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I8vec() (capnp.Int8List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int8List{List: p.List()}, err
}

func (s Z) HasI8vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI8vec(v capnp.Int8List) error {
	s.Struct.SetUint16(0, 20)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI8vec sets the i8vec field to a newly
// allocated capnp.Int8List, preferring placement in s's segment.
func (s Z) NewI8vec(n int32) (capnp.Int8List, error) {
	s.Struct.SetUint16(0, 20)
	l, err := capnp.NewInt8List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int8List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U64vec() (capnp.UInt64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt64List{List: p.List()}, err
}

func (s Z) HasU64vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU64vec(v capnp.UInt64List) error {
	s.Struct.SetUint16(0, 21)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU64vec sets the u64vec field to a newly
// allocated capnp.UInt64List, preferring placement in s's segment.
func (s Z) NewU64vec(n int32) (capnp.UInt64List, error) {
	s.Struct.SetUint16(0, 21)
	l, err := capnp.NewUInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U32vec() (capnp.UInt32List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt32List{List: p.List()}, err
}

func (s Z) HasU32vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU32vec(v capnp.UInt32List) error {
	s.Struct.SetUint16(0, 22)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU32vec sets the u32vec field to a newly
// allocated capnp.UInt32List, preferring placement in s's segment.
func (s Z) NewU32vec(n int32) (capnp.UInt32List, error) {
	s.Struct.SetUint16(0, 22)
	l, err := capnp.NewUInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt32List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U16vec() (capnp.UInt16List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt16List{List: p.List()}, err
}

func (s Z) HasU16vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU16vec(v capnp.UInt16List) error {
	s.Struct.SetUint16(0, 23)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU16vec sets the u16vec field to a newly
// allocated capnp.UInt16List, preferring placement in s's segment.
func (s Z) NewU16vec(n int32) (capnp.UInt16List, error) {
	s.Struct.SetUint16(0, 23)
	l, err := capnp.NewUInt16List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt16List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U8vec() (capnp.UInt8List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt8List{List: p.List()}, err
}

func (s Z) HasU8vec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU8vec(v capnp.UInt8List) error {
	s.Struct.SetUint16(0, 24)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU8vec sets the u8vec field to a newly
// allocated capnp.UInt8List, preferring placement in s's segment.
func (s Z) NewU8vec(n int32) (capnp.UInt8List, error) {
	s.Struct.SetUint16(0, 24)
	l, err := capnp.NewUInt8List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt8List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zvec() (Z_List, error) {
	p, err := s.Struct.Ptr(0)
	return Z_List{List: p.List()}, err
}

func (s Z) HasZvec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZvec(v Z_List) error {
	s.Struct.SetUint16(0, 25)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZvec sets the zvec field to a newly
// allocated Z_List, preferring placement in s's segment.
func (s Z) NewZvec(n int32) (Z_List, error) {
	s.Struct.SetUint16(0, 25)
	l, err := NewZ_List(s.Struct.Segment(), n)
	if err != nil {
		return Z_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zvecvec() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.PointerList{List: p.List()}, err
}

func (s Z) HasZvecvec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZvecvec(v capnp.PointerList) error {
	s.Struct.SetUint16(0, 26)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZvecvec sets the zvecvec field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s Z) NewZvecvec(n int32) (capnp.PointerList, error) {
	s.Struct.SetUint16(0, 26)
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zdate() (Zdate, error) {
	p, err := s.Struct.Ptr(0)
	return Zdate{Struct: p.Struct()}, err
}

func (s Z) HasZdate() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdate(v Zdate) error {
	s.Struct.SetUint16(0, 27)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewZdate sets the zdate field to a newly
// allocated Zdate struct, preferring placement in s's segment.
func (s Z) NewZdate() (Zdate, error) {
	s.Struct.SetUint16(0, 27)
	ss, err := NewZdate(s.Struct.Segment())
	if err != nil {
		return Zdate{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Zdata() (Zdata, error) {
	p, err := s.Struct.Ptr(0)
	return Zdata{Struct: p.Struct()}, err
}

func (s Z) HasZdata() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdata(v Zdata) error {
	s.Struct.SetUint16(0, 28)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewZdata sets the zdata field to a newly
// allocated Zdata struct, preferring placement in s's segment.
func (s Z) NewZdata() (Zdata, error) {
	s.Struct.SetUint16(0, 28)
	ss, err := NewZdata(s.Struct.Segment())
	if err != nil {
		return Zdata{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Aircraftvec() (Aircraft_List, error) {
	p, err := s.Struct.Ptr(0)
	return Aircraft_List{List: p.List()}, err
}

func (s Z) HasAircraftvec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetAircraftvec(v Aircraft_List) error {
	s.Struct.SetUint16(0, 29)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewAircraftvec sets the aircraftvec field to a newly
// allocated Aircraft_List, preferring placement in s's segment.
func (s Z) NewAircraftvec(n int32) (Aircraft_List, error) {
	s.Struct.SetUint16(0, 29)
	l, err := NewAircraft_List(s.Struct.Segment(), n)
	if err != nil {
		return Aircraft_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Aircraft() (Aircraft, error) {
	p, err := s.Struct.Ptr(0)
	return Aircraft{Struct: p.Struct()}, err
}

func (s Z) HasAircraft() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetAircraft(v Aircraft) error {
	s.Struct.SetUint16(0, 30)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAircraft sets the aircraft field to a newly
// allocated Aircraft struct, preferring placement in s's segment.
func (s Z) NewAircraft() (Aircraft, error) {
	s.Struct.SetUint16(0, 30)
	ss, err := NewAircraft(s.Struct.Segment())
	if err != nil {
		return Aircraft{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Regression() (Regression, error) {
	p, err := s.Struct.Ptr(0)
	return Regression{Struct: p.Struct()}, err
}

func (s Z) HasRegression() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetRegression(v Regression) error {
	s.Struct.SetUint16(0, 31)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRegression sets the regression field to a newly
// allocated Regression struct, preferring placement in s's segment.
func (s Z) NewRegression() (Regression, error) {
	s.Struct.SetUint16(0, 31)
	ss, err := NewRegression(s.Struct.Segment())
	if err != nil {
		return Regression{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Planebase() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s Z) HasPlanebase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetPlanebase(v PlaneBase) error {
	s.Struct.SetUint16(0, 32)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPlanebase sets the planebase field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s Z) NewPlanebase() (PlaneBase, error) {
	s.Struct.SetUint16(0, 32)
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Airport() Airport {
	return Airport(s.Struct.Uint16(8))
}

func (s Z) SetAirport(v Airport) {
	s.Struct.SetUint16(0, 33)
	s.Struct.SetUint16(8, uint16(v))
}

func (s Z) B737() (B737, error) {
	p, err := s.Struct.Ptr(0)
	return B737{Struct: p.Struct()}, err
}

func (s Z) HasB737() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetB737(v B737) error {
	s.Struct.SetUint16(0, 34)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewB737 sets the b737 field to a newly
// allocated B737 struct, preferring placement in s's segment.
func (s Z) NewB737() (B737, error) {
	s.Struct.SetUint16(0, 34)
	ss, err := NewB737(s.Struct.Segment())
	if err != nil {
		return B737{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) A320() (A320, error) {
	p, err := s.Struct.Ptr(0)
	return A320{Struct: p.Struct()}, err
}

func (s Z) HasA320() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetA320(v A320) error {
	s.Struct.SetUint16(0, 35)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewA320 sets the a320 field to a newly
// allocated A320 struct, preferring placement in s's segment.
func (s Z) NewA320() (A320, error) {
	s.Struct.SetUint16(0, 35)
	ss, err := NewA320(s.Struct.Segment())
	if err != nil {
		return A320{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) F16() (F16, error) {
	p, err := s.Struct.Ptr(0)
	return F16{Struct: p.Struct()}, err
}

func (s Z) HasF16() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetF16(v F16) error {
	s.Struct.SetUint16(0, 36)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewF16 sets the f16 field to a newly
// allocated F16 struct, preferring placement in s's segment.
func (s Z) NewF16() (F16, error) {
	s.Struct.SetUint16(0, 36)
	ss, err := NewF16(s.Struct.Segment())
	if err != nil {
		return F16{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Zdatevec() (Zdate_List, error) {
	p, err := s.Struct.Ptr(0)
	return Zdate_List{List: p.List()}, err
}

func (s Z) HasZdatevec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdatevec(v Zdate_List) error {
	s.Struct.SetUint16(0, 37)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZdatevec sets the zdatevec field to a newly
// allocated Zdate_List, preferring placement in s's segment.
func (s Z) NewZdatevec(n int32) (Zdate_List, error) {
	s.Struct.SetUint16(0, 37)
	l, err := NewZdate_List(s.Struct.Segment(), n)
	if err != nil {
		return Zdate_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zdatavec() (Zdata_List, error) {
	p, err := s.Struct.Ptr(0)
	return Zdata_List{List: p.List()}, err
}

func (s Z) HasZdatavec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdatavec(v Zdata_List) error {
	s.Struct.SetUint16(0, 38)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZdatavec sets the zdatavec field to a newly
// allocated Zdata_List, preferring placement in s's segment.
func (s Z) NewZdatavec(n int32) (Zdata_List, error) {
	s.Struct.SetUint16(0, 38)
	l, err := NewZdata_List(s.Struct.Segment(), n)
	if err != nil {
		return Zdata_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Boolvec() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.BitList{List: p.List()}, err
}

func (s Z) HasBoolvec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetBoolvec(v capnp.BitList) error {
	s.Struct.SetUint16(0, 39)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewBoolvec sets the boolvec field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s Z) NewBoolvec(n int32) (capnp.BitList, error) {
	s.Struct.SetUint16(0, 39)
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Z_List is a list of Z.
type Z_List struct{ capnp.List }

// NewZ creates a new list of Z.
func NewZ_List(s *capnp.Segment, sz int32) (Z_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Z_List{l}, err
}

func (s Z_List) At(i int) Z { return Z{s.List.Struct(i)} }

func (s Z_List) Set(i int, v Z) error { return s.List.SetStruct(i, v.Struct) }

// Z_Promise is a wrapper for a Z promised by a client call.
type Z_Promise struct{ *capnp.Pipeline }

func (p Z_Promise) Struct() (Z, error) {
	s, err := p.Pipeline.Struct()
	return Z{s}, err
}

func (p Z_Promise) Zz() Z_Promise {
	return Z_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Zdate() Zdate_Promise {
	return Zdate_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Zdata() Zdata_Promise {
	return Zdata_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Aircraft() Aircraft_Promise {
	return Aircraft_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Regression() Regression_Promise {
	return Regression_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Planebase() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) B737() B737_Promise {
	return B737_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) A320() A320_Promise {
	return A320_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) F16() F16_Promise {
	return F16_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Counter struct{ capnp.Struct }

func NewCounter(s *capnp.Segment) (Counter, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Counter{st}, err
}

func NewRootCounter(s *capnp.Segment) (Counter, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Counter{st}, err
}

func ReadRootCounter(msg *capnp.Message) (Counter, error) {
	root, err := msg.RootPtr()
	return Counter{root.Struct()}, err
}

func (s Counter) String() string {
	str, _ := text.Marshal(0x8748bc095e10cb5d, s.Struct)
	return str
}

func (s Counter) Size() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Counter) SetSize(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Counter) Words() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Counter) HasWords() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Counter) WordsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Counter) SetWords(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Counter) Wordlist() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s Counter) HasWordlist() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Counter) SetWordlist(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewWordlist sets the wordlist field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Counter) NewWordlist(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Counter_List is a list of Counter.
type Counter_List struct{ capnp.List }

// NewCounter creates a new list of Counter.
func NewCounter_List(s *capnp.Segment, sz int32) (Counter_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Counter_List{l}, err
}

func (s Counter_List) At(i int) Counter { return Counter{s.List.Struct(i)} }

func (s Counter_List) Set(i int, v Counter) error { return s.List.SetStruct(i, v.Struct) }

// Counter_Promise is a wrapper for a Counter promised by a client call.
type Counter_Promise struct{ *capnp.Pipeline }

func (p Counter_Promise) Struct() (Counter, error) {
	s, err := p.Pipeline.Struct()
	return Counter{s}, err
}

type Bag struct{ capnp.Struct }

func NewBag(s *capnp.Segment) (Bag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Bag{st}, err
}

func NewRootBag(s *capnp.Segment) (Bag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Bag{st}, err
}

func ReadRootBag(msg *capnp.Message) (Bag, error) {
	root, err := msg.RootPtr()
	return Bag{root.Struct()}, err
}

func (s Bag) String() string {
	str, _ := text.Marshal(0xd636fba4f188dabe, s.Struct)
	return str
}

func (s Bag) Counter() (Counter, error) {
	p, err := s.Struct.Ptr(0)
	return Counter{Struct: p.Struct()}, err
}

func (s Bag) HasCounter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Bag) SetCounter(v Counter) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCounter sets the counter field to a newly
// allocated Counter struct, preferring placement in s's segment.
func (s Bag) NewCounter() (Counter, error) {
	ss, err := NewCounter(s.Struct.Segment())
	if err != nil {
		return Counter{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Bag_List is a list of Bag.
type Bag_List struct{ capnp.List }

// NewBag creates a new list of Bag.
func NewBag_List(s *capnp.Segment, sz int32) (Bag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Bag_List{l}, err
}

func (s Bag_List) At(i int) Bag { return Bag{s.List.Struct(i)} }

func (s Bag_List) Set(i int, v Bag) error { return s.List.SetStruct(i, v.Struct) }

// Bag_Promise is a wrapper for a Bag promised by a client call.
type Bag_Promise struct{ *capnp.Pipeline }

func (p Bag_Promise) Struct() (Bag, error) {
	s, err := p.Pipeline.Struct()
	return Bag{s}, err
}

func (p Bag_Promise) Counter() Counter_Promise {
	return Counter_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Zserver struct{ capnp.Struct }

func NewZserver(s *capnp.Segment) (Zserver, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zserver{st}, err
}

func NewRootZserver(s *capnp.Segment) (Zserver, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zserver{st}, err
}

func ReadRootZserver(msg *capnp.Message) (Zserver, error) {
	root, err := msg.RootPtr()
	return Zserver{root.Struct()}, err
}

func (s Zserver) String() string {
	str, _ := text.Marshal(0xcc4411e60ba9c498, s.Struct)
	return str
}

func (s Zserver) Waitingjobs() (Zjob_List, error) {
	p, err := s.Struct.Ptr(0)
	return Zjob_List{List: p.List()}, err
}

func (s Zserver) HasWaitingjobs() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Zserver) SetWaitingjobs(v Zjob_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewWaitingjobs sets the waitingjobs field to a newly
// allocated Zjob_List, preferring placement in s's segment.
func (s Zserver) NewWaitingjobs(n int32) (Zjob_List, error) {
	l, err := NewZjob_List(s.Struct.Segment(), n)
	if err != nil {
		return Zjob_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Zserver_List is a list of Zserver.
type Zserver_List struct{ capnp.List }

// NewZserver creates a new list of Zserver.
func NewZserver_List(s *capnp.Segment, sz int32) (Zserver_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Zserver_List{l}, err
}

func (s Zserver_List) At(i int) Zserver { return Zserver{s.List.Struct(i)} }

func (s Zserver_List) Set(i int, v Zserver) error { return s.List.SetStruct(i, v.Struct) }

// Zserver_Promise is a wrapper for a Zserver promised by a client call.
type Zserver_Promise struct{ *capnp.Pipeline }

func (p Zserver_Promise) Struct() (Zserver, error) {
	s, err := p.Pipeline.Struct()
	return Zserver{s}, err
}

type Zjob struct{ capnp.Struct }

func NewZjob(s *capnp.Segment) (Zjob, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Zjob{st}, err
}

func NewRootZjob(s *capnp.Segment) (Zjob, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Zjob{st}, err
}

func ReadRootZjob(msg *capnp.Message) (Zjob, error) {
	root, err := msg.RootPtr()
	return Zjob{root.Struct()}, err
}

func (s Zjob) String() string {
	str, _ := text.Marshal(0xddd1416669fb7613, s.Struct)
	return str
}

func (s Zjob) Cmd() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Zjob) HasCmd() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Zjob) CmdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Zjob) SetCmd(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Zjob) Args() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s Zjob) HasArgs() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Zjob) SetArgs(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewArgs sets the args field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Zjob) NewArgs(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Zjob_List is a list of Zjob.
type Zjob_List struct{ capnp.List }

// NewZjob creates a new list of Zjob.
func NewZjob_List(s *capnp.Segment, sz int32) (Zjob_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Zjob_List{l}, err
}

func (s Zjob_List) At(i int) Zjob { return Zjob{s.List.Struct(i)} }

func (s Zjob_List) Set(i int, v Zjob) error { return s.List.SetStruct(i, v.Struct) }

// Zjob_Promise is a wrapper for a Zjob promised by a client call.
type Zjob_Promise struct{ *capnp.Pipeline }

func (p Zjob_Promise) Struct() (Zjob, error) {
	s, err := p.Pipeline.Struct()
	return Zjob{s}, err
}

type VerEmpty struct{ capnp.Struct }

func NewVerEmpty(s *capnp.Segment) (VerEmpty, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return VerEmpty{st}, err
}

func NewRootVerEmpty(s *capnp.Segment) (VerEmpty, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return VerEmpty{st}, err
}

func ReadRootVerEmpty(msg *capnp.Message) (VerEmpty, error) {
	root, err := msg.RootPtr()
	return VerEmpty{root.Struct()}, err
}

func (s VerEmpty) String() string {
	str, _ := text.Marshal(0x93c99951eacc72ff, s.Struct)
	return str
}

// VerEmpty_List is a list of VerEmpty.
type VerEmpty_List struct{ capnp.List }

// NewVerEmpty creates a new list of VerEmpty.
func NewVerEmpty_List(s *capnp.Segment, sz int32) (VerEmpty_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return VerEmpty_List{l}, err
}

func (s VerEmpty_List) At(i int) VerEmpty { return VerEmpty{s.List.Struct(i)} }

func (s VerEmpty_List) Set(i int, v VerEmpty) error { return s.List.SetStruct(i, v.Struct) }

// VerEmpty_Promise is a wrapper for a VerEmpty promised by a client call.
type VerEmpty_Promise struct{ *capnp.Pipeline }

func (p VerEmpty_Promise) Struct() (VerEmpty, error) {
	s, err := p.Pipeline.Struct()
	return VerEmpty{s}, err
}

type VerOneData struct{ capnp.Struct }

func NewVerOneData(s *capnp.Segment) (VerOneData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VerOneData{st}, err
}

func NewRootVerOneData(s *capnp.Segment) (VerOneData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VerOneData{st}, err
}

func ReadRootVerOneData(msg *capnp.Message) (VerOneData, error) {
	root, err := msg.RootPtr()
	return VerOneData{root.Struct()}, err
}

func (s VerOneData) String() string {
	str, _ := text.Marshal(0xfca3742893be4cde, s.Struct)
	return str
}

func (s VerOneData) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerOneData) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

// VerOneData_List is a list of VerOneData.
type VerOneData_List struct{ capnp.List }

// NewVerOneData creates a new list of VerOneData.
func NewVerOneData_List(s *capnp.Segment, sz int32) (VerOneData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return VerOneData_List{l}, err
}

func (s VerOneData_List) At(i int) VerOneData { return VerOneData{s.List.Struct(i)} }

func (s VerOneData_List) Set(i int, v VerOneData) error { return s.List.SetStruct(i, v.Struct) }

// VerOneData_Promise is a wrapper for a VerOneData promised by a client call.
type VerOneData_Promise struct{ *capnp.Pipeline }

func (p VerOneData_Promise) Struct() (VerOneData, error) {
	s, err := p.Pipeline.Struct()
	return VerOneData{s}, err
}

type VerTwoData struct{ capnp.Struct }

func NewVerTwoData(s *capnp.Segment) (VerTwoData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return VerTwoData{st}, err
}

func NewRootVerTwoData(s *capnp.Segment) (VerTwoData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return VerTwoData{st}, err
}

func ReadRootVerTwoData(msg *capnp.Message) (VerTwoData, error) {
	root, err := msg.RootPtr()
	return VerTwoData{root.Struct()}, err
}

func (s VerTwoData) String() string {
	str, _ := text.Marshal(0xf705dc45c94766fd, s.Struct)
	return str
}

func (s VerTwoData) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoData) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoData) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoData) SetDuo(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

// VerTwoData_List is a list of VerTwoData.
type VerTwoData_List struct{ capnp.List }

// NewVerTwoData creates a new list of VerTwoData.
func NewVerTwoData_List(s *capnp.Segment, sz int32) (VerTwoData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return VerTwoData_List{l}, err
}

func (s VerTwoData_List) At(i int) VerTwoData { return VerTwoData{s.List.Struct(i)} }

func (s VerTwoData_List) Set(i int, v VerTwoData) error { return s.List.SetStruct(i, v.Struct) }

// VerTwoData_Promise is a wrapper for a VerTwoData promised by a client call.
type VerTwoData_Promise struct{ *capnp.Pipeline }

func (p VerTwoData_Promise) Struct() (VerTwoData, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoData{s}, err
}

type VerOnePtr struct{ capnp.Struct }

func NewVerOnePtr(s *capnp.Segment) (VerOnePtr, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return VerOnePtr{st}, err
}

func NewRootVerOnePtr(s *capnp.Segment) (VerOnePtr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return VerOnePtr{st}, err
}

func ReadRootVerOnePtr(msg *capnp.Message) (VerOnePtr, error) {
	root, err := msg.RootPtr()
	return VerOnePtr{root.Struct()}, err
}

func (s VerOnePtr) String() string {
	str, _ := text.Marshal(0x94bf7df83408218d, s.Struct)
	return str
}

func (s VerOnePtr) Ptr() (VerOneData, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerOnePtr) HasPtr() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerOnePtr) SetPtr(v VerOneData) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr sets the ptr field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerOnePtr) NewPtr() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// VerOnePtr_List is a list of VerOnePtr.
type VerOnePtr_List struct{ capnp.List }

// NewVerOnePtr creates a new list of VerOnePtr.
func NewVerOnePtr_List(s *capnp.Segment, sz int32) (VerOnePtr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return VerOnePtr_List{l}, err
}

func (s VerOnePtr_List) At(i int) VerOnePtr { return VerOnePtr{s.List.Struct(i)} }

func (s VerOnePtr_List) Set(i int, v VerOnePtr) error { return s.List.SetStruct(i, v.Struct) }

// VerOnePtr_Promise is a wrapper for a VerOnePtr promised by a client call.
type VerOnePtr_Promise struct{ *capnp.Pipeline }

func (p VerOnePtr_Promise) Struct() (VerOnePtr, error) {
	s, err := p.Pipeline.Struct()
	return VerOnePtr{s}, err
}

func (p VerOnePtr_Promise) Ptr() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type VerTwoPtr struct{ capnp.Struct }

func NewVerTwoPtr(s *capnp.Segment) (VerTwoPtr, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerTwoPtr{st}, err
}

func NewRootVerTwoPtr(s *capnp.Segment) (VerTwoPtr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerTwoPtr{st}, err
}

func ReadRootVerTwoPtr(msg *capnp.Message) (VerTwoPtr, error) {
	root, err := msg.RootPtr()
	return VerTwoPtr{root.Struct()}, err
}

func (s VerTwoPtr) String() string {
	str, _ := text.Marshal(0xc95babe3bd394d2d, s.Struct)
	return str
}

func (s VerTwoPtr) Ptr1() (VerOneData, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoPtr) HasPtr1() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerTwoPtr) SetPtr1(v VerOneData) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoPtr) NewPtr1() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoPtr) Ptr2() (VerOneData, error) {
	p, err := s.Struct.Ptr(1)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoPtr) HasPtr2() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerTwoPtr) SetPtr2(v VerOneData) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoPtr) NewPtr2() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// VerTwoPtr_List is a list of VerTwoPtr.
type VerTwoPtr_List struct{ capnp.List }

// NewVerTwoPtr creates a new list of VerTwoPtr.
func NewVerTwoPtr_List(s *capnp.Segment, sz int32) (VerTwoPtr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return VerTwoPtr_List{l}, err
}

func (s VerTwoPtr_List) At(i int) VerTwoPtr { return VerTwoPtr{s.List.Struct(i)} }

func (s VerTwoPtr_List) Set(i int, v VerTwoPtr) error { return s.List.SetStruct(i, v.Struct) }

// VerTwoPtr_Promise is a wrapper for a VerTwoPtr promised by a client call.
type VerTwoPtr_Promise struct{ *capnp.Pipeline }

func (p VerTwoPtr_Promise) Struct() (VerTwoPtr, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoPtr{s}, err
}

func (p VerTwoPtr_Promise) Ptr1() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoPtr_Promise) Ptr2() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type VerTwoDataTwoPtr struct{ capnp.Struct }

func NewVerTwoDataTwoPtr(s *capnp.Segment) (VerTwoDataTwoPtr, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return VerTwoDataTwoPtr{st}, err
}

func NewRootVerTwoDataTwoPtr(s *capnp.Segment) (VerTwoDataTwoPtr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return VerTwoDataTwoPtr{st}, err
}

func ReadRootVerTwoDataTwoPtr(msg *capnp.Message) (VerTwoDataTwoPtr, error) {
	root, err := msg.RootPtr()
	return VerTwoDataTwoPtr{root.Struct()}, err
}

func (s VerTwoDataTwoPtr) String() string {
	str, _ := text.Marshal(0xb61ee2ecff34ca73, s.Struct)
	return str
}

func (s VerTwoDataTwoPtr) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoDataTwoPtr) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoDataTwoPtr) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoDataTwoPtr) SetDuo(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s VerTwoDataTwoPtr) Ptr1() (VerOneData, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoDataTwoPtr) HasPtr1() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerTwoDataTwoPtr) SetPtr1(v VerOneData) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoDataTwoPtr) NewPtr1() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoDataTwoPtr) Ptr2() (VerOneData, error) {
	p, err := s.Struct.Ptr(1)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoDataTwoPtr) HasPtr2() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerTwoDataTwoPtr) SetPtr2(v VerOneData) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoDataTwoPtr) NewPtr2() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// VerTwoDataTwoPtr_List is a list of VerTwoDataTwoPtr.
type VerTwoDataTwoPtr_List struct{ capnp.List }

// NewVerTwoDataTwoPtr creates a new list of VerTwoDataTwoPtr.
func NewVerTwoDataTwoPtr_List(s *capnp.Segment, sz int32) (VerTwoDataTwoPtr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return VerTwoDataTwoPtr_List{l}, err
}

func (s VerTwoDataTwoPtr_List) At(i int) VerTwoDataTwoPtr { return VerTwoDataTwoPtr{s.List.Struct(i)} }

func (s VerTwoDataTwoPtr_List) Set(i int, v VerTwoDataTwoPtr) error {
	return s.List.SetStruct(i, v.Struct)
}

// VerTwoDataTwoPtr_Promise is a wrapper for a VerTwoDataTwoPtr promised by a client call.
type VerTwoDataTwoPtr_Promise struct{ *capnp.Pipeline }

func (p VerTwoDataTwoPtr_Promise) Struct() (VerTwoDataTwoPtr, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoDataTwoPtr{s}, err
}

func (p VerTwoDataTwoPtr_Promise) Ptr1() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoDataTwoPtr_Promise) Ptr2() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HoldsVerEmptyList struct{ capnp.Struct }

func NewHoldsVerEmptyList(s *capnp.Segment) (HoldsVerEmptyList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerEmptyList{st}, err
}

func NewRootHoldsVerEmptyList(s *capnp.Segment) (HoldsVerEmptyList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerEmptyList{st}, err
}

func ReadRootHoldsVerEmptyList(msg *capnp.Message) (HoldsVerEmptyList, error) {
	root, err := msg.RootPtr()
	return HoldsVerEmptyList{root.Struct()}, err
}

func (s HoldsVerEmptyList) String() string {
	str, _ := text.Marshal(0xde9ed43cfaa83093, s.Struct)
	return str
}

func (s HoldsVerEmptyList) Mylist() (VerEmpty_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerEmpty_List{List: p.List()}, err
}

func (s HoldsVerEmptyList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerEmptyList) SetMylist(v VerEmpty_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerEmpty_List, preferring placement in s's segment.
func (s HoldsVerEmptyList) NewMylist(n int32) (VerEmpty_List, error) {
	l, err := NewVerEmpty_List(s.Struct.Segment(), n)
	if err != nil {
		return VerEmpty_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerEmptyList_List is a list of HoldsVerEmptyList.
type HoldsVerEmptyList_List struct{ capnp.List }

// NewHoldsVerEmptyList creates a new list of HoldsVerEmptyList.
func NewHoldsVerEmptyList_List(s *capnp.Segment, sz int32) (HoldsVerEmptyList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerEmptyList_List{l}, err
}

func (s HoldsVerEmptyList_List) At(i int) HoldsVerEmptyList {
	return HoldsVerEmptyList{s.List.Struct(i)}
}

func (s HoldsVerEmptyList_List) Set(i int, v HoldsVerEmptyList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerEmptyList_Promise is a wrapper for a HoldsVerEmptyList promised by a client call.
type HoldsVerEmptyList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerEmptyList_Promise) Struct() (HoldsVerEmptyList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerEmptyList{s}, err
}

type HoldsVerOneDataList struct{ capnp.Struct }

func NewHoldsVerOneDataList(s *capnp.Segment) (HoldsVerOneDataList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOneDataList{st}, err
}

func NewRootHoldsVerOneDataList(s *capnp.Segment) (HoldsVerOneDataList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOneDataList{st}, err
}

func ReadRootHoldsVerOneDataList(msg *capnp.Message) (HoldsVerOneDataList, error) {
	root, err := msg.RootPtr()
	return HoldsVerOneDataList{root.Struct()}, err
}

func (s HoldsVerOneDataList) String() string {
	str, _ := text.Marshal(0xabd055422a4d7df1, s.Struct)
	return str
}

func (s HoldsVerOneDataList) Mylist() (VerOneData_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData_List{List: p.List()}, err
}

func (s HoldsVerOneDataList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerOneDataList) SetMylist(v VerOneData_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerOneData_List, preferring placement in s's segment.
func (s HoldsVerOneDataList) NewMylist(n int32) (VerOneData_List, error) {
	l, err := NewVerOneData_List(s.Struct.Segment(), n)
	if err != nil {
		return VerOneData_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerOneDataList_List is a list of HoldsVerOneDataList.
type HoldsVerOneDataList_List struct{ capnp.List }

// NewHoldsVerOneDataList creates a new list of HoldsVerOneDataList.
func NewHoldsVerOneDataList_List(s *capnp.Segment, sz int32) (HoldsVerOneDataList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerOneDataList_List{l}, err
}

func (s HoldsVerOneDataList_List) At(i int) HoldsVerOneDataList {
	return HoldsVerOneDataList{s.List.Struct(i)}
}

func (s HoldsVerOneDataList_List) Set(i int, v HoldsVerOneDataList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerOneDataList_Promise is a wrapper for a HoldsVerOneDataList promised by a client call.
type HoldsVerOneDataList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerOneDataList_Promise) Struct() (HoldsVerOneDataList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerOneDataList{s}, err
}

type HoldsVerTwoDataList struct{ capnp.Struct }

func NewHoldsVerTwoDataList(s *capnp.Segment) (HoldsVerTwoDataList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoDataList{st}, err
}

func NewRootHoldsVerTwoDataList(s *capnp.Segment) (HoldsVerTwoDataList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoDataList{st}, err
}

func ReadRootHoldsVerTwoDataList(msg *capnp.Message) (HoldsVerTwoDataList, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoDataList{root.Struct()}, err
}

func (s HoldsVerTwoDataList) String() string {
	str, _ := text.Marshal(0xcbdc765fd5dff7ba, s.Struct)
	return str
}

func (s HoldsVerTwoDataList) Mylist() (VerTwoData_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoData_List{List: p.List()}, err
}

func (s HoldsVerTwoDataList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoDataList) SetMylist(v VerTwoData_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoData_List, preferring placement in s's segment.
func (s HoldsVerTwoDataList) NewMylist(n int32) (VerTwoData_List, error) {
	l, err := NewVerTwoData_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoData_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoDataList_List is a list of HoldsVerTwoDataList.
type HoldsVerTwoDataList_List struct{ capnp.List }

// NewHoldsVerTwoDataList creates a new list of HoldsVerTwoDataList.
func NewHoldsVerTwoDataList_List(s *capnp.Segment, sz int32) (HoldsVerTwoDataList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoDataList_List{l}, err
}

func (s HoldsVerTwoDataList_List) At(i int) HoldsVerTwoDataList {
	return HoldsVerTwoDataList{s.List.Struct(i)}
}

func (s HoldsVerTwoDataList_List) Set(i int, v HoldsVerTwoDataList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoDataList_Promise is a wrapper for a HoldsVerTwoDataList promised by a client call.
type HoldsVerTwoDataList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoDataList_Promise) Struct() (HoldsVerTwoDataList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoDataList{s}, err
}

type HoldsVerOnePtrList struct{ capnp.Struct }

func NewHoldsVerOnePtrList(s *capnp.Segment) (HoldsVerOnePtrList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOnePtrList{st}, err
}

func NewRootHoldsVerOnePtrList(s *capnp.Segment) (HoldsVerOnePtrList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOnePtrList{st}, err
}

func ReadRootHoldsVerOnePtrList(msg *capnp.Message) (HoldsVerOnePtrList, error) {
	root, err := msg.RootPtr()
	return HoldsVerOnePtrList{root.Struct()}, err
}

func (s HoldsVerOnePtrList) String() string {
	str, _ := text.Marshal(0xe508a29c83a059f8, s.Struct)
	return str
}

func (s HoldsVerOnePtrList) Mylist() (VerOnePtr_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerOnePtr_List{List: p.List()}, err
}

func (s HoldsVerOnePtrList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerOnePtrList) SetMylist(v VerOnePtr_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerOnePtr_List, preferring placement in s's segment.
func (s HoldsVerOnePtrList) NewMylist(n int32) (VerOnePtr_List, error) {
	l, err := NewVerOnePtr_List(s.Struct.Segment(), n)
	if err != nil {
		return VerOnePtr_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerOnePtrList_List is a list of HoldsVerOnePtrList.
type HoldsVerOnePtrList_List struct{ capnp.List }

// NewHoldsVerOnePtrList creates a new list of HoldsVerOnePtrList.
func NewHoldsVerOnePtrList_List(s *capnp.Segment, sz int32) (HoldsVerOnePtrList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerOnePtrList_List{l}, err
}

func (s HoldsVerOnePtrList_List) At(i int) HoldsVerOnePtrList {
	return HoldsVerOnePtrList{s.List.Struct(i)}
}

func (s HoldsVerOnePtrList_List) Set(i int, v HoldsVerOnePtrList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerOnePtrList_Promise is a wrapper for a HoldsVerOnePtrList promised by a client call.
type HoldsVerOnePtrList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerOnePtrList_Promise) Struct() (HoldsVerOnePtrList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerOnePtrList{s}, err
}

type HoldsVerTwoPtrList struct{ capnp.Struct }

func NewHoldsVerTwoPtrList(s *capnp.Segment) (HoldsVerTwoPtrList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoPtrList{st}, err
}

func NewRootHoldsVerTwoPtrList(s *capnp.Segment) (HoldsVerTwoPtrList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoPtrList{st}, err
}

func ReadRootHoldsVerTwoPtrList(msg *capnp.Message) (HoldsVerTwoPtrList, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoPtrList{root.Struct()}, err
}

func (s HoldsVerTwoPtrList) String() string {
	str, _ := text.Marshal(0xcf9beaca1cc180c8, s.Struct)
	return str
}

func (s HoldsVerTwoPtrList) Mylist() (VerTwoPtr_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoPtr_List{List: p.List()}, err
}

func (s HoldsVerTwoPtrList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoPtrList) SetMylist(v VerTwoPtr_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoPtr_List, preferring placement in s's segment.
func (s HoldsVerTwoPtrList) NewMylist(n int32) (VerTwoPtr_List, error) {
	l, err := NewVerTwoPtr_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoPtr_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoPtrList_List is a list of HoldsVerTwoPtrList.
type HoldsVerTwoPtrList_List struct{ capnp.List }

// NewHoldsVerTwoPtrList creates a new list of HoldsVerTwoPtrList.
func NewHoldsVerTwoPtrList_List(s *capnp.Segment, sz int32) (HoldsVerTwoPtrList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoPtrList_List{l}, err
}

func (s HoldsVerTwoPtrList_List) At(i int) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList{s.List.Struct(i)}
}

func (s HoldsVerTwoPtrList_List) Set(i int, v HoldsVerTwoPtrList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoPtrList_Promise is a wrapper for a HoldsVerTwoPtrList promised by a client call.
type HoldsVerTwoPtrList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoPtrList_Promise) Struct() (HoldsVerTwoPtrList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoPtrList{s}, err
}

type HoldsVerTwoTwoList struct{ capnp.Struct }

func NewHoldsVerTwoTwoList(s *capnp.Segment) (HoldsVerTwoTwoList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoList{st}, err
}

func NewRootHoldsVerTwoTwoList(s *capnp.Segment) (HoldsVerTwoTwoList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoList{st}, err
}

func ReadRootHoldsVerTwoTwoList(msg *capnp.Message) (HoldsVerTwoTwoList, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoTwoList{root.Struct()}, err
}

func (s HoldsVerTwoTwoList) String() string {
	str, _ := text.Marshal(0x95befe3f14606e6b, s.Struct)
	return str
}

func (s HoldsVerTwoTwoList) Mylist() (VerTwoDataTwoPtr_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoDataTwoPtr_List{List: p.List()}, err
}

func (s HoldsVerTwoTwoList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoTwoList) SetMylist(v VerTwoDataTwoPtr_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoDataTwoPtr_List, preferring placement in s's segment.
func (s HoldsVerTwoTwoList) NewMylist(n int32) (VerTwoDataTwoPtr_List, error) {
	l, err := NewVerTwoDataTwoPtr_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoDataTwoPtr_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoTwoList_List is a list of HoldsVerTwoTwoList.
type HoldsVerTwoTwoList_List struct{ capnp.List }

// NewHoldsVerTwoTwoList creates a new list of HoldsVerTwoTwoList.
func NewHoldsVerTwoTwoList_List(s *capnp.Segment, sz int32) (HoldsVerTwoTwoList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoTwoList_List{l}, err
}

func (s HoldsVerTwoTwoList_List) At(i int) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList{s.List.Struct(i)}
}

func (s HoldsVerTwoTwoList_List) Set(i int, v HoldsVerTwoTwoList) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoTwoList_Promise is a wrapper for a HoldsVerTwoTwoList promised by a client call.
type HoldsVerTwoTwoList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoTwoList_Promise) Struct() (HoldsVerTwoTwoList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoTwoList{s}, err
}

type HoldsVerTwoTwoPlus struct{ capnp.Struct }

func NewHoldsVerTwoTwoPlus(s *capnp.Segment) (HoldsVerTwoTwoPlus, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoPlus{st}, err
}

func NewRootHoldsVerTwoTwoPlus(s *capnp.Segment) (HoldsVerTwoTwoPlus, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoPlus{st}, err
}

func ReadRootHoldsVerTwoTwoPlus(msg *capnp.Message) (HoldsVerTwoTwoPlus, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoTwoPlus{root.Struct()}, err
}

func (s HoldsVerTwoTwoPlus) String() string {
	str, _ := text.Marshal(0x87c33f2330feb3d8, s.Struct)
	return str
}

func (s HoldsVerTwoTwoPlus) Mylist() (VerTwoTwoPlus_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoTwoPlus_List{List: p.List()}, err
}

func (s HoldsVerTwoTwoPlus) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoTwoPlus) SetMylist(v VerTwoTwoPlus_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoTwoPlus_List, preferring placement in s's segment.
func (s HoldsVerTwoTwoPlus) NewMylist(n int32) (VerTwoTwoPlus_List, error) {
	l, err := NewVerTwoTwoPlus_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoTwoPlus_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoTwoPlus_List is a list of HoldsVerTwoTwoPlus.
type HoldsVerTwoTwoPlus_List struct{ capnp.List }

// NewHoldsVerTwoTwoPlus creates a new list of HoldsVerTwoTwoPlus.
func NewHoldsVerTwoTwoPlus_List(s *capnp.Segment, sz int32) (HoldsVerTwoTwoPlus_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoTwoPlus_List{l}, err
}

func (s HoldsVerTwoTwoPlus_List) At(i int) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus{s.List.Struct(i)}
}

func (s HoldsVerTwoTwoPlus_List) Set(i int, v HoldsVerTwoTwoPlus) error {
	return s.List.SetStruct(i, v.Struct)
}

// HoldsVerTwoTwoPlus_Promise is a wrapper for a HoldsVerTwoTwoPlus promised by a client call.
type HoldsVerTwoTwoPlus_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoTwoPlus_Promise) Struct() (HoldsVerTwoTwoPlus, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoTwoPlus{s}, err
}

type VerTwoTwoPlus struct{ capnp.Struct }

func NewVerTwoTwoPlus(s *capnp.Segment) (VerTwoTwoPlus, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return VerTwoTwoPlus{st}, err
}

func NewRootVerTwoTwoPlus(s *capnp.Segment) (VerTwoTwoPlus, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return VerTwoTwoPlus{st}, err
}

func ReadRootVerTwoTwoPlus(msg *capnp.Message) (VerTwoTwoPlus, error) {
	root, err := msg.RootPtr()
	return VerTwoTwoPlus{root.Struct()}, err
}

func (s VerTwoTwoPlus) String() string {
	str, _ := text.Marshal(0xce44aee2d9e25049, s.Struct)
	return str
}

func (s VerTwoTwoPlus) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoTwoPlus) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoTwoPlus) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoTwoPlus) SetDuo(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s VerTwoTwoPlus) Ptr1() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoDataTwoPtr{Struct: p.Struct()}, err
}

func (s VerTwoTwoPlus) HasPtr1() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerTwoTwoPlus) SetPtr1(v VerTwoDataTwoPtr) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewPtr1() (VerTwoDataTwoPtr, error) {
	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoTwoPlus) Ptr2() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Ptr(1)
	return VerTwoDataTwoPtr{Struct: p.Struct()}, err
}

func (s VerTwoTwoPlus) HasPtr2() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerTwoTwoPlus) SetPtr2(v VerTwoDataTwoPtr) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewPtr2() (VerTwoDataTwoPtr, error) {
	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoTwoPlus) Tre() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s VerTwoTwoPlus) SetTre(v int64) {
	s.Struct.SetUint64(16, uint64(v))
}

func (s VerTwoTwoPlus) Lst3() (capnp.Int64List, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.Int64List{List: p.List()}, err
}

func (s VerTwoTwoPlus) HasLst3() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s VerTwoTwoPlus) SetLst3(v capnp.Int64List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewLst3 sets the lst3 field to a newly
// allocated capnp.Int64List, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewLst3(n int32) (capnp.Int64List, error) {
	l, err := capnp.NewInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int64List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// VerTwoTwoPlus_List is a list of VerTwoTwoPlus.
type VerTwoTwoPlus_List struct{ capnp.List }

// NewVerTwoTwoPlus creates a new list of VerTwoTwoPlus.
func NewVerTwoTwoPlus_List(s *capnp.Segment, sz int32) (VerTwoTwoPlus_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	return VerTwoTwoPlus_List{l}, err
}

func (s VerTwoTwoPlus_List) At(i int) VerTwoTwoPlus { return VerTwoTwoPlus{s.List.Struct(i)} }

func (s VerTwoTwoPlus_List) Set(i int, v VerTwoTwoPlus) error { return s.List.SetStruct(i, v.Struct) }

// VerTwoTwoPlus_Promise is a wrapper for a VerTwoTwoPlus promised by a client call.
type VerTwoTwoPlus_Promise struct{ *capnp.Pipeline }

func (p VerTwoTwoPlus_Promise) Struct() (VerTwoTwoPlus, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoTwoPlus{s}, err
}

func (p VerTwoTwoPlus_Promise) Ptr1() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoTwoPlus_Promise) Ptr2() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HoldsText struct{ capnp.Struct }

func NewHoldsText(s *capnp.Segment) (HoldsText, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HoldsText{st}, err
}

func NewRootHoldsText(s *capnp.Segment) (HoldsText, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HoldsText{st}, err
}

func ReadRootHoldsText(msg *capnp.Message) (HoldsText, error) {
	root, err := msg.RootPtr()
	return HoldsText{root.Struct()}, err
}

func (s HoldsText) String() string {
	str, _ := text.Marshal(0xe5817f849ff906dc, s.Struct)
	return str
}

func (s HoldsText) Txt() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HoldsText) HasTxt() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsText) TxtBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s HoldsText) SetTxt(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HoldsText) Lst() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s HoldsText) HasLst() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HoldsText) SetLst(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewLst sets the lst field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s HoldsText) NewLst(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s HoldsText) Lstlst() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.PointerList{List: p.List()}, err
}

func (s HoldsText) HasLstlst() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HoldsText) SetLstlst(v capnp.PointerList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewLstlst sets the lstlst field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s HoldsText) NewLstlst(n int32) (capnp.PointerList, error) {
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// HoldsText_List is a list of HoldsText.
type HoldsText_List struct{ capnp.List }

// NewHoldsText creates a new list of HoldsText.
func NewHoldsText_List(s *capnp.Segment, sz int32) (HoldsText_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return HoldsText_List{l}, err
}

func (s HoldsText_List) At(i int) HoldsText { return HoldsText{s.List.Struct(i)} }

func (s HoldsText_List) Set(i int, v HoldsText) error { return s.List.SetStruct(i, v.Struct) }

// HoldsText_Promise is a wrapper for a HoldsText promised by a client call.
type HoldsText_Promise struct{ *capnp.Pipeline }

func (p HoldsText_Promise) Struct() (HoldsText, error) {
	s, err := p.Pipeline.Struct()
	return HoldsText{s}, err
}

type WrapEmpty struct{ capnp.Struct }

func NewWrapEmpty(s *capnp.Segment) (WrapEmpty, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WrapEmpty{st}, err
}

func NewRootWrapEmpty(s *capnp.Segment) (WrapEmpty, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WrapEmpty{st}, err
}

func ReadRootWrapEmpty(msg *capnp.Message) (WrapEmpty, error) {
	root, err := msg.RootPtr()
	return WrapEmpty{root.Struct()}, err
}

func (s WrapEmpty) String() string {
	str, _ := text.Marshal(0x9ab599979b02ac59, s.Struct)
	return str
}

func (s WrapEmpty) MightNotBeReallyEmpty() (VerEmpty, error) {
	p, err := s.Struct.Ptr(0)
	return VerEmpty{Struct: p.Struct()}, err
}

func (s WrapEmpty) HasMightNotBeReallyEmpty() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WrapEmpty) SetMightNotBeReallyEmpty(v VerEmpty) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerEmpty struct, preferring placement in s's segment.
func (s WrapEmpty) NewMightNotBeReallyEmpty() (VerEmpty, error) {
	ss, err := NewVerEmpty(s.Struct.Segment())
	if err != nil {
		return VerEmpty{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// WrapEmpty_List is a list of WrapEmpty.
type WrapEmpty_List struct{ capnp.List }

// NewWrapEmpty creates a new list of WrapEmpty.
func NewWrapEmpty_List(s *capnp.Segment, sz int32) (WrapEmpty_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WrapEmpty_List{l}, err
}

func (s WrapEmpty_List) At(i int) WrapEmpty { return WrapEmpty{s.List.Struct(i)} }

func (s WrapEmpty_List) Set(i int, v WrapEmpty) error { return s.List.SetStruct(i, v.Struct) }

// WrapEmpty_Promise is a wrapper for a WrapEmpty promised by a client call.
type WrapEmpty_Promise struct{ *capnp.Pipeline }

func (p WrapEmpty_Promise) Struct() (WrapEmpty, error) {
	s, err := p.Pipeline.Struct()
	return WrapEmpty{s}, err
}

func (p WrapEmpty_Promise) MightNotBeReallyEmpty() VerEmpty_Promise {
	return VerEmpty_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Wrap2x2 struct{ capnp.Struct }

func NewWrap2x2(s *capnp.Segment) (Wrap2x2, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2{st}, err
}

func NewRootWrap2x2(s *capnp.Segment) (Wrap2x2, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2{st}, err
}

func ReadRootWrap2x2(msg *capnp.Message) (Wrap2x2, error) {
	root, err := msg.RootPtr()
	return Wrap2x2{root.Struct()}, err
}

func (s Wrap2x2) String() string {
	str, _ := text.Marshal(0xe1a2d1d51107bead, s.Struct)
	return str
}

func (s Wrap2x2) MightNotBeReallyEmpty() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoDataTwoPtr{Struct: p.Struct()}, err
}

func (s Wrap2x2) HasMightNotBeReallyEmpty() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Wrap2x2) SetMightNotBeReallyEmpty(v VerTwoDataTwoPtr) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s Wrap2x2) NewMightNotBeReallyEmpty() (VerTwoDataTwoPtr, error) {
	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Wrap2x2_List is a list of Wrap2x2.
type Wrap2x2_List struct{ capnp.List }

// NewWrap2x2 creates a new list of Wrap2x2.
func NewWrap2x2_List(s *capnp.Segment, sz int32) (Wrap2x2_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Wrap2x2_List{l}, err
}

func (s Wrap2x2_List) At(i int) Wrap2x2 { return Wrap2x2{s.List.Struct(i)} }

func (s Wrap2x2_List) Set(i int, v Wrap2x2) error { return s.List.SetStruct(i, v.Struct) }

// Wrap2x2_Promise is a wrapper for a Wrap2x2 promised by a client call.
type Wrap2x2_Promise struct{ *capnp.Pipeline }

func (p Wrap2x2_Promise) Struct() (Wrap2x2, error) {
	s, err := p.Pipeline.Struct()
	return Wrap2x2{s}, err
}

func (p Wrap2x2_Promise) MightNotBeReallyEmpty() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Wrap2x2plus struct{ capnp.Struct }

func NewWrap2x2plus(s *capnp.Segment) (Wrap2x2plus, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2plus{st}, err
}

func NewRootWrap2x2plus(s *capnp.Segment) (Wrap2x2plus, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2plus{st}, err
}

func ReadRootWrap2x2plus(msg *capnp.Message) (Wrap2x2plus, error) {
	root, err := msg.RootPtr()
	return Wrap2x2plus{root.Struct()}, err
}

func (s Wrap2x2plus) String() string {
	str, _ := text.Marshal(0xe684eb3aef1a6859, s.Struct)
	return str
}

func (s Wrap2x2plus) MightNotBeReallyEmpty() (VerTwoTwoPlus, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoTwoPlus{Struct: p.Struct()}, err
}

func (s Wrap2x2plus) HasMightNotBeReallyEmpty() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Wrap2x2plus) SetMightNotBeReallyEmpty(v VerTwoTwoPlus) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerTwoTwoPlus struct, preferring placement in s's segment.
func (s Wrap2x2plus) NewMightNotBeReallyEmpty() (VerTwoTwoPlus, error) {
	ss, err := NewVerTwoTwoPlus(s.Struct.Segment())
	if err != nil {
		return VerTwoTwoPlus{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Wrap2x2plus_List is a list of Wrap2x2plus.
type Wrap2x2plus_List struct{ capnp.List }

// NewWrap2x2plus creates a new list of Wrap2x2plus.
func NewWrap2x2plus_List(s *capnp.Segment, sz int32) (Wrap2x2plus_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Wrap2x2plus_List{l}, err
}

func (s Wrap2x2plus_List) At(i int) Wrap2x2plus { return Wrap2x2plus{s.List.Struct(i)} }

func (s Wrap2x2plus_List) Set(i int, v Wrap2x2plus) error { return s.List.SetStruct(i, v.Struct) }

// Wrap2x2plus_Promise is a wrapper for a Wrap2x2plus promised by a client call.
type Wrap2x2plus_Promise struct{ *capnp.Pipeline }

func (p Wrap2x2plus_Promise) Struct() (Wrap2x2plus, error) {
	s, err := p.Pipeline.Struct()
	return Wrap2x2plus{s}, err
}

func (p Wrap2x2plus_Promise) MightNotBeReallyEmpty() VerTwoTwoPlus_Promise {
	return VerTwoTwoPlus_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type VoidUnion struct{ capnp.Struct }
type VoidUnion_Which uint16

const (
	VoidUnion_Which_a VoidUnion_Which = 0
	VoidUnion_Which_b VoidUnion_Which = 1
)

func (w VoidUnion_Which) String() string {
	const s = "ab"
	switch w {
	case VoidUnion_Which_a:
		return s[0:1]
	case VoidUnion_Which_b:
		return s[1:2]

	}
	return "VoidUnion_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewVoidUnion(s *capnp.Segment) (VoidUnion, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VoidUnion{st}, err
}

func NewRootVoidUnion(s *capnp.Segment) (VoidUnion, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VoidUnion{st}, err
}

func ReadRootVoidUnion(msg *capnp.Message) (VoidUnion, error) {
	root, err := msg.RootPtr()
	return VoidUnion{root.Struct()}, err
}

func (s VoidUnion) String() string {
	str, _ := text.Marshal(0x8821cdb23640783a, s.Struct)
	return str
}

func (s VoidUnion) Which() VoidUnion_Which {
	return VoidUnion_Which(s.Struct.Uint16(0))
}
func (s VoidUnion) SetA() {
	s.Struct.SetUint16(0, 0)

}

func (s VoidUnion) SetB() {
	s.Struct.SetUint16(0, 1)

}

// VoidUnion_List is a list of VoidUnion.
type VoidUnion_List struct{ capnp.List }

// NewVoidUnion creates a new list of VoidUnion.
func NewVoidUnion_List(s *capnp.Segment, sz int32) (VoidUnion_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return VoidUnion_List{l}, err
}

func (s VoidUnion_List) At(i int) VoidUnion { return VoidUnion{s.List.Struct(i)} }

func (s VoidUnion_List) Set(i int, v VoidUnion) error { return s.List.SetStruct(i, v.Struct) }

// VoidUnion_Promise is a wrapper for a VoidUnion promised by a client call.
type VoidUnion_Promise struct{ *capnp.Pipeline }

func (p VoidUnion_Promise) Struct() (VoidUnion, error) {
	s, err := p.Pipeline.Struct()
	return VoidUnion{s}, err
}

type Nester1Capn struct{ capnp.Struct }

func NewNester1Capn(s *capnp.Segment) (Nester1Capn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Nester1Capn{st}, err
}

func NewRootNester1Capn(s *capnp.Segment) (Nester1Capn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Nester1Capn{st}, err
}

func ReadRootNester1Capn(msg *capnp.Message) (Nester1Capn, error) {
	root, err := msg.RootPtr()
	return Nester1Capn{root.Struct()}, err
}

func (s Nester1Capn) String() string {
	str, _ := text.Marshal(0xf14fad09425d081c, s.Struct)
	return str
}

func (s Nester1Capn) Strs() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s Nester1Capn) HasStrs() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Nester1Capn) SetStrs(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewStrs sets the strs field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Nester1Capn) NewStrs(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Nester1Capn_List is a list of Nester1Capn.
type Nester1Capn_List struct{ capnp.List }

// NewNester1Capn creates a new list of Nester1Capn.
func NewNester1Capn_List(s *capnp.Segment, sz int32) (Nester1Capn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Nester1Capn_List{l}, err
}

func (s Nester1Capn_List) At(i int) Nester1Capn { return Nester1Capn{s.List.Struct(i)} }

func (s Nester1Capn_List) Set(i int, v Nester1Capn) error { return s.List.SetStruct(i, v.Struct) }

// Nester1Capn_Promise is a wrapper for a Nester1Capn promised by a client call.
type Nester1Capn_Promise struct{ *capnp.Pipeline }

func (p Nester1Capn_Promise) Struct() (Nester1Capn, error) {
	s, err := p.Pipeline.Struct()
	return Nester1Capn{s}, err
}

type RWTestCapn struct{ capnp.Struct }

func NewRWTestCapn(s *capnp.Segment) (RWTestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return RWTestCapn{st}, err
}

func NewRootRWTestCapn(s *capnp.Segment) (RWTestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return RWTestCapn{st}, err
}

func ReadRootRWTestCapn(msg *capnp.Message) (RWTestCapn, error) {
	root, err := msg.RootPtr()
	return RWTestCapn{root.Struct()}, err
}

func (s RWTestCapn) String() string {
	str, _ := text.Marshal(0xf7ff4414476c186a, s.Struct)
	return str
}

func (s RWTestCapn) NestMatrix() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.PointerList{List: p.List()}, err
}

func (s RWTestCapn) HasNestMatrix() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RWTestCapn) SetNestMatrix(v capnp.PointerList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewNestMatrix sets the nestMatrix field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s RWTestCapn) NewNestMatrix(n int32) (capnp.PointerList, error) {
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// RWTestCapn_List is a list of RWTestCapn.
type RWTestCapn_List struct{ capnp.List }

// NewRWTestCapn creates a new list of RWTestCapn.
func NewRWTestCapn_List(s *capnp.Segment, sz int32) (RWTestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return RWTestCapn_List{l}, err
}

func (s RWTestCapn_List) At(i int) RWTestCapn { return RWTestCapn{s.List.Struct(i)} }

func (s RWTestCapn_List) Set(i int, v RWTestCapn) error { return s.List.SetStruct(i, v.Struct) }

// RWTestCapn_Promise is a wrapper for a RWTestCapn promised by a client call.
type RWTestCapn_Promise struct{ *capnp.Pipeline }

func (p RWTestCapn_Promise) Struct() (RWTestCapn, error) {
	s, err := p.Pipeline.Struct()
	return RWTestCapn{s}, err
}

type ListStructCapn struct{ capnp.Struct }

func NewListStructCapn(s *capnp.Segment) (ListStructCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ListStructCapn{st}, err
}

func NewRootListStructCapn(s *capnp.Segment) (ListStructCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ListStructCapn{st}, err
}

func ReadRootListStructCapn(msg *capnp.Message) (ListStructCapn, error) {
	root, err := msg.RootPtr()
	return ListStructCapn{root.Struct()}, err
}

func (s ListStructCapn) String() string {
	str, _ := text.Marshal(0xb1ac056ed7647011, s.Struct)
	return str
}

func (s ListStructCapn) Vec() (Nester1Capn_List, error) {
	p, err := s.Struct.Ptr(0)
	return Nester1Capn_List{List: p.List()}, err
}

func (s ListStructCapn) HasVec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ListStructCapn) SetVec(v Nester1Capn_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewVec sets the vec field to a newly
// allocated Nester1Capn_List, preferring placement in s's segment.
func (s ListStructCapn) NewVec(n int32) (Nester1Capn_List, error) {
	l, err := NewNester1Capn_List(s.Struct.Segment(), n)
	if err != nil {
		return Nester1Capn_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// ListStructCapn_List is a list of ListStructCapn.
type ListStructCapn_List struct{ capnp.List }

// NewListStructCapn creates a new list of ListStructCapn.
func NewListStructCapn_List(s *capnp.Segment, sz int32) (ListStructCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ListStructCapn_List{l}, err
}

func (s ListStructCapn_List) At(i int) ListStructCapn { return ListStructCapn{s.List.Struct(i)} }

func (s ListStructCapn_List) Set(i int, v ListStructCapn) error { return s.List.SetStruct(i, v.Struct) }

// ListStructCapn_Promise is a wrapper for a ListStructCapn promised by a client call.
type ListStructCapn_Promise struct{ *capnp.Pipeline }

func (p ListStructCapn_Promise) Struct() (ListStructCapn, error) {
	s, err := p.Pipeline.Struct()
	return ListStructCapn{s}, err
}

type Echo struct{ Client capnp.Client }

func (c Echo) Echo(ctx context.Context, params func(Echo_echo_Params) error, opts ...capnp.CallOption) Echo_echo_Results_Promise {
	if c.Client == nil {
		return Echo_echo_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x8e5322c1e9282534,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:Echo",
			MethodName:    "echo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Echo_echo_Params{Struct: s}) }
	}
	return Echo_echo_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Echo_Server interface {
	Echo(Echo_echo) error
}

func Echo_ServerToClient(s Echo_Server) Echo {
	c, _ := s.(server.Closer)
	return Echo{Client: server.New(Echo_Methods(nil, s), c)}
}

func Echo_Methods(methods []server.Method, s Echo_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x8e5322c1e9282534,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:Echo",
			MethodName:    "echo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Echo_echo{c, opts, Echo_echo_Params{Struct: p}, Echo_echo_Results{Struct: r}}
			return s.Echo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Echo_echo holds the arguments for a server call to Echo.echo.
type Echo_echo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Echo_echo_Params
	Results Echo_echo_Results
}

type Echo_echo_Params struct{ capnp.Struct }

func NewEcho_echo_Params(s *capnp.Segment) (Echo_echo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Params{st}, err
}

func NewRootEcho_echo_Params(s *capnp.Segment) (Echo_echo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Params{st}, err
}

func ReadRootEcho_echo_Params(msg *capnp.Message) (Echo_echo_Params, error) {
	root, err := msg.RootPtr()
	return Echo_echo_Params{root.Struct()}, err
}

func (s Echo_echo_Params) String() string {
	str, _ := text.Marshal(0x8a165fb4d71bf3a2, s.Struct)
	return str
}

func (s Echo_echo_Params) In() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Echo_echo_Params) HasIn() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Echo_echo_Params) InBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Echo_echo_Params) SetIn(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Echo_echo_Params_List is a list of Echo_echo_Params.
type Echo_echo_Params_List struct{ capnp.List }

// NewEcho_echo_Params creates a new list of Echo_echo_Params.
func NewEcho_echo_Params_List(s *capnp.Segment, sz int32) (Echo_echo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Echo_echo_Params_List{l}, err
}

func (s Echo_echo_Params_List) At(i int) Echo_echo_Params { return Echo_echo_Params{s.List.Struct(i)} }

func (s Echo_echo_Params_List) Set(i int, v Echo_echo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Echo_echo_Params_Promise is a wrapper for a Echo_echo_Params promised by a client call.
type Echo_echo_Params_Promise struct{ *capnp.Pipeline }

func (p Echo_echo_Params_Promise) Struct() (Echo_echo_Params, error) {
	s, err := p.Pipeline.Struct()
	return Echo_echo_Params{s}, err
}

type Echo_echo_Results struct{ capnp.Struct }

func NewEcho_echo_Results(s *capnp.Segment) (Echo_echo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Results{st}, err
}

func NewRootEcho_echo_Results(s *capnp.Segment) (Echo_echo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Results{st}, err
}

func ReadRootEcho_echo_Results(msg *capnp.Message) (Echo_echo_Results, error) {
	root, err := msg.RootPtr()
	return Echo_echo_Results{root.Struct()}, err
}

func (s Echo_echo_Results) String() string {
	str, _ := text.Marshal(0x9b37d729b9dd7b9d, s.Struct)
	return str
}

func (s Echo_echo_Results) Out() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Echo_echo_Results) HasOut() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Echo_echo_Results) OutBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Echo_echo_Results) SetOut(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Echo_echo_Results_List is a list of Echo_echo_Results.
type Echo_echo_Results_List struct{ capnp.List }

// NewEcho_echo_Results creates a new list of Echo_echo_Results.
func NewEcho_echo_Results_List(s *capnp.Segment, sz int32) (Echo_echo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Echo_echo_Results_List{l}, err
}

func (s Echo_echo_Results_List) At(i int) Echo_echo_Results {
	return Echo_echo_Results{s.List.Struct(i)}
}

func (s Echo_echo_Results_List) Set(i int, v Echo_echo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Echo_echo_Results_Promise is a wrapper for a Echo_echo_Results promised by a client call.
type Echo_echo_Results_Promise struct{ *capnp.Pipeline }

func (p Echo_echo_Results_Promise) Struct() (Echo_echo_Results, error) {
	s, err := p.Pipeline.Struct()
	return Echo_echo_Results{s}, err
}

type Hoth struct{ capnp.Struct }

func NewHoth(s *capnp.Segment) (Hoth, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hoth{st}, err
}

func NewRootHoth(s *capnp.Segment) (Hoth, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hoth{st}, err
}

func ReadRootHoth(msg *capnp.Message) (Hoth, error) {
	root, err := msg.RootPtr()
	return Hoth{root.Struct()}, err
}

func (s Hoth) String() string {
	str, _ := text.Marshal(0xad87da456fb0ebb9, s.Struct)
	return str
}

func (s Hoth) Base() (EchoBase, error) {
	p, err := s.Struct.Ptr(0)
	return EchoBase{Struct: p.Struct()}, err
}

func (s Hoth) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Hoth) SetBase(v EchoBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated EchoBase struct, preferring placement in s's segment.
func (s Hoth) NewBase() (EchoBase, error) {
	ss, err := NewEchoBase(s.Struct.Segment())
	if err != nil {
		return EchoBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Hoth_List is a list of Hoth.
type Hoth_List struct{ capnp.List }

// NewHoth creates a new list of Hoth.
func NewHoth_List(s *capnp.Segment, sz int32) (Hoth_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Hoth_List{l}, err
}

func (s Hoth_List) At(i int) Hoth { return Hoth{s.List.Struct(i)} }

func (s Hoth_List) Set(i int, v Hoth) error { return s.List.SetStruct(i, v.Struct) }

// Hoth_Promise is a wrapper for a Hoth promised by a client call.
type Hoth_Promise struct{ *capnp.Pipeline }

func (p Hoth_Promise) Struct() (Hoth, error) {
	s, err := p.Pipeline.Struct()
	return Hoth{s}, err
}

func (p Hoth_Promise) Base() EchoBase_Promise {
	return EchoBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EchoBase struct{ capnp.Struct }

func NewEchoBase(s *capnp.Segment) (EchoBase, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EchoBase{st}, err
}

func NewRootEchoBase(s *capnp.Segment) (EchoBase, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EchoBase{st}, err
}

func ReadRootEchoBase(msg *capnp.Message) (EchoBase, error) {
	root, err := msg.RootPtr()
	return EchoBase{root.Struct()}, err
}

func (s EchoBase) String() string {
	str, _ := text.Marshal(0xa8bf13fef2674866, s.Struct)
	return str
}

func (s EchoBase) Echo() Echo {
	p, _ := s.Struct.Ptr(0)
	return Echo{Client: p.Interface().Client()}
}

func (s EchoBase) HasEcho() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EchoBase) SetEcho(v Echo) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// EchoBase_List is a list of EchoBase.
type EchoBase_List struct{ capnp.List }

// NewEchoBase creates a new list of EchoBase.
func NewEchoBase_List(s *capnp.Segment, sz int32) (EchoBase_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EchoBase_List{l}, err
}

func (s EchoBase_List) At(i int) EchoBase { return EchoBase{s.List.Struct(i)} }

func (s EchoBase_List) Set(i int, v EchoBase) error { return s.List.SetStruct(i, v.Struct) }

// EchoBase_Promise is a wrapper for a EchoBase promised by a client call.
type EchoBase_Promise struct{ *capnp.Pipeline }

func (p EchoBase_Promise) Struct() (EchoBase, error) {
	s, err := p.Pipeline.Struct()
	return EchoBase{s}, err
}

func (p EchoBase_Promise) Echo() Echo {
	return Echo{Client: p.Pipeline.GetPipeline(0).Client()}
}

type StackingRoot struct{ capnp.Struct }

func NewStackingRoot(s *capnp.Segment) (StackingRoot, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return StackingRoot{st}, err
}

func NewRootStackingRoot(s *capnp.Segment) (StackingRoot, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return StackingRoot{st}, err
}

func ReadRootStackingRoot(msg *capnp.Message) (StackingRoot, error) {
	root, err := msg.RootPtr()
	return StackingRoot{root.Struct()}, err
}

func (s StackingRoot) String() string {
	str, _ := text.Marshal(0x8fae7b41c61fc890, s.Struct)
	return str
}

func (s StackingRoot) A() (StackingA, error) {
	p, err := s.Struct.Ptr(1)
	return StackingA{Struct: p.Struct()}, err
}

func (s StackingRoot) HasA() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s StackingRoot) SetA(v StackingA) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewA sets the a field to a newly
// allocated StackingA struct, preferring placement in s's segment.
func (s StackingRoot) NewA() (StackingA, error) {
	ss, err := NewStackingA(s.Struct.Segment())
	if err != nil {
		return StackingA{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s StackingRoot) AWithDefault() (StackingA, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return StackingA{}, err
	}
	ss, err := p.StructDefault(x_832bcc6686a26d56[64:96])
	return StackingA{Struct: ss}, err
}

func (s StackingRoot) HasAWithDefault() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s StackingRoot) SetAWithDefault(v StackingA) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAWithDefault sets the aWithDefault field to a newly
// allocated StackingA struct, preferring placement in s's segment.
func (s StackingRoot) NewAWithDefault() (StackingA, error) {
	ss, err := NewStackingA(s.Struct.Segment())
	if err != nil {
		return StackingA{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// StackingRoot_List is a list of StackingRoot.
type StackingRoot_List struct{ capnp.List }

// NewStackingRoot creates a new list of StackingRoot.
func NewStackingRoot_List(s *capnp.Segment, sz int32) (StackingRoot_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return StackingRoot_List{l}, err
}

func (s StackingRoot_List) At(i int) StackingRoot { return StackingRoot{s.List.Struct(i)} }

func (s StackingRoot_List) Set(i int, v StackingRoot) error { return s.List.SetStruct(i, v.Struct) }

// StackingRoot_Promise is a wrapper for a StackingRoot promised by a client call.
type StackingRoot_Promise struct{ *capnp.Pipeline }

func (p StackingRoot_Promise) Struct() (StackingRoot, error) {
	s, err := p.Pipeline.Struct()
	return StackingRoot{s}, err
}

func (p StackingRoot_Promise) A() StackingA_Promise {
	return StackingA_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p StackingRoot_Promise) AWithDefault() StackingA_Promise {
	return StackingA_Promise{Pipeline: p.Pipeline.GetPipelineDefault(0, x_832bcc6686a26d56[96:128])}
}

type StackingA struct{ capnp.Struct }

func NewStackingA(s *capnp.Segment) (StackingA, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StackingA{st}, err
}

func NewRootStackingA(s *capnp.Segment) (StackingA, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StackingA{st}, err
}

func ReadRootStackingA(msg *capnp.Message) (StackingA, error) {
	root, err := msg.RootPtr()
	return StackingA{root.Struct()}, err
}

func (s StackingA) String() string {
	str, _ := text.Marshal(0x9d3032ff86043b75, s.Struct)
	return str
}

func (s StackingA) Num() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s StackingA) SetNum(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

func (s StackingA) B() (StackingB, error) {
	p, err := s.Struct.Ptr(0)
	return StackingB{Struct: p.Struct()}, err
}

func (s StackingA) HasB() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s StackingA) SetB(v StackingB) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewB sets the b field to a newly
// allocated StackingB struct, preferring placement in s's segment.
func (s StackingA) NewB() (StackingB, error) {
	ss, err := NewStackingB(s.Struct.Segment())
	if err != nil {
		return StackingB{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// StackingA_List is a list of StackingA.
type StackingA_List struct{ capnp.List }

// NewStackingA creates a new list of StackingA.
func NewStackingA_List(s *capnp.Segment, sz int32) (StackingA_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return StackingA_List{l}, err
}

func (s StackingA_List) At(i int) StackingA { return StackingA{s.List.Struct(i)} }

func (s StackingA_List) Set(i int, v StackingA) error { return s.List.SetStruct(i, v.Struct) }

// StackingA_Promise is a wrapper for a StackingA promised by a client call.
type StackingA_Promise struct{ *capnp.Pipeline }

func (p StackingA_Promise) Struct() (StackingA, error) {
	s, err := p.Pipeline.Struct()
	return StackingA{s}, err
}

func (p StackingA_Promise) B() StackingB_Promise {
	return StackingB_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type StackingB struct{ capnp.Struct }

func NewStackingB(s *capnp.Segment) (StackingB, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return StackingB{st}, err
}

func NewRootStackingB(s *capnp.Segment) (StackingB, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return StackingB{st}, err
}

func ReadRootStackingB(msg *capnp.Message) (StackingB, error) {
	root, err := msg.RootPtr()
	return StackingB{root.Struct()}, err
}

func (s StackingB) String() string {
	str, _ := text.Marshal(0x85257b30d6edf8c5, s.Struct)
	return str
}

func (s StackingB) Num() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s StackingB) SetNum(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// StackingB_List is a list of StackingB.
type StackingB_List struct{ capnp.List }

// NewStackingB creates a new list of StackingB.
func NewStackingB_List(s *capnp.Segment, sz int32) (StackingB_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return StackingB_List{l}, err
}

func (s StackingB_List) At(i int) StackingB { return StackingB{s.List.Struct(i)} }

func (s StackingB_List) Set(i int, v StackingB) error { return s.List.SetStruct(i, v.Struct) }

// StackingB_Promise is a wrapper for a StackingB promised by a client call.
type StackingB_Promise struct{ *capnp.Pipeline }

func (p StackingB_Promise) Struct() (StackingB, error) {
	s, err := p.Pipeline.Struct()
	return StackingB{s}, err
}

type CallSequence struct{ Client capnp.Client }

func (c CallSequence) GetNumber(ctx context.Context, params func(CallSequence_getNumber_Params) error, opts ...capnp.CallOption) CallSequence_getNumber_Results_Promise {
	if c.Client == nil {
		return CallSequence_getNumber_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xabaedf5f7817c820,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:CallSequence",
			MethodName:    "getNumber",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(CallSequence_getNumber_Params{Struct: s}) }
	}
	return CallSequence_getNumber_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type CallSequence_Server interface {
	GetNumber(CallSequence_getNumber) error
}

func CallSequence_ServerToClient(s CallSequence_Server) CallSequence {
	c, _ := s.(server.Closer)
	return CallSequence{Client: server.New(CallSequence_Methods(nil, s), c)}
}

func CallSequence_Methods(methods []server.Method, s CallSequence_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xabaedf5f7817c820,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:CallSequence",
			MethodName:    "getNumber",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := CallSequence_getNumber{c, opts, CallSequence_getNumber_Params{Struct: p}, CallSequence_getNumber_Results{Struct: r}}
			return s.GetNumber(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// CallSequence_getNumber holds the arguments for a server call to CallSequence.getNumber.
type CallSequence_getNumber struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  CallSequence_getNumber_Params
	Results CallSequence_getNumber_Results
}

type CallSequence_getNumber_Params struct{ capnp.Struct }

func NewCallSequence_getNumber_Params(s *capnp.Segment) (CallSequence_getNumber_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CallSequence_getNumber_Params{st}, err
}

func NewRootCallSequence_getNumber_Params(s *capnp.Segment) (CallSequence_getNumber_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CallSequence_getNumber_Params{st}, err
}

func ReadRootCallSequence_getNumber_Params(msg *capnp.Message) (CallSequence_getNumber_Params, error) {
	root, err := msg.RootPtr()
	return CallSequence_getNumber_Params{root.Struct()}, err
}

func (s CallSequence_getNumber_Params) String() string {
	str, _ := text.Marshal(0xf58782f48a121998, s.Struct)
	return str
}

// CallSequence_getNumber_Params_List is a list of CallSequence_getNumber_Params.
type CallSequence_getNumber_Params_List struct{ capnp.List }

// NewCallSequence_getNumber_Params creates a new list of CallSequence_getNumber_Params.
func NewCallSequence_getNumber_Params_List(s *capnp.Segment, sz int32) (CallSequence_getNumber_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return CallSequence_getNumber_Params_List{l}, err
}

func (s CallSequence_getNumber_Params_List) At(i int) CallSequence_getNumber_Params {
	return CallSequence_getNumber_Params{s.List.Struct(i)}
}

func (s CallSequence_getNumber_Params_List) Set(i int, v CallSequence_getNumber_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// CallSequence_getNumber_Params_Promise is a wrapper for a CallSequence_getNumber_Params promised by a client call.
type CallSequence_getNumber_Params_Promise struct{ *capnp.Pipeline }

func (p CallSequence_getNumber_Params_Promise) Struct() (CallSequence_getNumber_Params, error) {
	s, err := p.Pipeline.Struct()
	return CallSequence_getNumber_Params{s}, err
}

type CallSequence_getNumber_Results struct{ capnp.Struct }

func NewCallSequence_getNumber_Results(s *capnp.Segment) (CallSequence_getNumber_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return CallSequence_getNumber_Results{st}, err
}

func NewRootCallSequence_getNumber_Results(s *capnp.Segment) (CallSequence_getNumber_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return CallSequence_getNumber_Results{st}, err
}

func ReadRootCallSequence_getNumber_Results(msg *capnp.Message) (CallSequence_getNumber_Results, error) {
	root, err := msg.RootPtr()
	return CallSequence_getNumber_Results{root.Struct()}, err
}

func (s CallSequence_getNumber_Results) String() string {
	str, _ := text.Marshal(0xa465f9502fd11e97, s.Struct)
	return str
}

func (s CallSequence_getNumber_Results) N() uint32 {
	return s.Struct.Uint32(0)
}

func (s CallSequence_getNumber_Results) SetN(v uint32) {
	s.Struct.SetUint32(0, v)
}

// CallSequence_getNumber_Results_List is a list of CallSequence_getNumber_Results.
type CallSequence_getNumber_Results_List struct{ capnp.List }

// NewCallSequence_getNumber_Results creates a new list of CallSequence_getNumber_Results.
func NewCallSequence_getNumber_Results_List(s *capnp.Segment, sz int32) (CallSequence_getNumber_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return CallSequence_getNumber_Results_List{l}, err
}

func (s CallSequence_getNumber_Results_List) At(i int) CallSequence_getNumber_Results {
	return CallSequence_getNumber_Results{s.List.Struct(i)}
}

func (s CallSequence_getNumber_Results_List) Set(i int, v CallSequence_getNumber_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// CallSequence_getNumber_Results_Promise is a wrapper for a CallSequence_getNumber_Results promised by a client call.
type CallSequence_getNumber_Results_Promise struct{ *capnp.Pipeline }

func (p CallSequence_getNumber_Results_Promise) Struct() (CallSequence_getNumber_Results, error) {
	s, err := p.Pipeline.Struct()
	return CallSequence_getNumber_Results{s}, err
}

type Defaults struct{ capnp.Struct }

func NewDefaults(s *capnp.Segment) (Defaults, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Defaults{st}, err
}

func NewRootDefaults(s *capnp.Segment) (Defaults, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Defaults{st}, err
}

func ReadRootDefaults(msg *capnp.Message) (Defaults, error) {
	root, err := msg.RootPtr()
	return Defaults{root.Struct()}, err
}

func (s Defaults) String() string {
	str, _ := text.Marshal(0x97e38948c61f878d, s.Struct)
	return str
}

func (s Defaults) Text() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextDefault("foo"), err
}

func (s Defaults) HasText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Defaults) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.DataDefault([]byte("foo" + "\x00"))
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Defaults) SetText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Defaults) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.DataDefault([]byte{0x62, 0x61, 0x72})), err
}

func (s Defaults) HasData() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Defaults) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s Defaults) Float() float32 {
	return math.Float32frombits(s.Struct.Uint32(0) ^ 0x4048f5c3)
}

func (s Defaults) SetFloat(v float32) {
	s.Struct.SetUint32(0, math.Float32bits(v)^0x4048f5c3)
}

func (s Defaults) Int() int32 {
	return int32(s.Struct.Uint32(4) ^ 4294967173)
}

func (s Defaults) SetInt(v int32) {
	s.Struct.SetUint32(4, uint32(v)^4294967173)
}

func (s Defaults) Uint() uint32 {
	return s.Struct.Uint32(8) ^ 42
}

func (s Defaults) SetUint(v uint32) {
	s.Struct.SetUint32(8, v^42)
}

// Defaults_List is a list of Defaults.
type Defaults_List struct{ capnp.List }

// NewDefaults creates a new list of Defaults.
func NewDefaults_List(s *capnp.Segment, sz int32) (Defaults_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return Defaults_List{l}, err
}

func (s Defaults_List) At(i int) Defaults { return Defaults{s.List.Struct(i)} }

func (s Defaults_List) Set(i int, v Defaults) error { return s.List.SetStruct(i, v.Struct) }

// Defaults_Promise is a wrapper for a Defaults promised by a client call.
type Defaults_Promise struct{ *capnp.Pipeline }

func (p Defaults_Promise) Struct() (Defaults, error) {
	s, err := p.Pipeline.Struct()
	return Defaults{s}, err
}

type BenchmarkA struct{ capnp.Struct }

func NewBenchmarkA(s *capnp.Segment) (BenchmarkA, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return BenchmarkA{st}, err
}

func NewRootBenchmarkA(s *capnp.Segment) (BenchmarkA, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return BenchmarkA{st}, err
}

func ReadRootBenchmarkA(msg *capnp.Message) (BenchmarkA, error) {
	root, err := msg.RootPtr()
	return BenchmarkA{root.Struct()}, err
}

func (s BenchmarkA) String() string {
	str, _ := text.Marshal(0xde2a1a960863c11c, s.Struct)
	return str
}

func (s BenchmarkA) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s BenchmarkA) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BenchmarkA) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s BenchmarkA) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s BenchmarkA) BirthDay() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s BenchmarkA) SetBirthDay(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s BenchmarkA) Phone() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s BenchmarkA) HasPhone() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s BenchmarkA) PhoneBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s BenchmarkA) SetPhone(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s BenchmarkA) Siblings() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s BenchmarkA) SetSiblings(v int32) {
	s.Struct.SetUint32(8, uint32(v))
}

func (s BenchmarkA) Spouse() bool {
	return s.Struct.Bit(96)
}

func (s BenchmarkA) SetSpouse(v bool) {
	s.Struct.SetBit(96, v)
}

func (s BenchmarkA) Money() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s BenchmarkA) SetMoney(v float64) {
	s.Struct.SetUint64(16, math.Float64bits(v))
}

// BenchmarkA_List is a list of BenchmarkA.
type BenchmarkA_List struct{ capnp.List }

// NewBenchmarkA creates a new list of BenchmarkA.
func NewBenchmarkA_List(s *capnp.Segment, sz int32) (BenchmarkA_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	return BenchmarkA_List{l}, err
}

func (s BenchmarkA_List) At(i int) BenchmarkA { return BenchmarkA{s.List.Struct(i)} }

func (s BenchmarkA_List) Set(i int, v BenchmarkA) error { return s.List.SetStruct(i, v.Struct) }

// BenchmarkA_Promise is a wrapper for a BenchmarkA promised by a client call.
type BenchmarkA_Promise struct{ *capnp.Pipeline }

func (p BenchmarkA_Promise) Struct() (BenchmarkA, error) {
	s, err := p.Pipeline.Struct()
	return BenchmarkA{s}, err
}

const schema_832bcc6686a26d56 = "x\xda\xacY\x0dpTU\x96~\xb7\x7f\xf2:\xe9t" +
	":/\xaf\x03\xf9\xc3\x98H\x10\xb2\x15'?\x0c\xa23" +
	"VHL\x06\xb4\xd0I\xd2 \xe2\x9a1\xaf\x93\x17\xd2" +
	"\xd8\x7ft\xbf\x86\x04\xc7\xc2\xd9\x81\x01,\xd8\x91\x12F" +
	"P\xd9e\xb2X+#\xb0\xe0\xc8\x96a\x85\x11\x07F" +
	"\x88\xb8\xa3\x16(\xb0\x8a\x8a\x8b\x8e(;\xc6\xd1\x15\xfc" +
	"\xeb=\xe7\xbe\xdf~\xfd\x1a\x94\xda*B^\xcew\xee" +
	"\xb9\xe7\x9e\xbf{\xee\xbd\xf5\x7fq\xcf\xb058g\xf2" +
	"\x0c\xd3\xf9\x8e3'u\xe8\xc2\xf9\xd7\xeb\xef\xabY\xc1" +
	"tz\x09I\xdd\x11\x1e\xfeU\xff\xd1\xbf\xfb%\xe3`" +
	"\x19\x86_\x9d;\xcao\xc8\xc5\xafu\xb9\xcd\x0cIu" +
	"\xbfT\xf8\xb3\xdc\xe7f\xad4\xf1:m\xc8\xb2;w" +
	"/?B\x99\xf7\xe4\xfe\x1b0\x9f\xf8\xfd\xb7\xf5\xd74" +
	"\xffq%\xc3y\x8d\xbc\x048\x9a\x84\xbc\"\xc2/\xca" +
	"C\xe6p\x1eJ\xbeqp\xc6\xb4\xa7_\xaeZe\x92" +
	"\xdc\xc6\xda\x80es\xde(\xbf\x8d2?\x91\xb7\x04\x98" +
	"\x87\xffV\xfe\xc63\xf7\x8c{\x90\xe1|\x84Q$\x12" +
	"\xb7\x8d0\x84w\xbaQ\xda\xd4\x9a\xc9\x1f\x1e\xa8\xf6\xff" +
	"#Lm\xd7\x85\x01\\\xe3\x1e\xe6\xeb\xdc(i\x8a{" +
	"&\xdf\x89_\xa9\x87\x0eW\xfe\xa9\xe5\xbe\x9d\xbf6\xe9" +
	"I\xd7t\x83\xfb\x0c\xdfN\xf9[\xdc8s*~\xf4" +
	"\\\xe7\xa6#\x0f\xa7\xf3R[ms\xbf\xc0\xef\x06V" +
	"{*\xff\xd8o\x0e\x16=U\xbf\x1e\x98\x1ci\xb3\xaf" +
	"u\x8f\xf2\x9bP\x9a\x7f\xbd\xdbN\xfc\xbbPe&\xb5" +
	"\xb6\xca5\xf5\xc2\xfd\x7fXoa'~\x07\x8c\x18\xa1" +
	"\xf3\xef\xa1\x0b\xbb7\xd2\xe3k\xfev\xff\x06+\x9b\x1e" +
	"w\x83M\xcfR\xe6w)\xf3\xda\x95\x95\x7f\x9a\xb5\xfa" +
	"\xbdG\xd0\xa66\xf3\xca\xb8\xfc\x17\xf8\xd2|\x1cW\x9c" +
	"_\x09\x96K\xcd\xdfn{\xec\x91M{\x1e\xb5R\xa3" +
	"\xce3\xca\xdf\xe0\xc1\xaf\x1fzP\xf2\xe6\xfb\xde\x1a\x99" +
	"\xf2\xc6\xf5\x8f\x19\x1d\xd0\xed\xc9C\x07\x88\x94\xe1\xf0\xdc" +
	"3\xce\xbd\xd7\xfe\xfa\xb1\x0c\x13\xac\x00I\xebP\x92\x7f" +
	"\x8d\x07L\xb0\xd1CM\x90\xfc\x91\xe3W\xa9\xc6\xfa\xcd" +
	"\xe6\xb8\xa2\x93\xaf\x86!\x1b\xe8\xe4\xeb<\xe8\x83G\xae" +
	"z\xf5\x07\x1d\x17\xc5\xadLg\x05\x81\xc1h\xfc\xa6\xf3" +
	"\x9e8N~\x91N\xde?k\xc1\xa7\xdf\xf2\x7fx\xd2" +
	"j%\xa5\x05/\xf0U\x05\xf85\xa1\x00y\xaf><" +
	"~\xf0\x9e\xb7w>\x95\x11)7\x15\x9c\xe1o\xa1\x8c" +
	"\xed\x053\xf9E\xf8\x95\x1a\xbb\xff\xb6\xda\xd6\xb9\xaf<" +
	"ee\xfd\xf9\x05e\x84\x0f\xd2\x01\"\x95<\xf2\xd1\xae" +
	"h\xfb\xa9\x95;\xac\xb4\xd8P0\xcco\xa6\xbc\x9b(" +
	"/\x17\xeb{#\xe2\xdc\xbe\xdb\x8aw\xa4\xe0S\xfe\x10" +
	"\xe5=@y\x97M\xfb\xd9\xf2\xee\xe9\x9f\xecF[\x19" +
	"Tv\xda\x91e\xac\xe05\xfe\x1bdn\xbaX0\x0f" +
	"\xbd\x9a\x18\x9d\x9a\xfa\xf8\xccU\xffn\x15\x03Mb\xa1" +
	"\x0d\xb2\xb0\x90fa\xe1\x07\x98X\xeeO\x87/\x8a\xa7" +
	"^\xb4\xd2\xa3\x85\xfb\x1d\x7f\x0bG\x0d\xc2\xa1\x1eu\xb7" +
	"\xdd\xb0\xef\xbd\xa7\xfe\xfe\x88U\xda\x88\xdc(\xbf\x88\xf2" +
	"\x869t\xd9\xde/\xde>~\xcf\xe27_\xb22\xdc" +
	"\x11\x0e\x0cw\x922\x1f\xa7\x827\x1e\xdc\xe6~\x9fk" +
	";j\xa5\x84\xb3h/\xef)\xc2\xaf\xdc\"\xe4}g" +
	"\xd2\xf4\x9e\xf7\x9e\xfe\xbd%\xefME\xc3|;\xe5m" +
	"\xa1\xbc\xb7t\x9c9yfg\xdb\x7fZ\x1aN(:" +
	"\xc7\x87\x91\xb9)XD\x0dw\xf8\x81\x03\x15\xa3\xe7\x1e" +
	"\xfb\xb3\x95\xca\x87x\xc8\xb4\xe3<\x8e{\x95G\xd1\xfb" +
	"O\xad\x1a\xdb\xfa\xd5\xb4\xd7\xad\xd4 \xbeG\xf9\\\x1f" +
	"U\xde\x87\xbc\xeb\xae\xff\xd7\x9e\xc8\x9f\x9f;\x81j8" +
	"\xcc\x86\xab\xf1\x8d\xf2\x0d\xc8\xdcT\xe7\xa3j\xac}\xe9" +
	"\xe4\x92U=kNZI^W<\xcco*\xa6\xf1" +
	"T\x8c\x92\xf9\xc5_\x05\xfb[^}\xcb\xca#{\x80" +
	"w\x1f\xe5\x1d)F\x8fT\x1c\xe8u\xfd\xa6\xac\xf6\xb4" +
	"\xd9\x18\xb2\x16\xe3^\xe3\x1b\xc6Q-\xc6Q-6\xcf" +
	"\x99\xb7\xe3?vv\x9c\xb6\xda#~1\xfew\xfc\xea" +
	"\xf1\xf8\xb5b<\x96\xfd\x87\xeb\x9f\xfc\xf2\xc7\xc7\xfe\xe9" +
	"\xb4\x95\xe1&\x94\xe4\x11\xbe\xae\x84\xd6\xdf\x12Ty\xc7" +
	"~\x96;\xfe\xea\xf0\xbbV\xcb\x9b_\xb2\x97\x17(o" +
	"7\xe5\x15\xfcME\x87\xce\x1d\xb1\xe4]Q\xf2(\xbf" +
	"\x96\xf2\xae\xa6\xbc\x17\xe6o\xf9\xe5\xe3\xc3\xae\xb3VJ" +
	"l+\x01\xef\x8dP\xe6=\x94y\xf7\x8bsO\xef," +
	"\xbc\xfd\xaciu\xed\x84u\x00\xcf\xd9\x92\x17\xf8\xf3\x94" +
	"\xfb\xc3\x12\xcc\x91\xaa\x1f\\(\xffzE7\x8a\xb6\xa5" +
	"U\x8c}\xa5{\xf9C\xa54QK\xd1\x10o\xe6\\" +
	"\xfc\xe7\xe5\xcb~a\xd6\x81\x86[q\xd9(_UF" +
	"\xcbP\x19\xf2\xce\x1f(\xfb\xeb\x8d\x1f-\x7f\xdfjm" +
	"g\xcbN\xf1c\x94\xf7|\x19\xdd\x03\x8e\xef\xdf\xba\xa3" +
	"l\xd1\x07\x19\xb5\xb5\xb8\x1c\x84\x96cm\xad(\x87\xda" +
	":\xb9\x9c\xd6\xd6\x93[^\xff\xd1#\x1fN:g\xaa" +
	"\x00\xb0\xba\xc98}\xf9\x83|\x0d\x8ei\xaa*\x7f\xd1" +
	"\x85\x81\xe1\xean\xcd\xdd\xf1\xd31+M\xd6\xd6\x9c\xe2" +
	"7\xd5\xd0\x80\xab\xa1\x99ZZ\xf4\xe0g\xff\xb0\xf2s" +
	"\x86\xabP\x0b\xf1\xa1\x9a\x85\x84q\xa4\xbe\xe9\x9fy\xa4" +
	"\xfdM\xe7\x17\xa6Ii\xc0\xec\xaey\x8d\xdfG\xa5\x8c" +
	"\xd4`(.,\x09\xcd\xf4\xb5\xa5\xbe\xb0\x9a\xb1x\xd2" +
	"k|\xd5$j\xa7I8\xe3\xe9\xd9\xfb\x1f\x9e,\xfd" +
	"\xcb\xd7V\x81(\x00o\x98\xf2\x06\x81\xb7.%\x04\xe3" +
	"\xbdq\xa1_\xb2]\xd7+\xc4\"\xb1\x1b\xfd\x92\xd0{" +
	"o0\xb2\xa0\x95a:\x08\xe9t\xd8\xc1\xbb\x0e\xd0\x9b" +
	"\xf3TC+\xe4\xb2\x93N\x9f\x8d\xb0\x91d\x988\x18" +
	"\x1b\xfc\x10M\x02Q$\xdc\xdc\x1cMF$1\x8e\xc3" +
	"\xf3\xb5\xe1\xed\xb50|\x06\x0c\x9fm#\x84\xf8pO" +
	"\xe2ni\x04Z\x1b\xd0:l\x84\xb3\x01\x11\xda\x18\xee" +
	"\xb6[\x818\x1b\x88\x036\xe2M\x04\x97\x8a\xc4\x09\x13" +
	"9\x19R\xb9$\x1a\xefK\x90|\xf8+\x1f\xa6\xc5\xbf" +
	"B\xc1\x84\xc4\x80S\x0b\x18\xd2a'\x14*0hd" +
	"W4\x9a\x15\x0d\xf5%\xee\x10\xe3s\x96D\xe1_G" +
	"(I\x12\xa6\xb5\xdd\xa8\xacm\xa2\x8d4\x87\x87P\xac" +
	"*\xb3P\xaf\x88\x0cI\x93\xaeZ\xec\x8eh\xb0on" +
	"$\x18\x8d\xc8\x16s\xd9\x1d\xf9\xa9\x14\x15;\xa5\x08\xc4" +
	"N\x04\xb1\xf56\xe2!\xdf\xa6\xe4U\xd7!u2P" +
	"\xa7\x82%\x04&\x87\x04\x98\x9c\x0c\x95\xdb{\x07\xa2\xd7" +
	"\x89\xf0\xdf\xc4\x0e!.\x84\x13\x8cQ\xdb2\xdd\x13\xf6" +
	"`D\xb3\x88\xd9\x11\xed,\x8c\x97\x17\xea\x84\x10W\xfb" +
	"A\xa2\xf6%\x1cW\xcb\xd88'\xeb\xc5yf\x10\xe0" +
	"\xcc\x1a\x0d]l4*)\xab#\x04\xfe\x11n\xcaB" +
	"}\x19\x85\x8aCo\xc0\xa5M\x05\xda\x0c\x1b\xc8\x9a\x17" +
	"\x94\x06\xda\xc4~\xc6+$C\x12\x18R\xeb_\xc0\x90" +
	"\x85\xd4X\x04\x82\x82\x08\x19\x90\x85\x8d\xc5x{8&" +
	"\x0d1VJ\xf6F#\x09i6\x0d\x05\x94\xa99N" +
	"\xab\xc7\xd4q\x0cG\xc6s.\xe2}\x9bu\x95\xe3\x7f" +
	"\x15V\x93\xfc4\"vH\xf1K\x86~L\x8a\x83l" +
	"-\xc5L\x0a[\x87\x1c(G\xa4\xef\x11rZ?\x92" +
	"%\xe4\xc0\xaah\xd3\x04\xd5\xd3'\xcb\x04\x97\xdc\x8f9" +
	"6\x082\x97C>\xd1$\x03\xe2\x0a$>\x00\xc45" +
	"\x10n6\xc81\xa0\xad\xc5\xc4[\x05\xb4\xf5\xc0h\x07" +
	"F\x98\x95[\x87\xab\\\x03\xc4\x8d@t\x00'\xc8\xe4" +
	"6\xe0\xe8\x87\x80\xf88d\xa3$\x0eJJ\xb4\x815" +
	"\xab\xd9\xfeh\xd4\xdb'H\x02\xf1\x00\xcd\x83\xb426" +
	" \xc4+\xfbCQA\"y\x8cm,\xef\x8f\x9f\xcf" +
	"\x9a\xc1\x106\x18\x91\xb0^\x8c9V\xa4R)\x86x" +
	"\x93HpA\xf4\xb9j-\x967/.\xc4dw\x9b" +
	"\x1d\xf14hS\x08\xdaT@|\x85\x83\x0b\x06\xa4\xdb" +
	"\xa3\x12i\x15\xbbD!\x14\x1a\xaa\xa4c\xc0|\xda\xf9" +
	"#\x8bs\xf4\xe4\xea\x12\x13\xd4\x8eL6oG\x93R" +
	"F~\xa5\xc5];\x94B9\xee\xbc\xfa\xa6\xc7\x10\xa7" +
	"\x970\xd9\xd3\xa9E+\x15\xca\x9cS\xaa\xf5J\xa1V" +
	"Gc\x9d0\x16\\\x12\x80%j\xe7Q\xd3\x12\x1dj" +
	"\x11\x06\x83\xf8\xc5EI1\xd2+^\xb7@\x94nO" +
	"\x86\x03b\x1c\xd6[I\x17l\\n\x91\xbe\\\x12A" +
	"\xa7\x10\x97\x85\xeah\xb3V!!\x9a=R\xab\x8f\xa6" +
	"u\x84p\xfa\xb9\x12t\xe3,Di\xba\xb1\xa0\x9c^" +
	"\x9f\xd4}\x92\xa8G\x17\x8e\xeb\x82\x08\xc9eS\xaa\xfe" +
	"\x0c\x89\xa7\x97)s\xb6A\xfa\xb6A<\xce\x0e\xda\x13" +
	"\xdf'\xdd\x8c\xc9\\`QHg\xb1Qi\xe0\x12\xeb" +
	"\x0e\x80]@\x8cv\xa2\xcaR\xc4\xb0D\xf9\xa5x\xb2" +
	"\xb7R\xba\x19\x08Yj\x0c(\xc8.\x16{u\xed\xb4" +
	"^#K1\xe8\x12\x17\xc4\xc5D\"\x18%Td\x89" +
	"&r\x13\xea\xb8\x1eDn\xd1\x83j3\xee\x1d\x1b\x81" +
	"\xb6\xd5\xb0\xe5\xfe\x16\x19\x1f\x07\xe2\xb3X\x0e \xf3\xed" +
	"@\xdc\x83\x16\xdb\x05\xc4\xa3X\x0e\x80\x13\xc4rGP" +
	"\xcb\x83@|\x05\x88N\xe0\x04\xc7q/#\xf10\x10" +
	"\x8f\xe9\xb6\xd0\xfaw\xd9\x16\xf6@=qCd\xb9!" +
	"\xf7\x03\"T\x0ceuny\xc7n\x8e\x85\x84\x88\x98" +
	"\xd0\xd7\xac5\x9b\xf2\x9a\xd9\xa1pR\x1d\xcf\x0e%\xfa" +
	"\xd4\xef\x8c@\x90+.\xc6\x00n\xf4P\xc9\xd1\"\x85" +
	"\x9aE\x04\xd4\xf4n\xb9\xb7\xe0T\x93\x88H\xec\x01b" +
	"H)\x90@\x0b\xa2E\xfa\x80\x16S\x0b$\x10\xc3H" +
	"\x1c\x00\xa2\x84.\x12B@\x84\x0d\x184\xeaKF\xd5" +
	"&\xc5\x0b\xdbCC\xe6\xfe\x80\xe4\xc6Kl\x1bj\xa4" +
	"\xddE\xeb\xe9%B\xcdPn-wIe\xd9\xe6\xf2" +
	"R\xab\x97\x17m\xddu\xb5z}\xb9B\xb5-v\xbb" +
	"+\xc9?\xad\x0f\xce\x92\x7fw5'\xc4\xf8b\xb9\xa3" +
	"4\x88\x0c\x80\xc8|\x109\x196\x83%BP\x82\xc2" +
	"\xba\x90a\xa3\x01C\x1ci\x87\xbd,\x92[\xd9\xeb\x9b" +
	"\xae\xbf|f\x9b\xa29\x8b\xe1\xe1\x9f\x17\x9a\xcb\x84)" +
	"\x0b\xab\xf5,\xd4l\xbf\xb9ZOC5\xe6\xb4,|" +
	"\xd2\x10sO q\x0b\x10\xb7\xab\x9b2\x10\xb7\xe1\xe8" +
	"\xad@\xdce\xc8\xc2\x1d\xc8\xf9$\x10\x0f^>:\x8d" +
	"\xfd\x85\xc1\xcd&2+\xc5\xb5\xe6\xdb\x1bJHM\xaa" +
	"]\x9d\x97\xef\xb3!\x08\xbfo\xd3\xa3]\x95d\xf3\x95" +
	"]X`\x12\xd7\xaa\xbbjY\xaf|\xea\x00A\xda\xfd" +
	"k\x16ou`\xb1\xa1{\x19\x93\xbd`\xea\xaejT" +
	"\\\xf5\x8c\xee\xaa\xdd\xb8\x8e\xedj\xc1\x9ca*\x98\xcf" +
	"\x19\\5\x82\xa7\x99ge\xafpN\xbb\xec\xaa\x03H" +
	"|^.\xad\xde\x88\x10\x16\xd5\x16\xa3r \x1a\xd6\xcb" +
	"`ZCA\xcbd\\\xc0 W}\xd2\xdc+D~" +
	"\x12\x1a\x02-A3X'\xacN\xe8\x0db\xd7\xc4\xa8" +
	",\xa9\xb00\xe8\x8f\x89b\x1f\xd2\xccES5l\x0b" +
	"\xdb\xd4X\x7f\xe5I\xa0e)\xbb0\x1a\xc8\xde\xd6d" +
	"\xd6\x9d6\x08\xd4\xdep\x9f\xbax\xaf\x10_\x90\xc8v" +
	"\x94S}\xd7\x0a\x0d\xcd@X\x88\xdfKZ\xbe\xc3f" +
	"w\xab!\xa3\xd4\xcd\xee\x89FCF\xa9\x9b\xdd\xb6[" +
	"\x95\xe4y\x06}\xd7#\xfb.\xcd\xcbj\x9a\xedi\xd4" +
	"\xbd\x9c\xe6\xbbT \x18\x87\xf3\x8e`4\x7fel " +
	"\x1a\xd19\x12\xc1@\x08\xfc\x97@\x0e\xa5\x99kN\xc4" +
	"\xa2I0\xaf\xe2\xc3\xca0\xf0\x0fe\xf5\x14\xdd\x1e\xc4" +
	"\xec\x07k\x0e\x034\xf3dmWN\xd6\xe8\x8dY@" +
	"\x9c\x03\xaa\x0f\x89B\\-\x118\xab4@r\xe0\xaf" +
	"\x1c,\x18\xc2\x90\xfa\x9d5\xc5i\xab\xad\x9d\xb9\xbek" +
	"\x8a\x1b\x1bs\xab\x14\x9f\x07\x11\x1ek\x1cl\xbc\xa2\xae" +
	"\xdfT\xbd2\x84\xff\xc4\xde0\xed\xca\xa3\xdc\xa2\xd5\xbc" +
	"\x82\x1a\xa7\xbdbd\xe9\xe5Z\x94\xbf\xd5\xbeE\xbdI" +
	"\x10j\xf5\xc6E\xbfI\x10k\xf5\xce\xc5c\xfb&\x95" +
	"\xd9\xbbx\xec_\xa7\x94\xe6\xa5Zo^\xbc\x8b\xa3\xc1" +
	">&\xc7\x1b\x80\xcd\x0f\x94\xd2\xae\x87\x95\xdd@\x80r" +
	"\x80\xba\xaa\x97\xaa\xcan\xd0\xdf0\x0d\xa8\xda\xfdb\x16" +
	"+\xb74\x07\xe3\xb1h\x9c\x1a\xa5B\x0e\xbbZz\x10" +
	"h\x07\x05\x88\x8d\xbb\x09\x7f\xd9\xb9\x1f\xe2/\x07W\x87" +
	"\xbf\x9c\\\x0d\xfe\xca\xe1&\x00\xa77\x02\x19\xc0.\xec" +
	"\xbf\x97\x0d\x09\x83l\xa2?\xca\x86\x92\x8b\xd9\xbe\xfe%" +
	"p\xfaLH\x19\x06\xa3\xee\x98\x03\xe7R9\x0e\x0di" +
	"QmL\x0b\xf5\xc2\xa9ZI\x8b\x1eL\x0b\xa5\x98w" +
	"\xa3\xc3\xee\x94O\xcd\xac\xa4\x9dp\x09\x1b\xd2=\xa7\x94" +
	"\xa3f \x19\xa8\x97\xa9U\xf3\xe4`\x8e\x85\xec\xc9\xc4" +
	"\x15E\xb4\xf1\xe6\xa90\xdb\x01\x14\x1a.y#3\xdd" +
	"x\x142L\xa1|\xdb\x91YFdcm\xd1\xe2\x8b" +
	"\xeft\x80\xe9\xfd\xb3\x1dv\xe2\xbf\xd3a\x081~\xae" +
	"\x03\x0e\x0c\xfe\x0e\x04\xee\x06`\x02D\x99\\G\xf9\xf9" +
	"\x0e0\xa6\x7f\x0e\"=\x88@\xa4\xc9\xc5\x94\xef\xa6\xc8" +
	"\x9d\x88\xf4!\xe2\xf8*%\x1f\x1fx\x81\"w#2" +
	"\x80\x88\xf3\xcb\x94\\Vy\x91\"=\x88\x84\x10\xc9\xb9" +
	"\x98r\xf8\xa0\x021|\x90\"}\x88\xc4\x10a/\xa4" +
	"\\>z\xed\x19\xa6\xba\x0d \"!\xe2\xfa\x02\xe7q" +
	"\x01\xb2\x88\x8e\x09!2\x88H\xee\xff\xe2<\xb9\x80$" +
	")\x12C\xe4\xe7\x88\xe4}\x8e\xf3\xe4\x012D\x11\x09" +
	"\x91\x07\x10q\x7f\x86\xf3\xb8\x01\xb9\x9f\xce3\x88\xc8r" +
	"D\xf2\xff\x96\x82-?\x1fo\xf5\xa9\xd9~\x8e\xc8*" +
	"4\x9b\xe7S0\x9b\x07o\xda)\xf0\x00\x02k\x10(" +
	"\x18\x03\xa0\x00/\xde)\xb0\x1c\x81\x87\x10\xf0~\x02\x80" +
	"\x17o\x8d\x1d\x10\x86 \x05\x80\xad\x08\x14\xfe\x15\x00\xf0" +
	"!\xff[\x0a<\x8e\xc0\xb3\x08p\xff\x03\x00\x87\xd7\xf2" +
	"\x14\xd8\x85\xc0Q\x04\x8a\xce\x03\x00\x87y\xfe\x08\x05\x0e" +
	"\"\xf0\x0e\x02\xfc\xc7\x00\xf0\x00\xbcE\x81\x13\x08|\x86" +
	"\x80\xef#\x00|\xf8 \xe6\x80\xfd\xc2\xff1\x02.'" +
	"\x00\xc5\xe7\x00(\xc6\xc7\x18'\x8c\xe8r\x02\xbd\x02\xe9" +
	"\xe3>\x04\xfa8| D\xba\xdf\x87@=\x02\xe3\xff" +
	"\x02\xc0x|\x03\xa5\xc0d\x04\xda\x10(\xf9\x00\x80\x12" +
	"|]r\xe2\x14?F\xe0N\x04J\xdf\x07\xa0\x14#" +
	"\xcc\x89\x16\xe9@ \x84@\xd9Y\x00\xca\xd0\xf1\xceV" +
	"t<\x02\x0f!P\xfe\xdf\x00\x94\xa3\xa9\xa8\xa8U\x08" +
	"\xacG\xa0\xe2=\x00*\xf0\xa9\x87\x02k\x10\xd8\x88\xc0" +
	"\x843\x00L\xc0\x8bxg\x00\x9f\x93\x11\xd8\x85\xc0U" +
	"\xef\x02p\x15\xbe\x1c;a\xe3\xf7oG\xe0Y\x04*" +
	"\xdf\x01\xa0\x12\x8d\xeb\x84$\xf1?\x83\xc0\xf3\x08\\\xfd" +
	"6\x00W\x03\xb0\xcf\xd9\x05\xc0s\x08\x1c\x06`B\xd5" +
	"i\x0c\xa0*@\x0eQ}\x9fG\xe4(\x0e\xa9~\x0b" +
	"\x86T\xa3?\xe8\x0a\x0f\"\xf0\x0a\x02\xd7\xbc\x09\xc05" +
	"\x00\xbcL\x81\xc3\x08\x1cC`\xe2\x7f\x010\x11\xdf\xcb" +
	"\x9c\x18\x8bG\x118\x81@\xcd)\x00j\xf0\xed\x8f\xea" +
	"{\x0c\x81O\x10\x98t\x12\x80I\xf8\xe8A\x81\x8f\x11" +
	"p\xe5\x00p\xed\x09\x00\xaeE\x0f\xe6\x80V]9\xe8" +
	"\xc1\x1c\xad\xf2\xdb\x97.\x85j\xa1\xbdy\xa8\x05~\xda" +
	"T\xed\x90\xdd\xdf\xd4\x88\x17x\xf0C\xd8 \xd0\x95\x86" +
	"\x86\x0d\x02]i]\xd8 l\x08J\x07a\x0fN\x87" +
	"\xa2`\x83\x1f\xc2&\x81=\x17\xbes\xf1\x1b\xd8\x95\x1b" +
	"%6\x09\xec,|\xb3\xc0\x9e\x9c\xaev\x18\xde@4" +
	"\x1aR\xdb\x1f\xe3\x0d# \xa1h@=\xeb6\x83r" +
	"\x86[\x11\xf5\xde\x00\xd44P\xf3\x14j0\x8d\xd7\xa9" +
	"R\xd3x\x1d*\xb5a\x9a\x81j\x97\xa9\x95\xc1\xe9\x06" +
	"\xa2MaM\xa6\x89\xcdU\xa9ib]*5M," +
	"\xab\x88M\x1a\xc5\xe6\xc8D\xef\xd2\xb4\xdb\x1e\xa3S\x80" +
	"\xb8\x0cQ\x03C6\xbe\xca\xa5\xd8\x19fl\x012\x1d" +
	"o\xd8\xb5Gk\xd3^\xc2\xa4_6\x99.^t6" +
	" \xa4\xa3($\xae\\=1\xf6h\x04`\xed\xd1]" +
	"\x81\xe9\xad\x0e\xf4U\x0c\xb1\xe8\xac\x96\x09r\xa3`:" +
	"\xe6x\xd1\xf1\xff\x0f}\x09\xb5\x07,\xcd\xf0<d~" +
	"\x11\x90\x99\x043\x93\xd1R\xe8\x01\x0cP\x83\x8dH\x96" +
	"\xdd\xfevhO\xc4x\xc3\xcd\x82=\xe3j\xafV\xef" +
	"\x08\xbd\x09)\x9e\xf5\x90s\x99\xcb\xdb\x0e\xc1\x8bOA" +
	"Yn\x1f\xa0\x0f \xc2w9~U\x1b\xae\x95\xb3\xdc" +
	"\x13d\xde1\xce\x9b\x03\xcb\xbbY\x88\x11\xf3\xda\xeeR" +
	".a\xa6C'\x13\x01\x9e\xdb\x04)\xce\xd8\x83\x83\x19" +
	"\x01{\xb9kL\xed\xfa\x96\x08\x97x}1(\xfc\x7f" +
	"\x01\x00\x00\xff\xff;\x92\x03\xf8"

func init() {
	schemas.Register(schema_832bcc6686a26d56,
		0x85257b30d6edf8c5,
		0x8748bc095e10cb5d,
		0x87c33f2330feb3d8,
		0x8821cdb23640783a,
		0x8a165fb4d71bf3a2,
		0x8e5322c1e9282534,
		0x8fae7b41c61fc890,
		0x93c99951eacc72ff,
		0x9430ab12c496d40c,
		0x94bf7df83408218d,
		0x95befe3f14606e6b,
		0x97e38948c61f878d,
		0x9ab599979b02ac59,
		0x9b37d729b9dd7b9d,
		0x9b8f27ba05e255c8,
		0x9d3032ff86043b75,
		0xa465f9502fd11e97,
		0xa8bf13fef2674866,
		0xabaedf5f7817c820,
		0xabd055422a4d7df1,
		0xad87da456fb0ebb9,
		0xb1ac056ed7647011,
		0xb1f0385d845e367f,
		0xb61ee2ecff34ca73,
		0xc7da65f9a2f20ba2,
		0xc95babe3bd394d2d,
		0xcbdc765fd5dff7ba,
		0xcc4411e60ba9c498,
		0xccb3b2e3603826e0,
		0xce44aee2d9e25049,
		0xcf9beaca1cc180c8,
		0xd636fba4f188dabe,
		0xd8bccf6e60a73791,
		0xd98c608877d9cb8d,
		0xddd1416669fb7613,
		0xde2a1a960863c11c,
		0xde50aebbad57549d,
		0xde9ed43cfaa83093,
		0xe1a2d1d51107bead,
		0xe1c9eac512335361,
		0xe508a29c83a059f8,
		0xe54e10aede55c7b1,
		0xe55d85fc1bf82f21,
		0xe5817f849ff906dc,
		0xe684eb3aef1a6859,
		0xe7711aada4bed56b,
		0xea26e9973bd6a0d9,
		0xf14fad09425d081c,
		0xf58782f48a121998,
		0xf705dc45c94766fd,
		0xf7ff4414476c186a,
		0xfca3742893be4cde)
}

var x_832bcc6686a26d56 = []byte{
	0, 0, 0, 0, 2, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0,
	223, 7, 8, 27, 0, 0, 0, 0,
	0, 0, 0, 0, 4, 0, 0, 0,
	1, 0, 0, 0, 23, 0, 0, 0,
	8, 0, 0, 0, 1, 0, 0, 0,
	223, 7, 8, 27, 0, 0, 0, 0,
	223, 7, 8, 28, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 0,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 0,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}
