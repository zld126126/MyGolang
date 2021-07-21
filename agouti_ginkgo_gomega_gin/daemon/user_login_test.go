package daemon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("login", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("test localhost:9090", func() {
		By("redirecting the home page", func() {
			Expect(page.Navigate("http://localhost:9090/index")).To(Succeed())
			//Expect(page).To(HaveURL("https://www.baidu.com"))
		})

		By("login success", func() {
			Eventually(page.Find("#myTitle")).Should(BeFound())
			Expect(page.Find("#myUserName").Fill("dongtech")).To(Succeed())
			Expect(page.Find("#myPassword").Fill("123456")).To(Succeed())
			Expect(page.Find("#myButton").Click())
			Eventually(page.Find("#myTitle")).Should(HaveText("登录成功"))
		})

		By("login failed", func() {
			Eventually(page.Find("#myTitle")).Should(BeFound())
			Expect(page.Find("#myUserName").Fill("dongtech")).To(Succeed())
			Expect(page.Find("#myPassword").Fill("654321")).To(Succeed())
			Expect(page.Find("#myButton").Click())
			Eventually(page.Find("#myTitle")).Should(HaveText("登录失败"))
		})
	})
})
