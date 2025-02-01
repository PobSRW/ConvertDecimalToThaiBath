package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func ConvertCurrency(inputs []decimal.Decimal) ([]string, error) {
	result := []string{}

	if len(inputs) == 0 {
		return result, nil
	}

	for _, v := range inputs {
		s := v.String()

		var isFractional bool
		if strings.Contains(s, DOT) {
			isFractional = true
		}

		idxTotal := len(s) - 1

		totalFractional := getTotalFractional(isFractional, idxTotal, s)
		if totalFractional > 2 {
			return nil, errors.New("number incorrect")
		}

		if isFractional {
			idxTotal = (idxTotal - 1) - totalFractional
		}

		var previosNum string
		var idx int
		var data = []string{}

		for i, r := range s {
			numStr := string(r)

			if numStr == DOT {
				idxTotal = 1
				data = append(data, BATH)
				idx = 0

				continue
			}

			var isInitNum bool
			if i == 0 {
				isInitNum = true
			}

			var isPreviosZero bool
			if previosNum == ZERO && !isInitNum {
				isPreviosZero = true
			}

			numTh := mapThaiNumber(numStr, idx, idxTotal, isInitNum, isPreviosZero)
			unit := mapThaiUnit(idxTotal-idx, numStr)

			data = append(data, fmt.Sprintf("%s%s", numTh, unit))

			previosNum = numStr
			idx++
		}

		data = append(data, addSuffix(isFractional))
		result = append(result, strings.Join(data, ""))
	}

	return result, nil
}

func mapThaiNumber(
	v string,
	idx, idxTotal int,
	isInitNum bool,
	isPreviosZero bool,
) string {
	index := idxTotal - idx

	idxTotal = idxTotal % 6
	miliionIndex := idxTotal - idx

	if v == ONE {
		if !isPreviosZero && !isInitNum && (miliionIndex == 0 || index == 0) {
			return "เอ็ด"
		}

		if miliionIndex == 1 || index == 1 {
			return ""
		}
	}

	if v == TWO && (miliionIndex == 1 || index == 1) {
		return "ยี่"
	}

	thaiNumberMap := map[string]string{
		ONE:   "หนึ่ง",
		TWO:   "สอง",
		THREE: "สาม",
		FOUR:  "สี่",
		FIVE:  "ห้า",
		SIX:   "หก",
		SEVEN: "เจ็ด",
		EIGHT: "แปด",
		NINE:  "เก้า",
	}

	return thaiNumberMap[v]
}

func mapThaiUnit(idx int, v string) string {
	var result string

	isMillion := idx >= 6
	if isMillion {
		idx = idx % 6
	}

	if v != ZERO {
		switch idx {
		case 1:
			result = "สิบ"
		case 2:
			result = "ร้อย"
		case 3:
			result = "พัน"
		case 4:
			result = "หมื่น"
		case 5:
			result = "แสน"
		}
	}

	if isMillion && idx == 0 {
		result = fmt.Sprintf("%sล้าน", result)
	}

	return result
}

func addSuffix(isFractional bool) string {
	if isFractional {
		return "สตางค์"
	}

	return BATH + "ถ้วน"
}

func getTotalFractional(isFractional bool, idxTotal int, value string) int {
	var result int

	if isFractional {
		for i := idxTotal; i >= 0; i-- {
			if string(value[i]) == DOT {
				break
			}

			result++
		}
	}

	return result
}
