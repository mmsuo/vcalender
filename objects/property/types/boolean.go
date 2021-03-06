package types

import "strings"

//   V Name:  BOOLEAN
//
//   Purpose:  This value type is used to identify properties that contain
//      either a "TRUE" or "FALSE" Boolean value.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       boolean    = "TRUE" / "FALSE"
//
//   Description:  These values are case-insensitive text.  No additional
//      content value encoding (i.e., BACKSLASH character encoding, see
//      Section 3.3.11) is defined for this value type.
//
//   Example:  The following is an example of a hypothetical property that
//      has a BOOLEAN value type:
//
//       TRUE

type Boolean struct {
	V bool
}

var (
	True  = Boolean{true}
	False = Boolean{false}
)

func (b Boolean) WriteValueToStrBuilder(s *strings.Builder) error {
	if b.V {
		s.WriteString("TRUE")
	}
	s.WriteString("FALSE")
	return nil
}
