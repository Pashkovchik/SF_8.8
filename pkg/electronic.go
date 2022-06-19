package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	Phone
	ButtonsCount() int
}

type Smartphone interface {
	Phone
	OS() string //название операционной системы
}

type applePhone struct { //Phone и Smartphone
	model string
}

type androidPhone struct { //Phone и Smartphone
	brand string //Samsung, LG, Xiaomi и так далее
	model string
}

type radioPhone struct { //Phone и StationPhone
	model        string
	brand        string
	buttonsCount int // количество кнопок
}

func (m *applePhone) Brand() string {
	return "Apple"
}

func (m *androidPhone) Brand() string {
	return m.brand
}

func (m *radioPhone) Brand() string {
	return m.brand
}

func (m *radioPhone) ButtonsCount() int {
	return m.buttonsCount
}

func (a *applePhone) Type() string {
	return "smartphone"
}

func (a *androidPhone) Type() string {
	return "smartphone"
}

func (a *radioPhone) Type() string {
	return "station"
}

func (a *applePhone) Model() string {
	return a.model
}

func (a *androidPhone) Model() string {
	return a.model
}

func (a *radioPhone) Model() string {
	return a.model
}

func (a *applePhone) OS() string {
	return "IOS"
}

func (a *androidPhone) OS() string {
	return "android"
}

func NewApplePhone(s string) *applePhone {
	return &applePhone{model: s}
}

func NewAndroidPhone(m, b string) *androidPhone {
	return &androidPhone{model: m, brand: b}
}

func NewRadioPhone(m, b string, c int) *radioPhone {
	return &radioPhone{model: m, brand: b, buttonsCount: c}
}

/*
type pet struct {
  name string
  age  int
}

type cat struct {
  pet
  furColor string //цвет шерсти
}

func NewCat(name string, age int, furColor string) cat {
  return cat{pet: pet{name: name, age: age}, furColor: furColor}
}
*/
