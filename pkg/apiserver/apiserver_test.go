package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/genridarkbkru/LinkShortenerApi/pkg/errors"
	"github.com/genridarkbkru/LinkShortenerApi/pkg/internal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiserver_GetShortUrl(t *testing.T) {
	psqlconn, tableName := internal.GetSqlconnAndTableName()
	addr, _ := internal.GetAddrProtocolWithDomain()
	s := NewServer(addr, psqlconn, tableName, false)

	testCases := []struct {
		name             string
		payload          interface{}
		expectedCode     int
		expectedShortUrl string
		expectedError    string
	}{
		{
			name: "create record",
			payload: map[string]string{
				"url": "user@example.org",
			},
			expectedCode:     http.StatusCreated,
			expectedShortUrl: internal.GetShortUrl("user@example.org"),
			expectedError:    "",
		},
		{
			name: "the record is already in the table",
			payload: map[string]string{
				"url": "user@example.org",
			},
			expectedCode:     http.StatusOK,
			expectedShortUrl: internal.GetShortUrl("user@example.org"),
			expectedError:    "",
		},
		{
			name: "length is zero symbols",
			payload: map[string]string{
				"ulr": "",
			},
			expectedCode:     http.StatusBadRequest,
			expectedShortUrl: "",
			expectedError:    errors.RangeOutLenUrl.Error(),
		},
		{
			name: "length is 500 symbols",
			payload: map[string]string{
				"ulr": "012345678910111213141516171819202122232425262728293031323334353637383940414243444546474849505152" +
					"535455565758596061626364656667686970717273747576777879808182838485868788899091929394959697989910010" +
					"110210310410510610710810911011111211311411511611711811912012112212312412512612712812913013113213313" +
					"413513613713813914014114214314414514614714814915015115215315415515615715815916016116216316416516616" +
					"716816917017117217317417517617717817918018118218318418518618718818919019119219319419519619719819920" +
					"020120220320420520620720820921021121221321421521621721821922022122222322422522622722822923023123223" +
					"323423523623723823924024124224324424524624724824925025125225325425525625725825926026126226326426526" +
					"626726826927027127227327427527627727827928028128228328428528628728828929029129229329429529629729829" +
					"930030130230330430530630730830931031131231331431531631731831932032132232332432532632732832933033133" +
					"233333433533633733833934034134234334434534634734834935035135235335435535635735835936036136236336436" +
					"536636736836937037137237337437537637737837938038138238338438538638738838939039139239339439539639739" +
					"839940040140240340440540640740840941041141241341441541641741841942042142242342442542642742842943043" +
					"143243343443543643743843944044144244344444544644744844945045145245345445545645745845946046146246346" +
					"446546646746846947047147247347447547647747847948048148248348448548648748848949049149249349449549649" +
					"7498499500",
			},
			expectedCode:     http.StatusBadRequest,
			expectedShortUrl: "",
			expectedError:    errors.RangeOutLenUrl.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/GetShortUrl", b)
			s.Handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
			if rec.Code == 200 || rec.Code == 201 {
				var actualShortUrl string
				json.NewDecoder(rec.Body).Decode(&actualShortUrl)
				assert.Equal(t, tc.expectedShortUrl, actualShortUrl)
			} else {
				actualError := rec.Body.String()
				assert.Equal(t, tc.expectedError, actualError[:len(actualError)-1])
			}
		})
	}
}

func TestApiserver_GetFullUrl(t *testing.T) {
	psqlconn, tableName := internal.GetSqlconnAndTableName()
	addr, _ := internal.GetAddrProtocolWithDomain()
	s := NewServer(addr, psqlconn, tableName, false)

	rec := httptest.NewRecorder()
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(map[string]string{"url": "user@example.org"})
	req := httptest.NewRequest(http.MethodPost, "/GetShortUrl", b)
	s.Handler.ServeHTTP(rec, req)

	testCases := []struct {
		name          string
		payload       map[string]string
		expectedCode  int
		expectedUrl   string
		expectedError string
	}{
		{
			name: "record in table",
			payload: map[string]string{
				"url": internal.GetShortUrl("user@example.org"),
			},
			expectedCode:  http.StatusFound,
			expectedUrl:   "user@example.org",
			expectedError: "",
		},
		{
			name: "record not in table",
			payload: map[string]string{
				"url": internal.GetShortUrl("user@example.org1"),
			},
			expectedCode:  http.StatusNotFound,
			expectedUrl:   "",
			expectedError: errors.RecordNotFound.Error(),
		},
		{
			name: "length shortUrl is not 10",
			payload: map[string]string{
				"ulr": "",
			},
			expectedCode:  http.StatusBadRequest,
			expectedUrl:   "",
			expectedError: errors.RangeOutLenShortUrl.Error(),
		},
		{
			name: "length shortUrl is not 10",
			payload: map[string]string{
				"ulr": "12345678901",
			},
			expectedCode:  http.StatusBadRequest,
			expectedUrl:   "",
			expectedError: errors.RangeOutLenShortUrl.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec = httptest.NewRecorder()
			b = &bytes.Buffer{}
			//json.NewEncoder(b).Encode(tc.payload)
			req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/GetFullUrl?short_url=%s", tc.payload["url"]), b)
			s.Handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
			if rec.Code == 302 {
				var actualUrl string
				json.NewDecoder(rec.Body).Decode(&actualUrl)
				assert.Equal(t, tc.expectedUrl, actualUrl)
			} else {
				actualError := rec.Body.String()
				assert.Equal(t, tc.expectedError, actualError[:len(actualError)-1])
			}
		})
	}
}
