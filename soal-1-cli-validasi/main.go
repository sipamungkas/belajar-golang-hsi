package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sanitize(input string) (string, error) {
	result := strings.TrimSpace(input)
	if len(result) < 1 {
		return "", errors.New("input tidak valid")
	}
	return result, nil
}

func validateName(input string) (string, error) {
	if len(input) < 1 {
		return "", errors.New("mohon masukkan nama anda")
	}
	name, err := sanitize(input)
	if err != nil {
		return "", err
	}
	return name, nil

}

func validateAge(input string) error {
	sanitizedInput, err := sanitize(input)
	if err != nil {
		return err
	}

	age, err := strconv.Atoi(sanitizedInput)
	if err != nil || age < 18 {
		return errors.New("umur tidak valid (minimal 18 tahun)")
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama: ")
	rawName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	name, err := validateName(rawName)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Print("Umur: ")
	rawAge, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	ageError := validateAge(rawAge)

	if ageError != nil {
		fmt.Printf("Error: %s", ageError)
		return
	}

	fmt.Printf("Selamat datang, %s!", name)

}
