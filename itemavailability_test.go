package schemaorg_test

import (
	"testing"

	"github.com/na4ma4/schemaorg"
)

func TestItemAvailabilityCompare(t *testing.T) {
	t.Parallel()

	itemList := []string{
		"instock",
		"INSTOCK",
		"InStock",
		"in stock",
		" in stock  ",
		"https://schema.org/InStock",
		"http://schema.org/InStock",
	}

	for _, e := range itemList {
		if !schemaorg.InStock.Equals(e) {
			t.Errorf("'%s' does not match %s", e, schemaorg.InStock)
		}
	}

	if schemaorg.OnlineOnly.Equals("only") {
		t.Errorf("'only' should not match %s", schemaorg.OnlineOnly)
	}

	if schemaorg.OnlineOnly.Equals("InStock") {
		t.Errorf("'InStock' should not match %s", schemaorg.OnlineOnly)
	}
}

func TestItemAvailabilityFromString(t *testing.T) {
	t.Parallel()

	itemList := map[string]schemaorg.ItemAvailability{
		"instock":             schemaorg.InStock,
		"Discontinued":        schemaorg.Discontinued,
		"PreSale":             schemaorg.PreSale,
		"LimitedAvailability": schemaorg.LimitedAvailability,
		"InStoreOnly":         schemaorg.InStoreOnly,
		"BackOrder":           schemaorg.BackOrder,
		"OnlineOnly":          schemaorg.OnlineOnly,
		"OutOfStock":          schemaorg.OutOfStock,
		"PreOrder":            schemaorg.PreOrder,
		"SoldOut":             schemaorg.SoldOut,
	}

	for key, value := range itemList {
		item, ok := schemaorg.ItemAvailabilityFromString(key)
		if !ok {
			t.Errorf("'%s' was not found", key)
		}

		if item != value {
			t.Errorf("'%s' returned '%s' instead of '%s'", key, item, value)
		}

		t.Logf("'%s' returned '%s'", key, item)
	}
}
