package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


var versaoApp string

var cmdVersao = &cobra.Command{
	Use:   "versao",
	Short: "Exibe a versão da aplicação",
	Long:  `Exibe a versão atual do assinador instalada no sistema.`,
	Run:   executarVersao,
}


func init() {
	cmdRaiz.AddCommand(cmdVersao)
}


func configurarVersao(v string) {
	versaoApp = v
}


func executarVersao(cmd *cobra.Command, args []string) {
	fmt.Print(`╔═╗┌─┐┌─┐┬┌┐┌┌─┐┌┬┐┌─┐┬─┐
╠═╣└─┐└─┐││││├─┤ │││ │├┬┘
╩ ╩└─┘└─┘┴┘└┘┴ ┴─┴┘└─┘┴└─
`)
	fmt.Printf("Versão: %s\n", versaoApp)
}
