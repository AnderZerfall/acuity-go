/// <reference types='vitest' />
import { crx } from '@crxjs/vite-plugin';
import { nxViteTsPaths } from '@nx/vite/plugins/nx-tsconfig-paths.plugin';
import react from '@vitejs/plugin-react';
import path, { resolve } from 'path';
import { defineConfig } from 'vite';
import manifest from './manifest' with { type: 'json' };

export default defineConfig({
  base: './',
  root: __dirname,
  plugins: [
    react(),
    nxViteTsPaths(),
    crx({
      manifest,
      contentScripts: {
        injectCss: true,
        preambleCode: false,
      },
    }),
  ],
  server: {
    port: 4200,
    fs: {
      allow: [path.resolve(__dirname), path.resolve(__dirname, '../../')],
    },
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true,
    rollupOptions: {
      input: {
        index: resolve(__dirname, 'index.html'),
        background: resolve(__dirname, './src/utils/service-worker.ts'),
        content: resolve(__dirname, './src/utils/content-worker.ts'),
      },
      output: {
        entryFileNames: '[name].js',
        chunkFileNames: 'assets/[name].js',
        assetFileNames: 'assets/[name].[ext]',
      },
    },
  },
});
