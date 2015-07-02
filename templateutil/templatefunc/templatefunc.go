package templatefunc

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/ostrost/ostent/params"
	"github.com/ostrost/ostent/system/operating"
	"github.com/ostrost/ostent/templateutil/templatepipe"
)

func (f JSXFuncs) classWord() string   { return "className" }
func (f JSXFuncs) colspanWord() string { return "colSpan" }
func (f JSXFuncs) forWord() string     { return "htmlFor" }

func (f HTMLFuncs) classWord() string   { return "class" }
func (f HTMLFuncs) colspanWord() string { return "colspan" }
func (f HTMLFuncs) forWord() string     { return "for" }

// jsxClose returns close tag markup as template.HTML.
func (f JSXFuncs) jsxClose(tag string) template.HTML { return template.HTML("</" + tag + ">") } // f is unused

// jsxClose returns empty template.HTML.
func (f HTMLFuncs) jsxClose(string) (empty template.HTML) { return } // f is unused

func (f JSXFuncs) droplink(value interface{}, args ...string) (interface{}, error) {
	// f is unused
	named, aclass := DropLinkArgs(args)
	vstring := ToString(value)
	_, pname := DotSplit(vstring)
	enums := params.NewParamsENUM(nil)
	ed := enums[pname].EnumDecodec
	return params.DropLink{
		AlignClass: aclass,
		Text:       ed.Text(named), // always static
		Href:       fmt.Sprintf("{%s.%s.%s}", vstring, named, "Href"),
		Class:      fmt.Sprintf("{%s.%s.%s}", vstring, named, "Class"),
		CaretClass: fmt.Sprintf("{%s.%s.%s}", vstring, named, "CaretClass"),
	}, nil
}

func (f HTMLFuncs) droplink(value interface{}, args ...string) (interface{}, error) {
	named, aclass := DropLinkArgs(args)
	ep, ok := value.(*params.EnumParam)
	if !ok {
		return nil, f.CastError("*params.EnumParam")
	}
	pname, uptr := ep.EnumDecodec.Unew()
	if err := uptr.Unmarshal(named, new(bool)); err != nil {
		return nil, err
	}
	l := ep.EncodeUint(pname, uptr)
	l.AlignClass = aclass
	return l, nil
}

func DropLinkArgs(args []string) (string, string) {
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

func (f JSXFuncs) usepercent(value interface{}) interface{} {
	ca := fmt.Sprintf(" %s={LabelClassColorPercent(%s)}", f.classWord(), f.uncurlv(value))
	return struct {
		Value     string
		ClassAttr template.HTMLAttr
	}{
		Value:     ToString(value),
		ClassAttr: template.HTMLAttr(ca),
	}
}

func (f HTMLFuncs) usepercent(value interface{}) interface{} {
	vstring := ToString(value)
	ca := fmt.Sprintf(" %s=%q", f.classWord(), LabelClassColorPercent(vstring))
	return struct {
		Value     string
		ClassAttr template.HTMLAttr
	}{
		Value:     vstring,
		ClassAttr: template.HTMLAttr(ca),
	}
}

// JSXFuncs has methods implementing Functor.
type JSXFuncs struct{}

// HTMLFuncs has methods implementing Functor.
type HTMLFuncs struct{}

// MakeMap is dull but required.
func (f JSXFuncs) MakeMap() template.FuncMap { return MakeMap(f) }

// MakeMap is dull but required.
func (f HTMLFuncs) MakeMap() template.FuncMap { return MakeMap(f) }

// MakeMap constructs template.FuncMap off f implementation.
func MakeMap(f Functor) template.FuncMap {
	return template.FuncMap{
		"rowsset": func(interface{}) string { return "" }, // empty pipeline
		// acepp overrides rowsset and adds setrows

		"droplink":   f.droplink,
		"usepercent": f.usepercent,
		"jsxClose":   f.jsxClose,
		"class":      f.classWord,
		"colspan":    f.colspanWord,
		"for":        f.forWord,
	}
}

// Funcs features functions for templates. In use in acepp and templates.
var Funcs = HTMLFuncs{}.MakeMap()

type Functor interface {
	MakeMap() template.FuncMap
	droplink(interface{}, ...string) (interface{}, error)
	usepercent(interface{}) interface{}
	jsxClose(string) template.HTML
	classWord() string
	colspanWord() string
	forWord() string
}

func init() {
	// check for Nota's interfaces compliance
	_ = interface {
		// operating (multiple types):
		BoolClassAttr(...string) (template.HTMLAttr, error)
		Clip(int, string, ...operating.ToStringer) (*operating.Clipped, error)
		KeyAttr(string) template.HTMLAttr

		FormActionAttr() interface{}                                        // Query
		BoolParamClassAttr(...string) (template.HTMLAttr, error)            // BoolParam
		DisabledAttr() interface{}                                          // BoolParam
		ToggleHrefAttr() interface{}                                        // BoolParam
		EnumClassAttr(string, string, ...string) (template.HTMLAttr, error) // EnumParam
		PeriodNameAttr() interface{}                                        // PeriodParam
		PeriodValueAttr() interface{}                                       // PeriodParam
		RefreshClassAttr(string) interface{}                                // PeriodParam
		LessHrefAttr() interface{}                                          // LimitParam
		MoreHrefAttr() interface{}                                          // LimitParam
	}(templatepipe.Nota(nil))
}

// DotSplit splits s by last ".".
func DotSplit(s string) (string, string) {
	if s == "" {
		return "", ""
	}
	i := len(s) - 1
	for i > 0 && s[i] != '.' {
		i--
	}
	return s[:i], s[i+1:]
}

// DotSplitV calls DotSplit with value's string.
func DotSplitV(value interface{}) (string, string) {
	return DotSplit(ToString(value))
}

func ToString(value interface{}) string {
	if s, ok := value.(string); ok {
		return s
	}
	return value.(templatepipe.Nota).ToString()
}

func (f HTMLFuncs) CastError(notype string) error {
	// f is unused
	return fmt.Errorf("Cannot convert into %s", notype)
}

func (f JSXFuncs) uncurl(s string) string {
	// f is unused
	return strings.TrimSuffix(strings.TrimPrefix(s, "{"), "}")
}

func (f JSXFuncs) uncurlv(v interface{}) string {
	return f.uncurl(ToString(v))
}

func LabelClassColorPercent(p string) string {
	if len(p) > 2 { // 100% and more
		return "label label-danger"
	}
	if len(p) > 1 {
		if p[0] == '9' {
			return "label label-danger"
		}
		if p[0] == '8' {
			return "label label-warning"
		}
		if p[0] == '1' {
			return "label label-success"
		}
		return "label label-info"
	}
	return "label label-success"
}

// SetKFunc constructs a func which
// sets k key to templatepipe.Curly(string (n))
// in passed interface{} (v) being a templatepipe.Nota.
// SetKFunc is used by acepp only.
func SetKFunc(k string) func(interface{}, string) interface{} {
	return func(v interface{}, n string) interface{} {
		if args := strings.Split(n, " "); len(args) > 1 {
			var list []string
			for _, arg := range args {
				list = append(list, templatepipe.Curl(arg))
			}
			v.(templatepipe.Nota)[k] = list
			return v
		}
		v.(templatepipe.Nota)[k] = templatepipe.Curl(n)
		return v
	}
}

// GetKFunc constructs a func which
// gets, deletes and returns k key
// in passed interface{} (v) being a templatepipe.Nota.
// GetKFunc is used by acepp only.
func GetKFunc(k string) func(interface{}) interface{} {
	return func(v interface{}) interface{} {
		h, ok := v.(templatepipe.Nota)
		if !ok {
			return "" // empty pipeline, affects dispatch
		}
		n := h[k]
		if args, ok := n.([]string); ok {
			if len(args) > 1 {
				h[k] = args[1:]
			} else {
				delete(h, k)
			}
			return args[0]
		}
		delete(h, k)
		return n // may also be empty, affects dispatch
	}
}
