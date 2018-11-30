package ravendb

import "math"

const (
	MoreLikeThisOptions_DEFAULT_MAXIMUM_NUMBER_OF_TOKENS_PARSED = 5000
	MoreLikeThisOptions_DEFAULT_MINIMUM_TERM_FREQUENCY          = 2
	MoreLikeThisOptions_DEFAULT_MINIMUM_DOCUMENT_FREQUENCY      = 5
	MoreLikeThisOptions_DEFAULT_MAXIMUM_DOCUMENT_FREQUENCY      = math.MaxInt32
	MoreLikeThisOptions_DEFAULT_BOOST                           = false
	MoreLikeThisOptions_DEFAULT_BOOST_FACTOR                    = 1
	MoreLikeThisOptions_DEFAULT_MINIMUM_WORD_LENGTH             = 0
	MoreLikeThisOptions_DEFAULT_MAXIMUM_WORD_LENGTH             = 0
	MoreLikeThisOptions_DEFAULT_MAXIMUM_QUERY_TERMS             = 25
)

var (
	MoreLikeThisOptions_defaultOptions = &MoreLikeThisOptions{}
)

type MoreLikeThisOptions struct {
	MinimumTermFrequency               *int     `json:"MinimumTermFrequency"`
	MaximumQueryTerms                  *int     `json:"MaximumQueryTerms"`
	MaximumNumberOfTokensParsed        *int     `json:"MaximumNumberOfTokensParsed"`
	MinimumWordLength                  *int     `json:"MinimumWordLength"`
	MaximumWordLength                  *int     `json:"MaximumWordLength"`
	MinimumDocumentFrequency           *int     `json:"MinimumDocumentFrequency"`
	MaximumDocumentFrequency           *int     `json:"MaximumDocumentFrequency"`
	MaximumDocumentFrequencyPercentage *int     `json:"MaximumDocumentFrequencyPercentage"`
	Boost                              *bool    `json:"Boost"`
	BoostFactor                        *float32 `json:"BoostFactor"`
	StopWordsDocumentID                *string  `json:"StopWordsDocumentId"`
	Fields                             []string `json:"Fields"`
}

func NewMoreLikeThisOptions() *MoreLikeThisOptions {
	return &MoreLikeThisOptions{}
}

func (o *MoreLikeThisOptions) SetMinimumTermFrequency(minimumTermFrequency int) {
	o.MinimumTermFrequency = &minimumTermFrequency
}

func (o *MoreLikeThisOptions) SetMaximumQueryTerms(maximumQueryTerms int) {
	o.MaximumQueryTerms = &maximumQueryTerms
}

func (o *MoreLikeThisOptions) SetMaximumNumberOfTokensParsed(maximumNumberOfTokensParsed int) {
	o.MaximumNumberOfTokensParsed = &maximumNumberOfTokensParsed
}

func (o *MoreLikeThisOptions) SetMinimumWordLength(minimumWordLength int) {
	o.MinimumWordLength = &minimumWordLength
}

func (o *MoreLikeThisOptions) SetMaximumWordLength(maximumWordLength int) {
	o.MaximumWordLength = &maximumWordLength
}

func (o *MoreLikeThisOptions) SetMinimumDocumentFrequency(minimumDocumentFrequency int) {
	o.MinimumDocumentFrequency = &minimumDocumentFrequency
}

func (o *MoreLikeThisOptions) SetMaximumDocumentFrequency(maximumDocumentFrequency int) {
	o.MaximumDocumentFrequency = &maximumDocumentFrequency
}

func (o *MoreLikeThisOptions) SetMaximumDocumentFrequencyPercentage(maximumDocumentFrequencyPercentage int) {
	o.MaximumDocumentFrequencyPercentage = &maximumDocumentFrequencyPercentage
}

func (o *MoreLikeThisOptions) SetBoost(boost bool) {
	o.Boost = &boost
}

func (o *MoreLikeThisOptions) SetBoostFactor(boostFactor float32) {
	o.BoostFactor = &boostFactor
}

func (o *MoreLikeThisOptions) SetStopWordsDocumentID(stopWordsDocumentId string) {
	o.StopWordsDocumentID = &stopWordsDocumentId
}
