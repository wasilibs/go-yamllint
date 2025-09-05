package runner

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"

	"github.com/wasilibs/go-yamllint/internal/wasm"
)

func Run(_ string, cmdArgs []string, stdin io.Reader, stdout io.Writer, _ io.Writer, cwd string) int {
	ctx := context.Background()

	rtCfg := wazero.NewRuntimeConfig()
	site, err := zip.NewReader(bytes.NewReader(wasm.Site), int64(len(wasm.Site)))
	if err != nil {
		log.Fatal(err)
	}

	uc, err := os.UserCacheDir()
	if err == nil {
		cache, err := wazero.NewCompilationCacheWithDir(filepath.Join(uc, "com.github.wasilibs"))
		if err == nil {
			rtCfg = rtCfg.WithCompilationCache(cache)
		}
	}
	rt := wazero.NewRuntimeWithConfig(ctx, rtCfg)

	wasi_snapshot_preview1.MustInstantiate(ctx, rt)

	args := []string{"python", ".venv/bin/yamllint"}
	args = append(args, cmdArgs...)

	libDir, _ := fs.Sub(site, "lib")
	venvDir, _ := fs.Sub(site, ".venv")

	cfg := wazero.NewModuleConfig().
		WithSysNanosleep().
		WithSysNanotime().
		WithSysWalltime().
		WithStdout(stdout).
		WithStdin(stdin).
		WithRandSource(rand.Reader).
		WithArgs(args...).
		WithFSConfig(wazero.NewFSConfig().
			WithFSMount(libDir, "lib").
			WithFSMount(venvDir, ".venv").
			WithDirMount(cwd, "/")).
		WithEnv("PYTHONPATH", ".venv/lib/python3.13/site-packages").
		WithEnv("PYTHONDONTWRITEBYTECODE", "1")
	for _, env := range os.Environ() {
		k, v, _ := strings.Cut(env, "=")
		cfg = cfg.WithEnv(k, v)
	}

	_, err = rt.InstantiateWithConfig(ctx, wasm.Python, cfg)
	if err != nil {
		sErr := &sys.ExitError{}
		if errors.As(err, &sErr) {
			return int(sErr.ExitCode())
		}
		log.Fatal(err)
	}
	return 0
}
