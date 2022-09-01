package pattern

import "fmt"

/*
	Паттерн применяется, чтобы избавиться от конструктора с большим количеством параметров.
	Так же, применяется, когда есть много представлений одного и того же объекта,
		и все их надо сконструировать похожим способом.

	+ поддерживает пошаговое создание объектов
	+ инкапсулирует сложную логику сборки
	+ нет дублирования кода последовательности шагов

	- усложнение кода
	- клиент взаимодействует с классом строителя??
*/

type PhoneBuilder interface {
	BuildDisplay()
	BuildProcessor()
	BuildMemory()
	BuildBody()
}

type Iphone10 struct {
	parts string
}

type Iphone10Builder struct {
	res string
}

func (i *Iphone10Builder) BuildDisplay() {
	i.res = i.res + "fulHD display."
}

func (i *Iphone10Builder) BuildProcessor() {
	i.res = i.res + "16 core cpu, "
}

func (i *Iphone10Builder) BuildMemory() {
	i.res = i.res + "32 gb ram, "
}

func (i *Iphone10Builder) BuildBody() {
	i.res = i.res + "Iphone 10 contains: "
}

func (i *Iphone10Builder) GetResult() *Iphone10 {
	return &Iphone10{parts: i.res}
}

type OldNokia struct {
	parts string
}

type OldNokiaBuilder struct {
	res string
}

func (n *OldNokiaBuilder) BuildDisplay() {
	n.res = n.res + "liquid cristal display."
}

func (n *OldNokiaBuilder) BuildProcessor() {
	n.res = n.res + "1 core cpu, "
}

func (n *OldNokiaBuilder) BuildMemory() {
	n.res = n.res + "256 kb ram, "
}

func (n *OldNokiaBuilder) BuildBody() {
	n.res = n.res + "Old Nokia contains: "
}

func (n *OldNokiaBuilder) GetResult() *OldNokia {
	return &OldNokia{parts: n.res}
}

type Director struct{}

func (d *Director) BuildPhone(b PhoneBuilder) {
	b.BuildBody()
	b.BuildMemory()
	b.BuildProcessor()
	b.BuildDisplay()
}

func AppBuilder() {
	d := Director{}

	iPhoneBr := &Iphone10Builder{}
	d.BuildPhone(iPhoneBr)
	fmt.Println(iPhoneBr.GetResult())

	OldNokiaBr := &OldNokiaBuilder{}
	d.BuildPhone(OldNokiaBr)
	fmt.Println(OldNokiaBr.GetResult())
}
