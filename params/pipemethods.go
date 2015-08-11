package params

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/ostrost/ostent/flags"
)

func SprintfAttr(format string, args ...interface{}) template.HTMLAttr {
	return template.HTMLAttr(fmt.Sprintf(format, args...))
}

// AttrActionForm is for template.
func (p Params) AttrActionForm() (interface{}, error) {
	qs, err := p.Encode()
	if err != nil {
		return nil, err
	}
	return SprintfAttr(" action=\"/form/%s\"", url.QueryEscape(qs)), nil
}

func (p Params) AttrClassP(num *Num, fstclass, sndclass string) template.HTMLAttr {
	return p.AttrClassN(!num.Head, fstclass, sndclass)
}

func (p Params) AttrClassZero(num *Num, fstclass, sndclass string) template.HTMLAttr {
	return p.AttrClassN(num.Body == 0, fstclass, sndclass)
}

func (p Params) AttrClassN(b bool, fstclass, sndclass string) template.HTMLAttr {
	// p is unused
	s := fstclass
	if !b {
		s = sndclass
	}
	return SprintfAttr(" class=%q", s)
}

/*
func (bp BoolParam) DisabledAttr() interface{} {
	if !bp.Value {
		return template.HTMLAttr("")
	}
	return SprintfAttr(" disabled=%q", "disabled")
}

func (bp BoolParam) ToggleHrefAttr() interface{} {
	return SprintfAttr(" href=\"%s\"", bp.EncodeToggle())
}
*/

func (p *Params) AttrClassParamsError(m MultiError, name, fstclass, sndclass string) template.HTMLAttr {
	_, ok := m[name]
	return p.AttrClassN(ok, fstclass, sndclass)
}

func (p Params) AttrClassT(num Num, cmp int, fstclass, sndclass string) template.HTMLAttr {
	return p.AttrClassN(num.Body == cmp, fstclass, sndclass)
}

// p is a pointer to flip (twice) the b.
func (p *Params) HrefToggle(b *bool) (string, error) {
	*b = !*b
	qs, err := p.Encode()
	*b = !*b
	return "?" + qs, err
}

func (p *Params) HrefToggleHead(num *Num) (string, error) {
	num.Head = !num.Head
	qs, err := p.Encode()
	num.Head = !num.Head
	return "?" + qs, err
}

type APlain struct {
	Href  string
	Text  string
	Badge string `json:",omitempty"`
}

type ALink struct {
	APlain
	ExtraClass string `json:"-"`
}

func (al ALink) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		APlain
		Class string `json:",omitempty"`
	}{
		APlain: al.APlain,
		Class:  al.Class(""),
	})
}

func (al ALink) Class(base string) string {
	if base == "" {
		return al.ExtraClass
	}
	return base + " " + al.ExtraClass
}

// p is a pointer to alter (and revert) v being part of p.
func (p *Params) AttrHrefToggle(v *bool) (interface{}, error) {
	href, err := p.HrefToggle(v)
	return SprintfAttr(" href=%q", href), err
}

// p is a pointer to alter (and revert) num being part of p.
func (p *Params) AttrHrefToggleHead(num *Num) (interface{}, error) {
	href, err := p.HrefToggleHead(num)
	return SprintfAttr(" href=%q", href), err
}

// TODO In Decoder, have a cache of type=>fieldName=>tag(got,splitted)
func (p Params) AttrNameRefresh(fieldName string) (interface{}, error) {
	field, ok := reflect.TypeOf(p).FieldByName(fieldName)
	if !ok {
		return nil, fmt.Errorf("Params has no field %q", fieldName)
	}
	tag := strings.Split(field.Tag.Get("schema"), ",")[0]
	return SprintfAttr(" name=%q", tag), nil
}

func (p Params) AttrValueRefresh(fieldName string) (interface{}, error) {
	field, ok := reflect.TypeOf(p).FieldByName(fieldName)
	if !ok {
		return nil, fmt.Errorf("Params has no field %q", fieldName)
	}
	tag := strings.Split(field.Tag.Get("schema"), ",")[0]
	values, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	v, ok := values[tag]
	if !ok || len(v) == 0 || v[0] == "" {
		return template.HTMLAttr(""), nil
	}
	return SprintfAttr(" value=%q", v[0]), nil
}

/*
func (ep EnumParam) EnumClassAttr(named, classif string, optelse ...string) (template.HTMLAttr, error) {
	classelse, err := EnumClassAttrArgs(optelse)
	if err != nil {
		return template.HTMLAttr(""), err
	}
	_, uptr := ep.EnumDecodec.Unew()
	if err := uptr.Unmarshal(named, new(bool)); err != nil {
		return template.HTMLAttr(""), err
	}
	if ep.Number.Uint != uptr.Touint() {
		classif = classelse
	}
	return SprintfAttr(" class=%q", classif), nil
}

func EnumClassAttrArgs(opt []string) (string, error) {
	if len(opt) == 1 {
		return opt[0], nil
	} else if len(opt) > 1 {
		return "", fmt.Errorf("number of args for EnumClassAttr: either 2 or 3 got %d",
			2+len(opt))
	}
	return "", nil
}
// */

type Varlink struct {
	AlignClass string
	CaretClass string
	LinkClass  string
	LinkHref   string
	LinkText   string `json:"-"` // static
}

func (p *Params) Variate(num *Num, cmp int, text, alignClass string) (Varlink, error) {
	vl := Varlink{LinkText: text, LinkClass: "state"}
	var toggleh bool
	if num.Body == cmp {
		vl.CaretClass = "caret"
		vl.LinkClass += " current"
		if num.Body == cmp {
			vl.LinkClass += " dropup"
			toggleh = true
		}
	}
	qs, err := p.EncodeN(num, cmp, &toggleh)
	if err != nil {
		return Varlink{}, err
	}
	vl.LinkHref = qs
	return vl, nil
}

/*
func (ep EnumParam) EnumLink(args ...string) (EnumLink, error) {
	named, aclass := EnumLinkArgs(args)
	pname, uptr := ep.EnumDecodec.Unew()
	if err := uptr.Unmarshal(named, new(bool)); err != nil {
		return EnumLink{}, err
	}
	l := ep.EncodeUint(pname, uptr)
	l.AlignClass = aclass
	return l, nil
}

func EnumLinkArgs(args []string) (string, string) {
	var named string
	if len(args) > 0 {
		named = args[0]
	}
	aclass := "text-right" // default
	if len(args) > 1 {
		aclass = ""
		if args[1] != "" {
			aclass = "text-" + args[1]
		}
	}
	return named, aclass
}

func (pp PeriodParam) PeriodNameAttr() interface{} {
	return SprintfAttr(" name=%q", pp.Pname)
}

func (pp PeriodParam) PeriodValueAttr() interface{} {
	if pp.Input == "" {
		return template.HTMLAttr("")
	}
	return SprintfAttr(" value=\"%s\"", pp.Input)
}

func (pp PeriodParam) RefreshClassAttr(classes string) interface{} {
	if pp.InputErrd {
		classes += " has-warning"
	}
	return SprintfAttr(" class=%q", classes)
}
*/

func (p *Params) EncodeN(num *Num, body int, thead *bool) (string, error) {
	copy := num.Body
	num.Body = body
	if thead != nil && *thead {
		num.Head = !num.Head
	}
	qs, err := p.Encode()
	num.Body = copy
	if thead != nil && *thead {
		num.Head = !num.Head
	}
	return "?" + qs, err
}

func Pow2Less(v int) int {
	switch v {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 1
	}
	g := math.Log2(float64(v))
	n := math.Floor(g)
	if n == g {
		n--
	}
	return int(math.Pow(2, n))
}

func Pow2More(v int) int {
	switch v {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	}
	if v <= 32768 { // up to 65536
		v = int(math.Pow(2, 1+math.Floor(math.Log2(float64(v)))))
	}
	return v
}

func (p *Params) ZeroN(num *Num) (ALink, error) { return p.LinkN(num, 0, "") }
func (p *Params) MoreN(num *Num) (ALink, error) { return p.LinkN(num, Pow2More(num.Body), "+") }
func (p *Params) LessN(num *Num) (ALink, error) { return p.LinkN(num, Pow2Less(num.Body), "-") }

func (p *Params) LinkN(num *Num, body int, badge string) (ALink, error) {
	href, err := p.EncodeN(num, body, nil)
	if err != nil {
		return ALink{}, err
	}
	var class string
	if badge == "" && num.Body == 0 { // "0" case && param is 0
		class = " disabled active"
	}
	if badge == "+" && body >= num.Limit {
		class = " disabled"
	}
	if badge == "-" && body == 0 {
		class = " disabled"
	}
	return ALink{
		APlain: APlain{
			Href:  href,
			Text:  fmt.Sprintf("%d", body),
			Badge: badge,
		},
		ExtraClass: class,
	}, nil
}

func (p *Params) EncodeD(dur *Duration, set time.Duration) (string, error) {
	copy := dur.D
	dur.D = set
	qs, err := p.Encode()
	dur.D = copy
	return "?" + qs, err
}

func DurationMore(dur Duration, step time.Duration) time.Duration {
	const s = time.Second
	const m = time.Second * 60
	var table = map[time.Duration]time.Duration{
		s:      2 * s,
		2 * s:  5 * s,
		5 * s:  10 * s,
		10 * s: 30 * s,
		30 * s: m,
		m:      2 * m,
		2 * m:  5 * m,
		5 * m:  10 * m,
		10 * m: 30 * m,
		30 * m: 60 * m,
	}
	if d, ok := table[dur.D]; ok {
		return d
	}
	return dur.D + step
}
func DurationLess(dur Duration, step time.Duration) time.Duration {
	const s = time.Second
	const m = time.Second * 60
	var table = map[time.Duration]time.Duration{
		s:      s,
		2 * s:  s,
		5 * s:  2 * s,
		10 * s: 5 * s,
		30 * s: 10 * s,
		m:      30 * s,
		2 * m:  m,
		5 * m:  2 * m,
		10 * m: 5 * m,
		30 * m: 10 * m,
		60 * m: 30 * m,
	}
	if d, ok := table[dur.D]; ok {
		return d
	}
	return dur.D - step
}

func (p *Params) MoreD(dur *Duration) (ALink, error) {
	return p.LinkD(dur, DurationMore(*dur, p.MinPeriod.Duration), "+")
}
func (p *Params) LessD(dur *Duration) (ALink, error) {
	return p.LinkD(dur, DurationLess(*dur, p.MinPeriod.Duration), "-")
}

func (p *Params) LinkD(dur *Duration, set time.Duration, badge string) (ALink, error) {
	href, err := p.EncodeD(dur, set)
	if err != nil {
		return ALink{}, err
	}
	var class string
	if badge == "-" && dur.D == p.MinPeriod.Duration {
		class = " disabled"
	}
	return ALink{
		APlain: APlain{
			Href:  href,
			Text:  flags.DurationString(set),
			Badge: badge,
		},
		ExtraClass: class,
	}, nil
}
