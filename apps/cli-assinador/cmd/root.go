package cmd

import (
	"os"
	"github.com/spf13/cobra"
)


var cmdRaiz = &cobra.Command{
	Use:   "assinador",
	Short: "CLI para assinatura digital de arquivos",
	Long: `╔═╗┌─┐┌─┐┬┌┐┌┌─┐┌┬┐┌─┐┬─┐
╠═╣└─┐└─┐││││├─┤ │││ │├┬┘
╩ ╩└─┘└─┘┴┘└┘┴ ┴─┴┘└─┘┴└─
	
Assinador é uma ferramenta de linha de comando para
assinar e validar assinaturas digitais de arquivos.`,
}


func Executar(versao string) {
	configurarVersao(versao)

	if err := cmdRaiz.Execute(); err != nil {
		os.Exit(1)
	}
}