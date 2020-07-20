package tinyml

import (
	"testing"
)

func assertToHTML(t *testing.T, cases map[string]string) {
	for key, expected := range cases {
		actual, err := ToHTML(key)
		if err != nil {
			t.Errorf("\n%v has error: %v", key, err)
		}

		if actual != expected {
			t.Errorf("\nexpected: %v\nactual  : %v", expected, actual)
		}
	}
}

func TestToHTML(t *testing.T) {
	cases := map[string]string{
		"消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。\n\n这是第二行": `<p>消息称 <span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span> 将于 5 月，在港股上市。</p><p>这是第二行</p>`,
		"消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。":          `<p>消息称 <span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span> 将于 5 月，在港股上市。</p>`,
		"[st]ST/US/BABA#阿里巴巴.US[/st]\n\n将于5月，在港股上市。":         `<p><span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span></p><p>将于 5 月，在港股上市。</p>`,
	}

	assertToHTML(t, cases)
}

func BenchmarkToHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 0.03ms/op
		ToHTML("消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。\n\n几家领头羊都处于第一步或者第二步。")
	}
}
