package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	mergeChunks("mediatest/media/uploads/1", "merged.mp4")
}

// Função para extrair o primeiro número encontrado no nome do arquivo
func extractNumber(fileName string) int {
	re := regexp.MustCompile(`\d+`)
	numStr := re.FindString(filepath.Base(fileName)) //string
	num, err := strconv.Atoi(numStr)                 //retorna um erro ou o numero presente na string
	if err != nil {
		return -1
	}
	return num
}

// Função para mesclar arquivos de fragmentos (chunks) em um único arquivo de saída
func mergeChunks(inputDir, outputFile string) error {
	chunks, err := filepath.Glob(filepath.Join(inputDir, "*.chunk"))
	if err != nil {
		return fmt.Errorf("fail tp find chunks %v", err)
	}

	// Ordena os chunks pelo número extraído do nome do arquivo
	sort.Slice(chunks, func(i, j int) bool {
		return extractNumber(chunks[i]) < extractNumber(chunks[j])
	})

	// Cria o arquivo de saída para mesclar os chunks
	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	// Itera sobre cada chunk e copia seu conteúdo para o arquivo de saída
	for _, chunk := range chunks {
		input, err := os.Open(chunk)
		if err != nil {
			return fmt.Errorf("failed to open chunk: %v", err)
		}

		// Lê o conteúdo do chunk e escreve no arquivo de saída
		_, err = output.ReadFrom(input)
		if err != nil {
			return fmt.Errorf("failed to read chunk %s to merge file: %v", chunk, err)
		}
		input.Close() // Fecha o arquivo chunk atual
	}

	return nil
}
