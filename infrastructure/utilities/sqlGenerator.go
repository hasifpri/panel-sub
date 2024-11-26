package infrastructureutilities

import (
	"fmt"
	"github.com/bytesaddict/dancok"
)

type SqlGenerator struct {
	TableName           string
	DefaultFieldForSort string
}

func NewSqlGenerator(tableName string, defaultFieldForSort string) *SqlGenerator {
	return &SqlGenerator{tableName, defaultFieldForSort}
}

func (g *SqlGenerator) Generate(param dancok.SelectParameter) string {
	result := "select * from (select ROW_NUMBER() OVER(" + g.ParseSort(param) + ") as RowNumber,* from \"" + g.TableName + "\" " + g.ParseFilter(param) + ") AS T"
	return result
}

func (g *SqlGenerator) Parse(param dancok.SelectParameter) string {
	result := g.ParseFilter(param) + g.ParseSort(param)

	return result
}

func (g *SqlGenerator) ParseFilter(param dancok.SelectParameter) string {
	filterText := ""
	if len(param.FilterDescriptors) > 0 {
		filterText = " WHERE "
		isFirstFilter := true
		for _, filter := range param.FilterDescriptors {
			if isFirstFilter {
				filterText = filterText + filter.FieldName
				isFirstFilter = false
			} else {
				filterText = filterText + " AND " + filter.FieldName
				//  Uncomment if need OR and AND config
				// if filter.Condition == dancok.And {
				// 	filterText = filterText + " AND " + filter.FieldName
				// } else {
				// 	filterText = filterText + " OR " + filter.FieldName
				// }
			}

			switch opt := filter.Operator; opt {
			case dancok.IsEqual:
				filterText = filterText + " = '" + filter.Value.(string) + "'"
			case dancok.IsNotEqual:
				filterText = filterText + " != '" + filter.Value.(string) + "'"
			case dancok.IsLessThan:
				filterText = filterText + " < " + filter.Value.(string)
			case dancok.IsLessThanOrEqual:
				filterText = filterText + " <= " + filter.Value.(string)
			case dancok.IsMoreThan:
				filterText = filterText + " > " + filter.Value.(string)
			case dancok.IsMoreThanOrEqual:
				filterText = filterText + " >= " + filter.Value.(string)
			case dancok.IsContain:
				filterText = filterText + " LIKE '%" + filter.Value.(string) + "%'"
			case dancok.IsBeginWith:
				filterText = filterText + " LIKE '" + filter.Value.(string) + "%'"
			case dancok.IsEndWith:
				filterText = filterText + " LIKE '%" + filter.Value.(string) + "'"
			case dancok.IsBetween:
				filterText = filterText + " BETWEEN '" + filter.Value.(string) + "' AND '" + filter.Value2.(string) + "'"
			case dancok.IsIn:
				filterText = filterText + " IN (" + ParseRangeValues(filter.RangeValues) + ")"
			case dancok.IsNotIn:
				filterText = filterText + " NOT IN (" + ParseRangeValues(filter.RangeValues) + ")"
			}
		}
	}

	if len(param.CompositeFilterDescriptors) > 0 {
		isFirstCompositeFilter := true
		for _, filter := range param.CompositeFilterDescriptors {
			if isFirstCompositeFilter {
				if filterText == "" {
					filterText = " WHERE ("
				} else {
					filterText = filterText + " " + string(filter.Condition) + " ("
				}
				isFirstCompositeFilter = false
			} else {
				filterText = filterText + " AND ("
				//  Uncomment if need OR and AND config
				// if filter.Condition == dancok.And {
				// 	filterText = filterText + " AND ("
				// } else {
				// 	filterText = filterText + " OR ("
				// }
			}

			isFirstItem := true
			for _, item := range filter.GroupFilterDescriptor.Items {
				if isFirstItem {
					switch opt := item.Operator; opt {
					case dancok.IsEqual:
						filterText = filterText + item.FieldName + " = '" + item.Value.(string) + "'"
					case dancok.IsNotEqual:
						filterText = filterText + item.FieldName + " != '" + item.Value.(string) + "'"
					case dancok.IsLessThan:
						filterText = filterText + item.FieldName + " < " + item.Value.(string)
					case dancok.IsLessThanOrEqual:
						filterText = filterText + item.FieldName + " <= " + item.Value.(string)
					case dancok.IsMoreThan:
						filterText = filterText + item.FieldName + " > " + item.Value.(string)
					case dancok.IsMoreThanOrEqual:
						filterText = filterText + item.FieldName + " >= " + item.Value.(string)
					}

					isFirstItem = false
				} else {
					switch opt := item.Operator; opt {
					case dancok.IsEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + item.FieldName + " = '" + item.Value.(string) + "'"
					case dancok.IsNotEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + item.FieldName + " != '" + item.Value.(string) + "'"
					case dancok.IsLessThan:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + item.FieldName + " < " + item.Value.(string)
					case dancok.IsLessThanOrEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + item.FieldName + " <= " + item.Value.(string)
					case dancok.IsMoreThan:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + item.FieldName + " > " + item.Value.(string)
					case dancok.IsMoreThanOrEqual:
						filterText = filterText + " " + string(filter.GroupFilterDescriptor.Condition) + " " + item.FieldName + " >= " + item.Value.(string)
					}
				}
			}

			filterText = filterText + ")"
		}
	}

	return filterText
}

func (g *SqlGenerator) ParseSort(param dancok.SelectParameter) string {
	sortText := " "

	if len(param.SortDescriptors) > 0 {
		isFirstSort := true
		sortText = sortText + "order by"
		for _, sort := range param.SortDescriptors {
			if sort.FieldName == "data_status" {
				sort.FieldName = "sys_row_status"
				sortText = sortText + " " + sort.FieldName
				isFirstSort = false
			} else if sort.FieldName == "created_at" {
				sort.FieldName = "sys_created_time"
				sortText = sortText + " " + sort.FieldName
				isFirstSort = false
			} else if sort.FieldName == "updated_at" {
				sort.FieldName = "sys_last_pending_time"
				sortText = sortText + " " + sort.FieldName
				isFirstSort = false
			} else if isFirstSort {
				sortText = sortText + " " + sort.FieldName
				isFirstSort = false
			} else {
				sortText = sortText + "," + sort.FieldName
			}

			if sort.SortDirection == dancok.Ascending {
				sortText = sortText + " asc"
			} else {
				sortText = sortText + " desc"
			}
		}
	} else {
		sortText = sortText + " order by " + g.DefaultFieldForSort + " desc"
	}

	return sortText
}

func ParseRangeValues(values []any) string {
	valueText := ""
	if len(values) > 0 {
		isFirstValue := true
		_, isStringType := values[0].(string)
		if isStringType {
			for _, v := range values {
				if isFirstValue {
					valueText = "'" + v.(string) + "'"
					isFirstValue = false
				} else {
					valueText = valueText + ",'" + v.(string) + "'"
				}
			}
		} else {
			for _, v := range values {
				if isFirstValue {
					valueText = fmt.Sprint(v.(int64))
					isFirstValue = false
				} else {
					valueText = valueText + "," + fmt.Sprint(v.(int64))
				}
			}
		}
	}
	return valueText
}
