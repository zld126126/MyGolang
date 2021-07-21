package baidu_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("baidu login", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("test baidu.com", func() {
		By("redirecting the home page", func() {
			Expect(page.Navigate("https://www.baidu.com/")).To(Succeed())
			//Expect(page).To(HaveURL("https://www.baidu.com"))
		})

		By("login", func() {
			Eventually(page.Find("#lh > a:nth-child(2)")).Should(HaveText("关于百度"))
		})
	})
})
