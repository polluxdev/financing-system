package helper

import (
	"fmt"
	"strings"
)

type SelectBuilder struct {
	Table    string
	Fields   string
	Function string
	Alias    string
}

type ConditionalBuilder struct {
	Table         string
	Column        string
	Value         interface{}
	FunctionValue string
	Logical       string
	Operator      string
}

type JoinBuilder struct {
	Type           string
	Table          string
	ForeignKey     string
	ReferenceTable string
	ReferenceKey   string
}

type GroupByBuilder struct {
	Table    string
	Fields   string
	Function string
}

func ConstructSelectClause(builder []SelectBuilder) []string {
	var result []string

	for _, v := range builder {
		if v.Fields == "" {
			return nil
		}

		switch {
		case v.Table != "" && v.Function != "" && v.Alias != "":
			result = append(result, fmt.Sprintf("%s(%s.%s) AS %s", v.Function, v.Table, v.Fields, v.Alias))
		case v.Table != "" && v.Function != "" && v.Alias == "":
			result = append(result, fmt.Sprintf("%s(%s.%s)", v.Function, v.Table, v.Fields))
		case v.Table != "" && v.Function == "" && v.Alias != "":
			result = append(result, fmt.Sprintf("%s.%s AS %s", v.Table, v.Fields, v.Alias))
		case v.Table != "" && v.Function == "" && v.Alias == "":
			result = append(result, fmt.Sprintf("%s.%s", v.Table, v.Fields))
		case v.Table == "" && v.Function != "" && v.Alias != "":
			result = append(result, fmt.Sprintf("%s(%s) AS %s", v.Function, v.Fields, v.Alias))
		case v.Table == "" && v.Function != "" && v.Alias == "":
			result = append(result, fmt.Sprintf("%s(%s)", v.Function, v.Fields))
		case v.Table == "" && v.Function == "" && v.Alias != "":
			result = append(result, fmt.Sprintf("%s AS %s", v.Fields, v.Alias))
		default:
			result = append(result, v.Fields)
		}
	}

	return result
}

func ConstructConditionalClause(builder []ConditionalBuilder) (string, []interface{}) {
	var (
		clauses       []string
		args          []interface{}
		defaultClause = "1 = 1 "
	)

	for _, v := range builder {
		if v.Column == "" || v.Logical == "" || v.Operator == "" {
			return defaultClause, nil
		}

		if v.FunctionValue == "" {
			v.FunctionValue = "?"
		}

		column := v.Column
		if v.Table != "" {
			column = fmt.Sprintf("%s.%s", v.Table, v.Column)
		}

		switch v.Logical {
		case "LIKE":
			clauses = append(clauses, fmt.Sprintf("%s LOWER(%s) %s %s", v.Operator, column, v.Logical, v.FunctionValue))
			args = append(args, fmt.Sprintf("%%%s%%", v.Value))
		case "IN":
			clauses = append(clauses, fmt.Sprintf("%s %s %s (%s)", v.Operator, column, v.Logical, v.FunctionValue))
			args = append(args, v.Value)
		case "BETWEEN":
			clauses = append(clauses, fmt.Sprintf("%s %s %s %s AND %s", v.Operator, column, v.Logical, v.FunctionValue, v.FunctionValue))
			slice, ok := v.Value.([]interface{})
			if !ok {
				return defaultClause, nil
			}

			args = append(args, slice...)
		case "JSON_OVERLAPS":
			clauses = append(clauses, fmt.Sprintf("%s %s(%s, %s)", v.Operator, v.Logical, column, v.FunctionValue))
			args = append(args, v.Value)
		default:
			if v.Value != nil {
				clauses = append(clauses, fmt.Sprintf("%s %s %s %s", v.Operator, column, v.Logical, v.FunctionValue))
				args = append(args, v.Value)
			} else {
				clauses = append(clauses, fmt.Sprintf("%s %s %s NULL", v.Operator, column, v.Logical))
			}
		}
	}

	return defaultClause + strings.Join(clauses, " "), args
}

func ConstructJoinClause(builder []JoinBuilder) []string {
	var result []string

	for _, v := range builder {
		if v.Type == "" || v.Table == "" || v.ForeignKey == "" || v.ReferenceTable == "" || v.ReferenceKey == "" {
			return nil
		}

		result = append(result, fmt.Sprintf("%s JOIN %s ON %s.%s = %s.%s", v.Type, v.Table, v.Table, v.ForeignKey, v.ReferenceTable, v.ReferenceKey))
	}

	return result
}

func ConstructGroupByClause(builder []GroupByBuilder) string {
	var clauses []string

	for _, v := range builder {
		if v.Fields == "" {
			return ""
		}

		switch {
		case v.Table != "" && v.Function != "":
			clauses = append(clauses, fmt.Sprintf("%s.%s %s", v.Table, v.Fields, v.Function))
		case v.Table != "" && v.Function == "":
			clauses = append(clauses, fmt.Sprintf("%s.%s", v.Table, v.Fields))
		case v.Table == "" && v.Function != "":
			clauses = append(clauses, fmt.Sprintf("%s %s", v.Fields, v.Function))
		default:
			clauses = append(clauses, v.Fields)
		}
	}

	return strings.Join(clauses, ", ")
}
