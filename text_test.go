package tinyml

import (
	"testing"
)

func assertToText(t *testing.T, cases map[string]string) {
	for key, expected := range cases {
		actual, err := ToText(key)
		if err != nil {
			t.Errorf("\n%v has error: %v", key, err)
		}

		if actual != expected {
			t.Errorf("\nexpected: %v\nactual  : %v", expected, actual)
		}
	}
}

func TestToText(t *testing.T) {
	cases := map[string]string{
		"在野蛮的战场上还是有些文明的微光在闪动，那就是人性所在，确实，那就是我们仅有的谦卑的温和的方式。":                    "在野蛮的战场上还是有些文明的微光在闪动，那就是人性所在，确实，那就是我们仅有的谦卑的温和的方式。",
		"你好，世界[st]ST/US/BABA#阿里巴巴.US[/st]港股上市\n\n这是第二行":                       "你好，世界 阿里巴巴.US 港股上市\n\n这是第二行",
		" [st]ST/US/BABA#阿里巴巴.US[/st]港股上市\n这是第二行[st]ST/HK/00700#腾讯集团.HK[/st]": "阿里巴巴.US 港股上市\n这是第二行 腾讯集团.HK",
	}

	assertToText(t, cases)
}

func BenchmarkToText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 0.02ms/op
		ToText("消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。")
	}
}
