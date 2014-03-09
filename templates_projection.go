package main

import (
	"errors"
	"fmt"
	"sort"
	"text/template"
)

func getProjectionTemplate(name string) (result *template.Template, err error) {
	t, found := ProjectionTemplates[name]
	if found {
		result = template.Must(template.New(name).Parse(t.Text))
	} else {
		err = errors.New(fmt.Sprintf("%s is not a known projection method", name))
	}
	return
}

func isProjectionMethod(s string) bool {
	_, ok := ProjectionTemplates[s]
	return ok
}

func getProjectionMethodKeys() (result []string) {
	for k := range ProjectionTemplates {
		result = append(result, k)
	}
	sort.Strings(result)
	return
}

var ProjectionTemplates = map[string]*Template{
	"Aggregate": &Template{
		Text: `
// {{.MethodName}} iterates over {{.Parent.Plural}}, operating on each element while maintaining ‘state’. See: http://clipperhouse.github.io/gen/#Aggregate
func (rcv {{.Parent.Plural}}) {{.MethodName}}(fn func({{.Type}}, {{.Parent.Pointer}}{{.Parent.Name}}) {{.Type}}) (result {{.Type}}) {
	for _, v := range rcv {
		result = fn(result, v)
	}
	return
}
`},

	"Average": &Template{
		Text: `
// {{.MethodName}} sums {{.Type}} over all elements and divides by len({{.Parent.Plural}}). See: http://clipperhouse.github.io/gen/#Average
func (rcv {{.Parent.Plural}}) {{.MethodName}}(fn func({{.Parent.Pointer}}{{.Parent.Name}}) {{.Type}}) (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine {{.MethodName}} of zero-length {{.Parent.Plural}}")
		return
	}
	for _, v := range rcv {
		result += fn(v)
	}
	result = result / {{.Type}}(l)
	return
}
`,
		RequiresNumeric: true,
	},

	"GroupBy": &Template{
		Text: `
// {{.MethodName}} groups elements into a map keyed by {{.Type}}. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv {{.Parent.Plural}}) {{.MethodName}}(fn func({{.Parent.Pointer}}{{.Parent.Name}}) {{.Type}}) map[{{.Type}}]{{.Parent.Plural}} {
	result := make(map[{{.Type}}]{{.Parent.Plural}})
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
`,
		RequiresComparable: true,
	},

	"Max": &Template{
		Text: `
// {{.MethodName}} selects the largest value of {{.Type}} in {{.Parent.Plural}}. Returns error on {{.Parent.Plural}} with no elements. See: http://clipperhouse.github.io/gen/#MaxCustom
func (rcv {{.Parent.Plural}}) {{.MethodName}}(fn func({{.Parent.Pointer}}{{.Parent.Name}}) {{.Type}}) (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine {{.MethodName}} of zero-length {{.Parent.Plural}}")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f > result {
				result = f
			}
		}
	}
	return
}
`,
		RequiresOrdered: true,
	},

	"Min": &Template{
		Text: `
// {{.MethodName}} selects the least value of {{.Type}} in {{.Parent.Plural}}. Returns error on {{.Parent.Plural}} with no elements. See: http://clipperhouse.github.io/gen/#MinCustom
func (rcv {{.Parent.Plural}}) {{.MethodName}}(fn func({{.Parent.Pointer}}{{.Parent.Name}}) {{.Type}}) (result {{.Type}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine {{.MethodName}} of zero-length {{.Parent.Plural}}")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}
`,
		RequiresOrdered: true,
	},

	"Select": &Template{
		Text: `
// {{.MethodName}} returns a slice of {{.Type}} in {{.Parent.Plural}}, projected by passed func. See: http://clipperhouse.github.io/gen/#Select
func (rcv {{.Parent.Plural}}) {{.MethodName}}(fn func({{.Parent.Pointer}}{{.Parent.Name}}) {{.Type}}) (result []{{.Type}}) {
	for _, v := range rcv {
		result = append(result, fn(v))
	}
	return
}
`,
	},

	"Sum": &Template{
		Text: `
// {{.MethodName}} sums {{.Type}} over elements in {{.Parent.Plural}}. See: http://clipperhouse.github.io/gen/#Sum
func (rcv {{.Parent.Plural}}) {{.MethodName}}(fn func({{.Parent.Pointer}}{{.Parent.Name}}) {{.Type}}) (result {{.Type}}) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}
`,
		RequiresNumeric: true,
	},
}