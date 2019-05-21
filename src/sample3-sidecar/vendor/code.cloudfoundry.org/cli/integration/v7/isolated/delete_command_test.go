package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("delete command", func() {
	var (
		orgName   string
		spaceName string
		appName   string
	)

	BeforeEach(func() {
		orgName = helpers.NewOrgName()
		spaceName = helpers.NewSpaceName()
		appName = helpers.PrefixedRandomName("app")
	})

	When("--help flag is set", func() {
		It("Displays command usage to output", func() {
			session := helpers.CF("delete", "--help")
			Eventually(session).Should(Say("NAME:"))
			Eventually(session).Should(Say("delete - Delete an app"))
			Eventually(session).Should(Say("USAGE:"))
			Eventually(session).Should(Say(`cf delete APP_NAME \[-r\] \[-f\]`))
			Eventually(session).Should(Say("OPTIONS:"))
			Eventually(session).Should(Say(`\s+-f\s+Force deletion without confirmation`))
			Eventually(session).Should(Say(`\s+-r\s+Also delete any mapped routes \[Not currently functional\]`))
			Eventually(session).Should(Exit(0))
		})
	})

	When("the app name is not provided", func() {
		It("tells the user that the app name is required, prints help text, and exits 1", func() {
			session := helpers.CF("delete")

			Eventually(session.Err).Should(Say("Incorrect Usage: the required argument `APP_NAME` was not provided"))
			Eventually(session).Should(Say("NAME:"))
			Eventually(session).Should(Exit(1))
		})
	})

	When("the environment is not setup correctly", func() {
		It("fails with the appropriate errors", func() {
			helpers.CheckEnvironmentTargetedCorrectly(true, true, ReadOnlyOrg, "delete", "-f", appName)
		})
	})

	When("the environment is setup correctly", func() {
		BeforeEach(func() {
			helpers.SetupCF(orgName, spaceName)
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(orgName)
		})

		When("the app does not exist", func() {
			When("the -f flag is provided", func() {
				It("it displays the app does not exist", func() {
					username, _ := helpers.GetCredentials()
					session := helpers.CF("delete", appName, "-f")
					Eventually(session).Should(Say("Deleting app %s in org %s / space %s as %s...", appName, orgName, spaceName, username))
					Eventually(session).Should(Say("App %s does not exist", appName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Exit(0))
				})
			})

			When("the -f flag not is provided", func() {
				var buffer *Buffer

				BeforeEach(func() {
					buffer = NewBuffer()
				})

				When("the user enters 'y'", func() {
					BeforeEach(func() {
						_, err := buffer.Write([]byte("y\n"))
						Expect(err).ToNot(HaveOccurred())
					})

					It("it displays the app does not exist", func() {
						username, _ := helpers.GetCredentials()
						session := helpers.CFWithStdin(buffer, "delete", appName)
						Eventually(session).Should(Say(`Really delete the app %s\? \[yN\]`, appName))
						Eventually(session).Should(Say("Deleting app %s in org %s / space %s as %s...", appName, orgName, spaceName, username))
						Eventually(session).Should(Say("App %s does not exist", appName))
						Eventually(session).Should(Say("OK"))
						Eventually(session).Should(Exit(0))
					})
				})

				When("the user enters 'n'", func() {
					BeforeEach(func() {
						_, err := buffer.Write([]byte("n\n"))
						Expect(err).ToNot(HaveOccurred())
					})

					It("does not delete the app", func() {
						session := helpers.CFWithStdin(buffer, "delete", appName)
						Eventually(session).Should(Say(`Really delete the app %s\? \[yN\]`, appName))
						Eventually(session).Should(Say("Delete cancelled"))
						Eventually(session).Should(Exit(0))
					})
				})

				When("the user enters the default input (hits return)", func() {
					BeforeEach(func() {
						_, err := buffer.Write([]byte("\n"))
						Expect(err).ToNot(HaveOccurred())
					})

					It("does not delete the app", func() {
						session := helpers.CFWithStdin(buffer, "delete", appName)
						Eventually(session).Should(Say(`Really delete the app %s\? \[yN\]`, appName))
						Eventually(session).Should(Say("Delete cancelled"))
						Eventually(session).Should(Exit(0))
					})
				})

				When("the user enters an invalid answer", func() {
					BeforeEach(func() {
						// The second '\n' is intentional. Otherwise the buffer will be
						// closed while the interaction is still waiting for input; it gets
						// an EOF and causes an error.
						_, err := buffer.Write([]byte("wat\n\n"))
						Expect(err).ToNot(HaveOccurred())
					})

					It("asks again", func() {
						session := helpers.CFWithStdin(buffer, "delete", appName)
						Eventually(session).Should(Say(`Really delete the app %s\? \[yN\]`, appName))
						Eventually(session).Should(Say(`invalid input \(not y, n, yes, or no\)`))
						Eventually(session).Should(Say(`Really delete the app %s\? \[yN\]`, appName))
						Eventually(session).Should(Exit(0))
					})
				})
			})
		})

		When("the app exists", func() {
			BeforeEach(func() {
				helpers.WithHelloWorldApp(func(appDir string) {
					Eventually(helpers.CustomCF(helpers.CFEnv{WorkingDirectory: appDir}, "push", appName)).Should(Exit(0))
				})
			})

			It("deletes the app", func() {
				session := helpers.CF("delete", appName, "-f")
				username, _ := helpers.GetCredentials()
				Eventually(session).Should(Say("Deleting app %s in org %s / space %s as %s...", appName, orgName, spaceName, username))
				Eventually(session).Should(Say("OK"))
				Eventually(session).Should(Exit(0))

				Eventually(helpers.CF("app", appName)).Should(Exit(1))
			})
		})
	})
})
