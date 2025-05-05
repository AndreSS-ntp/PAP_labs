package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab5/internal/domain"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("SOAP клиент для конвертации температуры")
	fmt.Println("Выберите операцию:")
	fmt.Println("1. Fahrenheit to Celsius")
	fmt.Println("2. Celsius to Fahrenheit")

	var choice int
	fmt.Scanln(&choice)

	var temp string
	fmt.Println("Введите температуру:")
	fmt.Scanln(&temp)

	var soapAction, soapRequest string

	switch choice {
	case 1:
		soapAction = "https://www.w3schools.com/xml/FahrenheitToCelsius"
		soapRequest = fmt.Sprintf(domain.FToCSoapRequest, temp)
	case 2:
		soapAction = "https://www.w3schools.com/xml/CelsiusToFahrenheit"
		soapRequest = fmt.Sprintf(domain.CToFSoapRequest, temp)
	default:
		log.Fatal("Неверный выбор")
	}

	fmt.Println("\nsend SOAP-request:")
	fmt.Println(soapRequest)

	req, err := http.NewRequest("POST", "https://www.w3schools.com/xml/tempconvert.asmx", bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		log.Fatal("Ошибка создания запроса:", err)
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Ошибка отправки запроса:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка чтения ответа:", err)
	}

	fmt.Println("\nSOAP-response:")
	fmt.Println(string(body))

	var result domain.SoapResponse
	err = xml.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("Ошибка парсинга ответа:", err)
	}

	fmt.Println("\nРезультат конвертации:")
	fmt.Println(result.Body.Response.Result)
}
