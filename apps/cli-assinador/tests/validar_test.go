package tests

import (
	"strings"
	"testing"
)

// TestValidarArquivoValido verifica que o comando validar conclui com sucesso
// quando um arquivo existente é informado.
func TestValidarArquivoValido(t *testing.T) {
	binario := compilarBinario(t)

	arquivo := criarArquivoTemp(t, "documento-*.txt", "conteúdo de teste")

	stdout, stderr, err := executarComando(binario, "validar", "--arquivo", arquivo)
	if err != nil {
		t.Fatalf("esperava sucesso, mas obteve erro: %v\nstderr: %s", err, stderr)
	}

	if !strings.Contains(stdout, "válida") {
		t.Errorf("esperava mensagem de assinatura válida na saída, mas obteve: %q", stdout)
	}
}

// TestValidarSemArquivo verifica que o comando validar falha com mensagem
// adequada quando a flag --arquivo não é informada.
func TestValidarSemArquivo(t *testing.T) {
	binario := compilarBinario(t)

	_, stderr, err := executarComando(binario, "validar")
	if err == nil {
		t.Fatal("esperava erro ao omitir --arquivo, mas o comando teve sucesso")
	}

	if !strings.Contains(stderr, "arquivo") {
		t.Errorf("esperava mensagem mencionando 'arquivo' no stderr, mas obteve: %q", stderr)
	}
}

// TestValidarArquivoInexistente verifica que o comando validar falha com
// mensagem adequada quando o arquivo informado não existe.
func TestValidarArquivoInexistente(t *testing.T) {
	binario := compilarBinario(t)

	_, stderr, err := executarComando(binario, "validar", "--arquivo", "nao-existe.txt")
	if err == nil {
		t.Fatal("esperava erro para arquivo inexistente, mas o comando teve sucesso")
	}

	if !strings.Contains(stderr, "não encontrado") {
		t.Errorf("esperava mensagem 'não encontrado' no stderr, mas obteve: %q", stderr)
	}
}

// TestValidarDiretorio verifica que o comando validar falha quando
// o caminho informado aponta para um diretório.
func TestValidarDiretorio(t *testing.T) {
	binario := compilarBinario(t)

	dir := t.TempDir()

	_, stderr, err := executarComando(binario, "validar", "--arquivo", dir)
	if err == nil {
		t.Fatal("esperava erro ao passar diretório como arquivo, mas o comando teve sucesso")
	}

	if !strings.Contains(stderr, "diretório") {
		t.Errorf("esperava mensagem mencionando 'diretório' no stderr, mas obteve: %q", stderr)
	}
}