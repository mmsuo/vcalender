package recurrence

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  RDATE
//
//   Purpose:  This property defines the list of DATE-TIME values for
//      recurring events, to-dos, journal entries, or time zone
//      definitions.
//
//   Value Type:  The default value type for this property is DATE-TIME.
//      The value type can be set to DATE or PERIOD.
//
//   Property Parameters:  IANA, non-standard, value data type, and time
//      zone identifier property parameters can be specified on this
//      property.
//
//   Conformance:  This property can be specified in recurring "VEVENT",
//      "VTODO", and "VJOURNAL" calendar components as well as in the
//      "STANDARD" and "DAYLIGHT" sub-components of the "VTIMEZONE"
//      calendar component.
//
//   Description:  This property can appear along with the "RRULE"
//      property to define an aggregate set of repeating occurrences.
//      When they both appear in a recurring component, the recurrence
//      instances are defined by the union of occurrences defined by both
//      the "RDATE" and "RRULE".
//
//      The recurrence dates, if specified, are used in computing the
//      recurrence set.  The recurrence set is the complete set of
//      recurrence instances for a calendar component.  The recurrence set
//      is generated by considering the initial "DTSTART" property along
//      with the "RRULE", "RDATE", and "EXDATE" properties contained
//      within the recurring component.  The "DTSTART" property defines
//      the first instance in the recurrence set.  The "DTSTART" property
//      value SHOULD match the pattern of the recurrence rule, if
//      specified.  The recurrence set generated with a "DTSTART" property
//      value that doesn't match the pattern of the rule is undefined.
//      The final recurrence set is generated by gathering all of the
//      start DATE-TIME values generated by any of the specified "RRULE"
//      and "RDATE" properties, and then excluding any start DATE-TIME
//      values specified by "EXDATE" properties.  This implies that start
//      DATE-TIME values specified by "EXDATE" properties take precedence
//      over those specified by inclusion properties (i.e., "RDATE" and
//      "RRULE").  Where duplicate instances are generated by the "RRULE"
//      and "RDATE" properties, only one recurrence is considered.
//      Duplicate instances are ignored.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       rdate      = "RDATE" rdtparam ":" rdtval *("," rdtval) CRLF
//
//       rdtparam   = *(
//                  ;
//                  ; The following are OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  (";" "VALUE" "=" ("DATE-TIME" / "DATE" / "PERIOD")) /
//                  (";" tzidparam) /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  (";" other-param)
//                  ;
//                  )
//
//       rdtval     = date-time / date / period
//       ;Value MUST match value type
//
//   Example:  The following are examples of this property:
//
//       RDATE:19970714T123000Z
//       RDATE;TZID=America/New_York:19970714T083000
//
//       RDATE;VALUE=PERIOD:19960403T020000Z/19960403T040000Z,
//        19960404T010000Z/PT3H
//
//       RDATE;VALUE=DATE:19970101,19970120,19970217,19970421
//        19970526,19970704,19970901,19971014,19971128,19971129,19971225

type RDate struct {
	Parameters []parameters.Parameter
	Values     []types.Value
}

func (r *RDate) Property() (string, error) {
	return properties.DefaultCreateMultiplePropertyFunc("RDATE", r.Parameters, r.Values), nil
}