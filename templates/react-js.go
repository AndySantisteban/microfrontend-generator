package templates

var INDEX_HTML_JSX_TEMPLATE = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My App</title>
</head>
<body>
    <div id="root"></div>
    <script type="module" src="/src/main.jsx"></script>
</body>
</html>
`

var MAIN_JSX_TEMPLATE = `
import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
)
`

var APP_JSX_TEMPLATE = `
import React from 'react'

function App() {
  return (
    <div>
      <h1>Hello, World!</h1>
    </div>
  )
}

export default App
`
