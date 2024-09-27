package templates

var INIT_JSON = `
{
  "projectName": "test",
  "port": 3001,
  "base": "/app",
  "repoBase": ".",
  "language": "javascript",
  "spa": false,
  "serveStatics": true,
  "proxy": [
    {
      "path": "/api",
      "target": "http://localhost:8000"
    },
    {
      "path": "/apps",
      "target": "http://localhost:8001"
    }
  ],
  "remotes": {
    "shared" : ["react","react-dom", "react-router-dom", "vue"],
    "exposes": [
      {
        "name": "App1",
        "remoteUrl": "http://localhost:5132",
        "rename": "/app1",
        "dir": "./app1/dist"
      },
      {
        "name": "App2",
        "remoteUrl": "http://localhost:5133",
        "rename": "/app2",
        "dir": "./app2/dist"
        
      },
      {
        "name": "App3",
        "remoteUrl": "http://localhost:5134",
        "rename": "/app3",
        "dir": "./app3/dist"
      }
    ]
  },
  "monorepo": {
    
  }
}
`
