package ravendb

type CircleCriteria struct {
	SpatialCriteriaCommon
	_radius      float64
	_latitude    float64
	_longitude   float64
	_radiusUnits SpatialUnits
}

func NewCircleCriteria(radius float64, latitude float64, longitude float64, radiusUnits SpatialUnits, relation SpatialRelation, distErrorPercent float64) *CircleCriteria {

	res := &CircleCriteria{
		_radius:      radius,
		_latitude:    latitude,
		_longitude:   longitude,
		_radiusUnits: radiusUnits,
	}
	res._relation = relation
	res._distanceErrorPct = distErrorPercent
	return res
}

func (c *CircleCriteria) ToQueryToken(fieldName string, addQueryParameter func(Object) string) queryToken {
	return c.SpatialCriteriaCommon.toQueryTokenCommon(c, fieldName, addQueryParameter)
}

func (c *CircleCriteria) GetShapeToken(addQueryParameter func(Object) string) *shapeToken {
	return ShapeToken_circle(addQueryParameter(c._radius), addQueryParameter(c._latitude),
		addQueryParameter(c._longitude), c._radiusUnits)
}
