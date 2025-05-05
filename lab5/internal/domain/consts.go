package domain

import "encoding/xml"

const FToCSoapRequest = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <FahrenheitToCelsius xmlns="https://www.w3schools.com/xml/">
      <Fahrenheit>%s</Fahrenheit>
    </FahrenheitToCelsius>
  </soap:Body>
</soap:Envelope>`

const CToFSoapRequest = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <CelsiusToFahrenheit xmlns="https://www.w3schools.com/xml/">
      <Celsius>%s</Celsius>
    </CelsiusToFahrenheit>
  </soap:Body>
</soap:Envelope>`

type SoapResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName  xml.Name `xml:"Body"`
		Response struct {
			XMLName xml.Name `xml:"FahrenheitToCelsiusResponse"`
			Result  string   `xml:"FahrenheitToCelsiusResult"`
		} `xml:"FahrenheitToCelsiusResponse"`
	} `xml:"Body"`
}
