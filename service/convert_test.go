package service_test

import (
	"errors"
	"sorawat-convert-currency-suffix/service"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	type given struct {
		inputs []decimal.Decimal
	}

	type expect struct {
		response []string
		err      error
	}

	type testCase struct {
		name   string
		given  given
		expect expect
	}

	testCases := []testCase{
		{
			name: "1,234",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(1234)},
			},
			expect: expect{
				response: []string{"หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
				err:      nil,
			},
		},
		{
			name: "1000.00",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(1000.00)},
			},
			expect: expect{
				response: []string{"หนึ่งพันบาทถ้วน"},
				err:      nil,
			},
		},
		{
			name: "33,333.75",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(33333.75)},
			},
			expect: expect{
				response: []string{"สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},
				err:      nil,
			},
		},
		{
			name: "321,132,521.2",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(321132521.2)},
			},
			expect: expect{
				response: []string{"สามร้อยยี่สิบเอ็ดล้านหนึ่งแสนสามหมื่นสองพันห้าร้อยยี่สิบเอ็ดบาทยี่สิบสตางค์"},
				err:      nil,
			},
		},
		{
			name: "1,100,000",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(1100000)},
			},
			expect: expect{
				response: []string{"หนึ่งล้านหนึ่งแสนบาทถ้วน"},
				err:      nil,
			},
		},
		{
			name: "12,345,678.99",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(12345678.99)},
			},
			expect: expect{
				response: []string{"สิบสองล้านสามแสนสี่หมื่นห้าพันหกร้อยเจ็ดสิบแปดบาทเก้าสิบเก้าสตางค์"},
				err:      nil,
			},
		},
		{
			name: "11.05",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(11.05)},
			},
			expect: expect{
				response: []string{"สิบเอ็ดบาทห้าสตางค์"},
				err:      nil,
			},
		},
		{
			name: "1,000,001",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(1000001)},
			},
			expect: expect{
				response: []string{"หนึ่งล้านหนึ่งบาทถ้วน"},
				err:      nil,
			},
		},
		{
			name: "2,000,011",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(2000011)},
			},
			expect: expect{
				response: []string{"สองล้านสิบเอ็ดบาทถ้วน"},
				err:      nil,
			},
		},
		{
			name: "return empty",
			given: given{
				inputs: []decimal.Decimal{},
			},
			expect: expect{
				response: []string{},
				err:      nil,
			},
		},
		{
			name: "return error",
			given: given{
				inputs: []decimal.Decimal{decimal.NewFromFloat(11.255)},
			},
			expect: expect{
				response: nil,
				err:      errors.New("number incorrect"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := service.ConvertCurrency(tc.given.inputs)
			if tc.expect.err != nil {
				assert.Equal(t, tc.expect.err, err)
				return
			}

			assert.Equal(t, tc.expect.response, response)
		})
	}
}
