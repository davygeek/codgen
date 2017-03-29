package generate

import (
	"fmt"
)

type FlagVar string

func (f *FlagVar) String() string {
	return fmt.Sprint(*f)

}

func (f *FlagVar) Set(value string) error {
	*f = FlagVar(value)
	return nil
}

var SQLDriver FlagVar
var SQLConn FlagVar
var Level FlagVar
var Tables FlagVar
var Fields FlagVar
