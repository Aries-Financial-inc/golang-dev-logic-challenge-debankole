package model

import (
	"errors"
	"time"
)

// OptionsContract structure for the request body
type OptionsContract struct {
	Type           OptionType // "call" or "put"
	StrikePrice    float64    // the strike price of the option
	Bid            float64    // bid price of the option
	Ask            float64    // ask price of the option
	ExpirationDate time.Time  // expiration date of the option
	LongShort      LongShort  // "long" or "short"
}

// OptionType type for the option type
type OptionType string

const (
	// CallOptionType is the call option type
	CallOptionType OptionType = "call"
	// PutOptionType is the put option type
	PutOptionType OptionType = "put"
)

// LongShort type for the long or short position of the option
type LongShort string

const (
	// Long is the long position
	Long LongShort = "long"
	// Short is the short position
	Short LongShort = "short"
)

// ErrOptionContractValidation is an error indicating an invalid option type.
var ErrOptionContractValidation = errors.New("invalid option contract")
