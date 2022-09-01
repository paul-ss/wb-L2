package pattern

import "fmt"

/*
	Фасад примеяется тогда, когда нужно организовать простой интерфейс к большой и сложной системе.
	Инкапсулирует в себе логику создания и взаиможействия с компонентами (и компонентов) сложной системы.

	+ Снимает с пользователя нагрузку при работе со сложной системой
	- Может превратиться в God object

	Пример: На сервер приходит картинка, нужно ее пожать на несколько форматов, сделать аву, ...
	Фасад:
		CompressPicture(path string) (*Compressed, error) - коннект и обращение к сервисам, ...
*/

type Formats struct {
	AvatarBig   string
	AvatarSmall string
	Post        string
}

// Facade

type FormatterFacade struct{}

func (f *FormatterFacade) FormatPicture(path string) (Formats, error) {
	filter := Filter{}
	path = filter.ApplyFilters(path, "some params")

	res := Formats{}

	cmp := Compressor{}
	res.Post = cmp.Compress(path, "post size")

	avaSmall := cmp.Compress(path, "ava small size")
	avaBig := cmp.Compress(path, "ava big size")

	shp := Shaper{}

	res.AvatarSmall = shp.CutToShape(avaSmall, "circle")
	res.AvatarBig = shp.CutToShape(avaBig, "circle")

	return res, nil
}

// Filter

type Filter struct{}

func (f *Filter) ApplyFilters(path string, params interface{}) string {
	return "newPath"
}

// Compressor

type Compressor struct{}

func (c *Compressor) Compress(path string, params interface{}) string {
	return "compressed"
}

// Shaper

type Shaper struct{}

func (s *Shaper) CutToShape(path string, params interface{}) string {
	return "shape"
}

func AppFacade() {
	fr := FormatterFacade{}
	formats, _ := fr.FormatPicture("path to pic")
	fmt.Println(formats)
}
