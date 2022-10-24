package fanhao

import (
	"testing"
)

func assertNormalizeResult(t *testing.T, src, dest string) {
	ret := Normalize(src)
	if ret != dest {
		t.Logf(" %v ---> %v , should be %v", src, ret, dest)
		t.Fail()
	}
}

func TestNormalize(t *testing.T) {
	assertNormalizeResult(t, "abc010_1.2K.avi", "ABC-010.avi")
	assertNormalizeResult(t, "etf-123.mp4", "ETF-123.mp4")
	assertNormalizeResult(t, "001-124.AVI", "001-124.avi")
	assertNormalizeResult(t, "T28-123.jpg", "T28-123.jpg")
	assertNormalizeResult(t, "abef-213r.gif", "ABEF-213.gif")
	assertNormalizeResult(t, "eft-124a.jpg", "EFT-124_A.jpg")
	assertNormalizeResult(t, "GFE-123-b.rmvb", "GFE-123_B.rmvb")
	assertNormalizeResult(t, "abs-104pl.jpg", "ABS-104.jpg")
	assertNormalizeResult(t, "EKDV-152 .jpg", "EKDV-152.jpg")
	assertNormalizeResult(t, "mide023.avi", "MIDE-023.avi")
	assertNormalizeResult(t, "MILD_753.jpg", "MILD-753.jpg")
	assertNormalizeResult(t, "sadr-052r.jpg", "SADR-052.jpg")
	assertNormalizeResult(t, "sadr-052rpl.jpg", "SADR-052.jpg")
	assertNormalizeResult(t, "ZDAD-28_ENG_01.rmvb", "ZDAD-28_ENG_01.rmvb")
	assertNormalizeResult(t, "COSQ-017_1.rmvb", "COSQ-017_1.rmvb")
	assertNormalizeResult(t, "ENFD-5401 Extra.rmvb", "ENFD-5401_E.rmvb")
	assertNormalizeResult(t, "heyzo_lt_0203.jpg", "HEYZO_LT-203.jpg")
	assertNormalizeResult(t, "abs-55a.avi", "ABS-055_A.avi")
	assertNormalizeResult(t, "abs-56_a.avi", "ABS-056_A.avi")
	assertNormalizeResult(t, "abs000055a.avi", "ABS-055_A.avi")
	assertNormalizeResult(t, "abs600.avi", "ABS-600.avi")
	assertNormalizeResult(t, "abs060.avi", "ABS-060.avi")
	assertNormalizeResult(t, "abs006.avi", "ABS-006.avi")
	assertNormalizeResult(t, "hD-abcd-088.avi", "ABCD-088.avi")
	assertNormalizeResult(t, "abc089.avi.avi", "ABC-089.avi")
	assertNormalizeResult(t, "abc090.mp4.mp4.mp4", "ABC-090.mp4")
	assertNormalizeResult(t, "[TtZz.Yy]abc-095.mp4", "ABC-095.mp4")
	assertNormalizeResult(t, "abcd618 m .avi", "ABCD-618.avi")
	assertNormalizeResult(t, "abcd619 nnn .avi", "ABCD-619.avi")
	assertNormalizeResult(t, "abc103++ .avi", "ABC-103.avi")
	assertNormalizeResult(t, "HND-581~kan224.com.mp4", "HND-581.mp4")
	assertNormalizeResult(t, "SW610.mp4", "SW-610.mp4")
	assertNormalizeResult(t, "opqr-123-7.mp4", "OPQR-123_7.mp4")
	assertNormalizeResult(t, "HD_apns-196.mp4", "APNS-196.mp4")
	assertNormalizeResult(t, "HD_apns-196(A)(B).mp4", "APNS-196.mp4")
	assertNormalizeResult(t, "CAWD-102C(720P)@18P2P.mp4", "CAWD-102_C.mp4")
	assertNormalizeResult(t, "@方块字@hello.com_ABC-999_1.2K.mp4", "ABC-999.mp4")
	assertNormalizeResult(t, "MDVR-018E.VR.mp4", "MDVR-018_E.mp4")
	assertNormalizeResult(t, "kiwvr-363 2.mp4", "KIWVR-363_2.mp4")
}
