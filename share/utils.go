package chshare

import (
	"errors"
	"fmt"
	"strings"
)

type StringList []string

func (s *StringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *StringList) Set(inputString string) error {
	*s = strings.Split(inputString, ",")
	if *s == nil {
		return errors.New(fmt.Sprintf("error while splitting: %s. Separator: , ", inputString))
	}
	return nil
}

func (s *StringList) Cast() {

}