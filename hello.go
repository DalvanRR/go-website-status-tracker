package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorings = 3
const delay = 5

func main() {
	showIntroduction()
	for {

		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			printLog()
		case 0:
			fmt.Println("Exiting program...")
			exitProgram()
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	var name string = "Dalvan"
	var age = 25
	version := 1.3
	fmt.Println("Hello Sr.", name, "Your age is", age)
	fmt.Println("This is the", version, "version")
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")
}

func readCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("The chosen command was", readCommand)
	fmt.Println("")

	return readCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readFileSites()

	for i := 0; i < monitorings; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func showLog() {
	fmt.Println("Showing logs...")
}

func exitProgram() {
	fmt.Println("Exiting program...")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error occurred:", err.Error())
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was successfully loaded!")
		logRegistration(site, true)
	} else {
		fmt.Println("Site:", site, "is with problems. Status Code:", resp.StatusCode)
		logRegistration(site, false)
	}

}

func readFileSites() []string {

	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func logRegistration(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + fmt.Sprint(status) + "\n")

	file.Close()

}
func printLog() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
