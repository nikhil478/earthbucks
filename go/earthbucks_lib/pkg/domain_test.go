package earthbucks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {
    t.Run("isValidDomain", func(t *testing.T) {
        tests := []struct {
            domainString string
            expected     bool
        }{
            {"earthbucks.com", true},
            {"earth-bucks.com", false},
            {"earthbucks.com.", false},
            {".earthbucks.com", false},
            {"node.node.node.node.earthbucks.com", true},
            {"node.node.node.node.node.node.node.node.node.earthbucks.com", false},
        }

        for _, test := range tests {
            t.Run(test.domainString, func(t *testing.T) {
                domain := DomainFromString(test.domainString)
                assert.Equal(t, test.expected, domain.IsValid())
            })
        }
    })
}