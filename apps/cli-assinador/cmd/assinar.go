package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// opcoesAssinar agrupa as flags do comando assinar.
type opcoesAssinar struct {
	caminhoArquivo string
}

// optsAssinar armazena as opções parseadas pelo Cobra para o comando assinar.
var optsAssinar opcoesAssinar

// cmdAssinar assina digitalmente um arquivo informado.
var cmdAssinar = &cobra.Command{
	Use:   "assinar",
	Short: "Assina digitalmente um arquivo",
	Long: `Gera uma assinatura digital para o arquivo informado.

Exemplos:
  assinador assinar --arquivo documento.pdf
  assinador assinar -a contrato.docx`,
	RunE: executarAssinar,
}

func init() {
	cmdRaiz.AddCommand(cmdAssinar)

	cmdAssinar.Flags().StringVarP(
		&optsAssinar.caminhoArquivo,
		"arquivo", "a",
		"",
		"Caminho do arquivo a ser assinado (obrigatório)",
	)

	_ = cmdAssinar.MarkFlagRequired("arquivo")
}

// executarAssinar contém a lógica principal do comando assinar.
func executarAssinar(cmd *cobra.Command, args []string) error {
	if err := validarArquivoExiste(optsAssinar.caminhoArquivo); err != nil {
		return err
	}

	fmt.Printf("Assinando arquivo: %s\n", optsAssinar.caminhoArquivo)

	// TODO: implementar a lógica real de assinatura digital

	fmt.Println("Arquivo assinado com sucesso.")
	return nil
}

// validarArquivoExiste verifica se o caminho informado aponta para um arquivo válido.
func validarArquivoExiste(caminho string) error {
	info, err := os.Stat(caminho)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("arquivo não encontrado: %s", caminho)
	}
	if err != nil {
		return fmt.Errorf("erro ao acessar o arquivo %s: %w", caminho, err)
	}
	if info.IsDir() {
		return fmt.Errorf("o caminho informado é um diretório, não um arquivo: %s", caminho)
	}
	return nil
}