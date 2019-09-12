package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

// HOMEPATH Directorio inicio
var HOMEPATH = "./"

// PREFIX Prefijo de nombramiento
var PREFIX = "C07_900116413_"

func main() {
	fmt.Println("------------------ ING DEVELOPERS -----------------------")
	fmt.Println("---------------------- RENAME ---------------------------")
	fmt.Println("----------------- INICIO DE PROCESO ---------------------")
	files, err := ioutil.ReadDir(HOMEPATH)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		fmt.Println("ARCHIVO/CARPETA:", file.Name())
		if file.IsDir() {
			err = encode(file.Name())
			if err != nil {
				log.Println(err)
			}
		}
	}
	fmt.Println("------------------- FIN DE PROCESO ---------------------")
	p := ""
	fmt.Scanf("%s", &p)
}

func encode(nameFolder string) error {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", HOMEPATH, nameFolder))
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("------------------ INICIA RENOMBRADO ", nameFolder, " -----------------------")
	c := 0
	for _, file := range files {
		fmt.Println("Archivo encontrado: ", file.Name())
		if filenameWithoutExtension(file.Name()) == "FACT" {
			err = rename(HOMEPATH+nameFolder+"/"+file.Name(), HOMEPATH+nameFolder+"/"+fmt.Sprintf("%s%s%s", PREFIX, nameFolder, path.Ext(file.Name())))
		} else {
			c++
			err = rename(HOMEPATH+nameFolder+"/"+file.Name(), HOMEPATH+nameFolder+"/"+fmt.Sprintf("%s%s_%s%s", PREFIX, nameFolder, strconv.Itoa(c), path.Ext(file.Name())))
		}
		if err != nil {
			log.Println(err)
			return err
		}
	}
	fmt.Println("------------------ FIN RENOMBRADO -----------------------")
	return nil
}
func rename(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func filenameWithoutExtension(name string) string {
	return strings.TrimSuffix(name, path.Ext(name))
}
