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
		"消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。\n\n这是第二行":                     `<p>消息称 <span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span> 将于 5 月，在港股上市。</p><p>这是第二行</p>`,
		"消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。":                              `<p>消息称 <span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span> 将于 5 月，在港股上市。</p>`,
		"[st]ST/US/BABA#阿里巴巴.US[/st]\n\n将于5月，在港股上市。":                             `<p><span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span></p><p>将于 5 月，在港股上市。</p>`,
		"[st]ST/US/BABA#阿里巴巴.US[/st]\n\n将于5月，在港股上市。\n并称消息准确。":                    `<p><span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span></p><p>将于 5 月，在港股上市。<br/>并称消息准确。</p>`,
		"测试普通的换行\n这是第二行":                                                         `<p>测试普通的换行<br/>这是第二行</p>`,
		"[st]ST/SZ/000014#沙河股份.SZ[/st] 测试中间带空格的情况 [st]ST/SH/601318#中国平安.SH[/st]": `<p><span class="security-tag" value="ST/SZ/000014#沙河股份.SZ" data-id="ST/SZ/000014">沙河股份.SZ</span> 测试中间带空格的情况 <span class="security-tag" value="ST/SH/601318#中国平安.SH" data-id="ST/SH/601318">中国平安.SH</span></p>`,
	}

	assertToHTML(t, cases)
}

func BenchmarkToHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 0.06ms/op
		ToHTML("消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。\n\n几家领头羊都处于第一步或者第二步。")
	}
}
