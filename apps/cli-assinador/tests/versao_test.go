package tests

import (
	"strings"
	"testing"
)

// TestVersaoExibeDev verifica que o subcomando "versao" exibe o valor
// padrão "dev" quando nenhuma injeção via -ldflags é realizada.
func TestVersaoExibeDev(t *testing.T) {
	binario := compilarBinario(t)

	stdout, stderr, err := executarComando(binario, "versao")
	if err != nil {
		t.Fatalf("erro ao executar 'assinador versao': %v\nstderr: %s", err, stderr)
	}

	resultado := strings.TrimSpace(stdout)
	if !strings.Contains(resultado, "dev") {
		t.Errorf("esperava saída contendo %q, mas obteve: %q", "dev", resultado)
	}
}