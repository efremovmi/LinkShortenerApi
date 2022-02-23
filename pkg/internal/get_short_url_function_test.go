package internal

import (
	"github.com/genridarkbkru/LinkShortenerApi/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SumTwoStrings(t *testing.T) {
	testCases := []struct {
		name          string
		leftValue     string
		rightValue    string
		expectedValue string
		expectedError error
	}{
		{
			name:          "empty strings",
			leftValue:     "",
			rightValue:    "",
			expectedValue: "",
			expectedError: nil,
		},
		{
			name:          "left string is empty",
			leftValue:     "",
			rightValue:    "ab",
			expectedValue: "",
			expectedError: nil,
		},
		{
			name:          "the length of the left string is greater than the right string",
			leftValue:     "abcdefghijklmnopqrstuvwxyz",
			rightValue:    "12345",
			expectedValue: "",
			expectedError: errors.RightLenLessLeftLen,
		},
		{
			name:          "two strings of the same length 1",
			leftValue:     "ab",
			rightValue:    "  ",
			expectedValue: "\x81\x82",
			expectedError: nil,
		},
		{
			name:          "two strings of the same lenght 2",
			leftValue:     "ab",
			rightValue:    "ba",
			expectedValue: "\xc3\xc3",
			expectedError: nil,
		},
		{
			name:          "two strings of the same length 3",
			leftValue:     "abcdefghijklmnopqrstuvwxyz",
			rightValue:    "01234567890123456789012345",
			expectedValue: "\x91\x93\x95\x97\x99\x9b\x9d\x9f\xa1\xa3\x9b\x9d\x9f\xa1\xa3\xa5\xa7\xa9\xab\xad\xa5\xa7\xa9\xab\xad\xaf",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualValue, actualErr := sumTwoStrings(tc.leftValue, tc.rightValue)
			assert.Equal(t, tc.expectedValue, actualValue)
			assert.Equal(t, tc.expectedError, actualErr)
		})
	}
}

func Test_GetHash(t *testing.T) {
	testCases := []struct {
		name           string
		inputString    string
		expectedValue  string
		expectedLength int
	}{
		{
			name:           "input string consists of zero characters",
			inputString:    "",
			expectedValue:  "B6f9h9ilsG",
			expectedLength: 10,
		},
		{
			name:           "input string consists of one characters",
			inputString:    "a",
			expectedValue:  "5jslfvILF_",
			expectedLength: 10,
		},
		{
			name:           "input string consists of two characters",
			inputString:    "ab",
			expectedValue:  "g5cbfOtWC2",
			expectedLength: 10,
		},
		{
			name:           "input string - link to Ozon.Fintech task",
			inputString:    "https://docs.google.com/document/d/1gPAgIpscDjXrczlDdzLfS-XJqpu59HjcgRgO0eRsTvM/edit",
			expectedValue:  "PDF9tLHmhw",
			expectedLength: 10,
		},
		{
			// Алгоритм не чувствителен к большим ссылкам, тк после 46 символа он перестает видеть строку (этот артефакт исправляет функция GetShortUrl)
			name:           "input string - link to Ozon.Fintech task with some fix",
			inputString:    "https://docs.google.com/document/d/1gPAgIpscDjXrczlDdzLfS-XJqpu59HjcgRgO0eRsTvM/edit1234",
			expectedValue:  "PDF9tLHmhw",
			expectedLength: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualValue := getHash(tc.inputString)
			assert.Equal(t, tc.expectedValue, actualValue)
			assert.Equal(t, tc.expectedLength, len(actualValue))

		})
	}

}

func Test_GetShortUrl(t *testing.T) {
	testCases := []struct {
		name             string
		inputUrl         string
		expectedShortUrl string
		expectedLength   int
	}{
		{
			name:             "input string - link to Ozon.Fintech task",
			inputUrl:         "https://docs.google.com/document/d/1gPAgIpscDjXrczlDdzLfS-XJqpu59HjcgRgO0eRsTvM/edit",
			expectedShortUrl: "8LcztdH2Fl",
			expectedLength:   10,
		},
		{
			name:             "input string - link to Ozon.Fintech task with some fix",
			inputUrl:         "https://docs.google.com/document/d/1gPAgIpscDjXrczlDdzLfS-XJqpu59HjcgRgO0eRsTvM/edit1234",
			expectedShortUrl: "GXIXFGcWC8",
			expectedLength:   10,
		},
		{
			name: "input string - link to linear algebra textbook",
			inputUrl: "https://docs.yandex.ru/docs/view?tm=1642176901&tld=ru&lang=ru&name=560.pdf&text=линейная%" +
				"20алгебра%20сборник%20задач&url=http%3A%2F%2Felibrary.sgu.ru%2Fuch_lit%2F560.pdf&lr=10708&mime=pdf&l10n" +
				"=ru&sign=fdff09fecd61494b4273e530cadab3fb&keyno=0&nosw=1&serpParams=tm%3D1642176901%26tld%3Dru%26lang%3" +
				"Dru%26name%3D560.pdf%26text%3D%25D0%25BB%25D0%25B8%25D0%25BD%25D0%25B5%25D0%25B9%25D0%25BD%25D0%25B0%25" +
				"D1%258F%2B%25D0%25B0%25D0%25BB%25D0%25B3%25D0%25B5%25D0%25B1%25D1%2580%25D0%25B0%2B%25D1%2581%25D0%25B1" +
				"%25D0%25BE%25D1%2580%25D0%25BD%25D0%25B8%25D0%25BA%2B%25D0%25B7%25D0%25B0%25D0%25B4%25D0%25B0%25D1%2587" +
				"%26url%3Dhttp%253A%2F%2Felibrary.sgu.ru%2Fuch_lit%2F560.pdf%26lr%3D10708%26mime%3Dpdf%26l10n%3Dru%26sig" +
				"n%3Dfdff09fecd61494b4273e530cadab3fb%26keyno%3D0%26nosw%3D1",

			expectedShortUrl: "GGiaF2cOSq",
			expectedLength:   10,
		},
		{
			name: "input string - link to linear algebra textbook with 3 characters changed in the middle",
			inputUrl: "https://docs.yandex.ru/docs/view?tm=1642176901&tld=ru&lang=ru&name=560.pdf&text=линейная%" +
				"20алгебра%20сборник%20задач&url=http%3A%2F%2Felibrary.sgu.ru%2Fuch_lit%2F560.pdf&lr=10708&mime=pdf&l10n" +
				"=ru&sign=fdff09fecd61494b4273e530cadab3fb&keyno=0&nosw=1&serpParams=tm%3D1642176901%26tld%3Dru%26lang%3" +
				"Dru%26name%3D560.pdf%26text%3D%25D0%fffB%25D0%25B8%25D0%25BD%25D0%25B5%25D0%25B9%25D0%25BD%25D0%25B0%25" +
				"D1%258F%2B%25D0%25B0%25D0%25BB%25D0%25B3%25D0%25B5%25D0%25B1%25D1%2580%25D0%25B0%2B%25D1%2581%25D0%25B1" +
				"%25D0%25BE%25D1%2580%25D0%25BD%25D0%25B8%25D0%25BA%2B%25D0%25B7%25D0%25B0%25D0%25B4%25D0%25B0%25D1%2587" +
				"%26url%3Dhttp%253A%2F%2Felibrary.sgu.ru%2Fuch_lit%2F560.pdf%26lr%3D10708%26mime%3Dpdf%26l10n%3Dru%26sig" +
				"n%3Dfdff09fecd61494b4273e530cadab3fb%26keyno%3D0%26nosw%3D1",

			expectedShortUrl: "blUDSpuZTd",
			expectedLength:   10,
		},
		{
			name: "input string - link to linear algebra textbook with 3 characters changed at the end",
			inputUrl: "https://docs.yandex.ru/docs/view?tm=1642176901&tld=ru&lang=ru&name=560.pdf&text=линейная%" +
				"20алгебра%20сборник%20задач&url=http%3A%2F%2Felibrary.sgu.ru%2Fuch_lit%2F560.pdf&lr=10708&mime=pdf&l10n" +
				"=ru&sign=fdff09fecd61494b4273e530cadab3fb&keyno=0&nosw=1&serpParams=tm%3D1642176901%26tld%3Dru%26lang%3" +
				"Dru%26name%3D560.pdf%26text%3D%25D0%25BB%25D0%25B8%25D0%25BD%25D0%25B5%25D0%25B9%25D0%25BD%25D0%25B0%25" +
				"D1%258F%2B%25D0%25B0%25D0%25BB%25D0%25B3%25D0%25B5%25D0%25B1%25D1%2580%25D0%25B0%2B%25D1%2581%25D0%25B1" +
				"%25D0%25BE%25D1%2580%25D0%25BD%25D0%25B8%25D0%25BA%2B%25D0%25B7%25D0%25B0%25D0%25B4%25D0%25B0%25D1%2587" +
				"%26url%3Dhttp%253A%2F%2Felibrary.sgu.ru%2Fuch_lit%2F560.pdf%26lr%3D10708%26mime%3Dpdf%26l10n%3Dru%26sig" +
				"n%3Dfdff09fecd61494b4273e530cadab3fb%26keyno%3D0%26nosw%fff",

			expectedShortUrl: "GdSvFRsXhx",
			expectedLength:   10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualShortUrl := GetShortUrl(tc.inputUrl)
			assert.Equal(t, tc.expectedShortUrl, actualShortUrl)
			assert.Equal(t, tc.expectedLength, len(actualShortUrl))
		})
	}
}
