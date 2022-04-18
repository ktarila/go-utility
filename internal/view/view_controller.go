package view

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ImgConvert struct {
	ImageType string `json:"name" xml:"name" form:"image_type"`
	ImageFile *multipart.FileHeader
}

type TwoFiles struct {
	ToProcess  *os.File
	Proccessed *os.File
}

func changeExtension(infile string, newExtension string) string {
	ext := path.Ext(infile)
	outfile := infile[0:len(infile)-len(ext)] + "." + newExtension

	return outfile
}

func getFilePaths(file *multipart.FileHeader, c *fiber.Ctx) (TwoFiles, error) {
	// Create a temporary file
	var files TwoFiles
	tempFile, err := ioutil.TempFile("", "*"+strip(file.Filename))
	if err != nil {
		return files, err
	}
	defer tempFile.Close()

	// Save the files to disk:
	if err := c.SaveFile(file, tempFile.Name()); err != nil {
		return files, err
	}

	// create processed file name
	processedFile, err := ioutil.TempFile("", "*processed"+strip(file.Filename))
	if err != nil {
		return files, err
	}
	defer processedFile.Close()

	files = TwoFiles{
		ToProcess:  tempFile,
		Proccessed: processedFile,
	}
	return files, err

}

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == '.' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func Home(c *fiber.Ctx) error {
	m := New("index", "Starter project for golang and stimulus js")
	return render(c, m)
}

func ImageConvert(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		ic := new(ImgConvert)

		if err := c.BodyParser(ic); err != nil {
			return err
		}

		if file, err := c.FormFile("fileUpload"); err == nil {
			files, err := getFilePaths(file, c)
			if err != nil {
				log.Fatal(err)
			}

			outfile := changeExtension(files.Proccessed.Name(), ic.ImageType)

			//  reduce file size here
			cmd := exec.Command("convert", files.ToProcess.Name(), outfile)
			err = cmd.Run()

			// send file to client
			newName := changeExtension(file.Filename, ic.ImageType)
			if err == nil {
				c.Set("Content-disposition", "attachment; filename=processed"+newName)
				return c.SendFile(outfile)
			} else {
				log.Fatal("shrink pdf os shell command failed")
			}
		}

		log.Println(ic.ImageType)

	}
	m := New("convert_image", "Image Convert")
	return render(c, m)
}

func ReducePdf(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		if form, err := c.FormFile("fileUpload"); err == nil {

			files, err := getFilePaths(form, c)
			if err != nil {
				log.Fatal(err)
			}

			//  reduce file size here
			cmd := exec.Command("shrinkpdf", files.ToProcess.Name(), files.Proccessed.Name())
			err = cmd.Run()

			// send file to client
			if err == nil {
				c.Set("Content-disposition", "attachment; filename=processed"+form.Filename)
				return c.SendFile(files.Proccessed.Name())
			} else {
				log.Fatal("shrink pdf os shell command failed")
			}
		} else {
			log.Print(err)
		}
	}
	m := New("index", "Reduce Pdf")

	return render(c, m)
}
