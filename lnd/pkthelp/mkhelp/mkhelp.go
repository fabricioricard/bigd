package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
type Field struct {
	Name        string
	Description []string
	Repeated    bool
	Type        Type
}

type Varient struct {
	Name        string
	Description []string
}

type Type struct {
	Name        string
	Description []string
	Fields      []Field
}

type Method struct {
	Name        string
	Service     string
	Description []string
	Req         Type
	Res         Type
}

var EnumVarientType Type = Type{
	Name: "ENUM_VARIENT",
}
*/
func desc(desc string, padding string) {
	if len(desc) > 0 {
		fmt.Printf("%sDescription: []string{\n", padding)
		for _, l := range strings.Split(desc, "\n") {
			fmt.Printf("%s    %s,\n", padding, strconv.Quote(l))
		}
		fmt.Printf("%s},\n", padding)
	}
}

func fixName(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: genhelp dirpath")
		os.Exit(100)
		return
	}
	path := os.Args[len(os.Args)-1]
	dir, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	files, err := dir.ReadDir(0)
	if err != nil {
		panic(err.Error())
	}
	var templates []Template
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".doc.json") {
			continue
		}
		//fmt.Println("Generating help " + path + "/" + f.Name())
		content, err := os.ReadFile(path + "/" + f.Name())
		if err != nil {
			panic(err.Error())
		}
		t := Template{}
		if err := json.Unmarshal(content, &t); err != nil {
			panic(err.Error())
		}
		templates = append(templates, t)
	}
	fmt.Printf("package pkthelp\n")
	for _, t := range templates {
		for _, s := range t.Scalars {
			fmt.Printf("func mk%s() Type {\n", s.ProtoType)
			fmt.Printf("    return Type{\n")
			fmt.Printf("        Name: %s,\n", strconv.Quote(s.GoType))
			fmt.Printf("    }\n")
			fmt.Printf("}\n")
		}
		break
	}
	for _, t := range templates {
		for _, f := range t.Files {
			for _, e := range f.Enums {
				fmt.Printf("func mk%s() Type {\n", fixName(e.FullName))
				fmt.Printf("    return Type{\n")
				fmt.Printf("        Name: %s,\n", strconv.Quote(fixName(e.FullName)))
				desc(e.Description, "        ")
				fmt.Printf("        Fields: []Field{\n")
				for _, v := range e.Values {
					fmt.Printf("            {\n")
					fmt.Printf("                Name: %s,\n", strconv.Quote(v.Name))
					desc(v.Description, "                ")
					fmt.Printf("                Type: EnumVarientType,\n")
					fmt.Printf("            },\n")
				}
				fmt.Printf("        },\n")
				fmt.Printf("    }\n")
				fmt.Printf("}\n")
			}
		}
	}
	for _, t := range templates {
		for _, f := range t.Files {
			for _, e := range f.Messages {
				fmt.Printf("func mk%s() Type {\n", fixName(e.FullName))
				fmt.Printf("    return Type{\n")
				fmt.Printf("        Name: %s,\n", strconv.Quote(fixName(e.FullName)))
				desc(e.Description, "        ")
				if len(e.Fields) > 0 {
					fmt.Printf("        Fields: []Field{\n")
					for _, f := range e.Fields {
						fmt.Printf("            {\n")
						fmt.Printf("                Name: %s,\n", strconv.Quote(f.Name))
						desc(f.Description, "                ")
						if f.Label == "repeated" {
							fmt.Printf("                Repeated: true,\n")
						}
						fmt.Printf("                Type: mk%s(),\n", fixName(f.FullType))
						fmt.Printf("            },\n")
					}
					fmt.Printf("        },\n")
				}
				fmt.Printf("    }\n")
				fmt.Printf("}\n")
			}
		}
	}

	var categoryRegexp *regexp.Regexp
	var shortDescriptionRegexp *regexp.Regexp

	categoryRegexp, err = regexp.Compile("\\$pld\\.category:\\s*`([^`]+)`")
	if err != nil {
		panic(err.Error())
	}

	shortDescriptionRegexp, err = regexp.Compile("\\$pld\\.short_description:\\s*`([^`]+)`")
	if err != nil {
		panic(err.Error())
	}

	for _, t := range templates {
		for _, f := range t.Files {
			for _, s := range f.Services {
				for _, m := range s.Methods {
					fmt.Printf("func %s_%s() Method {\n", s.Name, m.Name)
					fmt.Printf("    return Method{\n")
					fmt.Printf("        Name: %s,\n", strconv.Quote(m.Name))
					fmt.Printf("        Service: %s,\n", strconv.Quote(s.Name))
					if len(s.Description) > 0 {

						var match []string
						var matchIndex []int

						match = categoryRegexp.FindStringSubmatch(m.Description)
						if len(match) > 1 {
							fmt.Printf("        Category: %s,\n", strconv.Quote(match[1]))

							matchIndex = categoryRegexp.FindStringIndex(m.Description)
							m.Description = m.Description[0:matchIndex[0]] + m.Description[matchIndex[1]:]
						}

						match = shortDescriptionRegexp.FindStringSubmatch(m.Description)
						if len(match) > 1 {
							fmt.Printf("        ShortDescription: %s,\n", strconv.Quote(match[1]))

							matchIndex = shortDescriptionRegexp.FindStringIndex(m.Description)
							m.Description = m.Description[0:matchIndex[0]] + m.Description[matchIndex[1]:]
						}

						fmt.Printf("        Description: []string{\n")
						for _, s := range strings.Split(m.Description, "\n") {
							descriptionLine := strings.TrimSpace(s)
							if len(descriptionLine) > 0 {
								fmt.Printf("            %s,\n", strconv.Quote(descriptionLine))
							}
						}
						fmt.Printf("        },\n")
					}
					fmt.Printf("        Req: mk%s(),\n", fixName(m.RequestFullType))
					fmt.Printf("        Res: mk%s(),\n", fixName(m.ResponseFullType))
					fmt.Printf("    }\n")
					fmt.Printf("}\n")
				}
			}
		}
	}
}

/*
MIT License

Copyright (c) 2022 David Muto (pseudomuto)

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the “Software”),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software
is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// These are data structures which are copied from
// https://github.com/pseudomuto/protoc-gen-doc

// Template is a type for encapsulating all the parsed files, messages, fields, enums, services, extensions, etc. into
// an object that will be supplied to a go template.
type Template struct {
	// The files that were parsed
	Files []*File `json:"files"`
	// Details about the scalar values and their respective types in supported languages.
	Scalars []*ScalarValue `json:"scalarValueTypes"`
}

// ScalarValue contains information about scalar value types in protobuf. The common use case for this type is to know
// which language specific type maps to the protobuf type.
//
// For example, the protobuf type `int64` maps to `long` in C#, and `Bignum` in Ruby. For the full list, take a look at
// https://developers.google.com/protocol-buffers/docs/proto3#scalar
type ScalarValue struct {
	ProtoType  string `json:"protoType"`
	Notes      string `json:"notes"`
	CppType    string `json:"cppType"`
	CSharp     string `json:"csType"`
	GoType     string `json:"goType"`
	JavaType   string `json:"javaType"`
	PhpType    string `json:"phpType"`
	PythonType string `json:"pythonType"`
	RubyType   string `json:"rubyType"`
}

// ServiceMethod contains details about an individual method within a service.
type ServiceMethod struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	RequestType       string `json:"requestType"`
	RequestLongType   string `json:"requestLongType"`
	RequestFullType   string `json:"requestFullType"`
	RequestStreaming  bool   `json:"requestStreaming"`
	ResponseType      string `json:"responseType"`
	ResponseLongType  string `json:"responseLongType"`
	ResponseFullType  string `json:"responseFullType"`
	ResponseStreaming bool   `json:"responseStreaming"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Service contains details about a service definition within a proto file.
type Service struct {
	Name        string           `json:"name"`
	LongName    string           `json:"longName"`
	FullName    string           `json:"fullName"`
	Description string           `json:"description"`
	Methods     []*ServiceMethod `json:"methods"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// EnumValue contains details about an individual value within an enumeration.
type EnumValue struct {
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// MessageExtension contains details about message-scoped extensions in proto(2) files.
type MessageExtension struct {
	FileExtension

	ScopeType     string `json:"scopeType"`
	ScopeLongType string `json:"scopeLongType"`
	ScopeFullType string `json:"scopeFullType"`
}

// Enum contains details about enumerations. These can be either top level enums, or nested (defined within a message).
type Enum struct {
	Name        string       `json:"name"`
	LongName    string       `json:"longName"`
	FullName    string       `json:"fullName"`
	Description string       `json:"description"`
	Values      []*EnumValue `json:"values"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// MessageField contains details about an individual field within a message.
//
// In the case of proto3 files, DefaultValue will always be empty. Similarly, label will be empty unless the field is
// repeated (in which case it'll be "repeated").
type MessageField struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	LongType     string `json:"longType"`
	FullType     string `json:"fullType"`
	IsMap        bool   `json:"ismap"`
	IsOneof      bool   `json:"isoneof"`
	OneofDecl    string `json:"oneofdecl"`
	DefaultValue string `json:"defaultValue"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Message contains details about a protobuf message.
//
// In the case of proto3 files, HasExtensions will always be false, and Extensions will be empty.
type Message struct {
	Name        string `json:"name"`
	LongName    string `json:"longName"`
	FullName    string `json:"fullName"`
	Description string `json:"description"`

	HasExtensions bool `json:"hasExtensions"`
	HasFields     bool `json:"hasFields"`
	HasOneofs     bool `json:"hasOneofs"`

	Extensions []*MessageExtension `json:"extensions"`
	Fields     []*MessageField     `json:"fields"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// FileExtension contains details about top-level extensions within a proto(2) file.
type FileExtension struct {
	Name               string `json:"name"`
	LongName           string `json:"longName"`
	FullName           string `json:"fullName"`
	Description        string `json:"description"`
	Label              string `json:"label"`
	Type               string `json:"type"`
	LongType           string `json:"longType"`
	FullType           string `json:"fullType"`
	Number             int    `json:"number"`
	DefaultValue       string `json:"defaultValue"`
	ContainingType     string `json:"containingType"`
	ContainingLongType string `json:"containingLongType"`
	ContainingFullType string `json:"containingFullType"`
}

// File wraps all the relevant parsed info about a proto file. File objects guarantee that their top-level enums,
// extensions, messages, and services are sorted alphabetically based on their "long name". Other values (enum values,
// fields, service methods) will be in the order that they're defined within their respective proto files.
//
// In the case of proto3 files, HasExtensions will always be false, and Extensions will be empty.
type File struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Package     string `json:"package"`

	HasEnums      bool `json:"hasEnums"`
	HasExtensions bool `json:"hasExtensions"`
	HasMessages   bool `json:"hasMessages"`
	HasServices   bool `json:"hasServices"`

	Enums      orderedEnums      `json:"enums"`
	Extensions orderedExtensions `json:"extensions"`
	Messages   orderedMessages   `json:"messages"`
	Services   orderedServices   `json:"services"`

	Options map[string]interface{} `json:"options,omitempty"`
}

type orderedEnums []*Enum
type orderedExtensions []*FileExtension
type orderedMessages []*Message
type orderedServices []*Service
