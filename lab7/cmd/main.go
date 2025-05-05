package main

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"log"

	"github.com/go-ole/go-ole/oleutil"
)

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		log.Fatal("Ошибка создания Word:", err)
	}
	defer unknown.Release()

	word, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatal(err)
	}
	defer word.Release()

	oleutil.PutProperty(word, "Visible", true)

	docs := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	defer docs.Release()

	doc := oleutil.MustCallMethod(docs, "Add").ToIDispatch()
	defer doc.Release()

	selection := oleutil.MustGetProperty(word, "Selection").ToIDispatch()
	defer selection.Release()

	var labNumber, labName, studentName, group, date string
	fmt.Print("Номер лабораторной работы: ")
	fmt.Scanln(&labNumber)
	fmt.Print("Название работы: ")
	fmt.Scanln(&labName)
	fmt.Print("Ваше имя: ")
	fmt.Scanln(&studentName)
	fmt.Print("Ваша группа: ")
	fmt.Scanln(&group)
	fmt.Print("Дата выполнения (дд.мм.гггг): ")
	fmt.Scanln(&date)

	oleutil.MustCallMethod(selection, "TypeText", "Лабораторная работа №"+labNumber+"\n")
	oleutil.MustCallMethod(selection, "TypeText", "«"+labName+"»\n\n")
	oleutil.MustCallMethod(selection, "TypeText", "Выполнил: студент группы "+group+"\n")
	oleutil.MustCallMethod(selection, "TypeText", studentName+"\n\n")
	oleutil.MustCallMethod(selection, "TypeText", "Дата выполнения: "+date+"\n\n")

	fmt.Print("Имя файла для сохранения (без .docx): ")
	var filename string
	fmt.Scanln(&filename)
	oleutil.MustCallMethod(doc, "SaveAs", filename+".docx")

	fmt.Println("Отчет сохранен как", filename+".docx")
}
