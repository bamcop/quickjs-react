import React from 'react'
import { renderToString } from 'react-dom/server'
import App from './App'

globalThis.React = React

console.log(renderToString(<App />))
