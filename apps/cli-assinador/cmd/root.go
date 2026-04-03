package cmd

import (
	"os"
	"github.com/spf13/cobra"
)


var cmdRaiz = &cobra.Command{
	Use:   "assinador",
	Short: "CLI para assinatura digital de arquivos",
	Long: `assinador é uma ferramenta de linha de comando para
assinar e validar assinaturas digitais de arquivos.

Exemplos:
  assinador assinar --arquivo documento.pdf
  assinador validar --arquivo documento.pdf`,
}


func Executar(versao string) {
	configurarVersao(versao)

	if err := cmdRaiz.Execute(); err != nil {
		os.Exit(1)
	}
}