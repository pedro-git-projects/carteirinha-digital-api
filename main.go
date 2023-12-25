package main

import (
	"fmt"
	"os"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	a := Aluno{"pedro_meu_email@gmail.com", "falksjfdkaj12j31lj2"}

	jsn, err := a.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal struct with error: %v", err)
		return
	}

	qrc, err := qrcode.New(string(jsn))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate QR code with error: %v", err)
		return
	}

	w, err := standard.New("./assets/qr-code.jpeg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed create writer with: %v", err)
		return
	}

	if err = qrc.Save(w); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write QR code to file with: %v", err)
		return
	}

	fmt.Println(qrc)
}
