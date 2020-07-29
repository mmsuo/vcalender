package relationship

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
	"strings"
)

//   Property Name:  UID
//
//   Purpose:  This property defines the persistent, globally unique
//      identifier for the calendar component.
//
//   V Type:  TEXT
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  The property MUST be specified in the "VEVENT",
//      "VTODO", "VJOURNAL", or "VFREEBUSY" calendar components.
//
//   Description:  The "UID" itself MUST be a globally unique identifier.
//      The generator of the identifier MUST guarantee that the identifier
//      is unique.  There are several algorithms that can be used to
//      accomplish this.  A good method to assure uniqueness is to put the
//      domain name or a domain literal IP address of the host on which
//      the identifier was created on the right-hand side of an "@", and
//      on the left-hand side, put a combination of the current calendar
//      date and time of day (i.e., formatted in as a DATE-TIME value)
//      along with some other currently unique (perhaps sequential)
//      identifier available on the system (for example, a process id
//      number).  Using a DATE-TIME value on the left-hand side and a
//      domain name or domain literal on the right-hand side makes it
//      possible to guarantee uniqueness since no two hosts should be
//      using the same domain name or IP address at the same time.  Though
//      other algorithms will work, it is RECOMMENDED that the right-hand
//      side contain some domain identifier (either of the host itself or
//      otherwise) such that the generator of the message identifier can
//      guarantee the uniqueness of the left-hand side within the scope of
//      that domain.
//
//      This is the method for correlating scheduling messages with the
//      referenced "VEVENT", "VTODO", or "VJOURNAL" calendar component.
//
//      The full range of calendar components specified by a recurrence
//      set is referenced by referring to just the "UID" property value
//      corresponding to the calendar component.  The "RECURRENCE-ID"
//      property allows the reference to an individual instance within the
//      recurrence set.
//
//      This property is an important method for group-scheduling
//      applications to match requests with later replies, modifications,
//      or deletion requests.  Calendaring and scheduling applications
//      MUST generate this property in "VEVENT", "VTODO", and "VJOURNAL"
//      calendar components to assure interoperability with other group-
//      scheduling applications.  This identifier is created by the
//      calendar system that generates an iCalendar object.
//
//      Implementations MUST be able to receive and persist values of at
//      least 255 octets for this property, but they MUST NOT truncate
//      values in the middle of a UTF-8 multi-octet sequence.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       uid        = "UID" uidparam ":" text CRLF
//
//       uidparam   = *(";" other-param)
//
//   Example:  The following is an example of this property:
//
//       UID:19960401T080045Z-4000F192713-0052@example.com

type Uid struct {
	Parameters []parameters.Parameter
	Value      *types.Text
}

func (u *Uid) WritePropertyToStrBuilder(s *strings.Builder) error {
	return properties.DefaultCreatePropertyFunc("UID", u.Parameters, u.Value, s)
}

func NewUid(u string) *Uid {
	return &Uid{
		Value: &types.Text{V: u},
	}
}
