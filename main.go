package main

import (
    "fmt"
	"runtime"
	"github.com/0adri3n/3g-scan/network"
	"github.com/rivo/tview"
	"strings"

)

func main() {


	app := tview.NewApplication()
	form := tview.NewForm().
		AddTextArea("IP addresses", "", 40, 0, 0, nil).
		AddTextView("Notes", "Please enter one IP per line !", 40, 2, true, false).
		AddButton("Save", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("3g-scan").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
	ipsStr := form.GetFormItem(0).(*tview.TextArea).GetText()
	ips := strings.Split(ipsStr, "\n")

	fmt.Printf("%v\n", ips)

	fmt.Println("3g-scan started")

	fmt.Println("Please select an interface : ")
	interfaceName := network.SelectInterface()

    fmt.Println("Let's scan ip adresses !")

	for _, ip := range ips {
		
		fmt.Printf("\nScanning %v\n-----------------------------\n", ip)

		network.Pinger(ip)

		switch runtime.GOOS {
		case "windows":
			network.WindowsMaccer(ip)
		case "linux", "darwin":
			network.LinuxMaccer(ip, interfaceName)
		}
		fmt.Println("-----------------------------")

	}

	var exit string
	fmt.Println("\n\nPress any key then enter to exit...")
	fmt.Scan(&exit)


}