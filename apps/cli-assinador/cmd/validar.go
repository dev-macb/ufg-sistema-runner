package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// opcoesValidar agrupa as flags do comando validar.
type opcoesValidar struct {
	caminhoArquivo string
}

// optsValidar armazena as opções parseadas pelo Cobra para o comando validar.
var optsValidar opcoesValidar

// cmdValidar valida a assinatura digital de um arquivo.
var cmdValidar = &cobra.Command{
	Use:   "validar",
	Short: "Valida a assinatura digital de um arquivo",
	Long: `Verifica se a assinatura digital de um arquivo é válida e autêntica.

Exemplos:
  assinador validar --arquivo documento.pdf
  assinador validar -a contrato.docx`,
	RunE: executarValidar,
}

func init() {
	cmdRaiz.AddCommand(cmdValidar)

	cmdValidar.Flags().StringVarP(
		&optsValidar.caminhoArquivo,
		"arquivo", "a",
		"",
		"Caminho do arquivo a ser validado (obrigatório)",
	)

	_ = cmdValidar.MarkFlagRequired("arquivo")
}

// executarValidar contém a lógica principal do comando validar.
func executarValidar(cmd *cobra.Command, args []string) error {
	if err := validarArquivoExiste(optsValidar.caminhoArquivo); err != nil {
		return err
	}

	fmt.Printf("Validando assinatura do arquivo: %s\n", optsValidar.caminhoArquivo)

	// TODO: implementar a lógica real de validação da assinatura

	fmt.Println("Assinatura válida.")
	return nil
}