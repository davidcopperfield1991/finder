package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const INPUT_FILE = "../../.bash_history"

func readLines(path string) ([]string, error) {

	_, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println("file dastan shod")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func peydakon(loghat string) []string {
	arzesh, _ := readLines(INPUT_FILE)

	sort.Strings(arzesh)

	fmt.Scanln(loghat)

	fmt.Println(loghat)
	var javab []string
	for i := range arzesh {
		s := strings.Contains(arzesh[i], loghat)

		if s == true {
			fmt.Println(arzesh[i])
			javab = append(javab, arzesh[i])
		}
	}
	fmt.Println("javab")
	return javab
}

type UU struct {
	Matn map[string]interface{}
}

func find(c *fiber.Ctx) error {
	loghat := c.Params("loghat")
	peydakon(loghat)
	fmt.Println(loghat)
	result := peydakon(loghat)
	return c.JSON(result)
}

func Routers(app *fiber.App) {
	app.Get("/find/:loghat", find)
}

func main() {
	app := fiber.New()
	Routers(app)
	app.Listen(":3333")

}
