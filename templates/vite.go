package templates

var ViteConfigTemplate = `import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import federation from '@originjs/vite-plugin-federation';

export default defineConfig({
  plugins: [
    react(),
    federation({
      name: %s,
      filename: 'remoteEntry.js',
      exposes: {
        './App': './src/App.%s',
      },
      remotes: {
        %s
      },
      shared: %s
    })
  ],
  preview: {
    port: %d, 
    cors: {
      origin: "*"
    },
    proxy: {
      %s
    }
  },
  build: {
    modulePreload: false,
    target: "esnext",
    minify: false,
    cssCodeSplit: false,
    outDir: "dist",
    emptyOutDir: false,
  },
  base: '%s',
});
`
