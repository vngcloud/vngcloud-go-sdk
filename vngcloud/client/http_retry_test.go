package client

import (
	lerr "errors"
	lhttp "net/http"
	ltesting "testing"
	ltime "time"

	lreq "github.com/imroc/req/v3"
)

func respWithMethod(pmethod string) *lreq.Response {
	return &lreq.Response{Request: &lreq.Request{Method: pmethod}}
}

// TestRetryOnlyIdempotent: mac dinh cua req/v3 la `needRetry := err != nil`, tuc
// la thu lai MOI request gap loi transport - ke ca mot POST da timeout, ma khi do
// server rat co the da nhan va dang xu ly. Voi CreateBlockVolume thi thu lai la
// tao trung volume.
func TestRetryOnlyIdempotent(t *ltesting.T) {
	boom := lerr.New("dial tcp: i/o timeout")

	tcs := []struct {
		name   string
		method string
		err    error
		want   bool
	}{
		{"GET timeout thi thu lai", lhttp.MethodGet, boom, true},
		{"HEAD timeout thi thu lai", lhttp.MethodHead, boom, true},
		{"OPTIONS timeout thi thu lai", lhttp.MethodOptions, boom, true},
		{"TRACE timeout thi thu lai", lhttp.MethodTrace, boom, true},
		{"POST timeout thi KHONG thu lai", lhttp.MethodPost, boom, false},
		{"PUT timeout thi KHONG thu lai", lhttp.MethodPut, boom, false},
		{"PATCH timeout thi KHONG thu lai", lhttp.MethodPatch, boom, false},
		{"DELETE timeout thi KHONG thu lai", lhttp.MethodDelete, boom, false},
		{"GET thanh cong thi khong thu lai", lhttp.MethodGet, nil, false},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *ltesting.T) {
			if got := retryOnlyIdempotent(respWithMethod(tc.method), tc.err); got != tc.want {
				t.Fatalf("retryOnlyIdempotent(%s) = %v, muon %v", tc.method, got, tc.want)
			}
		})
	}
}

// TestRetryOnlyIdempotentHandlesNil: RetryConditionFunc duoc goi ca khi resp la
// nil (loi xay ra truoc khi co response), khong duoc panic o day.
func TestRetryOnlyIdempotentHandlesNil(t *ltesting.T) {
	boom := lerr.New("connection refused")

	for name, resp := range map[string]*lreq.Response{
		"resp nil":         nil,
		"resp.Request nil": {},
	} {
		if retryOnlyIdempotent(resp, boom) {
			t.Errorf("%s: khong xac dinh duoc method thi phai coi la khong thu lai", name)
		}
	}
}

// TestRetryIntervalIsNotNanoseconds khoa lai bug cu:
// SetCommonRetryFixedInterval(10) nhan hang so khong kieu nen thanh
// time.Duration(10) = 10 NANO giay, ba lan thu lai gan nhu tuc thi.
func TestRetryIntervalIsNotNanoseconds(t *ltesting.T) {
	if httpRetryIntervalMin < 100*ltime.Millisecond {
		t.Fatalf("khoang cho toi thieu = %v, qua ngan de goi la thu lai", httpRetryIntervalMin)
	}

	if httpRetryIntervalMax < httpRetryIntervalMin {
		t.Fatalf("tran %v nho hon san %v", httpRetryIntervalMax, httpRetryIntervalMin)
	}

	// Toan bo chuoi thu lai phai ngan hon nhieu so voi mot lan timeout, neu
	// khong thi rieng phan cho da lam caller treo them dang ke.
	if total := ltime.Duration(httpRetryCount) * httpRetryIntervalMax; total > 30*ltime.Second {
		t.Fatalf("tong thoi gian cho giua cac lan thu = %v, qua lau", total)
	}
}
