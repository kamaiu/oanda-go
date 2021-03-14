package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type ObjectType struct {
	Name     string
	Comments string
	Embed    string
	Type     string
	Fields   []*Field
}

type Field struct {
	FieldName string
	Name      string
	Type      string
	ParseType string
	Required  bool
	Default   string
	Comments  []string
}

type Enum struct {
	Name     string
	Comments string
	Options  []*EnumOption
}

type EnumOption struct {
	Name        string
	Description string
}

func main() {
	fmt.Println()
	generateObjects("./model/gen/position.txt", true)
	//generateObjects("tx.txt", true)
	//generateObjects("order_request.txt", true)
}

func generateObjects(path string, firstIsBase bool) {
	objects := parseObjectTypes(string(loadFile(path)))
	tx := objects[0]
	allFields := make(map[string]*Field)
	allFieldsSort := make([]string, 0, 32)
	embeddedFields := make(map[string]*Field)
	embeddedFieldsSort := make([]string, 0, 8)

	for _, field := range tx.Fields {
		fixFieldType(field)
		parseType(field)
		if allFields[field.Name] == nil {
			allFieldsSort = append(allFieldsSort, field.Name)
		}
		allFields[field.Name] = field
		embeddedFields[field.Name] = field
		embeddedFieldsSort = append(embeddedFieldsSort, field.Name)
	}
	for i := 1; i < len(objects); i++ {
		obj := objects[i]
		if firstIsBase {
			obj.Embed = objects[0].Name
		}
		fields := make([]*Field, 0, len(obj.Fields))
		for _, field := range obj.Fields {
			if strings.ToLower(field.Name) == "type" {
				obj.Type = field.Default
			}
			fixFieldType(field)
			parseType(field)
			if embeddedFields[field.Name] == nil {
				fields = append(fields, field)
			}

		}
		obj.Fields = fields
	}

	for _, obj := range objects {
		fmt.Println(generateStruct(obj))
	}

	for _, obj := range objects {
		for _, field := range obj.Fields {
			parseType(field)
			if existing := allFields[field.Name]; existing != nil {
				if existing.ParseType != field.ParseType {
					panic(fmt.Sprintf("mismatched fields: %s  <>  %s", existing.ParseType, field.ParseType))
				} else {
					//fmt.Println(existing.Name)
				}
			} else {
				allFieldsSort = append(allFieldsSort, field.Name)
			}
			allFields[field.Name] = field
		}
	}

	sort.Strings(embeddedFieldsSort)
	sort.Strings(allFieldsSort)
	parserFields := make([]*Field, 0, len(allFields))
	parserFieldMap := make(map[string]*Field)
	for _, fieldName := range allFieldsSort {
		field := allFields[fieldName]
		f := &Field{
			FieldName: field.FieldName,
			Name:      field.Name,
			Type:      field.ParseType,
			ParseType: field.ParseType,
			Comments:  nil,
		}
		parserFields = append(parserFields, f)
		parserFieldMap[field.Name] = field
	}
	parserStruct := &ObjectType{
		Name:     fmt.Sprintf("%sParser", objects[0].Name),
		Comments: "",
		Embed:    "",
		Fields:   parserFields,
	}
	fmt.Println(generateStruct(parserStruct))
	fmt.Println()

	b := &strings.Builder{}

	b.WriteString("// Example\n")
	b.WriteString("/*\n")
	b.WriteString("r := parser.Parse()\n")
	b.WriteString("switch v := r.(type) {\n")

	for i, obj := range objects {
		if i > 0 {
			b.WriteString(fmt.Sprintf("case *%s:\n", obj.Name))
		}
	}
	b.WriteString("}\n")
	b.WriteString("*/\n")

	b.WriteString(fmt.Sprintf("func (p *%s) Parse() interface{} {\n", parserStruct.Name))
	b.WriteString("    switch p.Type {\n")
	for i, obj := range objects {
		_ = i
		if i > 0 {
			b.WriteString(fmt.Sprintf("    case \"%s\":\n", obj.Type))
			b.WriteString(fmt.Sprintf("        return &%s{\n", obj.Name))
			b.WriteString(fmt.Sprintf("            %s: %s{\n", objects[0].Name, objects[0].Name))
			for _, field := range objects[0].Fields {
				pf := parserFieldMap[field.Name]
				if pf == nil {
					panic(fmt.Sprintf("could not find field: %s", field.Name))
				}
				if field.Type != field.ParseType {
					b.WriteString(fmt.Sprintf("                %s: %s(p.%s),\n", field.FieldName, field.Type, field.FieldName))
				} else {
					b.WriteString(fmt.Sprintf("                %s: p.%s,\n", field.FieldName, field.FieldName))
				}
			}
			b.WriteString(fmt.Sprintf("            },\n"))
			for _, field := range obj.Fields {
				pf := parserFieldMap[field.Name]
				if pf == nil {
					panic(fmt.Sprintf("could not find field: %s", field.Name))
				}
				if field.Type != field.ParseType {
					b.WriteString(fmt.Sprintf("            %s: %s(p.%s),\n", field.FieldName, field.Type, field.FieldName))
				} else {
					b.WriteString(fmt.Sprintf("            %s: p.%s,\n", field.FieldName, field.FieldName))
				}
			}
			b.WriteString("        }\n")
		}
	}
	b.WriteString("    }\n")
	b.WriteString("    return p\n")
	b.WriteString("}\n\n")

	fmt.Println(b.String())
}

func endsWith(s string, with string) bool {
	if len(with) > len(s) {
		return false
	}
	return s[len(s)-len(with):] == with
}

func parseType(field *Field) {
	if endsWith(field.Type, "Type") ||
		endsWith(field.Type, "Reason") ||
		(endsWith(field.Type, "ID") && field.Type[0:2] != "[]") ||
		endsWith(field.Type, "Condition") ||
		field.Type == "string" {
		field.ParseType = "string"
	} else {
		//switch field.Type {
		//case "DecimalNumber", "AccountUnits", "TimeInForce", "OrderPositionFill", "Currency":
		//	field.ParseType = "string"
		//}
		field.ParseType = field.Type
	}
}

func fixFieldType(field *Field) {
	typ := field.Type
	switch typ {
	case "AccountUnits", "TransactionType", "InstrumentName", "PriceValue", "TradeID", "ClientID", "DecimalNumber", "TimeInForce":

	case "boolean":
		typ = "bool"
	case "integer":
		typ = "int64"
	case "MarketOrderTradeClose", "MarketOrderPositionCloseout",
		"MarketOrderMarginCloseout",
		"TradeOpen",
		"TradeReduce",
		"TakeProfitDetails",
		"StopLossDetails",
		"TrailingStopLossDetails",
		"GuaranteedStopLossDetails",
		"MarketOrderDelayedTradeClose",
		"HomeConversionFactors",
		"ClientExtensions":
		typ = "*" + typ
	case "DateTime":
		typ = "*time.Time"
	case "Array[TradeID]":
		typ = "[]TradeID"
	case "Array[TradeReduce]":
		typ = "[]*TradeReduce"
	case "Array[PositionFinancing]":
		typ = "[]*PositionFinancing"
	case "Array[OpenTradeDividendAdjustment]":
		typ = "[]*OpenTradeDividendAdjustment"
	}
	field.Type = typ
}

func generateEnum(name, comments string, input string) string {
	b := &strings.Builder{}

	writeComments(comments, func(line string) {
		b.WriteString(fmt.Sprintf("// %s\n", line))
	})
	b.WriteString(fmt.Sprintf("type %s string\n", name))
	values := parseEnumOptions(input)
	if len(values) == 0 {
		return b.String()
	}
	b.WriteString("const (\n")
	for _, v := range values {
		if len(v.Description) > 0 {
			writeComments(v.Description, func(line string) {
				b.WriteString(fmt.Sprintf("    // %s\n", line))
			})
		}
		b.WriteString(fmt.Sprintf("    %s_%s %s = \"%s\"\n", name, v.Name, name, v.Name))
	}
	b.WriteString(")\n")
	return b.String()
}

func generateStruct(obj *ObjectType) string {
	b := &strings.Builder{}

	writeComments(obj.Comments, func(line string) {
		b.WriteString(fmt.Sprintf("// %s\n", line))
	})
	b.WriteString(fmt.Sprintf("type %s struct {\n", obj.Name))
	if len(obj.Embed) > 0 {
		b.WriteString(fmt.Sprintf("    %s\n", obj.Embed))
	}

	if len(obj.Fields) > 0 {
		for _, field := range obj.Fields {
			if len(field.Comments) > 0 {
				writeComments(strings.Join(field.Comments, " "), func(line string) {
					b.WriteString(fmt.Sprintf("    // %s\n", strings.TrimSpace(line)))
				})
			}

			b.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", field.FieldName, field.Type, field.Name))
		}
	}
	b.WriteString("}\n")
	return b.String()
}

func writeComments(comments string, fn func(line string)) {
	if len(comments) == 0 {
		return
	}
	lines := wordWrap(comments, 80)
	for _, line := range lines {
		if fn != nil {
			fn(line)
		}
	}
}

func parseEnumOptions(input string) []*EnumOption {
	lines := strings.Split(input, "\n")
	out := make([]*EnumOption, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
	loop:
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case ' ', '\t', '\r':
				out = append(out, &EnumOption{line[0:i], strings.TrimSpace(line[i+1:])})
				break loop
			}
		}
	}
	return out
}

func parseObjectFields(input string) []*Field {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	fields := make([]*Field, 0, len(lines))
	var comments []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			comments = append(comments, strings.TrimSpace(line[1:]))
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic("expected 2 parts for field declaration: " + line)
		}

		field := &Field{
			Name:     strings.TrimSpace(parts[0]),
			Comments: comments,
		}
		comments = nil

		field.FieldName = field.Name
		if len(field.Name) > 0 {
			field.FieldName = strings.ToUpper(field.Name[0:1]) + field.Name[1:]
		}

		typePart := strings.TrimSpace(parts[1])
		if len(typePart) < 3 {
			panic("expected type name: " + line)
		}
		if typePart[0] == '(' {
			typePart = typePart[1:]
		}
		if typePart[len(typePart)-1] == ',' {
			typePart = typePart[0 : len(typePart)-1]
		}
		if typePart[len(typePart)-1] == ')' {
			typePart = typePart[0 : len(typePart)-1]
		}
		typePart = strings.TrimSpace(typePart)
		typeParts := strings.Split(typePart, ",")

		field.Type = strings.TrimSpace(typeParts[0])

		if len(typeParts) > 0 {
			for i := 1; i < len(typeParts); i++ {
				p := strings.TrimSpace(typeParts[i])

				if p == "required" {
					field.Required = true
				} else {
					index := strings.Index(p, "default=")
					if index > -1 {
						field.Default = strings.TrimSpace(p[index+len("default="):])
					}
				}
			}
		}
		fields = append(fields, field)
		comments = nil
	}
	return fields
}

func parseObjectTypes(input string) []*ObjectType {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	type stateCode int
	const (
		StateBegin stateCode = iota
		StateOpenCurly
		StateCloseCurly
	)

	state := StateBegin
	ret := make([]*ObjectType, 0, 16)
	obj := &ObjectType{}
	mark := 0
	count := 0
	for _, line := range lines {
		lineSize := len(line) + 1
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			count += lineSize
			continue
		}
		switch state {
		case StateBegin:
			obj.Name, obj.Comments = parseNameAndComment(line)
			state = StateOpenCurly

		case StateOpenCurly:
			if line[0] != '{' {
				panic("expected {")
			}
			mark = count + lineSize + 1
			state = StateCloseCurly

		case StateCloseCurly:
			if line[0] == '}' {
				text := strings.TrimSpace(input[mark:count])

				if len(text) > 0 {
					obj.Fields = parseObjectFields(strings.TrimSpace(text))
				}
				ret = append(ret, obj)
				obj = &ObjectType{}
				state = StateBegin
			}
		}

		count += lineSize
	}
	return ret
}

func parseNameAndComment(line string) (string, string) {
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case ' ', '\t', '\r':
			return line[0:i], strings.TrimSpace(line[i+1:])
		}
	}
	return line, ""
}

func wordWrap(line string, maxLine int) []string {
	var lines []string
	c := 0
	mark := 0
	for i := 0; i < len(line); i++ {
		if c >= maxLine {
			switch line[i] {
			case ' ', '\t', '\r':
				lines = append(lines, line[mark:i])
				mark = i + 1
				c = 0
			default:
				c++
			}
		} else {
			c++
		}
	}

	end := strings.TrimSpace(line[mark:])
	if len(end) > 0 {
		lines = append(lines, end)
	}
	return lines
}

func loadFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}
