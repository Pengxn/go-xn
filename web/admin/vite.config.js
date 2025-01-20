import { join } from "node:path";
import { execSync } from "node:child_process";

import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

import pkg from '../../package.json' with { type: 'json' }

export default defineConfig({
    root: "web/admin",
    css: {
        postcss: 'web/admin/postcss.config.js'
    },
    plugins: [react()],
    define: {
        APP_NAME: JSON.stringify(pkg.name),
        APP_VERSION: JSON.stringify(pkg.version),
        COMMIT_HASH: JSON.stringify(execSync('git rev-parse --short HEAD').toString().trim()),
    },
    build: {
        outDir: 'dist',
        emptyOutDir: true,
    },
})
