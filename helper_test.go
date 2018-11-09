package govcloudair

import "testing"
import "github.com/stretchr/testify/assert"

func TestExtractID(t *testing.T) {
	assert.EqualValues(t, "", ExtractID(""))
	assert.EqualValues(t, "39867ab4-04e0-4b13-b468-08abcc1de810", ExtractID("39867ab4-04e0-4b13-b468-08abcc1de810"))
	assert.EqualValues(t, "39867ab4-04e0-4b13-b468-08abcc1de810", ExtractID("urn:vcloud:catalog:39867ab4-04e0-4b13-b468-08abcc1de810"))
}
