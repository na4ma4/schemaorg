package schemaorg

import (
	"strings"
)

// ItemAvailability
//
// SubClassOf: Enumeration
//
// A list of possible product availability options.
//
// https://github.com/schemaorg/schemaorg/blob/a378396bd144a084b1fc3364e0b1d7c55cb66476/data/schema.ttl#L1602
type ItemAvailability string

// Equals returns true if a ItemAvailability matches a string representation.
func (i ItemAvailability) Equals(availability string) bool {
	if strings.EqualFold(availability, string(i)) {
		return true
	}

	if strings.HasSuffix(availability, string(i[5:])) {
		return true
	}

	if idx := strings.LastIndex(string(i), "/"); idx >= 0 {
		if len(i) < idx+1 {
			return false
		}

		if strings.EqualFold(availability, string(i[idx+1:])) {
			return true
		}

		if strings.Contains(availability, " ") {
			availability = strings.ReplaceAll(availability, " ", "")

			return i.Equals(availability)
		}
	}

	return false
}

// String returns the string URI of the ItemAvailability.
func (i ItemAvailability) String() string {
	return string(i)
}

func ItemAvailabilityFromString(availability string) (ItemAvailability, bool) {
	if itemAvailabilityList == nil {
		itemAvailabilityList = []ItemAvailability{
			InStock,
			Discontinued,
			PreSale,
			LimitedAvailability,
			InStoreOnly,
			BackOrder,
			OnlineOnly,
			OutOfStock,
			PreOrder,
			SoldOut,
		}
	}

	for _, i := range itemAvailabilityList {
		if i.Equals(availability) {
			return i, true
		}
	}

	return ItemAvailability(availability), false
}

const (
	// InStock
	//
	// Indicates that the item is in stock.
	InStock ItemAvailability = "https://schema.org/InStock"

	// Discontinued
	//
	// Indicates that the item has been discontinued.
	Discontinued ItemAvailability = "https://schema.org/Discontinued"

	// PreSale
	//
	// Indicates that the item is available for ordering and delivery before general availability.
	PreSale ItemAvailability = "https://schema.org/PreSale"

	// LimitedAvailability
	//
	// Indicates that the item has limited availability.
	LimitedAvailability ItemAvailability = "https://schema.org/LimitedAvailability"

	// InStoreOnly
	//
	// Indicates that the item is available only at physical locations.
	InStoreOnly ItemAvailability = "https://schema.org/InStoreOnly"

	// BackOrder
	//
	// Indicates that the item is available on back order.
	BackOrder ItemAvailability = "https://schema.org/BackOrder"

	// OnlineOnly
	//
	// Indicates that the item is available only online.
	OnlineOnly ItemAvailability = "https://schema.org/OnlineOnly"

	// OutOfStock
	//
	// Indicates that the item is out of stock.
	OutOfStock ItemAvailability = "https://schema.org/OutOfStock"

	// PreOrder
	//
	// Indicates that the item is available for pre-order.
	PreOrder ItemAvailability = "https://schema.org/PreOrder"

	// SoldOut
	//
	// Indicates that the item has sold out.
	SoldOut ItemAvailability = "https://schema.org/SoldOut"
)

//nolint:gochecknoglobals // cache list of ItemAvailability options to reduce allocations.
var itemAvailabilityList []ItemAvailability
