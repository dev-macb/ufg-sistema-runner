package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestAssinarArquivoValido verifica que o comando assinar conclui com sucesso
// quando um arquivo existente é informado.
func TestAssinarArquivoValido(t *testing.T) {
	binario := compilarBinario(t)

	// Cria um arquivo temporário para ser assinado
	arquivo := criarArquivoTemp(t, "documento-*.txt", "conteúdo de teste")

	stdout, stderr, err := executarComando(binario, "assinar", "--arquivo", arquivo)
	if err != nil {
		t.Fatalf("esperava sucesso, mas obteve erro: %v\nstderr: %s", err, stderr)
	}

	if !strings.Contains(stdout, "sucesso") {
		t.Errorf("esperava mensagem de sucesso na saída, mas obteve: %q", stdout)
	}
}

// TestAssinarSemArquivo verifica que o comando assinar falha com mensagem
// adequada quando a flag --arquivo não é informada.
func TestAssinarSemArquivo(t *testing.T) {
	binario := compilarBinario(t)

	_, stderr, err := executarComando(binario, "assinar")
	if err == nil {
		t.Fatal("esperava erro ao omitir --arquivo, mas o comando teve sucesso")
	}

	if !strings.Contains(stderr, "arquivo") {
		t.Errorf("esperava mensagem mencionando 'arquivo' no stderr, mas obteve: %q", stderr)
	}
}

// TestAssinarArquivoInexistente verifica que o comando assinar falha com
// mensagem adequada quando o arquivo informado não existe.
func TestAssinarArquivoInexistente(t *testing.T) {
	binario := compilarBinario(t)

	_, stderr, err := executarComando(binario, "assinar", "--arquivo", "nao-existe.txt")
	if err == nil {
		t.Fatal("esperava erro para arquivo inexistente, mas o comando teve sucesso")
	}

	if !strings.Contains(stderr, "não encontrado") {
		t.Errorf("esperava mensagem 'não encontrado' no stderr, mas obteve: %q", stderr)
	}
}

// TestAssinarDiretorio verifica que o comando assinar falha quando
// o caminho informado aponta para um diretório.
func TestAssinarDiretorio(t *testing.T) {
	binario := compilarBinario(t)

	dir := t.TempDir()

	_, stderr, err := executarComando(binario, "assinar", "--arquivo", dir)
	if err == nil {
		t.Fatal("esperava erro ao passar diretório como arquivo, mas o comando teve sucesso")
	}

	if !strings.Contains(stderr, "diretório") {
		t.Errorf("esperava mensagem mencionando 'diretório' no stderr, mas obteve: %q", stderr)
	}
}

// criarArquivoTemp cria um arquivo temporário com o conteúdo informado
// e o remove automaticamente ao fim do teste.
func criarArquivoTemp(t *testing.T, padrao string, conteudo string) string {
	t.Helper()

	arquivo, err := os.CreateTemp("", padrao)
	if err != nil {
		t.Fatalf("erro ao criar arquivo temporário: %v", err)
	}

	if _, err := arquivo.WriteString(conteudo); err != nil {
		t.Fatalf("erro ao escrever no arquivo temporário: %v", err)
	}
	arquivo.Close()

	t.Cleanup(func() {
		os.Remove(filepath.Clean(arquivo.Name()))
	})

	return arquivo.Name()
}