/*
Copyright © 2024 FelipeMCassiano  felipemcassiano@gmail.com
*/
package cmd

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	var jpegFlag bool
	var pngFlag bool
	var renameFlag string

	rootCmd := &cobra.Command{
		Use:   "konbata",
		Short: "A image format converter",
		RunE: func(cmd *cobra.Command, args []string) error {
			imagePath := args[0]
			imageBytes, err := os.ReadFile(imagePath)

			var newImagePath string
			if err != nil {
				return err
			}

			if jpegFlag {
				jpegBytes, err := toJpeg(imageBytes)
				if err != nil {
					return err
				}
				outputName := renameFlag
				if len(outputName) < 1 {
					outputName = imagePath
				}

				jpegPath := fmt.Sprintf("%s.jpeg", strings.TrimSuffix(outputName, ".png"))

				if err := os.WriteFile(jpegPath, jpegBytes, os.ModePerm); err != nil {
					return err
				}

				newImagePath = jpegPath
			}

			if pngFlag {
				pngBytes, err := toPng(imageBytes)
				if err != nil {
					return err
				}

				outputName := renameFlag
				if len(outputName) > 1 {
					outputName = imagePath
				}

				pngPath := fmt.Sprintf("%s.png", strings.TrimSuffix(outputName, ".jpeg"))

				if err := os.WriteFile(pngPath, pngBytes, os.ModePerm); err != nil {
					return err
				}

				newImagePath = pngPath
			}

			if !pngFlag && !jpegFlag {
				return fmt.Errorf("A format is required")
			}

			fmt.Printf("Image conversion successful! \nSee %s \n", newImagePath)

			return nil
		},
	}

	rootCmd.Flags().BoolVar(&jpegFlag, "jpeg", false, "Convert the output format to JPEG")
	rootCmd.Flags().BoolVar(&pngFlag, "png", false, "Convert the output format to PNG")
	rootCmd.Flags().StringVarP(&renameFlag, "rename", "r", "", "Specify a new name for the converted output")
	return rootCmd
}

func Execute() {
	err := createRootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func toJpeg(imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	if contentType != "image/png" {
		return nil, fmt.Errorf("unable to convert %#v to jpeg", contentType)
	}

	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func toPng(imageBytes []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
