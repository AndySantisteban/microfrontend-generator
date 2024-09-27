package templates

var MAIN_TSX_TEMPLATE = `import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
)
`

var APP_TSX_TEMPLATE = `
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
