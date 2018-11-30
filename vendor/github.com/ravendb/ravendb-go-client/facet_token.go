package ravendb

import "strings"

var _ queryToken = &facetToken{}

type facetToken struct {
	_facetSetupDocumentId string
	_aggregateByFieldName string
	_alias                string
	_ranges               []string
	_optionsParameterName string

	_aggregations []*facetAggregationToken
}

func (t *facetToken) GetName() string {
	return firstNonEmptyString(t._alias, t._aggregateByFieldName)
}

func NewFacetTokenWithID(facetSetupDocumentId string) *facetToken {
	return &facetToken{
		_facetSetupDocumentId: facetSetupDocumentId,
	}
}

func NewFacetTokenAll(aggregateByFieldName string, alias string, ranges []string, optionsParameterName string) *facetToken {
	return &facetToken{
		_aggregateByFieldName: aggregateByFieldName,
		_alias:                alias,
		_ranges:               ranges,
		_optionsParameterName: optionsParameterName,
	}
}

func (t *facetToken) writeTo(writer *strings.Builder) {
	writer.WriteString("facet(")

	if t._facetSetupDocumentId != "" {
		writer.WriteString("id('")
		writer.WriteString(t._facetSetupDocumentId)
		writer.WriteString("'))")

		return
	}

	firstArgument := false

	if t._aggregateByFieldName != "" {
		writer.WriteString(t._aggregateByFieldName)
	} else if len(t._ranges) != 0 {
		firstInRange := true

		for _, rang := range t._ranges {
			if !firstInRange {
				writer.WriteString(", ")
			}

			firstInRange = false
			writer.WriteString(rang)
		}
	} else {
		firstArgument = true
	}

	for _, aggregation := range t._aggregations {
		if !firstArgument {
			writer.WriteString(", ")
		}
		firstArgument = false
		aggregation.writeTo(writer)
	}

	if stringIsNotBlank(t._optionsParameterName) {
		writer.WriteString(", $")
		writer.WriteString(t._optionsParameterName)
	}

	writer.WriteString(")")

	if stringIsBlank(t._alias) || t._alias == t._aggregateByFieldName {
		return
	}

	writer.WriteString(" as ")
	writer.WriteString(t._alias)
}

func createFacetToken(facetSetupDocumentId string) *facetToken {
	if stringIsWhitespace(facetSetupDocumentId) {
		//throw new IllegalArgumentException("facetSetupDocumentId cannot be null");
		panicIf(true, "facetSetupDocumentId cannot be null")
	}

	return NewFacetTokenWithID(facetSetupDocumentId)
}

func createFacetTokenWithFacet(facet *Facet, addQueryParameter func(Object) string) *facetToken {
	optionsParameterName := getOptionsParameterName(facet, addQueryParameter)
	token := NewFacetTokenAll(facet.GetFieldName(), facet.GetDisplayFieldName(), nil, optionsParameterName)

	applyAggregations(facet, token)
	return token
}

func createFacetTokenWithRangeFacet(facet *RangeFacet, addQueryParameter func(Object) string) *facetToken {
	optionsParameterName := getOptionsParameterName(facet, addQueryParameter)

	token := NewFacetTokenAll("", facet.GetDisplayFieldName(), facet.getRanges(), optionsParameterName)

	applyAggregations(facet, token)

	return token
}

func FacetToken_createWithGenericRangeFacet(facet *GenericRangeFacet, addQueryParameter func(Object) string) *facetToken {
	optionsParameterName := getOptionsParameterName(facet, addQueryParameter)

	var ranges []string
	for _, rangeBuilder := range facet.getRanges() {
		ranges = append(ranges, GenericRangeFacet_parse(rangeBuilder, addQueryParameter))
	}

	token := NewFacetTokenAll("", facet.GetDisplayFieldName(), ranges, optionsParameterName)

	applyAggregations(facet, token)
	return token
}

func FacetToken_createWithFacetBase(facet FacetBase, addQueryParameter func(Object) string) *facetToken {
	// this is just a dispatcher
	return facet.ToFacetToken(addQueryParameter)
}

func applyAggregations(facet FacetBase, token *facetToken) {
	m := facet.GetAggregations()

	for key, value := range m {
		var aggregationToken *facetAggregationToken
		switch key {
		case FacetAggregation_MAX:
			aggregationToken = facetAggregationTokenMax(value)
		case FacetAggregation_MIN:
			aggregationToken = facetAggregationTokenMin(value)
		case FacetAggregation_AVERAGE:
			aggregationToken = facetAggregationTokenAverage(value)
		case FacetAggregation_SUM:
			aggregationToken = facetAggregationTokenSum(value)
		default:
			panic("Unsupported aggregation method: " + key)
			//throw new NotImplementedException("Unsupported aggregation method: " + aggregation.getKey());
		}

		token._aggregations = append(token._aggregations, aggregationToken)
	}
}

func getOptionsParameterName(facet FacetBase, addQueryParameter func(Object) string) string {
	if facet.GetOptions() == nil || facet.GetOptions() == FacetOptions_getDefaultOptions() {
		return ""
	}
	return addQueryParameter(facet.GetOptions())
}

var _ queryToken = &facetAggregationToken{}

type facetAggregationToken struct {
	_fieldName   string
	_aggregation FacetAggregation
}

func newFacetAggregationToken(fieldName string, aggregation FacetAggregation) *facetAggregationToken {
	return &facetAggregationToken{
		_fieldName:   fieldName,
		_aggregation: aggregation,
	}
}

func (t *facetAggregationToken) writeTo(writer *strings.Builder) {
	switch t._aggregation {
	case FacetAggregation_MAX:
		writer.WriteString("max(")
		writer.WriteString(t._fieldName)
		writer.WriteString(")")
	case FacetAggregation_MIN:
		writer.WriteString("min(")
		writer.WriteString(t._fieldName)
		writer.WriteString(")")
	case FacetAggregation_AVERAGE:
		writer.WriteString("avg(")
		writer.WriteString(t._fieldName)
		writer.WriteString(")")
	case FacetAggregation_SUM:
		writer.WriteString("sum(")
		writer.WriteString(t._fieldName)
		writer.WriteString(")")
	default:
		panicIf(true, "Invalid aggregation mode: %s", t._aggregation)
		//throw new IllegalArgumentException("Invalid aggregation mode: " + _aggregation);
	}
}

func facetAggregationTokenMax(fieldName string) *facetAggregationToken {
	panicIf(stringIsWhitespace(fieldName), "FieldName can not be null")
	return newFacetAggregationToken(fieldName, FacetAggregation_MAX)
}

func facetAggregationTokenMin(fieldName string) *facetAggregationToken {
	panicIf(stringIsWhitespace(fieldName), "FieldName can not be null")
	return newFacetAggregationToken(fieldName, FacetAggregation_MIN)
}

func facetAggregationTokenAverage(fieldName string) *facetAggregationToken {
	panicIf(stringIsWhitespace(fieldName), "FieldName can not be null")
	return newFacetAggregationToken(fieldName, FacetAggregation_AVERAGE)
}

func facetAggregationTokenSum(fieldName string) *facetAggregationToken {
	panicIf(stringIsWhitespace(fieldName), "FieldName can not be null")
	return newFacetAggregationToken(fieldName, FacetAggregation_SUM)
}
