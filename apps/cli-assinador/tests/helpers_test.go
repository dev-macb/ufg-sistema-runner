// Pacote tests contém os testes de integração da CLI cli-assinador.
package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// compilarBinario compila o binário do projeto na raiz e retorna seu caminho.
// O binário é removido automaticamente ao fim do teste via t.Cleanup.
func compilarBinario(t *testing.T) string {
	t.Helper()

	_, arquivoAtual, _, _ := runtime.Caller(1)
	raizProjeto := filepath.Join(filepath.Dir(arquivoAtual), "..")

	nomeBinario := "assinador-teste"
	if isWindows() {
		nomeBinario += ".exe"
	}
	caminhoBinario := filepath.Join(raizProjeto, nomeBinario)

	build := exec.Command("go", "build", "-o", caminhoBinario, ".")
	build.Dir = raizProjeto
	if saida, err := build.CombinedOutput(); err != nil {
		t.Fatalf("erro ao compilar binário: %v\n%s", err, saida)
	}

	t.Cleanup(func() {
		os.Remove(caminhoBinario)
	})

	return caminhoBinario
}

// executarComando executa o binário com os argumentos informados e
// retorna stdout, stderr e o erro, se houver.
func executarComando(binario string, args ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command(binario, args...)

	var outBuf, errBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err = cmd.Run()
	return outBuf.String(), errBuf.String(), err
}

// isWindows retorna true se o sistema operacional for Windows.
func isWindows() bool {
	return os.PathSeparator == '\\'
}