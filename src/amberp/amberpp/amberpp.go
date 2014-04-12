package main
import (
	"os"
	"fmt"
	"flag"
	"bytes"
	"regexp"
	"strings"
	"io/ioutil"
	"text/template"
	"text/template/parse"

	"github.com/rzab/amber"
)

type tree struct {
	name string
	leaves []*tree // template/parse has nodes, we ought to have leaves
	parent *tree
	ranged bool
	keys []string
	decl string
}

func(top *tree) touch(words []string) {
	if len(words) == 0 {
		return
	}
	if nt := top.lookup(words[0]); nt != nil {
		nt.touch(words[1:])
		return
	}
	nt := &tree{name: words[0], parent: top}
	top.leaves = append(top.leaves, nt)
	nt.touch(words[1:])
}

func(top *tree) walk(words []string) *tree {
	if len(words) == 0 {
		return top
	}
	for _, leaf := range top.leaves {
		if leaf.name == words[0] {
			return leaf.walk(words[1:])
		}
	}
	return nil
}

func(top tree) lookup(name string) *tree {
	for _, leaf := range top.leaves {
		if name == leaf.name {
			return leaf
		}
	}
	return nil
}

func dotted(bottom *tree) string {
	if bottom == nil || bottom.name == "" {
		return ""
	}
	parent := dotted(bottom.parent)
	if parent == "" {
		return bottom.name
	}
	return parent + "." + bottom.name
}

func indent(top tree, level int) string {
	s := strings.Repeat(" ", level) +"["+ top.name +"]\n"
	level += 2
	for _, leaf := range top.leaves {
		s += indent(*leaf, level)
	}
	level -= 2
	return s
}

func(top tree) String() string {
	return indent(top, 0)
}

type hash map[string]interface{}

func mkmap(top tree) interface{} {
	if len(top.leaves) == 0 {
		return "{" + dotted(&top) + "}"
	}
	h := make(hash)
	for _, leaf := range top.leaves {
		if leaf.ranged {
			if len(leaf.keys) != 0 {
				kv := make(map[string]string)
				for _, k := range leaf.keys {
					kv[k] = "{"+ leaf.decl +"."+ k +"}"
				}
				h[leaf.name] = []map[string]string{kv}
			} else {
				h[leaf.name] = []string{}
			}
		} else {
			h[leaf.name] = mkmap(*leaf)
		}
	}
	return h
}

/* func string_hash(h interface{}) string {
	return hindent(h.(hash), 0)
}

func hindent(h hash, level int) string {
	s := ""
	for k, v := range h {
		s += strings.Repeat(" ", level) +"("+ k +")\n"
		vv, ok := v.(hash)
		if ok && len(vv) > 0 {
			level += 2
			s += hindent(vv, level)
			level -= 2
		} else {
			s += strings.Repeat(" ", level + 2) + fmt.Sprint(v) + "\n"
		}
	}
	return s
} // */

var bracketTrack = map[string]string{}

func bracketGet(key string) string {
	return bracketTrack[key]
}

func bracketSet(clause, value string) string {
	bracketTrack[clause] = value
	return ""
}

func bracketDel(key string) string {
	delete(bracketTrack, key)
	return ""
}

var bracketFuncs = map[string]interface{}{
	"GET": bracketGet,
	"SET": bracketSet,
	"DEL": bracketDel,
}

func main() {
	var  outputFile string
	var definesFile string
	var prettyPrint bool
	var jscriptMode bool

	flag.StringVar(& outputFile, "o",           "",    "Output file")
	flag.StringVar(& outputFile, "output",      "",    "Output file")
	flag.StringVar(&definesFile, "d",           "",    "Defines file")
	flag.StringVar(&definesFile, "defines",     "",    "Defines file")
	flag.  BoolVar(&prettyPrint, "pp",          false, "Pretty print")
	flag.  BoolVar(&prettyPrint, "prettyprint", false, "Pretty print")
	flag.  BoolVar(&jscriptMode, "j",           false, "Javascript mode")
	flag.  BoolVar(&jscriptMode, "javascript",  false, "Javascript mode")

	flag.Parse()
	inputFile := flag.Arg(0)

	if inputFile == "" {
		fmt.Fprintf(os.Stderr, "No input file specified.")
		flag.Usage()
		os.Exit(2)
	}

	inputText := ""
	if definesFile != "" {
		b, err := ioutil.ReadFile(definesFile)
		check(err)
		inputText += compile(b, prettyPrint)
	}
	b, err := ioutil.ReadFile(inputFile)
	check(err)
	inputText += compile(b, prettyPrint)

	fstplate, err := template.New("fst").Funcs(bracketFuncs).Delims("[[", "]]").Parse(inputText)
	check(err)
	fst := execute(fstplate, nil)

	if !jscriptMode {
		writeFile(outputFile, fst)
		return
	}

	sndplate, err := template.New("snd").Funcs(template.FuncMap(amber.FuncMap)).Parse(fst)
	check(err)

	m := data(sndplate.Tree) //; fmt.Printf("data => %+v\nstring_hash(data) => %+v", m, string_hash(m))
	snd := execute(sndplate, m)

	snd = regexp.MustCompile("</?script>").ReplaceAllLiteralString(snd, "")
	snd = strings.Replace(snd, " class=", " className=", -1) // fingers crossed

	writeFile(outputFile, snd)
}

func writeFile(optFilename, s string) {
	b := []byte(s)
	if optFilename != "" {
		check(ioutil.WriteFile(optFilename, b, 0644))
	} else {
		os.Stdout.Write(b)
	}
}

func compile(input []byte, prettyPrint bool) string {
	compiler := amber.New()
	compiler.PrettyPrint = prettyPrint // compiler.Options.PrettyPrint?

	check(compiler.Parse(string(input)))
	s, err := compiler.CompileString()
	check(err)
	return s
}

func execute(emplate *template.Template, data interface{}) string {
	buf := new(bytes.Buffer)
	check(emplate.Execute(buf, data))
	return buf.String()
}

func data(TREE *parse.Tree) interface{} {
	if TREE == nil || TREE.Root == nil {
		return "{}" // mkmap(tree{})
	}

	data := tree{}
	vars := map[string][]string{}

	for _, node := range TREE.Root.Nodes { // here we go
		switch node.Type() {
		case parse.NodeAction:
			actionNode := node.(*parse.ActionNode)
			decl := actionNode.Pipe.Decl

			for _, cmd := range actionNode.Pipe.Cmds {
				if cmd.NodeType != parse.NodeCommand {
					continue
				}
				for _, arg := range cmd.Args {
					var ident []string
					switch arg.Type() {

					case parse.NodeField:
						ident = arg.(*parse.FieldNode).Ident

						if len(decl) > 0 && len(decl[0].Ident) > 0 {
							vars[decl[0].Ident[0]] = ident
						}
						data.touch(ident)

					case parse.NodeVariable:
						ident = arg.(*parse.VariableNode).Ident

						if words, ok := vars[ident[0]]; ok {
							words := append(words, ident[1:]...)
							data.touch(words)
							if len(decl) > 0 && len(decl[0].Ident) > 0 {
								vars[decl[0].Ident[0]] = words
							}
						}
					}
				}
			}
		case parse.NodeRange:
			rangeNode := node.(*parse.RangeNode)
			decl := rangeNode.Pipe.Decl[len(rangeNode.Pipe.Decl) - 1].String()
			keys := []string{}

			for _, ifnode := range rangeNode.List.Nodes {
				switch ifnode.Type() {
				case parse.NodeAction:
					keys = append(keys, getKeys(decl, ifnode)...)
				case parse.NodeIf:
					for _, z := range ifnode.(*parse.IfNode).List.Nodes {
						if z.Type() == parse.NodeAction {
							keys = append(keys, getKeys(decl, z)...)
						}
					}
				}
			}

			// fml
			arg0 := rangeNode.Pipe.Cmds[0].Args[0].String()
			if words, ok := vars[arg0]; ok {
				if leaf := data.walk(words); leaf != nil {
					leaf.ranged = true
					leaf.keys = append(leaf.keys, keys...)
					leaf.decl = decl // redefined $
				}
			}
		}
	}
	return mkmap(data)
}

func getKeys(decl string, parseNode parse.Node) (keys []string) {
	for _, cmd := range parseNode.(*parse.ActionNode).Pipe.Cmds {
		if cmd.NodeType != parse.NodeCommand {
			continue
		}
		for _, arg := range cmd.Args {
			if arg.Type() != parse.NodeVariable {
				continue
			}
			ident := arg.(*parse.VariableNode).Ident
			if len(ident) < 2 || ident[0] != decl {
				continue
			}
			keys = append(keys, ident[1])
		}
	}
	return
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}












