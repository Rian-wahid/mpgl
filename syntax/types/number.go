package types

type Int struct {
	value int32
}

func (i *Int) Value() any {
	return i.value
}

func (i *Int) TypeName() string {
	return "int"
}

func (i *Int) Assign(v Type) {
	if v.TypeName() != "int" {
		panic("invalid type")
	}
	iv := v.Value().(int32)
	i.value = iv
}

func (i *Int) IsComparable() bool {
	return true
}

type Byte struct {
	value byte
}

func (i *Byte) Value() any {
	return i.value
}

func (i *Byte) TypeName() string {
	return "byte"
}

func (i *Byte) Assign(v Type) {
	if v.TypeName() != "byte" {
		panic("invalid type")
	}
	iv := v.Value().(byte)
	i.value = iv
}

func (i *Byte) IsComparable() bool {
	return true
}

type Int8 struct {
	value int8
}

func (i *Int8) Value() any {
	return i.value
}

func (i *Int8) TypeName() string {
	return "int8"
}

func (i *Int8) IsComparable() bool {
	return true
}

func (i *Int8) Assign(v Type) {
	if v.TypeName() != "int8" {
		panic("invalid type")
	}
	iv := v.Value().(int8)
	i.value = iv
}

type Int16 struct {
	value int16
}

func (i *Int16) Value() any {
	return i.value
}

func (i *Int16) TypeName() string {
	return "int16"
}

func (i *Int16) IsComparable() bool {
	return true
}

func (i *Int16) Assign(v Type) {
	if v.TypeName() != "int16" {
		panic("invalid type")
	}
	iv := v.Value().(int16)
	i.value = iv
}

type Int32 struct {
	value int32
}

func (i *Int32) Value() any {
	return i.value
}

func (i *Int32) TypeName() string {
	return "int32"
}

func (i *Int32) IsComparable() bool {
	return true
}

func (i *Int32) Assign(v Type) {
	if v.TypeName() != "int32" {
		panic("invalid type")
	}
	iv := v.Value().(int32)
	i.value = iv
}

type Int64 struct {
	value int64
}

func (i *Int64) Value() any {
	return i.value
}

func (i *Int64) TypeName() string {
	return "int64"
}

func (i *Int64) IsComparable() bool {
	return true
}

func (i *Int64) Assign(v Type) {
	if v.TypeName() != "int64" {
		panic("invalid type")
	}
	iv := v.Value().(int64)
	i.value = iv
}

type Uint struct {
	value uint32
}

func (i *Uint) Value() any {
	return i.value
}

func (i *Uint) TypeName() string {
	return "uint"
}

func (i *Uint) Assign(v Type) {
	if v.TypeName() != "uint" {
		panic("invalid type")
	}
	iv := v.Value().(uint32)
	i.value = iv
}

func (i *Uint) IsComparable() bool {
	return true
}

type Uint16 struct {
	value uint16
}

func (i *Uint16) Value() any {
	return i.value
}

func (i *Uint16) TypeName() string {
	return "uint16"
}

func (i *Uint16) Assign(v Type) {
	if v.TypeName() != "uint16" {
		panic("invalid type")
	}
	iv := v.Value().(uint16)
	i.value = iv
}

func (i *Uint16) IsComparable() bool {
	return true
}

type Uint32 struct {
	value uint32
}

func (i *Uint32) Value() any {
	return i.value
}

func (i *Uint32) TypeName() string {
	return "uint32"
}

func (i *Uint32) Assign(v Type) {
	if v.TypeName() != "uint32" {
		panic("invalid type")
	}
	iv := v.Value().(uint32)
	i.value = iv
}

func (i *Uint32) IsComparable() bool {
	return true
}

type Uint64 struct {
	value uint64
}

func (i *Uint64) Value() any {
	return i.value
}

func (i *Uint64) TypeName() string {
	return "uint64"
}

func (i *Uint64) Assign(v Type) {
	if v.TypeName() != "uint64" {
		panic("invalid type")
	}
	iv := v.Value().(uint64)
	i.value = iv
}

func (i *Uint64) IsComparable() bool {
	return true
}

type Float32 struct {
	value float32
}

func (i *Float32) Value() any {
	return i.value
}

func (i *Float32) TypeName() string {
	return "float32"
}

func (i *Float32) Assign(v Type) {
	if v.TypeName() != "float32" {
		panic("invalid type")
	}
	iv := v.Value().(float32)
	i.value = iv
}

func (i *Float32) IsComparable() bool {
	return true
}

type Float64 struct {
	value float64
}

func (i *Float64) Value() any {
	return i.value
}

func (i *Float64) TypeName() string {
	return "float64"
}

func (i *Float64) Assign(v Type) {
	if v.TypeName() != "float64" {
		panic("invalid type")
	}
	iv := v.Value().(float64)
	i.value = iv
}

func (i *Float64) IsComparable() bool {
	return true
}

type Number struct {
	value float64
}

func (i *Number) Value() any {
	return i.value
}

func (i *Number) TypeName() string {
	return "number"
}

func (i *Number) Assign(v Type) {
	if v.TypeName() != "number" {
		panic("invalid type")
	}
	iv := v.Value().(float64)
	i.value = iv
}

func (i *Number) IsComparable() bool {
	return true
}

func NewInt(i int32) *Int {

	return &Int{
		value: i,
	}
}

func NewByte(b byte) *Byte {

	return &Byte{
		value: b,
	}
}

func NewInt8(i int8) *Int8 {

	return &Int8{
		value: i,
	}
}

func NewInt16(i int16) *Int16 {

	return &Int16{
		value: i,
	}
}

func NewInt32(i int32) *Int32 {

	return &Int32{
		value: i,
	}
}

func NewInt64(i int64) *Int64 {

	return &Int64{
		value: i,
	}
}

func NewUint(ui uint32) *Uint {

	return &Uint{
		value: ui,
	}
}
func NewUint16(ui uint16) *Uint16 {

	return &Uint16{
		value: ui,
	}
}

func NewUint32(ui uint32) *Uint32 {

	return &Uint32{
		value: ui,
	}
}

func NewUint64(ui uint64) *Uint64 {

	return &Uint64{
		value: ui,
	}
}

func NewFloat32(f float32) *Float32 {

	return &Float32{
		value: f,
	}
}

func NewFloat64(f float64) *Float64 {

	return &Float64{
		value: f,
	}
}

func NewNumber(n float64) *Number {

	return &Number{
		value: n,
	}
}
