package templates

var PACKAGE_JSON_DEFAULT = `{
  "name": "%s",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "start": "vite preview",
    %s
    "build": "vite build",
    "build:watch": "vite build --watch"
  },
  "dependencies": {
    "@vitejs/plugin-vue": "^5.0.5",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "react-router-dom": "^6.23.1",
    "vue": "^3.4.29",
    "cors": "^2.8.5",
    "dotenv": "^16.4.5",
    "express": "^4.19.2",
    "http-proxy-middleware": "^3.0.0"
  },
  "devDependencies": {
    "@originjs/vite-plugin-federation": "^1.3.5",
    "@typescript-eslint/eslint-plugin": "^7.2.0",
    "@typescript-eslint/parser": "^7.2.0",
    "@vitejs/plugin-react": "^4.2.1",
    "vite": "^5.2.0"
  }
}
  
`
var PACKAGEJSON_WITH_TYPESCRIPT = `
{
  "name": "%s",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "start": "vite preview",
    %s
    "build": "tsc && vite build",
    "build:watch": "tsc && vite build --watch",
    "lint": "eslint . --ext ts,tsx --report-unused-disable-directives --max-warnings 0"
  },
  "dependencies": {
    "@vitejs/plugin-vue": "^5.0.5",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "react-router-dom": "^6.23.1",
    "vue": "^3.4.29",
    %s
  },
  "devDependencies": {
    "@originjs/vite-plugin-federation": "^1.3.5",
    "@types/node": "^20.14.6",
    "@types/react": "^18.2.66",
    "@types/react-dom": "^18.2.22",
    %s
    "@typescript-eslint/eslint-plugin": "^7.2.0",
    "@typescript-eslint/parser": "^7.2.0",
    "@vitejs/plugin-react": "^4.2.1",
    "eslint": "^8.57.0",
    "eslint-plugin-react-hooks": "^4.6.0",
    "eslint-plugin-react-refresh": "^0.4.6",
    "typescript": "^5.2.2",
    "vite": "^5.2.0"
  }
}
`
