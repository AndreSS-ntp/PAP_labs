package main

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"log"
	"path/filepath"
)

func main() {
	// Инициализация COM
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	word, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		log.Fatal("Ошибка создания Word:", err)
	}
	defer word.Release()

	wordApp := word.MustQueryInterface(ole.IID_IDispatch)
	defer wordApp.Release()

	oleutil.PutProperty(wordApp, "Visible", true)

	templatePath := filepath.Join("D:", "Study", "PAiPS", "PAP_labs", "lab7", "ТИТУЛЬНИК.docx")

	docs := oleutil.MustGetProperty(wordApp, "Documents").ToIDispatch()
	defer docs.Release()

	doc, err := oleutil.CallMethod(docs, "Open", templatePath)
	if err != nil {
		log.Fatal("Ошибка открытия шаблона:", err)
	}
	docDispatch := doc.ToIDispatch()
	defer docDispatch.Release()

	var labNumber, labName, studentName, date, group, teacher_name string
	fmt.Print("Номер лабораторной работы: ")
	fmt.Scanln(&labNumber)
	fmt.Print("Название работы: ")
	fmt.Scanln(&labName)
	fmt.Print("ФИО студента: ")
	fmt.Scanln(&studentName)
	fmt.Print("Студент группы: ")
	fmt.Scanln(&group)
	fmt.Print("Дата выполнения (дд.мм.гггг): ")
	fmt.Scanln(&date)
	fmt.Print("ФИО преподавателя: ")
	fmt.Scanln(&teacher_name)

	replaceText(docDispatch, "{{LAB_NUMBER}}", labNumber)
	replaceText(docDispatch, "{{LAB_NAME}}", labName)
	replaceText(docDispatch, "{{STUDENT_NAME}}", studentName)
	replaceText(docDispatch, "{{GROUP}}", group)
	replaceText(docDispatch, "{{TEACHER_NAME}}", teacher_name)
	replaceText(docDispatch, "{{DATE}}", date)

	fmt.Print("Имя файла для сохранения (без .docx): ")
	var filename string
	fmt.Scanln(&filename)
	resultPath := "D:\\Study\\PAiPS\\PAP_labs\\lab7\\" + filename + ".docx"
	oleutil.MustCallMethod(docDispatch, "SaveAs", resultPath)

	fmt.Printf("Отчет сохранен как: %s\n", resultPath)
	oleutil.MustCallMethod(wordApp, "Quit")
}

func replaceText(doc *ole.IDispatch, oldText, newText string) {
	content, err := oleutil.GetProperty(doc, "Content")
	if err != nil {
		log.Fatal("Ошибка получения Content:", err)
	}
	contentDispatch := content.ToIDispatch()
	defer contentDispatch.Release()

	find, err := oleutil.GetProperty(contentDispatch, "Find")
	if err != nil {
		log.Fatal("Ошибка получения Find:", err)
	}
	findDispatch := find.ToIDispatch()
	defer findDispatch.Release()

	oleutil.PutProperty(findDispatch, "Text", oldText)
	oleutil.PutProperty(findDispatch, "Replacement.Text", newText)

	if _, err := oleutil.CallMethod(findDispatch, "Execute", oldText, false, false, false, false, false, true, 1, false, newText, 2); err != nil {
		log.Fatal("Ошибка замены текста:", err)
	}
}
